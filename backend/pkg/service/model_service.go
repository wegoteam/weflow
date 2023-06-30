package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/common/utils"
	"github.com/wegoteam/weflow/pkg/model"
	wepkgSnowflake "github.com/wegoteam/wepkg/snowflake"
	"gorm.io/gorm"
	"time"
)

// GetModelList
// @Description: 查询模板列表
// @return []entity.ModelDetailResult
// @return error
func GetModelList(param *entity.ModelQueryBO) ([]entity.ModelDetailResult, error) {
	var models = make([]entity.ModelDetailResult, 0)
	var modelDetails []model.ModelDetail
	err := MysqlDB.Model(&model.ModelDetail{}).Scopes(BuildModelQuery(param)).Order("model_detail.create_time desc").Find(&modelDetails).Error
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

// PageModelList
// @Description: 分页查询模板列表
// @param: param
// @return *entity.Page[entity.ModelDetailResult]
// @return error
func PageModelList(param *entity.ModelPageBO) (*entity.Page[entity.ModelDetailResult], error) {
	var models = make([]entity.ModelDetailResult, 0)
	var modelDetails []model.ModelDetail
	var total int64
	countErr := MysqlDB.Model(&model.ModelDetail{}).Scopes(BuildModelPage(param)).Count(&total).Error
	if countErr != nil {
		hlog.Errorf("查询模板列表失败 error: %v", countErr)
		return nil, errors.New("查询模板列表失败")
	}
	if total == 0 {
		return &entity.Page[entity.ModelDetailResult]{
			Total:    total,
			Records:  models,
			PageNum:  param.PageNum,
			PageSize: param.PageSize,
		}, nil
	}
	// 查询模型列表
	modelErr := MysqlDB.Model(&model.ModelDetail{}).Scopes(BuildModelPage(param)).Order("model_detail.create_time desc").Find(&modelDetails).Error
	if modelErr != nil {
		hlog.Errorf("查询模板列表失败 error: %v", modelErr)
		return nil, errors.New("查询模板列表失败")
	}
	if utils.IsEmptySlice(modelDetails) {
		return &entity.Page[entity.ModelDetailResult]{
			Total:    total,
			Records:  models,
			PageNum:  param.PageNum,
			PageSize: param.PageSize,
		}, nil
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
	return &entity.Page[entity.ModelDetailResult]{
		Total:    total,
		Records:  models,
		PageNum:  param.PageNum,
		PageSize: param.PageSize,
	}, nil
}

// SaveModel
// @Description: 保存模板
// @param: param
// @return *base.Response
func SaveModel(param *entity.ModelSaveBO) (string, error) {
	if param.ModelID == "" {
		return saveDraftModel(param)
	}
	now := time.Now()
	modelID := param.ModelID
	// 查询模板是否存在
	var existModel = model.ModelDetail{}
	err := MysqlDB.Model(&model.ModelDetail{}).Where("model_id = ?", modelID).First(&existModel).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		hlog.Errorf("查询模板失败,模版ID输入有误，重新生成新的模板 error: %v", err)
		return saveDraftModel(param)
	}
	//判断模板的状态，为草稿状态的话，直接更新，不存在版本
	if existModel.Status == constant.ModelStatusDraft {
		return editDraftModel(param, &existModel)
	}
	//判断模板的状态，为发布或者停用状态的话，直接更新，新增流程定义和表单定义
	processDefID := wepkgSnowflake.GetSnowflakeId()
	formDefID := wepkgSnowflake.GetSnowflakeId()

	tx := MysqlDB.Begin()
	// 修改模板
	editModel := &model.ModelDetail{
		ModelTitle:   param.Base.ModelName,
		ProcessDefID: processDefID,
		FormDefID:    formDefID,
		ModelGroupID: param.Base.GroupID,
		IconURL:      param.Base.IconURL,
		Remark:       param.Base.Remark,
		Status:       constant.ModelStatusDraft,
		UpdateTime:   now,
		UpdateUser:   param.UserName,
	}
	editModelErr := tx.Model(&model.ModelDetail{}).Where("model_id = ?", param.ModelID).Updates(editModel).Error
	if editModelErr != nil {
		tx.Rollback()
		hlog.Errorf("保存模板失败 error: %v", editModelErr)
		return "", errors.New("保存模板失败")
	}
	// 删除流程定义
	delErr := tx.Where("process_def_id = ? and status = ?", existModel.ProcessDefID, constant.ModelStatusDraft).Delete(&model.ProcessDefInfo{}).Error
	if delErr != nil {
		tx.Rollback()
		hlog.Errorf("删除流程定义失败 error: %v", delErr)
		return "", errors.New("保存模板失败")
	}
	// 删除表单定义
	delFormErr := tx.Where("form_def_id = ? and status = ?", existModel.FormDefID, constant.ModelStatusDraft).Delete(&model.FormDefInfo{}).Error
	if delFormErr != nil {
		tx.Rollback()
		hlog.Errorf("删除表单定义失败 error: %v", delFormErr)
		return "", errors.New("保存模板失败")
	}
	// 添加流程定义
	addProcessDef := &model.ProcessDefInfo{
		ProcessDefID:   processDefID,
		ProcessDefName: param.Base.ModelName,
		Status:         constant.ModelStatusDraft,
		Remark:         param.Base.Remark,
		StructData:     param.FlowContent,
		CreateTime:     now,
		CreateUser:     param.UserName,
		UpdateTime:     now,
		UpdateUser:     param.UserName,
	}
	addProcessDefErr := tx.Model(&model.ProcessDefInfo{}).Save(addProcessDef).Error
	if addProcessDefErr != nil {
		tx.Rollback()
		hlog.Errorf("保存流程定义失败 error: %v", addProcessDefErr)
		return "", errors.New("保存流程失败")
	}
	//添加表单定义
	addFormDef := &model.FormDefInfo{
		FormDefID:   formDefID,
		FormDefName: param.Base.ModelName,
		HTMLContent: param.FormContent,
		HTMLPageURL: "",
		Status:      constant.ModelStatusDraft,
		Remark:      param.Base.Remark,
		CreateTime:  now,
		CreateUser:  param.UserName,
		UpdateTime:  now,
		UpdateUser:  param.UserName,
	}
	addFormDefErr := tx.Model(&model.FormDefInfo{}).Save(addFormDef).Error
	if addFormDefErr != nil {
		tx.Rollback()
		hlog.Errorf("保存表单定义失败 error: %v", addFormDefErr)
		return "", errors.New("保存表单失败")
	}
	tx.Commit()
	return modelID, nil
}

// editDraftModel
// @Description: 编辑草稿模板
// @param: param
// @return error
func editDraftModel(param *entity.ModelSaveBO, existModel *model.ModelDetail) (string, error) {
	now := time.Now()
	modelID := param.ModelID
	processDefID := existModel.ProcessDefID
	formDefID := existModel.FormDefID
	tx := MysqlDB.Begin()
	// 添加模板
	editModel := &model.ModelDetail{
		ModelID:      modelID,
		ModelTitle:   param.Base.ModelName,
		ProcessDefID: processDefID,
		FormDefID:    formDefID,
		ModelGroupID: param.Base.GroupID,
		IconURL:      param.Base.IconURL,
		Status:       constant.ModelStatusDraft,
		Remark:       param.Base.Remark,
		CreateTime:   now,
		CreateUser:   param.UserName,
		UpdateTime:   now,
		UpdateUser:   param.UserName,
	}
	editModelErr := tx.Model(&model.ModelDetail{}).Where("model_id = ?", modelID).Updates(editModel).Error
	if editModelErr != nil {
		tx.Rollback()
		hlog.Errorf("保存模板失败 error: %v", editModelErr)
		return "", errors.New("保存模板失败")
	}
	// 添加流程定义
	editProcessDef := &model.ProcessDefInfo{
		ProcessDefID:   processDefID,
		ProcessDefName: param.Base.ModelName,
		Status:         constant.ModelStatusDraft,
		Remark:         param.Base.Remark,
		StructData:     param.FlowContent,
		CreateTime:     now,
		CreateUser:     param.UserName,
		UpdateTime:     now,
		UpdateUser:     param.UserName,
	}
	editProcessDefErr := tx.Model(&model.ProcessDefInfo{}).Where("process_def_id = ?", processDefID).Updates(editProcessDef).Error
	if editProcessDefErr != nil {
		tx.Rollback()
		hlog.Errorf("保存流程定义失败 error: %v", editProcessDefErr)
		return "", errors.New("保存模板失败")
	}
	//添加表单定义
	editFormDef := &model.FormDefInfo{
		FormDefID:   formDefID,
		FormDefName: param.Base.ModelName,
		HTMLContent: param.FormContent,
		HTMLPageURL: "",
		Status:      constant.ModelStatusDraft,
		Remark:      param.Base.Remark,
		CreateTime:  now,
		CreateUser:  param.UserName,
		UpdateTime:  now,
		UpdateUser:  param.UserName,
	}
	editFormDefErr := tx.Model(&model.FormDefInfo{}).Where("form_def_id = ?", formDefID).Updates(editFormDef).Error
	if editFormDefErr != nil {
		tx.Rollback()
		hlog.Errorf("保存表单定义失败 error: %v", editFormDefErr)
		return "", errors.New("保存模板失败")
	}
	tx.Commit()
	return modelID, nil
}

// saveDraftModel
// @Description: 保存草稿
// @param: param
// @return error
func saveDraftModel(param *entity.ModelSaveBO) (string, error) {
	now := time.Now()
	modelID := wepkgSnowflake.GetSnowflakeId()
	processDefID := wepkgSnowflake.GetSnowflakeId()
	formDefID := wepkgSnowflake.GetSnowflakeId()
	tx := MysqlDB.Begin()
	// 添加模板
	addModel := &model.ModelDetail{
		ModelID:      modelID,
		ModelTitle:   param.Base.ModelName,
		ProcessDefID: processDefID,
		FormDefID:    formDefID,
		ModelGroupID: param.Base.GroupID,
		IconURL:      param.Base.IconURL,
		Status:       constant.ModelStatusDraft,
		Remark:       param.Base.Remark,
		CreateTime:   now,
		CreateUser:   param.UserName,
		UpdateTime:   now,
		UpdateUser:   param.UserName,
	}
	addModelErr := tx.Model(&model.ModelDetail{}).Save(addModel).Error
	if addModelErr != nil {
		tx.Rollback()
		hlog.Errorf("保存模板失败 error: %v", addModelErr)
		return "", errors.New("保存模板失败")
	}
	// 添加流程定义
	addProcessDef := &model.ProcessDefInfo{
		ProcessDefID:   processDefID,
		ProcessDefName: param.Base.ModelName,
		Status:         constant.ModelStatusDraft,
		Remark:         param.Base.Remark,
		StructData:     param.FlowContent,
		CreateTime:     now,
		CreateUser:     param.UserName,
		UpdateTime:     now,
		UpdateUser:     param.UserName,
	}
	addProcessDefErr := tx.Model(&model.ProcessDefInfo{}).Save(addProcessDef).Error
	if addProcessDefErr != nil {
		tx.Rollback()
		hlog.Errorf("保存流程定义失败 error: %v", addProcessDefErr)
		return "", errors.New("保存流程失败")
	}
	//添加表单定义
	addFormDef := &model.FormDefInfo{
		FormDefID:   formDefID,
		FormDefName: param.Base.ModelName,
		HTMLContent: param.FormContent,
		HTMLPageURL: "",
		Status:      constant.ModelStatusDraft,
		Remark:      param.Base.Remark,
		CreateTime:  now,
		CreateUser:  param.UserName,
		UpdateTime:  now,
		UpdateUser:  param.UserName,
	}
	addFormDefErr := tx.Model(&model.FormDefInfo{}).Save(addFormDef).Error
	if addFormDefErr != nil {
		tx.Rollback()
		hlog.Errorf("保存表单定义失败 error: %v", addFormDefErr)
		return "", errors.New("保存表单失败")
	}
	tx.Commit()
	return modelID, nil
}

// PublishModel
// @Description: 发布模型
// @param: param
// @return error
func PublishModel(param *entity.ModelSaveBO) error {
	if param.ModelID == "" {
		return publishFristDraftModel(param)
	}
	modelID := param.ModelID
	// 查询模板是否存在
	var existModel = model.ModelDetail{}
	err := MysqlDB.Model(&model.ModelDetail{}).Where("model_id = ?", modelID).First(&existModel).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		hlog.Errorf("查询模板失败,模版ID输入有误，重新生成新的模板 error: %v", err)
		return publishFristDraftModel(param)
	}
	//判断模板的状态，为草稿状态的话，直接发布，新增版本，修改模板状态，修改流程定义状态，修改表单定义状态
	//为已发布状态的话，新增版本, 修改模板绑定的流程定义和表单定义，修改流程定义状态，修改表单定义状态
	//为失效状态的话，新增版本，修改模板状态为发布状态
	switch existModel.Status {
	case constant.ModelStatusDraft:
		return publishExistDraftModel(param, &existModel)
	case constant.ModelStatusDeployed, constant.ModelStatusInvalid:
		return publishExistDeployedModel(param, &existModel)
	}

	return errors.New("发布模板失败")
}

