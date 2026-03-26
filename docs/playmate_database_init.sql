-- 游戏伙伴平台数据库初始化SQL
-- 日期：2026-03-25

-- 使用game_partner数据库
USE game_partner;

-- 创建用户表
CREATE TABLE IF NOT EXISTS game_partner_users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    avatar VARCHAR(255) DEFAULT '',
    nickname VARCHAR(50) NOT NULL,
    vip_level INT DEFAULT 0,
    balance FLOAT DEFAULT 0,
    coupon_count INT DEFAULT 0
);

-- 创建用户设置表
CREATE TABLE IF NOT EXISTS game_partner_user_settings (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    notifications TEXT,
    privacy TEXT,
    theme VARCHAR(50) DEFAULT 'light',
    language VARCHAR(20) DEFAULT 'zh-CN',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id)
);

-- 创建用户钱包表
CREATE TABLE IF NOT EXISTS game_partner_user_wallets (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    balance FLOAT DEFAULT 0,
    frozen FLOAT DEFAULT 0,
    total_income FLOAT DEFAULT 0,
    total_expense FLOAT DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id)
);

-- 创建交易记录表
CREATE TABLE IF NOT EXISTS game_partner_transactions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    type VARCHAR(20) NOT NULL,
    amount FLOAT NOT NULL,
    description VARCHAR(255) DEFAULT '',
    time DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id)
);

-- 创建社区帖子表
CREATE TABLE IF NOT EXISTS game_partner_community_posts (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    user_id INT NOT NULL,
    content TEXT NOT NULL,
    images VARCHAR(4096) DEFAULT '',
    likes INT DEFAULT 0,
    comments INT DEFAULT 0,
    game VARCHAR(100) DEFAULT '',
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id)
);

-- 创建评论表
CREATE TABLE IF NOT EXISTS game_partner_comments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    post_id INT NOT NULL,
    user_id INT NOT NULL,
    content TEXT NOT NULL,
    likes INT DEFAULT 0,
    FOREIGN KEY (post_id) REFERENCES game_partner_community_posts(id),
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id)
);

-- 创建推荐表
CREATE TABLE IF NOT EXISTS game_partner_recommendations (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    type VARCHAR(50) NOT NULL,
    data TEXT NOT NULL
);

-- 创建用户关注表
CREATE TABLE IF NOT EXISTS game_partner_user_follows (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    user_id INT NOT NULL,
    follow_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id),
    FOREIGN KEY (follow_id) REFERENCES game_partner_users(id)
);

-- 创建用户收藏表
CREATE TABLE IF NOT EXISTS game_partner_user_favorites (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    user_id INT NOT NULL,
    playmate_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id)
);

-- 创建用户浏览历史表
CREATE TABLE IF NOT EXISTS game_partner_user_browse_histories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    user_id INT NOT NULL,
    playmate_id INT NOT NULL,
    viewed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id)
);

-- 创建陪玩专家表
CREATE TABLE IF NOT EXISTS game_partner_playmates (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    user_id INT NOT NULL,
    nickname VARCHAR(50) NOT NULL,
    avatar VARCHAR(2048) DEFAULT '',
    rating FLOAT DEFAULT 0,
    price FLOAT DEFAULT 0,
    likes INT DEFAULT 0,
    tags VARCHAR(255) DEFAULT '',
    is_online BOOLEAN DEFAULT false,
    game VARCHAR(100) DEFAULT '',
    `rank` VARCHAR(50) DEFAULT '',
    gender VARCHAR(10) DEFAULT '',
    description TEXT,
    level INT DEFAULT 0,
    title VARCHAR(100) DEFAULT '',
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id)
);

-- 创建陪玩技能表
CREATE TABLE IF NOT EXISTS game_partner_playmate_skills (
    id INT PRIMARY KEY AUTO_INCREMENT,
    playmate_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    price FLOAT DEFAULT 0,
    level VARCHAR(50) DEFAULT '',
    description TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (playmate_id) REFERENCES game_partner_playmates(id)
);

-- 创建陪玩语音介绍表
CREATE TABLE IF NOT EXISTS game_partner_playmate_voice_introductions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    playmate_id INT NOT NULL,
    url VARCHAR(255) DEFAULT '',
    duration VARCHAR(20) DEFAULT '',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (playmate_id) REFERENCES game_partner_playmates(id)
);

-- 创建通知表
CREATE TABLE IF NOT EXISTS game_partner_notifications (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    user_id INT NOT NULL,
    type VARCHAR(50) NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    time DATETIME DEFAULT CURRENT_TIMESTAMP,
    `read` BOOLEAN DEFAULT false,
    order_id INT NULL,
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id)
);

-- 创建消息表
CREATE TABLE IF NOT EXISTS game_partner_messages (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    from_user_id INT NOT NULL,
    to_user_id INT NOT NULL,
    content TEXT NOT NULL,
    time DATETIME DEFAULT CURRENT_TIMESTAMP,
    `read` BOOLEAN DEFAULT false,
    FOREIGN KEY (from_user_id) REFERENCES game_partner_users(id),
    FOREIGN KEY (to_user_id) REFERENCES game_partner_users(id)
);

