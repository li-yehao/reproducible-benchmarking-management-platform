package Job

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Job_listRouter struct {
}

// InitJob_listRouter 初始化 Job_list 路由信息
func (s *Job_listRouter) InitJob_listRouter(Router *gin.RouterGroup) {
	jlRouter := Router.Group("jl").Use(middleware.OperationRecord())
	jlRouterWithoutRecord := Router.Group("jl")
	var jlApi = v1.ApiGroupApp.JobApiGroup.Job_listApi
	{
		jlRouter.POST("createJob_list", jlApi.CreateJob_list)   // 新建Job_list
		jlRouter.DELETE("deleteJob_list", jlApi.DeleteJob_list) // 删除Job_list
		jlRouter.DELETE("deleteJob_listByIds", jlApi.DeleteJob_listByIds) // 批量删除Job_list
		jlRouter.PUT("updateJob_list", jlApi.UpdateJob_list)    // 更新Job_list
	}
	{
		jlRouterWithoutRecord.GET("findJob_list", jlApi.FindJob_list)        // 根据ID获取Job_list
		jlRouterWithoutRecord.GET("getJob_listList", jlApi.GetJob_listList)  // 获取Job_list列表
		jlRouterWithoutRecord.GET("executeJob_list", jlApi.ExecuteJob_list)  // 根据ID执行Job_list
		jlRouterWithoutRecord.GET("cancelJob_list", jlApi.CancelJob_list)    // 根据ID执行Job_list
	}
}
