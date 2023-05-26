package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"reflect"
	"testing"
)

type TsetParam struct {
	Param1 string
	Param2 int64
	Param3 float64
	Param4 []string
}

func TestGetParamType(t *testing.T) {

	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = 1
	instTaskParamMap["testparam2"] = 1.22222222222222222222222222222222222222
	instTaskParamMap["testparam3"] = 22222222222222222
	instTaskParamMap["testparam4"] = "testparam4"
	instTaskParamMap["testparam5"] = 1.2222

	var slice = make([]string, 0)
	var slice2 = make([]TsetParam, 0)
	slice = append(slice, "test1")

	instTaskParamMap["testparam6"] = slice

	var tsetParam = &TsetParam{
		Param1: "test1",
		Param2: 1,
		Param3: 22222222222,
		Param4: slice,
	}
	instTaskParamMap["testparam7"] = tsetParam

	slice2 = append(slice2, *tsetParam)
	instTaskParamMap["testparam8"] = slice2
	for key, val := range instTaskParamMap {
		paramType := GetParamType(val)
		t := reflect.TypeOf(val).String()

		kind := reflect.ValueOf(val).Kind()
		hlog.Infof("val 的类型是 %v kind=%v", t, kind)
		hlog.Infof("key=%v  val=%v   valType=%v", key, val, paramType)
	}

}

func TestGetInstTaskParam(t *testing.T) {
	instTaskParamMap := GetInstTaskParam("421397709668421")
	hlog.Infof("instTaskParamMap是 %v", instTaskParamMap)
}

func TestGetModelVersion(t *testing.T) {
	modelVersion := GetModelVersion("420915317174341", "1681335332954505235")
	hlog.Infof("GetModelVersion= %v", modelVersion)

	modelVersionList := GetModelVersionList("420915317174341", "1681335332954505235")
	hlog.Infof("GetModelVersionList= %v", modelVersionList)

	modelVersion2 := GetEnableModelVersion("420915317174341")
	hlog.Infof("GetEnableModelVersion= %v", modelVersion2)
}
