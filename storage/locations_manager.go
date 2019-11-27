package storage

import (
	"fmt"

	"github.com/mariacastro96/go_quiz/locations"
)

type LocationsRepo interface {
	Insert(data locations.Location) error
	GetByID(id string) (locations.Location, error)
}

type LocationsManager struct {
	PostgresRepo LocationsRepo
	FileRepo     LocationsRepo
}

func (m LocationsManager) Save(location locations.Location) error {
	if postgresErr := m.PostgresRepo.Insert(location); postgresErr != nil {
		if fileErr := m.FileRepo.Insert(location); fileErr != nil {
			return fmt.Errorf("postgres error: %s. file error: %s", postgresErr.Error(), fileErr.Error())
		}
	}
	return nil
}

func (m LocationsManager) Find(id string) (locations.Location, error) {
	loc, postgresErr := m.PostgresRepo.GetByID(id)
	if postgresErr != nil {
		loc, fileErr := m.PostgresRepo.GetByID(id)
		if fileErr != nil {
			return locations.Location{}, fmt.Errorf("postgres error: %s. file error: %s", postgresErr.Error(), fileErr.Error())
		}
		return loc, nil
	}
	return loc, nil
}