-- 创建聊天消息表
CREATE TABLE IF NOT EXISTS game_partner_chat_messages (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    `from` VARCHAR(20) NOT NULL,
    content TEXT NOT NULL,
    time DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 插入用户数据
INSERT INTO game_partner_users (username, password, phone, avatar, nickname, vip_level, balance, coupon_count) VALUES
('user001', 'password123', '13800138001', 'https://example.com/avatar1.jpg', '游戏达人', 2, 1000, 5),
('user002', 'password123', '13800138002', 'https://example.com/avatar2.jpg', '电竞高手', 1, 500, 3),
('user003', 'password123', '13800138003', 'https://example.com/avatar3.jpg', '休闲玩家', 0, 200, 1),
('user004', 'password123', '13800138004', 'https://example.com/avatar4.jpg', '陪玩专家', 3, 2000, 10),
('user005', 'password123', '13800138005', 'https://example.com/avatar5.jpg', '游戏主播', 2, 1500, 7),
('user006', 'password123', '13800138006', 'https://example.com/avatar6.jpg', '新手玩家', 0, 100, 0),
('user007', 'password123', '13800138007', 'https://example.com/avatar7.jpg', '团战大师', 1, 800, 4),
('user008', 'password123', '13800138008', 'https://example.com/avatar8.jpg', '单挑王', 2, 1200, 6),
('user009', 'password123', '13800138009', 'https://example.com/avatar9.jpg', '战略大师', 1, 600, 2),
('user010', 'password123', '13800138010', 'https://example.com/avatar10.jpg', '游戏解说', 3, 2500, 12),
('user011', 'password123', '13800138011', 'https://example.com/avatar11.jpg', '职业选手', 3, 3000, 15),
('user012', 'password123', '13800138012', 'https://example.com/avatar12.jpg', '游戏设计师', 2, 1800, 8),
('user013', 'password123', '13800138013', 'https://example.com/avatar13.jpg', '休闲娱乐', 0, 150, 0),
('user014', 'password123', '13800138014', 'https://example.com/avatar14.jpg', '竞技达人', 2, 1300, 7),
('user015', 'password123', '13800138015', 'https://example.com/avatar15.jpg', '游戏测评', 1, 700, 3),
('user016', 'password123', '13800138016', 'https://example.com/avatar16.jpg', '新手指导', 1, 650, 2),
('user017', 'password123', '13800138017', 'https://example.com/avatar17.jpg', '团队领袖', 2, 1400, 8),
('user018', 'password123', '13800138018', 'https://example.com/avatar18.jpg', '游戏爱好者', 0, 250, 1),
('user019', 'password123', '13800138019', 'https://example.com/avatar19.jpg', '技术大神', 3, 2800, 14),
('user020', 'password123', '13800138020', 'https://example.com/avatar20.jpg', '游戏达人', 1, 900, 4);

-- 插入用户设置数据
INSERT INTO game_partner_user_settings (user_id, notifications, privacy, theme, language) VALUES
(1, '{"push": true, "email": true, "sms": false}', '{"profile": "public", "activity": "public", "messages": "friends"}', 'dark', 'zh-CN'),
(2, '{"push": true, "email": false, "sms": false}', '{"profile": "public", "activity": "private", "messages": "everyone"}', 'light', 'zh-CN'),
(3, '{"push": false, "email": false, "sms": false}', '{"profile": "private", "activity": "private", "messages": "friends"}', 'light', 'zh-CN'),
(4, '{"push": true, "email": true, "sms": true}', '{"profile": "public", "activity": "public", "messages": "everyone"}', 'dark', 'zh-CN'),
(5, '{"push": true, "email": true, "sms": false}', '{"profile": "public", "activity": "public", "messages": "everyone"}', 'dark', 'zh-CN'),
(6, '{"push": false, "email": false, "sms": false}', '{"profile": "private", "activity": "private", "messages": "friends"}', 'light', 'zh-CN'),
(7, '{"push": true, "email": false, "sms": false}', '{"profile": "public", "activity": "public", "messages": "everyone"}', 'light', 'zh-CN'),
(8, '{"push": true, "email": true, "sms": false}', '{"profile": "public", "activity": "public", "messages": "everyone"}', 'dark', 'zh-CN'),
(9, '{"push": false, "email": false, "sms": false}', '{"profile": "private", "activity": "private", "messages": "friends"}', 'light', 'zh-CN'),
(10, '{"push": true, "email": true, "sms": true}', '{"profile": "public", "activity": "public", "messages": "everyone"}', 'dark', 'zh-CN'),
(11, '{"push": true, "email": true, "sms": true}', '{"profile": "public", "activity": "public", "messages": "everyone"}', 'dark', 'zh-CN'),
(12, '{"push": true, "email": false, "sms": false}', '{"profile": "public", "activity": "private", "messages": "friends"}', 'light', 'zh-CN'),
(13, '{"push": false, "email": false, "sms": false}', '{"profile": "private", "activity": "private", "messages": "friends"}', 'light', 'zh-CN'),
(14, '{"push": true, "email": true, "sms": false}', '{"profile": "public", "activity": "public", "messages": "everyone"}', 'dark', 'zh-CN'),
(15, '{"push": false, "email": false, "sms": false}', '{"profile": "private", "activity": "private", "messages": "friends"}', 'light', 'zh-CN'),
(16, '{"push": true, "email": false, "sms": false}', '{"profile": "public", "activity": "public", "messages": "everyone"}', 'light', 'zh-CN'),
(17, '{"push": true, "email": true, "sms": false}', '{"profile": "public", "activity": "public", "messages": "everyone"}', 'dark', 'zh-CN'),
(18, '{"push": false, "email": false, "sms": false}', '{"profile": "private", "activity": "private", "messages": "friends"}', 'light', 'zh-CN'),
(19, '{"push": true, "email": true, "sms": true}', '{"profile": "public", "activity": "public", "messages": "everyone"}', 'dark', 'zh-CN'),
(20, '{"push": true, "email": false, "sms": false}', '{"profile": "public", "activity": "public", "messages": "everyone"}', 'light', 'zh-CN');

-- 插入用户钱包数据
INSERT INTO game_partner_user_wallets (user_id, balance, frozen, total_income, total_expense) VALUES
(1, 1000, 0, 5000, 4000),
(2, 500, 0, 2000, 1500),
(3, 200, 0, 800, 600),
(4, 2000, 0, 10000, 8000),
(5, 1500, 0, 7000, 5500),
(6, 100, 0, 300, 200),
(7, 800, 0, 3500, 2700),
(8, 1200, 0, 6000, 4800),
(9, 600, 0, 2500, 1900),
(10, 2500, 0, 12000, 9500),
(11, 3000, 0, 15000, 12000),
(12, 1800, 0, 8000, 6200),
(13, 150, 0, 500, 350),
(14, 1300, 0, 6500, 5200),
(15, 700, 0, 3000, 2300),
(16, 650, 0, 2800, 2150),
(17, 1400, 0, 6800, 5400),
(18, 250, 0, 900, 650),
(19, 2800, 0, 14000, 11200),
(20, 900, 0, 4000, 3100);

-- 插入交易记录数据
INSERT INTO game_partner_transactions (user_id, type, amount, description, time) VALUES
(1, 'income', 200, '陪玩服务收入', '2026-03-20 10:00:00'),
(1, 'expense', 50, '购买游戏道具', '2026-03-21 15:30:00'),
(2, 'income', 150, '陪玩服务收入', '2026-03-20 11:00:00'),
(3, 'expense', 30, '充值游戏币', '2026-03-22 09:15:00'),
(4, 'income', 300, '陪玩服务收入', '2026-03-19 14:45:00'),
(4, 'expense', 80, '购买装备', '2026-03-21 16:20:00'),
(5, 'income', 250, '直播打赏收入', '2026-03-18 20:30:00'),
(6, 'expense', 20, '购买游戏皮肤', '2026-03-23 10:45:00'),
(7, 'income', 180, '陪玩服务收入', '2026-03-20 12:15:00'),
(8, 'income', 220, '陪玩服务收入', '2026-03-19 15:30:00'),
(9, 'expense', 40, '购买游戏通行证', '2026-03-22 11:20:00'),
(10, 'income', 400, '直播打赏收入', '2026-03-18 21:45:00'),
(11, 'income', 500, '比赛奖金', '2026-03-17 16:30:00'),
(12, 'expense', 60, '购买设计素材', '2026-03-23 14:20:00'),
(13, 'expense', 15, '购买游戏道具', '2026-03-24 09:30:00'),
(14, 'income', 280, '陪玩服务收入', '2026-03-20 13:45:00'),
(15, 'income', 120, '游戏测评收入', '2026-03-19 10:15:00'),
(16, 'expense', 25, '购买游戏指南', '2026-03-22 15:40:00'),
(17, 'income', 320, '团队比赛奖金', '2026-03-18 14:30:00'),
(18, 'expense', 35, '充值会员', '2026-03-24 11:25:00'),
(19, 'income', 450, '技术指导收入', '2026-03-17 15:20:00'),
(20, 'income', 200, '陪玩服务收入', '2026-03-20 16:45:00');

-- 插入社区帖子数据
INSERT INTO game_partner_community_posts (user_id, content, images, likes, comments, game) VALUES
(1, '今天和朋友们一起玩了一整天的英雄联盟，感觉状态很好！', 'https://example.com/img1.jpg,https://example.com/img2.jpg', 15, 8, '英雄联盟'),
(2, '王者荣耀新赛季开始了，有没有一起冲分的小伙伴？', 'https://example.com/img3.jpg', 12, 6, '王者荣耀'),
(3, '休闲玩家一枚，想找个一起玩和平精英的队友', 'https://example.com/img4.jpg,https://example.com/img5.jpg', 8, 4, '和平精英'),
(4, '作为陪玩专家，今天接了5单，收入不错！', 'https://example.com/img6.jpg', 20, 12, '英雄联盟'),
(5, '今天直播了5个小时，感谢大家的支持！', 'https://example.com/img7.jpg,https://example.com/img8.jpg', 25, 15, '王者荣耀'),
(6, '刚接触游戏，有什么新手攻略吗？', 'https://example.com/img9.jpg', 5, 3, '和平精英'),
(7, '团战配合真的很重要，今天和队友配合得很默契', 'https://example.com/img10.jpg,https://example.com/img11.jpg', 18, 10, '英雄联盟'),
(8, '单挑我从来没怕过谁，有 challenger 吗？', 'https://example.com/img12.jpg', 16, 9, '王者荣耀'),
(9, '玩游戏还是要讲究策略，无脑冲只会送人头', 'https://example.com/img13.jpg', 14, 7, '和平精英'),
(10, '今天解说了一场精彩的比赛，大家觉得怎么样？', 'https://example.com/img14.jpg,https://example.com/img15.jpg', 30, 18, '英雄联盟'),
(11, '职业选手的日常训练真的很辛苦', 'https://example.com/img16.jpg', 28, 16, '王者荣耀'),
(12, '作为游戏设计师，我觉得游戏平衡真的很重要', 'https://example.com/img17.jpg', 22, 13, '和平精英'),
(13, '休闲娱乐为主，玩游戏就是图个开心', 'https://example.com/img18.jpg', 7, 3, '英雄联盟'),
(14, '竞技游戏的魅力就在于不断挑战自己', 'https://example.com/img19.jpg,https://example.com/img20.jpg', 21, 14, '王者荣耀'),
(15, '最近测试了几款新游戏，有不错的推荐', 'https://example.com/img21.jpg', 19, 11, '和平精英'),
(16, '新手指导第5期：如何快速上手游戏', 'https://example.com/img22.jpg,https://example.com/img23.jpg', 13, 8, '英雄联盟'),
(17, '团队领袖的职责就是带领大家走向胜利', 'https://example.com/img24.jpg', 24, 15, '王者荣耀'),
(18, '游戏爱好者的日常：每天玩两小时', 'https://example.com/img25.jpg', 9, 5, '和平精英'),
(19, '技术才是硬道理，操作决定一切', 'https://example.com/img26.jpg,https://example.com/img27.jpg', 27, 17, '英雄联盟'),
(20, '游戏达人分享：如何提高游戏水平', 'https://example.com/img28.jpg', 17, 10, '王者荣耀');

-- 插入评论数据
INSERT INTO game_partner_comments (post_id, user_id, content, likes) VALUES
(1, 2, '一起玩啊，我也喜欢英雄联盟', 5),
(1, 3, '羡慕你们能玩一整天', 3),
(2, 1, '我也想冲分，加个好友', 4),
(2, 4, '新赛季确实需要队友', 2),
(3, 5, '我可以和你一起玩', 3),
(3, 6, '新手求带', 2),
(4, 7, '厉害了，一天5单', 6),
(4, 8, '陪玩收入怎么样？', 4),
(5, 9, '主播辛苦了', 7),
(5, 10, '支持你！', 5),
(6, 11, '多看看攻略视频', 3),
(6, 12, '新手慢慢来', 2),
(7, 13, '团战配合确实重要', 4),
(7, 14, '默契是练出来的', 3),
(8, 15, '我来挑战你', 5),
(8, 16, '单挑王好厉害', 3),
(9, 17, '说得对，策略很重要', 4),
(9, 18, '无脑冲确实不行', 2),
(10, 19, '解说得很精彩', 6),
(10, 20, '期待下一次解说', 4),
(11, 1, '职业选手不容易', 5),
(11, 2, '加油！', 3),
(12, 3, '游戏平衡确实重要', 4),
(12, 4, '作为玩家也觉得平衡很重要', 2),
(13, 5, '玩游戏就是图开心', 3),
(13, 6, '同意，休闲娱乐为主', 2),
(14, 7, '挑战自己才有意思', 4),
(14, 8, '竞技游戏的魅力所在', 3),
(15, 9, '期待你的推荐', 5),
(15, 10, '新游戏测评很期待', 3),
(16, 11, '新手指导很有用', 4),
(16, 12, '感谢分享', 2),
(17, 13, '团队领袖责任重大', 3),
(17, 14, '带领团队胜利的感觉很棒', 2),
(18, 15, '每天两小时刚好', 3),
(18, 16, '适度游戏益脑', 2),
(19, 17, '技术确实重要', 4),
(19, 18, '操作决定一切', 3),
(20, 19, '期待你的分享', 5),
(20, 20, '学习了', 2);

-- 插入推荐数据
INSERT INTO game_partner_recommendations (type, data) VALUES
('playmate', '{"id": 1, "name": "游戏达人", "rating": 4.9, "price": 50, "game": "英雄联盟"}'),
('playmate', '{"id": 2, "name": "电竞高手", "rating": 4.8, "price": 60, "game": "王者荣耀"}'),
('playmate', '{"id": 3, "name": "陪玩专家", "rating": 4.95, "price": 80, "game": "和平精英"}'),
('playmate', '{"id": 4, "name": "团战大师", "rating": 4.7, "price": 45, "game": "英雄联盟"}'),
('playmate', '{"id": 5, "name": "单挑王", "rating": 4.85, "price": 55, "game": "王者荣耀"}'),
('activity', '{"id": 1, "title": "周末开黑派对", "time": "2026-03-28 14:00", "game": "英雄联盟"}'),
('activity', '{"id": 2, "title": "王者荣耀争霸赛", "time": "2026-03-29 15:00", "game": "王者荣耀"}'),
('activity', '{"id": 3, "title": "和平精英生存挑战", "time": "2026-03-30 16:00", "game": "和平精英"}'),
('playmate', '{"id": 6, "name": "战略大师", "rating": 4.75, "price": 48, "game": "和平精英"}'),
('playmate', '{"id": 7, "name": "游戏解说", "rating": 4.9, "price": 70, "game": "英雄联盟"}'),
('activity', '{"id": 4, "title": "新手训练营", "time": "2026-04-01 10:00", "game": "英雄联盟"}'),
('playmate', '{"id": 8, "name": "职业选手", "rating": 5.0, "price": 100, "game": "王者荣耀"}'),
('playmate', '{"id": 9, "name": "团队领袖", "rating": 4.8, "price": 65, "game": "英雄联盟"}'),
('activity', '{"id": 5, "title": "陪玩技能大赛", "time": "2026-04-02 14:00", "game": "王者荣耀"}'),
('playmate', '{"id": 10, "name": "技术大神", "rating": 4.95, "price": 90, "game": "和平精英"}'),
('activity', '{"id": 6, "title": "游戏设计分享会", "time": "2026-04-03 15:00", "game": "所有游戏"}'),
('playmate', '{"id": 11, "name": "游戏测评", "rating": 4.7, "price": 40, "game": "王者荣耀"}'),
('playmate', '{"id": 12, "name": "新手指导", "rating": 4.6, "price": 30, "game": "和平精英"}'),
('activity', '{"id": 7, "title": "电竞明星见面会", "time": "2026-04-04 16:00", "game": "英雄联盟"}'),
('playmate', '{"id": 13, "name": "游戏爱好者", "rating": 4.5, "price": 25, "game": "王者荣耀"}');

-- 插入用户关注数据
INSERT INTO game_partner_user_follows (user_id, follow_id) VALUES
(1, 2),
(1, 4),
(2, 1),
(2, 3),
(3, 2),
(3, 5),
(4, 1),
(4, 6),
(5, 3),
(5, 7),
(6, 4),
(6, 8),
(7, 5),
(7, 9),
(8, 6),
(8, 10),
(9, 7),
(9, 11),
(10, 8),
(10, 12),
(11, 9),
(11, 13),
(12, 10),
(12, 14),
(13, 11),
(13, 15),
(14, 12),
(14, 16),
(15, 13),
(15, 17),
(16, 14),
(16, 18),
(17, 15),
(17, 19),
(18, 16),
(18, 20),
(19, 17),
(19, 1),
(20, 18),
(20, 2);

-- 插入陪玩专家数据
INSERT INTO game_partner_playmates (user_id, nickname, avatar, rating, price, likes, tags, is_online, game, `rank`, gender, description, level, title) VALUES
(4, '陪玩专家', 'https://example.com/playmate1.jpg', 4.95, 80, 150, '英雄联盟,王者荣耀,和平精英', true, '英雄联盟', '钻石', '男', '专业陪玩，5年游戏经验，擅长各种位置', 5, '钻石陪玩'),
(1, '游戏达人', 'https://example.com/playmate2.jpg', 4.9, 50, 120, '英雄联盟,王者荣耀', false, '王者荣耀', '星耀', '男', '全能玩家，擅长团队配合', 4, '星耀玩家'),
(2, '电竞高手', 'https://example.com/playmate3.jpg', 4.8, 60, 100, '王者荣耀,和平精英', true, '和平精英', '王牌', '女', '职业电竞选手，反应迅速', 5, '职业选手'),
(7, '团战大师', 'https://example.com/playmate4.jpg', 4.7, 45, 80, '英雄联盟,王者荣耀', false, '英雄联盟', '白金', '男', '擅长组织团战，团队意识强', 3, '团战专家'),
(8, '单挑王', 'https://example.com/playmate5.jpg', 4.85, 55, 90, '王者荣耀,英雄联盟', true, '王者荣耀', '王者', '男', '单挑无敌，操作犀利', 5, '王者选手'),
(9, '战略大师', 'https://example.com/playmate6.jpg', 4.75, 48, 70, '和平精英,英雄联盟', false, '和平精英', '皇冠', '男', '擅长制定战略，战术大师', 4, '战略专家'),
(10, '游戏解说', 'https://example.com/playmate7.jpg', 4.9, 70, 130, '英雄联盟,王者荣耀', true, '英雄联盟', '大师', '女', '专业游戏解说，声音甜美', 5, '金牌解说'),
(11, '职业选手', 'https://example.com/playmate8.jpg', 5.0, 100, 200, '王者荣耀,英雄联盟', false, '王者荣耀', '荣耀王者', '男', '前职业选手，技术顶尖', 6, '职业大神'),
(17, '团队领袖', 'https://example.com/playmate9.jpg', 4.8, 65, 110, '英雄联盟,王者荣耀', true, '英雄联盟', '钻石', '女', '擅长领导团队，沟通能力强', 4, '团队核心'),
(19, '技术大神', 'https://example.com/playmate10.jpg', 4.95, 90, 180, '和平精英,英雄联盟', false, '和平精英', '无敌战神', '男', '技术流玩家，操作天花板', 6, '技术天花板'),
(5, '游戏主播', 'https://example.com/playmate11.jpg', 4.85, 75, 140, '王者荣耀,和平精英', true, '王者荣耀', '星耀', '女', '人气游戏主播，直播风格幽默', 5, '人气主播'),
(12, '游戏设计师', 'https://example.com/playmate12.jpg', 4.7, 60, 85, '英雄联盟,王者荣耀', false, '英雄联盟', '白金', '男', '游戏设计师，对游戏理解深刻', 4, '游戏专家'),
(14, '竞技达人', 'https://example.com/playmate13.jpg', 4.8, 58, 95, '王者荣耀,和平精英', true, '和平精英', '王牌', '男', '竞技游戏爱好者，反应迅速', 4, '竞技高手'),
(15, '游戏测评', 'https://example.com/playmate14.jpg', 4.7, 40, 60, '王者荣耀,英雄联盟', false, '王者荣耀', '钻石', '女', '专业游戏测评师，游戏见识广', 3, '测评专家'),
(16, '新手指导', 'https://example.com/playmate15.jpg', 4.6, 30, 50, '和平精英,王者荣耀', true, '和平精英', '钻石', '男', '擅长指导新手，耐心细致', 3, '新手导师'),
(3, '休闲玩家', 'https://example.com/playmate16.jpg', 4.5, 25, 40, '和平精英,王者荣耀', false, '和平精英', '黄金', '女', '休闲娱乐为主，轻松愉快', 2, '休闲玩家'),
(6, '新手玩家', 'https://example.com/playmate17.jpg', 4.4, 20, 30, '王者荣耀,英雄联盟', true, '王者荣耀', '白银', '男', '新手玩家，正在努力提升', 1, '新手玩家'),
(13, '休闲娱乐', 'https://example.com/playmate18.jpg', 4.5, 28, 45, '和平精英,王者荣耀', false, '和平精英', '黄金', '女', '休闲娱乐，开心就好', 2, '娱乐玩家'),
(18, '游戏爱好者', 'https://example.com/playmate19.jpg', 4.5, 35, 55, '英雄联盟,和平精英', true, '英雄联盟', '黄金', '男', '游戏爱好者，广泛涉猎', 2, '爱好者'),
(20, '游戏达人', 'https://example.com/playmate20.jpg', 4.6, 42, 65, '王者荣耀,英雄联盟', false, '王者荣耀', '钻石', '男', '游戏达人，多游戏精通', 3, '多面手');

-- 插入陪玩技能数据
INSERT INTO game_partner_playmate_skills (playmate_id, name, price, level, description) VALUES
(1, '英雄联盟上单', 80, '钻石', '擅长上单各种英雄，对线强势'),
(1, '王者荣耀打野', 75, '星耀', '打野节奏大师，带动全场'),
(2, '王者荣耀中单', 50, '星耀', '中单法王，输出爆炸'),
(2, '英雄联盟ADC', 45, '白金', 'ADC走位精准，输出稳定'),
(3, '和平精英刚枪', 60, '王牌', '刚枪小王子，所向披靡'),
(3, '王者荣耀辅助', 55, '钻石', '辅助意识强，保护到位'),
(4, '英雄联盟团战', 45, '白金', '团战指挥，战术大师'),
(4, '王者荣耀射手', 40, '钻石', '射手输出机器，稳定Carry'),
(5, '王者荣耀单挑', 55, '王者', '单挑无敌，操作犀利'),
(5, '英雄联盟中单', 50, '钻石', '中单杀神，游走支援'),
(6, '和平精英策略', 48, '皇冠', '战略大师，吃鸡专家'),
(6, '英雄联盟打野', 43, '白金', '打野节奏，掌控全局'),
(7, '英雄联盟解说', 70, '大师', '专业解说，分析到位'),
(7, '王者荣耀指导', 65, '星耀', '游戏指导，提升技巧'),
(8, '王者荣耀职业', 100, '荣耀王者', '职业选手，技术顶尖'),
(8, '英雄联盟职业', 95, '大师', '职业级操作，意识一流'),
(9, '英雄联盟团队', 65, '钻石', '团队核心，指挥全局'),
(9, '王者荣耀团战', 60, '星耀', '团战发动机，带领胜利'),
(10, '和平精英技术', 90, '无敌战神', '技术天花板，操作无敌'),
(10, '英雄联盟技术', 85, '大师', '技术流玩家，细节拉满'),
(11, '王者荣耀主播', 75, '星耀', '直播互动，娱乐教学'),
(11, '和平精英教学', 70, '王牌', '游戏教学，快速提升'),
(12, '游戏设计', 60, '白金', '游戏设计，深度解析'),
(12, '英雄联盟分析', 55, '钻石', '游戏分析，战术指导'),
(13, '和平精英竞技', 58, '王牌', '竞技玩法，高水平对战'),
(13, '王者荣耀竞技', 53, '星耀', '竞技比赛，专业指导'),
(14, '游戏测评', 40, '钻石', '游戏测评，推荐分析'),
(14, '王者荣耀指导', 35, '钻石', '新手指导，快速上手'),
(15, '新手指导', 30, '钻石', '耐心指导，基础教学'),
(15, '和平精英入门', 25, '黄金', '入门教学，快速上手'),
(16, '休闲娱乐', 25, '黄金', '轻松娱乐，开心游戏'),
(16, '和平精英娱乐', 20, '黄金', '娱乐玩法，欢乐多多'),
(17, '新手教学', 20, '白银', '基础教学，从零开始'),
(17, '王者荣耀入门', 18, '白银', '入门指导，快速上手'),
(18, '休闲娱乐', 28, '黄金', '轻松愉快，娱乐为主'),
(18, '和平精英娱乐', 25, '黄金', '娱乐玩法，欢乐时光'),
(19, '游戏爱好', 35, '黄金', '广泛涉猎，多游戏精通'),
(19, '英雄联盟娱乐', 30, '黄金', '娱乐玩法，开心就好'),
(20, '多游戏精通', 42, '钻石', '多游戏精通，全能玩家'),
(20, '王者荣耀娱乐', 38, '钻石', '娱乐玩法，欢乐无限');

-- 插入陪玩语音介绍数据
INSERT INTO game_partner_playmate_voice_introductions (playmate_id, url, duration) VALUES
(1, 'https://example.com/voice1.mp3', '30s'),
(2, 'https://example.com/voice2.mp3', '25s'),
(3, 'https://example.com/voice3.mp3', '28s'),
(4, 'https://example.com/voice4.mp3', '22s'),
(5, 'https://example.com/voice5.mp3', '26s'),
(6, 'https://example.com/voice6.mp3', '24s'),
(7, 'https://example.com/voice7.mp3', '32s'),
(8, 'https://example.com/voice8.mp3', '29s'),
(9, 'https://example.com/voice9.mp3', '27s'),
(10, 'https://example.com/voice10.mp3', '31s'),
(11, 'https://example.com/voice11.mp3', '28s'),
(12, 'https://example.com/voice12.mp3', '25s'),
(13, 'https://example.com/voice13.mp3', '26s'),
(14, 'https://example.com/voice14.mp3', '23s'),
(15, 'https://example.com/voice15.mp3', '24s'),
(16, 'https://example.com/voice16.mp3', '22s'),
(17, 'https://example.com/voice17.mp3', '21s'),
(18, 'https://example.com/voice18.mp3', '23s'),
(19, 'https://example.com/voice19.mp3', '25s'),
(20, 'https://example.com/voice20.mp3', '27s');

-- 插入用户收藏数据
INSERT INTO game_partner_user_favorites (user_id, playmate_id) VALUES
(1, 1),
(1, 3),
(2, 2),
(2, 4),
(3, 3),
(3, 5),
(4, 1),
(4, 6),
(5, 2),
(5, 7),
(6, 4),
(6, 8),
(7, 5),
(7, 9),
(8, 6),
(8, 10),
(9, 7),
(9, 11),
(10, 8),
(10, 12),
(11, 9),
(11, 13),
(12, 10),
(12, 14),
(13, 11),
(13, 15),
(14, 12),
(14, 16),
(15, 13),
(15, 17),
(16, 14),
(16, 18),
(17, 15),
(17, 19),
(18, 16),
(18, 20),
(19, 17),
(19, 1),
(20, 18),
(20, 2);

-- 插入用户浏览历史数据
INSERT INTO game_partner_user_browse_histories (user_id, playmate_id, viewed_at) VALUES
(1, 1, '2026-03-24 10:00:00'),
(1, 2, '2026-03-24 11:30:00'),
(2, 3, '2026-03-24 12:15:00'),
(2, 4, '2026-03-24 13:45:00'),
(3, 5, '2026-03-24 14:30:00'),
(3, 6, '2026-03-24 15:20:00'),
(4, 7, '2026-03-24 16:45:00'),
(4, 8, '2026-03-24 17:15:00'),
(5, 9, '2026-03-24 18:30:00'),
(5, 10, '2026-03-24 19:45:00'),
(6, 11, '2026-03-24 20:15:00'),
(6, 12, '2026-03-24 21:30:00'),
(7, 13, '2026-03-25 09:00:00'),
(7, 14, '2026-03-25 10:30:00'),
(8, 15, '2026-03-25 11:15:00'),
(8, 16, '2026-03-25 12:45:00'),
(9, 17, '2026-03-25 13:30:00'),
(9, 18, '2026-03-25 14:15:00'),
(10, 19, '2026-03-25 15:30:00'),
(10, 20, '2026-03-25 16:45:00'),
(11, 1, '2026-03-25 17:20:00'),
(11, 2, '2026-03-25 18:45:00'),
(12, 3, '2026-03-25 19:30:00'),
(12, 4, '2026-03-25 20:15:00'),
(13, 5, '2026-03-25 21:30:00'),
(13, 6, '2026-03-25 22:15:00'),
(14, 7, '2026-03-26 09:30:00'),
(14, 8, '2026-03-26 10:45:00'),
(15, 9, '2026-03-26 11:30:00'),
(15, 10, '2026-03-26 12:15:00'),
(16, 11, '2026-03-26 13:45:00'),
(16, 12, '2026-03-26 14:30:00'),
(17, 13, '2026-03-26 15:15:00'),
(17, 14, '2026-03-26 16:30:00'),
(18, 15, '2026-03-26 17:45:00'),
(18, 16, '2026-03-26 18:30:00'),
(19, 17, '2026-03-26 19:15:00'),
(19, 18, '2026-03-26 20:30:00'),
(20, 19, '2026-03-26 21:45:00'),
(20, 20, '2026-03-26 22:30:00');

-- 插入通知数据
INSERT INTO game_partner_notifications (user_id, type, title, content, `read`) VALUES
(1, 'system', '系统通知', '欢迎使用游戏伙伴平台！', 0),
(1, 'promotion', '促销活动', '新用户首单8折优惠', 0),
(2, 'system', '系统通知', '账号已成功注册', 1),
(2, 'message', '新消息', '有人给你发了一条消息', 0),
(3, 'system', '系统通知', '欢迎加入游戏伙伴', 0),
(4, 'order', '订单通知', '您的订单已完成', 1),
(4, 'system', '系统通知', '账号安全提醒', 0),
(5, 'promotion', '活动通知', '周末双倍积分活动', 0),
(5, 'system', '系统通知', '您的直播已开始', 1),
(6, 'system', '系统通知', '新手引导', 0),
(7, 'message', '新消息', '有人关注了你', 0),
(7, 'system', '系统通知', '等级提升', 1),
(8, 'promotion', '优惠活动', '会员专享折扣', 0),
(8, 'system', '系统通知', '技能认证成功', 1),
(9, 'system', '系统通知', '欢迎使用平台', 0),
(10, 'order', '订单通知', '新订单提醒', 0),
(10, 'system', '系统通知', '直播数据统计', 1),
(11, 'system', '系统通知', '职业选手认证', 1),
(11, 'promotion', '赛事通知', '参加比赛赢奖金', 0),
(12, 'system', '系统通知', '设计作品审核通过', 1),
(12, 'message', '新消息', '有人评论了你的帖子', 0),
(13, 'system', '系统通知', '欢迎注册', 0),
(14, 'promotion', '活动通知', '好友邀请奖励', 0),
(14, 'system', '系统通知', '账号信息更新', 1),
(15, 'message', '新消息', '测评邀请', 0),
(15, 'system', '系统通知', '测评任务完成', 1),
(16, 'system', '系统通知', '新手指导认证', 1),
(16, 'promotion', '学习优惠', '技能提升课程折扣', 0),
(17, 'order', '订单通知', '订单已支付', 1),
(17, 'system', '系统通知', '团队认证成功', 0),
(18, 'system', '系统通知', '欢迎加入', 0),
(19, 'promotion', '技术优惠', '技术指导服务折扣', 0),
(19, 'system', '系统通知', '技术认证成功', 1),
(20, 'message', '新消息', '有人给你点赞', 0),
(20, 'system', '系统通知', '等级提升', 1);

-- 插入消息数据
INSERT INTO game_partner_messages (from_user_id, to_user_id, content, `read`) VALUES
(1, 2, '你好，一起玩游戏吗？', 0),
(2, 1, '好啊，什么时候？', 1),
(1, 4, '你是陪玩专家吗？', 0),
(4, 1, '是的，有什么可以帮助你的？', 1),
(2, 3, '一起冲分吧', 0),
(3, 2, '好的，我随时可以', 1),
(3, 5, '你是游戏主播吗？', 0),
(5, 3, '是的，欢迎来看我的直播', 1),
(4, 6, '需要游戏指导吗？', 0),
(6, 4, '是的，我是新手', 1),
(5, 7, '一起玩王者荣耀吧', 0),
(7, 5, '好啊，我擅长打野', 1),
(6, 8, '你是单挑王吗？', 0),
(8, 6, '是的，有什么挑战吗？', 1),
(7, 9, '团战配合很重要', 0),
(9, 7, '是的，我擅长制定战略', 1),
(8, 10, '你是游戏解说吗？', 0),
(10, 8, '是的，欢迎来看我的解说', 1),
(9, 11, '你是职业选手吗？', 0),
(11, 9, '是的，前职业选手', 1),
(10, 12, '游戏设计很重要', 0),
(12, 10, '是的，平衡很关键', 1),
(11, 13, '休闲娱乐也不错', 0),
(13, 11, '是的，开心就好', 1),
(12, 14, '竞技游戏很刺激', 0),
(14, 12, '是的，挑战自己', 1),
(13, 15, '游戏测评很专业', 0),
(15, 13, '谢谢，我会继续努力', 1),
(14, 16, '新手指导很重要', 0),
(16, 14, '是的，我会耐心指导', 1),
(15, 17, '团队合作很重要', 0),
(17, 15, '是的，团队领袖很关键', 1),
(16, 18, '休闲玩家也很快乐', 0),
(18, 16, '是的，游戏就是图开心', 1),
(17, 19, '技术很重要', 0),
(19, 17, '是的，操作决定一切', 1),
(18, 20, '游戏达人分享一下经验吧', 0),
(20, 18, '好的，一起交流', 1),
(19, 1, '技术指导需要吗？', 0),
(1, 19, '需要，谢谢', 1),
(20, 2, '一起玩王者荣耀吧', 0),
(2, 20, '好啊，我擅长中单', 1);

-- 插入聊天消息数据
INSERT INTO game_partner_chat_messages (`from`, content, time) VALUES
('self', '你好，在吗？', '2026-03-25 10:00:00'),
('other', '在的，有什么事？', '2026-03-25 10:01:00'),
('self', '想找你一起玩游戏', '2026-03-25 10:02:00'),
('other', '好啊，玩什么游戏？', '2026-03-25 10:03:00'),
('self', '英雄联盟', '2026-03-25 10:04:00'),
('other', '可以，什么时候？', '2026-03-25 10:05:00'),
('self', '现在就可以', '2026-03-25 10:06:00'),
('other', '好的，我马上上线', '2026-03-25 10:07:00'),
('self', '王者荣耀新赛季开始了', '2026-03-25 11:00:00'),
('other', '是啊，我已经开始冲分了', '2026-03-25 11:01:00'),
('self', '一起吗？', '2026-03-25 11:02:00'),
('other', '好啊，我打野', '2026-03-25 11:03:00'),
('self', '我中单', '2026-03-25 11:04:00'),
('other', '没问题', '2026-03-25 11:05:00'),
('self', '和平精英有人一起吗？', '2026-03-25 12:00:00'),
('other', '我来，刚枪小王子', '2026-03-25 12:01:00'),
('self', '太好了，组队', '2026-03-25 12:02:00'),
('other', '马上来', '2026-03-25 12:03:00'),
('self', '陪玩服务怎么样？', '2026-03-25 13:00:00'),
('other', '很好，专业又耐心', '2026-03-25 13:01:00'),
('self', '价格贵吗？', '2026-03-25 13:02:00'),
('other', '很合理，物有所值', '2026-03-25 13:03:00'),
('self', '推荐一个陪玩专家吧', '2026-03-25 13:04:00'),
('other', 'id1的陪玩专家不错', '2026-03-25 13:05:00'),
('self', '谢谢推荐', '2026-03-25 13:06:00'),
('other', '不客气', '2026-03-25 13:07:00'),
('self', '游戏主播的直播很精彩', '2026-03-25 14:00:00'),
('other', '是的，很专业', '2026-03-25 14:01:00'),
('self', '有什么游戏推荐吗？', '2026-03-25 14:02:00'),
('other', '最近新出的游戏不错', '2026-03-25 14:03:00'),
('self', '新手怎么快速上手？', '2026-03-25 15:00:00'),
('other', '多看看教程，多练习', '2026-03-25 15:01:00'),
('self', '好的，谢谢', '2026-03-25 15:02:00'),
('other', '不客气，加油', '2026-03-25 15:03:00'),
('self', '团战怎么配合？', '2026-03-25 16:00:00'),
('other', '沟通很重要，听指挥', '2026-03-25 16:01:00'),
('self', '明白了', '2026-03-25 16:02:00'),
('other', '多练习就好了', '2026-03-25 16:03:00'),
('self', '单挑有什么技巧？', '2026-03-25 17:00:00'),
('other', '走位，技能释放时机', '2026-03-25 17:01:00'),
('self', '谢谢指导', '2026-03-25 17:02:00'),
('other', '没问题', '2026-03-25 17:03:00');

-- 创建订单表
CREATE TABLE IF NOT EXISTS game_partner_orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    user_id INT NOT NULL,
    playmate_id INT NOT NULL,
    game VARCHAR(100) NOT NULL,
    skill VARCHAR(100) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    service_time VARCHAR(50) DEFAULT '',
    amount FLOAT DEFAULT 0,
    order_number VARCHAR(50) DEFAULT '',
    payment_method VARCHAR(20) DEFAULT '',
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id),
    FOREIGN KEY (playmate_id) REFERENCES game_partner_playmates(id)
);

