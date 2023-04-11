package dal

import (
	"github.com/wegoteam/weflow/internal/biz/dal/mysql"
	"github.com/wegoteam/weflow/internal/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
