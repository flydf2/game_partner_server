package api

// ApiGroup API组
type ApiGroup struct {
	PlaymateApi
	UserApi
	OrderApi
	NotificationApi
	MessageApi
	GameApi
	ActivityApi
	ReviewApi
	WithdrawalApi
	CategoryApi
	CommunityApi
	GameCategoryApi
}

var ApiGroupApp = new(ApiGroup)