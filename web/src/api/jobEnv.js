import service from '@/utils/request'

// @Tags Job_env
// @Summary 创建Job_env
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job_env true "创建Job_env"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /je/createJob_env [post]
export const createJob_env = (data) => {
  return service({
    url: '/je/createJob_env',
    method: 'post',
    data
  })
}

// @Tags Job_env
// @Summary 删除Job_env
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job_env true "删除Job_env"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /je/deleteJob_env [delete]
export const deleteJob_env = (data) => {
  return service({
    url: '/je/deleteJob_env',
    method: 'delete',
    data
  })
}

// @Tags Job_env
// @Summary 删除Job_env
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Job_env"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /je/deleteJob_env [delete]
export const deleteJob_envByIds = (data) => {
  return service({
    url: '/je/deleteJob_envByIds',
    method: 'delete',
    data
  })
}

// @Tags Job_env
// @Summary 更新Job_env
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job_env true "更新Job_env"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /je/updateJob_env [put]
export const updateJob_env = (data) => {
  return service({
    url: '/je/updateJob_env',
    method: 'put',
    data
  })
}

// @Tags Job_env
// @Summary 用id查询Job_env
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Job_env true "用id查询Job_env"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /je/findJob_env [get]
export const findJob_env = (params) => {
  return service({
    url: '/je/findJob_env',
    method: 'get',
    params
  })
}

// @Tags Job_env
// @Summary 分页获取Job_env列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Job_env列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /je/getJob_envList [get]
export const getJob_envList = (params) => {
  return service({
    url: '/je/getJob_envList',
    method: 'get',
    params
  })
}
