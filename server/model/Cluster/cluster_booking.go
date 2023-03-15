// 自动生成模板Cluster_booking
package Cluster

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Cluster_booking 结构体
type Cluster_booking struct {
      global.GVA_MODEL
      Cluster_name  string `json:"cluster_name" form:"cluster_name" gorm:"column:cluster_name;comment:;"`
      Booker_name  string `json:"booker_name" form:"booker_name" gorm:"column:booker_name;comment:;"`
      Enabled  *bool `json:"enabled" form:"enabled" gorm:"column:enabled;comment:;"`
      Booking_reason  string `json:"booking_reason" form:"booking_reason" gorm:"column:booking_reason;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Cluster_booking 表名
func (Cluster_booking) TableName() string {
  return "cluster_booking"
}

