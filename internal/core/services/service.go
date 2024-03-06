package services

import (
	"log"

	"github.com/josenymad/hexagonal-go-api/internal/ports"
)

type Service struct {
	database ports.Database
}

func NewService(db ports.Database) Service {

	return Service{
		database: db,
	}
}

func (service *Service) SendData(data interface{}) error {
	if err := service.database.PostData(data); err != nil {
		log.Println("Failed to sendData in service")
		return err
	}

	return nil
}
