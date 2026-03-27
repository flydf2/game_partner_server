-- 为评价表添加order_id字段
ALTER TABLE `game_partner_reviews` ADD COLUMN `order_id` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单ID' AFTER `playmate_id`;