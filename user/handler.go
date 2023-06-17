package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/status"
	user "hdfs/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	MysqlManager
}

type MysqlManager interface {
	CreateUser(username, password string) (bool, error)
	LoginCheck(username, password string) (bool, string, error)
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.CommonResponse, err error) {
	// TODO: Your code here...
	exist, err := s.CreateUser(req.Username, req.Password)
	if err != nil {
		klog.Errorf("register error: ", err)
		return nil, status.Err(codes.Internal, "register error")
	}
	if exist {
		return &user.CommonResponse{Message: "user already exist"}, nil
	}
	return &user.CommonResponse{Message: "register success"}, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	flag, username, err := s.LoginCheck(req.Username, req.Password)
	if err != nil {
		klog.Error("login error")
		return nil, status.Err(codes.Internal, "login error")
	}
	if !flag {
		klog.Info("wrong password")
		return nil, status.Err(codes.Internal, "wrong password")
	}
	return &user.LoginResponse{Username: username}, nil
}
