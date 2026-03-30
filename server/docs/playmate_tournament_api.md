# 赛事管理模块 API 文档

## 概述

本文档描述了赛事管理模块的后端 API 接口实现，对应前端 `/System/Volumes/Data/webcode/UX/GamePartner/src/api/mock-tournament.js` 中的 Mock 数据接口。

## 数据模型

### Tournament (赛事)
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 赛事ID |
| title | string | 赛事标题 |
| description | string | 赛事描述 |
| cover | string | 封面图片URL |
| game | string | 游戏名称 |
| gameId | uint | 游戏ID |
| status | string | 赛事状态: upcoming-报名中, ongoing-进行中, completed-已结束 |
| registerStart | time.Time | 报名开始时间 |
| registerEnd | time.Time | 报名结束时间 |
| matchStart | time.Time | 比赛开始时间 |
| matchEnd | time.Time | 比赛结束时间 |
| prize | string | 奖金池 |
| participants | int | 当前参赛人数 |
| maxTeams | int | 最大参赛队伍数 |
| rules | string | 赛事规则(JSON格式) |
| format | string | 赛制(BO3/BO5等) |
| minRank | string | 最低段位要求 |
| organizer | string | 主办方 |
| contactInfo | string | 联系方式 |

### TournamentTeam (参赛队伍)
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 队伍ID |
| tournamentId | uint | 所属赛事ID |
| name | string | 队伍名称 |
| avatar | string | 队伍头像URL |
| members | int | 成员数量 |
| leaderId | uint | 队长用户ID |
| rank | string | 最终排名 |
| status | string | 状态: registered-已报名, approved-已通过, rejected-已拒绝 |

### TournamentMatch (比赛)
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 比赛ID |
| tournamentId | uint | 所属赛事ID |
| round | int | 轮次 |
| matchTime | time.Time | 比赛时间 |
| team1Id | uint | 队伍1 ID |
| team1Name | string | 队伍1名称 |
| team1Avatar | string | 队伍1头像 |
| team2Id | uint | 队伍2 ID |
| team2Name | string | 队伍2名称 |
| team2Avatar | string | 队伍2头像 |
| score1 | *int | 队伍1比分 |
| score2 | *int | 队伍2比分 |
| status | string | 状态: upcoming-未开始, ongoing-进行中, completed-已结束 |
| winnerId | *uint | 获胜队伍ID |

### TournamentRegistration (报名记录)
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 报名ID |
| tournamentId | uint | 赛事ID |
| userId | uint | 用户ID |
| teamName | string | 队伍名称 |
| contactInfo | string | 联系信息 |
| membersInfo | string | 成员信息(JSON) |
| status | string | 状态: pending-待审核, approved-已通过, rejected-已拒绝 |
| reviewRemark | string | 审核备注 |

## API 接口

### 公开接口

#### 1. 获取赛事列表
```
GET /playmate/tournaments
```
参数:
- status: 赛事状态 (upcoming/ongoing/completed)
- game: 游戏名称
- gameId: 游戏ID
- keyword: 关键词搜索
- page: 页码 (默认1)
- pageSize: 每页数量 (默认10)

响应:
```json
{
  "code": 0,
  "data": {
    "list": [...],
    "total": 100
  },
  "msg": "获取成功"
}
```

#### 2. 获取赛事详情
```
GET /playmate/tournaments/:id
```

#### 3. 获取赛事参赛队伍
```
GET /playmate/tournaments/:id/teams
```

#### 4. 获取赛事比赛列表
```
GET /playmate/tournaments/:id/matches
```

### 需要登录的接口

#### 5. 报名参赛
```
POST /playmate/tournaments/join
```
请求体:
```json
{
  "tournamentId": 1,
  "teamName": "战队名称",
  "contactInfo": "联系方式",
  "membersInfo": "成员信息JSON"
}
```

### 管理接口

#### 6. 创建赛事
```
POST /playmate/tournaments
```

#### 7. 更新赛事
```
PUT /playmate/tournaments/:id
```

#### 8. 删除赛事
```
DELETE /playmate/tournaments/:id
```

#### 9. 创建参赛队伍
```
POST /playmate/tournaments/teams
```

#### 10. 更新参赛队伍
```
PUT /playmate/tournaments/teams/:id
```

#### 11. 删除参赛队伍
```
DELETE /playmate/tournaments/teams/:id
```

#### 12. 创建比赛
```
POST /playmate/tournaments/matches
```

#### 13. 更新比赛
```
PUT /playmate/tournaments/matches/:id
```

#### 14. 删除比赛
```
DELETE /playmate/tournaments/matches/:id
```

## 文件结构

```
server/plugin/playmate/
├── model/
│   └── tournament.go              # 赛事数据模型
├── model/request/
│   └── request.go                 # 包含 TournamentSearch, JoinTournamentRequest
├── service/
│   ├── enter.go                   # 添加 TournamentService
│   └── tournament_service.go      # 赛事服务实现
├── api/
│   ├── enter.go                   # 添加 TournamentApi
│   └── tournament_api.go          # 赛事API实现
├── router/
│   ├── enter.go                   # 添加 TournamentRouter
│   └── tournament_router.go       # 赛事路由定义
└── initialize/
    ├── gorm.go                    # 添加赛事表迁移
    └── router.go                  # 添加赛事路由初始化
```

## 数据库表

- `game_partner_tournaments` - 赛事表
- `game_partner_tournament_teams` - 参赛队伍表
- `game_partner_tournament_matches` - 比赛表
- `game_partner_tournament_registrations` - 报名记录表

## 状态说明

### 赛事状态 (Tournament.status)
- `upcoming` - 报名中
- `ongoing` - 进行中
- `completed` - 已结束

### 比赛状态 (TournamentMatch.status)
- `upcoming` - 未开始
- `ongoing` - 进行中
- `completed` - 已结束

### 报名状态 (TournamentRegistration.status)
- `pending` - 待审核
- `approved` - 已通过
- `rejected` - 已拒绝
