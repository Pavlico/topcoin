package conf

const TopApi = "top"
const ScoreApi = "score"
const PageParam = "page"
const LimitParam = "limit"
const TsymParam = "tsym"
const UsdCurrency = "USD"
const SuccessMessage = "Success"
const NoErrorCode = 0
const Convert = "convert"
const SkipInvalid = "skip_invalid"
const SymbolParam = "symbol"
const ApiTimeout = 5
const EmptyValue = ""
const CoinmarketUrl = "http://coinmarket:8060"
const TopCollectorUrl = "http://topcollector:8070"
const TopCollectorEndpoint = "/topcoins"
const CoinmarketScoreEndpoint = "/score?"
const CryptocompareUrl = "http://cryptocompare:8050"
const CryptocompareTopEndpoint = "/top100"
const EndPointTop = "/topcoins"
const TopCoinPort = ":8070"
const ScorePort = ":8060"
const AppPort = ":8080"
const TopPort = "8050"

var ApiConfig = map[string]ApiData{
	ScoreApi: {
		ApiAddress:        "https://pro-api.coinmarketcap.com",
		EndPoint:          "/v1/cryptocurrency/quotes/latest?",
		Credentials:       "18f09716-386e-440a-8906-551d607d1574",
		CredentialsHeader: "X-CMC_PRO_API_KEY",
		Options: map[string]string{
			Convert:     "USD",
			SkipInvalid: "true",
		},
	},
}

type ApiData struct {
	ApiAddress        string
	EndPoint          string
	Credentials       string
	CredentialsHeader string
	Options           map[string]string
}
