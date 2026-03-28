package request

// RewardOrderSearch 悬赏订单搜索请求
type RewardOrderSearch struct {
	Page           int    `json:"page" form:"page"`
	PageSize       int    `json:"pageSize" form:"pageSize"`
	Status         string `json:"status" form:"status"`
	Game           string `json:"game" form:"game"`
	PaymentMethod  string `json:"paymentMethod" form:"paymentMethod"`
	Keyword        string `json:"keyword" form:"keyword"`
}

// CreateRewardOrderRequest 创建悬赏订单请求
type CreateRewardOrderRequest struct {
	Game          string   `json:"game" binding:"required"`
	Content       string   `json:"content" binding:"required"`
	Reward        float64  `json:"reward" binding:"required,gt=0"`
	PaymentMethod string   `json:"paymentMethod" binding:"required,oneof=prepay postpay"`
	Requirements  []string `json:"requirements" binding:"required"`
	Tags          []string `json:"tags" binding:"required"`
}

// UpdateRewardOrderRequest 更新悬赏订单请求
type UpdateRewardOrderRequest struct {
	Game          string   `json:"game"`
	Content       string   `json:"content"`
	Reward        float64  `json:"reward" binding:"omitempty,gt=0"`
	PaymentMethod string   `json:"paymentMethod" binding:"omitempty,oneof=prepay postpay"`
	Status        string   `json:"status" binding:"omitempty,oneof=available ongoing completed draft cancelled expired"`
	Requirements  []string `json:"requirements"`
	Tags          []string `json:"tags"`
}

// GrabRewardOrderRequest 抢单请求
type GrabRewardOrderRequest struct {
	Recommendation string `json:"recommendation"`
	VoiceUrl       string `json:"voiceUrl"`
	RecordUrl      string `json:"recordUrl"`
}

// SelectApplicantRequest 选择抢单者请求
type SelectApplicantRequest struct {
	ApplicantID uint `json:"applicantId" binding:"required"`
}

// PayRewardOrderRequest 支付订单请求
type PayRewardOrderRequest struct {
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	PaymentMethod string  `json:"paymentMethod" binding:"required"`
	TransactionID string  `json:"transactionId" binding:"required"`
}

// ConfirmServiceRequest 确认服务请求
type ConfirmServiceRequest struct {
	Rating  int      `json:"rating" binding:"required,min=1,max=5"`
	Review  string   `json:"review"`
	Images  []string `json:"images"`
}