-- 创建订单确认表
CREATE TABLE IF NOT EXISTS game_partner_order_confirmations (
    id INT PRIMARY KEY AUTO_INCREMENT,
    order_id INT NOT NULL,
    price_per_hour FLOAT DEFAULT 0,
    duration INT DEFAULT 1,
    service_fee FLOAT DEFAULT 0,
    coupon_id INT NULL,
    total_amount FLOAT DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES game_partner_orders(id)
);

-- 创建奖励订单表
CREATE TABLE IF NOT EXISTS game_partner_reward_orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    user_id INT NOT NULL,
    game VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    reward FLOAT DEFAULT 0,
    payment_method VARCHAR(20) DEFAULT 'prepay',
    status VARCHAR(20) DEFAULT 'available',
    tags VARCHAR(255) DEFAULT '',
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id)
);

-- 创建优惠券表
CREATE TABLE IF NOT EXISTS game_partner_coupons (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    discount FLOAT DEFAULT 0,
    description TEXT,
    valid_until DATETIME NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 创建评价表
CREATE TABLE IF NOT EXISTS game_partner_reviews (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    user_id INT NOT NULL,
    playmate_id INT NOT NULL,
    rating INT DEFAULT 5,
    content TEXT NOT NULL,
    images VARCHAR(4096) DEFAULT '',
    tags VARCHAR(255) DEFAULT '',
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id),
    FOREIGN KEY (playmate_id) REFERENCES game_partner_playmates(id)
);

