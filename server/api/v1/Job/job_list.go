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

type Job_listApi struct {
}

var jlService = service.ServiceGroupApp.JobServiceGroup.Job_listService


// CreateJob_list 创建Job_list
// @Tags Job_list
// @Summary 创建Job_list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Job.Job_list true "创建Job_list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jl/createJob_list [post]
func (jlApi *Job_listApi) CreateJob_list(c *gin.Context) {
	var jl Job.Job_list
	err := c.ShouldBindJSON(&jl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    jl.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Cluster_name":{utils.NotEmpty()},
    }
	if err := utils.Verify(jl, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := jlService.CreateJob_list(jl); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteJob_list 删除Job_list
// @Tags Job_list
// @Summary 删除Job_list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Job.Job_list true "删除Job_list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jl/deleteJob_list [delete]
func (jlApi *Job_listApi) DeleteJob_list(c *gin.Context) {
	var jl Job.Job_list
	err := c.ShouldBindJSON(&jl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    jl.DeletedBy = utils.GetUserID(c)
	if err := jlService.DeleteJob_list(jl); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteJob_listByIds 批量删除Job_list
// @Tags Job_list
// @Summary 批量删除Job_list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Job_list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /jl/deleteJob_listByIds [delete]
func (jlApi *Job_listApi) DeleteJob_listByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := jlService.DeleteJob_listByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateJob_list 更新Job_list
// @Tags Job_list
// @Summary 更新Job_list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Job.Job_list true "更新Job_list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jl/updateJob_list [put]
func (jlApi *Job_listApi) UpdateJob_list(c *gin.Context) {
	var jl Job.Job_list
	err := c.ShouldBindJSON(&jl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    jl.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Cluster_name":{utils.NotEmpty()},
	}
    if err := utils.Verify(jl, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
	}
	if err := jlService.UpdateJob_list(jl); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindJob_list 用id查询Job_list
// @Tags Job_list
// @Summary 用id查询Job_list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Job.Job_list true "用id查询Job_list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jl/findJob_list [get]
func (jlApi *Job_listApi) FindJob_list(c *gin.Context) {
	var jl Job.Job_list
	err := c.ShouldBindQuery(&jl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rejl, err := jlService.GetJob_list(jl.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rejl": rejl}, c)
	}
}

// GetJob_listList 分页获取Job_list列表
// @Tags Job_list
// @Summary 分页获取Job_list列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query JobReq.Job_listSearch true "分页获取Job_list列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jl/getJob_listList [get]
func (jlApi *Job_listApi) GetJob_listList(c *gin.Context) {
	var pageInfo JobReq.Job_listSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := jlService.GetJob_listInfoList(pageInfo); err != nil {
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

// GetJob_listList 执行Job_list
// @Tags Job_list
// @Summary 执行Job_list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query JobReq.Job_listSearch true "执行Job_list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jl/executeJob_list [get]
func (jlApi *Job_listApi) ExecuteJob_list(c *gin.Context) {
	var jl Job.Job_list
	err := c.ShouldBindQuery(&jl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := jlService.ExecuteJob_list(jl); err != nil {
        global.GVA_LOG.Error("执行失败!", zap.Error(err))
		response.FailWithMessage("执行失败", c)
	} else {
		response.OkWithMessage("执行成功", c)
	}
	
}

// GetJob_listList 取消Job_list
// @Tags Job_list
// @Summary 取消Job_list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query JobReq.Job_listSearch true "取消Job_list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jl/cancelJob_list [get]
func (jlApi *Job_listApi) CancelJob_list(c *gin.Context) {
	var jl Job.Job_list
	err := c.ShouldBindQuery(&jl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := jlService.CancelJob_list(jl); err != nil {
        global.GVA_LOG.Error("取消失败!", zap.Error(err))
		response.FailWithMessage("取消失败", c)
	} else {
		response.OkWithMessage("取消成功", c)
	}
}