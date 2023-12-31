// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"hdfs/api/initialize"
	"hdfs/api/initialize/rpc"
)

func main() {
	rpc.InitRPC()
	r, info := initialize.InitRegistry()
	h := server.New(server.WithHostPorts(":8080"),
		server.WithRegistry(r, &info))

	register(h)
	h.Spin()
}
