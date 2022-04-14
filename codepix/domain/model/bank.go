package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base
	Code string `json: "code"`
	Name string `json: "name"`
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

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
