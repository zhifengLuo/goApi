CREATE TABLE `users` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(60) NOT NULL COMMENT '名称',
    `nickname` varchar(60) NOT NULL DEFAULT '' COMMENT '昵称',
    `sex` tinyint NOT NULL DEFAULT '1' COMMENT '性别，1男/2女/3其他',
    `avatar` varchar(120) NOT NULL DEFAULT '' COMMENT '头像',
    `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
    `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime COMMENT '更新时间',
    `deleted_at` datetime COMMENT '删除时间',
    `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态，1启用/2禁用',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unq_username` (`username`),
    KEY `idx_status` (`status`),
    KEY `idx_mobile` (`mobile`),
    KEY `idx_create` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户';

CREATE TABLE `user_profile` (
     `user_id` int unsigned NOT NULL COMMENT '用户ID',
     `birthday` varchar(12) NOT NULL DEFAULT '' COMMENT '生日',
     `extend` json NOT NULL default '' COMMENT '扩展',
     PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户扩展';

CREATE TABLE `user_oauth` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `user_id` int unsigned NOT NULL COMMENT '用户ID',
    `oauth_type` varchar(12) NOT NULL default '' COMMENT '登入类型，wx_official/wx_mini/qq/weibo/',
    `oauth_id` varchar(60) NOT NULL default '' COMMENT '第三方 uid 、openid 等',
    `unionid` varchar(60) NOT NULL default '' COMMENT '微信同一主体下的Unionid',
    `credential` varchar(60) NOT NULL default '' COMMENT '凭证',
    `created_at` datetime COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `idx_user` (`user_id`),
    KEY `idx_auth` (`oauth_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='第三方认证';

CREATE TABLE `role` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(60) NOT NULL COMMENT '名称',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime COMMENT '更新时间',
    `deleted_at` datetime COMMENT '删除时间',
    `csid` int COMMENT '',
    PRIMARY KEY (`id`),
    KEY `idx_name` (`name`),
    KEY `idx_create` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色';

CREATE TABLE `role_user` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `role_id` int NOT NULL,
    `user_id` int NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_role` (`role_id`),
    KEY `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色用户关系';

CREATE TABLE `permission` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(60) NOT NULL COMMENT '名称',
    `name_code` varchar(60) NOT NULL COMMENT '权限代码',
    `pid` int NOT NULL DEFAULT '0' COMMENT '父ID',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime COMMENT '更新时间',
    `deleted_at` datetime COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unq_code` (`name_code`),
    KEY `idx_name` (`name`),
    KEY `idx_create` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='权限';

CREATE TABLE `permission_role` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `role_id` int NOT NULL,
    `permission_id` int NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_role` (`role_id`),
    KEY `idx_permission` (`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='权限角色关系';