package Interface

type ITruoraStocksHttpRepository interface {
	FetchStocks(nextPage string) ([]byte, error)
}
