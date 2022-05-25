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

var ApiConfig = map[string]ApiData{
	TopApi: {
		ApiAddress:        "https://min-api.cryptocompare.com",
		EndPoint:          "/data/top/totalvolfull?",
		Credentials:       "93d94cc6b38c0c4ed61c744e74d9ecd14223f427b14aea9e7f68f9ca72a5ab7",
		CredentialsHeader: "x-api-key",
		PageParam:         "2",
		Options: map[string]string{
			PageParam:  "2",
			LimitParam: "100",
			TsymParam:  UsdCurrency,
		},
	},
}

type ApiData struct {
	ApiAddress        string
	EndPoint          string
	Credentials       string
	CredentialsHeader string
	PageParam         string
	Options           map[string]string
}
