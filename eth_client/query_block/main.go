package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/9e696c6371784cef89636f7aca01db3d")
	if err != nil {
		panic(err)
	}

	blockNumber := big.NewInt(5671744)
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		panic(err)
	}
	fmt.Println("Header number:", header.Number.Int64())
	fmt.Println("Header Time:", header.Time)
	fmt.Println("Header Difficulty:", header.Difficulty.Uint64())
	fmt.Println("Header Hash:", header.Hash().Hex())

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		panic(err)
	}
	fmt.Println("Block number:", block.Number().Uint64())
	fmt.Println("Block Time:", block.Time())
	fmt.Println("Block Difficulty:", block.Difficulty().Uint64())
	fmt.Println("Block Hash:", block.Hash().Hex())
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction count:", count)
}
