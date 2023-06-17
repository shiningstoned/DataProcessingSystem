package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Db       string
}

var mc MysqlConfig

func GetConfig() *MysqlConfig {
	return &mc
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./user/config")
	err := viper.ReadInConfig()
	if err != nil {
		klog.Fatalf("read in config failed: %s", err.Error())
	}
	err = viper.Unmarshal(&mc)
	if err != nil {
		klog.Fatalf("unmarshal config failed")
	}
}
