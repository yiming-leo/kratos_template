package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"kratos_test/internal/conf"
	"kratos_test/internal/data/mysql"
)

// Data 定义操作所有DB的结构体，统一将实现的DB接口放这进行IoC与DI
type Data struct {
	Mysql mysql.Mysql
}

// NewData 是Data结构体的初始化函数
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// 创建空Data实例
	d := &Data{}
	// 创建mysql客户端
	mysqlClient, err := mysql.NewMysql()
	if err != nil {
		return nil, nil, err
	}
	d.Mysql = mysqlClient
	// 创建redis客户端
	// 创建mongodb客户端
	// 创建influxdb客户端

	// 清理垃圾函数
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return d, cleanup, nil
}
