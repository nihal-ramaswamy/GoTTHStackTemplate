run: templ
	@go run ./cmd/app/main.go

build: templ
	@go build -o bin/app ./cmd/app/main.go

test: templ
	@go test -v ./...

templ:
	@templ generate

clean:
	@rm -rf bin
