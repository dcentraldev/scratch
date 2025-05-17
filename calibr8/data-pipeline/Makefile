
OUTPUT_FOLDER=bin
OUTPUT_BINARY=calibr8

include ./config.properties

.PHONY: init
init:
	go mod tidy
	go fmt ./...


.PHONY: run
run:
	go run cli/main.go \
		--dataset="test" \
		--helius_grpc_url=${HELIUS_GRPC_URL} \
		--helius_api_key=${HELIUS_API_KEY}

.PHONY: build
build: clean init
	go build -o ${OUTPUT_FOLDER}/${OUTPUT_BINARY} cli/main.go

.PHONY: clean
clean:
	rm -rf ${OUTPUT_FOLDER}