.PHONY: install	
install:
	@cd view && npm install -g pnpm && pnpm install


.PHONY: run
run: 
	@go run app/main.go


.PHONY: run-dev
run-dev:
	@go run app/main.go
	@cd view && pnpm dev
