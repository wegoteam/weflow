package model

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/weflow/internal/biz/entity/bo"
	"github.com/wegoteam/weflow/internal/consts"
	weflowApi "github.com/wegoteam/weflow/pkg/api"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
)

// GetModelList
// @Description: 获取模板列表
func GetModelList(param *entity.ModelQueryBO) *base.Response {
	modelDetails, err := weflowApi.GetModelList(param)
	var models = make([]bo.ModelDetailResult, 0)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	if utils.IsEmptySlice(modelDetails) {
		return base.OK(models)
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
			CreateTime:   utils.TimeToString(modelDetail.CreateTime),
			CreateUser:   modelDetail.CreateUser,
			UpdateTime:   utils.TimeToString(modelDetail.UpdateTime),
			UpdateUser:   modelDetail.UpdateUser,
		}
		models = append(models, modelBO)
	}
	return base.OK(models)
}

// PageModelList
// @Description: 分页获取模板列表
// @param: param
// @return *base.Response
func PageModelList(param *entity.ModelPageBO) *base.Response {
	pageResult, err := weflowApi.PageModelList(param)
	var models = make([]bo.ModelDetailResult, 0)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	if utils.IsEmptySlice(pageResult.Records) {
		return base.OK(models)
	}
	for _, modelDetail := range pageResult.Records {
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
			CreateTime:   utils.TimeToString(modelDetail.CreateTime),
			CreateUser:   modelDetail.CreateUser,
			UpdateTime:   utils.TimeToString(modelDetail.UpdateTime),
			UpdateUser:   modelDetail.UpdateUser,
		}
		models = append(models, modelBO)
	}
	page := &base.Page[bo.ModelDetailResult]{
		Total:    pageResult.Total,
		Records:  models,
		PageNum:  param.PageNum,
		PageSize: param.PageSize,
	}
	return base.OK(page)
}