// publishFristDraftModel
// @Description: 发布草稿模板
// @param: param
// @return error
func publishFristDraftModel(param *entity.ModelSaveBO) error {
	now := time.Now()
	modelID := wepkgSnowflake.GetSnowflakeId()
	processDefID := wepkgSnowflake.GetSnowflakeId()
	formDefID := wepkgSnowflake.GetSnowflakeId()
	versionID := wepkgSnowflake.GetSnowflakeId()
	tx := MysqlDB.Begin()
	// 添加模板
	addModel := &model.ModelDetail{
		ModelID:      modelID,
		ModelTitle:   param.Base.ModelName,
		ProcessDefID: processDefID,
		FormDefID:    formDefID,
		ModelGroupID: param.Base.GroupID,
		IconURL:      param.Base.IconURL,
		Status:       constant.ModelStatusDeployed,
		Remark:       param.Base.Remark,
		CreateTime:   now,
		CreateUser:   param.UserName,
		UpdateTime:   now,
		UpdateUser:   param.UserName,
	}
	addModelErr := tx.Model(&model.ModelDetail{}).Save(addModel).Error
	if addModelErr != nil {
		tx.Rollback()
		hlog.Errorf("保存模板失败 error: %v", addModelErr)
		return errors.New("发布模板失败")
	}
	//添加版本
	addModelVersion := &model.ModelVersion{
		ModelID:      modelID,
		ModelTitle:   param.Base.ModelName,
		VersionID:    versionID,
		ProcessDefID: processDefID,
		FormDefID:    formDefID,
		UseStatus:    constant.ModelVersionUseStatusUse,
		Remark:       param.Base.Remark,
		CreateTime:   now,
		CreateUser:   param.UserName,
		UpdateTime:   now,
		UpdateUser:   param.UserName,
		NoticeURL:    "",
		TitleProps:   "",
	}
	addModelVersionErr := tx.Model(&model.ModelVersion{}).Save(addModelVersion).Error
	if addModelVersionErr != nil {
		tx.Rollback()
		hlog.Errorf("保存模板版本失败 error: %v", addModelVersionErr)
		return errors.New("保存模板版本失败")
	}
	// 添加流程定义
	addProcessDef := &model.ProcessDefInfo{
		ProcessDefID:   processDefID,
		ProcessDefName: param.Base.ModelName,
		Status:         constant.ModelStatusDeployed,
		Remark:         param.Base.Remark,
		StructData:     param.FlowContent,
		CreateTime:     now,
		CreateUser:     param.UserName,
		UpdateTime:     now,
		UpdateUser:     param.UserName,
	}
	addProcessDefErr := tx.Model(&model.ProcessDefInfo{}).Save(addProcessDef).Error
	if addProcessDefErr != nil {
		tx.Rollback()
		hlog.Errorf("保存流程定义失败 error: %v", addProcessDefErr)
		return errors.New("保存流程失败")
	}
	//添加表单定义
	addFormDef := &model.FormDefInfo{
		FormDefID:   formDefID,
		FormDefName: param.Base.ModelName,
		HTMLContent: param.FormContent,
		HTMLPageURL: "",
		Status:      constant.ModelStatusDeployed,
		Remark:      param.Base.Remark,
		CreateTime:  now,
		CreateUser:  param.UserName,
		UpdateTime:  now,
		UpdateUser:  param.UserName,
	}
	addFormDefErr := tx.Model(&model.FormDefInfo{}).Save(addFormDef).Error
	if addFormDefErr != nil {
		tx.Rollback()
		hlog.Errorf("保存表单定义失败 error: %v", addFormDefErr)
		return errors.New("保存表单失败")
	}
	tx.Commit()
	return nil
}

