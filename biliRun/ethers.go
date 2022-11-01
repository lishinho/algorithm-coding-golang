package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// 测试通过eth客户端拿pending txn

const (
	AlchemyMainnetWssurl = "wss://eth-mainnet.g.alchemy.com/v2/oKmOQKbneVkxgHZfibs-iFhIlIAl6HDN"
)

func dial() {
	// golang eth for pending txn
	client, err := ethclient.Dial(AlchemyMainnetWssurl)
	if err != nil {
		log.Fatal(err)
	}

	r, err := rpc.DialContext(context.Background(), AlchemyMainnetWssurl)
	if err != nil {
		log.Fatal(err)
	}
	ec := gethclient.New(r)

	ch := make(chan common.Hash)
	sub, err := ec.SubscribePendingTransactions(context.Background(), ch)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case hash := <-ch:
			tx, isPending, err := client.TransactionByHash(context.Background(), hash)
			log.Println("pending: ", isPending, "hash: ", hash)
			if isPending && err == nil {
				log.Printf("%+v\n", common.Bytes2Hex(tx.Data()))
			}
		}
	}
}
