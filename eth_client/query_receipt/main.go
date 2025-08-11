package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"time"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/9e696c6371784cef89636f7aca01db3d")
	if err != nil {
		panic(err)
	}

	blockNumber := big.NewInt(5671744)
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")

	ctx := context.Background()
	receiptByHash, err := client.BlockReceipts(ctx, rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Minute) // 连续请求会报请求频繁的错误，这里停1分钟
	receiptByNum, err := client.BlockReceipts(ctx, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		panic(err)
	}
	fmt.Println(receiptByHash[0].TxHash.Hex() == receiptByNum[0].TxHash.Hex())

	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		fmt.Println(receipt.TxHash.Hex())
		fmt.Println(receipt.TransactionIndex)
		fmt.Println(receipt.ContractAddress.Hex())
		break
	}

	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		panic(err)
	}
	fmt.Println(receipt.Status)
	fmt.Println(receipt.Logs)
	fmt.Println(receipt.TxHash.Hex())
	fmt.Println(receipt.TransactionIndex)
	fmt.Println(receipt.ContractAddress.Hex())
}
