package comm

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wegoteam/wepkg/snowflake"
)

// GetSnowflake 获取雪花算法唯一ID
// @Summary 获取雪花算法唯一ID
// @Description 获取雪花算法唯一ID
// @Accept application/json
// @Produce application/json
// @Router /snowflake [get]
func GetSnowflake(c context.Context, ctx *app.RequestContext) {
	snowflakeId := snowflake.GetSnowflakeId()
	ctx.JSON(consts.StatusOK, utils.H{"snowflakeId": snowflakeId})
}
