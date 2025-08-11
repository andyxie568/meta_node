package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/9e696c6371784cef89636f7aca01db3d")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("SEPOLIA_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	ctx := context.Background()
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		panic(err)
	}

	value := big.NewInt(10000000000000)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		panic(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		panic(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		panic(err)
	}
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
}
