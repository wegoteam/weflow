package conf

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/kr/pretty"
	wepgkConfig "github.com/wegoteam/wepkg/config"
	"sync"
)

const (
	HERTZ   = "hertz"
	REDIS   = "redis"
	MONGO   = "mongo"
	SWAGGER = "swagger"
)

var (
	conf   *Config
	once   sync.Once
	config *wepgkConfig.Config
)

// Config
// @Description: 配置
type Config struct {
	Hertz   Hertz   `yaml:"hertz" json:"hertz"`
	MySQL   MySQL   `yaml:"mysql" json:"mysql"`
	Redis   Redis   `yaml:"redis" json:"redis"`
	Mongo   Mongo   `yaml:"mongo" json:"mongo"`
	Swagger Swagger `yaml:"swagger" json:"swagger"`
}

// MySQL
// @Description: MySQL配置
//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
type MySQL struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Db       string `yaml:"db" json:"db"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Charset  string `yaml:"charset" json:"charset"`
}

// Redis
// @Description: redis配置
type Redis struct {
	Address  string `yaml:"address" json:"address"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	DB       int    `yaml:"db" json:"db"`
}

// Hertz
// @Description: Hertz配置
type Hertz struct {
	Address         string `yaml:"address" json:"address"`
	BasePath        string `yaml:"basePath" json:"basePath"`
	EnablePprof     bool   `yaml:"enablePprof" json:"enablePprof"`
	EnableGzip      bool   `yaml:"enableGzip" json:"enableGzip"`
	EnableAccessLog bool   `yaml:"enableAccessLog "json:"enableAccessLog"`
	LogLevel        string `yaml:"logLevel" json:"logLevel"`
	LogFileName     string `yaml:"logFileName" json:"logFileName"`
	LogMaxSize      int    `yaml:"logMaxSize" json:"logMaxSize"`
	LogMaxBackups   int    `yaml:"logMaxBackups" json:"logMaxBackups"`
	LogMaxAge       int    `yaml:"logMaxAge" json:"logMaxAge"`
}

// Swagger
// @Description: Swagger文档配置
type Swagger struct {
	Enable      bool     `json:"enable"`
	Version     string   `json:"version"`
	Host        string   `json:"host"`
	BasePath    string   `json:"basePath"`
	Schemes     []string `json:"schemes"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
}

// Mongo
// @Description: Mongo配置
type Mongo struct {
	Address  string `yaml:"address" json:"address"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

// init
// @Description:
func init() {
	once.Do(func() {
		config = wepgkConfig.GetConfig()
		initConf()
	})
}

// GetConf gets configuration instance
func GetConf() *Config {
	return conf
}

// initConf
// @Description: 初始化配置
func initConf() {
	var hertz = &Hertz{}
	err := config.Load(HERTZ, hertz)
	if err != nil {
		panic(err)
	}
	var redis = &Redis{}
	redisErr := config.Load(REDIS, redis)
	if redisErr != nil {
		hlog.Errorf("redis 加载失败 err=%v", redisErr)
	}
	var mongo = &Mongo{}
	mongoErr := config.Load(MONGO, mongo)
	if mongoErr != nil {
		hlog.Errorf("mongo 加载失败 err=%v", mongoErr)
	}

	var swagger = &Swagger{}
	swaggerErr := config.Load(SWAGGER, swagger)
	if swaggerErr != nil {
		hlog.Errorf("swagger 加载失败 err=%v", swaggerErr)
	}
	conf = &Config{
		Hertz:   *hertz,
		Redis:   *redis,
		Mongo:   *mongo,
		Swagger: *swagger,
	}
	pretty.Printf("%+v\n", conf)
}

// LogLevel
// @Description: 获取日志级别
// @return hlog.Level
func LogLevel() hlog.Level {
	level := GetConf().Hertz.LogLevel
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}
