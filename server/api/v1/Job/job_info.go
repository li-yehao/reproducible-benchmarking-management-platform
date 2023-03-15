package Job

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Job"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    JobReq "github.com/flipped-aurora/gin-vue-admin/server/model/Job/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type Job_infoApi struct {
}

var jiService = service.ServiceGroupApp.JobServiceGroup.Job_infoService


// CreateJob_info 创建Job_info
// @Tags Job_info
// @Summary 创建Job_info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Job.Job_info true "创建Job_info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ji/createJob_info [post]
func (jiApi *Job_infoApi) CreateJob_info(c *gin.Context) {
	var ji Job.Job_info
	err := c.ShouldBindJSON(&ji)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    ji.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Cluster_name":{utils.NotEmpty()},
    }
	if err := utils.Verify(ji, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := jiService.CreateJob_info(ji); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteJob_info 删除Job_info
// @Tags Job_info
// @Summary 删除Job_info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Job.Job_info true "删除Job_info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ji/deleteJob_info [delete]
func (jiApi *Job_infoApi) DeleteJob_info(c *gin.Context) {
	var ji Job.Job_info
	err := c.ShouldBindJSON(&ji)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    ji.DeletedBy = utils.GetUserID(c)
	if err := jiService.DeleteJob_info(ji); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteJob_infoByIds 批量删除Job_info
// @Tags Job_info
// @Summary 批量删除Job_info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Job_info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ji/deleteJob_infoByIds [delete]
func (jiApi *Job_infoApi) DeleteJob_infoByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := jiService.DeleteJob_infoByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateJob_info 更新Job_info
// @Tags Job_info
// @Summary 更新Job_info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Job.Job_info true "更新Job_info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ji/updateJob_info [put]
func (jiApi *Job_infoApi) UpdateJob_info(c *gin.Context) {
	var ji Job.Job_info
	err := c.ShouldBindJSON(&ji)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    ji.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Cluster_name":{utils.NotEmpty()},
      }
    if err := utils.Verify(ji, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := jiService.UpdateJob_info(ji); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindJob_info 用id查询Job_info
// @Tags Job_info
// @Summary 用id查询Job_info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Job.Job_info true "用id查询Job_info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ji/findJob_info [get]
func (jiApi *Job_infoApi) FindJob_info(c *gin.Context) {
	var ji Job.Job_info
	err := c.ShouldBindQuery(&ji)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reji, err := jiService.GetJob_info(ji.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reji": reji}, c)
	}
}

// GetJob_infoList 分页获取Job_info列表
// @Tags Job_info
// @Summary 分页获取Job_info列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query JobReq.Job_infoSearch true "分页获取Job_info列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ji/getJob_infoList [get]
func (jiApi *Job_infoApi) GetJob_infoList(c *gin.Context) {
	var pageInfo JobReq.Job_infoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := jiService.GetJob_infoInfoList(pageInfo); err != nil {
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
