package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/urfave/cli/v2"
)

var benchmarkCmd = &cli.Command{
	Name:        "benchmark",
	Usage:       "benchmarks each day and part to a nice stdout table",
	Description: `benchmarks each day and part to a nice stdout table`,
	Action:      runBenchmark,
}

func runBenchmark(_ *cli.Context) error {
	// list all days and parts to benchmark
	// for each day and part, run the benchmark
	// print the results to stdout
	os.Setenv("_SUBMIT", "1")
	t := table.NewWriter()
	t.SetTitle("Benchmark Results AoC 2023")
	// print t.Render() to stdout
	t.SetOutputMirror(os.Stdout)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "Day", AutoMerge: true, Align: text.AlignLeft},
		{Name: "Part", AutoMerge: true, Align: text.AlignLeft},
		{Name: "Answer", AutoMerge: true, Align: text.AlignLeft},
		{Name: "Time", AutoMerge: true, Align: text.AlignLeft},
	})
	t.AppendHeader(table.Row{"Day", "Part", "Answer", "Time"})

	// walk the dir and run the tests.
	fnWalkDir := func(path string, d fs.DirEntry, _ error) error {
		if !strings.HasPrefix(path, "day") || !strings.HasSuffix(path, "test.go") {
			return nil
		}

		pathSplit := strings.Split(path, string(os.PathSeparator))
		basenameSplit := strings.Split(d.Name(), "_")[0]

		part := string(basenameSplit[len(basenameSplit)-1])
		dayDir := pathSplit[0]
		args := []string{
			"test",
			"-v",
			fmt.Sprintf("./%s/...", dayDir),
			fmt.Sprintf("-run=Test_solvePart%s", part),
			"-count=1",
		}

		// run the benchmark
		cmd := exec.Command("go", args...)
		out, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("error running benchmark: %w", err)
		}

		// parse the stdout and append
		answer, time := parseStdout(string(out))
		t.AppendRow(table.Row{dayDir, part, answer, time})
		return nil
	}

	if err := filepath.WalkDir(".", fnWalkDir); err != nil {
		return fmt.Errorf("error walking dir: %w", err)
	}

	t.Render()

	return nil
}
