import service from '@/utils/request'

// @Tags Cluster_booking
// @Summary 创建Cluster_booking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cluster_booking true "创建Cluster_booking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cb/createCluster_booking [post]
export const createCluster_booking = (data) => {
  return service({
    url: '/cb/createCluster_booking',
    method: 'post',
    data
  })
}

// @Tags Cluster_booking
// @Summary 删除Cluster_booking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cluster_booking true "删除Cluster_booking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cb/deleteCluster_booking [delete]
export const deleteCluster_booking = (data) => {
  return service({
    url: '/cb/deleteCluster_booking',
    method: 'delete',
    data
  })
}

// @Tags Cluster_booking
// @Summary 删除Cluster_booking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cluster_booking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cb/deleteCluster_booking [delete]
export const deleteCluster_bookingByIds = (data) => {
  return service({
    url: '/cb/deleteCluster_bookingByIds',
    method: 'delete',
    data
  })
}

// @Tags Cluster_booking
// @Summary 更新Cluster_booking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cluster_booking true "更新Cluster_booking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cb/updateCluster_booking [put]
export const updateCluster_booking = (data) => {
  return service({
    url: '/cb/updateCluster_booking',
    method: 'put',
    data
  })
}

// @Tags Cluster_booking
// @Summary 用id查询Cluster_booking
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Cluster_booking true "用id查询Cluster_booking"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cb/findCluster_booking [get]
export const findCluster_booking = (params) => {
  return service({
    url: '/cb/findCluster_booking',
    method: 'get',
    params
  })
}

// @Tags Cluster_booking
// @Summary 分页获取Cluster_booking列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Cluster_booking列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cb/getCluster_bookingList [get]
export const getCluster_bookingList = (params) => {
  return service({
    url: '/cb/getCluster_bookingList',
    method: 'get',
    params
  })
}
