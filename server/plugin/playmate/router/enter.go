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
}

var RouterGroupApp = new(RouterGroup)
