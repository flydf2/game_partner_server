-- 检查游戏数据
SELECT COUNT(*) as game_count FROM games;

-- 检查活动数据
SELECT COUNT(*) as activity_count FROM activities;

-- 检查分类数据
SELECT COUNT(*) as category_count FROM categories;

-- 检查游戏分类数据
SELECT COUNT(*) as game_category_count FROM game_categories;

-- 检查陪玩专家数据
SELECT COUNT(*) as playmate_count FROM playmates;

-- 检查陪玩技能数据
SELECT COUNT(*) as playmate_skill_count FROM playmate_skills;

-- 检查语音介绍数据
SELECT COUNT(*) as voice_intro_count FROM playmate_voice_introductions;

-- 检查社区帖子数据
SELECT COUNT(*) as community_post_count FROM community_posts;

-- 检查奖励订单数据
SELECT COUNT(*) as reward_order_count FROM reward_orders;
