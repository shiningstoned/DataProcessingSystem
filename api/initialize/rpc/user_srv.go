package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consul "github.com/kitex-contrib/registry-consul"
	user "hdfs/kitex_gen/user/userservice"
)

func InitUserService() user.Client {
	r, err := consul.NewConsulResolver("172.24.111.215:8500")
	if err != nil {
		hlog.Fatalf("new consul resolver failed: %s", err.Error())
	}

	client, err := user.NewClient(
		"user_srv",
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "user_srv",
		}))
	if err != nil {
		hlog.Fatalf("new user client failed: %s", err.Error())
	}
	return client
}
