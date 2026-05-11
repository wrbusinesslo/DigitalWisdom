package config

import "go.uber.org/dig"

type DigitalWisdomServer struct {
	dig.Out
	DBMS           DatabaseManageSystem `mapstructure:"DatabaseManageSystem"`
	ServiceAddress ServiceAddress       `mapstructure:"service_address"`
}

type ServiceAddress struct {
	DigitalWisdom string `mapstructure:"DigitalWisdom"`
}

type DatabaseManageSystem struct {
	MongoDBSystem  map[string]MongoDB  `mapstructure:"MongoDB"`
	RedisServer    map[string]Redis    `mapstructure:"Redis"`
	PostgresServer map[string]Postgres `mapstructure:"Postgres"`
}

type MongoDB struct {
	Host     string `mapstructure:"Host"`
	Port     string `mapstructure:"Port"`
	User     string `mapstructure:"User"`
	Password string `mapstructure:"Password"`
	Database string `mapstructure:"Database"`
}

type Redis struct {
	Host     string `mapstructure:"Host"`
	Port     string `mapstructure:"Port"`
	Password string `mapstructure:"Password"`
	Database int    `mapstructure:"Database"`
}

type Postgres struct {
	Host     string `mapstructure:"Host"`
	Port     string `mapstructure:"Port"`
	Account  string `mapstructure:"Account"`
	Password string `mapstructure:"Password"`
	Database string `mapstructure:"Database"`
	MaxIdle  int    `mapstructure:"MaxIdle"`
	MaxOpen  int    `mapstructure:"MaxOpen"`
}
