.PHONY: help
help: ## print make targets 
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: go-install-air
go-install-air: ## Installs the air build reload system using 'go install'
	go install github.com/air-verse/air@latest

.PHONY: get-install-air
get-install-air: ## Installs the air build reload system using cUrl
	curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

.PHONY: go-install-templ
go-install-templ: ## Installs the templ Templating system for Go
	go install github.com/a-h/templ/cmd/templ@latest

.PHONY: get-install-tailwindcss
get-install-tailwindcss: ## Installs the tailwindcss cli
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
	chmod +x tailwindcss-linux-x64
	mv tailwindcss-linux-x64 tailwindcss

.PHONY: tailwind-watch
tailwind-watch: ## compile tailwindcss and watch for changes
	./tailwindcss -i ./static/css/custom.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build: ## one-time compile tailwindcss styles
	./tailwindcss -i ./static/css/custom.css -o ./static/css/style.css

.PHONY: build
build: ## compile tailwindcss and templ files and build the project
	./tailwindcss -i ./static/css/custom.css -o ./static/css/style.css
	templ generate
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

.PHONY: watch
watch: ## build and watch the project with air
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch

