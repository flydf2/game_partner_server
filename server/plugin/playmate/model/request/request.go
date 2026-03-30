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
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
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
	Amount      string `json:"amount" binding:"required"`
	Method      string `json:"method" binding:"required"` // alipay, wechat, bank
	AccountInfo string `json:"accountInfo" binding:"required"` // 账户信息，如支付宝账号、银行卡号等
	Description string `json:"description" binding:"omitempty"`
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

// AddSkillRequest 添加技能请求
type AddSkillRequest struct {
	Game        string  `json:"game" binding:"required"`
	Skill       string  `json:"skill" binding:"required"`
	Level       string  `json:"level" binding:"required"` // beginner, intermediate, advanced, expert
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description"`
}

// UpdateSkillRequest 更新技能请求
type UpdateSkillRequest struct {
	Game        string  `json:"game"`
	Skill       string  `json:"skill"`
	Level       string  `json:"level"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// CreatePostRequest 创建帖子请求
type CreatePostRequest struct {
	Content string   `json:"content" binding:"required"`
	Images  []string `json:"images"`
	Game    string   `json:"game"`
}

// RechargeRequest 充值请求
type RechargeRequest struct {
	Amount    float64 `json:"amount" binding:"required,gt=0"`
	Method    string  `json:"method" binding:"required"` // alipay, wechat, balance
	OrderID   string  `json:"orderId" binding:"omitempty"`
	ReturnURL string  `json:"returnUrl" binding:"omitempty"`
	NotifyURL string  `json:"notifyUrl" binding:"omitempty"`
}

// SendSmsCodeRequest 发送短信验证码请求
type SendSmsCodeRequest struct {
	Phone string `json:"phone" binding:"required"`
}

// SkillSearch 技能搜索请求
type SkillSearch struct {
	Game     string `json:"game" form:"game"`
	Level    string `json:"level" form:"level"`
	Keyword  string `json:"keyword" form:"keyword"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

// LeaderboardSearch 排行榜搜索请求
type LeaderboardSearch struct {
	Game     string `json:"game" form:"game"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

// CreateBidRequest 创建投标请求
type CreateBidRequest struct {
	PostID  uint   `json:"postId" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// CreatePlaymateRequest 创建陪玩请求
type CreatePlaymateRequest struct {
	UserID      uint    `json:"userId" binding:"required"`
	Nickname    string  `json:"nickname" binding:"required"`
	Avatar      string  `json:"avatar"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Tags        string  `json:"tags"`
	Game        string  `json:"game" binding:"required"`
	Rank        string  `json:"rank"`
	Gender      string  `json:"gender"`
	Description string  `json:"description"`
	Level       int     `json:"level"`
	Title       string  `json:"title"`
}

// UpdatePlaymateRequest 更新陪玩请求
type UpdatePlaymateRequest struct {
	Nickname    string  `json:"nickname"`
	Avatar      string  `json:"avatar"`
	Price       float64 `json:"price" binding:"omitempty,gt=0"`
	Tags        string  `json:"tags"`
	IsOnline    bool    `json:"isOnline"`
	Game        string  `json:"game"`
	Rank        string  `json:"rank"`
	Gender      string  `json:"gender"`
	Description string  `json:"description"`
	Level       int     `json:"level"`
	Title       string  `json:"title"`
}

// ShareOrderRequest 分享订单请求
type ShareOrderRequest struct {
	Platform string `json:"platform" binding:"required"`
}
