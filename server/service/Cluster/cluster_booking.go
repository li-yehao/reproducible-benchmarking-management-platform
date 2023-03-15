package Cluster

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Cluster"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ClusterReq "github.com/flipped-aurora/gin-vue-admin/server/model/Cluster/request"
    "gorm.io/gorm"
)

type Cluster_bookingService struct {
}

// CreateCluster_booking 创建Cluster_booking记录
// Author [piexlmax](https://github.com/piexlmax)
func (cbService *Cluster_bookingService) CreateCluster_booking(cb Cluster.Cluster_booking) (err error) {
	err = global.GVA_DB.Create(&cb).Error
	return err
}

// DeleteCluster_booking 删除Cluster_booking记录
// Author [piexlmax](https://github.com/piexlmax)
func (cbService *Cluster_bookingService)DeleteCluster_booking(cb Cluster.Cluster_booking) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Cluster.Cluster_booking{}).Where("id = ?", cb.ID).Update("deleted_by", cb.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&cb).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCluster_bookingByIds 批量删除Cluster_booking记录
// Author [piexlmax](https://github.com/piexlmax)
func (cbService *Cluster_bookingService)DeleteCluster_bookingByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Cluster.Cluster_booking{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&Cluster.Cluster_booking{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCluster_booking 更新Cluster_booking记录
// Author [piexlmax](https://github.com/piexlmax)
func (cbService *Cluster_bookingService)UpdateCluster_booking(cb Cluster.Cluster_booking) (err error) {
	err = global.GVA_DB.Save(&cb).Error
	return err
}

// GetCluster_booking 根据id获取Cluster_booking记录
// Author [piexlmax](https://github.com/piexlmax)
func (cbService *Cluster_bookingService)GetCluster_booking(id uint) (cb Cluster.Cluster_booking, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cb).Error
	return
}

// GetCluster_bookingInfoList 分页获取Cluster_booking记录
// Author [piexlmax](https://github.com/piexlmax)
func (cbService *Cluster_bookingService)GetCluster_bookingInfoList(info ClusterReq.Cluster_bookingSearch) (list []Cluster.Cluster_booking, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Cluster.Cluster_booking{})
    var cbs []Cluster.Cluster_booking
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Cluster_name != "" {
        db = db.Where("cluster_name = ?",info.Cluster_name)
    }
    if info.Booker_name != "" {
        db = db.Where("booker_name = ?",info.Booker_name)
    }
    if info.Enabled != nil {
        db = db.Where("enabled = ?",info.Enabled)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&cbs).Error
	return  cbs, total, err
}
