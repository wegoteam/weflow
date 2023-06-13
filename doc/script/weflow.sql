
drop table if exists model_detail;
create table `model_detail` (
                                `id` bigint(20) not null auto_increment comment '唯一id',
                                `model_id` varchar(32) not null default '' comment '模板id',
                                `model_title` varchar(256) not null default '' comment '模板标题',
                                `process_def_id` varchar(32) not null default '' comment '流程定义id',
                                `form_def_id` varchar(32) not null default '' comment '表单定义id',
                                `model_group_id` varchar(32) not null default '' comment '模版组id',
                                `icon_url` varchar(256) not null default '' comment 'icon图标地址',
                                `status` tinyint(4) not null default '1' comment '模板状态【1：草稿；2：发布；3：停用】默认草稿',
                                `remark` varchar(512) not null default '' comment '描述',
                                `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                `create_user` varchar(50) not null default '' comment '创建人',
                                `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                `update_user` varchar(50) default '' comment '更新人',
                                primary key (`id`),
                                key `model_detail_id_index` (`model_id`,`status`)
) engine=innodb default charset=utf8mb4 comment='模板详情表';

drop table if exists model_ext_prop;
create table `model_ext_prop` (
                                  `id` bigint(20) not null auto_increment comment '唯一id',
                                  `model_id` varchar(32) not null default '' comment '模板id',
                                  `notice_url` varchar(256) not null default '' comment '回调通知推送url',
                                  `title_props` varchar(1024) not null default '' comment '标题配置',
                                  primary key (`id`),
                                  unique key `model_id_index` (`model_id`)
) engine=innodb default charset=utf8mb4 comment='模板详情扩展属性表';

drop table if exists model_version;
create table `model_version` (
                                 `id` bigint(20) not null auto_increment comment '唯一id',
                                 `model_id` varchar(32) not null default '' comment '模板id',
                                 `model_title` varchar(256) not null default '' comment '模板版本标题',
                                 `version_id` varchar(32) not null default '' comment '版本id',
                                 `process_def_id` varchar(32) not null default '' comment '流程定义id',
                                 `form_def_id` varchar(32) not null default '' comment '表单定义id',
                                 `use_status` tinyint(4) not null default '1' comment '使用状态【1：使用；2：未使用】',
                                 `remark` varchar(512) not null default '' comment '描述',
                                 `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                 `create_user` varchar(50) not null default '' comment '创建人',
                                 `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                 `update_user` varchar(50) default '' comment '更新人',
                                 `notice_url` varchar(256) not null default '' comment '回调通知推送url',
                                 `title_props` varchar(1024) not null default '' comment '标题配置',
                                 primary key (`id`),
                                 unique key `model_and_version_unique` (`model_id`,`version_id`)
) engine=innodb default charset=utf8mb4 comment='模板版本表';

drop table if exists model_auth;
create table `model_auth` (
                              `id` bigint(20) not null auto_increment comment '唯一id',
                              `model_id` varchar(32) not null default '' comment '模板id',
                              `auth_obj_type` tinyint(4) not null default '1' comment '授权对象类型【人员：1；部门（包含子部门）：2；部门（不含子部门）：3；角色（包含子角色）：4；角色（不包含子角色）：5】',
                              `obj_id` varchar(32) not null comment '授权对象id;根据授权对象类型取值',
                              `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                              `create_user` varchar(50) not null default '' comment '创建人',
                              `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                              `update_user` varchar(50) default '' comment '更新人',
                              primary key (`id`),
                              key `model_id_index` (`model_id`)
) engine=innodb default charset=utf8mb4 comment='模板授权表';


drop table if exists process_def_info;
create table `process_def_info` (
                                    `id` bigint(20) not null auto_increment comment '唯一id',
                                    `process_def_id` varchar(32) not null default '' comment '流程定义id',
                                    `process_def_name` varchar(128) not null default '' comment '流程定义名称',
                                    `status` tinyint(4) not null default '1' comment '状态【1：草稿；2：发布可用；3：停用】',
                                    `remark` varchar(512) not null default '' comment '描述',
                                    `struct_data` text comment '流程定义数据',
                                    `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                    `create_user` varchar(50) not null default '' comment '创建人',
                                    `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                    `update_user` varchar(50) default '' comment '更新人',
                                    primary key (`id`),
                                    key `process_def_id_index` (`process_def_id`)
) engine=innodb default charset=utf8mb4 comment='流程定义表';

