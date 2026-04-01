-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS `game_partner` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `game_partner`;

-- 用户表
CREATE TABLE IF NOT EXISTS `game_partner_users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  `username` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `phone` VARCHAR(20) NOT NULL,
  `avatar` VARCHAR(255) NOT NULL,
  `nickname` VARCHAR(255) NOT NULL,
  `gender` VARCHAR(10) NULL,
  `birthday` VARCHAR(20) NULL,
  `bio` TEXT NULL,
  `location` VARCHAR(255) NULL,
  `vip_level` INT NOT NULL DEFAULT 1,
  `balance` DECIMAL(10,2) NOT NULL DEFAULT 0,
  `coupon_count` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_username` (`username`),
  UNIQUE INDEX `idx_phone` (`phone`),
  INDEX `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 用户设置表
CREATE TABLE IF NOT EXISTS `game_partner_user_settings` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `notifications` TEXT NOT NULL,
  `privacy` TEXT NOT NULL,
  `theme` VARCHAR(50) NOT NULL DEFAULT 'light',
  `language` VARCHAR(20) NOT NULL DEFAULT 'zh-CN',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_user_id` (`user_id`),
  CONSTRAINT `fk_user_settings_user` FOREIGN KEY (`user_id`) REFERENCES `game_partner_users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 用户钱包表
CREATE TABLE IF NOT EXISTS `game_partner_user_wallets` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `balance` DECIMAL(10,2) NOT NULL DEFAULT 0,
  `frozen` DECIMAL(10,2) NOT NULL DEFAULT 0,
  `total_income` DECIMAL(10,2) NOT NULL DEFAULT 0,
  `total_expense` DECIMAL(10,2) NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_user_id` (`user_id`),
  CONSTRAINT `fk_user_wallets_user` FOREIGN KEY (`user_id`) REFERENCES `game_partner_users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 交易记录表
CREATE TABLE IF NOT EXISTS `game_partner_transactions` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `type` VARCHAR(20) NOT NULL, -- income, expense, income_pending, expense_pending
  `amount` DECIMAL(10,2) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_user_id` (`user_id`),
  INDEX `idx_time` (`time`),
  INDEX `idx_type` (`type`),
  CONSTRAINT `fk_transactions_user` FOREIGN KEY (`user_id`) REFERENCES `game_partner_users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 提现记录表
CREATE TABLE IF NOT EXISTS `game_partner_withdrawals` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `amount` DECIMAL(10,2) NOT NULL,
  `fee` DECIMAL(10,2) NOT NULL,
  `actual_amount` DECIMAL(10,2) NOT NULL,
  `method` VARCHAR(20) NOT NULL, -- wechat, alipay, bank
  `status` VARCHAR(20) NOT NULL, -- pending, processing, completed, failed
  `completed_at` DATETIME NULL,
  `failed_reason` VARCHAR(255) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_user_id` (`user_id`),
  INDEX `idx_status` (`status`),
  INDEX `idx_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_withdrawals_user` FOREIGN KEY (`user_id`) REFERENCES `game_partner_users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 评价表
CREATE TABLE IF NOT EXISTS `game_partner_reviews` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `playmate_id` BIGINT UNSIGNED NOT NULL,
  `order_id` BIGINT UNSIGNED NOT NULL DEFAULT 0,
  `rating` INT NOT NULL,
  `content` TEXT NOT NULL,
  `images` VARCHAR(4096) NULL,
  `tags` VARCHAR(255) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_user_id` (`user_id`),
  INDEX `idx_playmate_id` (`playmate_id`),
  INDEX `idx_order_id` (`order_id`),
  INDEX `idx_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_reviews_user` FOREIGN KEY (`user_id`) REFERENCES `game_partner_users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;