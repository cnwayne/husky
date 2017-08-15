package main

import (
	"fmt"
	"strings"

	conf "./modules/config"
	log "./modules/logging"
	http "./modules/server"
)

func main() {
	var err error
	/*-
	 * 初始化应用配置
	 */
	err = conf.InitConfiguration(defaultConfJSON())
	if nil != err {
		fmt.Println("初始化应用配置 : " + fmt.Sprint(err))
		return
	}
	/*-
	 * 初始化日志框架
	 */
	err = initLogger()
	if nil != err {
		fmt.Println("初始化日志框架 : " + fmt.Sprint(err))
		return
	}
	/*-
	 * 启动监听
	 */
	err = startup()
	if nil != err {
		fmt.Println("启动监听 : " + fmt.Sprint(err))
		return
	}
}

func defaultConfJSON() (JSONStr string) {
	//%{time:2006-01-02 15:04:05} [%{level:.4s}] %{shortfunc} ▶ %{message}
	JSONStr = `
	{
		"` + conf.AppLogFile + `" : ""
		, "` + conf.AppLogFormat + `" : "%{color}%{time:2006-01-02 15:04:05} [%{level:.4s}] %{shortfile} ▶ %{message}%{color:reset}"
		, "` + conf.AppLogLevel + `" : "debug"
		, "` + conf.AppCatchDir + `" : "/tmp/"
	}
	`
	return
}

func initLogger() (err error) {
	err = log.InitLogger(
		conf.GetValue(conf.AppLogFile),
		conf.GetValue(conf.AppLogFormat),
		strings.ToUpper(conf.GetValue(conf.AppLogLevel)))
	return
}

func startup() (err error) {
	err = http.Listen(":18080")
	return
}
