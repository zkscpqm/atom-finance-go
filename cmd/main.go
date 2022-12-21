package main

import (
	"fmt"
	"github.com/zkscpqm/atom-finance-go"
	"github.com/zkscpqm/atom-finance-go/market"
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

	fmt.Println("getting analyst estimate for AAPL...")
	err = client.DEBUGAnalystEstimates("AAPL", market.USA)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("done!")
}
