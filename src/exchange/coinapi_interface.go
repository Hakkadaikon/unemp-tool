package exchange

type CoinApiInterface interface {
	GetEndPoint() string
	BtcToJpy(raw []byte) (float64, error)
}
