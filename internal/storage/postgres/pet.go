package postgres

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/vet-clinic-back/info-service/internal/models"
)

const petsTable = "pet"

func (s *Storage) CreatePet(pet models.Pet) (uint, error) {
	query := fmt.Sprintf("INSERT INTO %s (animal_type, name, gender, age, weight, condition, behavior, research_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", petsTable)

	var id uint
	err := s.db.QueryRow(query, pet.AnimalType, pet.Name, pet.Gender, pet.Age, pet.Weight, pet.Condition, pet.Behavior, pet.ResearchStatus).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create pet: %w", err)
	}

	return id, nil
}

func (s *Storage) GetPet(pet models.Pet) (models.Pet, error) {
	log := s.log.WithField("op", "Storage.GetPet")

	stmt := s.psql.Select("id", "animal_type", "name", "gender", "age", "weight", "condition", "behavior", "research_status").From(petsTable)

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

	err = s.db.QueryRow(query, args...).Scan(&pet.ID, &pet.AnimalType, &pet.Name, &pet.Gender, &pet.Age, &pet.Weight, &pet.Condition, &pet.Behavior, &pet.ResearchStatus)
	if err != nil {
		return models.Pet{}, err
	}
	return pet, nil
}

func (s *Storage) GetAllPets() ([]models.Pet, error) {
	log := s.log.WithField("op", "Storage.GetAllPets")

	stmt := s.psql.Select("id", "animal_type", "name", "gender", "age", "weight", "condition", "behavior", "research_status").From(petsTable)

	query, args, err := stmt.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}

	log.Debug("query: ", query, " args: ", args)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute select query: %w", err)
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var pet models.Pet
		err := rows.Scan(&pet.ID, &pet.AnimalType, &pet.Name, &pet.Gender, &pet.Age, &pet.Weight, &pet.Condition, &pet.Behavior, &pet.ResearchStatus)
		if err != nil {
			return nil, fmt.Errorf("failed to scan pet: %w", err)
		}
		pets = append(pets, pet)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %w", err)
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
