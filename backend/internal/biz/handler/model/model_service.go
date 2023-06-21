package model

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/weflow/internal/biz/entity/bo"
	"github.com/wegoteam/weflow/internal/consts"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/service"
)

// GetModelList
// @Description: 获取模板列表
func GetModelList() *base.Response {
	modelDetails, err := service.GetModelList()
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
			CreateTime:   modelDetail.CreateTime,
			CreateUser:   modelDetail.CreateUser,
			UpdateTime:   modelDetail.UpdateTime,
			UpdateUser:   modelDetail.UpdateUser,
		}
		models = append(models, modelBO)
	}
	return base.OK(models)
}

// GetModelGroupList
// @Description: 获取模板组列表
// @return []bo.ModelGroupResult
func GetModelGroupList() *base.Response {
	var modelGroups = make([]bo.ModelGroupResult, 0)
	groups, err := service.GetModelGroupList()
	if err != nil {
		base.Fail(consts.ERROR, err.Error())
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
			CreateTime: group.CreateTime,
			CreateUser: group.CreateUser,
			UpdateTime: group.UpdateTime,
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
	err := service.AddModelGroup(param)
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
	err := service.EditModelGroup(param)
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
	err := service.DelModelGroup(param)
	if err != nil {
		return base.Fail(consts.ERROR, err.Error())
	}
	hlog.Infof("删除模板组成功,data=%v", param)
	return base.Success()
}
