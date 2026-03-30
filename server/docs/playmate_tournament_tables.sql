-- 赛事管理模块数据库表结构
-- 创建时间: 2026-03-31
-- 对应前端: /System/Volumes/Data/webcode/UX/GamePartner/src/api/mock-tournament.js

-- 赛事表
CREATE TABLE IF NOT EXISTS `game_partner_tournaments` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '赛事ID',
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    `deleted_at` DATETIME(3) NULL DEFAULT NULL COMMENT '删除时间',
    `title` VARCHAR(255) NOT NULL COMMENT '赛事标题',
    `description` TEXT NULL COMMENT '赛事描述',
    `cover` VARCHAR(2048) NULL COMMENT '封面图片URL',
    `game` VARCHAR(100) NULL COMMENT '游戏名称',
    `game_id` BIGINT UNSIGNED NULL DEFAULT 0 COMMENT '游戏ID',
    `status` VARCHAR(20) NOT NULL DEFAULT 'upcoming' COMMENT '赛事状态: upcoming-报名中, ongoing-进行中, completed-已结束',
    `register_start` DATETIME(3) NULL COMMENT '报名开始时间',
    `register_end` DATETIME(3) NULL COMMENT '报名结束时间',
    `match_start` DATETIME(3) NULL COMMENT '比赛开始时间',
    `match_end` DATETIME(3) NULL COMMENT '比赛结束时间',
    `prize` VARCHAR(255) NULL COMMENT '奖金池',
    `participants` INT NOT NULL DEFAULT 0 COMMENT '当前参赛人数',
    `max_teams` INT NOT NULL DEFAULT 128 COMMENT '最大参赛队伍数',
    `rules` TEXT NULL COMMENT '赛事规则(JSON格式)',
    `format` VARCHAR(50) NULL COMMENT '赛制: BO3, BO5等',
    `min_rank` VARCHAR(100) NULL COMMENT '最低段位要求',
    `organizer` VARCHAR(255) NULL COMMENT '主办方',
    `contact_info` VARCHAR(500) NULL COMMENT '联系方式',
    PRIMARY KEY (`id`),
    INDEX `idx_status` (`status`),
    INDEX `idx_game` (`game`),
    INDEX `idx_game_id` (`game_id`),
    INDEX `idx_deleted_at` (`deleted_at`),
    INDEX `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='赛事表';

-- 赛事参赛队伍表
CREATE TABLE IF NOT EXISTS `game_partner_tournament_teams` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '队伍ID',
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    `deleted_at` DATETIME(3) NULL DEFAULT NULL COMMENT '删除时间',
    `tournament_id` BIGINT UNSIGNED NOT NULL COMMENT '所属赛事ID',
    `name` VARCHAR(255) NOT NULL COMMENT '队伍名称',
    `avatar` VARCHAR(2048) NULL COMMENT '队伍头像URL',
    `members` INT NOT NULL DEFAULT 5 COMMENT '成员数量',
    `leader_id` BIGINT UNSIGNED NULL DEFAULT 0 COMMENT '队长用户ID',
    `rank` VARCHAR(50) NULL COMMENT '最终排名: 冠军, 亚军, 季军等',
    `status` VARCHAR(20) NOT NULL DEFAULT 'registered' COMMENT '状态: registered-已报名, approved-已通过, rejected-已拒绝',
    PRIMARY KEY (`id`),
    INDEX `idx_tournament_id` (`tournament_id`),
    INDEX `idx_leader_id` (`leader_id`),
    INDEX `idx_status` (`status`),
    INDEX `idx_deleted_at` (`deleted_at`),
    UNIQUE KEY `uk_tournament_name` (`tournament_id`, `name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='赛事参赛队伍表';

