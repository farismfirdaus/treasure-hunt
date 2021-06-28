package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var block [][]string
	block = append(block, []string{"#", "#", "#", "#", "#", "#", "#", "#"})
	block = append(block, []string{"#", ".", ".", ".", ".", ".", ".", "#"})
	block = append(block, []string{"#", ".", "#", "#", "#", ".", ".", "#"})
	block = append(block, []string{"#", ".", ".", ".", "#", ".", "#", "#"})
	block = append(block, []string{"#", "X", "#", ".", ".", ".", ".", "#"})
	block = append(block, []string{"#", "#", "#", "#", "#", "#", "#", "#"})

	// Validate slice of slice string to same number of columns on every rows
	// also validate the value, allowed -> "X" "#" "."
	err := validateMap(block)
	if err != nil {
		panic(err)
	}

	// find start potion, "X"
	y, x := findStartPosition(block)
	// find probability treasure location, return grid map & solution
	block, solution := TreasureHunt(block, y, x)

	// printing solution
	fmt.Printf("SOLUTION %d found [x, y] -> %v\n\n", len(solution), solution)

	// printing grid map
	for _, v := range block {
		for _, value := range v {
			if strings.EqualFold(value, "─") || strings.EqualFold(value, "┬") || strings.EqualFold(value, "├") || strings.EqualFold(value, "┌") {
				fmt.Printf("%v─", value)
			} else {
				fmt.Printf("%v ", value)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// TreasureHunt let you find probable cordinate for treasure!:
// TreasureHunt(mapValue)
func TreasureHunt(mapValue [][]string, y int, x int) (result [][]string, solution []string) {
	for yAxis := y - 1; yAxis >= 0; yAxis-- {
		if strings.EqualFold(mapValue[yAxis][x], "#") {
			break
		} else if strings.EqualFold(mapValue[yAxis][x+1], ".") && strings.EqualFold(mapValue[yAxis-1][x], "#") {
			mapValue[yAxis][x] = "┌"
			rightStep(mapValue, solution, yAxis, x)
		} else if strings.EqualFold(mapValue[yAxis][x+1], ".") {
			mapValue[yAxis][x] = "├"
			rightStep(mapValue, solution, yAxis, x)
		} else {
			mapValue[yAxis][x] = "│"
		}
	}
	return mapValue, solution
}

func downStep(mapValue [][]string, solution []string, y int, x int) ([][]string, []string) {
	for yAxis := y + 1; yAxis < len(mapValue); yAxis++ {
		if strings.EqualFold(mapValue[yAxis][x], "#") {
			break
		} else {
			if !existInSolution(solution, fmt.Sprintf("[%v,%v]", x, yAxis)) {
				solution = append(solution, fmt.Sprintf("[%v,%v]", x, yAxis))
				mapValue[yAxis][x] = "$"
			}
		}
	}
	return mapValue, solution
}

func rightStep(mapValue [][]string, solution []string, y int, x int) ([][]string, []string) {
	for xAxis := x + 1; xAxis < len(mapValue[0]); xAxis++ {
		if strings.EqualFold(mapValue[y][xAxis], "#") {
			break
		} else if strings.EqualFold(mapValue[y+1][xAxis], ".") && strings.EqualFold(mapValue[y][xAxis+1], "#") {
			mapValue[y][xAxis] = "┐"
			mapValue, solution = downStep(mapValue, solution, y, xAxis)
		} else if strings.EqualFold(mapValue[y+1][xAxis], ".") || strings.EqualFold(mapValue[y+1][xAxis], "─") || strings.EqualFold(mapValue[y+1][xAxis], "┐") {
			mapValue[y][xAxis] = "┬"
			mapValue, solution = downStep(mapValue, solution, y, xAxis)
		} else {
			mapValue[y][xAxis] = "─"
		}
	}
	return mapValue, solution
}

func existInSolution(value []string, needle string) bool {
	for _, v := range value {
		if strings.EqualFold(v, needle) {
			return true
		}
	}
	return false
}

func findStartPosition(value [][]string) (y int, x int) {
	for y := 0; y < len(value); y++ {
		for x := 0; x < len(value[0]); x++ {
			if strings.EqualFold(value[y][x], "X") {
				return y, x
			}
		}
	}
	return 0, 0
}

func validateMap(value [][]string) error {
	for i, v := range value {
		if len(v) != len(value[0]) {
			return errors.New(fmt.Sprintf("number of columns should be the same! row index [%d]: %v", i, v))
		}
		for index, val := range v {
			if val != "." && val != "#" && val != "X" {
				return errors.New(fmt.Sprintf("invalid value at [%d][%d]: %v", index, i, val))
			}
		}
	}
	return nil
}
