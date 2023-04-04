// 自动生成模板Cluster
package Cluster

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Cluster 结构体
type Cluster struct {
      global.GVA_MODEL
      Cluster_name  string `json:"cluster_name" form:"cluster_name" gorm:"column:cluster_name;comment:;"`
      Pmu_name  string `json:"pmu_name" form:"pmu_name" gorm:"column:pmu_name;comment:;"`
      Working_status  string `json:"working_status" form:"working_status" gorm:"column:working_status;type:enum('available', 'booked', 'working');comment:;"`
      Health_status  string `json:"health_status" form:"health_status" gorm:"column:health_status;comment:;"`
      Booker_name  string `json:"booker_name" form:"booker_name" gorm:"column:booker_name;comment:;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:;"`
}


// TableName Cluster 表名
func (Cluster) TableName() string {
  return "cluster_status"
}

