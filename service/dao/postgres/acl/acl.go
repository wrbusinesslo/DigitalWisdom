package acl

import (
	aclDaoModel "DigitalWisdom/service/dao/daoModels/acl"
	"context" // Import context package
	"fmt"
	"gorm.io/gorm"
)

type AclDAO struct {
	db *gorm.DB
}

type AclDao interface {
	CreateAccount(ctx context.Context, account *aclDaoModel.Account, accountPassword *aclDaoModel.AccountPassword) error
	GetAccountPassword(ctx context.Context, accountID string) (hashedPassword string, err error)
}

func NewAclDAO(db *gorm.DB) AclDao {
	return &AclDAO{db: db}
}

func (dao *AclDAO) GetAccountPassword(ctx context.Context, accountID string) (string, error) {
	var record aclDaoModel.AccountPassword

	err := dao.db.WithContext(ctx).
		Select(string(aclDaoModel.HashedPassword)).
		Where(fmt.Sprintf("%s = ?", aclDaoModel.AccountID), accountID).
		First(&record).Error

	if err != nil {
		return "", err
	}

	return record.HashedPassword, nil
}

func (dao *AclDAO) CreateAccount(ctx context.Context, account *aclDaoModel.Account, accountPassword *aclDaoModel.AccountPassword) error {

	tx := dao.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	err := tx.Create(account).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Create(accountPassword).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
