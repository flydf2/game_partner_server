package request

// PlaymateSearch 陪玩搜索请求
type PlaymateSearch struct {
	Game       string `json:"game" form:"game"`
	Online     *bool  `json:"online" form:"online"`
	PriceRange string `json:"priceRange" form:"priceRange"`
	Rank       string `json:"rank" form:"rank"`
	Gender     string `json:"gender" form:"gender"`
	Keyword    string `json:"keyword" form:"keyword"`
	SortBy     string `json:"sortBy" form:"sortBy"`
	Page       int    `json:"page" form:"page"`
	PageSize   int    `json:"pageSize" form:"pageSize"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

// UpdateProfileRequest 更新个人资料请求
type UpdateProfileRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
}

// UpdateSettingsRequest 更新设置请求
type UpdateSettingsRequest struct {
	Notifications map[string]bool `json:"notifications"`
	Privacy       map[string]bool `json:"privacy"`
	Theme         string         `json:"theme"`
	Language      string         `json:"language"`
}

// CreateOrderRequest 创建订单请求
type CreateOrderRequest struct {
	PlaymateID  uint    `json:"playmateId" binding:"required"`
	Game        string  `json:"game" binding:"required"`
	Skill       string  `json:"skill" binding:"required"`
	ServiceTime string  `json:"serviceTime" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
}

// SubmitReviewRequest 提交评价请求
type SubmitReviewRequest struct {
	PlaymateID uint     `json:"playmateId" binding:"required"`
	Rating     int      `json:"rating" binding:"required,min=1,max=5"`
	Content    string   `json:"content" binding:"required"`
	Images     []string `json:"images"`
	Tags       []string `json:"tags"`
}

// SubmitWithdrawalRequest 提交提现请求
type SubmitWithdrawalRequest struct {
	Amount float64 `json:"amount" binding:"required,min=10,max=50000"`
	Method string  `json:"method" binding:"required"`
}

// SendMessageRequest 发送消息请求
type SendMessageRequest struct {
	Content string `json:"content" binding:"required"`
}

// CommentPostRequest 评论帖子请求
type CommentPostRequest struct {
	Content string `json:"content" binding:"required"`
}

// GrabRewardOrderRequest 抢奖励订单请求
type GrabRewardOrderRequest struct {
	OrderID uint `json:"orderId" binding:"required"`
}