// publishExistDraftModel
// @Description: 发布已存在的模板
// @param: param
// @param: existModel
// @return error
func publishExistDraftModel(param *entity.ModelSaveBO, existModel *model.ModelDetail) error {
	now := time.Now()
	modelID := param.ModelID
	processDefID := existModel.ProcessDefID
	formDefID := existModel.FormDefID
	versionID := wepkgSnowflake.GetSnowflakeId()
	tx := MysqlDB.Begin()
	// 添加模板
	editModel := &model.ModelDetail{
		ModelID:      modelID,
		ModelTitle:   param.Base.ModelName,
		ProcessDefID: processDefID,
		FormDefID:    formDefID,
		ModelGroupID: param.Base.GroupID,
		IconURL:      param.Base.IconURL,
		Status:       constant.ModelStatusDeployed,
		Remark:       param.Base.Remark,
		CreateTime:   now,
		CreateUser:   param.UserName,
		UpdateTime:   now,
		UpdateUser:   param.UserName,
	}
	editModelErr := tx.Model(&model.ModelDetail{}).Where("model_id = ?", modelID).Updates(editModel).Error
	if editModelErr != nil {
		tx.Rollback()
		hlog.Errorf("发布模板版本失败 error: %v", editModelErr)
		return errors.New("发布模板版本失败")
	}
	//添加版本
	addModelVersion := &model.ModelVersion{
		ModelID:      modelID,
		ModelTitle:   param.Base.ModelName,
		VersionID:    versionID,
		ProcessDefID: processDefID,
		FormDefID:    formDefID,
		UseStatus:    constant.ModelVersionUseStatusUse,
		Remark:       param.Base.Remark,
		CreateTime:   now,
		CreateUser:   param.UserName,
		UpdateTime:   now,
		UpdateUser:   param.UserName,
		NoticeURL:    "",
		TitleProps:   "",
	}
	addModelVersionErr := tx.Model(&model.ModelVersion{}).Save(addModelVersion).Error
	if addModelVersionErr != nil {
		tx.Rollback()
		hlog.Errorf("添加版本失败 error: %v", addModelVersionErr)
		return errors.New("发布模板版本失败")
	}
	// 添加流程定义
	editProcessDef := &model.ProcessDefInfo{
		ProcessDefID:   processDefID,
		ProcessDefName: param.Base.ModelName,
		Status:         constant.ModelStatusDeployed,
		Remark:         param.Base.Remark,
		StructData:     param.FlowContent,
		CreateTime:     now,
		CreateUser:     param.UserName,
		UpdateTime:     now,
		UpdateUser:     param.UserName,
	}
	editProcessDefErr := tx.Model(&model.ProcessDefInfo{}).Where("process_def_id = ?", processDefID).Updates(editProcessDef).Error
	if editProcessDefErr != nil {
		tx.Rollback()
		hlog.Errorf("修改流程定义失败 error: %v", editProcessDefErr)
		return errors.New("保存模板失败")
	}
	//添加表单定义
	editFormDef := &model.FormDefInfo{
		FormDefID:   formDefID,
		FormDefName: param.Base.ModelName,
		HTMLContent: param.FormContent,
		HTMLPageURL: "",
		Status:      constant.ModelStatusDeployed,
		Remark:      param.Base.Remark,
		CreateTime:  now,
		CreateUser:  param.UserName,
		UpdateTime:  now,
		UpdateUser:  param.UserName,
	}
	editFormDefErr := tx.Model(&model.FormDefInfo{}).Where("form_def_id = ?", formDefID).Updates(editFormDef).Error
	if editFormDefErr != nil {
		tx.Rollback()
		hlog.Errorf("修改表单定义失败 error: %v", editFormDefErr)
		return errors.New("保存模板失败")
	}
	tx.Commit()
	return nil
}

