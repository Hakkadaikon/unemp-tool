package exchange

import (
	"encoding/json"
)

type CoinGecko struct{}

func (this *CoinGecko) GetEndPoint() string {
	return "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=jpy"
}

func (this *CoinGecko) BtcToJpy(raw []byte) (float64, error) {
	type Response struct {
		Bitcoin struct {
			Jpy float64 `json:"jpy"`
		} `json:"bitcoin"`
	}

	var result Response
	if err := json.Unmarshal(raw, &result); err != nil {
		return 0, err
	}

	return result.Bitcoin.Jpy, nil
}
