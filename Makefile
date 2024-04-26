test:
	@go test -v ./...

gen:
	@templ generate

mod:
	@go mod tidy

