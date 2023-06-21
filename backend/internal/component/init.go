package component

import (
	"github.com/wegoteam/weflow/internal/component/mysql"
	"github.com/wegoteam/weflow/internal/component/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
