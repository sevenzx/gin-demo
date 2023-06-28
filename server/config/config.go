// Author:      xuan
// Date:        2023/6/25
// Description:

package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type ServerConfig struct {
	MySQL MySQL `yaml:"mysql"`
	Port  int   `yaml:"port"`
}

type MySQL struct {
	DBName   string `yaml:"db-name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var Config ServerConfig

func InitSetting() (err error) {
	// 打开配置文件
	configFile, err := os.Open("server/config/config.yml")
	if err != nil {
		fmt.Printf("Unable to open the configuration file: %v\n", err)
		return err
	}

	// 解析配置文件
	err = yaml.NewDecoder(configFile).Decode(&Config)
	if err != nil {
		fmt.Printf("Unable to parse the configuration file: %v\n", err)
		return err
	}

	_ = configFile.Close()
	return nil
}
