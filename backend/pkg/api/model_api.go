package api

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/service"
)

// GetModelList
// @Description: 查询模板列表
// @return []entity.ModelDetailResult
// @return error
func GetModelList() ([]entity.ModelDetailResult, error) {
	return service.GetModelList()
}

// GetModelGroupList
// @Description: 查询模型分组
// @return []entity.ModelGroupResult
func GetModelGroupList() ([]entity.ModelGroupResult, error) {
	return service.GetModelGroupList()
}

// AddModelGroup
// @Description: 添加模型分组
// @param: param
// @return bool
func AddModelGroup(param *entity.ModelGroupAddBO) error {
	return service.AddModelGroup(param)
}

// EditModelGroup
// @Description: 编辑模型分组
// @param: param
// @return error
func EditModelGroup(param *entity.ModelGroupEditBO) error {
	return service.EditModelGroup(param)
}

// DelModelGroup
// @Description: 删除模型分组
// @param: param
// @return error
func DelModelGroup(param *entity.ModelGroupDelBO) error {
	return service.DelModelGroup(param)
}

// GetGroupModelDetails
// @Description: 获取分组模型详情
// @param: param
// @return []entity.GroupModelDetailsResult
// @return error
func GetGroupModelDetails(param *entity.GroupModelQueryBO) ([]entity.GroupModelDetailsResult, error) {
	return service.GetGroupModelDetails(param)
}
