package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Mysql Mysql `yaml:"mysql"`
	Gin   Gin   `yaml:"gin"`
	Redis Redis `yaml:"redis"`
}
type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Dbname   string `yaml:"dbname"`
	Timeout  string `yaml:"timeout"`
}

type Gin struct {
	Address string `yaml:"address"`
}

type Redis struct {
	Address string `yaml:"address"`
}

var Config Conf

// 初始化Config, 读取config.yaml文件
func init() {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("[config init error] ioutil.ReadFile 配置文件读取失败 " + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		fmt.Println("[config init error] yaml.Unmarshal 配置文件解析失败 " + err.Error())
		return
	}
	fmt.Println("[Success] 配置文件读取成功！！！")
}
