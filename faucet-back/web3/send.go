package web3

import (
	"errors"
	"log"
	"math/big"

	"faucet-app/web3/account"
	"faucet-app/web3/client"
)

func SendETH(to_hex string, amount_str string) (txhash *string, err error) {
	to, err := account.CommonAddress(to_hex)
	amount := big.NewInt(0)
	amount, success := amount.SetString(amount_str, 10)
	if err != nil || !success {
		return nil, errors.New("parameter format error")
	}

	tx, err := client.SendTransaction(kaccount, to, amount, nil, log.Default())
	if tx != nil {
		hash := tx.Hash().Hex()
		txhash = &hash
	}
	return txhash, err
}

func SendErc20Token(to_hex string, amount_str string) (txhash *string, err error) {
	to, err := account.CommonAddress(to_hex)
	amount := big.NewInt(0)
	amount, success := amount.SetString(amount_str, 10)
	if err != nil || !success {
		return nil, errors.New("parameter format error")
	}

	contractAddress, _ := account.CommonAddress(erc20Address)
	data, _ := contractAbi.Pack("transfer", to, amount)

	tx, err := client.SendTransaction(kaccount, contractAddress, nil, data, log.Default())
	if tx != nil {
		hash := tx.Hash().Hex()
		txhash = &hash
	}
	return txhash, err
}
