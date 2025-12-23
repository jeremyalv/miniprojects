package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

type Matrix[T any] [][]T

func CreateMatrix[T any](m, n int, initial T) Matrix[T] {
	if m <= 0 || n <= 0 {
		return nil
	}

	matrix := make(Matrix[T], m+1)

	for i := 0; i <= n; i++ {
		matrix[i] = make([]T, n+1)

		// Fill the row with the initial value
		for j := 0; j <= n; j++ {
			matrix[i][j] = initial
		}
	}

	return matrix
}

func ComputeLCSMatrix(c [][]int, x, y []string) [][]int {
	fmt.Println(c)
	fmt.Println("range x: ", len(x))
	fmt.Println("range y: ", len(y))

	for i, xi := range x {
		for j, yj := range y {
			ni := i+1
			nj := j+1
			fmt.Printf("i: %d, j: %d\n", ni, nj)
			
			if reflect.DeepEqual(xi, yj){
				c[ni][nj] = 1 + c[ni-1][nj-1]
			} else {
				c[ni][nj] = max(c[ni][nj-1], c[ni-1][nj])
			}
		}
	}

	return c
}

func BacktrackLCS(c [][]int, x, y []string, i, j int) string {
	// Backtrack the LCS length matrix to get the actual LCS
	if i == -1 || j == -1 {
		return ""
	}
	if x[i] == y[i] {
		return BacktrackLCS(c, x, y, i-1, j-1) + x[i]
	}
	if c[i][j-1] > c[i-1][j] {
		return BacktrackLCS(c, x, y, i, j-1)
	} else {
		return BacktrackLCS(c, x, y, i-1, j)
	}
}

func PrintDiff(c [][]int, x, y []string, i, j int) {
	i = i+1
	j = j+1
	if i < 0 && j < 0 {
		return
	}
	if i < 0 {
		PrintDiff(c, x, y, i, j-1)
		fmt.Println("+ " + y[j])
	}
	if j < 0 {
		PrintDiff(c, x, y, i-1, j)
		fmt.Println("- " + x[i])
	}
	if x[i-1] == y[j-1] {
		PrintDiff(c, x, y, i-1, j-1)
		fmt.Println(" " + x[i-1])
	}
	if c[i-1][j-1-1] >= c[i-1-1][j-1] {
		PrintDiff(c, x, y, i, j-1)
		fmt.Println("+ " + y[j-1])
	}
	PrintDiff(c, x, y, i-1, j)
	fmt.Println("- " + x[i-1])
}

func diff(x, y []string) {
	c := CreateMatrix(len(x), len(y), 0)
	m := ComputeLCSMatrix(c, x, y)
	PrintDiff(m, x, y, len(x)-1, len(y)-1)
}

func main() {
	f1, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Errorf("error")
		return
	}
	defer f1.Close()

	f2, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Errorf("error")
		return

	}
	defer f2.Close()

	sc1 := bufio.NewScanner(f1)
	in1 := make([]string, 0)
	for sc1.Scan() {
		in1 = append(in1, sc1.Text())
	}

	sc2 := bufio.NewScanner(f2)
	in2 := make([]string, 0)
	for sc2.Scan() {
		in2 = append(in2, sc2.Text())
	}

	fmt.Printf("in1: (%d) %s, \n in2: (%d) %s\n", len(in1), in1, len(in2), in2)
	diff(in1, in2)
}
