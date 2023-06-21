package model

import (
	"github.com/wegoteam/weflow/internal/entity/bo"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/service"
)

// GetModelList
// @Description: 获取模板列表
func GetModelList() []bo.ModelDetailResult {
	modelDetails := service.GetModelList()
	var models = make([]bo.ModelDetailResult, 0)
	if utils.IsEmptySlice(modelDetails) {
		return models
	}
	for _, modelDetail := range modelDetails {
		modelBO := bo.ModelDetailResult{
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
	return models
}

// GetModelGroupList
// @Description: 获取模板组列表
// @return []bo.ModelGroupResult
func GetModelGroupList() []bo.ModelGroupResult {
	return nil
}