-- 创建提现表
CREATE TABLE IF NOT EXISTS game_partner_withdrawals (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    user_id INT NOT NULL,
    amount FLOAT DEFAULT 0,
    fee FLOAT DEFAULT 0,
    actual_amount FLOAT DEFAULT 0,
    method VARCHAR(20) DEFAULT 'wechat',
    status VARCHAR(20) DEFAULT 'pending',
    completed_at DATETIME NULL,
    failed_reason TEXT,
    FOREIGN KEY (user_id) REFERENCES game_partner_users(id)
);

-- 创建游戏表
CREATE TABLE IF NOT EXISTS game_partner_games (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    name VARCHAR(100) NOT NULL,
    icon VARCHAR(255) DEFAULT '',
    category VARCHAR(50) DEFAULT '',
    image VARCHAR(2048) DEFAULT ''
);

-- 创建活动表
CREATE TABLE IF NOT EXISTS game_partner_activities (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    discount FLOAT DEFAULT 0,
    type VARCHAR(20) DEFAULT 'discount',
    valid_until DATETIME NOT NULL
);

-- 创建分类表
CREATE TABLE IF NOT EXISTS game_partner_categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    name VARCHAR(50) NOT NULL,
    icon VARCHAR(255) DEFAULT ''
);

-- 创建游戏分类表
CREATE TABLE IF NOT EXISTS game_partner_game_categories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    name VARCHAR(50) NOT NULL,
    label VARCHAR(50) NOT NULL
);

