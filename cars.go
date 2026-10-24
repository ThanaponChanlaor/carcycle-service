package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Store struct {
	db *sql.DB
}

var ErrNotFound = errors.New("not found")

const carColumns = `
	id, name, brand, color, license_plate,
	insurance_renewal_date, insurance_provider, act_insurance_renewal_date,
	last_gear_oil_change_mileage, last_gear_oil_change_date,
	last_engine_oil_change_mileage, last_engine_oil_change_date,
	last_brake_oil_change_date, last_tire_change_mileage, last_tire_change_date,
	created_at, updated_at
`

func scanCar(row interface{ Scan(...any) error }) (Car, error) {
	var c Car
	err := row.Scan(
		&c.ID, &c.Name, &c.Brand, &c.Color, &c.LicensePlate,
		&c.InsuranceRenewalDate, &c.InsuranceProvider, &c.ActInsuranceRenewalDate,
		&c.LastGearOilChangeMileage, &c.LastGearOilChangeDate,
		&c.LastEngineOilChangeMileage, &c.LastEngineOilChangeDate,
		&c.LastBrakeOilChangeDate, &c.LastTireChangeMileage, &c.LastTireChangeDate,
		&c.CreatedAt, &c.UpdatedAt,
	)
	return c, err
}

func (s *Store) GetAllCars() ([]Car, error) {
	rows, err := s.db.Query(`SELECT ` + carColumns + ` FROM cars ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := []Car{}
	for rows.Next() {
		c, err := scanCar(rows)
		if err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, rows.Err()
}

func (s *Store) GetCarByID(id string) (Car, error) {
	row := s.db.QueryRow(`SELECT `+carColumns+` FROM cars WHERE id = ?`, id)
	c, err := scanCar(row)
	if errors.Is(err, sql.ErrNoRows) {
		return Car{}, ErrNotFound
	}
	return c, err
}

func (s *Store) GetHistory(carID string) ([]CarHistoryEntry, error) {
	rows, err := s.db.Query(
		`SELECT id, car_id, changed_at, changed_by, changes_json, note
		 FROM car_history WHERE car_id = ? ORDER BY changed_at DESC`,
		carID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := []CarHistoryEntry{}
	for rows.Next() {
		var e CarHistoryEntry
		var changesJSON string
		var note sql.NullString
		if err := rows.Scan(&e.ID, &e.CarID, &e.ChangedAt, &e.ChangedBy, &changesJSON, &note); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(changesJSON), &e.Changes); err != nil {
			return nil, err
		}
		e.Note = note.String
		entries = append(entries, e)
	}
	return entries, rows.Err()
}

func (s *Store) saveCar(c Car) error {
	_, err := s.db.Exec(`
		UPDATE cars SET
			name = ?, brand = ?, color = ?, license_plate = ?,
			insurance_renewal_date = ?, insurance_provider = ?, act_insurance_renewal_date = ?,
			last_gear_oil_change_mileage = ?, last_gear_oil_change_date = ?,
			last_engine_oil_change_mileage = ?, last_engine_oil_change_date = ?,
			last_brake_oil_change_date = ?, last_tire_change_mileage = ?, last_tire_change_date = ?,
			updated_at = ?
		WHERE id = ?`,
		c.Name, c.Brand, c.Color, c.LicensePlate,
		c.InsuranceRenewalDate, c.InsuranceProvider, c.ActInsuranceRenewalDate,
		c.LastGearOilChangeMileage, c.LastGearOilChangeDate,
		c.LastEngineOilChangeMileage, c.LastEngineOilChangeDate,
		c.LastBrakeOilChangeDate, c.LastTireChangeMileage, c.LastTireChangeDate,
		c.UpdatedAt, c.ID,
	)
	return err
}

func (s *Store) insertHistory(e CarHistoryEntry) error {
	changesJSON, err := json.Marshal(e.Changes)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(
		`INSERT INTO car_history (id, car_id, changed_at, changed_by, changes_json, note) VALUES (?, ?, ?, ?, ?, ?)`,
		e.ID, e.CarID, e.ChangedAt, e.ChangedBy, string(changesJSON), e.Note,
	)
	return err
}

// UpdateCar applies whatever fields are present in `in`, diffs them against
// the current row, and — if anything actually changed — persists the update
// and appends a history entry. Mirrors the diff logic in the FE's mock
// (CarCycleUI/src/services/mockData.ts mockUpdateCar) field-for-field, so
// the `field` names in history entries match the FE's fieldLabels map.
func (s *Store) UpdateCar(id string, in CarUpdateInput) (Car, error) {
	current, err := s.GetCarByID(id)
	if err != nil {
		return Car{}, err
	}

	updated := current
	var changes []CarChangeRecord

	if in.Name != nil && *in.Name != updated.Name {
		changes = append(changes, CarChangeRecord{Field: "name", OldValue: updated.Name, NewValue: *in.Name})
		updated.Name = *in.Name
	}
	if in.Brand != nil && *in.Brand != updated.Brand {
		changes = append(changes, CarChangeRecord{Field: "brand", OldValue: updated.Brand, NewValue: *in.Brand})
		updated.Brand = *in.Brand
	}
	if in.Color != nil && *in.Color != updated.Color {
		changes = append(changes, CarChangeRecord{Field: "color", OldValue: updated.Color, NewValue: *in.Color})
		updated.Color = *in.Color
	}
	if in.LicensePlate != nil && *in.LicensePlate != updated.LicensePlate {
		changes = append(changes, CarChangeRecord{Field: "licensePlate", OldValue: updated.LicensePlate, NewValue: *in.LicensePlate})
		updated.LicensePlate = *in.LicensePlate
	}
	if in.InsuranceRenewalDate != nil && *in.InsuranceRenewalDate != updated.InsuranceRenewalDate {
		changes = append(changes, CarChangeRecord{Field: "insuranceRenewalDate", OldValue: updated.InsuranceRenewalDate, NewValue: *in.InsuranceRenewalDate})
		updated.InsuranceRenewalDate = *in.InsuranceRenewalDate
	}
	if in.InsuranceProvider != nil && *in.InsuranceProvider != updated.InsuranceProvider {
		changes = append(changes, CarChangeRecord{Field: "insuranceProvider", OldValue: updated.InsuranceProvider, NewValue: *in.InsuranceProvider})
		updated.InsuranceProvider = *in.InsuranceProvider
	}
	if in.ActInsuranceRenewalDate != nil && *in.ActInsuranceRenewalDate != updated.ActInsuranceRenewalDate {
		changes = append(changes, CarChangeRecord{Field: "actInsuranceRenewalDate", OldValue: updated.ActInsuranceRenewalDate, NewValue: *in.ActInsuranceRenewalDate})
		updated.ActInsuranceRenewalDate = *in.ActInsuranceRenewalDate
	}
	if in.LastGearOilChangeMileage != nil && *in.LastGearOilChangeMileage != updated.LastGearOilChangeMileage {
		changes = append(changes, CarChangeRecord{
			Field:    "lastGearOilChangeMileage",
			OldValue: strconv.Itoa(updated.LastGearOilChangeMileage),
			NewValue: strconv.Itoa(*in.LastGearOilChangeMileage),
		})
		updated.LastGearOilChangeMileage = *in.LastGearOilChangeMileage
	}
	if in.LastGearOilChangeDate != nil && *in.LastGearOilChangeDate != updated.LastGearOilChangeDate {
		changes = append(changes, CarChangeRecord{Field: "lastGearOilChangeDate", OldValue: updated.LastGearOilChangeDate, NewValue: *in.LastGearOilChangeDate})
		updated.LastGearOilChangeDate = *in.LastGearOilChangeDate
	}
	if in.LastEngineOilChangeMileage != nil && *in.LastEngineOilChangeMileage != updated.LastEngineOilChangeMileage {
		changes = append(changes, CarChangeRecord{
			Field:    "lastEngineOilChangeMileage",
			OldValue: strconv.Itoa(updated.LastEngineOilChangeMileage),
			NewValue: strconv.Itoa(*in.LastEngineOilChangeMileage),
		})
		updated.LastEngineOilChangeMileage = *in.LastEngineOilChangeMileage
	}
	if in.LastEngineOilChangeDate != nil && *in.LastEngineOilChangeDate != updated.LastEngineOilChangeDate {
		changes = append(changes, CarChangeRecord{Field: "lastEngineOilChangeDate", OldValue: updated.LastEngineOilChangeDate, NewValue: *in.LastEngineOilChangeDate})
		updated.LastEngineOilChangeDate = *in.LastEngineOilChangeDate
	}
	if in.LastBrakeOilChangeDate != nil && *in.LastBrakeOilChangeDate != updated.LastBrakeOilChangeDate {
		changes = append(changes, CarChangeRecord{Field: "lastBrakeOilChangeDate", OldValue: updated.LastBrakeOilChangeDate, NewValue: *in.LastBrakeOilChangeDate})
		updated.LastBrakeOilChangeDate = *in.LastBrakeOilChangeDate
	}
	if in.LastTireChangeMileage != nil && *in.LastTireChangeMileage != updated.LastTireChangeMileage {
		changes = append(changes, CarChangeRecord{
			Field:    "lastTireChangeMileage",
			OldValue: strconv.Itoa(updated.LastTireChangeMileage),
			NewValue: strconv.Itoa(*in.LastTireChangeMileage),
		})
		updated.LastTireChangeMileage = *in.LastTireChangeMileage
	}
	if in.LastTireChangeDate != nil && *in.LastTireChangeDate != updated.LastTireChangeDate {
		changes = append(changes, CarChangeRecord{Field: "lastTireChangeDate", OldValue: updated.LastTireChangeDate, NewValue: *in.LastTireChangeDate})
		updated.LastTireChangeDate = *in.LastTireChangeDate
	}

	if len(changes) == 0 {
		return updated, nil
	}

	updated.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

	if err := s.saveCar(updated); err != nil {
		return Car{}, err
	}

	entry := CarHistoryEntry{
		ID:        fmt.Sprintf("hist-%d", time.Now().UnixNano()),
		CarID:     id,
		ChangedAt: updated.UpdatedAt,
		ChangedBy: "CarCycle Staff",
		Changes:   changes,
	}
	if err := s.insertHistory(entry); err != nil {
		return Car{}, err
	}

	return updated, nil
}
