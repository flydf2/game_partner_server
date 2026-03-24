package service

// ServiceGroup 服务组
type ServiceGroup struct {
	PlaymateService
	UserService
	OrderService
	NotificationService
	MessageService
	GameService
	ActivityService
	ReviewService
	WithdrawalService
	CategoryService
	RecommendationService
	CommunityService
	GameCategoryService
	RewardOrderService
	UploadService
}

var ServiceGroupApp = new(ServiceGroup)