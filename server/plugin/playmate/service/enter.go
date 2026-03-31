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
	AppealService
	ExpertOrderSettingService
	TournamentService
	LeaderboardService
	StatsService
}

var ServiceGroupApp = new(ServiceGroup)
