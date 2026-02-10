CREATE TABLE `user` (
    `id` bigint unsigned NOT NULL,
    `username` varchar(30) NOT NULL COMMENT '名称',
    `nickname` varchar(60) NOT NULL DEFAULT '' COMMENT '昵称',
    `sex` tinyint NOT NULL DEFAULT 0 COMMENT '性别，0未知/1男/2女/3其他',
    `password` varchar(64) NOT NULL DEFAULT '' COMMENT '密码',
    `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
    `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '状态，0禁用/1启用/2锁定',
    `created_at` datetime default current_timestamp,
    `updated_at` datetime default current_timestamp on update current_timestamp,
    `deleted_time` int default 0 COMMENT '删除时间戳',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_username` (`username`),
    INDEX `idx_status` (`status`),
    INDEX `idx_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

CREATE TABLE `user_profile` (
    `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
    `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '头像',
    `birthday` date COMMENT '出生日期',
    `login_at` datetime COMMENT '上次登入时间',
    `login_num` int NOT NULL DEFAULT 0 COMMENT '登入次数',
    `extend` varchar(1024) COMMENT '扩展',
    PRIMARY KEY (`user_id`),
    INDEX `idx_login` (`login_at`,`login_num`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户资料';

CREATE TABLE `user_oauth` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
    `oauth_type` varchar(20) NOT NULL default '' COMMENT '登入类型，wx_official/wx_mini/qq/weibo',
    `oauth_id` varchar(128) NOT NULL default '' COMMENT '第三方 uid 、openid 等',
    `union_id` varchar(128) NOT NULL default '' COMMENT '微信同一主体下的Unionid',
    `credential` varchar(128) NOT NULL default '' COMMENT '凭证',
    `created_at` datetime default current_timestamp,
    `expired_at` datetime,
    PRIMARY KEY (`id`),
    INDEX `idx_user` (`user_id`),
    INDEX `idx_auth` (`oauth_type`,`oauth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='第三方认证';

CREATE TABLE `user_token` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `token_type` varchar(30) NOT NULL,
    `token_id` bigint unsigned NOT NULL,
    `token` varchar(128) NOT NULL,
    `refresh_token` varchar(128) NOT NULL,
    `created_at` datetime default current_timestamp,
    `updated_at` datetime default current_timestamp on update current_timestamp,
    `expired_at` datetime,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unq_token` (`token`),
    INDEX `idx_token_info` (`token_type`,`token_id`),
    INDEX `idx_expired` (`expired_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户token';

-- template
CREATE TABLE `tablename` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_by` varchar(30) not null default '',
    `updated_by` varchar(30) not null default '',
    `created_at` datetime default current_timestamp,
    `updated_at` datetime default current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unq_token` (`token`),
    INDEX `idx_e` (`e`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='';
