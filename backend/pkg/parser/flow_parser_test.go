package parser

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"github.com/wegoteam/weflow/pkg/model"
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	var data = "[{\"nodeModel\":1,\"nodeName\":\"发起人\",\"nodeId\":\"1640993392605401088\",\"parentId\":\"\"},{\"nodeModel\":2,\"nodeName\":\"审批\",\"nodeId\":\"1640993449224310784\",\"parentId\":\"\",\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"2\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"2\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\"},{\"nodeModel\":6,\"nodeName\":\"分支节点\",\"nodeId\":\"1640993508049424384\",\"parentId\":\"\",\"children\":[[{\"nodeModel\":5,\"nodeName\":\"条件1\",\"nodeId\":\"1640993508049424385\",\"parentId\":\"1640993508049424384\"},{\"nodeModel\":2,\"nodeName\":\"审批\",\"nodeId\":\"1640993526328201216\",\"parentId\":\"1640993508049424384\",\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"2\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"2\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\"}],[{\"nodeModel\":5,\"nodeName\":\"条件2\",\"nodeId\":\"1640993508049424386\",\"parentId\":\"1640993508049424384\"},{\"nodeModel\":4,\"nodeName\":\"知会\",\"nodeId\":\"1640993535555670016\",\"parentId\":\"1640993508049424384\",\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"1\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\"}]]},{\"nodeModel\":7,\"nodeName\":\"分支汇聚\",\"nodeId\":\"1640993508053618688\",\"parentId\":\"\"},{\"nodeModel\":8,\"nodeName\":\"流程结束\",\"nodeId\":\"1640993392605401089\",\"parentId\":\"\"}]"
	nodes := Parser(data)
	marshal, _ := sonic.Marshal(nodes)

	hlog.Infof("数据为%v", string(marshal))
}

/**
https://github.com/redis/go-redis/blob/master/example/scan-struct/main.go
*/
func TestRedis(t *testing.T) {
	ctx := context.Background()
	err := RedisCliet.Set(ctx, "map", "key", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := RedisCliet.Get(ctx, "map").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}

func TestDB(t *testing.T) {
	ctx := context.Background()
	var flowdef = &model.ProcessDefInfo{}
	MysqlDB.WithContext(ctx).Where(&model.ProcessDefInfo{ProcessDefID: "1640993392605401001"}).First(flowdef)
	marshal, _ := sonic.Marshal(flowdef)
	fmt.Println(string(marshal))
}

func TestGetProcessDefModel(t *testing.T) {
	ctx := context.Background()

	var data = "[{\"nodeModel\":1,\"nodeName\":\"发起人\",\"nodeId\":\"1640993392605401088\",\"parentId\":\"\"},{\"nodeModel\":2,\"nodeName\":\"审批\",\"nodeId\":\"1640993449224310784\",\"parentId\":\"\",\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"2\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"2\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\"},{\"nodeModel\":6,\"nodeName\":\"分支节点\",\"nodeId\":\"1640993508049424384\",\"parentId\":\"\",\"children\":[[{\"nodeModel\":5,\"nodeName\":\"条件1\",\"nodeId\":\"1640993508049424385\",\"parentId\":\"1640993508049424384\"},{\"nodeModel\":2,\"nodeName\":\"审批\",\"nodeId\":\"1640993526328201216\",\"parentId\":\"1640993508049424384\",\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"2\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"2\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\"}],[{\"nodeModel\":5,\"nodeName\":\"条件2\",\"nodeId\":\"1640993508049424386\",\"parentId\":\"1640993508049424384\"},{\"nodeModel\":4,\"nodeName\":\"知会\",\"nodeId\":\"1640993535555670016\",\"parentId\":\"1640993508049424384\",\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"1\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\"}]]},{\"nodeModel\":7,\"nodeName\":\"分支汇聚\",\"nodeId\":\"1640993508053618688\",\"parentId\":\"\"},{\"nodeModel\":8,\"nodeName\":\"流程结束\",\"nodeId\":\"1640993392605401089\",\"parentId\":\"\"}]"
	nodes := Parser(data)

	_, err := RedisCliet.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		for _, node := range *nodes {
			nodeStr, _ := sonic.Marshal(node)
			RedisCliet.HSet(ctx, "process_def", node.NodeID, string(nodeStr))
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	val := RedisCliet.Exists(ctx, "process_def").Val()
	fmt.Println(val)
	//https://redis.uptrace.dev/zh/guide/go-redis-pipelines.html#%E7%AE%A1%E9%81%93
	var nodeMapStr map[string]string
	nodeMapStr = RedisCliet.HGetAll(ctx, "process_def").Val()
	fmt.Println(nodeMapStr)

	//mapdata := map[string]string{}
	//err = RedisCliet.HGetAll(ctx, "process_def").Scan(&mapdata)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//cmds, err := client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
	//	keys := client.HKeys(ctx, "process_def")
	//	for _,val := range keys.Val() {
	//		client.HGet(val)
	//	}
	//	return nil
	//})
	//if err != nil {
	//	panic(err)
	//}
	//for _, cmd := range cmds {
	//	fmt.Println(cmd.(*redis.StringCmd).Val())
	//}
	fmt.Println(nodeMapStr)
}

func TestProcessDefModel(t *testing.T) {
	//processDefModel := GetProcessDefModel("1640993392605401001")
	processDefModel := GetProcessDefModel("1640993392605401002")

	nodeMap := make(map[string]interface{})
	for _, node := range *processDefModel.NodeModels {
		nodeMap[node.NodeID] = node
	}
	for k, v := range nodeMap {
		switch v.(type) {
		case entity.NodeModelBO:
			ctx := context.Background()
			err := RedisCliet.Set(ctx, k, v, time.Second*100).Err()
			fmt.Println(err)
			fmt.Println(k, v)
		}
	}
	if processDefModel != nil {
		marshal, _ := sonic.Marshal(processDefModel)
		fmt.Println(string(marshal))
	}

}
