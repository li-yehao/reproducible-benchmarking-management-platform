// 自动生成模板Job_info
package Job

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Job_info 结构体
type Job_info struct {
      global.GVA_MODEL
      Cluster_name  string `json:"cluster_name" form:"cluster_name" gorm:"column:cluster_name;comment:;"`
      Workload_name  string `json:"workload_name" form:"workload_name" gorm:"column:workload_name;comment:;"`
      Job_status  string `json:"job_status" form:"job_status" gorm:"column:job_status;type:enum('init', 'canceled', 'working', 'failure', 'done', 'processing', 'process_failure');comment:;"`
      Svrinfo  *bool `json:"svrinfo" form:"svrinfo" gorm:"column:svrinfo;comment:;"`
      Emon  *bool `json:"emon" form:"emon" gorm:"column:emon;comment:;"`
      Executor_name  string `json:"executor_name" form:"executor_name" gorm:"column:executor_name;comment:;"`
      Start_time  *time.Time `json:"start_time" form:"start_time" gorm:"column:start_time;comment:;"`
      End_time  *time.Time `json:"end_time" form:"end_time" gorm:"column:end_time;comment:;"`
      Results  string `json:"results" form:"results" gorm:"column:results;comment:;"`
      Log_path  string `json:"log_path" form:"log_path" gorm:"column:log_path;comment:;"`
      Svrinfo_path  string `json:"svrinfo_path" form:"svrinfo_path" gorm:"column:svrinfo_path;comment:;"`
      Emon_path  string `json:"emon_path" form:"emon_path" gorm:"column:emon_path;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Job_info 表名
func (Job_info) TableName() string {
  return "job_info"
}

