package Cluster

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClusterRouter struct {
}

// InitClusterRouter 初始化 Cluster 路由信息
func (s *ClusterRouter) InitClusterRouter(Router *gin.RouterGroup) {
	csRouter := Router.Group("cs").Use(middleware.OperationRecord())
	csRouterWithoutRecord := Router.Group("cs")
	var csApi = v1.ApiGroupApp.ClusterApiGroup.ClusterApi
	{
		csRouter.POST("createCluster", csApi.CreateCluster)   // 新建Cluster
		csRouter.DELETE("deleteCluster", csApi.DeleteCluster) // 删除Cluster
		csRouter.DELETE("deleteClusterByIds", csApi.DeleteClusterByIds) // 批量删除Cluster
		csRouter.PUT("updateCluster", csApi.UpdateCluster)    // 更新Cluster
	}
	{
		csRouterWithoutRecord.GET("findCluster", csApi.FindCluster)        // 根据ID获取Cluster
		csRouterWithoutRecord.GET("getClusterList", csApi.GetClusterList)  // 获取Cluster列表
	}
}
