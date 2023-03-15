package Cluster

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Cluster"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ClusterReq "github.com/flipped-aurora/gin-vue-admin/server/model/Cluster/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type Cluster_bookingApi struct {
}

var cbService = service.ServiceGroupApp.ClusterServiceGroup.Cluster_bookingService


// CreateCluster_booking 创建Cluster_booking
// @Tags Cluster_booking
// @Summary 创建Cluster_booking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Cluster.Cluster_booking true "创建Cluster_booking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cb/createCluster_booking [post]
func (cbApi *Cluster_bookingApi) CreateCluster_booking(c *gin.Context) {
	var cb Cluster.Cluster_booking
	err := c.ShouldBindJSON(&cb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cb.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Cluster_name":{utils.NotEmpty()},
    }
	if err := utils.Verify(cb, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := cbService.CreateCluster_booking(cb); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCluster_booking 删除Cluster_booking
// @Tags Cluster_booking
// @Summary 删除Cluster_booking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Cluster.Cluster_booking true "删除Cluster_booking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cb/deleteCluster_booking [delete]
func (cbApi *Cluster_bookingApi) DeleteCluster_booking(c *gin.Context) {
	var cb Cluster.Cluster_booking
	err := c.ShouldBindJSON(&cb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cb.DeletedBy = utils.GetUserID(c)
	if err := cbService.DeleteCluster_booking(cb); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCluster_bookingByIds 批量删除Cluster_booking
// @Tags Cluster_booking
// @Summary 批量删除Cluster_booking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cluster_booking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cb/deleteCluster_bookingByIds [delete]
func (cbApi *Cluster_bookingApi) DeleteCluster_bookingByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := cbService.DeleteCluster_bookingByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCluster_booking 更新Cluster_booking
// @Tags Cluster_booking
// @Summary 更新Cluster_booking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Cluster.Cluster_booking true "更新Cluster_booking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cb/updateCluster_booking [put]
func (cbApi *Cluster_bookingApi) UpdateCluster_booking(c *gin.Context) {
	var cb Cluster.Cluster_booking
	err := c.ShouldBindJSON(&cb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cb.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Cluster_name":{utils.NotEmpty()},
      }
    if err := utils.Verify(cb, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := cbService.UpdateCluster_booking(cb); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCluster_booking 用id查询Cluster_booking
// @Tags Cluster_booking
// @Summary 用id查询Cluster_booking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Cluster.Cluster_booking true "用id查询Cluster_booking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cb/findCluster_booking [get]
func (cbApi *Cluster_bookingApi) FindCluster_booking(c *gin.Context) {
	var cb Cluster.Cluster_booking
	err := c.ShouldBindQuery(&cb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recb, err := cbService.GetCluster_booking(cb.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recb": recb}, c)
	}
}

// GetCluster_bookingList 分页获取Cluster_booking列表
// @Tags Cluster_booking
// @Summary 分页获取Cluster_booking列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ClusterReq.Cluster_bookingSearch true "分页获取Cluster_booking列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cb/getCluster_bookingList [get]
func (cbApi *Cluster_bookingApi) GetCluster_bookingList(c *gin.Context) {
	var pageInfo ClusterReq.Cluster_bookingSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cbService.GetCluster_bookingInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
