package Cluster

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Cluster"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ClusterReq "github.com/flipped-aurora/gin-vue-admin/server/model/Cluster/request"
)

type ClusterService struct {
}

// CreateCluster 创建Cluster记录
// Author [piexlmax](https://github.com/piexlmax)
func (csService *ClusterService) CreateCluster(cs Cluster.Cluster) (err error) {
	err = global.GVA_DB.Create(&cs).Error
	return err
}

// DeleteCluster 删除Cluster记录
// Author [piexlmax](https://github.com/piexlmax)
func (csService *ClusterService)DeleteCluster(cs Cluster.Cluster) (err error) {
	err = global.GVA_DB.Delete(&cs).Error
	return err
}

// DeleteClusterByIds 批量删除Cluster记录
// Author [piexlmax](https://github.com/piexlmax)
func (csService *ClusterService)DeleteClusterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Cluster.Cluster{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCluster 更新Cluster记录
// Author [piexlmax](https://github.com/piexlmax)
func (csService *ClusterService)UpdateCluster(cs Cluster.Cluster) (err error) {
	err = global.GVA_DB.Save(&cs).Error
	return err
}

// GetCluster 根据id获取Cluster记录
// Author [piexlmax](https://github.com/piexlmax)
func (csService *ClusterService)GetCluster(id uint) (cs Cluster.Cluster, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cs).Error
	return
}

// GetClusterInfoList 分页获取Cluster记录
// Author [piexlmax](https://github.com/piexlmax)
func (csService *ClusterService)GetClusterInfoList(info ClusterReq.ClusterSearch) (list []Cluster.Cluster, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Cluster.Cluster{})
    var css []Cluster.Cluster
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Cluster_name != "" {
        db = db.Where("cluster_name = ?",info.Cluster_name)
    }
    if info.Pmu_name != "" {
        db = db.Where("pmu_name = ?",info.Pmu_name)
    }
    if info.Working_status != "" {
        db = db.Where("working_status = ?",info.Working_status)
    }
    if info.Health_status != "" {
        db = db.Where("health_status = ?",info.Health_status)
    }
    if info.Booker_name != "" {
        db = db.Where("booker_name = ?",info.Booker_name)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&css).Error
	return  css, total, err
}
