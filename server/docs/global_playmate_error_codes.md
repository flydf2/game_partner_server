# Playmate 插件错误码接口文档

## 概述

Playmate 插件使用独立的错误码体系，基础错误码为 `10000`，每个业务场景都有唯一的错误码，便于前端精确处理和排查问题。

---

## 一、用户模块 (User) - 错误码 10001~10011

### 1.1 用户登录/注册

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/auth/login | 用户名不存在 | 10001 | 用户不存在 |
| POST /playmate/auth/login | 密码错误 | 10002 | 用户名或密码错误 |
| POST /playmate/auth/register | 用户名已存在 | 10003 | 用户名已存在 |
| POST /playmate/auth/register | 手机号已注册 | 10004 | 手机号已被注册 |
| PUT /playmate/user/profile | 手机号已被使用 | 10004 | 手机号已被注册 |

### 1.2 用户关注

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/user/following/{userId} | 重复关注 | 10005 | 已经关注过该用户 |
| POST /playmate/user/following/{userId} | 关注自己 | 10011 | 不能关注自己 |
| DELETE /playmate/user/following/{userId} | 未关注该用户 | 10006 | 未关注该用户 |

### 1.3 用户收藏

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| DELETE /playmate/user/favorites/{favoriteId} | 收藏不存在 | 10007 | 收藏不存在 |

---

## 二、钱包模块 (Wallet) - 错误码 10008~10010

### 2.1 钱包余额

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/order/createOrder | 钱包不存在 | 10009 | 钱包不存在 |
| POST /playmate/order/createOrder | 余额不足 | 10008 | 余额不足 |

### 2.2 提现

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/withdrawal/submit | 钱包不存在 | 10009 | 钱包不存在 |
| POST /playmate/withdrawal/submit | 金额格式错误 | 10010 | 金额格式错误 |
| POST /playmate/withdrawal/submit | 余额不足 | 10008 | 余额不足 |

---

## 三、订单模块 (Order) - 错误码 10101~10113

### 3.1 订单创建

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/order/createOrder | 陪玩不存在 | 10102 | 陪玩不存在 |
| POST /playmate/order/createOrder | 钱包不存在 | 10009 | 钱包不存在 |
| POST /playmate/order/createOrder | 余额不足 | 10008 | 余额不足 |

### 3.2 订单取消

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/order/cancel/{orderId} | 订单不存在 | 10101 | 订单不存在 |
| POST /playmate/order/cancel/{orderId} | 无权操作 | 10110 | 无权操作此订单 |
| POST /playmate/order/cancel/{orderId} | 状态不允许取消 | 10111 | 该订单状态无法取消 |

### 3.3 订单确认

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/order/confirm/{orderId} | 订单不存在 | 10101 | 订单不存在 |
| POST /playmate/order/confirm/{orderId} | 无权操作 | 10110 | 无权操作此订单 |
| POST /playmate/order/confirm/{orderId} | 状态不允许确认 | 10109 | 订单状态不允许确认服务 |

### 3.4 订单接受/拒绝

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/order/accept/{orderId} | 订单不存在 | 10101 | 订单不存在 |
| POST /playmate/order/accept/{orderId} | 无权操作 | 10110 | 无权操作此订单 |
| POST /playmate/order/accept/{orderId} | 状态不允许接受 | 10112 | 该订单状态无法接受 |
| POST /playmate/order/reject/{orderId} | 订单不存在 | 10101 | 订单不存在 |
| POST /playmate/order/reject/{orderId} | 无权操作 | 10110 | 无权操作此订单 |
| POST /playmate/order/reject/{orderId} | 状态不允许拒绝 | 10113 | 该订单状态无法拒绝 |

---

## 四、悬赏订单模块 (RewardOrder) - 错误码 10101~10110

### 4.1 悬赏订单基础操作

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| GET /playmate/rewardOrder/detail/{orderId} | 订单不存在 | 10101 | 订单不存在 |
| PUT /playmate/rewardOrder/update/{orderId} | 订单不存在 | 10101 | 订单不存在 |
| DELETE /playmate/rewardOrder/delete/{orderId} | 订单不存在 | 10101 | 订单不存在 |
| POST /playmate/rewardOrder/share/{orderId} | 订单不存在 | 10101 | 订单不存在 |
| POST /playmate/rewardOrder/share/{orderId} | 无权操作 | 10110 | 无权操作此订单 |

### 4.2 悬赏订单抢单

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/rewardOrder/grab/{orderId} | 订单不存在 | 10101 | 订单不存在 |
| POST /playmate/rewardOrder/grab/{orderId} | 订单不可抢 | 10103 | 订单不可抢 |
| POST /playmate/rewardOrder/grab/{orderId} | 重复抢单 | 10104 | 您已经抢过此订单 |
| GET /playmate/rewardOrder/applicants/{orderId} | 订单不存在 | 10101 | 订单不存在 |

### 4.3 悬赏订单选择抢单者

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/rewardOrder/selectApplicant | 订单不存在 | 10101 | 订单不存在 |
| POST /playmate/rewardOrder/selectApplicant | 抢单申请不存在 | 10105 | 抢单申请不存在 |
| POST /playmate/rewardOrder/selectApplicant | 抢单申请不匹配 | 10106 | 抢单申请不属于该订单 |

### 4.4 悬赏订单支付

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/rewardOrder/pay/{orderId} | 订单不存在 | 10101 | 订单不存在 |
| POST /playmate/rewardOrder/pay/{orderId} | 状态不允许支付 | 10107 | 订单状态不允许支付 |
| POST /playmate/rewardOrder/pay/{orderId} | 金额不符 | 10108 | 支付金额与订单金额不符 |

