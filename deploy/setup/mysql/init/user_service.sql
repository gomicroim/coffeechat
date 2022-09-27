CREATE TABLE `devices` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created` timestamp NOT NULL,
  `updated` timestamp NOT NULL,
  `device_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '设备id',
  `app_version` int NOT NULL COMMENT '版本号',
  `os_version` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '系统版本',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `nick_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '昵称',
  `sex` tinyint NOT NULL COMMENT '性别 0:未设置 1:男 2:女 3:非二次元性别',
  `phone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '手机号',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '邮箱',
  `extra` json DEFAULT NULL COMMENT '额外信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;