package aclCtrl

import (
	aclDaoModel "DigitalWisdom/service/dao/daoModels/acl"
	aclPostgresDao "DigitalWisdom/service/dao/postgres/acl"
	boAcl "DigitalWisdom/service/internal/model/bo/acl"
	"context"
	"encoding/hex"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.uber.org/dig"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"time"
)

type aclCtrl struct {
	pack   aclCtrlPack
	aclDao aclPostgresDao.AclDao
}

type aclCtrlPack struct {
	dig.In
	PostgresDigitalWisdom *gorm.DB      `name:"postgres_digitalWisdom"`
	RedisDigitalWisdom    *redis.Client `name:"redis_digitalWisdom"`
}

type AclCtrl interface {
	Acl(ctx context.Context, args *boAcl.AclArgs) bool
	CreateAccount(ctx context.Context, args *boAcl.CreateAccountArgs) error
}

func NewAcl(pack aclCtrlPack) AclCtrl {
	return &aclCtrl{
		pack:   pack,
		aclDao: aclPostgresDao.NewAclDAO(pack.PostgresDigitalWisdom),
	}
}

func (ctrl *aclCtrl) Acl(ctx context.Context, args *boAcl.AclArgs) bool {

	storedHexHash, err := ctrl.aclDao.GetAccountPassword(ctx, args.Query.AccountID)
	if err != nil {
		return false
	}

	currentHashedBytes, err := scrypt.Key(
		[]byte(args.Query.Password),
		[]byte(args.Query.AccountID),
		16384, 8, 1, 32,
	)
	if err != nil {
		return false
	}

	currentHexHash := hex.EncodeToString(currentHashedBytes)

	if currentHexHash != storedHexHash {
		return false
	}

	return true
}

func (ctrl *aclCtrl) CreateAccount(ctx context.Context, args *boAcl.CreateAccountArgs) error {
	accountID := uuid.New().String()
	hashedPassword, err := scrypt.Key([]byte(args.Query.Password), []byte(accountID), 16384, 8, 1, 32)

	if err != nil {
		return err
	}

	account := &aclDaoModel.Account{
		AccountID: accountID,
		UserName:  args.Query.UserName,
		Email:     args.Query.Email,
		CreatedAt: time.Now(),
	}

	accountPassword := &aclDaoModel.AccountPassword{
		AccountID:      accountID,
		HashedPassword: hex.EncodeToString(hashedPassword),
		UpdatedAt:      time.Now(),
	}

	if err := ctrl.aclDao.CreateAccount(ctx, account, accountPassword); err != nil {
		return err
	}

	return nil
}
