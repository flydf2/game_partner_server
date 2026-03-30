package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/gin-gonic/gin"
)

type TournamentRouter struct{}

// InitTournamentRouter 初始化赛事路由
func (r *TournamentRouter) InitTournamentRouter(router *gin.RouterGroup) {
	tournamentRouter := router.Group("/tournaments")
	{
		// 公开接口
		tournamentRouter.GET("", api.ApiGroupApp.TournamentApi.GetTournamentList)
		tournamentRouter.GET("/:id", api.ApiGroupApp.TournamentApi.GetTournamentDetail)
		tournamentRouter.GET("/:id/teams", api.ApiGroupApp.TournamentApi.GetTournamentTeams)
		tournamentRouter.GET("/:id/matches", api.ApiGroupApp.TournamentApi.GetTournamentMatches)

		// 需要登录的接口
		tournamentRouter.POST("/join", api.ApiGroupApp.TournamentApi.JoinTournament)

		// 管理接口（需要管理员权限）
		tournamentRouter.POST("", api.ApiGroupApp.TournamentApi.CreateTournament)
		tournamentRouter.PUT("/:id", api.ApiGroupApp.TournamentApi.UpdateTournament)
		tournamentRouter.DELETE("/:id", api.ApiGroupApp.TournamentApi.DeleteTournament)

		// 队伍管理接口
		tournamentRouter.POST("/teams", api.ApiGroupApp.TournamentApi.CreateTournamentTeam)
		tournamentRouter.PUT("/teams/:id", api.ApiGroupApp.TournamentApi.UpdateTournamentTeam)
		tournamentRouter.DELETE("/teams/:id", api.ApiGroupApp.TournamentApi.DeleteTournamentTeam)

		// 比赛管理接口
		tournamentRouter.POST("/matches", api.ApiGroupApp.TournamentApi.CreateTournamentMatch)
		tournamentRouter.PUT("/matches/:id", api.ApiGroupApp.TournamentApi.UpdateTournamentMatch)
		tournamentRouter.DELETE("/matches/:id", api.ApiGroupApp.TournamentApi.DeleteTournamentMatch)
	}
}
