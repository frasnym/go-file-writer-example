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
| 10          | 6,066      | 194,561 ns/op             | 194.561µs         |
| 100         | 6,807      | 201,701 ns/op             | 201.701µs         |
| 1,000       | 5,788      | 222,166 ns/op             | 222.166µs         |
| 10,000      | 1,779      | 718,845 ns/op             | 718.845µs         |
| 100,000     | 247        | 4,280,770 ns/op           | 4.280ms           |
| 1,000,000   | 24         | 47,067,297 ns/op          | 47.067ms          |
| 10,000,000  | 4          | 315,281,906 ns/op         | 315.282ms         |

These benchmark results were obtained on a system with the following specifications:
- Operating System: Darwin
- Architecture: arm64
