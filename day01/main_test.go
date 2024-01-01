package day01_test

import (
	"os"
	"testing"

	"github.com/blntrsz/go-aoc-2023/day01"
)

var input1 = `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

func TestPartOneDummyInput(t *testing.T) {
  got := day01.PartOne(input1)
  expected := 142

  if got != expected {
    t.Errorf("expected %v get %v", expected, got)
  }
}

func getInput() string {
  file, _ := os.ReadFile("input.txt")

  return string(file) 
}

func TestPartOne(t *testing.T) {
  got := day01.PartOne(getInput())
  expected := 54081

  if got != expected {
    t.Errorf("expected %v get %v", expected, got)
  }
}

var input2 = `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func TestPartTwoDummyInput(t *testing.T) {
  got := day01.PartTwo(input2)
  expected := 281

  if got != expected {
    t.Errorf("expected %v get %v", expected, got)
  }
}

func TestPartTwo(t *testing.T) {
  got := day01.PartTwo(getInput())
  expected := 54649

  if got != expected {
    t.Errorf("expected %v get %v", expected, got)
  }
}

