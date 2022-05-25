package dataTypes

type ScoreData struct {
	Symbol string
	Score  float32
}

type ScoreResponse struct {
	Data map[string]ScoreResponseData `json:"data"`
	ScoreResponseError
}

type ScoreResponseData struct {
	Symbol string `json:"symbol"`
	Score  struct {
		Currency struct {
			Price float32 `json:"price"`
		} `json:"USD"`
	} `json:"quote"`
}

type ScoreResponseError struct {
	Status struct {
		ErrorCode int `json:"error_code"`
	} `json:"status"`
}
