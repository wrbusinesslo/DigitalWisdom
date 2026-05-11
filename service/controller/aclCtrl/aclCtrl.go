package aclCtrl

import (
	aclDaoModel "DigitalWisdom/service/dao/daoModels/acl"
	aclPostgresDao "DigitalWisdom/service/dao/postgres/acl"
	aclRedisDao "DigitalWisdom/service/dao/redisDao/acl"
	boAcl "DigitalWisdom/service/internal/model/bo/acl"
	"DigitalWisdom/service/internal/utils"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/dig"
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
	Get(ctx context.Context, args *boAcl.GetArgs) (*boAcl.GetReply, error)
	GetLogin(ctx context.Context, args *boAcl.GetArgs) (*boAcl.GetLoginReply, error)
	Update(ctx context.Context, args *boAcl.UpdateArgs) error
}

func NewAcl(pack aclCtrlPack) AclCtrl {
	return &aclCtrl{
		pack:   pack,
		aclDao: aclPostgresDao.NewAclDAO(pack.PostgresDigitalWisdom),
	}
}

func (ctrl *aclCtrl) Get(ctx context.Context, args *boAcl.GetArgs) (*boAcl.GetReply, error) {
	// Implement Get method using postgres
	return nil, nil
}
func (ctrl *aclCtrl) GetLogin(ctx context.Context, args *boAcl.GetArgs) (*boAcl.GetLoginReply, error) {
	aclRao := aclRedisDao.New(ctrl.pack.RedisDigitalWisdom)

	session, err := aclRao.Get(ctx, args.User.Username, args.User.Token)
	if err != nil {
		return nil, err
	}
	if session != nil {
		return &boAcl.GetLoginReply{Session: session}, nil
	}

	// Implement GetLogin method using postgres

	token := utils.GenerateToken()

	newSession := &aclDaoModel.UserSession{
		Username: args.User.Username,
		Token:    token,
	}

	if err := aclRao.Set(ctx, newSession, time.Minute*30); err != nil {
		return nil, err
	}

	return &boAcl.GetLoginReply{Session: newSession}, nil
}

func (ctrl *aclCtrl) Update(ctx context.Context, args *boAcl.UpdateArgs) error {
	// Implement Update method using postgres
	return nil
}
