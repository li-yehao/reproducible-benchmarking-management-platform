package Job

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Job_infoRouter struct {
}

// InitJob_infoRouter 初始化 Job_info 路由信息
func (s *Job_infoRouter) InitJob_infoRouter(Router *gin.RouterGroup) {
	jiRouter := Router.Group("ji").Use(middleware.OperationRecord())
	jiRouterWithoutRecord := Router.Group("ji")
	var jiApi = v1.ApiGroupApp.JobApiGroup.Job_infoApi
	{
		jiRouter.POST("createJob_info", jiApi.CreateJob_info)   // 新建Job_info
		jiRouter.DELETE("deleteJob_info", jiApi.DeleteJob_info) // 删除Job_info
		jiRouter.DELETE("deleteJob_infoByIds", jiApi.DeleteJob_infoByIds) // 批量删除Job_info
		jiRouter.PUT("updateJob_info", jiApi.UpdateJob_info)    // 更新Job_info
	}
	{
		jiRouterWithoutRecord.GET("findJob_info", jiApi.FindJob_info)        // 根据ID获取Job_info
		jiRouterWithoutRecord.GET("getJob_infoList", jiApi.GetJob_infoList)  // 获取Job_info列表
	}
}
