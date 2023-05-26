package service

import (
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/model"
)

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
