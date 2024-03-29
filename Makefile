SERVER_BINARY=server

GO=go

GOFLAGS=

DEST=build/

.PHONY: all clean clean-build build build-server build-server-prepare run watch watchGo watchVite lint lint-js lint-go

include .env

all: build

build:
	$(MAKE) clean-build
	$(MAKE) build-server 

build-server-prepare:
	templ generate
	pnpm vite build
	mv ${DEST}/public/assets/manifest.json ./internal/middlewares/manifest.json
	cp -r ./public ${DEST}/

build-server: build-server-prepare
	$(GO) build $(GOFLAGS) -o ${DEST}$(SERVER_BINARY) ./cmd/server/.

debug:
	templ generate
	$(GO) build -tags debug -o ${DEST}$(SERVER_BINARY) ./cmd/server/.

run: build-server
	cd ${DEST}; ./$(SERVER_BINARY)

dev:
	${MAKE} -j2 watchGo watchVite

watchGo:
	air

watchVite:
	bun vite --host

clean:
	$(GO) clean
	$(MAKE) clean-build

clean-build:
	rm -rf $(DEST)

lint:
	${MAKE} -j3 lint-go lint-js lint-typechecks

lint-go:
	templ generate
	golangci-lint run ./...

lint-js:
	pnpm run lint

lint-typechecks:
	pnpm run typechecks 

