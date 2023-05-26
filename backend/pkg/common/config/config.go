package config

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	MysqlDB     *gorm.DB
	RedisCliet  *redis.Client
	MongoClient *mongo.Client
	once        sync.Once
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
		initMysqlConfig()
		initRedisConfig()
		initMongoDBConfig()
		//initSlogConfig()
		hlog.Info("MySQL、Redis、MongoDB、slog初始化成功")
	})
}

// initMysqlConfig
// @Description: 初始化MySQL配置
func initMysqlConfig() {
	var err error
	MysqlDB, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/weflow?charset=utf8&parseTime=True&loc=Local"),
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
	RedisCliet = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	MongoClient = client
}

// initSlogConfig
// @Description: 初始化slog配置
func initSlogConfig() {
	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
	})
	defer slog.MustFlush()
	logFile := "./log/weflow.log"
	fileHandler, err := handler.NewRotateFileHandler(logFile, rotatefile.EveryDay, handler.WithLogLevels(slog.NormalLevels), handler.WithBackupNum(200))
	if err != nil {
		panic(err)
	}
	slog.PushHandler(fileHandler)

	//slog.Configure(func(logger *slog.SugaredLogger) {
	//	f := logger.Formatter.(*slog.TextFormatter)
	//	f.EnableColor = true
	//})
	//defer slog.MustFlush()
	//logFile := "./log/weflow.log"
	//config := handler.NewEmptyConfig(
	//	handler.WithLogfile(logFile),
	//	handler.WithBuffMode(handler.BuffModeLine),
	//	handler.WithBuffSize(1024*16),
	//	handler.WithCompress(false),
	//	handler.WithMaxSize(rotatefile.OneMByte*10),
	//	handler.WithLogLevels(slog.AllLevels),
	//	handler.WithBuffMode(handler.BuffModeBite),
	//	handler.WithRotateTime(rotatefile.EveryDay),
	//)
	//config.BackupTime = rotatefile.DefaultBackTime
	//config.BackupNum = 100
	//
	//h, err := config.CreateHandler()
	//if err != nil {
	//	panic(err)
	//}
	//slog.PushHandler(h)
}
