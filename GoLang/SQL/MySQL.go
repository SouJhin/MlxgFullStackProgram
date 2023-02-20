/*
 * @Author: SouJhin soujhin2022@gmail.com
 * @Date: 2023-02-18 10:05:41
 * @LastEditors: SouJhin soujhin2022@gmail.com
 * @LastEditTime: 2023-02-18 10:15:41
 * @FilePath: \MlxgFullStackProgram\GoLang\SQL\MySQL.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type MySQLInfo struct {
	userName  string
	password  string
	ipAddress string
	port      int
	dbName    string
	charset   string
}

func initSQL(err error) *xorm.Engine {
	conf := new(MySQLInfo)
	conf.charset = "utf8mb4"
	conf.userName = "root"
	conf.password = "mysql"
	conf.port = 3306
	conf.dbName = "user"
	conf.ipAddress = "127.0.0.1"
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		conf.userName, conf.password, conf.ipAddress, conf.port, conf.dbName, conf.charset)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("\"数据库连接失败\": %v\n", "数据库连接失败")
	}
	return engine
}
