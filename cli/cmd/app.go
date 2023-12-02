package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

var (
	// App is entrypoint to the cli
	App *cli.App
	// sessionToken is the cookie value for session
	sessionToken string
	// baseURL is the base url for advent of code
	baseURL = "https://adventofcode.com"
	// reAnswer is the stdout for correct answer
	reAnswer = regexp.MustCompile("got: ([0-9]+)")
	// reTime is the stdout for time taken
	reTime = regexp.MustCompile(`took: ([^---]+)`)
	// reTooQuick is the stdout for submitting too quickly
	reTooQuick = regexp.MustCompile("You gave an answer too recently.*to wait.")
	// reWrong is the stdout for wrong answer
	reWrong = regexp.MustCompile(`That's not the right answer.*?\.`)
	// reCorrect is the stdout for correct answer
	reCorrect = regexp.MustCompile("That's the right answer!")
	// reAlreadyDone is the stdout for already solved puzzle
	reAlreadyDone = regexp.MustCompile(`You don't seem to be solving.*\?`)
)

func init() {
	initEnv()

	App = &cli.App{}
	App.Name = "aoc"
	App.Usage = "Advent of Code CLI"
	App.Version = "0.0.1"
	App.EnableBashCompletion = true
	App.UsageText = `
Advent of Code CLI
`

	App.Commands = []*cli.Command{
		submitCmd,
		inputCmd,
	}
}

// initEnv loads the .env file and sets the sessionToken.
func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	if sessionToken = os.Getenv("SESSION_TOKEN"); sessionToken == "" {
		log.Fatal("SESSION_TOKEN not found in .env")
	}
}

// request makes a http request to the given url with the given method and body
// and returns the response body as a string.
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

// parseResponse checks the response string for known errors and returns the
// parsed message based on the response body.
func parseResponse(resp string) string {
	for _, r := range []*regexp.Regexp{reTooQuick, reWrong, reAlreadyDone, reCorrect} {
		m := r.FindString(resp)
		if m != "" {
			return m
		}
	}

	return ""
}
