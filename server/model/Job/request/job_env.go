package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Job"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type Job_envSearch struct{
    Job.Job_env
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