drop table if exists process_def_node;
create table `process_def_node` (
                                    `id` bigint(20) not null auto_increment comment '唯一id',
                                    `process_def_id` varchar(32) not null default '' comment '流程定义id',
                                    `node_id` varchar(32) not null default '' comment '节点id',
                                    `node_model` tinyint(4) not null default '1' comment '节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】',
                                    `node_name` varchar(128) not null default '' comment '节点名称',
                                    `parent_id` varchar(32) not null default '' comment '节点父ID',
                                    `approve_type` tinyint(4) not null default '1' comment '审批类型【人工审批：1；自动通过：2；自动拒绝】默认人工审批1',
                                    `none_handler` tinyint(4) not null default '1' comment '审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1',
                                    `appoint_handler` varchar(128) not null default '' comment '审批人为空时指定审批人ID',
                                    `handle_mode` tinyint(4) not null default '2' comment '审批方式【依次审批：1、会签（需要完成人数的审批人同意或拒绝才可完成节点）：2、或签（其中一名审批人同意或拒绝即可）：3】默认会签2',
                                    `finish_mode` int(11) not null default '0' comment '完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0）',
                                    `branch_mode` tinyint(4) not null default '2' comment '分支执行方式【单分支：1；多分支：2】默认多分支2',
                                    `default_branch` int(11) not null default '0' comment '单分支处理需要默认分支，在条件优先级无法处理时候执行默认分支，取值分支下标',
                                    `branch_level` int(11) not null default '0' comment '优先级，分支执行方式为多分支处理方式无优先级应为0',
                                    `condition_group` varchar(4000) not null default '' comment '条件组前端描述展示条件组',
                                    `condition_expr` varchar(4000) not null default '' comment '条件组解析后的表达式',
                                    `remark` varchar(512) not null default '' comment '节点描述',
                                    `pre_nodes` varchar(2000) not null default '' comment '上节点ID集合,多个用逗号隔开',
                                    `next_nodes` varchar(2000) not null default '' comment '下节点ID集合,多个用逗号隔开',
                                    `last_nodes` varchar(2000) not null default '' comment '尾节点ID集合,多个用逗号隔开',
                                    `index` int(11) not null default '0' comment '节点下标',
                                    `branch_index` int(11) not null default '0' comment '分支节点下标',
                                    primary key (`id`),
                                    key `process_def_id_index` (`process_def_id`)
) engine=innodb default charset=utf8mb4 comment='流程定义节点信息表';

drop table if exists process_def_node_formper;
create table `process_def_node_formper` (
                                         `id` bigint(20) not null auto_increment comment '唯一id',
                                         `process_def_id` varchar(32) not null default '' comment '流程定义id',
                                         `node_id` varchar(32) not null default '' comment '节点id',

                                         `elemId` varchar(128) not null default '' comment '处理人对象id;处理对象的id，根据处理人类型区分，如果操作员id、部门id等',
                                         `elemPId` varchar(128) not null default '' comment '处理人对象id;处理对象的id，根据处理人类型区分，如果操作员id、部门id等',
                                         `per` int(11) not null default '1' comment '处理人顺序;正序排序',

                                         primary key (`id`),
                                         key `process_def_id_index` (`process_def_id`,`node_id`)
) engine=innodb default charset=utf8mb4 comment='流程定义节点表单权限表';

