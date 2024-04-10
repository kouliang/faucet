package setting

import (
	"os"
)

var env_key = "FAUCET_CONF"

func SearchConfigPath() string {
	path := os.Getenv(env_key)
	if path == "" {
		generate_template()
	}
	return path
}

func generate_template() {
	template := `
title = "TapApp 配置文件"

name = "tap"
mode = ""        # "debug" "release" "test"
version = ""     # 版本号
host = ""        # gin监听地址
verifyHost = ""  # 任务验证地址

[web3]
rpc_url = ""          # web3服务的 rpc 地址
private_key = ""      # 分发代币的账户私钥
contract_address = "" # erc20 合约地址

[log]
level = ""           # log level
filename = ""        # 日志文件的位置
max_size = 200       # 日志文件的最大大小（以MB为单位），超过此大小将会进行切割
max_age = 30         # 保留旧文件的最大天数
max_backups = 7      # 保留旧文件的最大个数

[database]
dbpath = ""          #sqllite数据库文件路径
	`

	os.WriteFile("./faucet_config.toml", []byte(template), 0666)
}
