// Code generated by Kitex v0.5.2. DO NOT EDIT.

package fileservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	file "hdfs/kitex_gen/file"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetFiles(ctx context.Context, req *file.GetFilesRequest, callOptions ...callopt.Option) (r *file.GetFilesResponse, err error)
	RemoveRepeat(ctx context.Context, req *file.RemoveRepeatRequest, callOptions ...callopt.Option) (r *file.RemoveRepeatResponse, err error)
	SortByNum(ctx context.Context, req *file.SortByNumRequest, callOptions ...callopt.Option) (r *file.SortByNumResponse, err error)
	SortByTime(ctx context.Context, req *file.SortByNumRequest, callOptions ...callopt.Option) (r *file.SortByTimeResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFileServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFileServiceClient struct {
	*kClient
}

func (p *kFileServiceClient) GetFiles(ctx context.Context, req *file.GetFilesRequest, callOptions ...callopt.Option) (r *file.GetFilesResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFiles(ctx, req)
}

func (p *kFileServiceClient) RemoveRepeat(ctx context.Context, req *file.RemoveRepeatRequest, callOptions ...callopt.Option) (r *file.RemoveRepeatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RemoveRepeat(ctx, req)
}

func (p *kFileServiceClient) SortByNum(ctx context.Context, req *file.SortByNumRequest, callOptions ...callopt.Option) (r *file.SortByNumResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SortByNum(ctx, req)
}

func (p *kFileServiceClient) SortByTime(ctx context.Context, req *file.SortByNumRequest, callOptions ...callopt.Option) (r *file.SortByTimeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SortByTime(ctx, req)
}
