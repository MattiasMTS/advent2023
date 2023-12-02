package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/urfave/cli/v2"
)

var submitCmd = &cli.Command{
	Name:        "submit",
	Usage:       "Submit your solution to Advent of Code",
	Description: `Submit your solution to Advent of Code.`,
	Action:      runSubmit,
	Flags:       setSubmitFlags(),
}

// runSubmit runs the submit command.
func runSubmit(c *cli.Context) error {
	in := getPipedStdinData()
	if in == "" {
		return fmt.Errorf("error piping, got value: %q", in)
	}

	answer := reAnswer.FindStringSubmatch(in)
	time := reTime.FindStringSubmatch(in)

	fmt.Println(answer[1])
	fmt.Println(time[1])

	resp, err := submit(answer[1], c.String("year"), c.String("day"), c.String("part"))
	if err != nil {
		return fmt.Errorf("error submitting: %w", err)
	}

	if resp = parseResponse(resp); resp == "" {
		return fmt.Errorf("failed to parse response")
	}

	fmt.Println(resp)

	return nil
}

// submit calls the advent of code api to submit the answer for a specific day
// and part.
func submit(answer, year, day, part string) (string, error) {
	u := fmt.Sprintf("%s/%s/day/%s/answer", baseURL, year, day)
	v := url.Values{}
	v.Add("level", part)
	v.Add("answer", answer)

	return request(http.MethodPost, u, strings.NewReader(v.Encode()))
}

// setSubmitFlags sets the flags for the submit command.
func setSubmitFlags() []cli.Flag {
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
		&cli.StringFlag{
			Name:     "part",
			Aliases:  []string{"p"},
			Usage:    "Part of the puzzle to solve",
			Required: true,
			Value:    "1",
		},
	}
}
