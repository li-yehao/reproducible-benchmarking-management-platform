package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Cluster"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type Cluster_bookingSearch struct{
    Cluster.Cluster_booking
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
