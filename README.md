# weflow

#### 介绍
{**以下是 Gitee 平台说明，您可以替换此简介**
Gitee 是 OSCHINA 推出的基于 Git 的代码托管平台（同时支持 SVN）。专为开发者提供稳定、高效、安全的云端软件开发协作平台
无论是个人、团队、或是企业，都能够用 Gitee 实现代码托管、项目管理、协作开发。企业项目请看 [https://gitee.com/enterprises](https://gitee.com/enterprises)}

#### 软件架构

```json
|-- bin # 二进制文件目录
|-- cmd # 编译入口
|   `-- app
|-- deploy # 环境和部署相关目录
|   |-- docker-compose # docker-compose 容器编排目录
|   `-- kubernetes # k8s 编排配置目录
|-- docs # 文档目录
|-- etc # 配置文件目录
|-- internal
|   `-- app
|       |-- command # 命令行功能模块
|       |   |-- handler
|       |   `-- script # 临时脚本
|       |-- component # 功能组件，如：db, redis 等
|       |-- config # 配置模型
|       |-- cron # 定时任务功能模块
|       |   `-- job
|       |-- model # 数据库模型
|       |-- pkg # 功能类库
|       |-- repository # 数据处理层
|       |-- service # 业务逻辑层
|       |-- test
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
- 支持流程办理、退回、自由流、会签、并行、串行、服务任务等
- 支持退回任务，退回到指定环节，退回到上一步，退回到发起人
- 支持转办任务，将任务交接给他人办理，办理完成后继续下一步骤
- 支持委托任务，将任务委托给他人，他人办理完成后再回到委托人
- 支持智能提交，相同处理人自动跳过，支持自由指定下一步处理人
- 支持作废流程，允许发起人快速终止流程，管理员维护终止流程
- 支持自由流程，根据环节选择，自由跳转到指定环节，特事特办
- 支持流程撤回，下一步未办理的任务，可进行取回撤销重做任务
- 支持流程跟踪图，流程状态展现，流转信息，任务历史，任务分配信息
- 支持一个流程模型挂接多个业务单据，如某公司 8 种费用审批流程，表单不一样，但流程相同
- 支持一个表单挂接多个流程环节，以表单角度去管理流程，方便业务理解
- 支持全局表单，用于流程全局表单配置，目前支持内置表单、url 表单。如果不配置则发起流程会提示错误。
- 支持节点表单，节点表单配置。如果不配置默认使用全局表单。
- 流程事件脚本在线编写，包括：流程启动、完成、取消；任务分配、创建、结束等
- 流程脚本管理（Groovy、Beetl），在线编辑、自动完成、脚本测试、多语言脚本模板维护
- 我的待办任务处理，我的已办任务、我创建的任务查询、流程跟踪、审批记录查询
- 流程管控，在无关联表单情况下流程调试，如流程发起、挂起；流程定义、实例、任务等查询；任务办理，重定位等
- 支持流程组件标签定义（流程按钮、意见审批、下一步流程信息等）快速与自定义的业务表单建立关系。
- 支持版本化管理流程，新调整的流程业务不影响正在运行，未结束的流程继续流转。
- 支持任务加签、催办任务、传阅任务、流程委托设置、流水号管理、常用语管理

#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 使用说明

引入本地module
```go
require "wego2023/weflow" v0.0.1
replace "wego2023/weflow" => "../"
```

1.  cwgo语法
```shell
cwgo  model --db_type mysql --out_file dao_gen.go --out_dir ./backend/pkg/dao --dsn "root:root@tcp(localhost:3306)/weflow?charset=utf8&parseTime=True&loc=Local"

cwgo server --service hello --module github.com/wego2023/weflow/internal --type HTTP  --idl ./idl/hello.thrift
cwgo server --service hello --module github.com/wego2023/weflow/internal --type RPC  --idl ./idl/hello.thrift
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

2.  xxxx
3.  xxxx

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
