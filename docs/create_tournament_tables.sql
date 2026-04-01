-- 创建赛事表
CREATE TABLE IF NOT EXISTS `game_partner_tournaments` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(255) NOT NULL COMMENT '赛事标题',
  `description` text COMMENT '赛事描述',
  `cover` varchar(2048) COMMENT '封面图片',
  `game` varchar(255) COMMENT '游戏名称',
  `game_id` bigint UNSIGNED COMMENT '游戏ID',
  `status` varchar(255) NOT NULL DEFAULT 'upcoming' COMMENT '赛事状态 upcoming-报名中 ongoing-进行中 completed-已结束',
  `register_start` datetime NOT NULL COMMENT '报名开始时间',
  `register_end` datetime NOT NULL COMMENT '报名结束时间',
  `match_start` datetime NOT NULL COMMENT '比赛开始时间',
  `match_end` datetime NOT NULL COMMENT '比赛结束时间',
  `prize` varchar(255) COMMENT '奖金池',
  `participants` int NOT NULL DEFAULT 0 COMMENT '参赛人数',
  `max_teams` int NOT NULL DEFAULT 128 COMMENT '最大参赛队伍数',
  `rules` text COMMENT '赛事规则(JSON格式)',
  `format` varchar(255) COMMENT '赛制 BO3/BO5等',
  `min_rank` varchar(255) COMMENT '最低段位要求',
  `organizer` varchar(255) COMMENT '主办方',
  `contact_info` varchar(255) COMMENT '联系方式',
  PRIMARY KEY (`id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 创建赛事参赛队伍表
CREATE TABLE IF NOT EXISTS `game_partner_tournament_teams` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_at` datetime(3) DEFAULT NULL,
  `tournament_id` bigint UNSIGNED NOT NULL COMMENT '赛事ID',
  `name` varchar(255) NOT NULL COMMENT '队伍名称',
  `avatar` varchar(2048) COMMENT '队伍头像',
  `members` int NOT NULL DEFAULT 5 COMMENT '成员数量',
  `leader_id` bigint UNSIGNED COMMENT '队长用户ID',
  `rank` varchar(255) COMMENT '最终排名',
  `status` varchar(255) NOT NULL DEFAULT 'registered' COMMENT '状态 registered-已报名 approved-已通过 rejected-已拒绝',
  PRIMARY KEY (`id`),
  KEY `idx_tournament_id` (`tournament_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 创建赛事比赛表
CREATE TABLE IF NOT EXISTS `game_partner_tournament_matches` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_at` datetime(3) DEFAULT NULL,
  `tournament_id` bigint UNSIGNED NOT NULL COMMENT '赛事ID',
  `round` int COMMENT '轮次',
  `match_time` datetime NOT NULL COMMENT '比赛时间',
  `team1_id` bigint UNSIGNED COMMENT '队伍1ID',
  `team1_name` varchar(255) COMMENT '队伍1名称',
  `team1_avatar` varchar(2048) COMMENT '队伍1头像',
  `team2_id` bigint UNSIGNED COMMENT '队伍2ID',
  `team2_name` varchar(255) COMMENT '队伍2名称',
  `team2_avatar` varchar(2048) COMMENT '队伍2头像',
  `score1` int COMMENT '队伍1比分',
  `score2` int COMMENT '队伍2比分',
  `status` varchar(255) NOT NULL DEFAULT 'upcoming' COMMENT '比赛状态 upcoming-未开始 ongoing-进行中 completed-已结束',
  `winner_id` bigint UNSIGNED COMMENT '获胜队伍ID',
  PRIMARY KEY (`id`),
  KEY `idx_tournament_id` (`tournament_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 创建赛事报名记录表
CREATE TABLE IF NOT EXISTS `game_partner_tournament_registrations` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_at` datetime(3) DEFAULT NULL,
  `tournament_id` bigint UNSIGNED NOT NULL COMMENT '赛事ID',
  `user_id` bigint UNSIGNED NOT NULL COMMENT '用户ID',
  `team_name` varchar(255) NOT NULL COMMENT '队伍名称',
  `contact_info` varchar(255) COMMENT '联系信息',
  `members_info` text COMMENT '成员信息(JSON格式)',
  `status` varchar(255) NOT NULL DEFAULT 'pending' COMMENT '报名状态 pending-待审核 approved-已通过 rejected-已拒绝',
  `review_remark` varchar(255) COMMENT '审核备注',
  PRIMARY KEY (`id`),
  KEY `idx_tournament_id` (`tournament_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;