// publishExistDeployedModel
// @Description: 发布已存在的模板
// @param: param
// @param: existModel
// @return error
func publishExistDeployedModel(param *entity.ModelSaveBO, existModel *model.ModelDetail) error {
	now := time.Now()
	modelID := param.ModelID
	processDefID := wepkgSnowflake.GetSnowflakeId()
	formDefID := wepkgSnowflake.GetSnowflakeId()
	versionID := wepkgSnowflake.GetSnowflakeId()
	tx := MysqlDB.Begin()
	// 添加模板
	editModel := &model.ModelDetail{
		ModelID:      modelID,
		ModelTitle:   param.Base.ModelName,
		ProcessDefID: processDefID,
		FormDefID:    formDefID,
		ModelGroupID: param.Base.GroupID,
		IconURL:      param.Base.IconURL,
		Status:       constant.ModelStatusDeployed,
		Remark:       param.Base.Remark,
		UpdateTime:   now,
		UpdateUser:   param.UserName,
	}
	editModelErr := tx.Model(&model.ModelDetail{}).Where("model_id = ?", modelID).Updates(editModel).Error
	if editModelErr != nil {
		tx.Rollback()
		hlog.Errorf("发布模板版本失败 error: %v", editModelErr)
		return errors.New("发布模板版本失败")
	}
	//修改原来的版本为不使用状态
	editModelVersion := &model.ModelVersion{
		UseStatus:  constant.ModelVersionUseStatusUnUse,
		UpdateTime: now,
		UpdateUser: param.UserName,
	}
	editModelVersionErr := tx.Model(&model.ModelVersion{}).Where("model_id = ? and use_status = ?", modelID, constant.ModelVersionUseStatusUse).Updates(editModelVersion).Error
	if editModelVersionErr != nil {
		tx.Rollback()
		hlog.Errorf("修改版本状态失败 error: %v", editModelVersionErr)
		return errors.New("发布模板版本失败")
	}
	//添加版本
	addModelVersion := &model.ModelVersion{
		ModelID:      modelID,
		ModelTitle:   param.Base.ModelName,
		VersionID:    versionID,
		ProcessDefID: processDefID,
		FormDefID:    formDefID,
		UseStatus:    constant.ModelVersionUseStatusUse,
		Remark:       param.Base.Remark,
		CreateTime:   now,
		CreateUser:   param.UserName,
		UpdateTime:   now,
		UpdateUser:   param.UserName,
		NoticeURL:    "",
		TitleProps:   "",
	}
	addModelVersionErr := tx.Model(&model.ModelVersion{}).Save(addModelVersion).Error
	if addModelVersionErr != nil {
		tx.Rollback()
		hlog.Errorf("添加版本失败 error: %v", addModelVersionErr)
		return errors.New("发布模板版本失败")
	}
	// 添加流程定义
	addProcessDef := &model.ProcessDefInfo{
		ProcessDefID:   processDefID,
		ProcessDefName: param.Base.ModelName,
		Status:         constant.ModelStatusDeployed,
		Remark:         param.Base.Remark,
		StructData:     param.FlowContent,
		CreateTime:     now,
		CreateUser:     param.UserName,
		UpdateTime:     now,
		UpdateUser:     param.UserName,
	}
	addProcessDefErr := tx.Model(&model.ProcessDefInfo{}).Save(addProcessDef).Error
	if addProcessDefErr != nil {
		tx.Rollback()
		hlog.Errorf("修改流程定义失败 error: %v", addProcessDefErr)
		return errors.New("保存模板失败")
	}
	//添加表单定义
	addFormDef := &model.FormDefInfo{
		FormDefID:   formDefID,
		FormDefName: param.Base.ModelName,
		HTMLContent: param.FormContent,
		HTMLPageURL: "",
		Status:      constant.ModelStatusDeployed,
		Remark:      param.Base.Remark,
		CreateTime:  now,
		CreateUser:  param.UserName,
		UpdateTime:  now,
		UpdateUser:  param.UserName,
	}
	addFormDefErr := tx.Model(&model.FormDefInfo{}).Where("form_def_id = ?", formDefID).Updates(addFormDef).Error
	if addFormDefErr != nil {
		tx.Rollback()
		hlog.Errorf("修改表单定义失败 error: %v", addFormDefErr)
		return errors.New("保存模板失败")
	}
	tx.Commit()
	return nil
}

