package aclDaoModel

import (
	"time"
)

type Account struct {
	AccountID      string    `gorm:"primaryKey;column:account_id" json:"account_id"`
	UserName       string    `gorm:"column:user_name" json:"user_name"`
	Email          string    `gorm:"column:email" json:"email"`
	PhoneNumber    string    `gorm:"column:phone_number" json:"phone_number"`
	ValidateStatus int       `gorm:"column:validate_status" json:"validate_status"`
	Sex            string    `gorm:"column:sex" json:"sex,omitempty"`
	BirthDay       time.Time `gorm:"column:birth_day" json:"birth_day,omitempty"`
	IsActive       bool      `gorm:"column:is_active" json:"is_active"`
	LastLoginAt    time.Time `gorm:"column:last_login_at" json:"last_login_at,omitempty"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      time.Time `gorm:"column:deleted_at;index" json:"-"`
}

func (Account) TableName() string {
	return "acl.account"
}

type AccountPassword struct {
	AccountID      string    `gorm:"primaryKey;column:account_id"`
	HashedPassword string    `gorm:"column:hashed_password"`
	Password       string    `gorm:"column:password"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (AccountPassword) TableName() string {
	return "acl.account_password"
}

type Query struct {
	Account
	Password string
}

type ColumnName string

const (
	// Account 相關欄位
	AccountID      ColumnName = "account_id"
	UserName       ColumnName = "user_name"
	Email          ColumnName = "email"
	PhoneNumber    ColumnName = "phone_number"
	ValidateStatus ColumnName = "validate_status"
	Sex            ColumnName = "sex"
	BirthDay       ColumnName = "birth_day"
	IsActive       ColumnName = "is_active"
	LastLoginAt    ColumnName = "last_login_at"

	// AccountPassword 相關欄位
	HashedPassword ColumnName = "hashed_password"

	// 共有/通用時間欄位
	CreatedAt ColumnName = "created_at"
	UpdatedAt ColumnName = "updated_at"
	DeletedAt ColumnName = "deleted_at"
)
