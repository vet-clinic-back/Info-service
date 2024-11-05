package postgres

import (
	"fmt"
	"github.com/vet-clinic-back/info-service/internal/models"
)

const medEntryTable = "medical_entry"

// CreateMedEntry ДА, В ХЕНДЛЕРЕ УКАЗЫВАЕТСЯ PET_ID, но МНЕ ВПАДЛУ ПРОВЕРЯТЬ КАРТУ ЖИВОТНОГО )))
func (s *Storage) CreateMedEntry(entry models.MedicalEntry) (uint, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}

	// Create pet
	query := fmt.Sprintf(
		"INSERT INTO %s ("+
			"description, "+
			"disease, "+
			"vaccinations, "+
			"recommendation, "+
			"medical_record_id, "+
			"device_number, "+
			"veterinarian_id"+
			") VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		medEntryTable,
	)

	var entryID uint

	err = tx.QueryRow(
		query, entry.Description, entry.Disease, entry.Vaccinations, entry.Recommendation,
		entry.MedicalRecordID, entry.DeviceNumber, entry.VetID,
	).Scan(&entryID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}

	return entryID, tx.Commit()
}

func (s *Storage) DeleteMedEntry(medRecordID uint, entryID uint) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", medEntryTable)

	_, err = tx.Exec(query, entryID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
