OPENAPI_GENERATOR_CLI_VERSION=v6.0.0
DIST=dist
FRONTEND_DIR := $(dir $(abspath "."))

gen: gen/swagger

gen/swagger:
	docker run --rm \
		-w /work \
		-v $(FRONTEND_DIR)/../:/work openapitools/openapi-generator-cli:$(OPENAPI_GENERATOR_CLI_VERSION) 
		generate \
		-i ./swagger/swagger.yml \
		-g typescript-fetch \
		-o ./frontend/src/openapi