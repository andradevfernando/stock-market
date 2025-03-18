package Http

import (
	"api-stock-market/src/app/StockMarket.Domain/Models"
	"api-stock-market/src/app/StockMarket.Infrastructure/Repository/Http/Request"
	"api-stock-market/src/app/StockMarket.Infrastructure/Repository/Http/Response"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type OpenAiHttpRepository struct {
	apiKey        string
	openAiBaseUrl string
	model         string
	httpClient    *http.Client
}

func NewOpenAiHttpRepository(apiKey, model, baseUrl string) *OpenAiHttpRepository {
	return &OpenAiHttpRepository{
		apiKey:        apiKey,
		model:         model,
		openAiBaseUrl: baseUrl,
		httpClient:    &http.Client{},
	}
}
func (c OpenAiHttpRepository) GetInvestmentRecommendation(content Request.Content) (recommendation Models.InvestmentRecommendation, err error) {
	prompt := fmt.Sprintf(`Given the following data, 
should I invest in this stock or not? 
return the data in the following json format and only the json so i can use it in my vue web app:

{
"shouldInvest": true,
"reason": "Because of the high PE ratio",
"recommendation": "Buy this stock",
"comment": "This is a good stock",
"score": 100,
"ticker": "AAPL",
"pe_ratio": 10,
"pb_ratio": 10,
"dividend_yield": 10,
"price": 10,
"debt_to_equity": 10,
"market_cap": 10
}
Ticker: %s, Last year moving average: %.2f, PE Ratio: %.2f, PB Ratio: %.2f, Dividend Yield: %.2f, Price: %.2f, Debt to equity: %.2f, Market cap: %.2f`,
		content.Ticker, content.MovingAverage, content.PERatio, content.PBRatio, content.DividendYield, content.Price, content.DebtToEquity, content.MarketCap)

	requestBody := Request.OpenAIRequest{
		Model: c.model,
		Messages: []Request.Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return recommendation, fmt.Errorf("error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", c.openAiBaseUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return recommendation, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return recommendation, fmt.Errorf("error doing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return recommendation, fmt.Errorf("openai API returned status: %d", resp.StatusCode)
	}

	var response Response.OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return recommendation, fmt.Errorf("error decoding response: %v", err)
	}

	err = json.Unmarshal([]byte(cleanJSONString(response.Choices[0].Message.Content)), &recommendation)
	if err != nil {
		fmt.Println("error converting JSON:", err)
		return recommendation, fmt.Errorf("error decoding recommendation: %v", err)
	}
	if len(response.Choices) == 0 {
		return recommendation, fmt.Errorf("no choices in response")
	}

	return recommendation, nil
}
func cleanJSONString(input string) string {
	result := strings.TrimSpace(input)
	if strings.HasPrefix(result, "```json") {
		result = strings.TrimPrefix(result, "```json")
	}
	if strings.HasSuffix(result, "```") {
		result = strings.TrimSuffix(result, "```")
	}
	return strings.TrimSpace(result)
}
