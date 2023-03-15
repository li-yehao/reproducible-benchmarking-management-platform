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

type Job_envApi struct {
}

var jeService = service.ServiceGroupApp.JobServiceGroup.Job_envService


// CreateJob_env 创建Job_env
// @Tags Job_env
// @Summary 创建Job_env
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Job.Job_env true "创建Job_env"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /je/createJob_env [post]
func (jeApi *Job_envApi) CreateJob_env(c *gin.Context) {
	var je Job.Job_env
	err := c.ShouldBindJSON(&je)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    je.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Job_id":{utils.NotEmpty()},
    }
	if err := utils.Verify(je, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := jeService.CreateJob_env(je); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteJob_env 删除Job_env
// @Tags Job_env
// @Summary 删除Job_env
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Job.Job_env true "删除Job_env"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /je/deleteJob_env [delete]
func (jeApi *Job_envApi) DeleteJob_env(c *gin.Context) {
	var je Job.Job_env
	err := c.ShouldBindJSON(&je)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    je.DeletedBy = utils.GetUserID(c)
	if err := jeService.DeleteJob_env(je); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteJob_envByIds 批量删除Job_env
// @Tags Job_env
// @Summary 批量删除Job_env
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Job_env"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /je/deleteJob_envByIds [delete]
func (jeApi *Job_envApi) DeleteJob_envByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := jeService.DeleteJob_envByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateJob_env 更新Job_env
// @Tags Job_env
// @Summary 更新Job_env
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Job.Job_env true "更新Job_env"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /je/updateJob_env [put]
func (jeApi *Job_envApi) UpdateJob_env(c *gin.Context) {
	var je Job.Job_env
	err := c.ShouldBindJSON(&je)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    je.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Job_id":{utils.NotEmpty()},
      }
    if err := utils.Verify(je, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := jeService.UpdateJob_env(je); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindJob_env 用id查询Job_env
// @Tags Job_env
// @Summary 用id查询Job_env
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Job.Job_env true "用id查询Job_env"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /je/findJob_env [get]
func (jeApi *Job_envApi) FindJob_env(c *gin.Context) {
	var je Job.Job_env
	err := c.ShouldBindQuery(&je)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reje, err := jeService.GetJob_env(je.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reje": reje}, c)
	}
}

// GetJob_envList 分页获取Job_env列表
// @Tags Job_env
// @Summary 分页获取Job_env列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query JobReq.Job_envSearch true "分页获取Job_env列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /je/getJob_envList [get]
func (jeApi *Job_envApi) GetJob_envList(c *gin.Context) {
	var pageInfo JobReq.Job_envSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := jeService.GetJob_envInfoList(pageInfo); err != nil {
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
