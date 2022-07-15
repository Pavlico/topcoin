package conf

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

const (
	currentFolderPath = "."
	configFileName    = "app"
	configFileType    = "env"
)

type Config struct {
	AppConfig
	GRPCServerConfig
	HttpServerConfig
	DbCredentials
}

type AppConfig struct {
	SymbolParam              string `mapstructure:"HTTP_SYMBOL_PARAM_NAME"`
	CoinmarketUrl            string `mapstructure:"COINMARKET_URL"`
	CoinmarketScoreEndpoint  string `mapstructure:"COINMARKET_ENDPOINT"`
	CryptocompareUrl         string `mapstructure:"CRYPTOCOMPARE_URL"`
	CryptocompareTopEndpoint string `mapstructure:"CRYPTOCOMPARE_ENDPOINT"`
}

type DbCredentials struct {
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
	Hostname string `mapstructure:"HOSTNAME"`
	DbName   string `mapstructure:"DBNAME"`
}
type GRPCServerConfig struct {
	GrpcPort    string `mapstructure:"SCORE_PORT"`
	GrpcNetwork string `mapstructure:"GRPC_NETWORK"`
}
type HttpServerConfig struct {
	HttpPort     string `mapstructure:"TOP_COIN_PORT"`
	HttpUrl      string `mapstructure:"TOP_COIN_URL"`
	HttpEndpoint string `mapstructure:"TOP_COIN_ENDPOINT"`
}

var ServiceConfig Config

func LoadConfig(config ...interface{}) error {
	if len(config) == 0 {
		return errors.New("Config was nas specified")
	}
	viper.AddConfigPath(currentFolderPath)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileType)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	var grpcServerConfig GRPCServerConfig
	var appConfig AppConfig
	var httpServerConfig HttpServerConfig
	var dbCredentials DbCredentials
	for _, conf := range config {
		switch conf.(type) {
		case GRPCServerConfig:
			unmarhsallErr := viper.Unmarshal(&grpcServerConfig)
			if unmarhsallErr != nil {
				err = fmt.Errorf(err.Error(), unmarhsallErr)
			}
			ServiceConfig = Config{GRPCServerConfig: grpcServerConfig}
		case AppConfig:
			unmarhsallErr := viper.Unmarshal(&appConfig)
			if unmarhsallErr != nil {
				err = fmt.Errorf(err.Error(), unmarhsallErr)
			}
			ServiceConfig = Config{AppConfig: appConfig}
		case HttpServerConfig:
			unmarhsallErr := viper.Unmarshal(&httpServerConfig)
			if unmarhsallErr != nil {
				err = fmt.Errorf(err.Error(), unmarhsallErr)
			}
			ServiceConfig = Config{HttpServerConfig: httpServerConfig}
		case DbCredentials:
			unmarhsallErr := viper.Unmarshal(&dbCredentials)
			if unmarhsallErr != nil {
				err = fmt.Errorf(err.Error(), unmarhsallErr)
			}
			ServiceConfig = Config{DbCredentials: dbCredentials}
		}

	}
	if err != nil {
		return err
	}
	return nil
}
