package Job

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Job"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    JobReq "github.com/flipped-aurora/gin-vue-admin/server/model/Job/request"
    "gorm.io/gorm"
)

type Job_infoService struct {
}

// CreateJob_info 创建Job_info记录
// Author [piexlmax](https://github.com/piexlmax)
func (jiService *Job_infoService) CreateJob_info(ji Job.Job_info) (err error) {
	err = global.GVA_DB.Create(&ji).Error
	return err
}

// DeleteJob_info 删除Job_info记录
// Author [piexlmax](https://github.com/piexlmax)
func (jiService *Job_infoService)DeleteJob_info(ji Job.Job_info) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Job.Job_info{}).Where("id = ?", ji.ID).Update("deleted_by", ji.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&ji).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteJob_infoByIds 批量删除Job_info记录
// Author [piexlmax](https://github.com/piexlmax)
func (jiService *Job_infoService)DeleteJob_infoByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Job.Job_info{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&Job.Job_info{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateJob_info 更新Job_info记录
// Author [piexlmax](https://github.com/piexlmax)
func (jiService *Job_infoService)UpdateJob_info(ji Job.Job_info) (err error) {
	err = global.GVA_DB.Save(&ji).Error
	return err
}

// GetJob_info 根据id获取Job_info记录
// Author [piexlmax](https://github.com/piexlmax)
func (jiService *Job_infoService)GetJob_info(id uint) (ji Job.Job_info, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ji).Error
	return
}

// GetJob_infoInfoList 分页获取Job_info记录
// Author [piexlmax](https://github.com/piexlmax)
func (jiService *Job_infoService)GetJob_infoInfoList(info JobReq.Job_infoSearch) (list []Job.Job_info, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Job.Job_info{})
    var jis []Job.Job_info
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Cluster_name != "" {
        db = db.Where("cluster_name = ?",info.Cluster_name)
    }
    if info.Workload_name != "" {
        db = db.Where("workload_name = ?",info.Workload_name)
    }
    if info.Job_status != "" {
        db = db.Where("job_status = ?",info.Job_status)
    }
    if info.Svrinfo != nil {
        db = db.Where("svrinfo = ?",info.Svrinfo)
    }
    if info.Emon != nil {
        db = db.Where("emon = ?",info.Emon)
    }
    if info.Executor_name != "" {
        db = db.Where("executor_name = ?",info.Executor_name)
    }
        if info.StartStart_time != nil && info.EndStart_time != nil {
            db = db.Where("start_time BETWEEN ? AND ? ",info.StartStart_time,info.EndStart_time)
        }
        if info.StartEnd_time != nil && info.EndEnd_time != nil {
            db = db.Where("end_time BETWEEN ? AND ? ",info.StartEnd_time,info.EndEnd_time)
        }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&jis).Error
	return  jis, total, err
}
