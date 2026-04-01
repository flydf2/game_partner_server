-- 更新图片链接SQL语句
-- 使用本地真实图片链接随机更新数据库表中的图片字段

-- 使用game_partner数据库
USE game_partner;

-- 更新用户表的avatar字段（使用本地头像图片）
UPDATE game_partner_users SET avatar = ELT(
    FLOOR(RAND() * 10) + 1,
    '/public/avatar/p1.png',
    '/public/avatar/p2.png',
    '/public/avatar/p3.png',
    '/public/avatar/p4.png',
    '/public/avatar/p5.png',
    '/public/avatar/p6.png',
    '/public/avatar/p7.png',
    '/public/avatar/p8.png',
    '/public/avatar/p9.png',
    '/public/avatar/p10.png'
);

-- 更新陪玩专家表的avatar字段（使用本地头像图片）
UPDATE game_partner_playmates SET avatar = ELT(
    FLOOR(RAND() * 10) + 1,
    '/public/avatar/p1.png',
    '/public/avatar/p2.png',
    '/public/avatar/p3.png',
    '/public/avatar/p4.png',
    '/public/avatar/p5.png',
    '/public/avatar/p6.png',
    '/public/avatar/p7.png',
    '/public/avatar/p8.png',
    '/public/avatar/p9.png',
    '/public/avatar/p10.png'
);

-- 更新社区帖子表的images字段（随机选择1-3个图片链接，用逗号分隔）
UPDATE game_partner_community_posts SET images = CONCAT(
    ELT(
        FLOOR(RAND() * 16) + 1,
        '/public/avatar/p1.png',
        '/public/avatar/p2.png',
        '/public/avatar/p3.png',
        '/public/avatar/p4.png',
        '/public/avatar/p5.png',
        '/public/game/Dota_2.jpg',
        '/public/game/第五人格.jpg',
        '/public/game/和平精英.jpg',
        '/public/game/决战平安京.jpg',
        '/public/game/王者荣耀.jpg',
        '/public/game/英雄联盟.jpg',
        '/public/game/英雄联盟手游.jpg',
        '/public/game/原神.jpg',
        '/public/game/永劫无间.jpg',
        '/public/topic/topic1.png',
        '/public/topic/topic2.png'
    ),
    CASE WHEN RAND() > 0.5 THEN CONCAT(',', ELT(
        FLOOR(RAND() * 16) + 1,
        '/public/avatar/p1.png',
        '/public/avatar/p2.png',
        '/public/avatar/p3.png',
        '/public/avatar/p4.png',
        '/public/avatar/p5.png',
        '/public/game/Dota_2.jpg',
        '/public/game/第五人格.jpg',
        '/public/game/和平精英.jpg',
        '/public/game/决战平安京.jpg',
        '/public/game/王者荣耀.jpg',
        '/public/game/英雄联盟.jpg',
        '/public/game/英雄联盟手游.jpg',
        '/public/game/原神.jpg',
        '/public/game/永劫无间.jpg',
        '/public/topic/topic1.png',
        '/public/topic/topic2.png'
    )) ELSE '' END,
    CASE WHEN RAND() > 0.7 THEN CONCAT(',', ELT(
        FLOOR(RAND() * 16) + 1,
        '/public/avatar/p1.png',
        '/public/avatar/p2.png',
        '/public/avatar/p3.png',
        '/public/avatar/p4.png',
        '/public/avatar/p5.png',
        '/public/game/Dota_2.jpg',
        '/public/game/第五人格.jpg',
        '/public/game/和平精英.jpg',
        '/public/game/决战平安京.jpg',
        '/public/game/王者荣耀.jpg',
        '/public/game/英雄联盟.jpg',
        '/public/game/英雄联盟手游.jpg',
        '/public/game/原神.jpg',
        '/public/game/永劫无间.jpg',
        '/public/topic/topic1.png',
        '/public/topic/topic2.png'
    )) ELSE '' END
);

