# weflow


## 介绍

工作流（golang版本的流程引擎、规则引擎、表单引擎）、表单编辑器、流程编辑器

Workflow (a Golang-based flow engine, rule engine, form engine), form editor, process editor.

## 软件架构

```text
|-- bin # 二进制文件目录
|-- cmd # 编译入口
|   `-- app
|-- deploy # 环境和部署相关目录
|   |-- docker-compose # docker-compose 容器编排目录
|   `-- kubernetes # k8s 编排配置目录
|-- docs # 文档目录
|-- config # 配置文件目录
|-- internal
|   `-- app
|       |-- command # 命令行功能模块
|       |   |-- handler
|       |   `-- script # 临时脚本
|       |-- component # 功能组件，如：db, redis 等
|       |-- config # 配置模型
|       |-- cron # 定时任务功能模块
|       |   `-- job
|       |-- common # 数据库模型
|       |-- model # 数据库模型
|       |-- pkg # 功能类库
|       |-- dao # 数据处理层
|       |-- service # 业务逻辑层
|       |-- controller # 控制层
|       |-- router # 路由层
|       `-- transport
|           |-- grpc
|           |   |-- api # proto 文件目录
|           |   |-- handler # 控制层
|           |   `-- middleware # 中间件
|           `-- http
|               |-- api # swagger 文档
|               |-- handler # 控制层
|               |-- middleware # 中间件
|               `-- router # 路由
|-- logs # 日志目录
|-- pkg # 功能类库
`-- proto # 第三方 proto 文件目录

```



软件架构说明


支持特性
- 流程模板支持版本化管理，支持流程模板的导入导出，支持流程模板的复制、删除、发布、禁用、启用等
- 流程模板可以切换版本，支持流程模板的版本回退
- 流程支持发起、终止、挂起、恢复、删除等
- 用户任务支持同意、不同意、转办、委托、退回（上节点、任意、发起节点）、加签、减签、撤回、撤销、保存、会签、或签、并行、串行、自定义任务等
- 支持退回任务，退回到指定任意节点，退回到上一节点（上处理任务），退回到发起节点
- 支持转办任务，将任务交接给他人办理，办理完成后继续下一步骤
- 支持委托任务，将任务委托给他人，他人办理完成后再回到委托人
- 支持撤销任务，发起人可撤销未办理的任务，撤销后流程回到发起人
- 支持催办任务，处理人可以催办任务，催办后任务处理人会收到消息提醒
- 支持智能提交，相同处理人自动跳过，支持自由指定下一步处理人
- 支持作废流程，允许发起人快速终止流程，管理员维护终止流程
- 支持流程撤回，下一步未办理的任务，可进行取回撤销重做任务
- 支持流程跟踪图，流程实例，实例节点任务，实例用户任务，实例任务参数
- 支持节点表单权限，控制人员表单权限，如：只读、编辑、隐藏、必填等
- 待办任务，已办任务、已发任务、流程跟踪、审批记录
- 流程管控，在无关联表单情况下流程调试，如流程发起、挂起；流程定义、实例、任务等查询；任务办理，重定位等
- 支持流程组件标签定义（流程按钮、意见审批、下一步流程信息等）快速与自定义的业务表单建立关系。


## 设计文档



设计文档：

https://www.processon.com/view/link/6459ef517ca03d041ea38cba

### 流程编辑器设计

节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】

节点字段描述：

- 节点模型：nodeModel
- 节点名称：nodeName
- 节点ID：nodeId
- 父节点ID：parentId



#### 开始节点

```json
{
        "nodeModel":"节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】--数字",
        "nodeName":"发起节点",
        "nodeId":"节点ID",
        "parentId":"节点父ID"
    }
```



案例：

```json
{
        "nodeModel":1,
        "nodeName":"发起人",
        "nodeId":"1640993392605401001",
        "parentId":""
    }
```



#### 审批节点

