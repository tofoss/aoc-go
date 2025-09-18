# Advent of Code Go

## Creating a new day

```bash
make new YEAR=2021 DAY=1
```

This creates the day structure and adds the import to `main.go`.

## Running solutions

```bash
go run . -year 2021 -day 1           # Run with actual input
go run . -year 2021 -day 1 -example  # Run with example input
```

Input files go in `input/YEAR/DAY/actual.txt` and `input/YEAR/DAY/example.txt`.

## Auto-fetch inputs (optional)

To automatically download inputs, create a `.env` file with your session cookie:

1. Log in to [adventofcode.com](https://adventofcode.com)
2. Open dev tools (F12) → Application/Storage → Cookies
3. Copy the `session` cookie value
4. Add to `.env`: `AOC_SESSION_COOKIE=your_cookie_here`