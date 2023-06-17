package model

import (
	"github.com/wegoteam/weflow/base"
	"github.com/wegoteam/weflow/pkg/service"
)

// GetModelList
// @Description: 获取模板列表
// @param reqCtx
func GetModelList(reqCtx *base.ReqContext) {
	modelList := service.GetModelList()

	reqCtx.OK(modelList)
}
