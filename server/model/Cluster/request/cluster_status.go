package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Cluster"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ClusterSearch struct{
    Cluster.Cluster
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
