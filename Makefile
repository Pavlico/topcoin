build-image:
	cd ./deployments && docker-compose build

run-app:
	cd ./deployments && docker-compose up

run-app-bg:
	cd ./deployments && docker-compose up -d

bash:
	cd ./deployments && docker-compose exec app bash

lint:
	docker run --rm -v `pwd`:/app -w /app golangci/golangci-lint golangci-lint run -v

run-tests:
	cd ./deployments && docker-compose exec app go test ./...

run-tests-coverage:
	cd ./deployments && docker-compose exec app go test ./... -coverprofile tmp/coverage.out
