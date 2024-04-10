package main

import (
	"faucet-app/database"
	"faucet-app/route"
	"faucet-app/setting"
	"faucet-app/web3"
	"fmt"
	"log"
	"os"
)

func main() {
	var configPath string
	if len(os.Args) < 2 {
		configPath = setting.SearchConfigPath()
		if configPath == "" {
			log.Println("No configuration file found, you can use the FAUCET_CONF environment variable to specify the profile path. A profile template has been automatically generated for you, named faucet_config.toml.")
			return
		}
	} else {
		configPath = os.Args[1]
	}

	// 1.加载配置
	if err := setting.Init(configPath); err != nil {
		fmt.Printf("init setting failed, err:%v\n", err)
		return
	}

	// 2.初始化web3
	web3.Init(setting.Conf.Web3Config)

	// 3.初始化数据库
	if err := database.Init(setting.Conf.DatabaseConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer database.Close()

	// 4.注册路由 & 启动服务
	router := route.Init(setting.Conf.Mode, setting.Conf.LogConfig)
	router.Run(setting.Conf.Host)
}
