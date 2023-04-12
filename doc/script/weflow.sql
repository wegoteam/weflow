
drop table if exists model_detail;
create table model_detail(
                             `id` bigint not null auto_increment  comment '唯一id' ,
                             `model_id` varchar(32) not null  default '' comment '模板id' ,
                             `model_name` varchar(128) not null  default '' comment '模板名称' ,
                             `model_title` varchar(256) not null  default '' comment '模板标题' ,
                             `process_def_id` varchar(32) not null  default '' comment '流程定义id' ,
                             `form_def_id` varchar(32) not null  default '' comment '表单定义id' ,
                             `model_group_id` varchar(32) not null  default '' comment '模版组id' ,
                             `status` tinyint not null  default 1 comment '模板状态【1：草稿；2：发布；3：停用】' ,
                             `remark` varchar(512) not null  default '' comment '描述' ,
                             `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                             `create_user` varchar(50) not null default '' comment '创建人',
                             `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                             `update_user` varchar(50) default '' comment '更新人',
                             `notice_url` varchar(256) not null  default '' comment '回调通知推送url' ,
                             `title_props` varchar(1024) not null  default '' comment '标题配置' ,
                             primary key (id)
)engine=innodb default charset=utf8mb4 comment = '模板详情表';

drop table if exists model_version;
create table model_version(
                                `id` bigint not null auto_increment  comment '唯一id' ,
                                `model_id` varchar(32) not null  default '' comment '模板id' ,
                                `version_id` varchar(32) not null  default '' comment '版本id' ,
                                `version_num` varchar(32) not null  default '' comment '版本号' ,
                                `model_title` varchar(256) not null  default '' comment '模板标题' ,
                                `process_def_id` varchar(32) not null  default '' comment '流程定义id' ,
                                `form_def_id` varchar(32) not null  default '' comment '表单定义id' ,
                                `table_info` varchar(1024) not null  default '' comment '表单数据库表' ,
                                `use_status` tinyint not null  default 1 comment '使用状态【1：非当前使用；2：当前使用】' ,
                                `remark` varchar(512) not null  default '' comment '描述' ,
                                `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                `create_user` varchar(50) not null default '' comment '创建人',
                                `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                `update_user` varchar(50) default '' comment '更新人',
                                `notice_url` varchar(256) not null  default '' comment '回调通知推送url' ,
                                `title_props` varchar(1024) not null  default '' comment '标题配置' ,
                                primary key (id)
)engine=innodb default charset=utf8mb4 comment = '模板版本表';

drop table if exists model_auth;
create table model_auth(
                                  `id` bigint not null auto_increment  comment '唯一id' ,
                                  `model_id` varchar(32) not null  default '' comment '模板id' ,
                                  `auth_obj_type` tinyint not null  default 1 comment '授权对象类型【1：操作员；2：部门（包含子部门）；3：部门（不含子部门）】' ,
                                  `obj_id` varchar(32) not null  comment '授权对象id;根据授权对象类型取值' ,
                                  `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                  `create_user` varchar(50) not null default '' comment '创建人',
                                  `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                  `update_user` varchar(50) default '' comment '更新人',
                                  primary key (id)
)engine=innodb default charset=utf8mb4 comment = '模板授权表';

drop table if exists process_def_info;
create table process_def_info(
                            `id` bigint not null auto_increment  comment '唯一id' ,
                            `process_def_id` varchar(32) not null  default '' comment '流程定义id' ,
                            `process_def_name` varchar(128) not null  default '' comment '流程定义名称' ,
                            `status` tinyint not null  default 1 comment '状态【1：草稿；2：发布可用；3：停用】' ,
                            `remark` varchar(512) not null  default '' comment '描述' ,
                            `struct_data` text  comment '流程结构化数据' ,
                            `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                            `create_user` varchar(50) not null default '' comment '创建人',
                            `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                            `update_user` varchar(50) default '' comment '更新人',
                            primary key (id)
)engine=innodb default charset=utf8mb4 comment = '流程定义表';

