run_http:
	go mod tidy; go run cmd/http/*.go

swag:
	swag init -g cmd/http/main.go ./docs; swag fmt
