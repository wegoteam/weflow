package component

import (
	"github.com/wegoteam/weflow/internal/dal/mysql"
	"github.com/wegoteam/weflow/internal/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