drop table if exists process_def_node_info;
create table process_def_node_info(
                             `id` bigint not null auto_increment  comment '唯一id' ,
                             `process_def_id` varchar(32) not null  default '' comment '流程定义id' ,
                             `node_id` varchar(32) not null  default '' comment '节点id' ,
                             `node_type` tinyint not null  default 1 comment '节点类型;1：正常节点；2：开始节点；3：结束节点；4：汇聚节点；5：条件节点' ,
                             `node_name` varchar(128) not null  default '' comment '节点名称' ,
                             `forward_mode` tinyint not null  default 1 comment '进行模式【1：并行 2：串行】' ,
                             `complete_conn` int not null  default 0 comment '节点完成条件;通过的人数，0表示所有人通过，节点才算完成' ,
                             `permission_mode` tinyint not null  default 1 comment '权限模式【1：协同 2：知会 3：审批】' ,
                             `allow_add` tinyint not null  default 1 comment '允许加签【1：不能加签；2：允许加签】' ,
                             `process_mode` tinyint not null  default 1 comment '处理模式【1：人工； 2：自动】' ,
                             `bus_id` varchar(32) not null  default '' comment '业务id' ,
                             `bus_type` varchar(128) not null  default '' comment '业务类型' ,
                             `time_limit` int not null  default 0 comment '处理期限时长;单位秒，0表示无期限；' ,
                             `conn_data` varchar(3000) not null  default '' comment '条件表达式;条件节点才有条件表达式' ,
                             `form_per_data` varchar(3000) not null  default '' comment '表单权限数据;节点表单权限配置，json格式' ,
                             `remark` varchar(512) not null  default '' comment '节点描述' ,
                             `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                             `create_user` varchar(50) not null default '' comment '创建人',
                             `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                             `update_user` varchar(50) default '' comment '更新人',
                             primary key (id)
) engine=innodb default charset=utf8mb4 comment = '流程定义节点信息表';

drop table if exists process_def_node_handler;
create table process_def_node_handler(
                                `id` bigint not null auto_increment  comment '唯一id' ,
                                `process_def_id` varchar(32) not null  default '' comment '流程定义id' ,
                                `node_id` varchar(32) not null  default '' comment '节点id' ,
                                `handler_name` varchar(128) not null  default '' comment '处理人名称' ,
                                `handler_type` tinyint not null  default 1 comment '处理人类型【1：用户；2：部门；3：相对岗位；4：表单控件；5：部门岗位】' ,
                                `handler_id` varchar(128) not null  default '' comment '处理人对象id;处理对象的id，根据处理人类型区分，如果操作员id、部门id等' ,
                                `handler_sort` int not null  default 1 comment '处理人顺序;正序排序' ,
                                `obj_data` varchar(3000) not null  default '' comment '对象数据;依据处理人类型取值，相对岗位和表单控件使用该字段存json数据' ,
                                `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                `create_user` varchar(50) not null default '' comment '创建人',
                                `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                `update_user` varchar(50) default '' comment '更新人',
                                primary key (id)
)  engine=innodb default charset=utf8mb4  comment = '流程定义节点处理人表';

drop table if exists form_def_detail;
create table form_def_detail(
                             `id` bigint not null auto_increment  comment '唯一id' ,
                             `form_def_id` varchar(32) not null  default '' comment '表单模板id;唯一id' ,
                             `form_def_name` varchar(128) not null  default '' comment '表单名称' ,
                             `html_content` text  comment '表单定义结构化数据' ,
                             `html_page_url` varchar(256) not null  default '' comment 'html页面访问路径;html页面文件访问路径' ,
                             `status` tinyint not null  default 1 comment '模板状态【1：草稿；2：发布；3：停用】' ,
                             `remark` varchar(512) not null  default '' comment '表单描述' ,
                             `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                             `create_user` varchar(50) not null default '' comment '创建人',
                             `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                             `update_user` varchar(50) default '' comment '更新人',
                             primary key (id)
)  engine=innodb default charset=utf8mb4  comment = '表单定义信息表';

