package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/vet-clinic-back/info-service/internal/config"
	"github.com/vet-clinic-back/info-service/internal/logging"
)

type Storage struct {
	log  *logging.Logger
	db   *sql.DB
	psql squirrel.StatementBuilderType
}

func New(log *logging.Logger, cfg *config.DbConfig) *Storage {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("init postgres failed ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("init postgres failed ", err)
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &Storage{
		log:  log,
		db:   db,
		psql: psql,
	}
}

func (s *Storage) Shutdown() error {
	return s.db.Close()
}
