package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Cluster"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Job"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Cluster Cluster.RouterGroup
	Job     Job.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
