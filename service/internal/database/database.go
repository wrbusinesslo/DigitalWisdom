package database

import (
	"DigitalWisdom/service/internal/config"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

const (
	redisLocal    = "RichieRedis"
	postgresLocal = "RichiePostgres"
)

type DigitalWisdomOut struct {
	dig.Out
	RedisLocal    *redis.Client `name:"redis_byside"`
	PostgresLocal *gorm.DB      `name:"postgres_byside"`
}

func NewDigitalWisdom(ctx context.Context, dbms config.DatabaseManageSystem) DigitalWisdomOut {
	return DigitalWisdomOut{
		RedisLocal:    newRedis(ctx, redisLocal, dbms.RedisServer[redisLocal]),
		PostgresLocal: newPostgres(postgresLocal, dbms.PostgresServer[postgresLocal]),
	}
}
