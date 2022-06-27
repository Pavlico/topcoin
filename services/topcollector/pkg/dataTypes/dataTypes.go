package dataTypes

type CoinData struct {
	Symbol string  `json:"Symbol"`
	Rank   int     `json:"Rank"`
	Score  float32 `json:"Score"`
}

type TopData struct {
	Symbol string `json:"Symbol"`
	Rank   int    `json:"Rank"`
}

type ScoreData struct {
	Symbol string  `json:"Symbol"`
	Score  float32 `json:"Score"`
}
