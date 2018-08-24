package avida

import (
	"fmt"
	"strings"

	cmc "github.com/miguelmota/go-coinmarketcap"
)

var coinMap *CoinMap

// CoinMap presents a map with coin info from Coinmarketcap.
type CoinMap struct {
	Coins map[string]cmc.Coin
}

// NewCoinMap creates a new CoinMap and populate it with all currencies from coinmarketcap.
func NewCoinMap() (*CoinMap, error) {
	cm := new(CoinMap)
	cm.Coins = make(map[string]cmc.Coin)

	coins, err := cmc.GetAllCoinData(100)
	if err != nil {
		return nil, err
	}
	for _, coin := range coins {
		cm.Coins[strings.ToLower(coin.Symbol)] = coin
		cm.Coins[strings.ToLower(coin.Name)] = coin
	}

	return cm, nil
}

func init() {
	var err error
	coinMap, err = NewCoinMap()
	if err != nil {
		fmt.Println(err)
	}
}
