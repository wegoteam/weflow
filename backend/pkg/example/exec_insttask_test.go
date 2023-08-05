package example

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/exec"
	"github.com/wegoteam/weflow/pkg/model"
	"github.com/wegoteam/weflow/pkg/parser"
	"github.com/wegoteam/weflow/pkg/service"
	"github.com/wegoteam/wepkg/id/snowflake"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)

func TestStartInstTask(t *testing.T) {
	processDefModel, _ := parser.GetProcessDefModel("1640993392605401001")

	execution := &exec.Execution{}
	execution.ProcessDefModel = processDefModel
	execution.InstTaskID = snowflake.GetSnowflakeId()
	execution.InstTaskName = "测试流程"
	execution.InstTaskStatus = constant.InstanceTaskStatusDoing
	execution.Now = time.Now()
	//startNodeId := processDefModel.StartNodeId
	//startNode := processDefModel.NodeModelMap[startNodeId]

	//实例任务参数
	var instTaskParamMap = make(map[string]interface{})
	execution.InstTaskParamMap = instTaskParamMap
	//实例节点任务执行缓存数据
	var execNodeTaskMap = make(map[string]entity.ExecNodeTaskBO)
	execution.ExecNodeTaskMap = execNodeTaskMap
	//用户任务
	var userTasks = make([]entity.UserTaskBO, 0)
	execution.UserTasks = &userTasks
	//实例节点任务
	var instNodeTasks = make([]entity.InstNodeTaskBO, 0)
	execution.InstNodeTasks = &instNodeTasks
	//实例节点任务表单权限
	var taskFormPers = make([]entity.TaskFormPerBO, 0)
	execution.TaskFormPers = &taskFormPers
	//实例任务操作日志
	var instTaskOpLogs = make([]entity.InstTaskOpLogBO, 0)
	execution.InstTaskOpLogs = &instTaskOpLogs

	execution.CreateUserName = "xuch01"
	execution.CreateUserID = "547"
	execution.ProcessDefId = "1640993392605401001"
	execution.FormDefId = "1640993392605401001"
	//exec.execNode(&startNode, execution)

	hlog.Infof("执行结果%+v", execution)
}

func TestInstTaskExecution(t *testing.T) {
	processDefModel, _ := parser.GetProcessDefModel("1640993392605401001")

	execution := &exec.Execution{}
	execution.ProcessDefModel = processDefModel
	execution.InstTaskID = snowflake.GetSnowflakeId()
	execution.InstTaskName = "测试流程"
	execution.InstTaskStatus = constant.InstanceTaskStatusDoing
	execution.Now = time.Now()
	//startNodeId := processDefModel.StartNodeId
	//startNode := processDefModel.NodeModelMap[startNodeId]

	//实例任务参数
	var instTaskParamMap = make(map[string]interface{})
	execution.InstTaskParamMap = instTaskParamMap
	//实例节点任务执行缓存数据
	var execNodeTaskMap = make(map[string]entity.ExecNodeTaskBO)
	execution.ExecNodeTaskMap = execNodeTaskMap
	//用户任务
	var userTasks = make([]entity.UserTaskBO, 0)
	execution.UserTasks = &userTasks
	//实例节点任务
	var instNodeTasks = make([]entity.InstNodeTaskBO, 0)
	execution.InstNodeTasks = &instNodeTasks
	//实例节点任务表单权限
	var taskFormPers = make([]entity.TaskFormPerBO, 0)
	execution.TaskFormPers = &taskFormPers
	//实例任务操作日志
	var instTaskOpLogs = make([]entity.InstTaskOpLogBO, 0)
	execution.InstTaskOpLogs = &instTaskOpLogs

	execution.CreateUserName = "xuch01"
	execution.CreateUserID = "547"
	execution.ProcessDefId = "1640993392605401001"
	execution.FormDefId = "1640993392605401001"
	//exec.execNode(&startNode, execution)

	hlog.Infof("执行结果%+v", execution)
	//instTaskExecution := &exec.InstTaskExecution{
	//	Execution:  execution,
	//	ModelID:    "420915317174341",
	//	VersionID:  "1681335332954505235",
	//	OpUserID:   "547",
	//	OpUserName: "xuch01",
	//}
	//instTaskExecution.execStartInstData()
}

