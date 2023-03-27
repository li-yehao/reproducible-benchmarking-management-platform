package Job

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"

	// "strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Job"
	JobReq "github.com/flipped-aurora/gin-vue-admin/server/model/Job/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Cluster"
    ClusterReq "github.com/flipped-aurora/gin-vue-admin/server/model/Cluster/request"
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
    if info.Reason != "" {
        db = db.Where("reason LIKE ?","%"+ info.Reason+"%")
    }
    if info.Config != "" {
        db = db.Where("config LIKE ?","%"+ info.Config+"%")
    }
    if info.Job_status != "" {
        db = db.Where("job_status = ?",info.Job_status)
    }
    if info.Executor_name != "" {
        db = db.Where("executor_name = ?",info.Executor_name)
    }
    if info.Cmd_line != "" {
        db = db.Where("cmd_line LIKE ?","%"+ info.Cmd_line+"%")
    }
    if info.Results != "" {
        db = db.Where("results LIKE ?","%"+ info.Results+"%")
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

	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&jls).Error
	return  jls, total, err
}

// GetJob_list 根据id获取Job_list记录
// Author [piexlmax](https://github.com/piexlmax)
func (jlService *Job_listService)ExecuteJob_list(jl Job.Job_list) (err error) {
	err = global.GVA_DB.Where("id = ?", jl.ID).First(&jl).Error
    if err != nil {
        err = fmt.Errorf("id: %w", err)
        return
    }

    // todo: move to ~
    realCmd := fmt.Sprintf("bash ~/lyh/reproducible-benchmarking-management-platform/server/service/Job/dispatch.sh -l %s -c '%s' -i %d -r '%s'", jl.Cluster_name, jl.Cmd_line, jl.ID, jl.Reason)
    cmd := exec.Command("bash", "-c", realCmd)
    err = cmd.Start()
    if err != nil {
        err = fmt.Errorf("dispatch start: %w", err)
        return
    }

    // front-end will convert to GMT+8 automatically 
    loc, err := time.LoadLocation("Etc/GMT")
	if err != nil {
		err = fmt.Errorf("load location: %w", err)
		return
	}

    // use 'results' to store pid, to be able to cancel job based on pid
    pid := cmd.Process.Pid
    jl.Results = strconv.Itoa(pid)
    jl.Job_status = "WORKING"
    start_time := time.Now().In(loc).Format("2006-01-02 15:04:05")
    parsedStartTime, err := time.Parse("2006-01-02 15:04:05", start_time)
    jl.Start_time = &parsedStartTime
    if err != nil {
        err = fmt.Errorf("parse time: %w", err)
        return
    }
    err = global.GVA_DB.Save(&jl).Error
    if err != nil {
        err = fmt.Errorf("pid: %w", err)
        return
    }

    var stderr bytes.Buffer
	cmd.Stderr = &stderr
    err = cmd.Wait()
    if err != nil {
        err = fmt.Errorf("dispatch wait: %w, %s", err, stderr.String())
        return
    }

    // fetch results and status
    gatewayPath := fmt.Sprintf("http://10.238.153.58/nfsdata/rbmp/%d_%s/summary.json", jl.ID, jl.Reason)
    resultCmd := exec.Command("curl", gatewayPath)
    results, err := resultCmd.Output()
    if err != nil {
        err = fmt.Errorf("fetch results and status: %w", err)
        return
    }

    // update results and status
    jl.Results = string(results)
    var resultsMap map[string]interface{}
    json.Unmarshal(results, &resultsMap)
    if len(resultsMap) == 0 || resultsMap["status"] == nil {
        err = fmt.Errorf("results nil: %w", err)
        return
    }
    jl.Job_status = resultsMap["status"].(string)
    end_time := time.Now().In(loc).Format("2006-01-02 15:04:05")
    parsedEndTime, err := time.Parse("2006-01-02 15:04:05", end_time)
    jl.End_time = &parsedEndTime
    err = global.GVA_DB.Save(&jl).Error
    if err != nil {
        err = fmt.Errorf("update results and status: %w", err)
        return
    }

    // update cluster status from working to booked
    csService := Cluster.ClusterService{}
    info := ClusterReq.ClusterSearch{}
    info.Cluster_name = jl.Cluster_name
    css, _, err := csService.GetClusterInfoList(info)
    if err != nil {
        err = fmt.Errorf("get cluster status: %w", err)
        return
    }
    css[0].Working_status = "booked"
    err = csService.UpdateCluster(css[0])
    if err != nil {
        err = fmt.Errorf("update cluster status: %w", err)
        return
    }
	return
}

// GetJob_list 根据id获取Job_list记录
// Author [piexlmax](https://github.com/piexlmax)
func (jlService *Job_listService)CancelJob_list(jl Job.Job_list) (err error) {
	err = global.GVA_DB.Where("id = ?", jl.ID).First(&jl).Error
    if err != nil || jl.Job_status != "WORKING" {
        return
    }
    cmd := exec.Command("kill", "-9", jl.Results)
    err = cmd.Run()
    if err != nil {
        return
    }
    jl.Results = ""
    jl.Job_status = "CANCELED"
    loc, err := time.LoadLocation("Etc/GMT")
	if err != nil {
		err = fmt.Errorf("load location: %w", err)
		return
	}
    end_time := time.Now().In(loc).Format("2006-01-02 15:04:05")
    parsedTime, err := time.Parse("2006-01-02 15:04:05", end_time)
    jl.End_time = &parsedTime
    err = global.GVA_DB.Save(&jl).Error
	return
}