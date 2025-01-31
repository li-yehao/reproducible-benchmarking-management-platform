import service from '@/utils/request'

// @Tags Cluster
// @Summary 创建Cluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cluster true "创建Cluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cs/createCluster [post]
export const createCluster = (data) => {
  return service({
    url: '/cs/createCluster',
    method: 'post',
    data
  })
}

// @Tags Cluster
// @Summary 删除Cluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cluster true "删除Cluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cs/deleteCluster [delete]
export const deleteCluster = (data) => {
  return service({
    url: '/cs/deleteCluster',
    method: 'delete',
    data
  })
}

// @Tags Cluster
// @Summary 删除Cluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cs/deleteCluster [delete]
export const deleteClusterByIds = (data) => {
  return service({
    url: '/cs/deleteClusterByIds',
    method: 'delete',
    data
  })
}

// @Tags Cluster
// @Summary 更新Cluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cluster true "更新Cluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cs/updateCluster [put]
export const updateCluster = (data) => {
  return service({
    url: '/cs/updateCluster',
    method: 'put',
    data
  })
}

// @Tags Cluster
// @Summary 用id查询Cluster
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Cluster true "用id查询Cluster"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cs/findCluster [get]
export const findCluster = (params) => {
  return service({
    url: '/cs/findCluster',
    method: 'get',
    params
  })
}

// @Tags Cluster
// @Summary 分页获取Cluster列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Cluster列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cs/getClusterList [get]
export const getClusterList = (params) => {
  return service({
    url: '/cs/getClusterList',
    method: 'get',
    params
  })
}
