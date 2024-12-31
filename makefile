GO_CMD_MAIN = $(wildcard cmd/*.go)

run:
	go run $(GO_CMD_MAIN) server

gen-proto:
	protoc -I ./api \
		-I ../ \
        --go_out ./api --go_opt paths=source_relative \
        --go-grpc_out ./api --go-grpc_opt paths=source_relative \
        --grpc-gateway_out ./api --grpc-gateway_opt paths=source_relative \
        --validate_out=lang=go,paths=source_relative:./api \
        ./api/api.proto

.PHONY: gen-proto, run