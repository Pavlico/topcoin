package conf

type ApiData struct {
	ApiAddress   string
	Subdirectory string
	Credentials  string
	Options      map[string]string
}

const UsdCurrency = "USD"
const CryptocompareCredentials = "93d94cc6b38c0c4ed61c744e74d9ecd14223f427b14aea9e7f68f9ca72a5ab78"
const CoinmarketcapCredentials = "18f09716-386e-440a-8906-551d607d1574"
const CryptocompareApiAddress = "https://min-api.cryptocompare.com"
const CoinmarketcapAddress = "https://pro-api.coinmarketcap.com"
const TopSubDirectory = "/data/top/totalvolfull?"
const ScoreSubDirectory = "/v1/cryptocurrency/listings/latest?"
const TopApiFlag = "top"
const ScoreApiFlag = "score"
const AmountOfPages = 2
const PageLimit = "100"
const FirstPage = "0"
const SuccessMessage = "Success"
const NoErrorCode = 0
