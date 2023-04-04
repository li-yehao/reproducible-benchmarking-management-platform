package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Cluster"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Job"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	ClusterServiceGroup Cluster.ServiceGroup
	JobServiceGroup     Job.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
