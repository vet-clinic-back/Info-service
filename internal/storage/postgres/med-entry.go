package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
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

func (s *Storage) GetMedEntries(filter models.EntryReqFilter) ([]models.MedicalEntry, error) {
	query := squirrel.Select(
		fmt.Sprintf("%s.id, %s.entry_date, %s.description, %s.disease, %s.vaccinations, %s.recommendation, "+
			"%s.medical_record_id, %s.device_number, %s.veterinarian_id",
			medEntryTable, medEntryTable, medEntryTable, medEntryTable, medEntryTable,
			medEntryTable, medEntryTable, medEntryTable, medEntryTable, // ha ha ha ha LOL
		),
	).
		From(medEntryTable)

	if filter.EntryID != nil {
		query = query.Where(squirrel.Eq{fmt.Sprintf("%s.id", medEntryTable): *filter.EntryID})
	}
	if filter.PetID != nil {
		query = query.Join(fmt.Sprintf("%s ON %s.id = %s.medical_record_id",
			medRecordTable, medRecordTable, medEntryTable)).
			Where(squirrel.Eq{fmt.Sprintf("%s.pet_id", medRecordTable): *filter.PetID})
	}
	if filter.Limit != nil {
		query = query.Limit(uint64(*filter.Limit))
	}
	if filter.Offset != nil {
		query = query.Offset(uint64(*filter.Offset))
	}

	sqlQuery, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			s.log.WithField("sql", sqlQuery).Error(err)
		}
	}(rows)

	var entries []models.MedicalEntry
	for rows.Next() {
		var entry models.MedicalEntry
		err := rows.Scan(&entry.ID, &entry.EntryDate, &entry.Description, &entry.Disease, &entry.Vaccinations,
			&entry.Recommendation, &entry.MedicalRecordID, &entry.DeviceNumber, &entry.VetID)
		if err != nil {
			return []models.MedicalEntry{}, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
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
