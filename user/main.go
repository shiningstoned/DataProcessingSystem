package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	user "hdfs/kitex_gen/user/userservice"
	"hdfs/user/dal"
	"hdfs/user/initialize"
	"log"
)

func main() {
	initialize.InitConfig()
	db := initialize.InitMysql()
	r, info := initialize.InitRegistry()
	svr := user.NewServer(&UserServiceImpl{
		dal.NewUserManager(db),
	},
		server.WithRegistry(r),
		server.WithRegistryInfo(&info),
		server.WithServiceAddr(utils.NewNetAddr("tcp", "172.24.111.215:10000")),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user_srv"}),
	)

	err := svr.Run()

	if err != nil {
		log.Fatalln(err.Error())
	}
}