### 4.5 悬赏订单确认服务

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/rewardOrder/confirm/{orderId} | 订单不存在 | 10101 | 订单不存在 |
| POST /playmate/rewardOrder/confirm/{orderId} | 状态不允许确认 | 10109 | 订单状态不允许确认服务 |

---

## 五、申诉模块 (Appeal) - 错误码 10201~10203

### 5.1 申诉基础操作

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| GET /playmate/appeal/detail/{appealId} | 申诉不存在 | 10201 | 申诉不存在 |
| PUT /playmate/appeal/update/{appealId} | 申诉不存在 | 10201 | 申诉不存在 |
| PUT /playmate/appeal/update/{appealId} | 状态不允许更新 | 10202 | 只能更新待处理状态的申诉 |
| DELETE /playmate/appeal/delete/{appealId} | 申诉不存在 | 10201 | 申诉不存在 |

### 5.2 申诉处理

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/appeal/handle/{appealId} | 申诉不存在 | 10201 | 申诉不存在 |
| POST /playmate/appeal/handle/{appealId} | 申诉已处理 | 10203 | 该申诉已处理完成 |

---

## 六、通知模块 (Notification) - 错误码 10301

### 6.1 通知操作

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| PUT /playmate/notification/read/{notificationId} | 通知不存在 | 10301 | 通知不存在 |

---

## 七、消息模块 (Message) - 错误码 10401~10402

### 7.1 消息发送

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/message/send/{userId} | 用户不存在 | 10401 | 用户不存在 |

### 7.2 会话操作

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| PUT /playmate/message/conversation/read/{conversationId} | 会话不存在 | 10402 | 会话不存在 |

---

## 八、陪玩/专家模块 (Playmate) - 复用订单模块错误码

### 8.1 陪玩基础操作

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| GET /playmate/expert/detail/{expertId} | 陪玩不存在 | 10102 | 陪玩不存在 |
| PUT /playmate/skill/update/{skillId} | 技能不存在 | 10101 | 订单不存在 |
| GET /playmate/match/history/{id} | 匹配历史不存在 | 10101 | 订单不存在 |

### 8.2 专家关注

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/expert/follow/{expertId} | 重复关注 | 10005 | 已经关注过该用户 |
| POST /playmate/expert/follow/{expertId} | 未关注 | 10006 | 未关注该用户 |

### 8.3 评价

| 接口 | 场景 | 错误码 | 错误信息 |
|------|------|--------|----------|
| POST /playmate/review/submit | 陪玩不存在 | 10102 | 陪玩不存在 |
| POST /playmate/review/submit | 重复评价 | - | 已经评价过该陪玩 |

---

## 九、错误码速查表

| 错误码 | 常量名称 | 错误信息 |
|--------|----------|----------|
| 10001 | ErrUserNotFound | 用户不存在 |
| 10002 | ErrInvalidCredentials | 用户名或密码错误 |
| 10003 | ErrUserAlreadyExists | 用户名已存在 |
| 10004 | ErrPhoneAlreadyRegistered | 手机号已被注册 |
| 10005 | ErrAlreadyFollowed | 已经关注过该用户 |
| 10006 | ErrNotFollowed | 未关注该用户 |
| 10007 | ErrFavoriteNotFound | 收藏不存在 |
| 10008 | ErrInsufficientBalance | 余额不足 |
| 10009 | ErrWalletNotFound | 钱包不存在 |
| 10010 | ErrInvalidAmount | 金额格式错误 |
| 10011 | ErrCannotFollowSelf | 不能关注自己 |
| 10101 | ErrOrderNotFound | 订单不存在 |
| 10102 | ErrPlaymateNotFound | 陪玩不存在 |
| 10103 | ErrOrderNot抢able | 订单不可抢 |
| 10104 | ErrAlready抢edOrder | 您已经抢过此订单 |
| 10105 | Err抢单ApplicationNotFound | 抢单申请不存在 |
| 10106 | Err抢单ApplicationNotMatch | 抢单申请不属于该订单 |
| 10107 | ErrOrderStatusNotAllowPay | 订单状态不允许支付 |
| 10108 | ErrPayAmountMismatch | 支付金额与订单金额不符 |
| 10109 | ErrOrderStatusNotAllowConfirm | 订单状态不允许确认服务 |
| 10110 | ErrUnauthorizedOperation | 无权操作此订单 |
| 10111 | ErrOrderStatusNotAllowCancel | 该订单状态无法取消 |
| 10112 | ErrOrderStatusNotAllowAccept | 该订单状态无法接受 |
| 10113 | ErrOrderStatusNotAllowReject | 该订单状态无法拒绝 |
| 10201 | ErrAppealNotFound | 申诉不存在 |
| 10202 | ErrAppealStatusNotUpdatable | 只能更新待处理状态的申诉 |
| 10203 | ErrAppealAlreadyProcessed | 该申诉已处理完成 |
| 10301 | ErrNotificationNotFound | 通知不存在 |
| 10401 | ErrMessageUserNotFound | 用户不存在 |
| 10402 | ErrConversationNotFound | 会话不存在 |

---

## 十、响应格式示例

### 成功响应
```json
{
    "code": 0,
    "data": {},
    "msg": "操作成功"
}
```

### 错误响应
```json
{
    "code": 10001,
    "data": {},
    "msg": "用户不存在"
}
```

---

## 更新日志

- 2026-03-29: 初始版本，创建完整的错误码接口文档
