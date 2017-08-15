package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Conf 应用配置实例
var (
	conf      *Configuration
	parentDir string
)

// Configuration 应用配置
type Configuration struct {
	kv map[string]string
}

// InitConfiguration 初始化应用配置实例
func InitConfiguration(defaultJSON string) (err error) {
	if nil != conf {
		return
	}
	conf = &Configuration{
		kv: make(map[string]string),
	}
	err = fillDefault(defaultJSON)
	if nil != err {
		return
	}
	err = fillMaster()
	return
}

// GetValue 获取一个映射，值是字符串
func GetValue(key string) (value string) {
	value, isContain := conf.kv[key]
	if !isContain {
		fmt.Println("[ERROR] The mapping is not exists for : " + key)
	}
	return
}

// GetValues 获取一个映射，值是字符串数组，值的分隔符是","
func GetValues(key string) (values []string) {
	value, isContain := conf.kv[key]
	if !isContain {
		fmt.Println("[ERROR] The mapping is not exists for : " + key)
	}
	values = strings.Split(value, ",")
	for index, value := range values {
		values[index] = strings.TrimSpace(value)
	}
	return
}

// Set 设置一个映射
func Set(key string, value string) {
	conf.kv[key] = value
}

func fillDefault(defaultJSON string) (err error) {
	value, err := getMasterHost()
	if nil != err {
		return
	}
	Set(SysMasterHost, value)
	value, err = getLocalHost()
	if nil != err {
		return
	}
	Set(SysLocalHost, value)
	value, err = getServiceInfo()
	if nil != err {
		return
	}
	Set(SysServiceInfo, value)
	tmp, err := parseJSON(defaultJSON)
	if nil != err {
		return
	}
	for key, value := range tmp {
		Set(key, value)
	}
	return
}

func fillMaster() (err error) {
	jsonStr, err := request()
	if nil != err {
		return
	}
	tmp, err := parseJSON(jsonStr)
	if nil != err {
		return
	}
	for key, value := range tmp {
		Set(key, value)
	}
	return
}

// TODO
func request() (jsonStr string, err error) {
	jsonStr = `{}`
	return
}

func parseJSON(jsonStr string) (tmp map[string]string, err error) {
	err = json.Unmarshal([]byte(jsonStr), &tmp)
	return
}

func getMasterHost() (host string, err error) {
	host = os.Getenv(SysMasterHost)
	if "" == host {
		err = errors.New("[ERROR] ENV is not definded : " + SysMasterHost)
	}
	return
}

func getLocalHost() (host string, err error) {
	host = os.Getenv(SysLocalHost)
	if "" == host {
		err = errors.New("[ERROR] ENV is not definded : " + SysLocalHost)
	}
	return
}

func getServiceInfo() (info string, err error) {
	info = os.Getenv(SysServiceInfo)
	if "" == info {
		err = errors.New("[ERROR] ENV is not definded : " + SysServiceInfo)
	}
	return
}

func getCurrentDirectory() (err error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}
	currentDir := strings.Replace(dir, "\\", "/", -1)
	runes := []rune(currentDir)
	l := 0 + strings.LastIndex(currentDir, "/")
	if l > len(runes) {
		l = len(runes)
	}
	parentDir = string(runes[0:l])
	return
}
