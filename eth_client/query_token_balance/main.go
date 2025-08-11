package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"myproject/meta_node/eth_client/query_token_balance/token"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/9e696c6371784cef89636f7aca01db3d")
	if err != nil {
	}

	tokenAddress := common.HexToAddress("0x9b765d761a18fe739d72b0283a85cfcf7516e5b7")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0xdCee1d1D79461c7f87b1C19bfacce0e986038451")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	//name, err := instance.Name(&bind.CallOpts{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//symbol, err := instance.Symbol(&bind.CallOpts{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//decimals, err := instance.Decimals(&bind.CallOpts{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	//fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	//fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	//value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	//fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
