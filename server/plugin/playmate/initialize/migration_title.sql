-- 为悬赏订单表添加title字段
ALTER TABLE `reward_orders` ADD COLUMN `title` VARCHAR(200) DEFAULT NULL COMMENT '订单标题' AFTER `game`;