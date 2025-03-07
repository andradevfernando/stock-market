package Repository

import (
	"fmt"
	"io"
	"net/http"
)

type HttpRepository struct {
}

func NewHttpRepository() *HttpRepository {
	return &HttpRepository{}
}

func (h *HttpRepository) FetchStocks(urlRequest string) ([]byte, error) {
	req, err := http.NewRequest("GET", urlRequest, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer ")
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err == nil {
		body, _ := io.ReadAll(resp.Body)
		return body, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}
	return nil, nil
}
