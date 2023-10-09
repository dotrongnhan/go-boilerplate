## This Makefile saves some typing and groups some common commands

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

gen-rpc: ## Generate RPC Code
	cd app/$(SERVICE)/rpc/pb && \
	goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero && \
	sed -i '' 's/,omitempty//g' *.pb.go

gen-api: ## Generate API
	cd cmd/api/$(SERVICE)/desc && \
	goctl api go -api *.api -dir ../ --style=goZero

gen-swagger: ## Generate API docs
	goctl api plugin -plugin goctl-swagger="swagger -filename $(SERVICE).json -host $(HOST) -basepath /" -api cmd/$(SERVICE)/desc/*.api -dir ./docs/swagger