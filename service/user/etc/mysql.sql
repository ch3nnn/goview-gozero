CREATE TABLE `sys_user`
(
    `id`       int          NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `username` varchar(50)  NOT NULL COMMENT '用户名',
    `password` varchar(100) NOT NULL COMMENT '密码',
    `nickname` varchar(100) DEFAULT NULL COMMENT '昵称',
    `dep_id`   int          DEFAULT NULL,
    `pos_id`   varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表'
