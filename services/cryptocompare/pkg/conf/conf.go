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
	HttpRequestConf
	GRPCServerConfig
	HttpServerConfig
}

type AppConfig struct {
	PageParam      string `mapstructure:"HTTP_PAGE_PARAM_NAME"`
	LimitParam     string `mapstructure:"HTTP_LIMIT_PARAM_NAME"`
	TsymParam      string `mapstructure:"HTTP_TSYM_PARAM_NAME"`
	SuccessMessage string `mapstructure:"SUCCESS_MESSAGE"`
	SkipInvalid    string `mapstructure:"HTTP_SKIP_INVALID_PARAM_NAME"`
	SymbolParam    string `mapstructure:"HTTP_SYMBOL_PARAM_NAME"`
	ApiTimeout     int    `mapstructure:"API_TIMEOUT"`
}

type HttpRequestConf struct {
	RankRequestApiAddress        string `mapstructure:"RANK_REQUEST_API_ADDRESS"`
	RankRequestEndPoint          string `mapstructure:"RANK_REQUEST_ENDPOINT"`
	RankRequestCredentials       string `mapstructure:"RANK_REQUEST_CREDENTIALS"`
	RankRequestCredentialsHeader string `mapstructure:"RANK_REQUEST_CREDENTIALS_HEADER"`
	RankRequestPageParam         string `mapstructure:"RANK_REQUEST_PAGE_PARAM"`
	RankRequestTsymVal           string `mapstructure:"RANK_REQUEST_SKIP_TSYM_PARAM"`
	RankRequestLimitVal          string `mapstructure:"RANK_REQUEST_LIMIT_VAL"`
}

type GRPCServerConfig struct {
	GrpcPort    string `mapstructure:"RANK_PORT"`
	GrpcNetwork string `mapstructure:"GRPC_NETWORK"`
}

type HttpServerConfig struct {
	HttpPort     string `mapstructure:"RANK_PORT"`
	HttpUrl      string `mapstructure:"CRYPTOCOMPARE_URL"`
	HttpEndpoint string `mapstructure:"CRYPTOCOMPARE_ENDPOINT"`
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
	var httpRequestConfig HttpRequestConf
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
		case HttpRequestConf:
			unmarhsallErr := viper.Unmarshal(&httpRequestConfig)
			if unmarhsallErr != nil {
				err = fmt.Errorf(err.Error(), unmarhsallErr)
			}
			ServiceConfig = Config{HttpRequestConf: httpRequestConfig}
		}
	}
	if err != nil {
		return err
	}
	return nil
}
