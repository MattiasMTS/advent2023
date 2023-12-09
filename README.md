[![benchmark aoc](https://github.com/MattiasMTS/advent2023/actions/workflows/benchmark.yml/badge.svg)](https://github.com/MattiasMTS/advent2023/actions/workflows/benchmark.yml)

# ❄️ advent2023 ❄️

<!--toc:start-->

- [❄️ advent2023 ❄️](#️-advent2023-️)
- [🛠️ To get started 🛠️](#🛠️-to-get-started-🛠️)
- [🏃Useful commands 🏃](#🏃useful-commands-🏃)
- [🎄Repository Structure 🎄](#🎄repository-structure-🎄)
- [🏋️Benchmark 🏋️](#🏋️benchmark-🏋️)
<!--toc:end-->

https://adventofcode.com/2023

Gopher is coming to town in AoC 2023 🦫

# 🛠️ To get started 🛠️

Start by install `make` on your OS:

- linux: `sudo apt-get install make`
- macos: `brew install make`
- windows: `choco install make`

# 🏃Useful commands 🏃

Useful commands to get you started:

1. generate the daily task:

```shell
make gen
make gen day=9
make gen year=2022 day=9
```

By default it will generate today's task but can be changed via
`year` and `day` keywords.

2. test the code:

```shell
make test
make test day=4 part=2
make test day=4 part=2
```

By default it will test today's task and part 1, but can be changed via
the `day` and `part` keywords.

3. submit your answer

```
make submit
make submit day=13 part=1
```

By default it will submit today's task and part 1, but can be changed via
the `day` and `part` keywords.

# 🎄Repository Structure 🎄

The repository is structured as follows:

```tree
.
├── Makefile
├── README.md
├── cli
│   └── main.go
├── day01
│   ├── input.txt
│   ├── part1.go
│   ├── part1_test.go
│   ├── part2.go
│   └── part2_test.go
├── go.mod
├── go.sum
└── templates
    └── generate_aoc.go
```

- `./cli/` folder is in charge of interacting with https://adventofcode.com website
  for i) fetching the daily input and ii) submitting input.

  You can try it out by running:

  ```shell
  go run cli/main.go --help
  ```

- `./templates/` contains the go templates for generating each day.

- `./dayXX/` contains the daily problems. For each part, we have a solution file
  `partN.go` and a corresponding test file `partN_test.go`.

  The test file can look like this:

  ```go
  func Test_solvePart1(t *testing.T) {
  type args struct {
  	input func() (string, error)
  }
  tests := []struct {
  	args args
  	name string
  	want int
  }{
  	{
  		name: "solvePart1() with test input",
  		args: args{
  			input: func() (string, error) {
  				return "", nil // TODO: Add test input here.
  			},
  		},
  		want: 0, // TODO: Add expected output here.
  	},
  }
  ```

  where you have to manually add the input and answer on the TODOs.
  If you want to debug the `input.txt`, then just comment out the
  `want` field from the `test input` test case and add `want: 1`
  to the `input.txt` test case.

# 🏋️Benchmark 🏋️

```
+----------------------------------------+
| Benchmark Results AoC 2023             |
+-------+------+-----------+-------------+
| DAY   | PART | ANSWER    | TIME        |
+-------+------+-----------+-------------+
| day01 | 1    | 54644     | 258.222µs   |
|       |      |           |             |
|       | 2    | 53348     | 3.314306ms  |
|       |      |           |             |
| day02 | 1    | 2169      | 205.434µs   |
|       |      |           |             |
|       | 2    | 60948     | 224.389µs   |
|       |      |           |             |
| day03 | 1    | 527364    | 6.462963ms  |
|       |      |           |             |
|       | 2    | 79026871  | 6.097792ms  |
|       |      |           |             |
| day04 | 1    | 21158     | 444.359µs   |
|       |      |           |             |
|       | 2    | 6050769   | 623.714µs   |
|       |      |           |             |
| day06 | 1    | 1660968   | 10.81µs     |
|       |      |           |             |
|       | 2    | 26499773  | 30.791923ms |
|       |      |           |             |
| day07 | 1    | 253603890 | 1.347916ms  |
|       |      |           |             |
|       | 2    | 253630098 | 1.327467ms  |
|       |      |           |             |
+-------+------+-----------+-------------+
```

Benchmark table is generated by CI/CD or run `make benchmark`.
