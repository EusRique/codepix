package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	ID        string    `json:"id"`
	Code      string    `json: "code"`
	Name      string    `json: "name"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

//MÉTODO
func (bank *Bank) isValid() error {
	return nil
}

func newBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	return &bank, nil
}
