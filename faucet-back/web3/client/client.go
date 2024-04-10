package client

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"faucet-app/web3/account"
)

var visitAllow = true
var visitLock sync.Mutex

func unlock() {
	visitAllow = true
	visitLock.Unlock()
}

type logger interface {
	Println(a ...interface{})
	Printf(format string, a ...interface{})
}

type KLClient struct {
	Ethclient *ethclient.Client
	ChainID   *big.Int
}

func New(rpc string) (*KLClient, error) {
	client_, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	chainID_, err := client_.NetworkID(timeoutCtx())
	if err != nil {
		return nil, err
	}

	return &KLClient{
		Ethclient: client_,
		ChainID:   chainID_,
	}, nil
}

func (client *KLClient) Nonce(account common.Address) (uint64, error) {
	return client.Ethclient.PendingNonceAt(timeoutCtx(), account)
}

func (client *KLClient) SuggestGasPrice() (*big.Int, error) {
	return client.Ethclient.SuggestGasPrice(timeoutCtx())
}

func (client *KLClient) EstimateGas(account common.Address, to string, callData []byte) (uint64, error) {
	contractAddress := common.HexToAddress(to)
	return client.Ethclient.EstimateGas(timeoutCtx(), ethereum.CallMsg{
		From: account,
		To:   &contractAddress,
		Data: callData,
	})
}

func (client *KLClient) SendTransaction(klAccount *account.KLAccount, to common.Address, amount *big.Int, data []byte, log logger) (*types.Transaction, error) {

	if !visitAllow {
		return nil, errors.New("client is busy")
	}
	visitLock.Lock()
	visitAllow = false
	defer unlock()

	log.Println("===============================")
	nonce, err := client.Nonce(klAccount.AddressCommon)
	if err != nil {
		log.Println("get nonce error:", err.Error())
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice()
	if err != nil {
		gasPrice = big.NewInt(30000000000) // in wei (30 gwei)
	}

	gasLimit, err := client.EstimateGas(klAccount.AddressCommon, to.Hex(), data)
	if err != nil {
		log.Println("get gaslimit error:", err.Error())
		return nil, err
	}
	log.Printf("nonce:%d to:%v amount:%v gasLimit:%d gasPrice:%v\n", nonce, to, amount, gasLimit, gasPrice)

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(client.ChainID), klAccount.PrivateKey)
	if err != nil {
		log.Println("sign tx error:", err.Error())
		return nil, err
	}

	err = client.Ethclient.SendTransaction(timeoutCtx(), signedTx)
	if err != nil {
		log.Println("send transaction error:", err.Error())
		return signedTx, err
	}
	log.Println("tx broadcast:", signedTx.Hash().Hex())

	receipt, err := bind.WaitMined(timeoutCtx(), client.Ethclient, signedTx)
	if err != nil {
		log.Println("wait mined error:", err)
		if err.Error() == "context deadline exceeded" {
			return signedTx, err
		} else {
			return signedTx, err
		}
	} else {
		log.Printf("receipted - status:%d, blockNumber:%s\n", receipt.Status, receipt.BlockNumber.String())
		if receipt.Status == 1 {
			return signedTx, nil
		} else {
			return signedTx, fmt.Errorf("receipted fail, status: %d", receipt.Status)
		}
	}
}

func timeoutCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	return ctx
}
