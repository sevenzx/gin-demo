// Author:      xuan
// Date:        2023/6/25
// Description:

package setting

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type YmlConfig struct {
	Database Database `yaml:"database"`
	Port     int      `yaml:"port"`
}

type Database struct {
	DB       string `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var Config YmlConfig

func InitSetting() (err error) {
	// 打开配置文件
	configFile, err := os.Open("./conf/config.yml")
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
