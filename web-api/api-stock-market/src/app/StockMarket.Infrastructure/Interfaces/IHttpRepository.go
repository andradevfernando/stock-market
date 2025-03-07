package Interface

type IHttpRepository interface {
	FetchStocks(urlRequest string) ([]byte, error)
}
