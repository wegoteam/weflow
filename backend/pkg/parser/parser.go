package parser

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"github.com/wegoteam/weflow/pkg/common/config"
	"github.com/wegoteam/weflow/pkg/common/constant"
	"github.com/wegoteam/weflow/pkg/common/entity"
)

/**
流程引擎解析流程定义
节点模型【1：开始节点；2：审批节点；3：知会节点；4：自定义节点；5：条件节点；6：分支节点；7：汇聚节点；8：结束节点】
解析生成节点实例：上节点、下节点、节点下标、尾节点等基础信息

表单引擎解析表单定义
表单类型

*/

var (
	MysqlDB    = config.MysqlDB
	RedisCliet = config.RedisCliet
)

// GetProcessDefModel
//  @Description: 获取流程定义模型
//获取缓存的数据，不存在则部署
//  @param processDefId
//  @return *entity.ProcessDefModel
func GetProcessDefModel(processDefId string) (*entity.ProcessDefModel, error) {
	var processDefKey = constant.RedisProcessDefModel + processDefId
	ctx := context.Background()
	hasKey, existErr := RedisCliet.Exists(ctx, processDefKey).Result()
	if existErr != nil {
		hlog.Warnf("获取流程定义模型失败，错误信息：%s", existErr.Error())
		return nil, errors.New("获取流程定义模型失败")
	}
	//存在，获取Redis内存的数据
	if hasKey == constant.HasRedisProcessDefModel {
		return buildProcessDefOnRedis(processDefId)
	}
	return buildProcessDefOnDB(processDefId)
}

// GetFormDefModel
// @Description: 获取表单定义模型
func GetFormDefModel() {

}
