package config

import (
	"fmt"
	"xorm.io/xorm"
)

var SqlServer *xorm.Engine

type MysqlSqlParam struct {
	Host     string
	Port     string
	DataBase string
	UserName string
	Password string
}

// InitConfig 初始化项目配置
func InitConfig() {
	initDB()
}

func initDB() {
	var params MysqlSqlParam
	params.Host = "mysql-ksa-platbss-feature.hszq8.com"
	params.Port = "3306"
	params.DataBase = "hs_admin"
	params.UserName = "ksafeature_hs_admin"
	params.Password = "hs_admin"

	mySqlServer, err := InitMysql(params)
	SqlServer = mySqlServer
	if err != nil {
		fmt.Println("some errors - ", err.Error())
		panic(err.Error())
	}
	return
}
