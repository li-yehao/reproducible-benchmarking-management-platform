package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Job"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type Job_infoSearch struct{
    Job.Job_info
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartStart_time  *time.Time  `json:"startStart_time" form:"startStart_time"`
    EndStart_time  *time.Time  `json:"endStart_time" form:"endStart_time"`
    StartEnd_time  *time.Time  `json:"startEnd_time" form:"startEnd_time"`
    EndEnd_time  *time.Time  `json:"endEnd_time" form:"endEnd_time"`
    request.PageInfo
}
