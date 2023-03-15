// 自动生成模板Job_list
package Job

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Job_list 结构体
type Job_list struct {
      global.GVA_MODEL
      Cluster_name  string `json:"cluster_name" form:"cluster_name" gorm:"column:cluster_name;comment:;"`
      Executor_name  string `json:"executor_name" form:"executor_name" gorm:"column:executor_name;comment:;"`
      Start_time  *time.Time `json:"start_time" form:"start_time" gorm:"column:start_time;comment:;"`
      End_time  *time.Time `json:"end_time" form:"end_time" gorm:"column:end_time;comment:;"`
      Job_status  string `json:"job_status" form:"job_status" gorm:"column:job_status;type:enum('init', 'working', 'canceled', 'processing', 'process_failed', 'failed', 'done');comment:;"`
      Cmd_line  string `json:"cmd_line" form:"cmd_line" gorm:"column:cmd_line;comment:;"`
      Results  string `json:"results" form:"results" gorm:"column:results;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Job_list 表名
func (Job_list) TableName() string {
  return "job_list"
}

