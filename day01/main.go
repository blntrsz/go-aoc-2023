package day01

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func convertAndSum(lines []string) int {
  var newLines []string

  for _, line := range lines {
    chars := strings.Split(line, "")

    var newLine string
    for _, char := range chars {
      isMatching, err := regexp.MatchString("\\d", char) 

      if err != nil {
        continue
      }

      if isMatching {
        newLine += char
      }
    }

    newLines = append(newLines, newLine)
  }

  var sum int
  for _, line := range newLines {
    if len(line) < 1 {
      continue
    }

    if len(line) == 1 {
      s, err := strconv.Atoi(line + line)
      if err != nil {
        panic("Cannot convert to integer")
      }

      sum += s
      continue
    }

    hd := string(line[0])
    tl := string(line[len(line)-1])

    s, err := strconv.Atoi(hd + tl)

    if err != nil {
      panic("Cannot convert to integer")
    }

    sum += s
  }

  return sum
}

func PartOne(input string) int {
  lines := strings.Split(input, "\n")
  return convertAndSum(lines)
}

func convertNumbersToDigits(lines []string) []string {
  numbers := [9]string{ "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" }

  for lineIndex, line := range lines {

    var newLine string
    newLine = line
    for index, n := range numbers {
      re := regexp.MustCompile(n)

      newLine = re.ReplaceAllString(newLine, fmt.Sprintf("%s%s%s", n, fmt.Sprint(index+1), n))
    }

    lines[lineIndex] = newLine
  }

  for lineIndex, line := range lines {
    var newLine string
    newLine = line
    for index, n := range numbers {
      re := regexp.MustCompile(n)

      newLine = re.ReplaceAllString(newLine, fmt.Sprintf("%s", fmt.Sprint(index+1)))
    }

    lines[lineIndex] = newLine
  }

  return lines
}

func PartTwo(input string) int {
  lines := strings.Split(input, "\n")

  lines = convertNumbersToDigits(lines)

  return convertAndSum(lines)
}

