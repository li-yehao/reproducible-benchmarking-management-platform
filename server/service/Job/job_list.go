package Job

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Job"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    JobReq "github.com/flipped-aurora/gin-vue-admin/server/model/Job/request"
    "gorm.io/gorm"
)

type Job_listService struct {
}

// CreateJob_list 创建Job_list记录
// Author [piexlmax](https://github.com/piexlmax)
func (jlService *Job_listService) CreateJob_list(jl Job.Job_list) (err error) {
	err = global.GVA_DB.Create(&jl).Error
	return err
}

// DeleteJob_list 删除Job_list记录
// Author [piexlmax](https://github.com/piexlmax)
func (jlService *Job_listService)DeleteJob_list(jl Job.Job_list) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Job.Job_list{}).Where("id = ?", jl.ID).Update("deleted_by", jl.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&jl).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteJob_listByIds 批量删除Job_list记录
// Author [piexlmax](https://github.com/piexlmax)
func (jlService *Job_listService)DeleteJob_listByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Job.Job_list{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&Job.Job_list{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateJob_list 更新Job_list记录
// Author [piexlmax](https://github.com/piexlmax)
func (jlService *Job_listService)UpdateJob_list(jl Job.Job_list) (err error) {
	err = global.GVA_DB.Save(&jl).Error
	return err
}

// GetJob_list 根据id获取Job_list记录
// Author [piexlmax](https://github.com/piexlmax)
func (jlService *Job_listService)GetJob_list(id uint) (jl Job.Job_list, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&jl).Error
	return
}

// GetJob_listInfoList 分页获取Job_list记录
// Author [piexlmax](https://github.com/piexlmax)
func (jlService *Job_listService)GetJob_listInfoList(info JobReq.Job_listSearch) (list []Job.Job_list, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Job.Job_list{})
    var jls []Job.Job_list
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Cluster_name != "" {
        db = db.Where("cluster_name = ?",info.Cluster_name)
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
    if info.Job_status != "" {
        db = db.Where("job_status = ?",info.Job_status)
    }
    if info.Cmd_line != "" {
        db = db.Where("cmd_line LIKE ?","%"+ info.Cmd_line+"%")
    }
    if info.Results != "" {
        db = db.Where("results LIKE ?","%"+ info.Results+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&jls).Error
	return  jls, total, err
}
