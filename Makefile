.PHONY: run air-local build-app build-templ build-tailwind fetch-tailwindcss

TAILWIND_RELEASE = tailwindcss-macos-x64

run:
	air

air-local:
	make build-tailwind
	make build-templ
	make build-app

build-app:
	go build -o ./tmp/main ./cmd
	
build-templ:
	templ generate

build-tailwind: tmp/tailwindcss
	./tmp/tailwindcss -i ./web/static/styles/tailwind-input.css -o web/static/styles/tailwind-output.css --minify

fetch-tailwindcss:
	mkdir tmp
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/$(TAILWIND_RELEASE)
	chmod +x $(TAILWIND_RELEASE)
	mv $(TAILWIND_RELEASE) tmp/tailwindcss
