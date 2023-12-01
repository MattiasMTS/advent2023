// templates/generate_aoc.go

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

const part1TemplateContent = `package day{{.Day}}

func solvePart1(input string) int {
	// Implement your solution for part 1
	return 0
}
`

const part2TemplateContent = `package day{{.Day}}

func solvePart2(input string) int {
	// Implement your solution for part 2
	return 0
}
`

const part1TestTemplateContent = `package day{{.Day}}

import (
	"fmt"
	"os"
	"testing"
	"time"
)

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
		{
			name: "solvePart1() with input.txt",
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
			got := solvePart1(input)
			if got != tt.want && !ok {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}

			// print answer to stdout for piping
			if ok {
				fmt.Printf("got: %v\n", got)
				fmt.Printf("took: %v\n", time.Since(tn))
			}
		})
	}
}
`

const part2TestTemplateContent = `package day{{.Day}}

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
					return "", nil // TODO: Add test input here.
				},
			},
			want: 0, // TODO: Add expected output here.
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
				t.Errorf("solvePart2() = %v, want %v", got, tt.want)
			}

			// print answer to stdout for piping
			if ok {
				fmt.Printf("got: %v\n", got)
				fmt.Printf("took: %v\n", time.Since(tn))
			}
		})
	}
}
`

// errInvalidGenerateArgs is the error message when the arguments are invalid.
var errInvalidGenerateArgs = "Usage: go run templates/generate_aoc.go day=X"

func main() {
	if len(os.Args) < 2 {
		log.Fatal(errInvalidGenerateArgs)
	}

	d := os.Args[1][4:]
	if d == "" {
		log.Fatal(errInvalidGenerateArgs)
	}

	// convert day to integer
	day, err := strconv.Atoi(d)
	if err != nil {
		log.Fatal(err)
	}

	// Create folder for the day
	dayFolder := fmt.Sprintf("day%02d", day)
	if err = os.Mkdir(dayFolder, 0755); err != nil && !os.IsExist(err) {
		log.Fatal("Error creating folder:", err)
	}

	// Generate exercise and test files
	files := map[string]string{
		"part1.go":      part1TemplateContent,
		"part2.go":      part2TemplateContent,
		"part1_test.go": part1TestTemplateContent,
		"part2_test.go": part2TestTemplateContent,
	}
	for name, content := range files {
		if err := generateFile(filepath.Join(dayFolder, name), content, struct{ Day string }{Day: fmt.Sprintf("%02d", day)}); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Generated files in folder: %q\n", dayFolder)
}

func generateFile(filePath, templateContent string, data any) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating %s: %v", filePath, err)
	}
	defer file.Close()

	template := template.Must(template.New("template").Parse(templateContent))
	return template.Execute(file, data)
}
