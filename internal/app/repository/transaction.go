package repository

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/database"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	AddTransaction(transaction payload.TransactionPayload) (database.TransactionEntity, error)
}

type TransactionRepository struct {
	opt Option
}

func NewTransactionRepository(opt Option) ITransactionRepository {
	return &TransactionRepository{
		opt: opt,
	}
}

func (tr *TransactionRepository) AddTransaction(transaction payload.TransactionPayload) (database.TransactionEntity, error) {
	transactionEntity := database.TransactionEntity{
		UserID:     uint(transaction.UserID),
		TotalPrice: transaction.TotalPrice,
	}

	err := tr.opt.DbMysql.Transaction(func(tx *gorm.DB) error {
		tx.Model(&database.TransactionEntity{}).Create(&transactionEntity)

		for _, cartItem := range transaction.CartItems {
			product := database.ProductEntity{}
			tx.Model(&database.ProductEntity{}).Where("id = ?", cartItem.ProductID).First(&product)
			if product.Stock < cartItem.Quantity {
				return commons.ErrOutOfStock
			}

			product.Stock -= cartItem.Quantity
			tx.Model(&database.ProductEntity{}).Where("id = ?", cartItem.ProductID).Update("stock", product.Stock)
			item := &database.TransactionDetailEntity{
				TransactionID: uint(transactionEntity.ID),
				ProductID:     uint(cartItem.ProductID),
				Quantity:      cartItem.Quantity,
				Price:         cartItem.Price,
			}
			tx.Model(&database.TransactionDetailEntity{}).Create(item)

			// TODO: remove need?
			transactionEntity.ItemDetail = append(transactionEntity.ItemDetail, *item)

		}

		return nil
	})
	if err != nil {
		return transactionEntity, commons.ErrAddTransaction
	}

	return transactionEntity, nil
}