-- 插入订单数据
INSERT INTO game_partner_orders (user_id, playmate_id, game, skill, status, service_time, amount, order_number, payment_method) VALUES
(1, 1, '英雄联盟', '上单', 'completed', '2026-03-20 14:00-16:00', 160, 'ORD20260320001', 'wechat'),
(2, 2, '王者荣耀', '打野', 'completed', '2026-03-21 19:00-21:00', 120, 'ORD20260321002', 'alipay'),
(3, 3, '和平精英', '刚枪', 'completed', '2026-03-22 15:00-17:00', 100, 'ORD20260322003', 'wechat'),
(4, 4, '英雄联盟', '团战', 'completed', '2026-03-19 18:00-20:00', 90, 'ORD20260319004', 'alipay'),
(5, 5, '王者荣耀', '单挑', 'completed', '2026-03-18 20:00-22:00', 110, 'ORD20260318005', 'wechat'),
(6, 6, '和平精英', '策略', 'completed', '2026-03-23 16:00-18:00', 96, 'ORD20260323006', 'alipay'),
(7, 7, '英雄联盟', '解说', 'completed', '2026-03-20 10:00-12:00', 140, 'ORD20260320007', 'wechat'),
(8, 8, '王者荣耀', '职业', 'completed', '2026-03-21 14:00-16:00', 200, 'ORD20260321008', 'alipay'),
(9, 9, '英雄联盟', '团队', 'completed', '2026-03-22 19:00-21:00', 130, 'ORD20260322009', 'wechat'),
(10, 10, '和平精英', '技术', 'completed', '2026-03-19 15:00-17:00', 180, 'ORD20260319010', 'alipay'),
(11, 11, '王者荣耀', '主播', 'completed', '2026-03-18 16:00-18:00', 150, 'ORD20260318011', 'wechat'),
(12, 12, '英雄联盟', '设计', 'completed', '2026-03-23 10:00-12:00', 120, 'ORD20260323012', 'alipay'),
(13, 13, '和平精英', '竞技', 'completed', '2026-03-20 13:00-15:00', 116, 'ORD20260320013', 'wechat'),
(14, 14, '王者荣耀', '测评', 'completed', '2026-03-21 10:00-12:00', 80, 'ORD20260321014', 'alipay'),
(15, 15, '和平精英', '新手指导', 'completed', '2026-03-22 14:00-16:00', 60, 'ORD20260322015', 'wechat'),
(16, 16, '和平精英', '休闲娱乐', 'completed', '2026-03-19 19:00-21:00', 50, 'ORD20260319016', 'alipay'),
(17, 17, '王者荣耀', '新手教学', 'completed', '2026-03-18 14:00-16:00', 40, 'ORD20260318017', 'wechat'),
(18, 18, '和平精英', '休闲娱乐', 'completed', '2026-03-23 19:00-21:00', 56, 'ORD20260323018', 'alipay'),
(19, 19, '英雄联盟', '游戏爱好', 'completed', '2026-03-20 16:00-18:00', 70, 'ORD20260320019', 'wechat'),
(20, 20, '王者荣耀', '多游戏精通', 'completed', '2026-03-21 16:00-18:00', 84, 'ORD20260321020', 'alipay');

