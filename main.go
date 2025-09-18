package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/tofoss/aoc-go/2021/day01"
	"github.com/tofoss/aoc-go/pkg/aoc"
	"github.com/tofoss/aoc-go/pkg/registry"
)

func main() {
	var year, day int
	var useExample bool

	flag.IntVar(&year, "year", time.Now().Year(), "Year to run")
	flag.IntVar(&day, "day", time.Now().Day(), "Day to run")
	flag.BoolVar(&useExample, "example", false, "Use example input")
	flag.Parse()

	godotenv.Load()

	err := runDay(year, day, useExample)

	if err != nil {
		fmt.Println(err)
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

	input, err := aoc.FetchInput(year, day, useExample)
	if err != nil {
		return err
	}

	solver, err := registry.Registry[year][day](input)

	if err != nil {
		return err
	}

	timeBeforeP1 := time.Now()
	fmt.Printf("%s Year %d, Day %d %s\n", selectedEmoji, year, day, selectedEmoji)
	p1, err := solver.Part1()
	if err != nil {
		return err
	}
	fmt.Print(p1)
	timeAfterP1 := time.Now()
	duration := timeAfterP1.Sub(timeBeforeP1)
	fmt.Printf("time elapsed: %v\n", duration)

	p2, err := solver.Part2()
	if err != nil {
		return err
	}
	fmt.Print(p2)
	timeAfterP2 := time.Now()
	duration = timeAfterP2.Sub(timeAfterP1)
	fmt.Printf("time elapsed: %v\n", duration)
	fmt.Printf("total time elapsed: %v\n", time.Since(timeBeforeP1))

	return nil
}
