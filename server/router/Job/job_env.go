package Job

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Job_envRouter struct {
}

// InitJob_envRouter 初始化 Job_env 路由信息
func (s *Job_envRouter) InitJob_envRouter(Router *gin.RouterGroup) {
	jeRouter := Router.Group("je").Use(middleware.OperationRecord())
	jeRouterWithoutRecord := Router.Group("je")
	var jeApi = v1.ApiGroupApp.JobApiGroup.Job_envApi
	{
		jeRouter.POST("createJob_env", jeApi.CreateJob_env)   // 新建Job_env
		jeRouter.DELETE("deleteJob_env", jeApi.DeleteJob_env) // 删除Job_env
		jeRouter.DELETE("deleteJob_envByIds", jeApi.DeleteJob_envByIds) // 批量删除Job_env
		jeRouter.PUT("updateJob_env", jeApi.UpdateJob_env)    // 更新Job_env
	}
	{
		jeRouterWithoutRecord.GET("findJob_env", jeApi.FindJob_env)        // 根据ID获取Job_env
		jeRouterWithoutRecord.GET("getJob_envList", jeApi.GetJob_envList)  // 获取Job_env列表
	}
}
