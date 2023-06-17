package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consul "github.com/kitex-contrib/registry-consul"
	file "hdfs/kitex_gen/file/fileservice"
)

func InitFileService() file.Client {
	r, err := consul.NewConsulResolver("172.24.111.215:8500")
	if err != nil {
		hlog.Fatalf("new consul resolver failed: %s", err.Error())
	}

	client, err := file.NewClient(
		"file_srv",
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "file_srv",
		}))
	if err != nil {
		hlog.Fatalf("new file client failed: %s", err.Error())
	}
	return client
}