drop table if exists process_def_node_user;
create table `process_def_node_user` (
                                         `id` bigint(20) not null auto_increment comment '唯一id',
                                         `process_def_id` varchar(32) not null default '' comment '流程定义id',
                                         `node_id` varchar(32) not null default '' comment '节点id',
                                         `type` tinyint(4) not null default '1' comment '常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】',
                                         `strategy` tinyint(4) not null default '1' comment '处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】',
                                         `node_user_name` varchar(128) not null default '' comment '处理人名称',
                                         `node_user_id` varchar(128) not null default '' comment '处理人id',
                                         `sort` int(11) not null default '1' comment '处理人顺序;正序排序',
                                         `obj` varchar(4000) not null default '' comment '扩展字段，设计中可忽略',
                                         `relative` varchar(4000) not null default '' comment '相对发起人的直属主管，设计中可忽略',
                                         primary key (`id`),
                                         key `process_def_id_index` (`process_def_id`,`node_id`)
) engine=innodb default charset=utf8mb4 comment='流程定义节点用户表';

drop table if exists form_def_info;
create table `form_def_info` (
                                 `id` bigint(20) not null auto_increment comment '唯一id',
                                 `form_def_id` varchar(32) not null default '' comment '表单模板id;唯一id',
                                 `form_def_name` varchar(128) not null default '' comment '表单名称',
                                 `html_content` text comment '表单定义结构化数据',
                                 `html_page_url` varchar(256) not null default '' comment 'html页面访问路径;html页面文件访问路径',
                                 `status` tinyint(4) not null default '1' comment '模板状态【1：草稿；2：发布；3：停用】',
                                 `remark` varchar(512) not null default '' comment '表单描述',
                                 `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                 `create_user` varchar(50) not null default '' comment '创建人',
                                 `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                 `update_user` varchar(50) default '' comment '更新人',
                                 primary key (`id`),
                                 key `form_def_id_index` (`form_def_id`)
) engine=innodb default charset=utf8mb4 comment='表单定义信息表';

drop table if exists form_def_element;
create table `form_def_element` (
                                    `id` bigint(20) not null auto_increment comment '唯一id',
                                    `form_def_id` varchar(32) not null default '' comment '表单模板id;唯一id',
                                    `ele_name` varchar(128) not null default '' comment '元素名称',
                                    `ele_id` varchar(128) not null default '' comment '元素标识id;表单内元素唯一id，可通过该id关联表单内元素',
                                    `label_type` varchar(16) not null default 'input' comment '元素标签类型;html标签：input、select等',
                                    `ele_type` tinyint(4) not null default '1' comment '元素类型【1：文本；2：复选：3：下拉；4：单选】',
                                    `ele_default` varchar(1024) not null default '' comment '元素默认值',
                                    `remark` varchar(512) not null default '' comment '元素描述',
                                    `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                    `create_user` varchar(50) not null default '' comment '创建人',
                                    `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                    `update_user` varchar(50) default '' comment '更新人',
                                    primary key (`id`),
                                    key `form_def_id_index` (`form_def_id`)
) engine=innodb default charset=utf8mb4 comment='表单定义元素表';


drop table if exists inst_task_detail;
create table `inst_task_detail` (
                                    `id` bigint(20) not null auto_increment comment '唯一id',
                                    `inst_task_id` varchar(32) not null default '' comment '实例任务id',
                                    `model_id` varchar(32) not null default '' comment '模板id',
                                    `process_def_id` varchar(32) not null default '' comment '流程定义id',
                                    `form_def_id` varchar(32) not null default '' comment '表单定义id',
                                    `version_id` varchar(32) not null default '' comment '版本id',
                                    `task_name` varchar(512) not null default '' comment '实例任务名称',
                                    `status` tinyint(4) not null default '1' comment '任务状态【1：创建中(草稿)；2：进行中； 3：终止； 4：完成； 5：挂起；6：回退】',
                                    `remark` varchar(512) not null default '' comment '描述',
                                    `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                    `create_user_id` varchar(32) not null default '' comment '创建人id',
                                    `create_user_name` varchar(128) not null default '' comment '创建人名称',
                                    `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                    `update_user_id` varchar(32) default '' comment '更新人id',
                                    `update_user_name` varchar(128) default '' comment '更新人名称',
                                    `start_time` timestamp not null default current_timestamp on update current_timestamp comment '发起时间',
                                    `end_time` timestamp not null default current_timestamp on update current_timestamp comment '结束时间',
                                    primary key (`id`),
                                    unique key `inst_user_task_unique` (`inst_task_id`),
                                    key `inst_task_id_index` (`inst_task_id`)
) engine=innodb default charset=utf8mb4 comment='执行实例任务信息表';