-- 插入订单确认数据
INSERT INTO game_partner_order_confirmations (order_id, price_per_hour, duration, service_fee, total_amount) VALUES
(1, 80, 2, 0, 160),
(2, 60, 2, 0, 120),
(3, 50, 2, 0, 100),
(4, 45, 2, 0, 90),
(5, 55, 2, 0, 110),
(6, 48, 2, 0, 96),
(7, 70, 2, 0, 140),
(8, 100, 2, 0, 200),
(9, 65, 2, 0, 130),
(10, 90, 2, 0, 180),
(11, 75, 2, 0, 150),
(12, 60, 2, 0, 120),
(13, 58, 2, 0, 116),
(14, 40, 2, 0, 80),
(15, 30, 2, 0, 60),
(16, 25, 2, 0, 50),
(17, 20, 2, 0, 40),
(18, 28, 2, 0, 56),
(19, 35, 2, 0, 70),
(20, 42, 2, 0, 84);

-- 插入奖励订单数据
INSERT INTO game_partner_reward_orders (user_id, game, content, reward, payment_method, status, tags) VALUES
(1, '英雄联盟', '寻找钻石以上段位的陪玩，一起冲分', 200, 'postpay', 'completed', '钻石,冲分,英雄联盟'),
(2, '王者荣耀', '寻找王者段位的陪玩，打排位赛', 250, 'postpay', 'completed', '王者,排位,王者荣耀'),
(3, '和平精英', '寻找王牌段位的陪玩，一起吃鸡', 150, 'postpay', 'completed', '王牌,吃鸡,和平精英'),
(4, '英雄联盟', '寻找大师段位的陪玩，学习技术', 300, 'prepay', 'completed', '大师,技术,英雄联盟'),
(5, '王者荣耀', '寻找荣耀王者的陪玩，学习意识', 350, 'prepay', 'completed', '荣耀王者,意识,王者荣耀'),
(6, '和平精英', '寻找无敌战神的陪玩，学习操作', 280, 'postpay', 'completed', '无敌战神,操作,和平精英'),
(7, '英雄联盟', '寻找职业选手的陪玩，学习战术', 400, 'prepay', 'completed', '职业选手,战术,英雄联盟'),
(8, '王者荣耀', '寻找主播的陪玩，一起开黑', 220, 'postpay', 'completed', '主播,开黑,王者荣耀'),
(9, '和平精英', '寻找主播的陪玩，一起娱乐', 180, 'postpay', 'completed', '主播,娱乐,和平精英'),
(10, '英雄联盟', '寻找解说的陪玩，学习游戏理解', 260, 'prepay', 'completed', '解说,游戏理解,英雄联盟'),
(11, '王者荣耀', '寻找教练的陪玩，提升技术', 320, 'prepay', 'completed', '教练,技术提升,王者荣耀'),
(12, '和平精英', '寻找战术大师的陪玩，学习策略', 240, 'postpay', 'completed', '战术大师,策略,和平精英'),
(13, '英雄联盟', '寻找团战大师的陪玩，学习配合', 200, 'postpay', 'completed', '团战大师,配合,英雄联盟'),
(14, '王者荣耀', '寻找单挑王的陪玩，学习操作', 280, 'prepay', 'completed', '单挑王,操作,王者荣耀'),
(15, '和平精英', '寻找刚枪王的陪玩，学习枪法', 220, 'postpay', 'completed', '刚枪王,枪法,和平精英'),
(16, '英雄联盟', '寻找ADC大神的陪玩，学习走位', 250, 'prepay', 'completed', 'ADC,走位,英雄联盟'),
(17, '王者荣耀', '寻找中单法王的陪玩，学习输出', 260, 'postpay', 'completed', '中单,输出,王者荣耀'),
(18, '和平精英', '寻找狙击手的陪玩，学习瞄准', 230, 'postpay', 'completed', '狙击手,瞄准,和平精英'),
(19, '英雄联盟', '寻找上单霸主的陪玩，学习对线', 240, 'prepay', 'completed', '上单,对线,英雄联盟'),
(20, '王者荣耀', '寻找辅助大师的陪玩，学习保护', 210, 'postpay', 'completed', '辅助,保护,王者荣耀');

