

create table `sso_application` (
                                   `id` bigint not null auto_increment  comment '唯一id' ,
                                   `app_name` varchar(64) character set utf8mb4 collate utf8mb4_bin not null comment '系统名称',
                                   `app_instance_id` varchar(32) character set utf8mb4 collate utf8mb4_bin not null comment '系统唯一id',
                                   `app_owner` varchar(32) character set utf8mb4 collate utf8mb4_bin default null comment '系统所有者',
                                   `status` smallint(6) not null comment '状态',
                                   `create_time` bigint(20) not null comment '创建时间',
                                   `create_user_id` varchar(32) character set utf8mb4 collate utf8mb4_bin not null comment '创建人',
                                   `last_update_time` bigint(20) default null comment '最近更新时间',
                                   `last_update_user_id` varchar(32) character set utf8mb4 collate utf8mb4_bin default null comment '最近更新人',
                                   `app_code` varchar(32) character set utf8mb4 collate utf8mb4_bin not null comment '系统编码',
                                   `app_owner_email` varchar(127) character set utf8mb4 collate utf8mb4_bin default null comment '系统所有者邮箱',
                                   `app_remark` varchar(255) character set utf8mb4 collate utf8mb4_bin default null comment '描述',
                                   `app_type` smallint(6) not null comment '类型',
                                   primary key (`id`),
                                   unique key `app_code` (`app_code`),
                                   unique key `app_instance_id` (`app_instance_id`),
                                   key `idx_res_application` (`create_time`)
) engine=innodb default charset=utf8mb4;

create table `sso_application_auth_ref` (
                                            `id` bigint not null auto_increment  comment '唯一id' ,
                                            `create_time` bigint(20) not null,
                                            `create_user_id` varchar(32) not null,
                                            `last_update_time` bigint(20) default null,
                                            `last_update_user_id` varchar(32) default null,
                                            `application_id` varchar(32) not null,
                                            `auth_app_id` varchar(32) not null,
                                            `status` int(11) not null,
                                            primary key (`id`),
                                            unique key `idx_res_app_auth` (`application_id`,`auth_app_id`)
) engine=innodb default charset=utf8mb4;

create table `sso_application_logs` (
                                        `id` bigint not null auto_increment  comment '唯一id' ,
                                        `create_time` bigint(20) not null,
                                        `create_user_id` varchar(32) not null,
                                        `last_update_time` bigint(20) default null,
                                        `last_update_user_id` varchar(32) default null,
                                        `application_id` varchar(32) not null,
                                        `auth_app_ids` longtext,
                                        `opt_type` varchar(6) default null,
                                        `status` int(11) not null,
                                        primary key (`id`),
                                        key `idx_res_app_logs` (`create_time`)
) engine=innodb default charset=utf8mb4;

create table `sso_application_provide` (
                                           `id` bigint not null auto_increment  comment '唯一id' ,
                                           `create_time` bigint(20) not null,
                                           `create_user_id` varchar(32) not null,
                                           `last_update_time` bigint(20) default null,
                                           `last_update_user_id` varchar(32) default null,
                                           `application_id` varchar(32) not null,
                                           `service_code` varchar(32) not null,
                                           `icon` varchar(64) default null,
                                           `open` smallint(6) default null,
                                           `remark` varchar(255) default null,
                                           `service_name` varchar(127) default null,
                                           `service_path` varchar(255) default null,
                                           `status` smallint(6) not null,
                                           primary key (`id`),
                                           key `idx_res_app_privide` (`create_time`)
) engine=innodb default charset=utf8mb4;


