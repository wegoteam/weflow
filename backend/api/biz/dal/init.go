package dal

import (
	"github.com/wegoteam/weflow/api/biz/dal/mysql"
	"github.com/wegoteam/weflow/api/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