drop table if exists form_def_element;
create table form_def_element(
                                `id` bigint not null auto_increment  comment '唯一id' ,
                                `form_def_id` varchar(32) not null  default '' comment '表单模板id;唯一id' ,
                                `ele_name` varchar(128) not null  default '' comment '元素名称' ,
                                `ele_id` varchar(128) not null  default '' comment '元素标识id;表单内元素唯一id，可通过该id关联表单内元素' ,
                                `label_type` varchar(16) not null  default 'input' comment '元素标签类型;html标签：input、select等' ,
                                `ele_type` tinyint not null  default 1 comment '元素类型【1：文本；2：复选：3：下拉；4：单选】' ,
                                `ele_default` varchar(1024) not null  default '' comment '元素默认值' ,
                                `remark` varchar(512) not null  default '' comment '元素描述' ,
                                `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                `create_user` varchar(50) not null default '' comment '创建人',
                                `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                `update_user` varchar(50) default '' comment '更新人',
                                primary key (id)
)  engine=innodb default charset=utf8mb4  comment = '表单定义元素表';


drop table if exists inst_task_detail;
create table inst_task_detail(
                             `id` bigint not null auto_increment  comment '唯一id' ,
                             `inst_task_id` varchar(32) not null  default '' comment '实例任务id' ,
                             `process_def_id` varchar(32) not null  default '' comment '流程定义id' ,
                             `model_id` varchar(32) not null  default '' comment '模板id' ,
                             `form_def_id` varchar(32) not null  default '' comment '表单定义id' ,
                             `version_id` bigint not null  default 0 comment '版本id' ,
                             `version_num` varchar(24) not null  default '' comment '版本号' ,
                             `create_src` int not null  default 1 comment '创建来源【1：系统发起；2：API发起】' ,
                             `task_name` varchar(512) not null  default '' comment '实例任务名称' ,
                             `status` tinyint not null  default 1 comment '任务状态【1：创建中；2：进行中； 3：终止； 4：完成； 5：挂起；6：草稿】' ,
                             `remark` varchar(512) not null  default '' comment '描述' ,
                             `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                             `create_user` varchar(50) not null default '' comment '创建人',
                             `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                             `update_user` varchar(50) default '' comment '更新人',
                             `start_time` timestamp not null default current_timestamp on update current_timestamp comment '发起时间' ,
                             `end_time` timestamp not null default current_timestamp on update current_timestamp comment '结束时间' ,
                             `source_id` varchar(128) not null  default '' comment '来源id，界面发起人或者api对应的应用' ,
                             primary key (id)
) engine=innodb default charset=utf8mb4  comment = '执行实例任务信息表';

drop table if exists inst_node_task;
create table inst_node_task(
                             `id` bigint not null auto_increment  comment '唯一id' ,
                             `inst_task_id` varchar(32) not null  default '' comment '实例任务id' ,
                             `node_task_id` varchar(32) not null  default '' comment '节点任务id' ,
                             `node_id` varchar(32) not null  default '' comment '节点任务id' ,
                             `parent_id` varchar(32) not null  default '' comment '父节点id' ,
                             `node_type` tinyint not null  default 1 comment '节点类型【1：正常节点；2：开始节点；3：结束节点；4：汇聚节点；5：条件节点；6：分支节点】' ,
                             `node_name` varchar(128) not null  default '' comment '节点名称' ,
                             `forward_mode` tinyint not null  default 1 comment '进行模式【1：并行 2：串行】' ,
                             `complete_conn` int not null  default 0 comment '节点完成条件;通过的人数，0表示所有人通过，节点才算完成' ,
                             `permission_mode` tinyint not null  default 1 comment '权限模式【1：协同 2：知会 3：审批；4：业务】' ,
                             `allow_add` tinyint not null  default 1 comment '允许加签【1：不能加签；2：允许加签】' ,
                             `process_mode` tinyint not null  default 1 comment '处理模式【1：人工； 2：自动；3：自动转人工】' ,
                             `time_limit` bigint not null  default 0 comment '处理期限;格式：yyyymmddhhmm 可直接指定到期限的具体时间，期限支持到分钟； 0表示无期限' ,
                             `conn_data` varchar(3000) not null  default '' comment '条件数据;前端生成，json格式' ,
                             `form_per_data` varchar(3000) not null  default '' comment '表单权限数据;节点表单权限配置，json格式' ,
                             `status` tinyint not null  default 1 comment '任务状态【0：未开始；1：处理中；2：完成；3：回退；4：终止；5：条件验证通过；6：条件验证不通过】' ,
                             `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                             `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                             primary key (id)
) engine=innodb default charset=utf8mb4  comment = '实例节点任务表';

