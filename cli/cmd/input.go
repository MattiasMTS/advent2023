package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli/v2"
)

var inputCmd = &cli.Command{
	Name:        "input",
	Usage:       "Get input for a specific day",
	Description: `Get input for a specific day.`,
	Action:      runInput,
	Flags:       setInputFlags(),
}

// runInput runs the input command.
func runInput(c *cli.Context) error {
	d := c.String("day")
	day, err := strconv.Atoi(d)
	if err != nil {
		log.Fatal(err)
	}

	b, err := getInput(c.String("year"), fmt.Sprintf("%d", day))
	if err != nil {
		return fmt.Errorf("error getting input: %w", err)
	}

	dayDir := filepath.Join(fmt.Sprintf("day%02d", day), "input.txt")
	if err = os.MkdirAll(filepath.Dir(dayDir), 0755); err != nil {
		return fmt.Errorf("error creating folder: %w", err)
	}
	return os.WriteFile(dayDir, []byte(b), 0644)
}

// getInput calls the advent of code api to get the input for a specific day.
func getInput(year, day string) (string, error) {
	u := fmt.Sprintf("%s/%s/day/%s/input", baseURL, year, day)
	return request(http.MethodGet, u, nil)
}

// setInputFlags sets the flags for the input command.
func setInputFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "day",
			Aliases:  []string{"d"},
			Usage:    "Day of the puzzle to solve",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "year",
			Aliases:  []string{"y"},
			Usage:    "Year of the puzzle to solve",
			Required: true,
		},
	}
}