func TestTransformInstTaskParam(t *testing.T) {
	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = "testparam1"
	instTaskParamMap["testparam2"] = "testparam22222"
	instTaskParamMap["testparam3"] = "testparam33333"
	instTaskParamMap["testparam3"] = "testparam4"

	instTaskExecution := &exec.InstTaskExecution{
		Execution:  &exec.Execution{},
		ModelID:    "420915317174341",
		VersionID:  "1681335332954505235",
		OpUserID:   "547",
		OpUserName: "xuch01",
	}
	instTaskParamMap["instTaskExecution"] = instTaskExecution
	instTaskID := snowflake.GetSnowflakeId()
	//instTaskID := "421395986214981"
	instTaskParam := service.TransformInstTaskParam(instTaskID, instTaskParamMap, time.Now())

	exec.MysqlDB.Debug().CreateInBatches(instTaskParam, len(instTaskParam))
	taskParamMap, _ := service.GetInstTaskParamMap(instTaskID)
	hlog.Infof("taskParamMap=%+v", taskParamMap)
	var instTaskParam2 []model.InstTaskParam
	exec.MysqlDB.Debug().Where("inst_task_id = ?", instTaskID).Find(&instTaskParam2)

	instTaskParamForBsonD := transformInstTaskParamForD(instTaskID, instTaskParamMap, time.Now())

	filter := bson.M{"inst_task_id": instTaskID}
	collection := exec.MongoClient.Database("weflow").Collection("inst_task_param")
	hlog.Infof("查询数据为%+v", collection)
	//opts := options.InsertMany().SetOrdered(false)
	//_, err := collection.InsertMany(context.TODO(), instTaskParamForBsonD, opts)

	_, err := collection.UpdateMany(context.TODO(), instTaskParamForBsonD, filter)
	if err != nil {
		hlog.Fatal(err)
	}

	hlog.Infof("执行结果%+v，查询数据为%+v", instTaskParam, instTaskParam2)
}

func transformInstTaskParamForD(instTaskID string, instTaskParamMap map[string]interface{}, now time.Time) []interface{} {
	var instTaskParams = make([]interface{}, 0)

	if instTaskParamMap == nil || len(instTaskParamMap) == 0 {
		return instTaskParams
	}
	for paramId, paramVal := range instTaskParamMap {
		instTaskParam := bson.D{
			{"inst_task_id", instTaskID},
			{"param_id", paramId},
			{"param_name", ""},
			{"param_value", paramVal},
			{"create_time", now},
			{"update_time", now},
		}
		instTaskParams = append(instTaskParams, instTaskParam)
	}

	return instTaskParams
}

func transformInstTaskParamForM(instTaskID string, instTaskParamMap map[string]interface{}, now time.Time) []interface{} {
	var instTaskParams = make([]interface{}, 0)

	if instTaskParamMap == nil || len(instTaskParamMap) == 0 {
		return instTaskParams
	}
	for paramId, paramVal := range instTaskParamMap {
		instTaskParam := bson.D{
			{"inst_task_id", instTaskID},
			{"param_id", paramId},
			{"param_name", ""},
			{"param_value", paramVal},
			{"create_time", now},
			{"update_time", now},
		}
		instTaskParams = append(instTaskParams, instTaskParam)
	}

	return instTaskParams
}

type Params struct { // user type for example
	ID         primitive.ObjectID `bson:"_id"`
	InstTaskID string             `bson:"inst_task_id" `
	ParamID    string             `bson:"param_id" `
	ParamName  string             `bson:"param_name"`
	ParamValue interface{}        `bson:"param_value"`
	UpdateTime time.Time          `bson:"update_time"`
	CreateTime time.Time          `bson:"create_time"`
}

