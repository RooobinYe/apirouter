CREATE TABLE `apikeys` (
  `id` varchar(36) NOT NULL COMMENT 'UUID',
  `user_id` bigint NOT NULL,
  `name` varchar(255) NOT NULL,
  `api_key` varchar(255) NOT NULL UNIQUE,
  `created_at` bigint NOT NULL,
  `status` varchar(20) NOT NULL DEFAULT 'active' COMMENT 'active, inactive, expired',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_api_key` (`api_key`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;