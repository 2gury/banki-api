CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.61.0
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=banks-api/cmd/app

all: format build test lint

build: bindir
	go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" -o ${BINDIR}/app ${PACKAGE}

test:
	GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn go test ./...

run:
	GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn go run ${PACKAGE}

lint: install-lint
	${LINTBIN} run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN:=$(CURDIR)/bin
bin:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Устанавливаем proto описания google/googleapis
vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis vendor-proto/googleapis &&\
 	cd vendor-proto/googleapis &&\
	git sparse-checkout set --no-cone google/api &&\
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/googleapis/google/api vendor-proto/google
	rm -rf vendor-proto/googleapis

# Устанавливаем proto описания google/protobuf
vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
	cd vendor-proto/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/protobuf/src/google/protobuf vendor-proto/google
	rm -rf vendor-proto/protobuf

generate: 
	generate-banks-api
	generate-reviews-api

generate-banks-api: bin vendor-proto/google/api vendor-proto/google/protobuf
	mkdir -p pkg/banks
	$(PROTOC) -I api/banks -I vendor-proto \
	--go_out pkg/banks --go_opt paths=source_relative \
	--go-grpc_out pkg/banks --go-grpc_opt paths=source_relative \
	--grpc-gateway_out pkg/banks --grpc-gateway_opt paths=source_relative \
	--validate_out="lang=go,paths=source_relative:pkg/banks" \
	api/banks/banks.proto

generate-reviews-api: bin vendor-proto/google/api vendor-proto/google/protobuf
	mkdir -p pkg/reviews
	$(PROTOC) -I api/reviews -I vendor-proto \
	--go_out pkg/reviews --go_opt paths=source_relative \
	--go-grpc_out pkg/reviews --go-grpc_opt paths=source_relative \
	--grpc-gateway_out pkg/reviews --grpc-gateway_opt paths=source_relative \
	--validate_out="lang=go,paths=source_relative:pkg/reviews" \
	api/reviews/reviews.proto

WEB_APP = web
BUILD_DIR = $(PWD)/bin

.PHONY: test bench run all

all: web cli

web:
	go build -o $(BUILD_DIR)/$(WEB_APP) cmd/$(WEB_APP)/main.go

run: web
	$(BUILD_DIR)/$(WEB_APP)

test:
	go test -v -race ./...

bench:
	go test -bench=. -benchmem ./...