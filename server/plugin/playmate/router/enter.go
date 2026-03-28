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
}

var RouterGroupApp = new(RouterGroup)
