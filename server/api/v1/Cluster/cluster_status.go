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

type ClusterApi struct {
}

var csService = service.ServiceGroupApp.ClusterServiceGroup.ClusterService


// CreateCluster 创建Cluster
// @Tags Cluster
// @Summary 创建Cluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Cluster.Cluster true "创建Cluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cs/createCluster [post]
func (csApi *ClusterApi) CreateCluster(c *gin.Context) {
	var cs Cluster.Cluster
	err := c.ShouldBindJSON(&cs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Cluster_name":{utils.NotEmpty()},
    }
	if err := utils.Verify(cs, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := csService.CreateCluster(cs); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCluster 删除Cluster
// @Tags Cluster
// @Summary 删除Cluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Cluster.Cluster true "删除Cluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cs/deleteCluster [delete]
func (csApi *ClusterApi) DeleteCluster(c *gin.Context) {
	var cs Cluster.Cluster
	err := c.ShouldBindJSON(&cs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csService.DeleteCluster(cs); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteClusterByIds 批量删除Cluster
// @Tags Cluster
// @Summary 批量删除Cluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cs/deleteClusterByIds [delete]
func (csApi *ClusterApi) DeleteClusterByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := csService.DeleteClusterByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCluster 更新Cluster
// @Tags Cluster
// @Summary 更新Cluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Cluster.Cluster true "更新Cluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cs/updateCluster [put]
func (csApi *ClusterApi) UpdateCluster(c *gin.Context) {
	var cs Cluster.Cluster
	err := c.ShouldBindJSON(&cs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Cluster_name":{utils.NotEmpty()},
      }
    if err := utils.Verify(cs, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := csService.UpdateCluster(cs); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCluster 用id查询Cluster
// @Tags Cluster
// @Summary 用id查询Cluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Cluster.Cluster true "用id查询Cluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cs/findCluster [get]
func (csApi *ClusterApi) FindCluster(c *gin.Context) {
	var cs Cluster.Cluster
	err := c.ShouldBindQuery(&cs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recs, err := csService.GetCluster(cs.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recs": recs}, c)
	}
}

// GetClusterList 分页获取Cluster列表
// @Tags Cluster
// @Summary 分页获取Cluster列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ClusterReq.ClusterSearch true "分页获取Cluster列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cs/getClusterList [get]
func (csApi *ClusterApi) GetClusterList(c *gin.Context) {
	var pageInfo ClusterReq.ClusterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := csService.GetClusterInfoList(pageInfo); err != nil {
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
