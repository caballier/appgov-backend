package model

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

// Application is the structure of this component
type Application struct {
	ID          uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Name        string    `json:"name" example:"application name"`
	Description string    `json:"description" example:"application description"`
}

// List of error messages
var (
	ErrNameInvalid = errors.New("name is empty")
)

// AddApplication is the structure needed to add an application
type AddApplication struct {
	Name        string `json:"name" example:"application name"`
	Description string `json:"description" example:"application description"`
}

// Validation for AddApplication validate an application before to add it
func (a AddApplication) Validation() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// UpdateApplication is the structure needed to update an application
type UpdateApplication struct {
	Name        string `json:"name" example:"application name"`
	Description string `json:"description" example:"application description"`
}

// Validation for UpdateApplication function that validate for update
func (a UpdateApplication) Validation() error {
	switch {
	case len(a.Name) == 0:
		return ErrNameInvalid
	default:
		return nil
	}
}

// ApplicationsAll list all applications
func ApplicationsAll(q string) ([]Application, error) {
	return nil, nil
}

// ApplicationOne list all applications
func ApplicationOne(id int) (Application, error) {
	return Application{}, nil
}

// Insert add an application
func (a Application) Insert() (int, error) {
	return 0, nil
}

// Delete remove an application
func Delete(id int) error {
	return nil
}

// Update apply changes in an application
func (a Application) Update() error {
	return nil
}
