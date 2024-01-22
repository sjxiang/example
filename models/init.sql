
CREATE DATABASE IF NOT EXISTS `simple_bank` COLLATE utf8mb4_general_ci;

USE `simple_bank`;

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username`        varchar(255) NOT NULL,
  `hashed_password` varchar(255) NOT NULL,
  `full_name`       varchar(255) NOT NULL,
  `email`           varchar(255) NOT NULL,
  `role`            varchar(255) NOT NULL,
  `is_email_verified`   tinyint(1) NOT NULL DEFAULT '0',
  `created_at`          timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `password_changed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '密码最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`),
  UNIQUE KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SELECT * FROM `user`\G


-- 当你插入 Insert 一行时，如果没有提供 update_at 的值，它将被设置为插入时的时间戳。
-- 当你更新 Update 这一行时，update_at 将自动更新为当前时间戳。



SET SESSION transaction isolation level read committed; read uncommitted repeatable read

SELECT @@transaction_isolation;

BEGIN;

COMMIT;