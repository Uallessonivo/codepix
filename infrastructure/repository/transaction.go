package repository

import (
	"fmt"
	"github.com/Uallessonivo/codepix/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (t TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := t.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (t TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := t.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (t TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transactionModel model.Transaction

	t.Db.Preload("AccountFrom.Bank").First(&transactionModel, "id = ?", id)

	if transactionModel.ID == "" {
		return nil, fmt.Errorf("transaction not found")
	}

	return &transactionModel, nil
}
