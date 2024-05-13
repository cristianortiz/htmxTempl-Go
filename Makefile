run: build 
	@./bin/app

build: templ
	@go build -o bin/app cmd/app/main.go

templ:
	@templ generate