create table `sso_application_service` (
                                           `id` bigint not null auto_increment  comment '唯一id' ,
                                           `create_time` bigint(20) not null,
                                           `create_user_id` varchar(32) not null,
                                           `last_update_time` bigint(20) default null,
                                           `last_update_user_id` varchar(32) default null,
                                           `application_id` varchar(32) not null,
                                           `service_name` varchar(255) not null,
                                           `service_code` varchar(32) not null,
                                           `remark` varchar(255) default null,
                                           `status` smallint(6) not null,
                                           primary key (`id`),
                                           unique key `service_code` (`service_code`),
                                           key `idx_res_app_service` (`create_time`)
) engine=innodb default charset=utf8;
create table `sso_auth_resource_ref` (
                                         `id` bigint not null auto_increment  comment '唯一id' ,
                                         `authorization_id` varchar(32) character set utf8mb4 collate utf8mb4_bin not null comment '权限id',
                                         `resource_id` varchar(32) character set utf8mb4 collate utf8mb4_bin not null comment '资源id',
                                         `type` smallint(6) not null comment '0：数据权限项，1：资源',
                                         `status` smallint(6) not null comment '0：有效，1：超级管理员或父级置为失效，2：权限拥有者置为失效',
                                         `create_time` bigint(20) not null comment '创建时间',
                                         `create_user_id` varchar(32) character set utf8mb4 collate utf8mb4_bin not null comment '创建用户',
                                         `last_update_time` bigint(20) default null comment '上次修改时间',
                                         `last_update_user_id` varchar(32) character set utf8mb4 collate utf8mb4_bin default null comment '上次修改用户',
                                         `application_id` varchar(32) character set utf8mb4 collate utf8mb4_bin default null comment '应用id',
                                         primary key (`id`),
                                         unique key `idx_res_auth_res` (`authorization_id`,`resource_id`,`type`) using btree
) engine=innodb default charset=utf8mb4;
create table `sso_authorization` (
                                     `id` bigint not null auto_increment  comment '唯一id' ,
                                     `name` varchar(255) collate utf8mb4_bin not null comment '名称',
                                     `parent_id` varchar(32) collate utf8mb4_bin default null comment '父id',
                                     `remark` varchar(255) collate utf8mb4_bin default null comment '描述',
                                     `status` smallint(6) not null comment '状态',
                                     `cust_id` varchar(32) collate utf8mb4_bin default null,
                                     `create_time` bigint(20) not null comment '创建时间',
                                     `create_user_id` varchar(32) collate utf8mb4_bin not null comment '创建用户',
                                     `last_update_time` bigint(20) default null comment '上次修改时间',
                                     `last_update_user_id` varchar(32) collate utf8mb4_bin default null comment '上次修改用户',
                                     `application_id` varchar(32) collate utf8mb4_bin default null comment '应用id',
                                     `code` varchar(32) collate utf8mb4_bin not null comment 'code',
                                     `order_by` smallint(6) default null comment '排序字段',
                                     primary key (`id`),
                                     unique key `code` (`code`),
                                     key `idx_res_authorization` (`create_time`)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_bin;

create table `sso_enterprise` (
                                  `id` bigint not null auto_increment  comment '唯一id' ,
                                  `create_time` bigint(20) not null,
                                  `create_user_id` varchar(32) not null,
                                  `last_update_time` bigint(20) default null,
                                  `last_update_user_id` varchar(32) default null,
                                  `employees_num` varchar(32) default null,
                                  `enter_address` varchar(64) default null,
                                  `enter_contacts` varchar(32) default null,
                                  `enter_name` varchar(64) default null,
                                  `enter_trade` varchar(32) default null,
                                  `enter_license` varchar(32) default null,
                                  primary key (`id`),
                                  key `idx_res_enterprise` (`create_time`)
) engine=innodb default charset=utf8mb4;
create table `sso_group` (
                             `id` bigint not null auto_increment  comment '唯一id' ,
                             `cust_id` varchar(32) default null,
                             `create_time` bigint(20) not null,
                             `create_user_id` varchar(32) not null,
                             `last_update_time` bigint(20) default null,
                             `last_update_user_id` varchar(32) default null,
                             `application_id` varchar(32) default null,
                             `code` varchar(64) not null,
                             `name` varchar(255) default null,
                             `remark` varchar(255) default null,
                             `status` smallint(6) not null,
                             primary key (`id`),
                             unique key `code` (`code`),
                             key `idx_res_group` (`create_time`)
) engine=innodb default charset=utf8mb4;

create table `sso_group_role_ref` (
                                      `id` bigint not null auto_increment  comment '唯一id' ,
                                      `create_time` bigint(20) not null,
                                      `create_user_id` varchar(32) not null,
                                      `last_update_time` bigint(20) default null,
                                      `last_update_user_id` varchar(32) default null,
                                      `application_id` varchar(32) default null,
                                      `group_id` varchar(32) not null,
                                      `role_id` varchar(32) not null,
                                      primary key (`id`),
                                      unique key `idx_res_group_role` (`group_id`,`role_id`)
) engine=innodb default charset=utf8mb4;
create table `sso_group_user_ref` (
                                      `id` bigint not null auto_increment  comment '唯一id' ,
                                      `create_time` bigint(20) not null,
                                      `create_user_id` varchar(32) not null,
                                      `last_update_time` bigint(20) default null,
                                      `last_update_user_id` varchar(32) default null,
                                      `application_id` varchar(32) default null,
                                      `group_id` varchar(32) not null,
                                      `user_id` varchar(32) not null,
                                      primary key (`id`),
                                      unique key `idx_res_group_user` (`group_id`,`user_id`)
) engine=innodb default charset=utf8mb4;
create table `sso_organization` (
                                    `id` bigint not null auto_increment  comment '唯一id' ,
                                    `name` varchar(255) collate utf8mb4_bin default null comment '机构名称',
                                    `parent_id` varchar(32) collate utf8mb4_bin default null comment '父id',
                                    `create_time` bigint(20) not null comment '创建时间',
                                    `create_user_id` varchar(32) collate utf8mb4_bin not null comment '创建用户',
                                    `last_update_time` bigint(20) default null comment '上次修改时间',
                                    `last_update_user_id` varchar(32) collate utf8mb4_bin default null comment '上次修改用户',
                                    `address` varchar(255) collate utf8mb4_bin default null comment '地址',
                                    `application_id` varchar(32) collate utf8mb4_bin default null comment '应用id',
                                    `code` varchar(32) collate utf8mb4_bin not null comment 'code',
                                    `cust_id` varchar(32) collate utf8mb4_bin default null,
                                    `fax` varchar(24) collate utf8mb4_bin default null,
                                    `link_man_dept` varchar(64) collate utf8mb4_bin default null comment '联系人部门',
                                    `link_man_email` varchar(255) collate utf8mb4_bin default null comment '联系人邮箱',
                                    `link_man_name` varchar(255) collate utf8mb4_bin default null comment '联系人名称',
                                    `link_man_pos` varchar(64) collate utf8mb4_bin default null,
                                    `link_man_tel` varchar(64) collate utf8mb4_bin default null comment '联系人手机号',
                                    `rank` smallint(6) default '0' comment '等级',
                                    `remark` varchar(255) collate utf8mb4_bin default null comment '描述',
                                    `status` smallint(6) not null default '0' comment '状态',
                                    `telephone` varchar(32) collate utf8mb4_bin default null,
                                    `type` smallint(6) not null default '0' comment '0：经营单位；1：机构/部门',
                                    `zip_code` varchar(16) collate utf8mb4_bin default null,
                                    `path` varchar(1024) collate utf8mb4_bin default null,
                                    `organization_number` varchar(6) collate utf8mb4_bin default null comment '机构编号',
                                    `admin_user_id` varchar(32) collate utf8mb4_bin default null comment '管理员id',
                                    `area` smallint(6) not null comment '区域（0：区域，1：非区域）',
                                    `order_by` smallint(6) default null comment '排序字段',
                                    primary key (`id`),
                                    unique key `code` (`code`),
                                    key `idx_res_organization` (`create_time`)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_bin;
create table `sso_resource` (
                                `id` bigint not null auto_increment  comment '唯一id' ,
                                `name` varchar(255) character set utf8mb4 collate utf8mb4_bin not null comment '资源名称',
                                `parent_id` varchar(32) character set utf8mb4 collate utf8mb4_bin default null comment '父id',
                                `status` smallint(6) not null comment '状态',
                                `create_time` bigint(20) not null comment '创建时间',
                                `create_user_id` varchar(32) character set utf8mb4 collate utf8mb4_bin not null comment '创建用户id',
                                `last_update_time` bigint(20) default null comment '上次修改时间',
                                `last_update_user_id` varchar(32) character set utf8mb4 collate utf8mb4_bin default null comment '上次修改用户',
                                `application_id` varchar(32) character set utf8mb4 collate utf8mb4_bin default null comment '应用id',
                                `code` varchar(32) character set utf8mb4 collate utf8mb4_bin not null comment 'code',
                                `order_by` smallint(6) default null comment '排序',
                                `picture_name` varchar(255) character set utf8mb4 collate utf8mb4_bin default null comment '图标',
                                `remark` varchar(255) character set utf8mb4 collate utf8mb4_bin default null comment '描述',
                                `type` smallint(6) not null comment '类型',
                                `url` varchar(255) character set utf8mb4 collate utf8mb4_bin default null comment '路由',
                                `show_status` smallint(6) default null comment '显示状态，0：显示，1：不显示',
                                `have_menu` smallint(6) default null comment '是否拥有，0：非拥有，1：拥有',
                                primary key (`id`),
                                unique key `code` (`code`),
                                key `idx_res_resource` (`create_time`)
) engine=innodb default charset=utf8mb4;

create table `sso_role` (
                            `id` bigint not null auto_increment  comment '唯一id' ,
                            `parent_id` varchar(32) collate utf8mb4_bin default null comment '父id',
                            `name` varchar(255) collate utf8mb4_bin default null comment '角色名称',
                            `cust_id` varchar(32) collate utf8mb4_bin default null,
                            `application_id` varchar(32) collate utf8mb4_bin not null comment '应用id',
                            `remark` varchar(255) collate utf8mb4_bin default null comment '描述',
                            `create_time` bigint(20) not null comment '创建时间',
                            `code` varchar(64) collate utf8mb4_bin not null comment 'code',
                            `create_user_id` varchar(32) collate utf8mb4_bin default null comment '创建用户',
                            `last_update_time` bigint(20) default null comment '上次修改时间',
                            `last_update_user_id` varchar(32) collate utf8mb4_bin default null comment '上次修改用户',
                            `status` smallint(6) not null comment '角色状态',
                            `order_by` smallint(6) default null comment '排序字段',
                            `type` smallint(6) not null comment '0：功能角色，1：数据角色',
                            primary key (`id`),
                            unique key `code` (`code`),
                            key `idx_res_role` (`create_time`)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_bin;

create table `sso_role_authorization_ref` (
                                              `id` bigint not null auto_increment  comment '唯一id' ,
                                              `create_time` bigint(20) not null,
                                              `create_user_id` varchar(32) not null,
                                              `last_update_time` bigint(20) default null,
                                              `last_update_user_id` varchar(32) default null,
                                              `application_id` varchar(32) default null,
                                              `authorization_id` varchar(32) not null,
                                              `role_id` varchar(32) not null,
                                              `status` smallint(6) not null comment '0：有效，1：超级管理员或父级置为失效，2：权限拥有者置为失效',
                                              primary key (`id`),
                                              unique key `idx_res_role_auth` (`role_id`,`authorization_id`)
) engine=innodb default charset=utf8mb4;
create table `sso_role_resource_ref` (
                                         `id` bigint not null auto_increment  comment '唯一id' ,
                                         `create_time` bigint(20) not null comment '创建时间',
                                         `create_user_id` varchar(32) collate utf8mb4_bin not null comment '创建用户',
                                         `last_update_time` bigint(20) default null comment '上次修改时间',
                                         `last_update_user_id` varchar(32) collate utf8mb4_bin default null comment '上次修改用户',
                                         `application_id` varchar(32) collate utf8mb4_bin default null comment '应用id',
                                         `role_id` varchar(32) collate utf8mb4_bin not null comment '角色名称',
                                         `resource_id` varchar(32) collate utf8mb4_bin not null comment '资源id',
                                         `status` smallint(6) not null comment '0：有效，1：超级管理员或父级置为失效，2：权限拥有者置为失效',
                                         `type` smallint(6) not null comment '0：数据权限项，1：资源',
                                         primary key (`id`),
                                         unique key `idx_res_auth_res` (`role_id`,`resource_id`,`type`) using btree
) engine=innodb default charset=utf8mb4 collate=utf8mb4_bin;
create table `sso_user` (
                            `id` bigint not null auto_increment  comment '唯一id' ,
                            `create_time` bigint(20) not null comment '创建时间',
                            `create_user_id` varchar(32) character set utf8mb4 collate utf8mb4_bin not null comment '创建用户',
                            `platfrom_userid` varchar(10) character set utf8mb4 collate utf8mb4_bin not null,
                            `last_update_time` bigint(20) default null comment '上次修改时间',
                            `last_update_user_id` varchar(32) character set utf8mb4 collate utf8mb4_bin default null comment '上次修改用户',
                            `company_id` varchar(32) character set utf8mb4 collate utf8mb4_bin default null,
                            `organization_id` varchar(32) character set utf8mb4 collate utf8mb4_bin default null,
                            `name` varchar(255) character set utf8mb4 collate utf8mb4_bin default null comment '用户名',
                            `user_type` smallint(6) not null comment '用户类型',
                            `node_id` varchar(32) character set utf8mb4 collate utf8mb4_bin default null,
                            `sex` smallint(6) default null comment '性别',
                            `birthday` varchar(32) character set utf8mb4 collate utf8mb4_bin default null comment '生日',
                            `email` varchar(255) character set utf8mb4 collate utf8mb4_bin default null comment '邮箱',
                            `position` varchar(255) character set utf8mb4 collate utf8mb4_bin default null comment '位置',
                            `remark` varchar(255) character set utf8mb4 collate utf8mb4_bin default null comment '描述',
                            `tel_phone` varchar(32) character set utf8mb4 collate utf8mb4_bin default null comment '手机号',
                            `user_code` varchar(32) character set utf8mb4 collate utf8mb4_bin not null,
                            `user_level` smallint(6) default null,
                            `face_url` mediumblob comment '头像',
                            `plat_person_id` varchar(32) character set utf8mb4 collate utf8mb4_bin default null comment '工号',
                            `sign` varchar(255) character set utf8mb4 collate utf8mb4_bin default null,
                            `brief_introd` varchar(255) character set utf8mb4 collate utf8mb4_bin default null,
                            primary key (`id`),
                            unique key `user_code` (`user_code`),
                            unique key `platfrom_userid` (`platfrom_userid`),
                            unique key `tel_phone_uk` (`tel_phone`) using btree,
                            key `idx_res_user` (`create_time`)
) engine=innodb default charset=utf8mb4;
create table `sso_user_account` (
                                    `id` bigint not null auto_increment  comment '唯一id' ,
                                    `account_expired` bigint(20) default null,
                                    `account_locked` smallint(6) not null,
                                    `create_time` bigint(20) not null comment '创建时间',
                                    `credential` varchar(512) default null,
                                    `credential_expired` bigint(20) default null,
                                    `identifier` varchar(32) not null comment '账号',
                                    `identity_type` smallint(6) not null comment '账号类型 1:username,2:email,3:phone,4:wechat,5:qq',
                                    `last_login_ip` varchar(15) default null comment '上次登录ip',
                                    `last_login_time` varchar(32) default null comment '上次登录时间',
                                    `last_update_time` bigint(20) default null comment '上次修改时间',
                                    `login_mode` smallint(6) not null,
                                    `status` smallint(6) not null comment '用户状态',
                                    `user_id` varchar(32) not null comment '用户id',
                                    `credential_update_time` bigint(20) default null comment '上次密码修改时间',
                                    `first_pwd_error_time` bigint(20) default null comment '第一次密码错误时间',
                                    `login_pwd_error_time` smallint(6) default null comment '登录密码错误的次数',
                                    primary key (`id`),
                                    unique key `identifier` (`identifier`),
                                    key `idx_res_account` (`last_update_time`)
) engine=innodb default charset=utf8mb4;

create table `sso_user_role_ref` (
                                     `id` bigint not null auto_increment  comment '唯一id' ,
                                     `application_id` varchar(32) collate utf8mb4_bin default null,
                                     `create_time` bigint(20) not null comment '创建时间',
                                     `create_user_id` varchar(32) collate utf8mb4_bin not null comment '创建用户',
                                     `last_update_time` bigint(20) default null comment '上次修改时间',
                                     `last_update_user_id` varchar(32) collate utf8mb4_bin default null comment '上次修改用户',
                                     `role_id` varchar(32) collate utf8mb4_bin not null comment '角色id',
                                     `user_id` varchar(32) collate utf8mb4_bin not null comment '用户id',
                                     `type` smallint(6) not null comment '0：功能角色，1：数据角色',
                                     primary key (`id`),
                                     unique key `idx_res_user_role` (`user_id`,`role_id`,`type`) using btree
) engine=innodb default charset=utf8mb4 collate=utf8mb4_bin;




