package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
)

func InitMysql() *gorm.DB {
	username := GetConfig().Username
	password := GetConfig().Password
	host := GetConfig().Host
	port := GetConfig().Port
	database := GetConfig().Db
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		klog.Fatalf("connect to mysql failed: %s", err.Error())
	}
	return db

}