drop table if exists inst_handler_task;
create table inst_handler_task(
                                `id` bigint not null auto_increment  comment '唯一id' ,
                                `inst_task_id` varchar(32) not null  default '' comment '实例任务id' ,
                                `node_task_id` varchar(32) not null  default '' comment '节点任务id' ,
                                `node_id` varchar(32) not null  default '' comment '节点任务id' ,
                                `handler_task_id` varchar(32) not null  default '' comment '处理人任务id' ,
                                `node_handler_id` varchar(32) not null  default '' comment '节点处理人id' ,
                                `node_handler_name` varchar(128) not null  default '' comment '处理人名称' ,
                                `node_handler_type` tinyint not null  default 1 comment '处理人类型【1：操作员；2：部门；3：相对岗位；4：表单控件；5：角色；6：岗位；7：组织；8：自定义】' ,
                                `op_origin` tinyint not null  default 1 comment '操作来源【1：正常；2：加签】' ,
                                `time_limit` bigint not null  default 0 comment '处理期限;格式：yyyymmddhhmm 可直接指定到期限的具体时间，期限支持到分钟； 0表示无期限' ,
                                `status` tinyint not null  default 1 comment '任务状态【1：处理中；2：完成；3：回退；4：终止】' ,
                                `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                `handle_time` timestamp null default current_timestamp on update current_timestamp comment '处理时间',
                                `op_user_id` varchar(128) not null  default '' comment '操作用户id' ,
                                `op_user_name` varchar(128) not null  default '' comment '操作用户名称' ,
                                `handler_sort` int not null  default 1 comment '处理人排序;处理人当前的处理排序' ,
                                `opinion` tinyint not null  default 1 comment '处理意见【1：未发表；2：已阅；3：同意；4：不同意】' ,
                                `opinion_desc` varchar(3000) not null  default '' comment '处理意见描述' ,
                                primary key (id)
) engine=innodb default charset=utf8mb4 comment = '实例节点处理人任务表';

drop table if exists inst_handler_task_opinion;
create table inst_handler_task_opinion(
                                   `id` bigint not null auto_increment  comment '唯一id' ,
                                   `inst_task_id` varchar(32) not null  default '' comment '实例任务id' ,
                                   `node_task_id` varchar(32) not null  default '' comment '节点任务id' ,
                                   `node_id` varchar(32) not null  default '' comment '节点id' ,
                                   `opinion_id` varchar(32) not null  default '' comment '意见id' ,
                                   `opinion` tinyint not null  default 1 comment '处理意见【1：未处理；2：已阅；3：同意；4：不同意；5：回退；6：终止】',
                                   `opinion_desc` varchar(3000) not null  default '' comment '处理意见描述' ,
                                   `op_user_id` varchar(128) not null  default '' comment '操作用户id' ,
                                   `op_user_name` varchar(128) not null  default '' comment '操作用户名称' ,
                                   `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                   `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                   `opinion_time` timestamp null default current_timestamp on update current_timestamp comment '发表意见时间',
                                   primary key (id)
) engine=innodb default charset=utf8mb4 comment = '实例节点处理人任务意见表';

drop table if exists inst_task_param;
create table inst_task_param(
                              `id` bigint not null auto_increment  comment '唯一id' ,
                              `inst_task_id` varchar(32) not null  default '' comment '实例任务id' ,
                              `param_id` varchar(128) not null  default '' comment '参数id;参数的唯一标识' ,
                              `param_name` varchar(128) not null  default '' comment '参数名称;参数的名称' ,
                              `param_value` varchar(4000) not null  default '' comment '参数值' ,
                              `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                              `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                              `param_data_type` varchar(12) not null  default 'string' comment '参数数据类型【string：字符串；int ：整形数值；float：浮点型数值；object：对象；array：数组；decimal：金额；long：长整型；table：表格；boolean：布尔】' ,
                              `param_binary` blob(8000)  comment '参数二进制值;可存二进制值对象' ,
                              primary key (id)
) engine=innodb default charset=utf8mb4 comment = '实例任务参数表(备注可用MongoDB扩展)';

drop table if exists inst_task_param_attr;
create table inst_task_param_attr(
                                `id` bigint not null auto_increment  comment '唯一id' ,
                                `inst_task_id` varchar(32) not null  default '' comment '实例任务id' ,
                                `param_id` varchar(128) not null  default '' comment '参数id' ,
                                `param_name` varchar(128) not null  default '' comment '参数名称' ,
                                `param_attr` varchar(4000) not null  default '' comment '参数值' ,
                                `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                primary key (id)
) engine=innodb default charset=utf8mb4 comment = '实例任务参数表(备注可用MongoDB扩展)';

