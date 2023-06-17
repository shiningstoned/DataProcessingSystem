package rpc

import "hdfs/api/global"

func InitRPC() {
	global.UserClient = InitUserService()
	global.FileClient = InitFileService()
}
