# Playmate Reward Orders API 修改记录

## 修改概述
为 playmate 插件添加 reward-orders 完整的 CRUD 接口功能。

## 后端修改

### 1. 新增文件

#### server/plugin/playmate/api/reward_order_api.go
- 创建奖励订单API层
- 实现接口：
  - `GetRewardOrders` - 获取奖励订单列表 (GET /playmate/reward-orders)
  - `GetRewardOrderDetail` - 获取奖励订单详情 (GET /playmate/reward-orders/:id)
  - `CreateRewardOrder` - 创建奖励订单 (POST /playmate/reward-orders)
  - `UpdateRewardOrder` - 更新奖励订单 (PUT /playmate/reward-orders/:id)
  - `DeleteRewardOrder` - 删除奖励订单 (DELETE /playmate/reward-orders/:id)
  - `GrabRewardOrder` - 抢奖励订单 (POST /playmate/reward-orders/:id/grab)

#### server/plugin/playmate/router/reward_order_router.go
- 创建奖励订单路由层
- 注册路由：
  - GET /reward-orders - 获取列表
  - GET /reward-orders/:id - 获取详情
  - POST /reward-orders - 创建订单
  - PUT /reward-orders/:id - 更新订单
  - DELETE /reward-orders/:id - 删除订单
  - POST /reward-orders/:id/grab - 抢单

### 2. 修改文件

#### server/plugin/playmate/api/enter.go
- 在 ApiGroup 结构体中添加 `RewardOrderApi` 字段

#### server/plugin/playmate/router/enter.go
- 在 RouterGroup 结构体中添加 `RewardOrderRouter` 字段

#### server/plugin/playmate/initialize/router.go
- 添加奖励订单路由初始化调用：`router.RouterGroupApp.InitRewardOrderRouter(routerGroup)`

#### server/plugin/playmate/model/request/request.go
- 新增 `RewardOrderSearch` 结构体用于奖励订单搜索请求参数
- 新增 `CreateRewardOrderRequest` 结构体用于创建奖励订单请求
- 新增 `UpdateRewardOrderRequest` 结构体用于更新奖励订单请求

#### server/plugin/playmate/service/reward_order_service.go
- 实现完整的CRUD业务逻辑：
  - `GetRewardOrders` - 查询列表（支持分页、筛选、搜索）
  - `GetRewardOrderDetail` - 查询详情
  - `CreateRewardOrder` - 创建订单
  - `UpdateRewardOrder` - 更新订单
  - `DeleteRewardOrder` - 删除订单
  - `GrabRewardOrder` - 抢单业务逻辑

## 前端修改

### docs/vue_api/playmate.js
前端API已包含 reward-orders 完整的CRUD接口：
- `getRewardOrders(params)` - 获取奖励订单列表
- `getRewardOrderDetail(orderId)` - 获取奖励订单详情
- `createRewardOrder(data)` - 创建奖励订单
- `updateRewardOrder(orderId, data)` - 更新奖励订单
- `deleteRewardOrder(orderId)` - 删除奖励订单
- `grabRewardOrder(orderId)` - 抢奖励订单

## API 接口列表

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | /playmate/reward-orders | 获取奖励订单列表 |
| GET | /playmate/reward-orders/:id | 获取奖励订单详情 |
| POST | /playmate/reward-orders | 创建奖励订单 |
| PUT | /playmate/reward-orders/:id | 更新奖励订单 |
| DELETE | /playmate/reward-orders/:id | 删除奖励订单 |
| POST | /playmate/reward-orders/:id/grab | 抢奖励订单 |

## 请求参数

### GetRewardOrders (列表查询)
- `game` (string, optional): 游戏筛选
- `status` (string, optional): 订单状态筛选 (available, grabbed, completed, cancelled)
- `paymentMethod` (string, optional): 支付方式筛选 (prepay, postpay)
- `keyword` (string, optional): 关键词搜索
- `page` (int, optional): 页码，默认1
- `pageSize` (int, optional): 每页数量，默认20

### GetRewardOrderDetail (详情查询)
- `id` (path parameter): 订单ID

### CreateRewardOrder (创建)
- `game` (string, required): 游戏名称
- `content` (string, required): 订单内容描述
- `reward` (number, required): 奖励金额，必须大于0
- `paymentMethod` (string, required): 支付方式 (prepay/postpay)
- `tags` (string, optional): 标签，逗号分隔

### UpdateRewardOrder (更新)
- `id` (path parameter): 订单ID
- `game` (string, optional): 游戏名称
- `content` (string, optional): 订单内容描述
- `reward` (number, optional): 奖励金额
- `paymentMethod` (string, optional): 支付方式
- `status` (string, optional): 订单状态
- `tags` (string, optional): 标签

### DeleteRewardOrder (删除)
- `id` (path parameter): 订单ID

### GrabRewardOrder (抢单)
- `id` (path parameter): 订单ID

## 数据模型

### RewardOrder
```go
type RewardOrder struct {
    ID            uint           `json:"id"`
    CreatedAt     time.Time      `json:"createdAt"`
    UpdatedAt     time.Time      `json:"updatedAt"`
    UserID        uint           `json:"userId"`
    Game          string         `json:"game"`
    Content       string         `json:"content"`
    Reward        float64        `json:"reward"`
    PaymentMethod string         `json:"paymentMethod"` // prepay, postpay
    Status        string         `json:"status"`       // available, grabbed, completed, cancelled
    Tags          string         `json:"tags"`
}
```

## 状态说明

- `available` - 可抢单状态
- `grabbed` - 已被抢
- `completed` - 已完成
- `cancelled` - 已取消

## 修改日期
2026-03-25