// InvalidModel
// @Description: 停用模板
// @param: modelID 模板ID
// @return error
func InvalidModel(modelID string) error {
	if modelID == "" {
		return errors.New("模版ID输入有误")
	}
	// 查询模板是否存在
	var existModel = model.ModelDetail{}
	err := MysqlDB.Model(&model.ModelDetail{}).Where("model_id = ?", modelID).First(&existModel).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		hlog.Errorf("查询模板失败 error: %v", err)
		return errors.New("模版不存在")
	}
	//发布状态才允许停用
	if existModel.Status != constant.ModelStatusDeployed {
		return errors.New("只有发布状态的模板才能停用")
	}
	now := time.Now()
	// 修改模板状态
	editModel := &model.ModelDetail{
		Status:     constant.ModelStatusInvalid,
		UpdateTime: now,
	}
	editModelErr := MysqlDB.Model(&model.ModelDetail{}).Where("model_id = ?", modelID).Updates(editModel).Error
	if editModelErr != nil {
		hlog.Errorf("停用模板失败 error: %v", editModelErr)
		return errors.New("停用模板失败")
	}
	return nil
}

// ReleaseModelVersion
// @Description: 上线模板版本
// @param: versionID 版本ID
// @return error
func ReleaseModelVersion(versionID string) error {
	if versionID == "" {
		return errors.New("模版版本ID输入有误")
	}
	// 查询模板版本是否存在
	var existModelVersion = model.ModelVersion{}
	err := MysqlDB.Model(&model.ModelVersion{}).Where("version_id = ?", versionID).First(&existModelVersion).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		hlog.Errorf("查询模板版本失败 error: %v", err)
		return errors.New("模版版本不存在")
	}
	if existModelVersion.UseStatus == constant.ModelVersionUseStatusUse {
		return errors.New("模版版本已上线")
	}
	now := time.Now()
	editModelVersion := &model.ModelVersion{
		UseStatus:  constant.ModelVersionUseStatusUnUse,
		UpdateTime: now,
	}
	editModelVersionErr := MysqlDB.Model(&model.ModelVersion{}).Where("model_id = ? and use_status = ?", existModelVersion.ModelID, constant.ModelVersionUseStatusUse).Updates(editModelVersion).Error
	if editModelVersionErr != nil {
		hlog.Errorf("上线模板版本失败 error: %v", editModelVersionErr)
		return errors.New("上线模板版本失败")
	}
	// 修改模板版本状态
	editModelVersionStatus := &model.ModelVersion{
		UseStatus:  constant.ModelVersionUseStatusUse,
		UpdateTime: now,
	}
	editModelVersionStatusErr := MysqlDB.Model(&model.ModelVersion{}).Where("version_id = ?", versionID).Updates(editModelVersionStatus).Error
	if editModelVersionStatusErr != nil {
		hlog.Errorf("上线模板版本失败 error: %v", editModelVersionStatusErr)
		return errors.New("上线模板版本失败")
	}
	return nil
}

