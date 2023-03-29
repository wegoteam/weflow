package dal

import (
	"github.com/wego2023/weflow/internal/biz/dal/mysql"
	"github.com/wego2023/weflow/internal/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
