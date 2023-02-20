package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func Init() *xorm.Engine {
	var (
		userName  string = "root"
		password  string = "mysql"
		ipAddress string = "127.0.0.1"
		port      int    = 3306
		dbName    string = "test"
		charset   string = "utf8mb4"
	)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("\"数据库连接失败\": %v\n", "数据库连接失败")
	}
	err = engine.Sync(new(Food), new(Profit), new(User))
	if err != nil {
		println("connect failed")
	}
	return engine
}
