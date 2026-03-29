# Playmate 插件错误码文档

## 概述

Playmate 插件使用独立的错误码体系，基础错误码为 `10000`，避免与系统其他模块的错误码冲突。

## 错误码范围

| 模块 | 错误码范围 |
|------|-----------|
| 用户模块 | 10001 - 10100 |
| 订单模块 | 10101 - 10200 |
| 申诉模块 | 10201 - 10300 |
| 通知模块 | 10301 - 10400 |
| 消息模块 | 10401 - 10500 |

## 错误码详细定义

### 用户模块错误码 (ErrUser*)

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

### 订单模块错误码 (ErrOrder*)

| 错误码 | 常量名称 | 错误信息 |
|--------|----------|----------|
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

### 申诉模块错误码 (ErrAppeal*)

| 错误码 | 常量名称 | 错误信息 |
|--------|----------|----------|
| 10201 | ErrAppealNotFound | 申诉不存在 |
| 10202 | ErrAppealStatusNotUpdatable | 只能更新待处理状态的申诉 |
| 10203 | ErrAppealAlreadyProcessed | 该申诉已处理完成 |

### 通知模块错误码 (ErrNotification*)

| 错误码 | 常量名称 | 错误信息 |
|--------|----------|----------|
| 10301 | ErrNotificationNotFound | 通知不存在 |

### 消息模块错误码 (ErrMessage*)

| 错误码 | 常量名称 | 错误信息 |
|--------|----------|----------|
| 10401 | ErrMessageUserNotFound | 用户不存在 |
| 10402 | ErrConversationNotFound | 会话不存在 |

## 使用方式

### Service 层返回错误

```go
import "github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"

// 返回带错误码的错误
return response.NewPlaymateError(response.ErrUserNotFound)

// 返回带自定义消息的错误码
return response.NewPlaymateErrorWithMsg(response.ErrUserNotFound, "自定义错误信息")
```

### API 层处理错误

```go
import "github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"

// 使用 FailWithCode 自动映射错误信息
if err != nil {
    response.FailWithCode(response.ErrUserNotFound, c)
    return
}

// 使用 FailWithError 自动识别 PlaymateError 类型
if err != nil {
    response.FailWithError(err, c)
    return
}
```

## 响应格式

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

## 更新日志

- 2026-03-29: 初始版本，创建独立的错误码体系
