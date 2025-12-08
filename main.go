package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/tofoss/aoc-go/2021/day01"
	_ "github.com/tofoss/aoc-go/2025/day01"
	_ "github.com/tofoss/aoc-go/2025/day02"
	_ "github.com/tofoss/aoc-go/2025/day03"
	_ "github.com/tofoss/aoc-go/2025/day04"
	_ "github.com/tofoss/aoc-go/2025/day05"
	_ "github.com/tofoss/aoc-go/2025/day06"
	_ "github.com/tofoss/aoc-go/2025/day07"
	_ "github.com/tofoss/aoc-go/2025/day08"
	"github.com/tofoss/aoc-go/pkg/aoc"
	"github.com/tofoss/aoc-go/pkg/registry"
)

func main() {
	var year, day, part int
	var useExample bool
	var testInput string

	flag.IntVar(&year, "year", time.Now().Year(), "Year to run")
	flag.IntVar(&day, "day", time.Now().Day(), "Day to run")
	flag.IntVar(&part, "part", 0, "Part to run")
	flag.BoolVar(&useExample, "example", false, "Use example input")
	flag.StringVar(&testInput, "input", "", "Use test input")
	flag.Parse()

	godotenv.Load()
	debug.SetMemoryLimit(16 << 30)

	err := runDay(year, day, part, useExample, testInput)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runDay(year, day, part int, useExample bool, testInput string) error {
	emoji := []string{
		"🎄",
		"🎅",
		"🦌",
		"⛄",
		"❄️",
		"🎁",
		"🌟",
		"🕯️",
		"🔔",
		"🛷",
		"🧣",
		"🧤",
		"🍪",
		"🥛",
		"🪵",
		"🔥",
		"🎶",
		"🌌",
		"🦉",
		"🌙",
		"✨",
		"🧑‍🎄",
		"🐧",
		"🎂",
		"🎇",
	}

	selectedEmoji := emoji[(day-1)%len(emoji)]

	runPart1, runPart2 := true, true
	switch part {
	case 1:
		runPart1 = true
		runPart2 = false
	case 2:
		runPart1 = false
		runPart2 = true
	}

	input, err := aoc.FetchInput(year, day, useExample)
	if err != nil {
		return err
	}
	if testInput != "" {
		input = strings.Split(testInput, "\n")
	}

	solver, err := registry.Registry[year][day](input)

	if err != nil {
		return err
	}

	fmt.Printf("%s Year %d, Day %d %s\n", selectedEmoji, year, day, selectedEmoji)

	startTime := time.Now()
	if runPart1 {
		p1, err := solver.Part1()
		if err != nil {
			return err
		}
		fmt.Printf("part one: %s\n", p1)
		timeAfterP1 := time.Now()
		duration := timeAfterP1.Sub(startTime)
		fmt.Printf("time elapsed: %v\n", duration)
	}

	if runPart2 {
		p2StartTime := time.Now()
		p2, err := solver.Part2()
		if err != nil {
			return err
		}
		fmt.Printf("part two: %s\n", p2)
		timeAfterP2 := time.Now()
		duration := timeAfterP2.Sub(p2StartTime)
		fmt.Printf("time elapsed: %v\n", duration)
	}
	fmt.Printf("total time elapsed: %v\n", time.Since(startTime))

	return nil
}
