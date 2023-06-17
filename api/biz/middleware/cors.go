package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Cors() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		method := c.Request.Method()
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if string(method) == "Option" {
			c.JSON(http.StatusOK, "ok")
		}
		c.Next(ctx)
	}
}
