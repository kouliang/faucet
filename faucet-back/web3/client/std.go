package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"faucet-app/web3/account"
)

var std *KLClient

func InitStd(rpc string) (err error) {
	std, err = New(rpc)
	return
}

func ChainID() *big.Int {
	return std.ChainID
}

func Nonce(account common.Address) (uint64, error) {
	return std.Nonce(account)
}

func SuggestGasPrice() (*big.Int, error) {
	return std.SuggestGasPrice()
}

func EstimateGas(account common.Address, to string, callData []byte) (uint64, error) {
	return std.EstimateGas(account, to, callData)
}

func SendTransaction(klAccount *account.KLAccount, to common.Address, amount *big.Int, data []byte, log logger) (*types.Transaction, error) {
	return std.SendTransaction(klAccount, to, amount, data, log)
}
