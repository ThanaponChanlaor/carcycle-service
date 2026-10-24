package main

// Ported from CarCycleUI/src/services/mockData.ts so the real backend
// shows the same demo fleet the FE mock did.

var seedCars = []Car{
	{
		ID: "car-001", Name: "Camry", Brand: "Toyota", Color: "White Pearl",
		LicensePlate:         "กข 1234 กรุงเทพมหานคร",
		InsuranceRenewalDate: "2027-03-15", InsuranceProvider: "วิริยะประกันภัย ชั้น 1", ActInsuranceRenewalDate: "2027-01-10",
		LastGearOilChangeMileage: 40000, LastGearOilChangeDate: "2026-02-10",
		LastEngineOilChangeMileage: 48000, LastEngineOilChangeDate: "2026-06-20",
		LastBrakeOilChangeDate: "2026-05-10",
		LastTireChangeMileage:  35000, LastTireChangeDate: "2026-01-20",
		CreatedAt: "2025-05-01T02:00:00Z", UpdatedAt: "2026-06-20T04:30:00Z",
	},
	{
		ID: "car-002", Name: "CR-V", Brand: "Honda", Color: "Modern Steel",
		LicensePlate:         "ขค 5678 นนทบุรี",
		InsuranceRenewalDate: "2026-08-20", InsuranceProvider: "กรุงเทพประกันภัย ชั้น 1", ActInsuranceRenewalDate: "2026-08-05",
		LastGearOilChangeMileage: 55000, LastGearOilChangeDate: "2025-11-05",
		LastEngineOilChangeMileage: 62000, LastEngineOilChangeDate: "2026-05-12",
		LastBrakeOilChangeDate: "2025-09-15",
		LastTireChangeMileage:  50000, LastTireChangeDate: "2025-10-01",
		CreatedAt: "2025-04-18T02:00:00Z", UpdatedAt: "2026-07-20T07:15:00Z",
	},
	{
		ID: "car-003", Name: "Ranger", Brand: "Ford", Color: "Race Red",
		LicensePlate:         "คง 9012 ชลบุรี",
		InsuranceRenewalDate: "2027-06-01", InsuranceProvider: "เมืองไทยประกันภัย ชั้น 1", ActInsuranceRenewalDate: "2027-05-01",
		LastGearOilChangeMileage: 15000, LastGearOilChangeDate: "2026-04-01",
		LastEngineOilChangeMileage: 18000, LastEngineOilChangeDate: "2026-07-01",
		LastBrakeOilChangeDate: "2026-07-05",
		LastTireChangeMileage:  10000, LastTireChangeDate: "2026-05-01",
		CreatedAt: "2026-02-05T02:00:00Z", UpdatedAt: "2026-07-01T02:00:00Z",
	},
	{
		ID: "car-004", Name: "CX-5", Brand: "Mazda", Color: "Soul Red Crystal",
		LicensePlate:         "งจ 3456 เชียงใหม่",
		InsuranceRenewalDate: "2026-09-10", InsuranceProvider: "ทิพยประกันภัย ชั้น 2+", ActInsuranceRenewalDate: "2026-09-01",
		LastGearOilChangeMileage: 78000, LastGearOilChangeDate: "2025-12-15",
		LastEngineOilChangeMileage: 95000, LastEngineOilChangeDate: "2026-06-30",
		LastBrakeOilChangeDate: "2025-08-20",
		LastTireChangeMileage:  70000, LastTireChangeDate: "2025-11-01",
		CreatedAt: "2024-03-10T02:00:00Z", UpdatedAt: "2026-07-25T09:00:00Z",
	},
	{
		ID: "car-005", Name: "Seltos", Brand: "Kia", Color: "Gravity Gray",
		LicensePlate:         "จฉ 7890 ปทุมธานี",
		InsuranceRenewalDate: "2027-02-14", InsuranceProvider: "สินมั่นคงประกันภัย ชั้น 1", ActInsuranceRenewalDate: "2027-01-20",
		LastGearOilChangeMileage: 20000, LastGearOilChangeDate: "2026-03-01",
		LastEngineOilChangeMileage: 24000, LastEngineOilChangeDate: "2026-06-10",
		LastBrakeOilChangeDate: "2026-04-01",
		LastTireChangeMileage:  18000, LastTireChangeDate: "2026-02-15",
		CreatedAt: "2025-06-01T02:00:00Z", UpdatedAt: "2026-07-10T05:00:00Z",
	},
	{
		ID: "car-006", Name: "D-Max", Brand: "Isuzu", Color: "Silver Metallic",
		LicensePlate:         "ฉช 2345 สมุทรปราการ",
		InsuranceRenewalDate: "2027-07-01", InsuranceProvider: "แอกซ่าประกันภัย ชั้น 1", ActInsuranceRenewalDate: "2027-06-15",
		LastGearOilChangeMileage: 60000, LastGearOilChangeDate: "2026-01-20",
		LastEngineOilChangeMileage: 65000, LastEngineOilChangeDate: "2026-07-10",
		LastBrakeOilChangeDate: "2026-06-01",
		LastTireChangeMileage:  55000, LastTireChangeDate: "2025-12-01",
		CreatedAt: "2025-07-15T02:00:00Z", UpdatedAt: "2026-07-15T02:00:00Z",
	},
	{
		ID: "car-007", Name: "Tucson", Brand: "Hyundai", Color: "Amazon Gray",
		LicensePlate:         "ชซ 6789 ขอนแก่น",
		InsuranceRenewalDate: "2026-08-10", InsuranceProvider: "ธนชาตประกันภัย ชั้น 2+", ActInsuranceRenewalDate: "2026-08-01",
		LastGearOilChangeMileage: 42000, LastGearOilChangeDate: "2025-10-05",
		LastEngineOilChangeMileage: 50000, LastEngineOilChangeDate: "2026-04-25",
		LastBrakeOilChangeDate: "2025-10-10",
		LastTireChangeMileage:  38000, LastTireChangeDate: "2025-09-20",
		CreatedAt: "2024-04-02T02:00:00Z", UpdatedAt: "2026-04-25T06:45:00Z",
	},
	{
		ID: "car-008", Name: "C300", Brand: "Mercedes-Benz", Color: "Obsidian Black",
		LicensePlate:         "ซฌ 0123 ภูเก็ต",
		InsuranceRenewalDate: "2027-04-05", InsuranceProvider: "อลิอันซ์ อยุธยา ชั้น 1", ActInsuranceRenewalDate: "2027-03-01",
		LastGearOilChangeMileage: 12000, LastGearOilChangeDate: "2026-03-20",
		LastEngineOilChangeMileage: 14000, LastEngineOilChangeDate: "2026-06-05",
		LastBrakeOilChangeDate: "2026-03-15",
		LastTireChangeMileage:  8000, LastTireChangeDate: "2026-04-10",
		CreatedAt: "2025-06-20T02:00:00Z", UpdatedAt: "2026-06-05T03:00:00Z",
	},
}

