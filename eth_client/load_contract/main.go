package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"myproject/meta_node/eth_client/deploy_contract/Store"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		panic(err)
	}

	storeContract, err := Store.NewStore(common.HexToAddress("0xacb36ecff2b55c01340d70378a8c0d3074a17981"), client)
	if err != nil {
		panic(err)
	}
	fmt.Println(json.Marshal(storeContract))
}
