package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 0
)

const (
	PlaymateBaseCode = 10000
)

const (
	ErrUserNotFound = iota + 1 + PlaymateBaseCode
	ErrInvalidCredentials
	ErrUserAlreadyExists
	ErrPhoneAlreadyRegistered
	ErrAlreadyFollowed
	ErrNotFollowed
	ErrFavoriteNotFound
	ErrInsufficientBalance
	ErrWalletNotFound
	ErrInvalidAmount
	ErrCannotFollowSelf
	ErrWithdrawalNotFound
	ErrInvalidStatus
)

const (
	ErrOrderNotFound = iota + 101 + PlaymateBaseCode
	ErrPlaymateNotFound
	ErrOrderNotGrabble
	ErrAlreadyGrabbedOrder
	ErrGrabOrderApplicationNotFound
	ErrGrabOrderApplicationNotMatch
	ErrOrderStatusNotAllowPay
	ErrPayAmountMismatch
	ErrOrderStatusNotAllowConfirm
	ErrUnauthorizedOperation
	ErrOrderStatusNotAllowCancel
	ErrOrderStatusNotAllowAccept
	ErrOrderStatusNotAllowReject
	ErrGrabOrderStatusNotAllowWithdraw
)

const (
	ErrAppealNotFound = iota + 201 + PlaymateBaseCode
	ErrAppealStatusNotUpdatable
	ErrAppealAlreadyProcessed
)

const (
	ErrNotificationNotFound = iota + 301 + PlaymateBaseCode
)

const (
	ErrMessageUserNotFound = iota + 401 + PlaymateBaseCode
	ErrConversationNotFound
)

var playmateErrorMessages = map[int]string{
	SUCCESS:                   "成功",
	ErrUserNotFound:           "用户不存在",
	ErrInvalidCredentials:     "用户名或密码错误",
	ErrUserAlreadyExists:      "用户名已存在",
	ErrPhoneAlreadyRegistered: "手机号已被注册",
	ErrAlreadyFollowed:        "已经关注过该用户",
	ErrNotFollowed:            "未关注该用户",
	ErrFavoriteNotFound:       "收藏不存在",
	ErrInsufficientBalance:    "余额不足",
	ErrWalletNotFound:         "钱包不存在",
	ErrInvalidAmount:          "金额格式错误",
	ErrCannotFollowSelf:       "不能关注自己",
	ErrWithdrawalNotFound:     "提现记录不存在",
	ErrInvalidStatus:          "无效的状态",

	ErrOrderNotFound:              "订单不存在",
	ErrPlaymateNotFound:           "陪玩不存在",
	ErrOrderNotGrabble:            "订单不可抢",
	ErrAlreadyGrabbedOrder:        "您已经抢过此订单",
	ErrGrabOrderApplicationNotFound: "抢单申请不存在",
	ErrGrabOrderApplicationNotMatch: "抢单申请不属于该订单",
	ErrOrderStatusNotAllowPay:     "订单状态不允许支付",
	ErrPayAmountMismatch:          "支付金额与订单金额不符",
	ErrOrderStatusNotAllowConfirm: "订单状态不允许确认服务",
	ErrUnauthorizedOperation:      "无权操作此订单",
	ErrOrderStatusNotAllowCancel:  "该订单状态无法取消",
	ErrOrderStatusNotAllowAccept:  "该订单状态无法接受",
	ErrOrderStatusNotAllowReject:  "该订单状态无法拒绝",
	ErrGrabOrderStatusNotAllowWithdraw: "该抢单状态无法撤回",
	ErrAppealNotFound:             "申诉不存在",
	ErrAppealStatusNotUpdatable:   "只能更新待处理状态的申诉",
	ErrAppealAlreadyProcessed:     "该申诉已处理完成",
	ErrNotificationNotFound:       "通知不存在",
	ErrMessageUserNotFound:        "用户不存在",
	ErrConversationNotFound:       "会话不存在",
}

type PlaymateError struct {
	Code int
	Msg  string
}

func (e *PlaymateError) Error() string {
	return e.Msg
}

func NewPlaymateError(code int) *PlaymateError {
	return &PlaymateError{
		Code: code,
		Msg:  playmateErrorMessages[code],
	}
}

func NewPlaymateErrorWithMsg(code int, msg string) *PlaymateError {
	return &PlaymateError{
		Code: code,
		Msg:  msg,
	}
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(PlaymateBaseCode, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(PlaymateBaseCode, map[string]interface{}{}, message, c)
}

func FailWithCode(code int, c *gin.Context) {
	msg := playmateErrorMessages[code]
	if msg == "" {
		msg = "未知错误"
	}
	Result(code, map[string]interface{}{}, msg, c)
}

func FailWithCodeAndMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(PlaymateBaseCode, data, message, c)
}

func FailWithError(err error, c *gin.Context) {
	if pe, ok := err.(*PlaymateError); ok {
		FailWithCode(pe.Code, c)
		return
	}
	FailWithMessage(err.Error(), c)
}

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		PlaymateBaseCode,
		nil,
		message,
	})
}

func GetErrorMessage(code int) string {
	if msg, ok := playmateErrorMessages[code]; ok {
		return msg
	}
	return fmt.Sprintf("未知错误码: %d", code)
}

func IsPlaymateErrorCode(code int) bool {
	_, ok := playmateErrorMessages[code]
	return ok
}
