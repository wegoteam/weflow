package model

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/service"
)

// GetModelList
// @Description: 获取模板列表
func GetModelList() []entity.ModelDetailResult {
	modelList := service.GetModelList()
	return modelList
}
