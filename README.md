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