# Playmate 排行榜功能升级文档

## 修改概述

为陪玩系统的排行榜功能增加以下字段：
1. **类型** - 支持周榜(weekly)和月榜(monthly)两种类型
2. **名称** - 存储榜单名称，可以是话题名称、游戏名称等

## 后端修改

### 1. 新增模型文件

**文件**: `server/plugin/playmate/model/leaderboard.go`

新增两个模型：
- `Leaderboard` - 排行榜主表，包含名称、类型、游戏、时间范围等字段
- `LeaderboardItem` - 排行榜条目表，存储排名、陪玩信息、评分等

```go
// LeaderboardType 排行榜类型
type LeaderboardType string

const (
    LeaderboardTypeWeekly  LeaderboardType = "weekly"  // 周榜
    LeaderboardTypeMonthly LeaderboardType = "monthly" // 月榜
)
```

### 2. 新增请求模型

**文件**: `server/plugin/playmate/model/request/request.go`

新增请求结构体：
- `CreateLeaderboardRequest` - 创建排行榜请求
- `UpdateLeaderboardRequest` - 更新排行榜请求

### 3. 新增服务层

**文件**: `server/plugin/playmate/service/leaderboard_service.go`

主要功能：
- `GetLeaderboards` - 获取排行榜列表（支持按类型和游戏过滤）
- `GetLeaderboardById` - 根据ID获取排行榜详情
- `GetLeaderboardWithItems` - 获取排行榜及其条目
- `CreateLeaderboard` - 创建排行榜
- `UpdateLeaderboard` - 更新排行榜
- `DeleteLeaderboard` - 删除排行榜
- `GenerateLeaderboard` - 根据陪玩数据自动生成排名
- `AddLeaderboardItem` - 添加排行榜条目
- `GetLeaderboardItems` - 获取排行榜条目列表

### 4. 新增API层

**文件**: `server/plugin/playmate/api/leaderboard_api.go`

接口列表：
- `GET /playmate/leaderboards` - 获取排行榜列表
- `GET /playmate/leaderboards/:id` - 获取排行榜详情
- `GET /playmate/leaderboards/:id/items` - 获取排行榜及其条目
- `GET /playmate/leaderboards/:id/items-only` - 仅获取排行榜条目
- `POST /playmate/leaderboards` - 创建排行榜
- `PUT /playmate/leaderboards/:id` - 更新排行榜
- `DELETE /playmate/leaderboards/:id` - 删除排行榜
- `POST /playmate/leaderboards/:id/generate` - 生成排行榜数据

### 5. 新增路由

**文件**: `server/plugin/playmate/router/leaderboard_router.go`

### 6. 注册服务

修改文件：
- `server/plugin/playmate/service/enter.go` - 注册 LeaderboardService
- `server/plugin/playmate/api/enter.go` - 注册 LeaderboardApi
- `server/plugin/playmate/router/enter.go` - 注册 LeaderboardRouter
- `server/plugin/playmate/initialize/router.go` - 初始化排行榜路由
- `server/plugin/playmate/initialize/gorm.go` - 注册数据库模型

## 前端修改

### 1. 新增API接口

**文件**: `web/src/api/plugin/playmate.js`

新增接口函数：
- `getLeaderboards` - 获取排行榜列表
- `getLeaderboardById` - 获取排行榜详情
- `getLeaderboardWithItems` - 获取排行榜及其条目
- `createLeaderboard` - 创建排行榜
- `updateLeaderboard` - 更新排行榜
- `deleteLeaderboard` - 删除排行榜
- `generateLeaderboard` - 生成排行榜

### 2. 新增管理页面

**文件**: `web/src/plugin/playmate/view/leaderboardList.vue`

页面功能：
- 排行榜列表展示（支持按类型和游戏搜索）
- 创建/编辑排行榜（包含名称、类型、游戏、描述、时间范围、状态等字段）
- 查看排行榜详情（展示排名、陪玩信息、评分、订单数、点赞数）
- 生成排行榜数据
- 删除排行榜

## 数据库表结构

### game_partner_leaderboards 表

| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | uint | 主键 |
| name | varchar(255) | 榜单名称（话题名称、游戏名称等） |
| type | varchar(20) | 榜单类型（weekly-周榜, monthly-月榜） |
| game | varchar(100) | 关联游戏 |
| start_time | datetime | 开始时间 |
| end_time | datetime | 结束时间 |
| description | text | 描述 |
| status | int | 状态（1-启用，0-禁用） |
| sort_order | int | 排序顺序 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

### game_partner_leaderboard_items 表

| 字段名 | 类型 | 说明 |
|--------|------|------|
| id | uint | 主键 |
| leaderboard_id | uint | 排行榜ID |
| playmate_id | uint | 陪玩ID |
| rank | int | 排名 |
| score | float | 评分 |
| order_count | int | 订单数 |
| rating | float | 评分 |
| likes | int | 点赞数 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |

## 使用说明

### 创建排行榜

1. 进入排行榜管理页面
2. 点击"创建排行榜"按钮
3. 填写榜单信息：
   - 榜单名称：如"王者荣耀热门陪玩榜"、"话题挑战榜"等
   - 榜单类型：选择"周榜"或"月榜"
   - 关联游戏：选择关联的游戏
   - 描述：榜单的详细说明
   - 时间范围：榜单的有效期
   - 状态：启用或禁用
   - 排序：榜单的显示顺序

### 生成排行榜数据

1. 在排行榜列表中找到需要生成的榜单
2. 点击"生成榜单"按钮
3. 系统会自动根据陪玩的评分、订单数、点赞数等数据计算排名

### 查看排行榜

1. 点击"查看"按钮
2. 可以看到榜单的详细信息
3. 包括排名、陪玩头像、昵称、评分、订单数、点赞数等

## 注意事项

1. 排行榜类型一旦创建后建议不要修改，以免影响历史数据统计
2. 生成榜单会重新计算所有排名，原有数据会被清除
3. 禁用的榜单不会在前端展示
4. 榜单名称可以根据业务需求灵活设置，如话题名称、游戏名称、活动名称等
