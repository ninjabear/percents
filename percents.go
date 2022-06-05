package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func roundToPlate(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func Percents(weight int, breakdowns []int) []float64 {
	result := make([]float64, len(breakdowns))

	for i, breakdown := range breakdowns {
		result[i] = float64(weight) * (float64(breakdown) / 100.0)
	}

	return result
}

func RenderPercents(percents []float64) []string {
	result := make([]string, len(percents))
	roundTo := 2.5
	for i, p := range percents {
		result[i] = fmt.Sprintf("%g", roundToPlate(p, roundTo))
	}
	return result
}

func main() {
	var weight int

	if len(os.Args) == 1 {
		var err error
		for err != nil || weight <= 0 {
			fmt.Print("Enter weight: ")
			_, err = fmt.Scanf("%d\n", &weight)
		}
	} else {
		w, e := strconv.Atoi(os.Args[1])
		if e != nil {
			os.Stderr.WriteString("Invalid command line argument")
			os.Exit(1)
		} else {
			weight = w
		}
	}

	breakdowns := []int{100, 95, 90, 80, 70, 60, 50, 40, 30}
	lines := RenderPercents(Percents(weight, breakdowns))
	for i, line := range lines {
		fmt.Printf("%3d%%: %s\r\n", breakdowns[i], line)
	}
}
