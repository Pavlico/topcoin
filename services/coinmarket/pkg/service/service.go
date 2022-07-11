package service

import (
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/conf"
	"github.com/Pavlico/topcoin/services/coinmarket/pkg/flags"
	grpcserver "github.com/Pavlico/topcoin/services/coinmarket/pkg/grpc/server"
	httpserver "github.com/Pavlico/topcoin/services/coinmarket/pkg/http/server"
)

const (
	grpcType    = "grpc"
	httpType    = "http"
	defaultType = httpType
)

func Run() error {
	flagData := flags.GetFlagData()
	if err := flagData.ValidateFlags(); err != nil {
		return err
	}

	if flagData.CommunicationType == defaultType {
		if err := conf.LoadConfig(conf.AppConfig{}, conf.HttpRequestConf{}, conf.HttpServerConfig{}); err != nil {
			return err
		}
		if err := httpserver.Serve(); err != nil {
			return err
		}
	}
	if flagData.CommunicationType == grpcType {
		if err := conf.LoadConfig(conf.AppConfig{}, conf.HttpRequestConf{}, conf.GRPCServerConfig{}); err != nil {
			return err
		}
		if err := grpcserver.Serve(); err != nil {

		}
	}
	return nil
}