// SaveModel
// @Description: 保存模板
// @param: param
// @return *base.Response
func SaveModel(param *entity.ModelSaveBO) *base.Response {
	modelID, err := weflowApi.SaveModel(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	res := bo.ModelSaveResult{
		ModelID: modelID,
	}
	return base.OK(res)
}

// PublishModel
// @Description: 发布模板
// @param: param
// @return *base.Response
func PublishModel(param *entity.ModelSaveBO) *base.Response {
	err := weflowApi.PublishModel(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	return base.OK(true)
}

// InvalidModel
// @Description: 停用模板
// @param: modelID 模板ID
// @return *base.Response
func InvalidModel(modelID string) *base.Response {
	err := weflowApi.InvalidModel(modelID)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	return base.OK(true)
}

// ReleaseModelVersion
// @Description: 上线模板版本
// @param: versionID 版本ID
// @return *base.Response
func ReleaseModelVersion(versionID string) *base.Response {
	err := weflowApi.ReleaseModelVersion(versionID)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	return base.OK(true)
}

// GetModelVersionList
// @Description: 获取模板版本列表
// @param: modelID
// @return *base.Response
func GetModelVersionList(modelID string) *base.Response {
	versionList, err := weflowApi.GetModelVersionList(modelID)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	modelVersions := make([]bo.ModelVersionResult, 0)
	if utils.IsEmptySlice(versionList) {
		return base.OK(modelVersions)
	}
	for _, version := range versionList {
		var modelVersionBO = &bo.ModelVersionResult{
			ID:           version.ID,
			ModelID:      version.ModelID,
			ModelTitle:   version.ModelTitle,
			VersionID:    version.VersionID,
			ProcessDefID: version.ProcessDefID,
			FormDefID:    version.FormDefID,
			UseStatus:    version.UseStatus,
			Remark:       version.Remark,
			CreateTime:   utils.TimeToString(version.CreateTime),
			CreateUser:   version.CreateUser,
			UpdateTime:   utils.TimeToString(version.UpdateTime),
			UpdateUser:   version.UpdateUser,
			NoticeURL:    version.NoticeURL,
			TitleProps:   version.TitleProps,
		}
		modelVersions = append(modelVersions, *modelVersionBO)
	}
	return base.OK(modelVersions)
}

// GetModelGroupList
// @Description: 获取模板组列表
// @return []bo.ModelGroupResult
func GetModelGroupList() *base.Response {
	var modelGroups = make([]bo.ModelGroupResult, 0)
	groups, err := weflowApi.GetModelGroupList()
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	if utils.IsEmptySlice(groups) {
		return base.OK(modelGroups)
	}
	for _, group := range groups {
		var modelGroupBO = &bo.ModelGroupResult{
			ID:         group.ID,
			GroupID:    group.GroupID,
			GroupName:  group.GroupName,
			Remark:     group.Remark,
			CreateTime: utils.TimeToString(group.CreateTime),
			CreateUser: group.CreateUser,
			UpdateTime: utils.TimeToString(group.UpdateTime),
			UpdateUser: group.UpdateUser,
		}
		modelGroups = append(modelGroups, *modelGroupBO)
	}
	return base.OK(modelGroups)
}

// AddModelGroup
// @Description: 添加模板组
// @param: param
// @return bool
func AddModelGroup(param *entity.ModelGroupAddBO) *base.Response {
	err := weflowApi.AddModelGroup(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	hlog.Infof("添加模板组成功,data=%v", param)
	return base.Success()
}

// EditModelGroup
// @Description: 编辑模板组
// @param: param
// @return bool
func EditModelGroup(param *entity.ModelGroupEditBO) *base.Response {
	err := weflowApi.EditModelGroup(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	hlog.Infof("编辑模板组成功,data=%v", param)
	return base.Success()
}

// DelModelGroup
// @Description: 删除模板组
// @param: param
// @return bool
func DelModelGroup(param *entity.ModelGroupDelBO) *base.Response {
	err := weflowApi.DelModelGroup(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	hlog.Infof("删除模板组成功,data=%v", param)
	return base.Success()
}

// GetGroupModelDetails
// @Description: 获取所有组的所有模版
// @param: param
// @return *base.Response
func GetGroupModelDetails(param *entity.GroupModelQueryBO) *base.Response {
	modelGroupsResult, err := weflowApi.GetGroupModelDetails(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	var groupModelDetails = make([]bo.GroupModelDetailsResult, 0)
	if modelGroupsResult == nil {
		return base.OK(groupModelDetails)
	}
	for _, group := range modelGroupsResult {
		models := make([]bo.ModelDetailResult, 0)
		if utils.IsNotEmptySlice(group.Models) {
			for _, model := range group.Models {
				var modelDetailBO = &bo.ModelDetailResult{
					ID:           model.ID,
					ModelID:      model.ModelID,
					ModelTitle:   model.ModelTitle,
					ProcessDefID: model.ProcessDefID,
					FormDefID:    model.FormDefID,
					ModelGroupID: model.ModelGroupID,
					IconURL:      model.IconURL,
					Status:       model.Status,
					Remark:       model.Remark,
					CreateTime:   utils.TimeToString(model.CreateTime),
					CreateUser:   model.CreateUser,
					UpdateTime:   utils.TimeToString(model.UpdateTime),
					UpdateUser:   model.UpdateUser,
				}
				models = append(models, *modelDetailBO)
			}
		}
		var modelGroupBO = &bo.GroupModelDetailsResult{
			ID:         group.ID,
			GroupID:    group.GroupID,
			GroupName:  group.GroupName,
			Remark:     group.Remark,
			CreateTime: utils.TimeToString(group.CreateTime),
			CreateUser: group.CreateUser,
			UpdateTime: utils.TimeToString(group.UpdateTime),
			UpdateUser: group.UpdateUser,
			Models:     models,
		}
		groupModelDetails = append(groupModelDetails, *modelGroupBO)
	}
	return base.OK(groupModelDetails)
}

// GetModelAndVersionInfo
// @Description: 获取模板和版本信息
// @param: modelID 模板ID
// @param: versionID 版本ID
// @return *base.Response
func GetModelAndVersionInfo(modelID, versionID string) *base.Response {
	modelDetail, err := weflowApi.GetModelAndVersionInfo(modelID, versionID)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	modelResult := &bo.ModelAndVersionInfoResult{
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
		FlowContent:  modelDetail.FlowContent,
		FormContent:  modelDetail.FormContent,
	}
	return base.OK(modelResult)
}
