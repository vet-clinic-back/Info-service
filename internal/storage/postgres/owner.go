package postgres

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/vet-clinic-back/info-service/internal/models"
)

const ownersTable = "owner"

func (s *Storage) CreateOwner(owner models.Owner) (uint, error) {
	query := fmt.Sprintf("INSERT INTO %s (full_name, email, phone, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", ownersTable)

	var id uint
	err := s.db.QueryRow(query, owner.FullName, owner.Email, owner.Phone, owner.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create owner: %w", err)
	}

	return id, nil
}

func (s *Storage) GetOwner(owner models.Owner) (models.Owner, error) {
	log := s.log.WithField("op", "Storage.GetOwner")

	stmt := s.psql.Select("id", "full_name", "email", "phone").From(ownersTable)

	if owner.ID != 0 {
		stmt = stmt.Where(squirrel.Eq{"id": owner.ID})
	}
	if owner.Email != "" {
		stmt = stmt.Where(squirrel.Eq{"email": owner.Email})
	}
	if owner.Phone != "" {
		stmt = stmt.Where(squirrel.Eq{"phone": owner.Phone})
	}
	if owner.Password != "" {
		stmt = stmt.Where(squirrel.Eq{"password_hash": owner.Password})
	}

	query, args, err := stmt.ToSql()
	if err != nil {
		return models.Owner{}, err
	}

	log.Debug("query: ", query, " args: ", args)

	err = s.db.QueryRow(query, args...).Scan(&owner.ID, &owner.FullName, &owner.Email, &owner.Phone)
	if err != nil {
		return models.Owner{}, err
	}
	return owner, nil
}

func (s *Storage) GetAllOwners() ([]models.Owner, error) {
	log := s.log.WithField("op", "Storage.GetAllOwners")

	stmt := s.psql.Select("id", "full_name", "email", "phone").From(ownersTable)

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

	var owners []models.Owner
	for rows.Next() {
		var owner models.Owner
		if err := rows.Scan(&owner.ID, &owner.FullName, &owner.Email, &owner.Phone); err != nil {
			return nil, fmt.Errorf("failed to scan owner: %w", err)
		}
		owners = append(owners, owner)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %w", err)
	}

	return owners, nil
}

func (s *Storage) UpdateOwner(owner models.Owner) (models.Owner, error) {
	log := s.log.WithField("op", "Storage.Updateowner")

	stmt := s.psql.Update(ownersTable).Where(squirrel.Eq{"id": owner.ID})

	if owner.Email != "" {
		stmt = stmt.Set("email", owner.Email)
	}
	if owner.Phone != "" {
		stmt = stmt.Set("phone", owner.Phone)
	}
	if owner.Password != "" {
		stmt = stmt.Set("password_hash", owner.Password)
	}

	stmt = stmt.Where(squirrel.Eq{"id": owner.ID})
	query, args, err := stmt.ToSql()
	if err != nil {
		return models.Owner{}, fmt.Errorf("failed to build update query: %w", err)
	}

	log.Debug("query: ", query, " args: ", args)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		return models.Owner{}, fmt.Errorf("failed to update owner: %w", err)
	}

	return s.GetOwner(owner)
}

func (s *Storage) DeleteOwner(id uint) error {
	log := s.log.WithField("op", "Storage.Deleteowner")

	stmt := s.psql.Delete(ownersTable).Where(squirrel.Eq{"id": id})

	query, args, err := stmt.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build delete query: %w", err)
	}

	log.Debug("query: ", query, " args: ", args)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete owner: %w", err)
	}

	return nil
}
