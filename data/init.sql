CREATE TABLE `user` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(60) NOT NULL COMMENT '名称',
    `nickname` varchar(60) NOT NULL DEFAULT '' COMMENT '昵称',
    `sex` tinyint NOT NULL DEFAULT '1' COMMENT '性别，1男/2女/3其他',
    `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
    `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime COMMENT '更新时间',
    `deleted_at` datetime COMMENT '删除时间',
    `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态，0禁用/1锁定/2启用',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unq_username` (`username`),
    KEY `idx_status` (`status`),
    KEY `idx_mobile` (`mobile`),
    KEY `idx_create` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

CREATE TABLE `user_profile` (
     `user_id` int unsigned NOT NULL COMMENT '用户ID',
     `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '头像',
     `birthday` varchar(12) NOT NULL DEFAULT '' COMMENT '生日',
     `login_time` datetime COMMENT '上次登入时间',
     `login_num` int NOT NULL DEFAULT 0 COMMENT '登入次数',
     `extend` json COMMENT '扩展',
     PRIMARY KEY (`user_id`),
     KEY `idx_login` (`login_time`,`login_num`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户扩展';

CREATE TABLE `user_token` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `tokenable_type` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
    `tokenable_id` int unsigned NOT NULL,
    `token` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
    `last_used_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unq_token` (`token`),
    KEY `idx_tokenable` (`tokenable_type`,`tokenable_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='第三方认证';

CREATE TABLE `user_oauth` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `user_id` int unsigned NOT NULL COMMENT '用户ID',
    `oauth_type` varchar(20) NOT NULL default '' COMMENT '登入类型，wx_official/wx_mini/qq/weibo/',
    `oauth_id` varchar(60) NOT NULL default '' COMMENT '第三方 uid 、openid 等',
    `union_id` varchar(60) NOT NULL default '' COMMENT '微信同一主体下的Unionid',
    `credential` varchar(60) NOT NULL default '' COMMENT '凭证',
    `created_at` datetime COMMENT '创建时间',
    `expired_at` datetime COMMENT '过期时间',
    PRIMARY KEY (`id`),
    KEY `idx_user` (`user_id`),
    KEY `idx_auth` (`oauth_type`,`oauth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='第三方认证';

CREATE TABLE `casbin_rule` (
    `id` int NOT NULL AUTO_INCREMENT,
    `ptype` varchar(120) DEFAULT NULL,
    `v0` varchar(120) DEFAULT NULL,
    `v1` varchar(120) DEFAULT NULL,
    `v2` varchar(120) DEFAULT NULL,
    `v3` varchar(120) DEFAULT NULL,
    `v4` varchar(120) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `ptype` (`ptype`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8