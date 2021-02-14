/*
* @Time    : 2021-02-05 19:41
* @Author  : CoderCharm
* @File    : mysql.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 获取 mysql 配置
**/

package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

//获取 mysql 的配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	_ = config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}
