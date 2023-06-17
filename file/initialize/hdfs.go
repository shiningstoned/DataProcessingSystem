package initialize

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/colinmarc/hdfs/v2"
)

func InitHDFS() *hdfs.Client {
	client, err := hdfs.New("192.168.254.128:9000")
	if err != nil {
		hlog.Fatalf("connect to hdfs failed")
	}
	return client
}
