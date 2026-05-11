package acl

import (
	aclDaoModel "DigitalWisdom/service/dao/daoModels/acl"
	"gorm.io/gorm"
)

type AclDAO struct {
	db *gorm.DB
}

// AclDao defines the interface for acl controller
type AclDao interface {
	CreateUser(args *aclDaoModel.User) error
}

func NewAclDAO(db *gorm.DB) *AclDAO {
	return &AclDAO{db: db}
}

// CreateUser creates a new user record in the database
func (dao *AclDAO) CreateUser(args *aclDaoModel.User) error {
	return dao.db.Create(args).Error
}
