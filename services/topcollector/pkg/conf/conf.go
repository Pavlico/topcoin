package conf

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
		Hostname: "appDb:8094",
		DbName:   "coins",
	}
}
