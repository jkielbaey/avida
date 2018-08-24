package avida

import (
	"errors"

	"github.com/pdepip/go-binance/binance"
)

// Exchange represents a single crypto exchange.
type Exchange struct {
	Exchange  string
	APIKey    string
	APISecret string
}

// GetPositions returns all positions on a given exchange.
func (e *Exchange) GetPositions() (*[]Position, error) {
	if e.Exchange == "binance" {
		return getBinancePositions(e.APIKey, e.APISecret)
	}
	return nil, errors.New("unsupported exchange")
}

func getBinancePositions(apiKey string, apiSecret string) (*[]Position, error) {
	client := binance.New(apiKey, apiSecret)

	binancePositions, err := client.GetPositions()
	if err != nil {
		return nil, err
	}

	var positions []Position
	for _, pos := range binancePositions {
		p := NewPosition(pos.Asset, pos.Free+pos.Locked)
		positions = append(positions, *p)
	}
	return &positions, nil
}
