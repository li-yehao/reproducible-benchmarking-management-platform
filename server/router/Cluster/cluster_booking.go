package Cluster

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Cluster_bookingRouter struct {
}

// InitCluster_bookingRouter 初始化 Cluster_booking 路由信息
func (s *Cluster_bookingRouter) InitCluster_bookingRouter(Router *gin.RouterGroup) {
	cbRouter := Router.Group("cb").Use(middleware.OperationRecord())
	cbRouterWithoutRecord := Router.Group("cb")
	var cbApi = v1.ApiGroupApp.ClusterApiGroup.Cluster_bookingApi
	{
		cbRouter.POST("createCluster_booking", cbApi.CreateCluster_booking)   // 新建Cluster_booking
		cbRouter.DELETE("deleteCluster_booking", cbApi.DeleteCluster_booking) // 删除Cluster_booking
		cbRouter.DELETE("deleteCluster_bookingByIds", cbApi.DeleteCluster_bookingByIds) // 批量删除Cluster_booking
		cbRouter.PUT("updateCluster_booking", cbApi.UpdateCluster_booking)    // 更新Cluster_booking
	}
	{
		cbRouterWithoutRecord.GET("findCluster_booking", cbApi.FindCluster_booking)        // 根据ID获取Cluster_booking
		cbRouterWithoutRecord.GET("getCluster_bookingList", cbApi.GetCluster_bookingList)  // 获取Cluster_booking列表
	}
}
