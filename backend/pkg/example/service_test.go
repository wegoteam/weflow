package example

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/model"
	"github.com/wegoteam/weflow/pkg/service"
	"gorm.io/gorm"
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
		paramType := service.GetParamType(val)
		t := reflect.TypeOf(val).String()

		kind := reflect.ValueOf(val).Kind()
		hlog.Infof("val 的类型是 %v kind=%v", t, kind)
		hlog.Infof("key=%v  val=%v   valType=%v", key, val, paramType)
	}

}

func TestGetInstTaskParam(t *testing.T) {
	instTaskParamMap, _ := service.GetInstTaskParamMap("421397709668421")
	hlog.Infof("instTaskParamMap是 %v", instTaskParamMap)
}

func TestGetModelVersion(t *testing.T) {
	modelVersion := service.GetModelVersion("420915317174341", "1681335332954505235")
	hlog.Infof("GetModelVersion= %v", modelVersion)

	modelVersionList, _ := service.GetModelVersionList("433984855478341")
	hlog.Infof("GetModelVersionList= %v", modelVersionList)

	modelVersion2 := service.GetEnableModelVersion("420915317174341")
	hlog.Infof("GetEnableModelVersion= %v", modelVersion2)
}

func TestGetInstNodeUserTask(t *testing.T) {
	instNodeUserTask := service.GetInstNodeUserTask("424136865722437")
	hlog.Infof("instNodeUserTask= %v", instNodeUserTask)
}

func TestGetExecNodeTaskMap(t *testing.T) {
	execNodeTaskMap := service.GetExecNodeTaskMap("425247133954117")
	hlog.Infof("execNodeTaskMap= %v", execNodeTaskMap)
}

func TestGetTodoUserTask(t *testing.T) {
	param := &entity.UserTaskQueryBO{
		UserID:          "547",
		PageNum:         2,
		PageSize:        1,
		InstStatus:      2,
		CreateTimeStart: "2020-8-5 13:14:15",
		CreateTimeEnd:   "2024-8-5 13:14:15",
	}
	userTask := service.GetTodoUserTasks("547")
	page, _ := service.PageTodoUserTasks(param)
	hlog.Infof("userTask= %v", userTask)
	hlog.Infof("page= %v", page)
}

func TestGetDoneUserTask(t *testing.T) {
	userTask, _ := service.GetDoneUserTasks("547")
	hlog.Infof("userTask= %v", userTask)
}

func TestGetInitiatingInstTask(t *testing.T) {
	initiatingInstTasks := service.GetInitiatingInstTasks("547")
	hlog.Infof("initiatingInstTasks= %v", initiatingInstTasks)
}

func TestGetDraftInstTask(t *testing.T) {
	draftInstTask := service.GetDraftInstTask("547")
	hlog.Infof("draftInstTask= %v", draftInstTask)
}

func TestGetModelList(t *testing.T) {
	var existModel = model.ModelDetail{}
	err := service.MysqlDB.Model(&model.ModelDetail{}).Where("model_id = ?", "1").First(&existModel).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		hlog.Errorf("查询模板失败,模版ID输入有误，重新生成新的模板 error: %v", err)
	}

	var existModel2 = []model.ModelDetail{}
	service.MysqlDB.Model(&model.ModelDetail{}).Where("model_id = ?", "1").Find(&existModel2)
	hlog.Infof("existModel2= %v", existModel2)
}

func TestGetRoleTree(t *testing.T) {
	param := &entity.RoleQueryBO{}
	roleTree, err := service.GetRoleTree(param)
	if err != nil {
		hlog.Errorf("查询角色树失败 error: %v", err)
	}
	hlog.Infof("roleTree= %s", roleTree)
}

func TestGetAllRoleUserList(t *testing.T) {
	roleUserList, err := service.GetAllRoleUserTree()
	if err != nil {
		hlog.Errorf("查询角色用户列表失败 error: %v", err)
	}
	hlog.Infof("roleUserList= %s", roleUserList)
}