func TestMongodb(t *testing.T) {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	mgoCli, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := mgoCli.Database("weflow").Collection("inst_task_param")
	hlog.Infof("查询数据为%+v", collection)

	cursor, _ := collection.Find(context.Background(), bson.D{{"inst_task_id", "421395986214981"}})

	//var results []bson.M
	var results []Params
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		hlog.Infof("查询数据为%+v", result)
	}
	hlog.Infof("查询数据为%+v", cursor)

}

func TestStartProcessInstTask(t *testing.T) {
	var createUserName = "xuch01"
	var createUserID = "547"
	var modelID = "420915317174341"
	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = "testparam1"
	instTaskParamMap["testparam2"] = "testparam22222"
	instTaskParamMap["testparam3"] = "testparam33333"
	instTaskParamMap["testparam3"] = "testparam4"
	instTaskId, _ := exec.Start(modelID, createUserID, createUserName, instTaskParamMap)
	hlog.Infof("执行结果:%v", instTaskId)
}

func TestStop(t *testing.T) {
	var createUserName = "xuch01"
	var createUserID = "547"
	var modelID = "420915317174341"
	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = "testparam1"
	instTaskParamMap["testparam2"] = "testparam2"
	instTaskParamMap["testparam3"] = "testparam3"
	instTaskParamMap["testparam4"] = "testparam4"
	instTaskId, _ := exec.Start(modelID, createUserID, createUserName, instTaskParamMap)
	hlog.Infof("执行结果:%v", instTaskId)

	userID := "547"
	userName := "xuch01"
	desc := "测试"
	err := exec.Stop(instTaskId, userID, userName, desc)
	hlog.Infof("执行结果:%v", err)
}

func TestSuspend(t *testing.T) {
	var createUserName = "xuch01"
	var createUserID = "547"
	var modelID = "420915317174341"
	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = "testparam1"
	instTaskParamMap["testparam2"] = "testparam2"
	instTaskParamMap["testparam3"] = "testparam3"
	instTaskParamMap["testparam4"] = "testparam4"
	instTaskId, _ := exec.Start(modelID, createUserID, createUserName, instTaskParamMap)
	hlog.Infof("执行结果:%v", instTaskId)

	userID := "547"
	userName := "xuch01"
	desc := "测试"
	err := exec.Suspend(instTaskId, userID, userName, desc)
	hlog.Infof("执行结果:%v", err)
}

func TestSesume(t *testing.T) {
	var createUserName = "xuch01"
	var createUserID = "547"
	var modelID = "420915317174341"
	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = "testparam1"
	instTaskParamMap["testparam2"] = "testparam2"
	instTaskParamMap["testparam3"] = "testparam3"
	instTaskParamMap["testparam4"] = "testparam4"
	instTaskId, _ := exec.Start(modelID, createUserID, createUserName, instTaskParamMap)
	hlog.Infof("执行结果:%v", instTaskId)

	userID := "547"
	userName := "xuch01"
	desc := "测试"
	err := exec.Suspend(instTaskId, userID, userName, desc)
	hlog.Infof("执行结果:%v", err)

	err2 := exec.Sesume(instTaskId, userID, userName, desc)
	hlog.Infof("执行结果:%v", err2)
}

func TestDelete(t *testing.T) {
	var createUserName = "xuch01"
	var createUserID = "547"
	var modelID = "420915317174341"
	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = "testparam1"
	instTaskParamMap["testparam2"] = "testparam2"
	instTaskParamMap["testparam3"] = "testparam3"
	instTaskParamMap["testparam4"] = "testparam4"
	instTaskId, _ := exec.Start(modelID, createUserID, createUserName, instTaskParamMap)
	hlog.Infof("执行结果:%v", instTaskId)

	userID := "547"
	userName := "xuch01"
	desc := "测试"
	err := exec.Suspend(instTaskId, userID, userName, desc)
	hlog.Infof("执行结果:%v", err)

	err2 := exec.Delete(instTaskId, userID, userName, desc)
	hlog.Infof("执行结果:%v", err2)
}
