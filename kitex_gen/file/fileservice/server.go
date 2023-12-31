// Code generated by Kitex v0.5.2. DO NOT EDIT.
package fileservice

import (
	server "github.com/cloudwego/kitex/server"
	file "hdfs/kitex_gen/file"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler file.FileService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
