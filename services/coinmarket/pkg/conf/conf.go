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
	NoErrorCode int    `mapstructure:"NO_ERROR_CODE"`
	Convert     string `mapstructure:"HTTP_CONVERT_PARAM_NAME"`
	SkipInvalid string `mapstructure:"HTTP_SKIP_INVALID_PARAM_NAME"`
	SymbolParam string `mapstructure:"HTTP_SYMBOL_PARAM_NAME"`
	ApiTimeout  int    `mapstructure:"API_TIMEOUT"`
}

type HttpRequestConf struct {
	ScoreRequestApiAddress        string `mapstructure:"SCORE_REQUEST_API_ADDRESS"`
	ScoreRequestEndPoint          string `mapstructure:"SCORE_REQUEST_ENDPOINT"`
	ScoreRequestCredentials       string `mapstructure:"SCORE_REQUEST_CREDENTIALS"`
	ScoreRequestCredentialsHeader string `mapstructure:"SCORE_REQUEST_CREDENTIALS_HEADER"`
	ScoreRequestConvertVal        string `mapstructure:"SCORE_REQUEST_CONVERT_VAL"`
	ScoreRequestSkipInvalidVal    string `mapstructure:"SCORE_REQUEST_SKIP_INVALID_VAL"`
}

type GRPCServerConfig struct {
	GrpcPort    string `mapstructure:"SCORE_PORT"`
	GrpcNetwork string `mapstructure:"GRPC_NETWORK"`
}
type HttpServerConfig struct {
	HttpPort     string `mapstructure:"SCORE_PORT"`
	HttpUrl      string `mapstructure:"COINMARKET_URL"`
	HttpEndpoint string `mapstructure:"COINMARKET_ENDPOINT"`
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