drop table if exists inst_node_task;
create table `inst_node_task` (
                                  `id` bigint(20) not null auto_increment comment '唯一id',
                                  `inst_task_id` varchar(32) not null default '' comment '实例任务id',
                                  `node_task_id` varchar(32) not null default '' comment '节点任务id',
                                  `node_id` varchar(32) not null default '' comment '节点id',
                                  `parent_id` varchar(32) not null default '' comment '父节点id',
                                  `node_model` tinyint(4) not null default '1' comment '节点模型【1：开始节点；2：审批节点；3：办理节点；4：抄送节点；5：自定义节点；6：条件节点；7：分支节点；8：汇聚节点；9：结束节点】',
                                  `node_name` varchar(128) not null default '' comment '节点名称',
                                  `approve_type` tinyint(4) not null default '1' comment '审批类型【人工审批：1；自动通过：2；自动拒绝】默认人工审批1',
                                  `none_handler` tinyint(4) not null default '1' comment '审批人为空时【自动通过：1；自动转交管理员：2；指定审批人：3】默认自动通过1',
                                  `appoint_handler` varchar(128) not null default '' comment '审批人为空时指定审批人ID',
                                  `handle_mode` tinyint(4) not null default '2' comment '审批方式【依次审批：1；会签（需要完成人数的审批人同意或拒绝才可完成节点）：2；或签（其中一名审批人同意或拒绝即可）：3】默认会签2',
                                  `finish_mode` int(11) not null default '0' comment '完成人数：依次审批默认0所有人不可选人，会签默认0所有人（可选人大于0），或签默认1一个人（可选人大于0）',
                                  `branch_mode` tinyint(4) not null default '2' comment '分支执行方式【单分支：1；多分支：2】默认多分支2',
                                  `default_branch` int(11) not null default '0' comment '单分支处理需要默认分支，在条件优先级无法处理时候执行默认分支，取值分支下标',
                                  `branch_level` int(11) not null default '0' comment '优先级，分支执行方式为多分支处理方式无优先级应为0',
                                  `condition_group` varchar(4000) not null default '' comment '条件组前端描述展示条件组',
                                  `condition_expr` varchar(4000) not null default '' comment '条件组解析后的表达式',
                                  `remark` varchar(512) not null default '' comment '节点描述',
                                  `status` tinyint(4) not null default '1' comment '任务状态【1：未开始；2：处理中；3：完成；4：回退；5：终止；6：不通过】',
                                  `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                  `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                  primary key (`id`),
                                  unique key `inst_user_task_unique` (`inst_task_id`,`node_task_id`),
                                  key `inst_node_task_id_index` (`inst_task_id`,`node_task_id`)
) engine=innodb default charset=utf8mb4 comment='实例节点任务表';

drop table if exists inst_node_task_formper;
create table `inst_node_task_formper` (
                                            `id` bigint(20) not null auto_increment comment '唯一id',
                                            `inst_task_id` varchar(32) not null default '' comment '实例任务id',
                                            `node_task_id` varchar(32) not null default '' comment '节点任务id',
                                            `node_id` varchar(32) not null default '' comment '节点id',
                                            `elemId` varchar(128) not null default '' comment '表单元素ID',
                                            `elemPId` varchar(128) not null default '' comment '表单元素父ID',
                                            `per` tinyint(4) not null default '2' comment '表单权限【可编辑：1；只读：2；隐藏：3;必填：4】默认只读2',

                                            primary key (`id`),
                                            key `process_def_id_index` (`inst_task_id`,`node_task_id`)
) engine=innodb default charset=utf8mb4 comment='实例节点任务表单权限表';

