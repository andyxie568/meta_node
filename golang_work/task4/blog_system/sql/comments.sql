CREATE TABLE `comments` (
                            `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增唯一标识',
                            `content` text NOT NULL COMMENT '评论内容',
                            `user_id` int NOT NULL COMMENT '用户ID，关联 users 表的 id',
                            `post_id` int NOT NULL COMMENT '文章ID，关联 posts 表的 id',
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论创建时间',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章评论表';