package main

import (
	"context"
	"fmt"
	"github.com/zkscpqm/atom-finance-go"
	"github.com/zkscpqm/atom-finance-go/pkg/market"
	"os"
)

func main() {
	fmt.Println("parsing config...")
	cfg, err := atom.NewConfig(".test/config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("done!")

	fmt.Println("getting initializing client...")
	client, err := atom.NewClient(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer client.Close()

	fmt.Println("done!")

	for _, ticker := range []string{"NKE"} {
		fmt.Println("getting analyst estimate for", ticker)
		resp, err := client.AnalystEstimates(context.Background(), ticker, market.USA)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(resp.Json())
	}

	fmt.Println("done!")
}
