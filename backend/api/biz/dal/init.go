package dal

import (
	"github.com/wego2023/weflow/api/biz/dal/mysql"
	"github.com/wego2023/weflow/api/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
