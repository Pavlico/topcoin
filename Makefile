build-image:
	cd ./deployments && docker-compose build

run-app:
	cd ./deployments && docker-compose up

run-app-bg:
	cd ./deployments && docker-compose up -d

bash:
	cd ./deployments && docker-compose exec app bash

bash-grpc:
	cd ./deployments && docker-compose exec grpcapp bash

lint:
	docker run --rm -v `pwd`:/app -w /app golangci/golangci-lint golangci-lint run -v

run-tests:
	cd ./deployments && docker-compose exec app go test ./...

run-tests-coverage:
	cd ./deployments && docker-compose exec app go test ./... -coverprofile tmp/coverage.out

.PHONY: protos

protos:
	 protoc -I internal/protos/ internal/protos/coins.proto --go_out=plugins=grpc:internal/protos/coins

protos-cryptocompare:
	 protoc -I services/cryptocompare/pkg/protos/ services/cryptocompare/pkg/protos/cryptocompare.proto --go_out=plugins=grpc:services/cryptocompare/pkg/grpc/protos/cryptocompare

protos-coinmarket:
	 protoc -I services/coinmarket/pkg/protos/ services/coinmarket/pkg/protos/coinmarket.proto --go_out=plugins=grpc:services/coinmarket/pkg/grpc/protos/coinmarket

protos-topcollector:
	 protoc -I services/topcollector/pkg/grpc/protos/ services/topcollector/pkg/grpc/protos/topcollector.proto --go_out=plugins=grpc:services/topcollector/pkg/grpc/protos/topcollector