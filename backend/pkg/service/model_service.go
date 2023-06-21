package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/model"
	"github.com/wegoteam/wepkg/snowflake"
)

// GetModelList
// @Description: 查询模板列表
// @return []entity.ModelDetailResult
// @return error
func GetModelList() ([]entity.ModelDetailResult, error) {
	var models = make([]entity.ModelDetailResult, 0)
	var modelDetails []model.ModelDetail
	err := MysqlDB.Where("").Find(&modelDetails).Error
	if err != nil {
		hlog.Errorf("查询模板列表失败 error: %v", err)
		return models, errors.New("查询模板列表失败")
	}
	if utils.IsEmptySlice(modelDetails) {
		return models, nil
	}
	for _, modelDetail := range modelDetails {
		modelBO := entity.ModelDetailResult{
			ID:           modelDetail.ID,
			ModelID:      modelDetail.ModelID,
			ModelTitle:   modelDetail.ModelTitle,
			ProcessDefID: modelDetail.ProcessDefID,
			FormDefID:    modelDetail.FormDefID,
			ModelGroupID: modelDetail.ModelGroupID,
			IconURL:      modelDetail.IconURL,
			Status:       modelDetail.Status,
			Remark:       modelDetail.Remark,
			CreateTime:   modelDetail.CreateTime,
			CreateUser:   modelDetail.CreateUser,
			UpdateTime:   modelDetail.UpdateTime,
			UpdateUser:   modelDetail.UpdateUser,
		}
		models = append(models, modelBO)
	}

	return models, nil
}

// GetModelVersionList
// @Description: 获取模型版本
// @param modelID
// @param versionID
// @return []entity.ModelVersionResult
func GetModelVersionList(modelID, versionID string) []entity.ModelVersionResult {

	var modelVersions = make([]entity.ModelVersionResult, 0)

	var versionList []model.ModelVersion
	MysqlDB.Where("model_id = ? and version_id = ?", modelID, versionID).Find(&versionList)

	if versionList == nil {
		return modelVersions
	}
	for _, version := range versionList {
		var modelVersionBO = &entity.ModelVersionResult{
			ID:           version.ID,
			ModelID:      version.ModelID,
			ModelTitle:   version.ModelTitle,
			VersionID:    version.VersionID,
			ProcessDefID: version.ProcessDefID,
			FormDefID:    version.FormDefID,
			UseStatus:    version.UseStatus,
			Remark:       version.Remark,
			CreateTime:   version.CreateTime,
			CreateUser:   version.CreateUser,
			UpdateTime:   version.UpdateTime,
			UpdateUser:   version.UpdateUser,
			NoticeURL:    version.NoticeURL,
			TitleProps:   version.TitleProps,
		}
		modelVersions = append(modelVersions, *modelVersionBO)
	}
	return modelVersions
}

// GetModelVersion
// @Description: 获取模型版本
// @param modelID
// @param versionID
// @return []entity.ModelVersionResult
func GetModelVersion(modelID, versionID string) *entity.ModelVersionResult {

	var version = &model.ModelVersion{}
	MysqlDB.Where("model_id = ? and version_id = ?", modelID, versionID).Find(version)

	if version == nil {
		return nil
	}

	var modelVersionBO = &entity.ModelVersionResult{
		ID:           version.ID,
		ModelID:      version.ModelID,
		ModelTitle:   version.ModelTitle,
		VersionID:    version.VersionID,
		ProcessDefID: version.ProcessDefID,
		FormDefID:    version.FormDefID,
		UseStatus:    version.UseStatus,
		Remark:       version.Remark,
		CreateTime:   version.CreateTime,
		CreateUser:   version.CreateUser,
		UpdateTime:   version.UpdateTime,
		UpdateUser:   version.UpdateUser,
		NoticeURL:    version.NoticeURL,
		TitleProps:   version.TitleProps,
	}
	return modelVersionBO
}

