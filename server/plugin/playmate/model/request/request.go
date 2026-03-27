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
	Theme         string          `json:"theme"`
	Language      string          `json:"language"`
}

// CreateOrderRequest 创建订单请求
type CreateOrderRequest struct {
	PlaymateID  uint    `json:"playmateId" binding:"required"`
	SkillID     uint    `json:"skillId" binding:"omitempty"`
	Game        string  `json:"game" binding:"required"`
	Skill       string  `json:"skill" binding:"required"`
	ServiceTime string  `json:"serviceTime" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required,min=1"`
}

// SubmitReviewRequest 提交评价请求
type SubmitReviewRequest struct {
	PlaymateID uint     `json:"playmateId" binding:"required"`
	OrderID    uint     `json:"orderId" binding:"omitempty"`
	Rating     int      `json:"rating" binding:"required,min=1,max=5"`
	Content    string   `json:"content" binding:"required"`
	Images     []string `json:"images"`
	Tags       []string `json:"tags"`
}

// SubmitWithdrawalRequest 提交提现请求
type SubmitWithdrawalRequest struct {
	Amount string `json:"amount" binding:"required"`
	Method string `json:"method" binding:"required"`
}

// SendMessageRequest 发送消息请求
type SendMessageRequest struct {
	Content string `json:"content" binding:"required"`
	Type    string `json:"type" binding:"omitempty,oneof=text image voice system"` // text, image, voice, system
}

// CommentPostRequest 评论帖子请求
type CommentPostRequest struct {
	Content string `json:"content" binding:"required"`
}

// GrabRewardOrderRequest 抢奖励订单请求
type GrabRewardOrderRequest struct {
	OrderID uint `json:"orderId" binding:"required"`
}

// RewardOrderSearch 奖励订单搜索请求
type RewardOrderSearch struct {
	Game          string `json:"game" form:"game"`
	Status        string `json:"status" form:"status"`
	PaymentMethod string `json:"paymentMethod" form:"paymentMethod"`
	Keyword       string `json:"keyword" form:"keyword"`
	Page          int    `json:"page" form:"page"`
	PageSize      int    `json:"pageSize" form:"pageSize"`
}

// CreateRewardOrderRequest 创建奖励订单请求
type CreateRewardOrderRequest struct {
	Game          string  `json:"game" binding:"required"`
	Content       string  `json:"content" binding:"required"`
	Reward        float64 `json:"reward" binding:"required,gt=0"`
	PaymentMethod string  `json:"paymentMethod" binding:"required"`
	Tags          string  `json:"tags"`
}

// UpdateRewardOrderRequest 更新奖励订单请求
type UpdateRewardOrderRequest struct {
	Game          string  `json:"game"`
	Content       string  `json:"content"`
	Reward        float64 `json:"reward"`
	PaymentMethod string  `json:"paymentMethod"`
	Status        string  `json:"status"`
	Tags          string  `json:"tags"`
}

