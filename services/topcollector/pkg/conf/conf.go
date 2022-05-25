package conf

const TopApi = "top"
const ScoreApi = "score"
const PageParam = "page"
const LimitParam = "limit"
const TsymParam = "tsym"
const UsdCurrency = "USD"
const SuccessMessage = "Success"
const NoErrorCode = 0
const Aux = "aux"
const ApiTimeout = 2

var ApiConfig = map[string]ApiData{
	TopApi: {
		ApiAddress:        "https://min-api.cryptocompare.com",
		EndPoint:          "/data/top/totalvolfull?",
		Credentials:       "93d94cc6b38c0c4ed61c744e74d9ecd14223f427b14aea9e7f68f9ca72a5ab7",
		CredentialsHeader: "x-api-key",
		Options: map[string]string{
			PageParam:  "2",
			LimitParam: "100",
			TsymParam:  UsdCurrency,
		},
	},
	ScoreApi: {
		ApiAddress:        "https://pro-api.coinmarketcap.com",
		EndPoint:          "/v1/cryptocurrency/quotes/latest?",
		Credentials:       "18f09716-386e-440a-8906-551d607d1574",
		CredentialsHeader: "X-CMC_PRO_API_KEY",
		Options: map[string]string{
			Aux: "",
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
