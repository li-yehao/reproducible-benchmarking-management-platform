// 自动生成模板Job_env
package Job

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Job_env 结构体
type Job_env struct {
      global.GVA_MODEL
      Job_id  *int `json:"job_id" form:"job_id" gorm:"column:job_id;comment:;"`
      Catlog  string `json:"catlog" form:"catlog" gorm:"column:catlog;comment:;"`
      Sub_catlog  string `json:"sub_catlog" form:"sub_catlog" gorm:"column:sub_catlog;comment:;"`
      Key_info  string `json:"key_info" form:"key_info" gorm:"column:key_info;comment:;"`
      Info_value  string `json:"info_value" form:"info_value" gorm:"column:info_value;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Job_env 表名
func (Job_env) TableName() string {
  return "job_env"
}

