// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/colinmarc/hdfs/v2"
	"hdfs/api/biz/middleware"
	"hdfs/api/global"
	"hdfs/kitex_gen/file"
	"hdfs/kitex_gen/user"
	"io"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api "hdfs/api/biz/model/api"
)

// Login .
// @router /user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := global.UserClient.Login(ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
		return
	}
	j := middleware.NewJWT()
	token, err := j.CreateToken(resp.Username)
	if err != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
	}

	c.JSON(consts.StatusOK, &api.LoginResponse{Token: token})
}

// Regiter .
// @router /user/register [POST]
func Regiter(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := global.UserClient.Register(ctx, &user.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(consts.StatusBadRequest, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// UploadFile .
// @router /file/upload [POST]
func UploadFile(ctx context.Context, c *app.RequestContext) {
	var err error
	username, _ := c.Get("username")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	src, err := fileHeader.Open()
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	defer src.Close()
	client, err := hdfs.New("192.168.254.128:9000")
	dst, err := client.Create("/user/" + username.(string) + fileHeader.Filename)
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	c.JSON(consts.StatusOK, "upload file success")
}

// DownloadFile .
// @router /file/download [GET]
func DownloadFile(ctx context.Context, c *app.RequestContext) {
	var err error
	username, _ := c.Get("username")
	var req api.FileDownloadRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	dst, err := os.Create("temp/" + req.Filename)
	if err != nil {
		hlog.Errorf("create temp file failed")
		return
	}
	defer dst.Close()
	client, err := hdfs.New("192.168.254.128:9000")
	src, err := client.Open("/user/" + username.(string) + req.Filename)
	if err != nil {
		hlog.Errorf("get hdfs file failed")
		return
	}
	defer src.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		hlog.Errorf("copy hdfs file failed")
		return
	}

	c.File("temp/" + req.Filename)
	os.Remove("temp/" + req.Filename)
}

// GetDir .
// @router /file/gitdir [GET]
func GetDir(ctx context.Context, c *app.RequestContext) {
	var err error
	username, _ := c.Get("username")
	var req api.GetDirRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	files, err := global.FileClient.GetFiles(ctx, &file.GetFilesRequest{Username: username.(string)})
	if err != nil {
		hlog.Errorf("get file failed")
		return
	}

	c.JSON(consts.StatusOK, files)
}

// RemoveRepeat .
// @router /file/rmrepeat [POST]
func RemoveRepeat(ctx context.Context, c *app.RequestContext) {
	var err error
	username, _ := c.Get("username")
	var req api.RemoveRepeatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	path := "/user/" + username.(string) + "/" + req.Filename
	resp, err := global.FileClient.RemoveRepeat(ctx, &file.RemoveRepeatRequest{Filename: path})
	if err != nil {
		hlog.Errorf("file remove repeat failed")
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// SortByNum .
// @router /file/sort [POST]
func SortByNum(ctx context.Context, c *app.RequestContext) {
	var err error
	username, _ := c.Get("username")
	var req api.SortByNumRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	path := "/user/" + username.(string) + "/" + req.Filename

	resp, err := global.FileClient.SortByNum(ctx, &file.SortByNumRequest{Filename: path})
	if err != nil {
		hlog.Errorf("file sort by num error")
		return
	}

	c.JSON(consts.StatusOK, resp)
}