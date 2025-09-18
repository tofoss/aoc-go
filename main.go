package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/tofoss/aoc-go/2021/day01"
	"github.com/tofoss/aoc-go/pkg/registry"
)

func main() {
	var year, day int
	var useExample bool

	flag.IntVar(&year, "year", time.Now().Year(), "Year to run")
	flag.IntVar(&day, "day", time.Now().Day(), "Day to run")
	flag.BoolVar(&useExample, "example", false, "Use example input")
	flag.Parse()

	err := runDay(year, day, useExample)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func runDay(year, day int, useExample bool) error {
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

	input := fmt.Sprintf("input/%d/%02d/", year, day)

	if useExample {
		input = input + "example.txt"
	} else {
		input = input + "actual.txt"
	}

	solver, err := registry.Registry[year][day](input)

	if err != nil {
		return err
	}

	fmt.Printf("%s Year %d, Day %d %s\n", selectedEmoji, year, day, selectedEmoji)
	p1, err := solver.Part1()
	if err != nil {
		return err
	}
	fmt.Print(p1)

	p2, err := solver.Part2()
	if err != nil {
		return err
	}
	fmt.Print(p2)

	return nil
}
