package conf

const SymbolParam = "symbols"
const CoinmarketUrl = "http://coinmarket:8060"
const CoinmarketScoreEndpoint = "/score?"
const CryptocompareUrl = "http://cryptocompare:8050"
const CryptocompareTopEndpoint = "/top100"
const TopCoinEndpoint = "/topcoins"
const TopCoinPort = ":8070"

type dbCredentials struct {
	Username string
	Password string
	Hostname string
	DbName   string
}

func GetDbCredentials() dbCredentials {
	return dbCredentials{
		Username: "admin",
		Password: "password123",
		Hostname: "db:3306",
		DbName:   "coins",
	}
}