-- 赛事比赛表
CREATE TABLE IF NOT EXISTS `game_partner_tournament_matches` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '比赛ID',
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    `deleted_at` DATETIME(3) NULL DEFAULT NULL COMMENT '删除时间',
    `tournament_id` BIGINT UNSIGNED NOT NULL COMMENT '所属赛事ID',
    `round` INT NOT NULL DEFAULT 1 COMMENT '轮次',
    `match_time` DATETIME(3) NULL COMMENT '比赛时间',
    `team1_id` BIGINT UNSIGNED NULL DEFAULT 0 COMMENT '队伍1 ID',
    `team1_name` VARCHAR(255) NULL COMMENT '队伍1名称',
    `team1_avatar` VARCHAR(2048) NULL COMMENT '队伍1头像',
    `team2_id` BIGINT UNSIGNED NULL DEFAULT 0 COMMENT '队伍2 ID',
    `team2_name` VARCHAR(255) NULL COMMENT '队伍2名称',
    `team2_avatar` VARCHAR(2048) NULL COMMENT '队伍2头像',
    `score1` INT NULL COMMENT '队伍1比分',
    `score2` INT NULL COMMENT '队伍2比分',
    `status` VARCHAR(20) NOT NULL DEFAULT 'upcoming' COMMENT '比赛状态: upcoming-未开始, ongoing-进行中, completed-已结束',
    `winner_id` BIGINT UNSIGNED NULL COMMENT '获胜队伍ID',
    PRIMARY KEY (`id`),
    INDEX `idx_tournament_id` (`tournament_id`),
    INDEX `idx_team1_id` (`team1_id`),
    INDEX `idx_team2_id` (`team2_id`),
    INDEX `idx_status` (`status`),
    INDEX `idx_match_time` (`match_time`),
    INDEX `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='赛事比赛表';

-- 赛事报名记录表
CREATE TABLE IF NOT EXISTS `game_partner_tournament_registrations` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '报名ID',
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    `deleted_at` DATETIME(3) NULL DEFAULT NULL COMMENT '删除时间',
    `tournament_id` BIGINT UNSIGNED NOT NULL COMMENT '赛事ID',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `team_name` VARCHAR(255) NOT NULL COMMENT '队伍名称',
    `contact_info` VARCHAR(500) NULL COMMENT '联系信息',
    `members_info` TEXT NULL COMMENT '成员信息(JSON格式)',
    `status` VARCHAR(20) NOT NULL DEFAULT 'pending' COMMENT '报名状态: pending-待审核, approved-已通过, rejected-已拒绝',
    `review_remark` VARCHAR(500) NULL COMMENT '审核备注',
    PRIMARY KEY (`id`),
    INDEX `idx_tournament_id` (`tournament_id`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_status` (`status`),
    INDEX `idx_deleted_at` (`deleted_at`),
    UNIQUE KEY `uk_tournament_user` (`tournament_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='赛事报名记录表';

-- 插入示例数据
-- 赛事示例数据
INSERT INTO `game_partner_tournaments` (
    `id`, `title`, `description`, `cover`, `game`, `status`, 
    `register_start`, `register_end`, `match_start`, `match_end`,
    `prize`, `participants`, `max_teams`, `rules`, `format`, `organizer`
) VALUES 
(
    1, 
    '王者荣耀春季联赛', 
    '王者荣耀春季联赛是官方举办的大型电竞赛事，汇聚了全服最顶尖的选手。赛事分为海选赛、淘汰赛和决赛三个阶段，最终决出冠军。冠军将获得丰厚奖金和限定皮肤。',
    'https://example.com/tournament-cover-1.jpg',
    '王者荣耀',
    'upcoming',
    '2026-04-01 00:00:00', '2026-04-05 23:59:59',
    '2026-04-10 00:00:00', '2026-04-15 23:59:59',
    '¥10,000', 128, 256,
    '["采用BO3淘汰赛制", "参赛者需达到指定段位要求", "禁止使用任何作弊手段", "最终解释权归主办方所有"]',
    'BO3',
    'GamePartner官方'
),
(
    2,
    '英雄联盟全球总决赛预选赛',
    '英雄联盟全球总决赛预选赛是国内最高水平的电竞赛事之一，汇聚了各路高手。赛事分为小组赛、淘汰赛和决赛三个阶段。',
    'https://example.com/tournament-cover-2.jpg',
    '英雄联盟',
    'ongoing',
    '2026-03-01 00:00:00', '2026-03-10 23:59:59',
    '2026-03-15 00:00:00', '2026-03-25 23:59:59',
    '¥50,000', 256, 512,
    '["采用BO5淘汰赛制", "必须使用正式服账号参赛", "禁止使用任何第三方软件", "保持良好的游戏态度"]',
    'BO5',
    'GamePartner官方'
),
(
    3,
    '绝地求生杯首届邀请赛',
    '绝地求生杯首届邀请赛，邀请了全服最顶尖的选手参与。赛事采用积分制，最终根据总积分排名。',
    'https://example.com/tournament-cover-3.jpg',
    '绝地求生',
    'completed',
    '2026-02-01 00:00:00', '2026-02-10 23:59:59',
    '2026-02-15 00:00:00', '2026-02-20 23:59:59',
    '¥30,000', 64, 128,
    '["采用多轮积分制", "每轮比赛独立计分", "禁止开挂作弊", "最终解释权归主办方所有"]',
    '积分制',
    'GamePartner官方'
);

-- 参赛队伍示例数据
INSERT INTO `game_partner_tournament_teams` (
    `id`, `tournament_id`, `name`, `avatar`, `members`, `rank`, `status`
) VALUES
(1, 1, 'WX战队', 'https://randomuser.me/api/portraits/men/1.jpg', 5, '冠军', 'approved'),
(2, 1, 'QG战队', 'https://randomuser.me/api/portraits/men/2.jpg', 5, '亚军', 'approved'),
(3, 1, 'AG战队', 'https://randomuser.me/api/portraits/men/3.jpg', 5, NULL, 'approved'),
(4, 1, 'RNG战队', 'https://randomuser.me/api/portraits/men/4.jpg', 5, NULL, 'approved'),
(5, 2, 'EDG战队', 'https://randomuser.me/api/portraits/men/5.jpg', 5, '冠军', 'approved'),
(6, 2, 'RNG战队', 'https://randomuser.me/api/portraits/men/6.jpg', 5, '亚军', 'approved'),
(7, 2, 'WE战队', 'https://randomuser.me/api/portraits/men/7.jpg', 5, NULL, 'approved'),
(8, 2, 'BLG战队', 'https://randomuser.me/api/portraits/men/8.jpg', 5, NULL, 'approved'),
(9, 3, '4AM战队', 'https://randomuser.me/api/portraits/men/9.jpg', 4, '冠军', 'approved'),
(10, 3, 'IFTY战队', 'https://randomuser.me/api/portraits/men/10.jpg', 4, '亚军', 'approved');

-- 比赛示例数据
INSERT INTO `game_partner_tournament_matches` (
    `id`, `tournament_id`, `round`, `match_time`,
    `team1_id`, `team1_name`, `team1_avatar`,
    `team2_id`, `team2_name`, `team2_avatar`,
    `score1`, `score2`, `status`, `winner_id`
) VALUES
(1, 1, 1, '2026-04-10 14:00:00', 1, 'WX战队', 'https://randomuser.me/api/portraits/men/1.jpg', 2, 'QG战队', 'https://randomuser.me/api/portraits/men/2.jpg', 2, 1, 'completed', 1),
(2, 1, 1, '2026-04-10 16:00:00', 3, 'AG战队', 'https://randomuser.me/api/portraits/men/3.jpg', 4, 'RNG战队', 'https://randomuser.me/api/portraits/men/4.jpg', NULL, NULL, 'upcoming', NULL),
(3, 2, 1, '2026-03-15 10:00:00', 5, 'EDG战队', 'https://randomuser.me/api/portraits/men/5.jpg', 6, 'RNG战队', 'https://randomuser.me/api/portraits/men/6.jpg', 3, 2, 'completed', 5),
(4, 2, 1, '2026-03-15 14:00:00', 7, 'WE战队', 'https://randomuser.me/api/portraits/men/7.jpg', 8, 'BLG战队', 'https://randomuser.me/api/portraits/men/8.jpg', 1, 3, 'completed', 8),
(5, 2, 2, '2026-03-20 14:00:00', 5, 'EDG战队', 'https://randomuser.me/api/portraits/men/5.jpg', 8, 'BLG战队', 'https://randomuser.me/api/portraits/men/8.jpg', NULL, NULL, 'ongoing', NULL),
(6, 3, 1, '2026-02-15 10:00:00', 9, '4AM战队', 'https://randomuser.me/api/portraits/men/9.jpg', 10, 'IFTY战队', 'https://randomuser.me/api/portraits/men/10.jpg', 2, 0, 'completed', 9);

-- 重置自增ID
ALTER TABLE `game_partner_tournaments` AUTO_INCREMENT = 4;
ALTER TABLE `game_partner_tournament_teams` AUTO_INCREMENT = 11;
ALTER TABLE `game_partner_tournament_matches` AUTO_INCREMENT = 7;
