package conf

import "github.com/spf13/viper"

const TopApi = "top"
const PageParam = "page"
const LimitParam = "limit"
const TsymParam = "tsym"
const UsdCurrency = "USD"
const SuccessMessage = "Success"
const NoErrorCode = 0
const Convert = "convert"
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

type Config struct {
	TopApi         string
	PageParam      string
	LimitParam     string
	TsymParam      string
	UsdCurrency    string
	SuccessMessage string
	NoErrorCode    int
	Convert        string
	SymbolParam    string
	ApiTimeout     int
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("service")
	v.SetConfigType("json")
	// viper.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		return v, err
	}

	return v, nil
}
