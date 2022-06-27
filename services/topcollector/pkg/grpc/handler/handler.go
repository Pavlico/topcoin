package handler

import (
	"context"

	protos "github.com/Pavlico/topcoin/services/topcollector/pkg/grpc/protos/topcollector"
	"github.com/Pavlico/topcoin/services/topcollector/pkg/grpc/service"
)

type TopCollector struct {
}

func NewCoinList() *TopCollector {
	return &TopCollector{}
}
func (c *TopCollector) GetMergedData(ctx context.Context, cr *protos.TopCoinRequest) (*protos.TopCoinResponse, error) {
	mergedData, err := service.GetMergedData()
	if err != nil {
		return nil, err
	}
	return &protos.TopCoinResponse{Coin: mergedData}, nil
}
