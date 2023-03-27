# weflow

#### 介绍
{**以下是 Gitee 平台说明，您可以替换此简介**
Gitee 是 OSCHINA 推出的基于 Git 的代码托管平台（同时支持 SVN）。专为开发者提供稳定、高效、安全的云端软件开发协作平台
无论是个人、团队、或是企业，都能够用 Gitee 实现代码托管、项目管理、协作开发。企业项目请看 [https://gitee.com/enterprises](https://gitee.com/enterprises)}

#### 软件架构
软件架构说明

预览：http://47.109.22.115:10090/task/done

支持特性
支持流程办理、退回、自由流、会签、并行、串行、服务任务等
支持退回任务，退回到指定环节，退回到上一步，退回到发起人
支持转办任务，将任务交接给他人办理，办理完成后继续下一步骤
支持委托任务，将任务委托给他人，他人办理完成后再回到委托人
支持智能提交，相同处理人自动跳过，支持自由指定下一步处理人
支持作废流程，允许发起人快速终止流程，管理员维护终止流程
支持自由流程，根据环节选择，自由跳转到指定环节，特事特办
支持流程撤回，下一步未办理的任务，可进行取回撤销重做任务
支持流程跟踪图，流程状态展现，流转信息，任务历史，任务分配信息
支持一个流程模型挂接多个业务单据，如某公司 8 种费用审批流程，表单不一样，但流程相同
支持一个表单挂接多个流程环节，以表单角度去管理流程，方便业务理解
支持全局表单，用于流程全局表单配置，目前支持内置表单、url 表单。如果不配置则发起流程会提示错误。
支持节点表单，节点表单配置。如果不配置默认使用全局表单。
流程事件脚本在线编写，包括：流程启动、完成、取消；任务分配、创建、结束等
流程脚本管理（Groovy、Beetl），在线编辑、自动完成、脚本测试、多语言脚本模板维护
我的待办任务处理，我的已办任务、我创建的任务查询、流程跟踪、审批记录查询
流程管控，在无关联表单情况下流程调试，如流程发起、挂起；流程定义、实例、任务等查询；任务办理，重定位等
支持流程组件标签定义（流程按钮、意见审批、下一步流程信息等）快速与自定义的业务表单建立关系。
支持版本化管理流程，新调整的流程业务不影响正在运行，未结束的流程继续流转。
支持任务加签、催办任务、传阅任务、流程委托设置、流水号管理、常用语管理

#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 使用说明

1.  cwgo语法
```shell
cwgo  model --db_type mysql --out_dir ./backend/pkg/dao --dsn "root:root@tcp(localhost:3306)/weflow?charset=utf8&parseTime=True&loc=Local"
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
