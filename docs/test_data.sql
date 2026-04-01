-- 清空现有数据
DELETE FROM game_partner_tournament_matches;
DELETE FROM game_partner_tournament_teams;
DELETE FROM game_partner_tournament_registrations;
DELETE FROM game_partner_tournaments;

-- 赛事数据
INSERT INTO game_partner_tournaments (id, title, description, cover, game, status, register_start, register_end, match_start, match_end, prize, participants, rules) VALUES 
(1, '王者荣耀春季联赛', '王者荣耀春季联赛是官方举办的大型电竞赛事，汇聚了全服最顶尖的选手。赛事分为海选赛、淘汰赛和决赛三个阶段，最终决出冠军。冠军将获得丰厚奖金和限定皮肤。', 'https://lh3.googleusercontent.com/aida-public/AB6AXuCJCdVR4SWlVmZS2MfQxv4neWDHVSFi46iFu7fIlPuVcLUzoMaMncSZHdLUkeV6uno5pL3MvmUa7m5q3jQFEiIa1tlYhxIxW43ul0iN37eeYHJZEAZD_Nonsn3SrL3j3htSrp0l3GDtWDA4dsGL-GTGfRJU5k7W99I3RnHFfw_bieodydRDntxEspPO8D_yu3K5n-8DSp_x_AKb77wFMUJe9DzUntWS-mfd2UsJzyUft_2rZPiJ1jRBVJQqDMGpVHTdG6aRdOaK2tU', '王者荣耀', 'upcoming', '2026-04-01 00:00:00', '2026-04-05 23:59:59', '2026-04-10 00:00:00', '2026-04-15 23:59:59', '¥10,000', 128, '["采用BO3淘汰赛制", "参赛者需达到指定段位要求", "禁止使用任何作弊手段", "最终解释权归主办方所有"]'),
(2, '英雄联盟全球总决赛预选赛', '英雄联盟全球总决赛预选赛是国内最高水平的电竞赛事之一，汇聚了各路高手。赛事分为小组赛、淘汰赛和决赛三个阶段。', 'https://lh3.googleusercontent.com/aida-public/AB6AXuDHdJE8Ql8TnJGS2MfQxv4neWDHVSFi46iFu7fIlPuVcLUzoMaMncSZHdLUkeV6uno5pL3MvmUa7m5q3jQFEiIa1tlYhxIxW43ul0iN37eeYHJZEAZD_Nonsn3SrL3j3htSrp0l3GDtWDA4dsGL-GTGfRJU5k7W99I3RnHFfw_bieodydRDntxEspPO8D_yu3K5n-8DSp_x_AKb77wFMUJe9DzUntWS-mfd2UsJzyUft_2rZPiJ1jRBVJQqDMGpVHTdG6aRdOaK2tU', '英雄联盟', 'ongoing', '2026-03-01 00:00:00', '2026-03-10 23:59:59', '2026-03-15 00:00:00', '2026-03-25 23:59:59', '¥50,000', 256, '["采用BO5淘汰赛制", "必须使用正式服账号参赛", "禁止使用任何第三方软件", "保持良好的游戏态度"]'),
(3, '绝地求生杯首届邀请赛', '绝地求生杯首届邀请赛，邀请了全服最顶尖的选手参与。赛事采用积分制，最终根据总积分排名。', 'https://lh3.googleusercontent.com/aida-public/AB6AXuEIfKf9Rm9UoJHS3NfRyx5ofXEOJWIh46jGu8jIoPwVcMUz1obSndTZHmMklI6vunp6pL4NvnUb8n6q4kRGjJhJa2tlYhxIxW43ul0iN37eeYHJZEAZD_Nonsn3SrL3j3htSrp0l3GDtWDA4dsGL-GTGfRJU5k7W99I3RnHFfw_bieodydRDntxEspPO8D_yu3K5n-8DSp_x_AKb77wFMUJe9DzUntWS-mfd2UsJzyUft_2rZPiJ1jRBVJQqDMGpVHTdG6aRdOaK2tU', '绝地求生', 'completed', '2026-02-01 00:00:00', '2026-02-10 23:59:59', '2026-02-15 00:00:00', '2026-02-20 23:59:59', '¥30,000', 64, '["采用多轮积分制", "每轮比赛独立计分", "禁止开挂作弊", "最终解释权归主办方所有"]');

-- 队伍数据
INSERT INTO game_partner_tournament_teams (tournament_id, name, avatar, members, `rank`, status) VALUES
(1, 'WX战队', 'https://randomuser.me/api/portraits/men/1.jpg', 5, '冠军', 'approved'),
(1, 'QG战队', 'https://randomuser.me/api/portraits/men/2.jpg', 5, '亚军', 'approved'),
(1, 'AG战队', 'https://randomuser.me/api/portraits/men/3.jpg', 5, NULL, 'approved'),
(1, 'RNG战队', 'https://randomuser.me/api/portraits/men/4.jpg', 5, NULL, 'approved'),
(2, 'EDG战队', 'https://randomuser.me/api/portraits/men/5.jpg', 5, '冠军', 'approved'),
(2, 'RNG战队', 'https://randomuser.me/api/portraits/men/6.jpg', 5, '亚军', 'approved'),
(2, 'WE战队', 'https://randomuser.me/api/portraits/men/7.jpg', 5, NULL, 'approved'),
(2, 'BLG战队', 'https://randomuser.me/api/portraits/men/8.jpg', 5, NULL, 'approved'),
(3, '4AM战队', 'https://randomuser.me/api/portraits/men/9.jpg', 4, '冠军', 'approved'),
(3, 'IFTY战队', 'https://randomuser.me/api/portraits/men/10.jpg', 4, '亚军', 'approved');

-- 比赛数据
INSERT INTO game_partner_tournament_matches (tournament_id, round, match_time, team1_name, team1_avatar, team2_name, team2_avatar, score1, score2, status, winner_id) VALUES
(1, 1, '2026-04-10 14:00:00', 'WX战队', 'https://randomuser.me/api/portraits/men/1.jpg', 'QG战队', 'https://randomuser.me/api/portraits/men/2.jpg', 2, 1, 'completed', 1),
(1, 1, '2026-04-10 16:00:00', 'AG战队', 'https://randomuser.me/api/portraits/men/3.jpg', 'RNG战队', 'https://randomuser.me/api/portraits/men/4.jpg', NULL, NULL, 'upcoming', NULL),
(2, 1, '2026-03-15 10:00:00', 'EDG战队', 'https://randomuser.me/api/portraits/men/5.jpg', 'RNG战队', 'https://randomuser.me/api/portraits/men/6.jpg', 3, 2, 'completed', 5),
(2, 1, '2026-03-15 14:00:00', 'WE战队', 'https://randomuser.me/api/portraits/men/7.jpg', 'BLG战队', 'https://randomuser.me/api/portraits/men/8.jpg', 1, 3, 'completed', 8),
(2, 2, '2026-03-20 14:00:00', 'EDG战队', 'https://randomuser.me/api/portraits/men/5.jpg', 'BLG战队', 'https://randomuser.me/api/portraits/men/8.jpg', NULL, NULL, 'ongoing', NULL),
(3, 1, '2026-02-15 10:00:00', '4AM战队', 'https://randomuser.me/api/portraits/men/9.jpg', 'IFTY战队', 'https://randomuser.me/api/portraits/men/10.jpg', 2, 0, 'completed', 9);