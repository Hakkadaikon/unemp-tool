package exchange

import (
	"unemp-tool/middleware"
	"unemp-tool/myerror"
)

type Bitcoin struct {
	oneBtcToJpy       float64
	httpClientAdapter middleware.HttpClient
	coinapi           CoinApiInterface
}

func (this *Bitcoin) SetHttpClient(client middleware.HttpClientInterface) {
	this.httpClientAdapter.SetHttpClient(client)
}

func (this *Bitcoin) SetCoinApi(coinapi CoinApiInterface) {
	this.coinapi = coinapi
}

func (this *Bitcoin) OneBtcToJpy() (float64, error) {
	if this.oneBtcToJpy != 0 {
		return this.oneBtcToJpy, nil
	}

	if this.coinapi == nil {
		this.coinapi = &CoinGecko{}
	}

	endpoint := this.coinapi.GetEndPoint()
	resp, err := this.httpClientAdapter.Get(endpoint)
	if err != nil {
		return 0, err
	}

	oneBtcToJpy, err := this.coinapi.BtcToJpy([]byte(resp))
	if err != nil {
		return 0, err
	}

	this.oneBtcToJpy = oneBtcToJpy
	return oneBtcToJpy, nil
}

func (this *Bitcoin) SatoshiToBtc(satoshi float64) float64 {
	return satoshi / 100000000
}

func (this *Bitcoin) BtcToSatoshi(satoshi float64) float64 {
	return satoshi * 100000000
}

func (this *Bitcoin) JpyToBtc(jpy float64) (float64, error) {
	btcToJpy, err := this.OneBtcToJpy()
	if err != nil {
		return 0, err
	}

	jpyToBtc := jpy / btcToJpy
	return jpyToBtc, nil
}

func (this *Bitcoin) JpyToSatoshi(jpy float64) (float64, error) {
	if jpy <= 0 {
		return 0, myerror.New("jpy is 0")
	}

	jpyToBtc, err := this.JpyToBtc(jpy)
	if err != nil {
		return 0, err
	}

	if jpyToBtc <= 0 {
		return 0, myerror.New("jpy to btc is 0")
	}

	jpyToSatoshi := this.BtcToSatoshi(jpyToBtc)
	return jpyToSatoshi, nil
}