```json
{
        "nodeModel":"节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】--数字",
        "nodeName":"审批节点",
        "nodeId":"节点ID",
        "parentId":"节点父ID",
        "formPer":[{
            "elemId": "表单元素ID",
            "elemPId": "表单元素父ID",
            "per": "表单权限【可编辑：1；只读：2；隐藏：3】默认只读2--数字"
        }],
        "approveType": "审批类型【人工审批：1；自动通过：2；自动拒绝】默认人工审批1--数字",
        "nodeSetting":{
            "execCheck":"审批执行校验，设计中可忽略",
            "timeout":"审批限时处理，设计中可忽略"
        },
        "nodeHandler": {
            "type": "常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】--数字",
            "handlers": [{
                "id": "处理人ID",
                "name": "处理人名称",
                "sort": "排序--数字"
            }],
            "strategy":"处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】--数字",
            "obj": "扩展字段，设计中可忽略",
            "relative":"相对发起人的直属主管，设计中可忽略"
        },
        "noneHandler":"审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1--数字",
        "appointHandler":"审批人为空时指定审批人ID",
        "handleMode":"审批方式【依次审批：1、会签（需要完成人数的审批人同意或拒绝才可完成节点）：2、或签（其中一名审批人同意或拒绝即可）：3】默认会签2--数字",
        "finishMode":"完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0）--数字"
    }
```



案例：

```json
{
        "nodeModel":2,
        "nodeName":"审批节点1",
        "nodeId":"1640993392605401002",
        "parentId":"",
        "formPer":[{
            "elemId": "param1",
            "elemPId": "",
            "per": 1
        }],
        "approveType": 1,
        "nodeSetting":{
            "execCheck":"",
            "timeout":""
        },
        "nodeHandler": {
            "type": "1",
            "handlers": [{
                "id": "547",
                "name": "xuch01",
                "sort": 1
            }],
            "strategy":1,
            "obj": "",
            "relative":""
        },
        "noneHandler":1,
        "appointHandler":"",
        "handleMode":2,
        "finishMode":0
    }
```



#### 办理节点

```json
{
        "nodeModel":"节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】--数字",
        "nodeName":"办理节点",
        "nodeId":"节点ID",
        "parentId":"节点父ID",
        "formPer":[{
            "elemId": "表单元素ID",
            "elemPId": "表单元素父ID",
            "per": "表单权限【可编辑：1；只读：2；隐藏：3】默认只读2--数字"
        }],
        "nodeSetting":{
            "timeout":"审批限时处理，设计中可忽略"
        },
        "nodeHandler": {
            "type": "常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】--数字",
            "handlers": [{
                "id": "处理人ID",
                "name": "处理人名称",
                "sort": "排序--数字"
            }],
            "strategy":"处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】--数字",
            "obj": "扩展字段，设计中可忽略",
            "relative":"相对发起人的直属主管，设计中可忽略"
        },
        "noneHandler":"审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1",
        "appointHandler":"审批人为空时指定审批人：3，指定审批人ID",
        "handleMode":"审批方式【依次审批：1、会签（需要完成人数的审批人同意或拒绝才可完成节点）：2、或签（其中一名审批人同意或拒绝即可）：3】默认会签2",
        "finishMode":"完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0）"
    }
```



案例：

```json
{
                    "nodeModel":3,
                    "nodeName":"办理节点1",
                    "nodeId":"1640993392605401005",
                    "parentId":"1640993392605401003",
                    "formPer":[{
                        "elemId": "param1",
                        "elemPId": "",
                        "per": 1
                    }],
                    "nodeSetting":{
                        "timeout":""
                    },
                    "nodeHandler": {
                        "type": "1",
                        "handlers": [{
                            "id": "547",
                            "name": "xuch01",
                            "sort": 1
                        }],
                        "strategy":1,
                        "obj": "",
                        "relative":""
                    },
                    "noneHandler":1,
                    "appointHandler":"",
                    "handleMode":2,
                    "finishMode":0
                }
```





#### 抄送节点

