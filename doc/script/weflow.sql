create table `flow_def_info` (
                                 `id` int(11) not null comment '主键',
                                 `name` varchar(100) collate utf8mb4_bin not null default '' comment '流程定义名称',
                                 `flow_def_id` varchar(32) collate utf8mb4_bin not null default '' comment '流程定义id',
                                 `state` smallint(4) not null comment '状态',
                                 `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                 `create_user` varchar(50) collate utf8mb4_bin not null comment '创建人',
                                 `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                 `update_user` varchar(50) collate utf8mb4_bin default null comment '更新人',
                                 primary key (`id`)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_bin comment='流程定义信息表';

create table `flow_inst_info` (
                                  `id` int(11) not null comment '主键',
                                  `name` varchar(100) collate utf8mb4_bin not null default '' comment '流程实例名称',
                                  `flow_inst_id` varchar(32) collate utf8mb4_bin not null default '' comment '流程实例id',
                                  `state` smallint(4) not null comment '状态',
                                  `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                  `create_user` varchar(50) collate utf8mb4_bin not null comment '创建人',
                                  `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                  `update_user` varchar(50) collate utf8mb4_bin default null comment '更新人',
                                  primary key (`id`)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_bin comment='流程实例信息表';

create table `flow_inst_task` (
                                  `id` int(11) not null comment '主键',
                                  `name` varchar(100) collate utf8mb4_bin not null default '' comment '流程实例任务名称',
                                  `flow_inst_id` varchar(32) collate utf8mb4_bin not null default '' comment '流程实例id',
                                  `flow_inst_task_id` varchar(32) collate utf8mb4_bin not null default '' comment '流程实例任务id',
                                  `state` smallint(4) not null comment '状态',
                                  `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                  `create_user` varchar(50) collate utf8mb4_bin not null comment '创建人',
                                  `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                  `update_user` varchar(50) collate utf8mb4_bin default null comment '更新人',
                                  primary key (`id`)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_bin comment='流程实例任务表';

create table `flow_task_operator` (
                                      `id` int(11) not null comment '主键',
                                      `name` varchar(100) collate utf8mb4_bin not null default '' comment '流程实例任务名称',
                                      `flow_inst_id` varchar(32) collate utf8mb4_bin not null default '' comment '流程实例id',
                                      `flow_inst_task_id` varchar(32) collate utf8mb4_bin not null default '' comment '流程实例任务id',
                                      `task_operator_id` varchar(32) collate utf8mb4_bin not null default '' comment '任务id',
                                      `user_id` varchar(32) collate utf8mb4_bin not null default '' comment '用户id',
                                      `state` smallint(4) not null comment '状态',
                                      `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                      `create_user` varchar(50) collate utf8mb4_bin not null comment '创建人',
                                      `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                      `update_user` varchar(50) collate utf8mb4_bin default null comment '更新人',
                                      primary key (`id`)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_bin comment='流程任务用户操作表';


create table `form_def_info` (
                                 `id` int(11) not null comment '主键',
                                 `name` varchar(100) collate utf8mb4_bin not null default '' comment '表单定义名称',
                                 `form_def_id` varchar(32) collate utf8mb4_bin not null default '' comment '表单定义id',
                                 `state` smallint(4) not null comment '状态',
                                 `create_time` timestamp not null default current_timestamp on update current_timestamp comment '创建时间',
                                 `create_user` varchar(50) collate utf8mb4_bin not null comment '创建人',
                                 `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
                                 `update_user` varchar(50) collate utf8mb4_bin default null comment '更新人',
                                 primary key (`id`)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_bin comment='表单定义信息表';