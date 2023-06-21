package comm

import (
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/wepkg/snowflake"
)

// GetSnowflake
// @Description: 获取雪花算法唯一ID
func GetSnowflake() *base.Response {
	snowflakeId := snowflake.GetSnowflakeId()
	return base.OK(snowflakeId)
}