drop table if exists inst_task_op_log;
create table inst_task_op_log(
                                 `id` bigint not null auto_increment  comment '唯一id' ,
                                 `inst_task_id` varchar(32) not null  default '' comment '实例任务id' ,
                                 `node_id` varchar(32) not null  default '' comment '节点任务id' ,
                                 `node_name` varchar(128) not null  default '' comment '节点名称' ,
                                 `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                 `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                 `type` tinyint not null  default 1 comment '类型【1：节点；2：任务；3：其他】' ,
                                 `remark` varchar(1024) not null  default '' comment '描述' ,
                                 primary key (id)
) engine=innodb default charset=utf8mb4 comment = '实例任务操作日志表';

drop table if exists model_group;
create table model_group(
                            `id` bigint not null auto_increment  comment '唯一id' ,
                            `group_id` varchar(32) not null  default '' comment '组id' ,
                            `group_name` varchar(128) not null  default '' comment '组名称' ,
                            `remark` varchar(512) not null  default '' comment '描述' ,
                            `create_user` varchar(50) not null default '' comment '创建人',
                            `update_user` varchar(50) not null  default '' comment '更新人' ,
                            `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                            `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                            primary key (id)
) engine=innodb default charset=utf8mb4 comment = '模型组表';

drop table if exists organization_info;
create table organization_info(
                             `id` bigint not null auto_increment  comment '唯一id' ,
                             `org_id` varchar(32) not null  default '' comment '组id' ,
                             `parent_id` varchar(32) not null  default '' comment '组父id' ,
                             `org_name` varchar(128) not null  default '' comment '组名称' ,
                             `status` tinyint not null  default 1 comment '状态【1：未启用；2：已启用；3：锁定；】' ,
                             `remark` varchar(512) not null  default '' comment '描述' ,
                             `create_user` varchar(50) not null default '' comment '创建人',
                             `update_user` varchar(50) not null  default '' comment '更新人' ,
                             `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                             `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                             primary key (id)
) engine=innodb default charset=utf8mb4 comment = '组织信息表';

drop table if exists role_info;
create table role_info(
                              `id` bigint not null auto_increment  comment '唯一id' ,
                              `role_id` varchar(32) not null  default '' comment '角色id' ,
                              `parent_id` varchar(32) not null  default '' comment '角色父id' ,
                              `role_name` varchar(128) not null  default '' comment '角色名称' ,
                              `status` tinyint not null  default 1 comment '状态【1：未启用；2：已启用；3：锁定；】' ,
                              `remark` varchar(512) not null  default '' comment '描述' ,
                              `create_user` varchar(50) not null default '' comment '创建人',
                              `update_user` varchar(50) not null  default '' comment '更新人' ,
                              `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                              `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                              primary key (id)
) engine=innodb default charset=utf8mb4 comment = '角色信息表';

