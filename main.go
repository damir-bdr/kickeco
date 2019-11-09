package main

import (
	"log"

	"github.com/damir-bdr/kickeco/clicmd"

	"github.com/damir-bdr/kickeco/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	GATEWAY              = "https://mainnet.infura.io/v3/3b3b55afe03647b8a528ba48a71211eb"
	KickCoinContractAddr = "0x27695e09149adc738a978e9a678f99e4c39e9eb9"
)

func main() {
	conn, err := ethclient.Dial(GATEWAY)
	if err != nil {
		log.Fatalf("Could not connect to Ethereum gateway: %v\n", err)
	}
	defer conn.Close()

	accountAddress := common.HexToAddress(KickCoinContractAddr)
	_, err = contract.NewCSToken(accountAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v\n", err)
	}

	clicmd.Execute()
}
