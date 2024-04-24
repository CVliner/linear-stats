package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go data.txt")
		os.Exit(1)
	}

	path := os.Args[1]

	data, err := ReadFile(path)
	if err != nil {
		log.Fatal("Error reading data:", err)
	}

	a, b := Linear_Regression(data)
	pcc := Pearson(data)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", a, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", pcc)
}

// ReadFile reads data from a file and returns a slice of integers.
func ReadFile(path string) ([]float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func Linear_Regression(data []float64) (float64, float64) {
	length := float64(len(data))

	var X float64
	var Y float64
	var XY float64
	var XX float64

	for i, y := range data {
		x := float64(i)
		X += x
		Y += y
		XY += x * y
		XX += x * x
	}

	a := (length*XY - X*Y) / (length*XX - X*X)
	b := (Y - a*X) / length

	return a, b
}

func Pearson(data []float64) float64 {
	length := float64(len(data))

	var X float64
	var Y float64
	var XY float64
	var XX float64
	var YY float64

	for i, y := range data {
		x := float64(i)
		X += x
		Y += y
		XY += x * y
		XX += x * x
		YY += y * y
	}

	num := length*XY - X*Y
	den := (length*XX - X*X) * (length*YY - Y*Y)

	if den == 0 {
		return 0
	}

	return num / math.Sqrt(den)
}