// GetEnableModelVersion
// @Description: 获取发布的模型版本
// @param modelID
// @param versionID
// @return []entity.ModelVersionResult
func GetEnableModelVersion(modelID string) *entity.ModelVersionResult {
	var version = &model.ModelVersion{}
	MysqlDB.Where("model_id = ? and use_status = ?", modelID, 1).Find(version)
	if version == nil {
		return nil
	}
	var modelVersionBO = &entity.ModelVersionResult{
		ID:           version.ID,
		ModelID:      version.ModelID,
		ModelTitle:   version.ModelTitle,
		VersionID:    version.VersionID,
		ProcessDefID: version.ProcessDefID,
		FormDefID:    version.FormDefID,
		UseStatus:    version.UseStatus,
		Remark:       version.Remark,
		CreateTime:   version.CreateTime,
		CreateUser:   version.CreateUser,
		UpdateTime:   version.UpdateTime,
		UpdateUser:   version.UpdateUser,
		NoticeURL:    version.NoticeURL,
		TitleProps:   version.TitleProps,
	}
	return modelVersionBO
}

// GetModelGroupList
// @Description: 查询模型分组
// @return []entity.ModelGroupResult
func GetModelGroupList() ([]entity.ModelGroupResult, error) {
	var modelGroups = make([]entity.ModelGroupResult, 0)
	var groups []model.ModelGroup
	err := MysqlDB.Where("").Find(&groups).Error
	if err != nil {
		hlog.Errorf("查询模型分组失败 error:%s", err.Error())
		return modelGroups, errors.New("查询模型分组失败")
	}
	if utils.IsEmptySlice(groups) {
		return modelGroups, nil
	}
	for _, group := range groups {
		var modelGroupBO = &entity.ModelGroupResult{
			ID:         group.ID,
			GroupID:    group.GroupID,
			GroupName:  group.GroupName,
			Remark:     group.Remark,
			CreateTime: group.CreateTime,
			CreateUser: group.CreateUser,
			UpdateTime: group.UpdateTime,
			UpdateUser: group.UpdateUser,
		}
		modelGroups = append(modelGroups, *modelGroupBO)
	}
	return modelGroups, nil
}

// AddModelGroup
// @Description: 添加模型分组
// @param: param
// @return bool
func AddModelGroup(param *entity.ModelGroupAddBO) error {
	modelGroup := &model.ModelGroup{
		GroupID:    snowflake.GetSnowflakeId(),
		GroupName:  param.GroupName,
		Remark:     param.Remark,
		CreateTime: param.CreateTime,
		CreateUser: param.CreateUser,
		UpdateTime: param.UpdateTime,
		UpdateUser: param.UpdateUser,
	}
	err := MysqlDB.Create(modelGroup).Error
	if err != nil {
		hlog.Errorf("添加模型分组失败 error:%s", err.Error())
		return errors.New("添加模型分组失败")
	}
	return nil
}

// EditModelGroup
// @Description: 编辑模型分组
// @param: param
// @return error
func EditModelGroup(param *entity.ModelGroupEditBO) error {
	modelGroup := &model.ModelGroup{
		GroupName:  param.GroupName,
		Remark:     param.Remark,
		UpdateTime: param.UpdateTime,
		UpdateUser: param.UpdateUser,
	}
	err := MysqlDB.Where("group_id = ?", param.GroupID).Updates(modelGroup).Error
	if err != nil {
		hlog.Errorf("编辑模型分组失败 error:%s", err.Error())
		return errors.New("编辑模型分组失败")
	}
	return nil
}

// DelModelGroup
// @Description: 删除模型分组
// @param: param
// @return error
func DelModelGroup(param *entity.ModelGroupDelBO) error {
	err := MysqlDB.Where("group_id = ?", param.GroupID).Delete(&model.ModelGroup{}).Error
	if err != nil {
		hlog.Errorf("删除模型分组失败 error:%s", err.Error())
		return errors.New("删除模型分组失败")
	}
	return nil
}
