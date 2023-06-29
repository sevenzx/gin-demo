// Author:      xuan
// Date:        2023/6/25
// Description:

package config

import (
	_ "embed"
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

//go:embed config.yml
var EmbedConfig []byte

func InitSetting() (err error) {
	// 先判断根目录下是否有config.yml
	var configFile *os.File
	configFile, err = os.Open("config.yml")
	if err != nil {
		// 使用EmbedConfig
		err = yaml.Unmarshal(EmbedConfig, &Config)
		if err != nil {
			return err
		}
		fmt.Println("use embed config")
		return nil
	} else {
		// 使用config.yml
		err = yaml.NewDecoder(configFile).Decode(&Config)
		if err != nil {
			fmt.Printf("Unable to parse the configuration file: %v\n", err)
			return err
		}
		_ = configFile.Close()
		fmt.Println("use config.yml")
		return nil
	}
}
