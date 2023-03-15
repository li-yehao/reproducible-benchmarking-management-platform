import service from '@/utils/request'

// @Tags Job_list
// @Summary 创建Job_list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job_list true "创建Job_list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jl/createJob_list [post]
export const createJob_list = (data) => {
  return service({
    url: '/jl/createJob_list',
    method: 'post',
    data
  })
}

// @Tags Job_list
// @Summary 删除Job_list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job_list true "删除Job_list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jl/deleteJob_list [delete]
export const deleteJob_list = (data) => {
  return service({
    url: '/jl/deleteJob_list',
    method: 'delete',
    data
  })
}

// @Tags Job_list
// @Summary 删除Job_list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Job_list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /jl/deleteJob_list [delete]
export const deleteJob_listByIds = (data) => {
  return service({
    url: '/jl/deleteJob_listByIds',
    method: 'delete',
    data
  })
}

// @Tags Job_list
// @Summary 更新Job_list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Job_list true "更新Job_list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /jl/updateJob_list [put]
export const updateJob_list = (data) => {
  return service({
    url: '/jl/updateJob_list',
    method: 'put',
    data
  })
}

// @Tags Job_list
// @Summary 用id查询Job_list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Job_list true "用id查询Job_list"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /jl/findJob_list [get]
export const findJob_list = (params) => {
  return service({
    url: '/jl/findJob_list',
    method: 'get',
    params
  })
}

// @Tags Job_list
// @Summary 分页获取Job_list列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Job_list列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /jl/getJob_listList [get]
export const getJob_listList = (params) => {
  return service({
    url: '/jl/getJob_listList',
    method: 'get',
    params
  })
}
