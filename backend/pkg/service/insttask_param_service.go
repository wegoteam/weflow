package service

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/model"
	"time"
)

// TransformInstTaskParam
// @Description: 转换实例任务参数
// @param: instTaskParamMap
// @return []model.InstTaskParam
func TransformInstTaskParam(instTaskID string, instTaskParamMap map[string]interface{}, now time.Time) []model.InstTaskParam {
	var instTaskParams = make([]model.InstTaskParam, 0)

	if instTaskParamMap == nil || len(instTaskParamMap) == 0 {
		return instTaskParams
	}
	for paramId, paramVal := range instTaskParamMap {
		paramByte, err := sonic.Marshal(paramVal)
		if err != nil {
			hlog.Errorf("实例任务[%s]参数[%s]转换失败", instTaskID, paramId, err)
			continue
		}
		var instTaskParam = &model.InstTaskParam{
			InstTaskID:    instTaskID,
			ParamID:       paramId,
			ParamName:     "",
			ParamDataType: GetParamType(paramVal),
			ParamBinary:   paramByte,
			ParamValue:    "",
			CreateTime:    now,
			UpdateTime:    now,
		}
		instTaskParams = append(instTaskParams, *instTaskParam)
	}

	return instTaskParams
}

// GetInstTaskParamMap
// @Description: 查询实例任务参数
// @param: instTaskParamMap
// @return []model.InstTaskParam
func GetInstTaskParamMap(instTaskID string) (map[string]interface{}, error) {
	var instTaskParamMap = make(map[string]interface{})
	var instTaskParams []model.InstTaskParam
	err := MysqlDB.Model(&model.InstTaskParam{}).Where("inst_task_id = ?", instTaskID).Find(&instTaskParams).Error
	if err != nil {
		hlog.Errorf("查询实例任务[%s]参数失败", instTaskID, err)
		return instTaskParamMap, errors.New("查询实例任务参数失败")
	}
	if instTaskParams == nil {
		return instTaskParamMap, nil
	}
	for _, instTaskParam := range instTaskParams {
		var obj interface{}
		err := sonic.Unmarshal(instTaskParam.ParamBinary, &obj)
		if err != nil {
			hlog.Errorf("实例任务[%s]参数[%s]转换失败", instTaskID, instTaskParam.ParamID, err)
			continue
		}
		instTaskParamMap[instTaskParam.ParamID] = obj
	}
	return instTaskParamMap, nil
}

// GetInstTaskParams
// @Description: 查询实例任务参数
// @param: instTaskID 实例任务id
// @return []entity.InstTaskParamResult
// @return error
func GetInstTaskParams(instTaskID string) ([]entity.InstTaskParamResult, error) {
	var instTaskParams []model.InstTaskParam
	err := MysqlDB.Model(&model.InstTaskParam{}).Where("inst_task_id = ?", instTaskID).Find(&instTaskParams).Error
	if err != nil {
		hlog.Errorf("查询实例任务[%s]参数失败", instTaskID, err)
		return nil, errors.New("查询实例任务参数失败")
	}
	var instTaskParamResults = make([]entity.InstTaskParamResult, 0)
	if instTaskParams == nil {
		return instTaskParamResults, nil
	}
	for _, instTaskParam := range instTaskParams {
		param := entity.InstTaskParamResult{
			ID:            instTaskParam.ID,
			InstTaskID:    instTaskParam.InstTaskID,
			ParamID:       instTaskParam.ParamID,
			ParamName:     instTaskParam.ParamName,
			ParamValue:    instTaskParam.ParamValue,
			CreateTime:    instTaskParam.CreateTime,
			UpdateTime:    instTaskParam.UpdateTime,
			ParamDataType: instTaskParam.ParamDataType,
			ParamBinary:   instTaskParam.ParamBinary,
		}
		instTaskParamResults = append(instTaskParamResults, param)
	}
	return instTaskParamResults, nil
}

// GetParamType
// @Description: 获取参数类型
// @param: value
// @return string
func GetParamType(value interface{}) string {
	if value == nil {
		return "string"
	}
	switch value.(type) {
	case float64:
		return "float64"
	case float32:
		return "float32"
	case int:
		return "int"
	case uint:
		return "uint"
	case int8:
		return "int8"
	case uint8:
		return "uint8"
	case int16:
		return "int16"
	case uint16:
		return "uint16"
	case int32:
		return "int32"
	case uint32:
		return "uint32"
	case int64:
		return "int64"
	case uint64:
		return "uint64"
	case string:
		return "string"
	case []string:
		return "array"
	case []int:
		return "array"
	case bool:
		return "bool"
	case []byte:
		return "[]byte"
	default:
		return "object"
	}
}