drop table if exists inst_user_task;
create table `inst_user_task` (
                                  `id` bigint(20) not null auto_increment comment '唯一id',
                                  `inst_task_id` varchar(32) not null default '' comment '实例任务id',
                                  `node_task_id` varchar(32) not null default '' comment '节点任务id',
                                  `node_id` varchar(32) not null default '' comment '节点任务id',
                                  `user_task_id` varchar(32) not null default '' comment '处理人任务id',
                                  `type` tinyint(4) not null default '1' comment '常用审批人【指定成员：1；发起人自己：2；发起人自选：3：角色：4；部门：5】主管（相对岗位）【直属主管：1；部门主管：2；连续多级主管：3；部门控件对应主管：4】其他【表单人员控件：1；部门控件：2；角色控件：3】',
                                  `strategy` tinyint(4) not null default '1' comment '处理人策略【常用审批人：1；主管（相对岗位）：2；其他：3】',
                                  `node_user_name` varchar(128) not null default '' comment '处理人名称',
                                  `node_user_id` varchar(128) not null default '' comment '处理人id',
                                  `sort` int(11) not null default '1' comment '处理人顺序;正序排序',
                                  `obj` varchar(4000) not null default '' comment '扩展字段，设计中可忽略',
                                  `relative` varchar(4000) not null default '' comment '相对发起人的直属主管，设计中可忽略',
                                  `status` tinyint(4) not null default '1' comment '实例用户任务状态【1：处理中；2：完成（同意）；3：不通过（不同意）；4：回退；5：终止】',
                                  `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                  `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                  `handle_time` timestamp null default current_timestamp on update current_timestamp comment '处理时间',
                                  `op_user_id` varchar(128) not null default '' comment '操作用户id',
                                  `op_user_name` varchar(128) not null default '' comment '操作用户名称',
                                  `opinion` tinyint(4) not null default '1' comment '任务处理意见【1：未发表；2：同意；3：不同意；4：xxx】',
                                  `opinion_desc` varchar(3000) not null default '' comment '处理意见描述',
                                  primary key (`id`),
                                  unique key `inst_user_task_unique` (`inst_task_id`,`user_task_id`),
                                  key `inst_user_task_index` (`inst_task_id`,`node_task_id`)
) engine=innodb default charset=utf8mb4 comment='实例用户任务表';


drop table if exists inst_user_task_opinion;
create table `inst_user_task_opinion` (
                                          `id` bigint(20) not null auto_increment comment '唯一id',
                                          `inst_task_id` varchar(32) not null default '' comment '实例任务id',
                                          `node_task_id` varchar(32) not null default '' comment '节点任务id',
                                          `user_task_id` varchar(32) not null default '' comment '用户任务id',
                                          `node_id` varchar(32) not null default '' comment '节点id',
                                          `opinion_id` varchar(32) not null default '' comment '意见id',
                                          `opinion` tinyint(4) not null default '1' comment '处理意见【1：未处理；2：同意；3：不同意；4：回退；5：终止；6：撤回】',
                                          `opinion_desc` varchar(3000) not null default '' comment '处理意见描述',
                                          `op_user_id` varchar(128) not null default '' comment '操作用户id',
                                          `op_user_name` varchar(128) not null default '' comment '操作用户名称',
                                          `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                          `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                          `opinion_time` timestamp null default current_timestamp on update current_timestamp comment '发表意见时间',
                                          primary key (`id`),
                                          key `inst_user_task_id_index` (`inst_task_id`)
) engine=innodb default charset=utf8mb4 comment='实例用户任务意见表';

drop table if exists inst_task_param;
create table `inst_task_param` (
                                   `id` bigint(20) not null auto_increment comment '唯一id',
                                   `inst_task_id` varchar(32) not null default '' comment '实例任务id',
                                   `param_id` varchar(128) not null default '' comment '参数id;参数的唯一标识',
                                   `param_name` varchar(128) not null default '' comment '参数名称;参数的名称',
                                   `param_value` varchar(4000) not null default '' comment '参数值',
                                   `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                   `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                   `param_data_type` varchar(12) not null default 'string' comment '参数数据类型【string：字符串；int ：整形数值；float：浮点型数值；object：对象；array：数组；decimal：金额；long：长整型；table：表格；boolean：布尔】',
                                   `param_binary` blob comment '参数二进制值;可存二进制值对象',
                                   primary key (`id`),
                                   key `inst_task_id_index` (`inst_task_id`)
) engine=innodb default charset=utf8mb4 comment='实例任务参数表(备注可用mongodb扩展)';

