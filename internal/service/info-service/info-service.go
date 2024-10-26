package infoservice

import (
	"github.com/vet-clinic-back/info-service/internal/logging"
	"github.com/vet-clinic-back/info-service/internal/storage"
)

type InfoService struct {
	log     *logging.Logger
	storage storage.Info
}

func New(log *logging.Logger, storage storage.Info) *InfoService {
	return &InfoService{log: log, storage: storage}
}
