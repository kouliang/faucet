package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Host    string    `mapstructure:"host"`

	VerifyHost string `mapstructure:"verifyHost"`

	*Web3Config     `mapstructure:"web3"`
	*LogConfig      `mapstructure:"log"`
	*DatabaseConfig `mapstructure:"database"`
}

type Web3Config struct {
	RpcUrl          string `mapstructure:"rpc_url"`
	PKey            string `mapstructure:"private_key"`
	ContractAddress string `mapstructure:"contract_address"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type DatabaseConfig struct {
	// Host         string `mapstructure:"host"`
	// User         string `mapstructure:"user"`
	// Password     string `mapstructure:"password"`
	// DbName       string `mapstructure:"dbname"`
	// Port         int    `mapstructure:"port"`
	// MaxOpenConns int    `mapstructure:"max_open_conns"`
	// MaxIdleConns int    `mapstructure:"max_idle_conns"`

	DBPath string `mapstructure:"dbpath"`
}

func Init(file string) (err error) {
	viper.SetConfigFile(file)

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	err = viper.Unmarshal(Conf)
	if err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	return
}
