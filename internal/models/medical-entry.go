package models

// CREATE TABLE IF NOT EXISTS medical_entry (
// id INTEGER PRIMARY KEY DEFAULT nextval('medical_entry_id_seq'),
// entry_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// description TEXT,
// disease TEXT,
// vaccinations TEXT,
// recommendation TEXT,
// medical_record_id INTEGER REFERENCES medical_record(id),
// device_number INTEGER REFERENCES device(id)
// );

type MedicalEntry struct {
	ID              uint   `json:"id"`
	EntryDate       string `json:"entry_date"`
	Description     string `json:"description"`
	Disease         string `json:"disease"`
	Vaccinations    string `json:"vaccinations"`
	Recommendation  string `json:"recommendation"`
	MedicalRecordID uint   `json:"medical_record_id"`
	DeviceNumber    uint   `json:"device_number"`
	VetID           uint   `json:"vet_id"`
}
