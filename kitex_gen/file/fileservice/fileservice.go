// Code generated by Kitex v0.5.2. DO NOT EDIT.

package fileservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	file "hdfs/kitex_gen/file"
)

func serviceInfo() *kitex.ServiceInfo {
	return fileServiceServiceInfo
}

var fileServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FileService"
	handlerType := (*file.FileService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetFiles":     kitex.NewMethodInfo(getFilesHandler, newFileServiceGetFilesArgs, newFileServiceGetFilesResult, false),
		"RemoveRepeat": kitex.NewMethodInfo(removeRepeatHandler, newFileServiceRemoveRepeatArgs, newFileServiceRemoveRepeatResult, false),
		"SortByNum":    kitex.NewMethodInfo(sortByNumHandler, newFileServiceSortByNumArgs, newFileServiceSortByNumResult, false),
		"SortByTime":   kitex.NewMethodInfo(sortByTimeHandler, newFileServiceSortByTimeArgs, newFileServiceSortByTimeResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "file",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func getFilesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*file.FileServiceGetFilesArgs)
	realResult := result.(*file.FileServiceGetFilesResult)
	success, err := handler.(file.FileService).GetFiles(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFileServiceGetFilesArgs() interface{} {
	return file.NewFileServiceGetFilesArgs()
}

func newFileServiceGetFilesResult() interface{} {
	return file.NewFileServiceGetFilesResult()
}

func removeRepeatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*file.FileServiceRemoveRepeatArgs)
	realResult := result.(*file.FileServiceRemoveRepeatResult)
	success, err := handler.(file.FileService).RemoveRepeat(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFileServiceRemoveRepeatArgs() interface{} {
	return file.NewFileServiceRemoveRepeatArgs()
}

func newFileServiceRemoveRepeatResult() interface{} {
	return file.NewFileServiceRemoveRepeatResult()
}

func sortByNumHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*file.FileServiceSortByNumArgs)
	realResult := result.(*file.FileServiceSortByNumResult)
	success, err := handler.(file.FileService).SortByNum(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFileServiceSortByNumArgs() interface{} {
	return file.NewFileServiceSortByNumArgs()
}

func newFileServiceSortByNumResult() interface{} {
	return file.NewFileServiceSortByNumResult()
}

func sortByTimeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*file.FileServiceSortByTimeArgs)
	realResult := result.(*file.FileServiceSortByTimeResult)
	success, err := handler.(file.FileService).SortByTime(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFileServiceSortByTimeArgs() interface{} {
	return file.NewFileServiceSortByTimeArgs()
}

func newFileServiceSortByTimeResult() interface{} {
	return file.NewFileServiceSortByTimeResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetFiles(ctx context.Context, req *file.GetFilesRequest) (r *file.GetFilesResponse, err error) {
	var _args file.FileServiceGetFilesArgs
	_args.Req = req
	var _result file.FileServiceGetFilesResult
	if err = p.c.Call(ctx, "GetFiles", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RemoveRepeat(ctx context.Context, req *file.RemoveRepeatRequest) (r *file.RemoveRepeatResponse, err error) {
	var _args file.FileServiceRemoveRepeatArgs
	_args.Req = req
	var _result file.FileServiceRemoveRepeatResult
	if err = p.c.Call(ctx, "RemoveRepeat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SortByNum(ctx context.Context, req *file.SortByNumRequest) (r *file.SortByNumResponse, err error) {
	var _args file.FileServiceSortByNumArgs
	_args.Req = req
	var _result file.FileServiceSortByNumResult
	if err = p.c.Call(ctx, "SortByNum", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SortByTime(ctx context.Context, req *file.SortByNumRequest) (r *file.SortByTimeResponse, err error) {
	var _args file.FileServiceSortByTimeArgs
	_args.Req = req
	var _result file.FileServiceSortByTimeResult
	if err = p.c.Call(ctx, "SortByTime", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
