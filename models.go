package main

// Car mirrors the FE's `Car` type in CarCycleUI/src/types/index.ts.
// JSON tags must stay camelCase and in sync with that file.
type Car struct {
	ID                         string `json:"id"`
	Name                       string `json:"name"`
	Brand                      string `json:"brand"`
	Color                      string `json:"color"`
	LicensePlate               string `json:"licensePlate"`
	InsuranceRenewalDate       string `json:"insuranceRenewalDate"`
	InsuranceProvider          string `json:"insuranceProvider"`
	ActInsuranceRenewalDate    string `json:"actInsuranceRenewalDate"`
	LastGearOilChangeMileage   int    `json:"lastGearOilChangeMileage"`
	LastGearOilChangeDate      string `json:"lastGearOilChangeDate"`
	LastEngineOilChangeMileage int    `json:"lastEngineOilChangeMileage"`
	LastEngineOilChangeDate    string `json:"lastEngineOilChangeDate"`
	LastBrakeOilChangeDate     string `json:"lastBrakeOilChangeDate"`
	LastTireChangeMileage      int    `json:"lastTireChangeMileage"`
	LastTireChangeDate         string `json:"lastTireChangeDate"`
	CreatedAt                  string `json:"createdAt"`
	UpdatedAt                  string `json:"updatedAt"`
}

// CarChangeRecord mirrors the FE's `CarChangeRecord` type.
type CarChangeRecord struct {
	Field    string `json:"field"`
	OldValue string `json:"oldValue"`
	NewValue string `json:"newValue"`
}

// CarHistoryEntry mirrors the FE's `CarHistoryEntry` type.
type CarHistoryEntry struct {
	ID        string            `json:"id"`
	CarID     string            `json:"carId"`
	ChangedAt string            `json:"changedAt"`
	ChangedBy string            `json:"changedBy"`
	Changes   []CarChangeRecord `json:"changes"`
	Note      string            `json:"note,omitempty"`
}

// CarUpdateInput mirrors the FE's `CarUpdateInput` type. Pointers distinguish
// "field not sent" from "field sent as empty string/zero", even though the FE
// form currently always sends every field.
type CarUpdateInput struct {
	Name                       *string `json:"name"`
	Brand                      *string `json:"brand"`
	Color                      *string `json:"color"`
	LicensePlate               *string `json:"licensePlate"`
	InsuranceRenewalDate       *string `json:"insuranceRenewalDate"`
	InsuranceProvider          *string `json:"insuranceProvider"`
	ActInsuranceRenewalDate    *string `json:"actInsuranceRenewalDate"`
	LastGearOilChangeMileage   *int    `json:"lastGearOilChangeMileage"`
	LastGearOilChangeDate      *string `json:"lastGearOilChangeDate"`
	LastEngineOilChangeMileage *int    `json:"lastEngineOilChangeMileage"`
	LastEngineOilChangeDate    *string `json:"lastEngineOilChangeDate"`
	LastBrakeOilChangeDate     *string `json:"lastBrakeOilChangeDate"`
	LastTireChangeMileage      *int    `json:"lastTireChangeMileage"`
	LastTireChangeDate         *string `json:"lastTireChangeDate"`
}
