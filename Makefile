build:
	@go build -o bin/url-shortener-service cmd/main/main.go

run: buld
	@./bin/url-shortener-service