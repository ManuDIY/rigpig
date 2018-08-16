package internal

import "fmt"

type Wallet struct {
	Name    string
	Coin    string
	Address string
}

var (
	Wallets []Wallet
)

func testWallets() {
	btcwallet := Wallet{
		Name:    "BTC Mining",
		Coin:    "BTC",
		Address: "ABCH5VYJKvEXCkAayNxURhSXCRnNnAQ8K5",
	}

	pgnWallet := Wallet{
		Name:    "PigeonCoin Mining",
		Coin:    "PGN",
		Address: "PBABCEEBfBFyGjTBrUaYwaTuJJzXT1JKFH",
	}

	fmt.Print(btcwallet, pgnWallet)

}