drop table if exists inst_task_param_attr;
create table `inst_task_param_attr` (
                                        `id` bigint(20) not null auto_increment comment '唯一id',
                                        `inst_task_id` varchar(32) not null default '' comment '实例任务id',
                                        `param_id` varchar(128) not null default '' comment '参数id',
                                        `param_name` varchar(128) not null default '' comment '参数名称',
                                        `param_attr` varchar(4000) not null default '' comment '参数值',
                                        `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                        `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                        primary key (`id`),
                                        key `inst_task_id_index` (`inst_task_id`)
) engine=innodb default charset=utf8mb4 comment='实例任务参数表(备注可用mongodb扩展)';

drop table if exists inst_task_op_log;
create table `inst_task_op_log` (
                                    `id` bigint(20) not null auto_increment comment '唯一id',
                                    `inst_task_id` varchar(32) not null default '' comment '实例任务id',
                                    `node_id` varchar(32) not null default '' comment '节点任务id',
                                    `node_name` varchar(128) not null default '' comment '节点名称',
                                    `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                    `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                    `type` tinyint(4) not null default '1' comment '类型【1：节点；2：任务；3：其他】',
                                    `remark` varchar(1024) not null default '' comment '描述',
                                    primary key (`id`),
                                    key `inst_task_id_index` (`inst_task_id`)
) engine=innodb default charset=utf8mb4 comment='实例任务操作日志表';

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
                            primary key (id),
                            key `group_id_index` (`group_id`)
) engine=innodb default charset=utf8mb4 comment = '模型组表';

drop table if exists organization_info;
create table `organization_info` (
                                     `id` bigint(20) not null auto_increment comment '唯一id',
                                     `org_id` varchar(32) not null default '' comment '组id',
                                     `parent_id` varchar(32) not null default '' comment '组父id',
                                     `org_name` varchar(128) not null default '' comment '组名称',
                                     `status` tinyint(4) not null default '1' comment '状态【1：未启用；2：已启用；3：锁定；】',
                                     `remark` varchar(512) not null default '' comment '描述',
                                     `create_user` varchar(50) not null default '' comment '创建人',
                                     `update_user` varchar(50) not null default '' comment '更新人',
                                     `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                     `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                     primary key (`id`),
                                     key `organization_org_id_index` (`org_id`)
) engine=innodb default charset=utf8mb4 comment='组织信息表';

drop table if exists role_info;
create table `role_info` (
                             `id` bigint(20) not null auto_increment comment '唯一id',
                             `role_id` varchar(32) not null default '' comment '角色id',
                             `parent_id` varchar(32) not null default '' comment '角色父id',
                             `role_name` varchar(128) not null default '' comment '角色名称',
                             `status` tinyint(4) not null default '1' comment '状态【1：未启用；2：已启用；3：锁定；】',
                             `remark` varchar(512) not null default '' comment '描述',
                             `create_user` varchar(50) not null default '' comment '创建人',
                             `update_user` varchar(50) not null default '' comment '更新人',
                             `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                             `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                             primary key (`id`),
                             key `role_id_index` (`role_id`)
) engine=innodb default charset=utf8mb4 comment='角色信息表';

drop table if exists user_info;
create table `user_info` (
                             `id` bigint(20) not null auto_increment comment '唯一id',
                             `user_id` varchar(32)  not null default '' comment '用户id',
                             `user_name` varchar(128)  not null default '' comment '用户名称',
                             `password` varchar(128)  not null default '' comment '密码',
                             `phone` varchar(11)  not null default '' comment '手机号',
                             `email` varchar(50)  not null default '' comment '邮箱',
                             `org_id` varchar(32) not null default '' comment '组织id',
                             `status` tinyint(4) not null default '1' comment '状态【1：未启用；2：已启用；3：锁定；】',
                             `remark` varchar(512) not null default '' comment '描述',
                             `create_user` varchar(50) not null default '' comment '创建人',
                             `update_user` varchar(50) not null default '' comment '更新人',
                             `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                             `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                             primary key (`id`),
                             key `user_id_index` (`user_id`)
) engine=innodb default charset=utf8mb4 comment='用户信息表';

