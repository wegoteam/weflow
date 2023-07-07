package config

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
	wepgkConfig "github.com/wegoteam/wepkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

const (
	MySQL = "mysql"
	REDIS = "redis"
	MONGO = "mongo"
)

var (
	MysqlDB     *gorm.DB
	RedisCliet  *redis.Client
	MongoClient *mongo.Client
	once        sync.Once
	config      *wepgkConfig.Config
)

// init
// @Description: 初始化配置
func init() {
	InitConfig()
}

// InitConfig
// @Description: 初始化配置
func InitConfig() {
	once.Do(func() {
		config = wepgkConfig.GetConfig()
		initMysqlConfig()
		initRedisConfig()
		//initMongoDBConfig()
		hlog.Info("MySQL、Redis、MongoDB 初始化成功")
	})
}

// initMysqlConfig
// @Description: 初始化MySQL配置
func initMysqlConfig() {
	var mysqlConfig = &wepgkConfig.MySQL{}
	config.Load(MySQL, mysqlConfig)
	dns := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Db, mysqlConfig.Charset)
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dns),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
}

// initRedisConfig
// @Description: 初始化Redis配置
func initRedisConfig() {
	var redisConfig = &wepgkConfig.Redis{}
	config.Load(REDIS, redisConfig)
	RedisCliet = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	})
	_, rediserr := RedisCliet.Ping(context.Background()).Result()
	if rediserr != nil {
		panic(rediserr)
	}
}

// initMongoDBConfig
// @Description: 初始化MongoDB配置
//mongodb://user:password@localhost:27017/?authSource=admin
func initMongoDBConfig() {
	var mongoConfig = &wepgkConfig.Mongo{}
	config.Load(MONGO, mongoConfig)
	url := fmt.Sprintf("mongodb://%s", mongoConfig.Address)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	MongoClient = client
}
