package router

type RouterGroup struct {
	PlaymateRouter
	UserRouter
	OrderRouter
	NotificationRouter
	MessageRouter
	GameRouter
	ActivityRouter
	ReviewRouter
	WithdrawalRouter
	CommunityRouter
	CategoryRouter
	GameCategoryRouter
	RewardOrderRouter
	AppealRouter
	UploadRouter
	TestToolRouter
}

var RouterGroupApp = new(RouterGroup)