```json
{
        "nodeModel":"节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】--数字",
        "nodeName":"抄送节点",
        "nodeId":"节点ID",
        "parentId":"节点父ID",
        "formPer":[{
            "elemId": "表单元素ID",
            "elemPId": "表单元素父ID",
            "per": "表单权限【可编辑：1；只读：2；隐藏：3】默认只读2--数字"
        }],
        "nodeSetting":{
            "allowNotify":"允许发起人添加抄送人，先忽略"
        },
        "nodeHandler": {
            "type": "常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】--数字",
            "handlers": [{
                "id": "处理人ID",
                "name": "处理人名称",
                "sort": "排序--数字"
            }],
            "strategy":"处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】--数字",
            "obj": "扩展字段，设计中可忽略",
            "relative":"相对发起人的直属主管，设计中可忽略"
        }
    }
```



案例：

```json
{
                    "nodeModel":4,
                    "nodeName":"抄送节点1",
                    "nodeId":"1640993392605401007",
                    "parentId":"1640993392605401003",
                    "formPer":[{
                        "elemId": "param1",
                        "elemPId": "",
                        "per": 1
                    }],
                    "nodeSetting":{
                        "allowNotify":""
                    },
                    "nodeHandler": {
                        "type": "1",
                        "handlers": [{
                            "id": "547",
                            "name": "xuch01",
                            "sort": 1
                        }],
                        "strategy":1,
                        "obj": "",
                        "relative":""
                    }
                }
```



#### 分支节点

```json
{
        "nodeModel":"节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】--数字",
        "nodeName":"分支节点",
        "nodeId":"节点ID",
        "parentId":"节点父ID",
        "branchMode":"分支执行方式【单分支：1；多分支：2】默认多分支2--数字",
        "defaultBranch":"单分支处理需要默认分支，在条件优先级无法处理时候执行默认分支，取值分支下标--数字",
        "children":[
            [
                {
                    "nodeModel":"节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】--数字",
                    "nodeName":"条件节点",
                    "nodeId":"节点ID",
                    "parentId":"节点父ID对应分支节点ID",
                    "level":"优先级，分支执行方式为多分支处理方式无优先级应为0--数字",
                    "conditionGroup":"条件组前端描述展示条件组",
                    "conditionExpr":"条件组解析后的表达式"
                },
                {
                    "nodeModel":"节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】--数字",
                    "nodeName":"抄送节点",
                    "nodeId":"节点ID",
                    "parentId":"节点父ID",
                    "...":""
                }
            ]
        ]
    }
```







#### 条件节点

```json
{
                    "nodeModel":"节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】--数字",
                    "nodeName":"条件节点",
                    "nodeId":"节点ID",
                    "parentId":"节点父ID对应分支节点ID",
                    "level":"优先级，分支执行方式为多分支处理方式无优先级应为0--数字",
                    "conditionGroup":"条件组前端描述展示条件组",
                    "conditionExpr":"条件组解析后的表达式"
                }
```



#### 汇聚节点

```json
{
        "nodeModel":"节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】--数字",
        "nodeName":"汇聚节点",
        "nodeId":"节点ID",
        "parentId":"节点父ID"
    }
```



案例：

```json
{
        "nodeModel":8,
        "nodeName":"分支汇聚",
        "nodeId":"1640993392605401008",
        "parentId":""
    }
```





#### 自定义节点

```json
{
        "nodeModel":"节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】--数字",
        "nodeName":"自定义节点",
        "nodeId":"节点ID",
        "parentId":"节点父ID"
    }

设计中
```



#### 结束节点

```json
{
        "nodeModel":"节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】--数字",
        "nodeName":"结束节点",
        "nodeId":"节点ID",
        "parentId":"节点父ID"
    }
```



案例

```json
{
        "nodeModel":9,
        "nodeName":"流程结束",
        "nodeId":"1640993392605401009",
        "parentId":""
    }
```





## 使用说明

引入本地module
```go
require "github.com/wegoteam/weflow" v0.0.1
replace "github.com/wegoteam/weflow" => "../"
```