drop table if exists user_role_link;
create table `user_role_link` (
                                  `id` bigint(20) not null auto_increment comment '唯一id',
                                  `role_id` varchar(32) not null default '' comment '角色id',
                                  `user_id` varchar(32) not null default '' comment '用户id',
                                  `status` tinyint(4) not null default '1' comment '状态【1：未启用；2：已启用】',
                                  `remark` varchar(512) not null default '' comment '描述',
                                  `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                  `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                  primary key (`id`),
                                  key `user_role_id_index` (`user_id`,`role_id`)
) engine=innodb default charset=utf8mb4 comment='用户角色关联表';

-- 流程定义
INSERT INTO `process_def_info` (`process_def_id`, `process_def_name`, `status`, `remark`, `struct_data`, `create_time`, `create_user`, `update_time`, `update_user`) VALUES ('1640993392605401001', '测试流程定义', 1, '测试流程定义', '[
    {
        "nodeModel":1,
        "nodeName":"发起人",
        "nodeId":"1640993392605401001",
        "parentId":""
    },
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
            "type": 1,
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
    },
    {
        "nodeModel":7,
        "nodeName":"分支节点",
        "nodeId":"1640993392605401003",
        "parentId":"",
        "branchMode":2,
        "defaultBranch":0,
        "children":[
            [
                {
                    "nodeModel":6,
                    "nodeName":"条件1",
                    "nodeId":"1640993392605401004",
                    "parentId":"1640993392605401003",
                    "level":0,
                    "conditionGroup":"",
                    "conditionExpr":"",
                    "conditionVal":""
                },
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
                        "type": 1,
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
            ],
            [
                {
                    "nodeModel":6,
                    "nodeName":"条件2",
                    "nodeId":"1640993392605401006",
                    "parentId":"1640993392605401003",
                    "level":0,
                    "conditionGroup":"",
                    "conditionExpr":"",
                    "conditionVal":""
                },
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
                        "type": 1,
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
            ]
        ]
    },
    {
        "nodeModel":8,
        "nodeName":"分支汇聚",
        "nodeId":"1640993392605401008",
        "parentId":""
    },
    {
        "nodeModel":9,
        "nodeName":"流程结束",
        "nodeId":"1640993392605401009",
        "parentId":""
    }
]', '2023-04-12 15:19:02', '', '2023-04-12 15:19:02', '');

INSERT INTO `process_def_info` (`process_def_id`, `process_def_name`, `status`, `remark`, `struct_data`, `create_time`, `create_user`, `update_time`, `update_user`) VALUES ('1640993392605401002', '测试流程定义2', 1, '测试流程定义2', '[
    {
        "nodeModel":1,
        "nodeName":"发起人",
        "nodeId":"1640993392605401001",
        "parentId":""
    },
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
            "type": 1,
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
    },
    {
        "nodeModel":7,
        "nodeName":"分支节点",
        "nodeId":"1640993392605401003",
        "parentId":"",
        "branchMode":2,
        "defaultBranch":0,
        "children":[
            [
                {
                    "nodeModel":6,
                    "nodeName":"条件1",
                    "nodeId":"1640993392605401004",
                    "parentId":"1640993392605401003",
                    "level":0,
                    "conditionGroup":"",
                    "conditionExpr":"",
                    "conditionVal":""
                },
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
                        "type": 1,
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
            ],
            [
                {
                    "nodeModel":6,
                    "nodeName":"条件2",
                    "nodeId":"1640993392605401006",
                    "parentId":"1640993392605401003",
                    "level":0,
                    "conditionGroup":"",
                    "conditionExpr":"",
                    "conditionVal":""
                },
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
                        "type": 1,
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
            ]
        ]
    },
    {
        "nodeModel":8,
        "nodeName":"分支汇聚",
        "nodeId":"1640993392605401008",
        "parentId":""
    },
    {
        "nodeModel":9,
        "nodeName":"流程结束",
        "nodeId":"1640993392605401009",
        "parentId":""
    }
]', '2023-04-12 15:19:02', '', '2023-04-12 15:19:02', '');

