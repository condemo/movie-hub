binary-name=movie-hub
data-service=movie-data

build:
	@GOOS=windows GOARCH=amd64 go build -o ./bin/${binary-name}-win.exe ./cmd/main.go
	@GOOS=linux GOARCH=amd64 go build -o ./bin/${binary-name}-linux ./cmd/main.go
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/${binary-name}-darwin ./cmd/main.go

run: build
	@./bin/${binary-name}-linux

arm-build:
	@GOOS=linux GOARCH=arm64 go build -o ./bin/${binary-name}-arm64 ./services/rest/cmd/main.go
	@GOOS=linux GOARCH=arm64 go build -o ./bin/${data-service}-arm64 ./services/data_handler/cmd/main.go

arm-run: arm-build
	@./bin/${binary-name}-arm64

arm-install: arm-build
	@cp -f ./bin/${binary-name}-arm64 ~/services/movie-hub/${binary-name}
	@cp -f ./bin/${data-service}-arm64 ~/services/movie-hub/${data-service}

protogen:
	@protoc \
		--proto_path=proto "proto/data_handler.proto" \
		--go_out=services/common/protogen/pb --go_opt=paths=source_relative \
		--go-grpc_out=services/common/protogen/pb \
		--go-grpc_opt=paths=source_relative

test:
	@godotenv -f .env go test ./... -v

clean:
	@rm -rf ./bin/*
	@go clean

data-build:
	@GOOS=linux GOARCH=amd64 go build -o ./bin/${data-service}-linux_64 ./services/data_handler/cmd/main.go

data-run: data-build
	@./bin/${data-service}-linux_64


arm-data-build:
	@GOOS=linux GOARCH=arm64 go build -o ./bin/${data-service}-linux_arm64 ./services/data_handler/cmd/main.go

arm-data-run: arm-data-build
	@./bin/${data-service}-linux_arm64 -addr=:6100

kill-services:
	@lsof -t -i:5200 | xargs -r kill
	@lsof -t -i:5300 | xargs -r kill
