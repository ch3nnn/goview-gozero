CREATE TABLE `screen_project`
(
    `id`        int NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`      varchar(255) DEFAULT NULL COMMENT '大屏名称',
    `state`     int          DEFAULT NULL COMMENT '发布状态(-1 未发布  1 已发布)',
    `user_id`   int          DEFAULT NULL COMMENT '创建用户ID',
    `index_img` varchar(255) DEFAULT NULL COMMENT '缩略图',
    `remark`    text COMMENT '备注',
    `is_del`    tinyint(1) DEFAULT '0' COMMENT ' 是否删除(0 未删除 1 已删除)',
    `create_at` datetime     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='大屏信息'


CREATE TABLE `screen_datum`
(
    `id`         int NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `project_id` int      DEFAULT NULL COMMENT 'project ID',
    `user_id`    int      DEFAULT NULL COMMENT '用户 ID',
    `content`    text COMMENT '内容数据',
    `create_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='大屏数据'
