package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gonum/matrix/dense"
)

func parseMatrix(input string) (*dense.Mat64, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWhitespace)
	matrixData := make([][]float64, 0, scanner.Len()/2)

	for scanner.Scan() {
		row := make([]float64, 0, scanner.Len()/2)
		for scanner.Scan() {
			val, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				return nil, fmt.Errorf("error parsing float: %v", err)
			}
			row = append(row, val)
		}
		matrixData = append(matrixData, row)
	}

	matrix := dense.New(len(matrixData), len(matrixData[0]))
	for i := 0; i < len(matrixData); i++ {
		for j := 0; j < len(matrixData[i]); j++ {
			matrix.Set(i, j, matrixData[i][j])
		}
	}
	return matrix, nil
}

func multiplyMatrices(m1, m2 *dense.Mat64) (*dense.Mat64, error) {
	if m1.Cols() != m2.Rows() {
		return nil, fmt.Errorf("number of columns in m1 must equal number of rows in m2")
	}

	product := dense.New(m1.Rows(), m2.Cols())
	product.Mul(m1, m2)
	return product, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Matrix Operations Calculator")
	fmt.Println("Enter matrices in the form: <rows> <columns> <matrix_elements>")
	fmt.Println("Example for a 2x3 matrix: 2 3 1 2 3 4 5 6")
	fmt.Println("Enter 'exit' to quit the program.")

	for {
		fmt.Print("Enter matrices (separate by a new line):\nMatrix 1:\n")
		matrix1Input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}

		fmt.Print("Matrix 2:\n")
		matrix2Input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}

		input := strings.TrimSpace(matrix1Input + matrix2Input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting calculator.")
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Invalid input format.")
			continue
		}

		rows1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error parsing number of rows in matrix 1.")
			continue
		}

		cols1, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error parsing number of columns in matrix 1.")
			continue
		}

		rows2, err := strconv.Atoi(parts[3])
		if err != nil {
			fmt.Println("Error parsing number of rows in matrix 2.")
			continue
		}

		cols2, err := strconv.Atoi(parts[4])
		if err != nil {
			fmt.Println("Error parsing number of columns in matrix 2.")
			continue
		}

		if rows1 != rows2 {
			fmt.Println("Number of rows in both matrices must be equal.")
			continue
		}

		if cols1 != cols2 {
			fmt.Println("Number of columns in both matrices must be equal.")
			continue
		}

		matrix1, err := parseMatrix(matrix1Input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		matrix2, err := parseMatrix(matrix2Input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := multiplyMatrices(matrix1, matrix2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Result of matrix multiplication:")
		for i := 0; i < result.Rows(); i++ {
			for j := 0; j < result.Cols(); j++ {
				fmt.Printf("%.6f ", result.At(i, j))
			}
			fmt.Println()
		}
	}
}
