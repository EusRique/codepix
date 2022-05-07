package repository

import (
	"github.com/EusRique/codepix/domain/model"
	"github.com/jinzhu/gorm"
)

type PixKeyRepositoryDb struct {
	DB *gorm.DB
}

func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := r.DB.Create(bank).Error
	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := r.DB.Create(account).Error
	if err != nil {
		return err
	}

	return nil
}
