CREATE TABLE `sys_user`
(
    `id`       int          NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `username` varchar(50)  NOT NULL COMMENT '用户名',
    `password` varchar(100) NOT NULL COMMENT '密码',
    `nickname` varchar(100) DEFAULT NULL COMMENT '昵称',
    `dep_id`   int          DEFAULT NULL,
    `pos_id`   varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';


CREATE TABLE `screen_project`
(
    `id`        int NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`      varchar(255) DEFAULT NULL COMMENT '大屏名称',
    `state`     int          DEFAULT '1' COMMENT '发布状态(-1 未发布  1 已发布)',
    `user_id`   int          DEFAULT NULL COMMENT '创建用户ID',
    `index_img` varchar(255) DEFAULT NULL COMMENT '缩略图',
    `remark`    text COMMENT '备注',
    `is_del`    tinyint(1) DEFAULT '0' COMMENT ' 是否删除(0 未删除 1 已删除)',
    `create_at` datetime     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='大屏信息';


CREATE TABLE `screen_datum`
(
    `id`         int NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `project_id` int      DEFAULT NULL COMMENT 'project ID',
    `user_id`    int      DEFAULT NULL COMMENT '用户 ID',
    `content`    text COMMENT '内容数据',
    `create_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='大屏数据';



INSERT INTO goview.sys_user (id, username, password, nickname, dep_id, pos_id)
VALUES (1, 'admin', '21232f297a57a5a743894a0e4a801fc3', '管理员', 2, '410792368778907648');
