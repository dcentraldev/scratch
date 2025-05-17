
OUTPUT_FOLDER=bin
OUTPUT_BINARY=calibr8
GCP_PROJECT=kunal-scratch
GCP_REGION=us-central1

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


.PHONY: run-df
run-df: clean init
	go run cli/main.go \
		--dataset="test" \
		--helius_grpc_url=${HELIUS_GRPC_URL} \
		--helius_api_key=${HELIUS_API_KEY} \
		--project=${GCP_PROJECT} \
		--runner=dataflow \
		--region=${GCP_REGION} \
		--temp_location="gs://${GCP_PROJECT}/temp/" \
		--max_num_workers=1 \
		--worker_machine_type=n1-standard-2 \
		--staging_location="gs://${GCP_PROJECT}/staging/"
