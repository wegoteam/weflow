package api

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/service"
)

// GetModelList
// @Description: 查询模板列表
// @param: param
// @return []entity.ModelDetailResult
// @return error
func GetModelList(param *entity.ModelQueryBO) ([]entity.ModelDetailResult, error) {
	return service.GetModelList(param)
}

// PageModelList
// @Description: 分页查询模板列表
// @param: param
// @return []entity.ModelDetailResult
// @return error
func PageModelList(param *entity.ModelPageBO) (*entity.Page[entity.ModelDetailResult], error) {
	return service.PageModelList(param)
}

// GetModelGroupList
// @Description: 查询模型分组
// @return []entity.ModelGroupResult
func GetModelGroupList() ([]entity.ModelGroupResult, error) {
	return service.GetModelGroupList()
}

// SaveModel
// @Description: 保存模板
// @param: param
// @return *base.Response
func SaveModel(param *entity.ModelSaveBO) (string, error) {
	return service.SaveModel(param)
}

// PublishModel
// @Description: 发布模板
// @param: param
// @return error
func PublishModel(param *entity.ModelSaveBO) error {
	return service.PublishModel(param)
}

// InvalidModel
// @Description: 停用模板
// @param: modelID
// @return error
func InvalidModel(modelID string) error {
	return service.InvalidModel(modelID)
}

// ReleaseModelVersion
// @Description: 上线模板版本
// @param: versionID 版本ID
// @return error
func ReleaseModelVersion(versionID string) error {
	return service.ReleaseModelVersion(versionID)
}

// GetModelVersionList
// @Description: 获取模型版本列表
// @param: modelID
// @return []entity.ModelVersionResult
func GetModelVersionList(modelID string) ([]entity.ModelVersionResult, error) {
	return service.GetModelVersionList(modelID)
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

// GetModelAndVersionInfo
// @Description: 获取模板和版本信息
// @param: modelID 模板ID
// @param: versionID 版本ID
// @return *entity.ModelDetailResult
// @return error
func GetModelAndVersionInfo(modelID, versionID string) (*entity.ModelAndVersionInfoResult, error) {
	return service.GetModelAndVersionInfo(modelID, versionID)
}
