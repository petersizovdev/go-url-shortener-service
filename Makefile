build:
	@go build -o bin/url-shortener-service cmd/main/main.go

run: build
	@./bin/url-shortener-service