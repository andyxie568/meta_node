package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://eth-mainnet.g.alchemy.com/v2/nEi1vhiza6JeDEdreutlg")
	if err != nil {
		panic(err)
	}

	headers := make(chan *types.Header)
	ctx := context.Background()
	sub, err := client.SubscribeNewHead(ctx, headers)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())
			block, err := client.BlockByHash(ctx, header.Hash())
			if err != nil {
				panic(err)
			}

			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Number().Uint64())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
		}
	}
}
