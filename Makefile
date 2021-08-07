## ----------------------------------------------------------------------
## The purpose of this Makefile is to simplify common development tasks.
## ----------------------------------------------------------------------
##

.PHONY:help
help: ## Show this help.
##
	@sed -ne '/@sed/!s/##//p' $(MAKEFILE_LIST)

.PHONY:run
run: ## run the code.
##
	go run .

.PHONY:build
build: ## build the code.
##
	go build -o hackpoints

.PHONY:clean
clean: ## clean up
	rm ./hackpoints

.PHONY:test
test: ## test the code.
##
	go test ./...

.PHONY:swagger
swagger: ## Generate swagger doc
	swagger generate spec -o ./docs/swaggerui/swagger.json --scan-models
