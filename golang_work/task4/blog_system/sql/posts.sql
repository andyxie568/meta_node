CREATE TABLE `posts` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                         `title` varchar(255) NOT NULL COMMENT '文章标题',
                         `content` text NOT NULL COMMENT '文章内容',
                         `user_id` int NOT NULL COMMENT '关联用户 ID',
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='博客文章表';