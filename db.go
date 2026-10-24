package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "modernc.org/sqlite"
)

func mustOpenDB(path string) *sql.DB {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatalf("open database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("connect to database: %v", err)
	}
	if err := migrate(db); err != nil {
		log.Fatalf("migrate database: %v", err)
	}
	if err := seedIfEmpty(db); err != nil {
		log.Fatalf("seed database: %v", err)
	}
	return db
}

func migrate(db *sql.DB) error {
	// Two separate Exec calls rather than one multi-statement string —
	// safer across sql drivers than relying on multi-statement support.
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS cars (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			brand TEXT NOT NULL,
			color TEXT NOT NULL,
			license_plate TEXT NOT NULL,
			insurance_renewal_date TEXT NOT NULL,
			insurance_provider TEXT NOT NULL,
			act_insurance_renewal_date TEXT NOT NULL,
			last_gear_oil_change_mileage INTEGER NOT NULL,
			last_gear_oil_change_date TEXT NOT NULL,
			last_engine_oil_change_mileage INTEGER NOT NULL,
			last_engine_oil_change_date TEXT NOT NULL,
			last_brake_oil_change_date TEXT NOT NULL,
			last_tire_change_mileage INTEGER NOT NULL,
			last_tire_change_date TEXT NOT NULL,
			created_at TEXT NOT NULL,
			updated_at TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS car_history (
			id TEXT PRIMARY KEY,
			car_id TEXT NOT NULL REFERENCES cars(id),
			changed_at TEXT NOT NULL,
			changed_by TEXT NOT NULL,
			changes_json TEXT NOT NULL,
			note TEXT
		)
	`)
	return err
}

// seedIfEmpty loads the same demo fleet used by the FE's mock data
// (CarCycleUI/src/services/mockData.ts) so switching VITE_USE_MOCK_DATA
// from true to false doesn't change what you see.
func seedIfEmpty(db *sql.DB) error {
	var count int
	if err := db.QueryRow(`SELECT COUNT(*) FROM cars`).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	for _, c := range seedCars {
		_, err := db.Exec(`
			INSERT INTO cars (
				id, name, brand, color, license_plate,
				insurance_renewal_date, insurance_provider, act_insurance_renewal_date,
				last_gear_oil_change_mileage, last_gear_oil_change_date,
				last_engine_oil_change_mileage, last_engine_oil_change_date,
				last_brake_oil_change_date, last_tire_change_mileage, last_tire_change_date,
				created_at, updated_at
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			c.ID, c.Name, c.Brand, c.Color, c.LicensePlate,
			c.InsuranceRenewalDate, c.InsuranceProvider, c.ActInsuranceRenewalDate,
			c.LastGearOilChangeMileage, c.LastGearOilChangeDate,
			c.LastEngineOilChangeMileage, c.LastEngineOilChangeDate,
			c.LastBrakeOilChangeDate, c.LastTireChangeMileage, c.LastTireChangeDate,
			c.CreatedAt, c.UpdatedAt,
		)
		if err != nil {
			return err
		}
	}

	for _, h := range seedHistory {
		changesJSON, err := json.Marshal(h.Changes)
		if err != nil {
			return err
		}
		_, err = db.Exec(
			`INSERT INTO car_history (id, car_id, changed_at, changed_by, changes_json, note) VALUES (?, ?, ?, ?, ?, ?)`,
			h.ID, h.CarID, h.ChangedAt, h.ChangedBy, string(changesJSON), h.Note,
		)
		if err != nil {
			return err
		}
	}

	log.Printf("seeded %d cars and %d history entries", len(seedCars), len(seedHistory))
	return nil
}
