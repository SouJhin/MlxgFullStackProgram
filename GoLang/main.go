/*
 * @Author: SouJhin soujhin2022@gmail.com
 * @Date: 2023-02-20 09:27:32
 * @LastEditors: SouJhin soujhin2022@gmail.com
 * @LastEditTime: 2023-02-20 09:53:17
 * @FilePath: \MlxgFullStackProgram\GoLang\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	router "server/Router"
)

func main() {
	r := router.Router()
	r.Run()
}
