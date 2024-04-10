package web3

import (
	"faucet-app/setting"
	"faucet-app/web3/account"
	"faucet-app/web3/client"
)

var kaccount *account.KLAccount
var erc20Address string

func Init(cfg *setting.Web3Config) {
	initContractAbi()
	erc20Address = cfg.ContractAddress

	err := client.InitStd(cfg.RpcUrl)
	if err != nil {
		panic(err)
	}

	kaccount, err = account.New(cfg.PKey)
	if err != nil {
		panic(err)
	}
}
