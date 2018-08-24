package avida

import (
	"strings"
)

// Position presents a single position in either a wallet or an exchange.
type Position struct {
	Symbol string
	Amount float64
}

// NewPosition creates a new Position.
func NewPosition(symbol string, amount float64) *Position {
	p := new(Position)
	p.Symbol = symbol
	p.Amount = amount
	return p
}

// GetValueUSD will calculate the value of a ColdAsset based on the current
// price of the coin on CoinMarketCap.
func (p *Position) GetValueUSD() float64 {
	c := coinMap.Coins[strings.ToLower(p.Symbol)]
	return p.Amount * float64(c.PriceUSD)
}
