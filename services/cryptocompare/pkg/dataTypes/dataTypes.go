package dataTypes

type TopData struct {
	Symbol string
	Rank   int
}

type TopResponse struct {
	Data []TopResponseData `json:"Data"`
	TopResponseError
}
type TopResponseData struct {
	CoinInfo struct {
		Symbol string `json:"Name"`
	} `json:"CoinInfo"`
}
type TopResponseError struct {
	Message string `json:"Message"`
}
