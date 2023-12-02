package day02

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_solvePart2(t *testing.T) {
	type args struct {
		input func() (string, error)
	}
	tests := []struct {
		args args
		name string
		want int
	}{
		{
			name: "solvePart2() with test input",
			args: args{
				input: func() (string, error) {
					return `
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`, nil
				},
			},
			want: 2286,
		},
		{
			name: "solvePart2() with input.txt",
			args: args{
				input: func() (string, error) {
					out, err := os.ReadFile("input.txt")
					if err != nil {
						return "", err
					}
					return string(out), nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := tt.args.input()
			if err != nil {
				t.Fatal(err)
			}
			_, ok := os.LookupEnv("_SUBMIT")
			if tt.want == 0 && !ok {
				t.Skip()
			}

			if tt.want != 0 && ok {
				t.Skip()
			}

			tn := time.Now()
			got := solvePart2(input)
			if got != tt.want && !ok {
				t.Errorf("got = %v, want %v", got, tt.want)
			}

			// print answer to stdout for piping
			if ok {
				fmt.Printf("got: %v\n", got)
				fmt.Printf("took: %v\n", time.Since(tn))
			}
		})
	}
}
