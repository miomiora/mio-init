package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Mysql Mysql `yaml:"mysql"`
	Gin   Gin   `yaml:"gin"`
}
type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Timeout  string `yaml:"timeout"`
}

type Gin struct {
	Port string `yaml:"port"`
}

var Config Conf

// 初始化Config, 读取config.yaml文件
func init() {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("配置文件读取失败 " + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		fmt.Println("配置文件解析失败 " + err.Error())
	}
	fmt.Println("配置文件读取成功！")
}
