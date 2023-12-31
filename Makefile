export

SAMPLE_DATA_FOLDER=output

# ==============================================================================
# Sequential Writing
## example: make sequential TOTAL=10000000 FILE_NAME=tenmillion.txt
.PHONY: sequential sequential_benchmark

sequential:
	@if [ -z "$(TOTAL)" ]; then echo >&2 "Please set TOTAL via the variable TOTAL"; exit 2; fi
	@if [ -z "$(FILE_NAME)" ]; then echo >&2 "Please set FILE_NAME via the variable FILE_NAME"; exit 2; fi
	rm -f "${SAMPLE_DATA_FOLDER}/${FILE_NAME}"
	go run main.go --lines=$(TOTAL) --filename="${SAMPLE_DATA_FOLDER}/sequential_${FILE_NAME}" --method=sequential

sequential_benchmark:
	go test ./filewriter/sequential -bench=. > ./filewriter/sequential/benchmark_results.txt

# ==============================================================================
# Parallel Processing
## example: make parallel TOTAL=10000000 FILE_NAME=tenmillion.txt
.PHONY: parallel parallel_benchmark

parallel:
	@if [ -z "$(TOTAL)" ]; then echo >&2 "Please set TOTAL via the variable TOTAL"; exit 2; fi
	@if [ -z "$(FILE_NAME)" ]; then echo >&2 "Please set FILE_NAME via the variable FILE_NAME"; exit 2; fi
	rm -f "${SAMPLE_DATA_FOLDER}/${FILE_NAME}"
	go run main.go --lines=$(TOTAL) --filename="${SAMPLE_DATA_FOLDER}/parallel_${FILE_NAME}" --method=parallel

parallel_benchmark:
	go test ./filewriter/parallel -bench=. > ./filewriter/parallel/benchmark_results.txt

# ==============================================================================
# File Chunking and Parallel Writing
## example: make parallelchunk TOTAL=10000000 FILE_NAME=tenmillion.txt
.PHONY: parallelchunk parallelchunk_benchmark

parallelchunk:
	@if [ -z "$(TOTAL)" ]; then echo >&2 "Please set TOTAL via the variable TOTAL"; exit 2; fi
	@if [ -z "$(FILE_NAME)" ]; then echo >&2 "Please set FILE_NAME via the variable FILE_NAME"; exit 2; fi
	rm -f "${SAMPLE_DATA_FOLDER}/${FILE_NAME}"
	go run main.go --lines=$(TOTAL) --filename="${SAMPLE_DATA_FOLDER}/parallelchunk_${FILE_NAME}" --method=parallelchunk

parallelchunk_benchmark:
	go test ./filewriter/parallelchunk -bench=. > ./filewriter/parallelchunk/benchmark_results.txt

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

.PHONY: mock-gen
## mock-gen: generate mock using go generate
mock-gen:
	go generate ./...