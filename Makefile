export image=echo:1.0
export web=echo-web:1.0

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

gen: ## Stub proto
	protoc -I proto proto/echo.proto --go_out=plugins=grpc:proto --dart_out=grpc:client/dart/lib/src/generated

build-server: ## Build server docker image
	docker build -t ${image} -f server/Dockerfile .

build-web: ## Build web server image
	cd client/dart/app; \
	webdev build; \
	docker build -t ${web} .

print-pod: ## Print pod name
	@kubectl get po --selector="app=${pod}" -o jsonpath='{.items[*].metadata.name}'

