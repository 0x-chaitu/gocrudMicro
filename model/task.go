package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Task struct {
	Id      int    `json:"id"`
	Name    string `json:"taskname"`
	Created string `json:"created"`
}

// validate data before its persisited in DB
func (t Task) Validate() error {
	return validation.ValidateStruct(
		&t,
		validation.Field(&t.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&t.Created, validation.Required),
	)
}

// validate data after it has persisted in DB
func (t Task) ValidatePersisted() error {
	return validation.ValidateStruct(
		&t,
		validation.Field(&t.Id, validation.Required, is.Int),
		validation.Field(&t.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&t.Created, validation.Required),
	)
}
