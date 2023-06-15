package service

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/model"
	"reflect"
	"time"
)

// TransformInstTaskParam
// @Description: 转换实例任务参数
// @param instTaskParamMap
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

// GetInstTaskParam
// @Description: 转换实例任务参数
// @param instTaskParamMap
// @return []model.InstTaskParam
func GetInstTaskParam(instTaskID string) map[string]interface{} {
	var instTaskParamMap = make(map[string]interface{})
	var instTaskParams []model.InstTaskParam
	MysqlDB.Model(&model.InstTaskParam{}).Where("inst_task_id = ?", instTaskID).Find(&instTaskParams)
	if instTaskParams == nil {
		return instTaskParamMap
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
	return instTaskParamMap
}

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

func GetParamTypes(value interface{}) string {
	if value == nil {
		return "string"
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Float64:
		return "float64"
	case reflect.Float32:
		return "float32"
	case reflect.Int:
		return "int"
	case reflect.Uint:
		return "uint"
	case reflect.Int8:
		return "int8"
	case reflect.Uint8:
		return "uint8"
	case reflect.Int16:
		return "int16"
	case reflect.Uint16:
		return "uint16"
	case reflect.Int32:
		return "int32"
	case reflect.Uint32:
		return "uint32"
	case reflect.Int64:
		return "int64"
	case reflect.Uint64:
		return "uint64"
	case reflect.String:
		return "string"
	case reflect.Array:
		return "array"
	case reflect.Map:
		return "map"
	case reflect.Bool:
		return "bool"
	default:
		return "object"
	}
}
