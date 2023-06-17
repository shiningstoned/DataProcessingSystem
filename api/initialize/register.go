package initialize

import (
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/hashicorp/consul/api"
	consul "github.com/hertz-contrib/registry/consul"
)

func InitRegistry() (registry.Registry, registry.Info) {
	config := api.DefaultConfig()
	config.Address = "172.24.111.215:8500"
	client, err := api.NewClient(config)
	if err != nil {
		hlog.Fatalf("new consul client failed: %s", err.Error())
	}

	r := consul.NewConsulRegister(client,
		consul.WithCheck(&api.AgentServiceCheck{
			Interval:                       "5s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "15s",
		}))
	info := registry.Info{
		ServiceName: "api",
		Addr:        utils.NewNetAddr("tcp", "172.24.111.215:8500"),
		Weight:      10,
	}
	return r, info
}
