.PHONY: pay 
pay: ## run payment-service
	go run payserv/serve/server.go
.PHONY: api
api: ## run backend-api
	go run backapi/main.go