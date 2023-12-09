package day07

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
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`, nil
				},
			},
			want: 5905,
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
