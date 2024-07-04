test:
	@go test -v ./...

gen:
	@templ generate

mod:
	@go mod tidy

mongo:
	@docker start mongo 

run:
	@go run main.go