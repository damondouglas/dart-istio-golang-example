export image=echo:1.0
export proxy=proxy:1.0

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

gen: ## Stub proto
	protoc -I proto proto/echo.proto --go_out=plugins=grpc:proto --dart_out=grpc:client/dart/lib/src/generated

build-server: ## Build docker image
	docker build -t ${image} -f server/Dockerfile .

sh-server: ## sh into image
	docker run -it ${image} sh

run-server: ## run image locally
	docker run -it -p 9090:9090 ${image}

build-proxy: ## Build proxy image
	docker build -t ${proxy} -f proxy/Dockerfile .

run-proxy: ## Run proxy locally
	docker run -p 9901:9901 -p 8080:8080 -it ${proxy}

run-web: ## Run web server
	cd client/dart/app; \
	webdev serve web:5000