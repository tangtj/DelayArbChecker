package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

var SequencerUrl = "https://arb1-sequencer.arbitrum.io/rpc"

var url string

var count int

func main() {

	parseArgs()

	fmt.Printf("[ url ]  %s\n", url)

	ctx, cal := context.WithCancel(context.Background())
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		fmt.Printf("dial url : %s  error : %s", url, err)
	}

	go func() {
		// 3s 调用一次 chain id 方法

		t := time.NewTicker(2 * time.Second)
		defer t.Stop()

		c := 0

		for {
			select {
			case <-ctx.Done():
				fmt.Printf("ticker done\n")
				return
			case <-t.C:
				// 计量开始和结束时间
				start := time.Now().UnixMilli()
				_, _ = client.ChainID(context.Background())
				end := time.Now().UnixMilli()

				c++

				fmt.Printf("[ %d ] %d ms\n", c, end-start)
				if c > count {
					break
				}
			}
			cal()
		}
	}()

	// 阻塞主线程
	<-ctx.Done()

}

func parseArgs() {
	flag.StringVar(&url, "url", SequencerUrl, "rpc url")
	flag.IntVar(&count, "count", 10, "count")
	flag.Parse()
}
