package aoc

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/tofoss/aoc-go/pkg/readers"
)

func FetchInput(year, day int, example bool) ([]string, error) {
	input, err := fetchFromCache(year, day, example)
	if err == nil {
		return input, nil
	}

	if example {
		input, err = readUserInput()
	} else {
		input, err = fetchFromAoc(year, day)
	}

	err = saveInput(year, day, input)

	if err != nil {
		return nil, err
	}

	return input, nil
}

func fetchFromAoc(year, day int) ([]string, error) {
	sessionCookie := os.Getenv("AOC_SESSION_COOKIE")

	if sessionCookie == "" || sessionCookie == "your_session_cookie_here" {
		fmt.Println("AOC_SESSION_COOKIE not found")
		return readUserInput()
	}

	return downloadInput(year, day, sessionCookie)
}

func downloadInput(year, day int, sessionCookie string) ([]string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", sessionCookie))
	req.Header.Add("User-Agent", "github.com/tofoss/aoc-go")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to download input: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(body)), "\n"), nil
}

func saveInput(year, day int, input []string) error {
	dir := fmt.Sprintf("input/%d/%02d/", year, day)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	filename := filepath.Join(dir, "actual.txt")
	return os.WriteFile(filename, []byte(strings.Join(input, "\n")), 0644)
}

func fetchFromCache(year, day int, example bool) ([]string, error) {
	path := fmt.Sprintf("input/%d/%02d/", year, day)

	if example {
		path += "example.txt"
	} else {
		path += "actual.txt"
	}

	return readers.ReadLines(path)
}

func readUserInput() ([]string, error) {
	fmt.Println("Paste your input below, then press Enter twice when done:")

	scanner := bufio.NewScanner(os.Stdin)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
