/*
 * @Author: SouJhin soujhin2022@gmail.com
 * @Date: 2023-02-20 09:30:27
 * @LastEditors: SouJhin soujhin2022@gmail.com
 * @LastEditTime: 2023-02-20 09:31:51
 * @FilePath: \MlxgFullStackProgram\GoLang\Model\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

type User struct {
	Id          int `xorm:"not nul"`
	Name        string
	Score       int
	IsStudent   bool
	IsVip       bool
	PhoneNumber string
	Address     string
	Avatar      string
}
