package dal

import (
	"hello/biz/dal/mysql"
	"hello/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