// GetModelVersionList
// @Description: 获取模型版本
// @param: modelID
// @param: versionID
// @return []entity.ModelVersionResult
func GetModelVersionList(modelID string) ([]entity.ModelVersionResult, error) {
	var modelVersions = make([]entity.ModelVersionResult, 0)
	var versionList []model.ModelVersion
	err := MysqlDB.Where("model_id = ? ", modelID).Find(&versionList).Error
	if err != nil {
		hlog.Errorf("查询模板版本失败 error: %v", err)
		return modelVersions, errors.New("查询模板版本失败")
	}
	if utils.IsEmptySlice(versionList) {
		return modelVersions, nil
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
	return modelVersions, nil
}

// GetModelVersion
// @Description: 获取模型版本
// @param: modelID
// @param: versionID
// @return []entity.ModelVersionResult
func GetModelVersion(modelID, versionID string) *entity.ModelVersionResult {
	var version = model.ModelVersion{}
	err := MysqlDB.Where("model_id = ? and version_id = ?", modelID, versionID).First(&version).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		hlog.Errorf("查询模板版本失败 error: %v", err)
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
// @param: modelID
// @param: versionID
// @return []entity.ModelVersionResult
func GetEnableModelVersion(modelID string) *entity.ModelVersionResult {
	var version = model.ModelVersion{}
	err := MysqlDB.Where("model_id = ? and use_status = ?", modelID, 1).First(&version).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		hlog.Errorf("查询模板版本失败 error: %v", err)
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
	err := MysqlDB.Where("").Order("id desc").Find(&groups).Error
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
		GroupID:    wepkgSnowflake.GetSnowflakeId(),
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

// GetGroupModelDetails
// @Description: 获取分组模型详情
// @param: param
// @return []entity.GroupModelDetailsResult
// @return error
func GetGroupModelDetails(param *entity.GroupModelQueryBO) ([]entity.GroupModelDetailsResult, error) {
	var groupModelDetails = make([]entity.GroupModelDetailsResult, 0)
	var modelGroups []model.ModelGroup
	groupErr := MysqlDB.Model(&model.ModelGroup{}).Order("create_time desc").Find(&modelGroups).Error
	if groupErr != nil {
		hlog.Errorf("查询模板列表失败 error: %v", groupErr)
		return groupModelDetails, errors.New("查询模板列表失败")
	}
	if utils.IsEmptySlice(modelGroups) {
		return groupModelDetails, nil
	}
	//获取所有模型详情,根据组ID分组
	modelDetailsMap, modelErr := getAllModelDetailsMap(param)
	if modelErr != nil {
		hlog.Errorf("查询模板列表失败 error: %v", modelErr)
		return groupModelDetails, errors.New("查询模板列表失败")
	}
	for _, group := range modelGroups {
		models, ok := modelDetailsMap[group.GroupID]
		if !ok {
			models = make([]entity.ModelDetailResult, 0)
		}
		var modelGroupBO = &entity.GroupModelDetailsResult{
			ID:         group.ID,
			GroupID:    group.GroupID,
			GroupName:  group.GroupName,
			Remark:     group.Remark,
			CreateTime: group.CreateTime,
			CreateUser: group.CreateUser,
			UpdateTime: group.UpdateTime,
			UpdateUser: group.UpdateUser,
			Models:     models,
		}
		groupModelDetails = append(groupModelDetails, *modelGroupBO)
	}
	return groupModelDetails, nil
}

// getAllModelDetailsMap
// @Description: 获取所有模型详情,根据组ID分组
// @return map[string][]entity.ModelDetailResult
// @return error
func getAllModelDetailsMap(param *entity.GroupModelQueryBO) (map[string][]entity.ModelDetailResult, error) {
	//key:组ID val:模板详情
	modelDetailsMap := make(map[string][]entity.ModelDetailResult)
	var modelDetails []model.ModelDetail
	modelErr := MysqlDB.Model(&model.ModelDetail{}).Scopes(BuildGroupModelQuery(param)).Order("model_detail.create_time desc").Find(&modelDetails).Error
	if modelErr != nil {
		hlog.Errorf("查询模板列表失败 error: %v", modelErr)
		return modelDetailsMap, errors.New("查询模板列表失败")
	}
	if utils.IsEmptySlice(modelDetails) {
		return modelDetailsMap, nil
	}
	for _, modelDetail := range modelDetails {
		//是否存在该组
		models, ok := modelDetailsMap[modelDetail.ModelGroupID]
		if !ok {
			models = make([]entity.ModelDetailResult, 0)
		}
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
		modelDetailsMap[modelDetail.ModelGroupID] = models
	}
	return modelDetailsMap, nil
}

// BuildGroupModelQuery
// @Description: 构建模型查询条件
// @param: param
// @return func(db *gorm.DB) *gorm.DB
func BuildGroupModelQuery(param *entity.GroupModelQueryBO) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tx := db
		if utils.IsStrNotBlank(param.ModelName) {
			tx = db.Where("model_detail.model_title like ?", "%"+param.ModelName+"%")
		}
		return tx
	}
}

// BuildModelQuery
// @Description: 构建模型查询条件
// @param: param
// @return func(db *gorm.DB) *gorm.DB
func BuildModelQuery(param *entity.ModelQueryBO) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tx := db
		if utils.IsStrNotBlank(param.ModelName) {
			tx = db.Where("model_detail.model_title like ?", "%"+param.ModelName+"%")
		}
		if param.Status != 0 {
			tx = db.Where("model_detail.status = ?", param.Status)
		}
		return tx
	}
}

// BuildModelPage
// @Description: 构建模型分页查询条件
// @param: param
// @return func(db *gorm.DB) *gorm.DB
func BuildModelPage(param *entity.ModelPageBO) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tx := db
		if utils.IsStrNotBlank(param.ModelName) {
			tx = db.Where("model_detail.model_title like ?", "%"+param.ModelName+"%")
		}
		if param.Status != 0 {
			tx = db.Where("model_detail.status = ?", param.Status)
		}
		return tx
	}
}
