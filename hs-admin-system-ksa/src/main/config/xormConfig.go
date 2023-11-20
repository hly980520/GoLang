package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

var mysql *xorm.Engine

// InitMysql 初始化mysql连接
func InitMysql(params MysqlSqlParam) (*xorm.Engine, error) {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		params.UserName, params.Password, params.Host, params.Port, params.DataBase)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("mysql", "初始化数据库，创建连接异常:", err.Error())
		return nil, err
	}
	engine.TZLocation, _ = time.LoadLocation("Asia/Chongqing")
	engine.SetMaxIdleConns(3)
	engine.SetMaxOpenConns(20)
	engine.SetConnMaxLifetime(0)
	engine.ShowSQL(true)
	mysql = engine
	timer := time.NewTicker(time.Minute * 10)
	go func(conn *xorm.Engine) {
		for _ = range timer.C {
			if err = mysql.Ping(); err != nil {
				MySQLAutoConnect()
			}
		}
	}(mysql)
	return mysql, nil
}

func autoConnectMySQL(tryTimes int, maxTryTimes int) int {
	tryTimes++
	if tryTimes <= maxTryTimes {
		if mysql.Ping() != nil {
			message := fmt.Sprintf("数据库连接失败,已重连%d次", tryTimes)
			fmt.Println("mysql", message)
		}
		tryTimes = autoConnectMySQL(tryTimes, maxTryTimes)
	}
	return tryTimes
}

func MySQLAutoConnect() {
	autoConnectMySQL(0, 5)
}