INSERT INTO `form_def_info` (`form_def_id`, `form_def_name`, `html_content`, `html_page_url`, `status`, `remark`, `create_time`, `create_user`, `update_time`, `update_user`) VALUES ('1681467241063842637', '测试表单', NULL, '', 1, '', '2023-05-24 13:36:54', 'xuch01', '2023-05-24 13:36:54', 'xuch01');

INSERT INTO `organization_info` (`org_id`, `parent_id`, `org_name`, `status`, `remark`, `create_user`, `update_user`, `create_time`, `update_time`) VALUES ('420627966730317', '0', 'wego', 1, '', 'xuch01', 'xuch01', '2023-05-24 11:35:13', '2023-05-24 11:35:13');
INSERT INTO `role_info` (`role_id`, `parent_id`, `role_name`, `status`, `remark`, `create_user`, `update_user`, `create_time`, `update_time`) VALUES ('420627966730315', '0', '管理员', 1, '', 'xuch01', 'xuch01', '2023-05-24 11:34:16', '2023-05-24 11:34:16');
INSERT INTO `user_info` ( `user_id`, `user_name`, `password`, `phone`, `email`, `org_id`, `status`, `remark`, `create_user`, `update_user`, `create_time`, `update_time`) VALUES ('547', 'xuch01', 'xuch01', '13800138000', '13800138000@163.com', '420627966730317', 1, '', 'xuch01', 'xuch01', '2023-05-24 10:54:02', '2023-05-24 10:54:02');
INSERT INTO `user_info` ( `user_id`, `user_name`, `password`, `phone`, `email`, `org_id`, `status`, `remark`, `create_user`, `update_user`, `create_time`, `update_time`) VALUES ('420627966730316', 'xuch02', 'xuch02', '13800138001', '13800138001@163.com', '420627966730317', 1, '', 'xuch02', 'xuch02', '2023-05-24 10:54:04', '2023-05-24 10:54:04');
INSERT INTO `user_role_link` (`role_id`, `user_id`, `status`, `remark`, `create_time`, `update_time`) VALUES ('420627966730315', '547', 1, '', '2023-05-23 17:54:35', '2023-05-23 17:54:35');
INSERT INTO `user_role_link` (`role_id`, `user_id`, `status`, `remark`, `create_time`, `update_time`) VALUES ('420627966730315', '420627966730316', 1, '', '2023-05-23 17:54:54', '2023-05-23 17:54:54');

INSERT INTO `model_detail` (`model_id`, `model_title`, `process_def_id`, `form_def_id`, `model_group_id`, `icon_url`, `status`, `remark`, `create_time`, `create_user`, `update_time`, `update_user`) VALUES ('420915317174341', '测试模板', '1640993392605401001', '1681467241063842637', '1665958955971575812', '', 1, '', '2023-05-24 13:37:14', 'xuch01', '2023-05-24 13:37:14', 'xuch01');
INSERT INTO `model_group` (`group_id`, `group_name`, `remark`, `create_user`, `update_user`, `create_time`, `update_time`) VALUES ('1665958955971575812', '测试组', '', 'xuch01', 'xuch01', '2023-05-24 13:32:46', '2023-05-24 13:32:46');
INSERT INTO `model_version` (`model_id`, `model_title`, `version_id`, `process_def_id`, `form_def_id`, `use_status`, `remark`, `create_time`, `create_user`, `update_time`, `update_user`, `notice_url`, `title_props`) VALUES ('420915317174341', '测试模板', '1681335332954505235', '1640993392605401001', '1681467241063842637',  1, '', '2023-05-24 13:37:09', '', '2023-05-24 13:37:09', '', '', '');