-- 更新评价表的images字段（随机选择1-2个图片链接，用逗号分隔）
UPDATE game_partner_reviews SET images = CONCAT(
    ELT(
        FLOOR(RAND() * 16) + 1,
        '/public/avatar/p1.png',
        '/public/avatar/p2.png',
        '/public/avatar/p3.png',
        '/public/avatar/p4.png',
        '/public/avatar/p5.png',
        '/public/game/Dota_2.jpg',
        '/public/game/第五人格.jpg',
        '/public/game/和平精英.jpg',
        '/public/game/决战平安京.jpg',
        '/public/game/王者荣耀.jpg',
        '/public/game/英雄联盟.jpg',
        '/public/game/英雄联盟手游.jpg',
        '/public/game/原神.jpg',
        '/public/game/永劫无间.jpg',
        '/public/topic/topic1.png',
        '/public/topic/topic2.png'
    ),
    CASE WHEN RAND() > 0.5 THEN CONCAT(',', ELT(
        FLOOR(RAND() * 16) + 1,
        '/public/avatar/p1.png',
        '/public/avatar/p2.png',
        '/public/avatar/p3.png',
        '/public/avatar/p4.png',
        '/public/avatar/p5.png',
        '/public/game/Dota_2.jpg',
        '/public/game/第五人格.jpg',
        '/public/game/和平精英.jpg',
        '/public/game/决战平安京.jpg',
        '/public/game/王者荣耀.jpg',
        '/public/game/英雄联盟.jpg',
        '/public/game/英雄联盟手游.jpg',
        '/public/game/原神.jpg',
        '/public/game/永劫无间.jpg',
        '/public/topic/topic1.png',
        '/public/topic/topic2.png'
    )) ELSE '' END
);

-- 更新游戏表的icon和image字段（使用本地游戏图片）
UPDATE game_partner_games SET 
    icon = ELT(
        FLOOR(RAND() * 9) + 1,
        '/public/game/Dota_2.jpg',
        '/public/game/第五人格.jpg',
        '/public/game/和平精英.jpg',
        '/public/game/决战平安京.jpg',
        '/public/game/王者荣耀.jpg',
        '/public/game/英雄联盟.jpg',
        '/public/game/英雄联盟手游.jpg',
        '/public/game/原神.jpg',
        '/public/game/永劫无间.jpg'
    ),
    image = ELT(
        FLOOR(RAND() * 9) + 1,
        '/public/game/Dota_2.jpg',
        '/public/game/第五人格.jpg',
        '/public/game/和平精英.jpg',
        '/public/game/决战平安京.jpg',
        '/public/game/王者荣耀.jpg',
        '/public/game/英雄联盟.jpg',
        '/public/game/英雄联盟手游.jpg',
        '/public/game/原神.jpg',
        '/public/game/永劫无间.jpg'
    );

-- 更新分类表的icon字段（使用本地话题图片）
UPDATE game_partner_categories SET icon = ELT(
    FLOOR(RAND() * 7) + 1,
    '/public/topic/topic1.png',
    '/public/topic/topic2.png',
    '/public/topic/topic3.png',
    '/public/topic/topic4.png',
    '/public/topic/topic5.png',
    '/public/topic/topic6.png',
    '/public/topic/topic7.png'
);

-- 查看更新结果
SELECT 'Users' as table_name, COUNT(*) as updated_rows FROM game_partner_users WHERE avatar != '';
SELECT 'Playmates' as table_name, COUNT(*) as updated_rows FROM game_partner_playmates WHERE avatar != '';
SELECT 'Community Posts' as table_name, COUNT(*) as updated_rows FROM game_partner_community_posts WHERE images != '';
SELECT 'Reviews' as table_name, COUNT(*) as updated_rows FROM game_partner_reviews WHERE images != '';
SELECT 'Games' as table_name, COUNT(*) as updated_rows FROM game_partner_games WHERE icon != '' AND image != '';
SELECT 'Categories' as table_name, COUNT(*) as updated_rows FROM game_partner_categories WHERE icon != '';