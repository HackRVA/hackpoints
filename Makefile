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

.PHONY:test
test: ## test the code.
##
	go test ./...
