# go-file-writer-example
An example to write file in Go

## Benchmark Results

Here are the benchmark results for each package:

### Sequential File Writer

| Lines       | Operations | Nanoseconds per Operation | Time per Operation |
|-------------|------------|---------------------------|-------------------|
| 10          | 27,145     | 47,018 ns/op              | 47.018µs           |
| 100         | 24,067     | 53,568 ns/op              | 53.568µs           |
| 1,000       | 8,190      | 148,878 ns/op             | 148.878µs          |
| 10,000      | 1,185      | 1,047,948 ns/op           | 1.048ms            |
| 100,000     | 124        | 9,348,030 ns/op           | 9.348ms            |
| 1,000,000   | 13         | 88,569,997 ns/op          | 88.570ms           |
| 10,000,000  | 2          | 928,590,208 ns/op         | 928.590ms          |

### Parallel File Writer

| Lines       | Operations | Nanoseconds per Operation | Time per Operation |
|-------------|------------|---------------------------|-------------------|
| 10          | 17,680     | 67,004 ns/op              | 67.004µs           |
| 100         | 17,887     | 66,847 ns/op              | 66.847µs           |
| 1,000       | 10,000     | 100,178 ns/op             | 100.178µs          |
| 10,000      | 2,041      | 598,333 ns/op             | 598.333µs          |
| 100,000     | 253        | 4,808,686 ns/op           | 4.808ms            |
| 1,000,000   | 28         | 41,799,996 ns/op          | 41.800ms           |
| 10,000,000  | 3          | 477,522,000 ns/op         | 477.522ms          |

### Parallel Chunk File Writer

| Lines       | Operations | Nanoseconds per Operation | Time per Operation |
|-------------|------------|---------------------------|-------------------|
| 10          | 10,332     | 131,032 ns/op             | 131.032µs          |
| 100         | 10,000     | 134,346 ns/op             | 134.346µs          |
| 1,000       | 6,606      | 194,377 ns/op             | 194.377µs          |
| 10,000      | 2,287      | 539,779 ns/op             | 539.779µs          |
| 100,000     | 265        | 4,284,090 ns/op           | 4.284ms            |
| 1,000,000   | 31         | 37,744,698 ns/op          | 37.745ms           |
| 10,000,000  | 3          | 452,873,153 ns/op         | 452.873ms          |

These benchmark results were obtained on a system with the following specifications:
- Operating System: Darwin
- Architecture: arm64