-- 插入优惠券数据
INSERT INTO game_partner_coupons (name, discount, description, valid_until) VALUES
('新用户首单8折', 0.8, '新用户首次下单享受8折优惠', '2026-12-31 23:59:59'),
('周末特惠9折', 0.9, '周末下单享受9折优惠', '2026-12-31 23:59:59'),
('会员专属7折', 0.7, 'VIP会员专享7折优惠', '2026-12-31 23:59:59'),
('满100减20', 0.8, '订单满100元立减20元', '2026-12-31 23:59:59'),
('节日特惠85折', 0.85, '节假日下单享受85折优惠', '2026-12-31 23:59:59'),
('推荐好友返现', 0.9, '推荐好友注册并下单，双方享受9折优惠', '2026-12-31 23:59:59'),
('连续下单优惠', 0.85, '连续下单3次及以上享受85折优惠', '2026-12-31 23:59:59'),
('游戏专属优惠', 0.8, '指定游戏下单享受8折优惠', '2026-12-31 23:59:59'),
('陪玩专家优惠', 0.9, '选择陪玩专家服务享受9折优惠', '2026-12-31 23:59:59'),
('深夜特惠75折', 0.75, '22:00-06:00下单享受75折优惠', '2026-12-31 23:59:59'),
('工作日优惠', 0.9, '工作日下单享受9折优惠', '2026-12-31 23:59:59'),
('长期合作优惠', 0.7, '累计下单10次及以上享受7折优惠', '2026-12-31 23:59:59'),
('特殊活动优惠', 0.8, '参与平台活动获得的8折优惠券', '2026-12-31 23:59:59'),
('生日特惠', 0.7, '生日当月下单享受7折优惠', '2026-12-31 23:59:59'),
('首次评价返现', 0.9, '首次评价订单后获得9折优惠券', '2026-12-31 23:59:59'),
('分享优惠', 0.85, '分享平台获得85折优惠券', '2026-12-31 23:59:59'),
('充值优惠', 0.8, '充值满1000元获得8折优惠券', '2026-12-31 23:59:59'),
('季节特惠', 0.85, '季节性活动85折优惠', '2026-12-31 23:59:59'),
('新手礼包', 0.75, '新手注册获得75折优惠券', '2026-12-31 23:59:59'),
('活动专属', 0.8, '特定活动期间的8折优惠券', '2026-12-31 23:59:59');

-- 插入评价数据
INSERT INTO game_partner_reviews (user_id, playmate_id, rating, content, images, tags) VALUES
(1, 1, 5, '非常专业的陪玩，技术很好，沟通也很愉快！', 'https://example.com/review1.jpg', '专业,技术好,沟通愉快'),
(2, 2, 5, '打野节奏很好，带着我赢了很多局，非常满意！', 'https://example.com/review2.jpg', '节奏好,胜率高,满意'),
(3, 3, 4, '刚枪很厉害，就是说话有点少，总体还是不错的', 'https://example.com/review3.jpg', '刚枪厉害,话少,不错'),
(4, 4, 5, '团战指挥很专业，团队配合默契，赢了很多团', 'https://example.com/review4.jpg', '指挥专业,配合默契,团战强'),
(5, 5, 5, '单挑真的无敌，学到了很多技巧，非常感谢！', 'https://example.com/review5.jpg', '单挑无敌,技巧多,感谢'),
(6, 6, 4, '策略制定很合理，就是反应有点慢，总体满意', 'https://example.com/review6.jpg', '策略合理,反应慢,满意'),
(7, 7, 5, '解说很专业，讲解详细，学到了很多游戏知识', 'https://example.com/review7.jpg', '解说专业,讲解详细,知识丰富'),
(8, 8, 5, '职业选手就是不一样，技术顶尖，意识一流', 'https://example.com/review8.jpg', '技术顶尖,意识一流,职业水准'),
(9, 9, 4, '团队指挥不错，就是有时候脾气有点急', 'https://example.com/review9.jpg', '指挥不错,脾气急,还行'),
(10, 10, 5, '技术真的天花板，操作太秀了，学到了很多', 'https://example.com/review10.jpg', '技术天花板,操作秀,学习'),
(11, 11, 5, '主播很有趣，直播效果好，玩得很开心', 'https://example.com/review11.jpg', '有趣,直播效果好,开心'),
(12, 12, 4, '游戏理解很深，就是话有点多，总体不错', 'https://example.com/review12.jpg', '理解深,话多,不错'),
(13, 13, 5, '竞技水平很高，反应迅速，配合默契', 'https://example.com/review13.jpg', '水平高,反应快,配合好'),
(14, 14, 4, '测评很专业，分析详细，就是价格有点贵', 'https://example.com/review14.jpg', '专业,详细,价格贵'),
(15, 15, 5, '新手指导很耐心，讲解详细，进步很大', 'https://example.com/review15.jpg', '耐心,详细,进步大'),
(16, 16, 4, '休闲娱乐还可以，就是技术一般，玩得开心', 'https://example.com/review16.jpg', '休闲,技术一般,开心'),
(17, 17, 4, '新手教学很耐心，就是经验不足，总体满意', 'https://example.com/review17.jpg', '耐心,经验不足,满意'),
(18, 18, 5, '休闲娱乐非常愉快，氛围好，玩得开心', 'https://example.com/review18.jpg', '愉快,氛围好,开心'),
(19, 19, 4, '游戏知识丰富，就是操作一般，学到了很多', 'https://example.com/review19.jpg', '知识丰富,操作一般,学习'),
(20, 20, 5, '多游戏都很精通，技术全面，非常满意', 'https://example.com/review20.jpg', '全面,精通,满意');