1.  cwgo语法
```shell
cwgo  model --db_type mysql --out_file dao_gen.go --out_dir ./backend/pkg/dao --dsn "root:root@tcp(localhost:3306)/weflow?charset=utf8&parseTime=True&loc=Local"

cwgo server --service flow --module github.com/wegoteam/weflow/internal --type HTTP  --idl ./idl/flow.thrift
cwgo server --service flow --module github.com/wegoteam/weflow/api --type RPC  --idl ./idl/flow.thrift
```

```shell
$ cwgo -h
NAME:
   cwgo - All in one tools for CloudWeGo

USAGE:
   cwgo [global options] command [command options] [arguments...]

COMMANDS:
   init      交互式命令行
   server    生成 RPC 或者 HTTP Server
   client    生成 RPC 或者 HTTP Client
   model     生成 DB Model
   fallback  回退到 kitex 或者 hz 工具

GLOBAL OPTIONS:
   --verbose      打开冗余日志模式
   --version, -v  打印工具版本


数据库模型语法
USAGE:
   cwgo model [command options] [arguments...]

OPTIONS:
   --dsn value                        Specify the database source name. (https://gorm.io/docs/connecting_to_the_database.html)
   --db_type value                    Specify database type. (mysql or sqlserver or sqlite or postgres) (default: mysql)
   --out_dir value                    Specify output directory (default: biz/dao/query)
   --out_file value                   Specify output filename (default: gen.go)
   --tables value [ --tables value ]  Specify databases tables
   --unittest                         Specify generate unit test (default: false)
   --only_model                       Specify only generate model code (default: false)
   --model_pkg value                  Specify model package name
   --nullable                         Specify generate with pointer when field is nullable (default: false)
   --type_tag                         Specify generate field with gorm column type tag (default: false)
   --index_tag                        Specify generate field with gorm index tag (default: false)
   --help, -h                         show help (default: false)
解释：
--dsn         指定数据库 DSN
--db_type     指定数据库类型
--out_dir     指定输出文件夹，默认 biz/dao/query
--out_file    指定输出文件名，默认 gen.go
--tables      指定数据库表名称
--unittest    是否生成单测，默认不生成
--only_model  是否只生成 model 代码，默认关闭
--model_pkg   指定 model package 名
--nullable    当字段为 null 时，指定是否生成指针，默认关闭
--type_tag    是否给字段生成 gorm column type tag，默认不生成  
--index_tag   是否给字段生成 gorm index tag，默认不生成  

 
客户端、服务端语法
USAGE:
   cwgo client [command options] [arguments...]

OPTIONS:
   --service value                                                              Specify the service name.
   --type value                                                                 Specify the generate type. (RPC or HTTP) (default: "RPC")
   --module value, --mod value                                                  Specify the Go module name to generate go.mod.
   --idl value                                                                  Specify the IDL file path. (.thrift or .proto)
   --out_dir value, -o value                                                    Specify the output path. (default: biz/http)
   --registry value                                                             Specify the registry, default is None
   --proto_search_path value, -I value [ --proto_search_path value, -I value ]  Add an IDL search path for includes. (Valid only if idl is protobuf)
   --pass value [ --pass value ]                                                pass param to hz or kitex
   --help, -h                                                                   show help (default: false)
解释：
--service    指定服务名称
--type       指定生成类型
--module     指定生成 module 名称
--idl        指定 IDL 文件路径
--out_dir    指定输出路径
--template   指定 layout 模板路径
--registry   指定服务注册组件
--proto_search_path 添加 IDL 搜索路径，只对 pb 生效
--pass value 传递给 hz 和 kitex 的参数

```

2.swagger语法
```shell
swag init --parseDependency --parseInternal --parseDepth 6 --instanceName "weflow"

swag init
```
http://localhost:8081/weflow/swagger/index.html

https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format

https://github.com/swaggo/swag/blob/master/README_zh-CN.md