// ActivitySearch 活动搜索请求
type ActivitySearch struct {
	Status    string `json:"status" form:"status"`
	Type      string `json:"type" form:"type"`
	StartTime string `json:"startTime" form:"startTime"`
	EndTime   string `json:"endTime" form:"endTime"`
	Keyword   string `json:"keyword" form:"keyword"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
}

// OrderSearch 订单搜索请求
type OrderSearch struct {
	Status    string  `json:"status" form:"status"`
	Game      string  `json:"game" form:"game"`
	StartTime string  `json:"startTime" form:"startTime"`
	EndTime   string  `json:"endTime" form:"endTime"`
	MinAmount float64 `json:"minAmount" form:"minAmount"`
	MaxAmount float64 `json:"maxAmount" form:"maxAmount"`
	Quantity  int     `json:"quantity" form:"quantity"`
	Keyword   string  `json:"keyword" form:"keyword"`
	Page      int     `json:"page" form:"page"`
	PageSize  int     `json:"pageSize" form:"pageSize"`
}

// CategorySearch 分类搜索请求
type CategorySearch struct {
	Status   string `json:"status" form:"status"`
	Type     string `json:"type" form:"type"`
	ParentID uint   `json:"parentId" form:"parentId"`
	Keyword  string `json:"keyword" form:"keyword"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

// CommunitySearch 社区搜索请求
type CommunitySearch struct {
	Type      string `json:"type" form:"type"`
	Status    string `json:"status" form:"status"`
	UserID    uint   `json:"userId" form:"userId"`
	Game      string `json:"game" form:"game"`
	Keyword   string `json:"keyword" form:"keyword"`
	StartTime string `json:"startTime" form:"startTime"`
	EndTime   string `json:"endTime" form:"endTime"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
}

// GameSearch 游戏搜索请求
type GameSearch struct {
	CategoryIDs []uint `json:"categoryIds" form:"categoryIds"`
	Status      string `json:"status" form:"status"`
	Platform    string `json:"platform" form:"platform"`
	Keyword     string `json:"keyword" form:"keyword"`
	Page        int    `json:"page" form:"page"`
	PageSize    int    `json:"pageSize" form:"pageSize"`
}

// MessageSearch 消息搜索请求
type MessageSearch struct {
	Type       string `json:"type" form:"type"`
	Status     string `json:"status" form:"status"`
	SenderID   uint   `json:"senderId" form:"senderId"`
	ReceiverID uint   `json:"receiverId" form:"receiverId"`
	StartTime  string `json:"startTime" form:"startTime"`
	EndTime    string `json:"endTime" form:"endTime"`
	Keyword    string `json:"keyword" form:"keyword"`
	Page       int    `json:"page" form:"page"`
	PageSize   int    `json:"pageSize" form:"pageSize"`
}

// NotificationSearch 通知搜索请求
type NotificationSearch struct {
	Type      string `json:"type" form:"type"`
	Status    string `json:"status" form:"status"`
	StartTime string `json:"startTime" form:"startTime"`
	EndTime   string `json:"endTime" form:"endTime"`
	Keyword   string `json:"keyword" form:"keyword"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
}

// ReviewSearch 评价搜索请求
type ReviewSearch struct {
	PlaymateID uint   `json:"playmateId" form:"playmateId"`
	MinRating  int    `json:"minRating" form:"minRating"`
	MaxRating  int    `json:"maxRating" form:"maxRating"`
	Game       string `json:"game" form:"game"`
	Keyword    string `json:"keyword" form:"keyword"`
	StartTime  string `json:"startTime" form:"startTime"`
	EndTime    string `json:"endTime" form:"endTime"`
	Page       int    `json:"page" form:"page"`
	PageSize   int    `json:"pageSize" form:"pageSize"`
}

// WithdrawalSearch 提现搜索请求
type WithdrawalSearch struct {
	Status    string  `json:"status" form:"status"`
	Method    string  `json:"method" form:"method"`
	MinAmount float64 `json:"minAmount" form:"minAmount"`
	MaxAmount float64 `json:"maxAmount" form:"maxAmount"`
	StartTime string  `json:"startTime" form:"startTime"`
	EndTime   string  `json:"endTime" form:"endTime"`
	Page      int     `json:"page" form:"page"`
	PageSize  int     `json:"pageSize" form:"pageSize"`
}

// GameCategorySearch 游戏分类搜索请求
type GameCategorySearch struct {
	Status   string `json:"status" form:"status"`
	ParentID uint   `json:"parentId" form:"parentId"`
	Keyword  string `json:"keyword" form:"keyword"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

// AppealSearch 申诉搜索请求
type AppealSearch struct {
	Type      string `json:"type" form:"type"`
	Status    string `json:"status" form:"status"`
	Priority  string `json:"priority" form:"priority"`
	UserID    uint   `json:"userId" form:"userId"`
	OrderID   uint   `json:"orderId" form:"orderId"`
	Keyword   string `json:"keyword" form:"keyword"`
	StartTime string `json:"startTime" form:"startTime"`
	EndTime   string `json:"endTime" form:"endTime"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
}

// CreateAppealRequest 创建申诉请求
type CreateAppealRequest struct {
	OrderID     uint   `json:"orderId" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Images      string `json:"images"`
	ContactInfo string `json:"contactInfo"`
	Priority    string `json:"priority"`
}

// UpdateAppealRequest 更新申诉请求
type UpdateAppealRequest struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Images      string `json:"images"`
	Status      string `json:"status"`
	Response    string `json:"response"`
	ContactInfo string `json:"contactInfo"`
	Priority    string `json:"priority"`
}

// HandleAppealRequest 处理申诉请求
type HandleAppealRequest struct {
	Status   string `json:"status" binding:"required"` // resolved, rejected
	Response string `json:"response" binding:"required"`
}
