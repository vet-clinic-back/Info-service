package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/vet-clinic-back/info-service/internal/models"
)

const petsTable = "pet"
const vetTable = "veterinarian"
const medicalRecordTable = "medical_record"

// CreatePetWithCard creates pet -> creates card. on fail do not create each.
func (s *Storage) CreatePetWithCard(pet models.Pet, ownderID uint, vetID uint) (uint, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}

	// Create pet
	query := fmt.Sprintf(
		"INSERT INTO %s (animal_type, name, gender, age, weight, condition, behavior, research_status) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		petsTable,
	)

	var petID uint
	if err = tx.QueryRow(
		query, pet.AnimalType, pet.Name, pet.Gender, pet.Age, pet.Weight, pet.Condition, pet.Behavior, pet.ResearchStatus,
	).Scan(&petID); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.log.Errorf("failed to rollback transaction: %v", rollbackErr)
		}
		return 0, fmt.Errorf("failed to create pet: %w", err)
	}

	// Create medical record
	query = fmt.Sprintf("INSERT INTO %s "+
		"(veterinarian_id, owner_id, pet_id) "+
		"VALUES ($1, $2, $3)", medicalRecordTable)

	_, err = tx.Exec(query, vetID, ownderID, petID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, fmt.Errorf("failed to insert pet: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return petID, nil
}

func (s *Storage) GetPet(pet models.Pet) (models.Pet, error) {
	log := s.log.WithField("op", "Storage.GetPet")

	stmt := s.psql.Select(
		"id", "animal_type", "name", "gender", "age", "weight", "condition", "behavior", "research_status",
	).From(petsTable)

	if pet.ID != 0 {
		stmt = stmt.Where(squirrel.Eq{"id": pet.ID})
	}
	if pet.AnimalType != "" {
		stmt = stmt.Where(squirrel.Eq{"animal_type": pet.AnimalType})
	}
	if pet.Name != "" {
		stmt = stmt.Where(squirrel.Eq{"name": pet.Name})
	}
	if pet.Gender != "" {
		stmt = stmt.Where(squirrel.Eq{"gender": pet.Gender})
	}
	if pet.Age != 0 {
		stmt = stmt.Where(squirrel.Eq{"age": pet.Age})
	}
	if pet.Weight != 0 {
		stmt = stmt.Where(squirrel.Eq{"weight": pet.Weight})
	}
	if pet.Condition != "" {
		stmt = stmt.Where(squirrel.Eq{"condition": pet.Condition})
	}
	if pet.Behavior != "" {
		stmt = stmt.Where(squirrel.Eq{"behavior": pet.Behavior})
	}
	if pet.ResearchStatus != "" {
		stmt = stmt.Where(squirrel.Eq{"research_status": pet.ResearchStatus})
	}

	query, args, err := stmt.ToSql()
	if err != nil {
		return models.Pet{}, err
	}

	log.Debug("query: ", query, " args: ", args)

	err = s.db.QueryRow(query, args...).Scan(
		&pet.ID,
		&pet.AnimalType,
		&pet.Name,
		&pet.Gender,
		&pet.Age,
		&pet.Weight,
		&pet.Condition,
		&pet.Behavior,
		&pet.ResearchStatus,
	)
	if err != nil {
		return models.Pet{}, err
	}
	return pet, nil
}

func (s *Storage) GetPetsWithOwnerAndVet(filter models.PetReqFilter) ([]models.OutputPetDTO, error) {
	query := squirrel.Select(
		"pet.id", "pet.animal_type", "pet.name", "pet.gender", "pet.age", "pet.weight",
		"pet.condition", "pet.behavior", "pet.research_status",
		"medical_record.owner_id",
		"medical_record.veterinarian_id",
	).
		From(petsTable).
		Join(fmt.Sprintf("%s ON %s.id = %s.pet_id", medicalRecordTable, petsTable, medicalRecordTable)).
		Join(fmt.Sprintf("%s ON %s.owner_id = %s.id", ownersTable, medicalRecordTable, ownersTable)).
		Join(fmt.Sprintf("%s ON %s.veterinarian_id = %s.id", vetTable, medicalRecordTable, vetTable))

	// Apply filters only if they are non-nil
	if filter.PetID != nil {
		query = query.Where(squirrel.Eq{fmt.Sprintf("%s.id", petsTable): *filter.PetID})
	}
	if filter.OwnerID != nil {
		query = query.Where(squirrel.Eq{fmt.Sprintf("%s.owner_id", medicalRecordTable): *filter.OwnerID})
	}
	if filter.VetID != nil {
		query = query.Where(squirrel.Eq{fmt.Sprintf("%s.veterinarian_id", medicalRecordTable): *filter.VetID})
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

	s.log.WithField("op", "Storage.GetPetsWithOwnerAndVet").WithField("sql", sqlQuery).Info("sql")

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

	var pets []models.OutputPetDTO
	for rows.Next() {
		var pet models.OutputPetDTO
		err := rows.Scan(
			&pet.Pet.ID, &pet.Pet.AnimalType, &pet.Pet.Name, &pet.Pet.Gender, &pet.Pet.Age,
			&pet.Pet.Weight, &pet.Pet.Condition, &pet.Pet.Behavior, &pet.Pet.ResearchStatus,
			&pet.OwnerID, &pet.VetID,
		)
		if err != nil {
			return []models.OutputPetDTO{}, err
		}
		pets = append(pets, pet)
	}

	return pets, nil
}

func (s *Storage) UpdatePet(pet models.Pet) (models.Pet, error) {
	log := s.log.WithField("op", "Storage.UpdatePet")

	stmt := s.psql.Update(petsTable).Where(squirrel.Eq{"id": pet.ID})

	if pet.AnimalType != "" {
		stmt = stmt.Set("animal_type", pet.AnimalType)
	}
	if pet.Name != "" {
		stmt = stmt.Set("name", pet.Name)
	}
	if pet.Gender != "" {
		stmt = stmt.Set("gender", pet.Gender)
	}
	if pet.Age != 0 {
		stmt = stmt.Set("age", pet.Age)
	}
	if pet.Weight != 0 {
		stmt = stmt.Set("weight", pet.Weight)
	}
	if pet.Condition != "" {
		stmt = stmt.Set("condition", pet.Condition)
	}
	if pet.Behavior != "" {
		stmt = stmt.Set("behavior", pet.Behavior)
	}
	if pet.ResearchStatus != "" {
		stmt = stmt.Set("research_status", pet.ResearchStatus)
	}

	stmt = stmt.Where(squirrel.Eq{"id": pet.ID})
	query, args, err := stmt.ToSql()
	if err != nil {
		return models.Pet{}, fmt.Errorf("failed to build update query: %w", err)
	}

	log.Debug("query: ", query, " args: ", args)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		return models.Pet{}, fmt.Errorf("failed to update pet: %w", err)
	}

	return s.GetPet(pet)
}

func (s *Storage) DeletePet(id uint) error {
	log := s.log.WithField("op", "Storage.DeletePet")

	stmt := s.psql.Delete(petsTable).Where(squirrel.Eq{"id": id})

	query, args, err := stmt.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build delete query: %w", err)
	}

	log.Debug("query: ", query, " args: ", args)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete pet: %w", err)
	}

	return nil
}
