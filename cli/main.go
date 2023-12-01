package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

var (
	// App is entrypoint to the cli
	App *cli.App
	// reAnswer is the stdout for correct answer
	reAnswer = regexp.MustCompile("got: ([0-9]+)")
	// reTime is the stdout for time taken
	reTime = regexp.MustCompile(`took: ([^---]+)`)
	// errTooQuick is the stdout for submitting too quickly
	errTooQuick = regexp.MustCompile("You gave an answer too recently.*to wait.")
	// errWrong is the stdout for wrong answer
	errWrong = regexp.MustCompile(`That's not the right answer.*?\.`)
	// errCorrect is the stdout for correct answer
	errCorrect = regexp.MustCompile("That's the right answer!")
	// errAlreadyDone is the stdout for already solved puzzle
	errAlreadyDone = regexp.MustCompile(`You don't seem to be solving.*\?`)
	// sessionToken is the cookie value for session
	sessionToken string
	// baseURL is the base url for advent of code
	baseURL = "https://adventofcode.com"
)

func main() {
	initEnv()

	if err := App.Run(os.Args); err != nil {
		log.Println(err)
	}

	// cleanup
	os.Unsetenv("_SUBMIT")
	os.Exit(0)
}

func init() {
	App = &cli.App{}
	App.Name = "aoc"
	App.Usage = "Advent of Code CLI"
	App.Version = "0.0.1"
	App.EnableBashCompletion = true
	App.UsageText = `
Advent of Code CLI
`
	submitCmd := &cli.Command{
		Name:        "submit",
		Usage:       "Submit your solution to Advent of Code",
		Description: `Submit your solution to Advent of Code.`,
		Action:      submitAction,
	}

	submitCmd.Flags = []cli.Flag{
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

	inputCmd := &cli.Command{
		Name:        "input",
		Usage:       "Get input for a specific day",
		Description: `Get input for a specific day.`,
		Action:      getAction,
	}

	inputCmd.Flags = []cli.Flag{
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

	App.Commands = []*cli.Command{
		submitCmd,
		inputCmd,
	}
}

func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	if sessionToken = os.Getenv("SESSION_TOKEN"); sessionToken == "" {
		log.Fatal("SESSION_TOKEN not found in .env")
	}
}

func getAction(c *cli.Context) error {
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

func submitAction(c *cli.Context) error {
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

	if resp = checkResponse(resp); resp == "" {
		return fmt.Errorf("failed to parse response")
	}

	fmt.Println(resp)

	return nil
}

// getPipedStdinData reads the stdin and returns the available data as a string
// if and only if the data was piped to the process
func getPipedStdinData() string {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return ""
	}
	stdinData := ""
	if (fi.Mode()&os.ModeCharDevice) == 0 && fi.Size() > 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdinData = fmt.Sprintf("%s%s", stdinData, scanner.Text())
		}
	}
	return stdinData
}

func request(method, url string, b io.Reader) (string, error) {
	client := http.DefaultClient

	// setup http client
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionToken,
	})

	if method == http.MethodPost {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	// Read response Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getInput(year, day string) (string, error) {
	u := fmt.Sprintf("%s/%s/day/%s/input", baseURL, year, day)
	return request(http.MethodGet, u, nil)
}

func submit(answer, year, day, part string) (string, error) {
	u := fmt.Sprintf("%s/%s/day/%s/answer", baseURL, year, day)
	v := url.Values{}
	v.Add("level", part)
	v.Add("answer", answer)

	return request(http.MethodPost, u, strings.NewReader(v.Encode()))
}

func checkResponse(resp string) string {
	for _, r := range []*regexp.Regexp{errTooQuick, errWrong, errAlreadyDone, errCorrect} {
		m := r.FindString(resp)
		if m != "" {
			return m
		}
	}

	return ""
}
