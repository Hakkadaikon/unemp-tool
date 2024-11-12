package exchange

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bitcoin struct {
	oneBtcToJpy float64
}

func (this Bitcoin) OneBtcToJpy() (float64, error) {
	if this.oneBtcToJpy != 0 {
		return this.oneBtcToJpy, nil
	}

	type CoinGeckoResponse struct {
		Bitcoin struct {
			Jpy float64 `json:"jpy"`
		} `json:"bitcoin"`
	}

	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=jpy"
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	this.oneBtcToJpy = result.Bitcoin.Jpy
	return result.Bitcoin.Jpy, nil
}

func (this Bitcoin) SatoshiToBtc(satoshi float64) float64 {
	return satoshi / 100000000
}

func (this Bitcoin) BtcToSatoshi(satoshi float64) float64 {
	return satoshi * 100000000
}

func (this Bitcoin) JpyToBtc(jpy float64) (float64, error) {
	btcToJpy, err := this.OneBtcToJpy()
	if err != nil {
		return 0, err
	}

	jpyToBtc := jpy / btcToJpy
	return jpyToBtc, nil
}

func (this Bitcoin) JpyToSatoshi(jpy float64) (float64, error) {
	if jpy <= 0 {
		return 0, fmt.Errorf("jpy is 0")
	}

	jpyToBtc, err := this.JpyToBtc(jpy)
	if err != nil {
		return 0, err
	}

	if jpyToBtc <= 0 {
		return 0, fmt.Errorf("jpy to btc is 0")
	}

	jpyToSatoshi := this.BtcToSatoshi(jpyToBtc)
	return jpyToSatoshi, nil
}
