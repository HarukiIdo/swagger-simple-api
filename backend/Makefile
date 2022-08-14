.PHONY: install
install:
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: gen
gen:
	swagger generate model -f ../swagger/swagger.yml -m generated_swagger -t ./swagger