drop table if exists user_info;
create table user_info(
                          `id` bigint not null auto_increment  comment '唯一id' ,
                          `user_id` varchar(32) not null  default '' comment '用户id' ,
                          `user_name` varchar(128) not null  default '' comment '用户名称' ,
                          `password` varchar(128) not null  default '' comment '密码' ,
                          `phone` varchar(11) not null  default '' comment '手机号' ,
                          `email` varchar(50) not null  default '' comment '邮箱' ,
                          `org_id` varchar(32) not null  default '' comment '组织id' ,
                          `status` tinyint not null  default 1 comment '状态【1：未启用；2：已启用；3：锁定；】' ,
                          `remark` varchar(512) not null  default '' comment '描述' ,
                          `create_user` varchar(50) not null default '' comment '创建人',
                          `update_user` varchar(50) not null  default '' comment '更新人' ,
                          `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                          `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                          primary key (id)
)  comment = '用户信息表';

drop table if exists user_role_link;
create table user_role_link(
                          `id` bigint not null auto_increment  comment '唯一id' ,
                          `role_id` varchar(32) not null  default '' comment '角色id' ,
                          `user_id` varchar(32) not null  default '' comment '用户id' ,
                          `status` tinyint not null  default 1 comment '状态【1：未启用；2：已启用】' ,
                          `remark` varchar(512) not null  default '' comment '描述' ,
                          `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                          `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                          primary key (id)
) engine=innodb default charset=utf8mb4 comment = '用户角色关联表';

-- 流程定义
INSERT INTO `process_def_info` (`process_def_id`, `process_def_name`, `status`, `remark`, `struct_data`, `create_time`, `create_user`, `update_time`, `update_user`) VALUES ('1640993392605401001', '测试流程定义', 1, '测试流程定义', '[{\"nodeModel\":1,\"nodeName\":\"发起人\",\"nodeId\":\"1640993392605401088\",\"parentId\":\"\"},{\"nodeModel\":2,\"nodeName\":\"审批\",\"nodeId\":\"1640993449224310784\",\"parentId\":\"\",\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"2\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"2\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\"},{\"nodeModel\":6,\"nodeName\":\"分支节点\",\"nodeId\":\"1640993508049424384\",\"parentId\":\"\",\"children\":[[{\"nodeModel\":5,\"nodeName\":\"条件1\",\"nodeId\":\"1640993508049424385\",\"parentId\":\"1640993508049424384\"},{\"nodeModel\":2,\"nodeName\":\"审批\",\"nodeId\":\"1640993526328201216\",\"parentId\":\"1640993508049424384\",\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"2\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"2\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\"}],[{\"nodeModel\":5,\"nodeName\":\"条件2\",\"nodeId\":\"1640993508049424386\",\"parentId\":\"1640993508049424384\"},{\"nodeModel\":4,\"nodeName\":\"知会\",\"nodeId\":\"1640993535555670016\",\"parentId\":\"1640993508049424384\",\"perData\":\"{\\\\\\\"1640993433239818240\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"1640993434883985408\\\\\\\":\\\\\\\"1\\\\\\\"}\",\"handlerList\":\"[{\\\\\\\"handlerId\\\\\\\":\\\\\\\"547\\\\\\\",\\\\\\\"handlerName\\\\\\\":\\\\\\\"xuch01\\\\\\\",\\\\\\\"handlerType\\\\\\\":1,\\\\\\\"handlerSort\\\\\\\":1}]\"}]]},{\"nodeModel\":7,\"nodeName\":\"分支汇聚\",\"nodeId\":\"1640993508053618688\",\"parentId\":\"\"},{\"nodeModel\":8,\"nodeName\":\"流程结束\",\"nodeId\":\"1640993392605401089\",\"parentId\":\"\"}]', '2023-04-12 15:19:02', '', '2023-04-12 15:19:02', '');
