package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"hdfs/file/initialize"
	"hdfs/file/pkg"
	file "hdfs/kitex_gen/file/fileservice"
	"log"
)

func main() {
	r, info := initialize.InitRegistry()
	hdfs := initialize.InitHDFS()
	svr := file.NewServer(&FileServiceImpl{
		HDFS: pkg.NewHDFS(hdfs),
	},
		server.WithRegistry(r),
		server.WithRegistryInfo(&info),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "file_srv"}),
		server.WithServiceAddr(utils.NewNetAddr("tcp", "172.24.111.215.10001")))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
