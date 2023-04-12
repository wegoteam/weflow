package config

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	err        error
	RedisCliet *redis.Client
)

func init() {
	DB, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/weflow?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	RedisCliet = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