var seedHistory = []CarHistoryEntry{
	{
		ID: "hist-001", CarID: "car-001", ChangedAt: "2026-06-20T04:30:00Z", ChangedBy: "ช่างบริการ - สมชาย",
		Changes: []CarChangeRecord{
			{Field: "lastEngineOilChangeMileage", OldValue: "40000", NewValue: "48000"},
			{Field: "lastEngineOilChangeDate", OldValue: "2026-01-15", NewValue: "2026-06-20"},
		},
		Note: "เปลี่ยนน้ำมันเครื่องตามกำหนด",
	},
	{
		ID: "hist-002", CarID: "car-002", ChangedAt: "2026-07-20T07:15:00Z", ChangedBy: "ฝ่ายธุรการ - นภา",
		Changes: []CarChangeRecord{
			{Field: "insuranceRenewalDate", OldValue: "2025-08-20", NewValue: "2026-08-20"},
			{Field: "insuranceProvider", OldValue: "วิริยะประกันภัย ชั้น 1", NewValue: "กรุงเทพประกันภัย ชั้น 1"},
		},
		Note: "เปลี่ยนบริษัทประกันเมื่อต่ออายุ",
	},
	{
		ID: "hist-003", CarID: "car-002", ChangedAt: "2025-11-05T01:00:00Z", ChangedBy: "ช่างบริการ - วิชัย",
		Changes: []CarChangeRecord{
			{Field: "lastGearOilChangeMileage", OldValue: "48000", NewValue: "55000"},
			{Field: "lastGearOilChangeDate", OldValue: "2025-05-01", NewValue: "2025-11-05"},
		},
	},
	{
		ID: "hist-004", CarID: "car-004", ChangedAt: "2026-07-25T09:00:00Z", ChangedBy: "ฝ่ายธุรการ - นภา",
		Changes: []CarChangeRecord{
			{Field: "actInsuranceRenewalDate", OldValue: "2025-09-01", NewValue: "2026-09-01"},
		},
	},
	{
		ID: "hist-005", CarID: "car-007", ChangedAt: "2026-04-25T06:45:00Z", ChangedBy: "ช่างบริการ - สมชาย",
		Changes: []CarChangeRecord{
			{Field: "lastEngineOilChangeMileage", OldValue: "44000", NewValue: "50000"},
			{Field: "lastEngineOilChangeDate", OldValue: "2025-11-01", NewValue: "2026-04-25"},
		},
	},
	{
		ID: "hist-006", CarID: "car-005", ChangedAt: "2026-07-10T05:00:00Z", ChangedBy: "ฝ่ายธุรการ - นภา",
		Changes: []CarChangeRecord{
			{Field: "insuranceRenewalDate", OldValue: "2026-02-14", NewValue: "2027-02-14"},
		},
	},
}
