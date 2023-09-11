export

SAMPLE_DATA_FOLDER=output

# ==============================================================================
# Sequential Writing
## example: make sequential_writing TOTAL=10000000 FILE_NAME=tenmillion.txt
.PHONY: sequential_writing sequential_writing_benchmark

sequential_writing:
	@if [ -z "$(TOTAL)" ]; then echo >&2 "Please set TOTAL via the variable TOTAL"; exit 2; fi
	@if [ -z "$(FILE_NAME)" ]; then echo >&2 "Please set FILE_NAME via the variable FILE_NAME"; exit 2; fi
	rm -f "${SAMPLE_DATA_FOLDER}/${FILE_NAME}"
	echo "Generating file ${SAMPLE_DATA_FOLDER}/${FILE_NAME}..."
	go run sequential_writing/sequential_writing.go --lines=$(TOTAL) -filename="${SAMPLE_DATA_FOLDER}/sequential_writing_${FILE_NAME}"
	echo "Finished generate ${SAMPLE_DATA_FOLDER}/${FILE_NAME}."

sequential_writing_benchmark:
	cd sequential_writing && go test -bench=. .

# ==============================================================================
# Asynchronous I/O
## example: make asynchronous_io TOTAL=10000000 FILE_NAME=tenmillion.txt
.PHONY: asynchronous_io asynchronous_io_benchmark

asynchronous_io:
	@if [ -z "$(TOTAL)" ]; then echo >&2 "Please set TOTAL via the variable TOTAL"; exit 2; fi
	@if [ -z "$(FILE_NAME)" ]; then echo >&2 "Please set FILE_NAME via the variable FILE_NAME"; exit 2; fi
	rm -f "${SAMPLE_DATA_FOLDER}/${FILE_NAME}"
	echo "Generating file ${SAMPLE_DATA_FOLDER}/${FILE_NAME}..."
	go run main.go --lines=$(TOTAL) --filename="${SAMPLE_DATA_FOLDER}/asynchronous_io_${FILE_NAME}" --method=asynchronous_io
	echo "Finished generate ${SAMPLE_DATA_FOLDER}/${FILE_NAME}."

asynchronous_io_benchmark:
	go test ./file_writer/asynchronous_io -bench=. > ./file_writer/asynchronous_io/benchmark_results.txt

# ==============================================================================
# Tests

.PHONY: test
## test: runs tests
test:
	@ go test -v ./...

.PHONY: test-coverage
## test-coverage: run unit tests and generate coverage report in html format
test-coverage:
	@ go test -coverprofile=coverage.out ./...  && go tool cover -html=coverage.out