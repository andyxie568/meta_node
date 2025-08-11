package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/nEi1vhiza6JeDEdreutlg")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		panic(err)
	}

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		panic(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())
		fmt.Println(tx.Value().String())
		fmt.Println(tx.Gas())
		fmt.Println(tx.GasPrice().Uint64())
		fmt.Println(tx.Nonce())
		fmt.Println(tx.Data())
		fmt.Println(tx.To().Hex())

		sender, err := types.Sender(types.NewEIP155Signer(chainID), tx)
		if err != nil {
			panic(err)
		}
		fmt.Println("sender", sender.Hex())

		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			panic(err)
		}

		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		break
	}

	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(ctx, blockHash)
	if err != nil {
		panic(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(ctx, blockHash, idx)
		if err != nil {
			panic(err)
		}
		fmt.Println(tx.Hash().Hex())
		break
	}

	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	tx, isPending, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		panic(err)
	}

	fmt.Println(isPending)
	fmt.Println(tx.Hash().Hex())
}
