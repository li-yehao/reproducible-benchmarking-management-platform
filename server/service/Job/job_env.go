package Job

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Job"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    JobReq "github.com/flipped-aurora/gin-vue-admin/server/model/Job/request"
    "gorm.io/gorm"
)

type Job_envService struct {
}

// CreateJob_env 创建Job_env记录
// Author [piexlmax](https://github.com/piexlmax)
func (jeService *Job_envService) CreateJob_env(je Job.Job_env) (err error) {
	err = global.GVA_DB.Create(&je).Error
	return err
}

// DeleteJob_env 删除Job_env记录
// Author [piexlmax](https://github.com/piexlmax)
func (jeService *Job_envService)DeleteJob_env(je Job.Job_env) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Job.Job_env{}).Where("id = ?", je.ID).Update("deleted_by", je.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&je).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteJob_envByIds 批量删除Job_env记录
// Author [piexlmax](https://github.com/piexlmax)
func (jeService *Job_envService)DeleteJob_envByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Job.Job_env{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&Job.Job_env{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateJob_env 更新Job_env记录
// Author [piexlmax](https://github.com/piexlmax)
func (jeService *Job_envService)UpdateJob_env(je Job.Job_env) (err error) {
	err = global.GVA_DB.Save(&je).Error
	return err
}

// GetJob_env 根据id获取Job_env记录
// Author [piexlmax](https://github.com/piexlmax)
func (jeService *Job_envService)GetJob_env(id uint) (je Job.Job_env, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&je).Error
	return
}

// GetJob_envInfoList 分页获取Job_env记录
// Author [piexlmax](https://github.com/piexlmax)
func (jeService *Job_envService)GetJob_envInfoList(info JobReq.Job_envSearch) (list []Job.Job_env, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Job.Job_env{})
    var jes []Job.Job_env
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Job_id != nil {
        db = db.Where("job_id = ?",info.Job_id)
    }
    if info.Catlog != "" {
        db = db.Where("catlog = ?",info.Catlog)
    }
    if info.Sub_catlog != "" {
        db = db.Where("sub_catlog = ?",info.Sub_catlog)
    }
    if info.Key_info != "" {
        db = db.Where("key_info = ?",info.Key_info)
    }
    if info.Info_value != "" {
        db = db.Where("info_value = ?",info.Info_value)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&jes).Error
	return  jes, total, err
}
