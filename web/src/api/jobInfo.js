import service from '@/utils/request'

// @Tags Job_info
// @Summary 创建Job_info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job_info true "创建Job_info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ji/createJob_info [post]
export const createJob_info = (data) => {
  return service({
    url: '/ji/createJob_info',
    method: 'post',
    data
  })
}

// @Tags Job_info
// @Summary 删除Job_info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job_info true "删除Job_info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ji/deleteJob_info [delete]
export const deleteJob_info = (data) => {
  return service({
    url: '/ji/deleteJob_info',
    method: 'delete',
    data
  })
}

// @Tags Job_info
// @Summary 删除Job_info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Job_info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ji/deleteJob_info [delete]
export const deleteJob_infoByIds = (data) => {
  return service({
    url: '/ji/deleteJob_infoByIds',
    method: 'delete',
    data
  })
}

// @Tags Job_info
// @Summary 更新Job_info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job_info true "更新Job_info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ji/updateJob_info [put]
export const updateJob_info = (data) => {
  return service({
    url: '/ji/updateJob_info',
    method: 'put',
    data
  })
}

// @Tags Job_info
// @Summary 用id查询Job_info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Job_info true "用id查询Job_info"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ji/findJob_info [get]
export const findJob_info = (params) => {
  return service({
    url: '/ji/findJob_info',
    method: 'get',
    params
  })
}

// @Tags Job_info
// @Summary 分页获取Job_info列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Job_info列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ji/getJob_infoList [get]
export const getJob_infoList = (params) => {
  return service({
    url: '/ji/getJob_infoList',
    method: 'get',
    params
  })
}