-- 插入提现数据
INSERT INTO game_partner_withdrawals (user_id, amount, fee, actual_amount, method, status, completed_at, failed_reason) VALUES
(1, 500, 5, 495, 'wechat', 'completed', '2026-03-20 10:00:00', NULL),
(2, 300, 3, 297, 'alipay', 'completed', '2026-03-21 15:30:00', NULL),
(3, 200, 2, 198, 'wechat', 'completed', '2026-03-22 09:15:00', NULL),
(4, 1000, 10, 990, 'alipay', 'completed', '2026-03-19 14:45:00', NULL),
(5, 800, 8, 792, 'wechat', 'completed', '2026-03-18 20:30:00', NULL),
(6, 100, 1, 99, 'alipay', 'completed', '2026-03-23 10:45:00', NULL),
(7, 400, 4, 396, 'wechat', 'completed', '2026-03-20 12:15:00', NULL),
(8, 600, 6, 594, 'alipay', 'completed', '2026-03-19 15:30:00', NULL),
(9, 300, 3, 297, 'wechat', 'completed', '2026-03-22 11:20:00', NULL),
(10, 1200, 12, 1188, 'alipay', 'completed', '2026-03-18 21:45:00', NULL),
(11, 1500, 15, 1485, 'wechat', 'completed', '2026-03-17 16:30:00', NULL),
(12, 900, 9, 891, 'alipay', 'completed', '2026-03-23 14:20:00', NULL),
(13, 150, 1.5, 148.5, 'wechat', 'completed', '2026-03-24 09:30:00', NULL),
(14, 700, 7, 693, 'alipay', 'completed', '2026-03-20 13:45:00', NULL),
(15, 400, 4, 396, 'wechat', 'completed', '2026-03-19 10:15:00', NULL),
(16, 300, 3, 297, 'alipay', 'completed', '2026-03-22 15:40:00', NULL),
(17, 800, 8, 792, 'wechat', 'completed', '2026-03-18 14:30:00', NULL),
(18, 200, 2, 198, 'alipay', 'completed', '2026-03-24 11:25:00', NULL),
(19, 1400, 14, 1386, 'wechat', 'completed', '2026-03-17 15:20:00', NULL),
(20, 500, 5, 495, 'alipay', 'completed', '2026-03-20 16:45:00', NULL);

-- 插入游戏数据
INSERT INTO game_partner_games (name, icon, category, image) VALUES
('英雄联盟', 'https://example.com/icons/lol.png', 'MOBA', 'https://example.com/images/lol.jpg'),
('王者荣耀', 'https://example.com/icons/wzry.png', 'MOBA', 'https://example.com/images/wzry.jpg'),
('和平精英', 'https://example.com/icons/hpjy.png', 'FPS', 'https://example.com/images/hpjy.jpg'),
('绝地求生', 'https://example.com/icons/pubg.png', 'FPS', 'https://example.com/images/pubg.jpg'),
('CS2', 'https://example.com/icons/cs2.png', 'FPS', 'https://example.com/images/cs2.jpg'),
('DOTA2', 'https://example.com/icons/dota2.png', 'MOBA', 'https://example.com/images/dota2.jpg'),
('守望先锋', 'https://example.com/icons/ow.png', 'FPS', 'https://example.com/images/ow.jpg'),
('原神', 'https://example.com/icons/ys.png', 'RPG', 'https://example.com/images/ys.jpg'),
('塞尔达传说', 'https://example.com/icons/zelda.png', 'RPG', 'https://example.com/images/zelda.jpg'),
('超级马里奥', 'https://example.com/icons/mario.png', 'Platformer', 'https://example.com/images/mario.jpg'),
('刺客信条', 'https://example.com/icons/ac.png', 'Action', 'https://example.com/images/ac.jpg'),
('GTA5', 'https://example.com/icons/gta5.png', 'Open World', 'https://example.com/images/gta5.jpg'),
('赛博朋克2077', 'https://example.com/icons/cyberpunk.png', 'RPG', 'https://example.com/images/cyberpunk.jpg'),
('英雄联盟手游', 'https://example.com/icons/lolm.png', 'MOBA', 'https://example.com/images/lolm.jpg'),
('王者荣耀国际版', 'https://example.com/icons/arena.png', 'MOBA', 'https://example.com/images/arena.jpg'),
('穿越火线', 'https://example.com/icons/cf.png', 'FPS', 'https://example.com/images/cf.jpg'),
('使命召唤', 'https://example.com/icons/cod.png', 'FPS', 'https://example.com/images/cod.jpg'),
('堡垒之夜', 'https://example.com/icons/fortnite.png', 'Battle Royale', 'https://example.com/images/fortnite.jpg'),
('Apex英雄', 'https://example.com/icons/apex.png', 'Battle Royale', 'https://example.com/images/apex.jpg'),
('VALORANT', 'https://example.com/icons/valorant.png', 'FPS', 'https://example.com/images/valorant.jpg');

-- 插入活动数据
INSERT INTO game_partner_activities (title, description, discount, type, valid_until) VALUES
('周末开黑派对', '周末组队开黑，享受额外折扣', 0.8, 'weekend', '2026-12-31 23:59:59'),
('王者荣耀争霸赛', '参与王者荣耀比赛，赢取丰厚奖励', 0.9, 'competition', '2026-12-31 23:59:59'),
('和平精英生存挑战', '和平精英生存挑战活动，挑战自我', 0.85, 'event', '2026-12-31 23:59:59'),
('新手训练营', '新手专属训练营，快速上手游戏', 0.7, 'newbie', '2026-12-31 23:59:59'),
('陪玩技能大赛', '陪玩技能大赛，展示专业水平', 0.8, 'competition', '2026-12-31 23:59:59'),
('游戏设计分享会', '游戏设计分享会，了解游戏开发', 0.9, 'event', '2026-12-31 23:59:59'),
('电竞明星见面会', '电竞明星见面会，近距离接触偶像', 0.95, 'event', '2026-12-31 23:59:59'),
('夏日狂欢节', '夏日游戏狂欢节，多重优惠等你来', 0.75, 'seasonal', '2026-12-31 23:59:59'),
('国庆特惠', '国庆期间下单享受特别优惠', 0.8, 'holiday', '2026-12-31 23:59:59'),
('周年庆典', '平台周年庆典，豪礼送不停', 0.7, 'anniversary', '2026-12-31 23:59:59'),
('寒假活动', '寒假期间的特别活动', 0.85, 'seasonal', '2026-12-31 23:59:59'),
('春节特惠', '春节期间的特别优惠', 0.7, 'holiday', '2026-12-31 23:59:59'),
('情人节活动', '情人节特别活动，双人开黑优惠', 0.8, 'holiday', '2026-12-31 23:59:59'),
('清明节活动', '清明节特别活动', 0.9, 'holiday', '2026-12-31 23:59:59'),
('劳动节特惠', '劳动节期间的特别优惠', 0.85, 'holiday', '2026-12-31 23:59:59'),
('端午节活动', '端午节特别活动', 0.9, 'holiday', '2026-12-31 23:59:59'),
('中秋节特惠', '中秋节期间的特别优惠', 0.8, 'holiday', '2026-12-31 23:59:59'),
('重阳节活动', '重阳节特别活动', 0.9, 'holiday', '2026-12-31 23:59:59'),
('双11特惠', '双11期间的特别优惠', 0.6, 'promotion', '2026-12-31 23:59:59'),
('圣诞节活动', '圣诞节特别活动', 0.8, 'holiday', '2026-12-31 23:59:59');

-- 插入分类数据
INSERT INTO game_partner_categories (name, icon) VALUES
('MOBA', 'https://example.com/icons/moba.png'),
('FPS', 'https://example.com/icons/fps.png'),
('RPG', 'https://example.com/icons/rpg.png'),
('Action', 'https://example.com/icons/action.png'),
('Adventure', 'https://example.com/icons/adventure.png'),
('Strategy', 'https://example.com/icons/strategy.png'),
('Simulation', 'https://example.com/icons/simulation.png'),
('Puzzle', 'https://example.com/icons/puzzle.png'),
('Sports', 'https://example.com/icons/sports.png'),
('Racing', 'https://example.com/icons/racing.png'),
('Fighting', 'https://example.com/icons/fighting.png'),
('Platformer', 'https://example.com/icons/platformer.png'),
('Musical', 'https://example.com/icons/musical.png'),
('Casual', 'https://example.com/icons/casual.png'),
('Board', 'https://example.com/icons/board.png'),
('Card', 'https://example.com/icons/card.png'),
('Educational', 'https://example.com/icons/educational.png'),
('Horror', 'https://example.com/icons/horror.png'),
('Sci-Fi', 'https://example.com/icons/scifi.png'),
('Fantasy', 'https://example.com/icons/fantasy.png');

-- 插入游戏分类数据
INSERT INTO game_partner_game_categories (name, label) VALUES
('MOBA', '多人在线战术竞技'),
('FPS', '第一人称射击'),
('RPG', '角色扮演'),
('Action', '动作'),
('Adventure', '冒险'),
('Strategy', '策略'),
('Simulation', '模拟'),
('Puzzle', '益智'),
('Sports', '体育'),
('Racing', '竞速'),
('Fighting', '格斗'),
('Platformer', '平台'),
('Musical', '音乐'),
('Casual', '休闲'),
('Board', '桌面'),
('Card', '卡牌'),
('Educational', '教育'),
('Horror', '恐怖'),
('Sci-Fi', '科幻'),
('Fantasy', '奇幻');