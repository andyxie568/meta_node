package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/9e696c6371784cef89636f7aca01db3d")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	account := common.HexToAddress("0xdCee1d1D79461c7f87b1C19bfacce0e986038451")
	balance, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(balance)

	blackNumber := big.NewInt(8894050)
	balanceAt, err := client.BalanceAt(ctx, account, blackNumber)
	if err != nil {
		panic(err)
	}
	fmt.Println(balanceAt)

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)
	pendingBalance, err := client.PendingBalanceAt(ctx, account)
	if err != nil {
		panic(err)
	}
	fmt.Println(pendingBalance)
}
