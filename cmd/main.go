package main

import (
	"fmt"
	"time"

	"github.com/deforceHK/goghostex/okex"
)

func main() {
	fmt.Println("Hello, World!")

	var wsOK = okex.WSMarketOKEx{
		&okex.WSTradeOKEx{
			RecvHandler: func(msg string) {
				fmt.Println(msg)
			},
		},
	}

	wsOK.Start()

	var op = okex.WSOpOKEx{
		Op: "subscribe",
		Args: []map[string]string{
			{
				"channel": "books5",
				"instId":  "BTC-USDT-SWAP",
			},
			//{
			//	"channel": "books5",
			//	"instId":  "ETH-USDT-SWAP",
			//},
		},
	}
	wsOK.Subscribe(op)
	time.Sleep(5 * time.Second)
	select {}
}
