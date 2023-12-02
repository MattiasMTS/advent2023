# ❄️ advent2023 ❄️

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

By default it will test today's and part 1, but can be changed via
the `day` and `part` keywords.

3. submit your answer

```
make submit
make submit day=13 part=1
```

By default it will submit today's and part 1, but can be changed via
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


## Benchmark

+------------------------------------+
