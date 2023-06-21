package comm

import (
	"github.com/wegoteam/wepkg/snowflake"
)

// GetSnowflake
// @Description: 获取雪花算法唯一ID
func GetSnowflake() string {
	snowflakeId := snowflake.GetSnowflakeId()
	return snowflakeId
}
