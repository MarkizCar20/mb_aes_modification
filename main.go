package main

import (
	"fmt"
	"strconv"
)

func convertToHexMatrix(word string) [4][4]string {
	var matrix [4][4]string
	paddedWord := []byte(word)
	if len(word) < 16 {
		paddedWord = append(paddedWord, make([]byte, 16-len(word))...)
	}

	for i := 0; i < 16; i++ {
		matrix[i/4][i%4] = fmt.Sprintf("%02x", paddedWord[i])
	}
	return matrix
}

func printMatrix(matrix [4][4]string) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%s ", val)
		}
		fmt.Println()
	}
}

func addMatrix(m1 [4][4]string, m2 [4][4]string) [4][4]string {
	var result [4][4]string

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			val1, err1 := strconv.ParseInt(m1[i][j], 16, 64)
			val2, err2 := strconv.ParseInt(m2[i][j], 16, 64)

			if err1 != nil || err2 != nil {
				fmt.Println("Error converting hex to int", err1, err2)
			}
			sum := val1 + val2
			result[i][j] = fmt.Sprintf("%02x", sum)
		}
	}

	return result
}

func main() {
	fmt.Println("Hello go")

	word := "Two One Nine Two"
	key := "1234567891234567"
	wordMatrix := convertToHexMatrix(word)
	keyMatrix := convertToHexMatrix(key)
	sumMatrix := addMatrix(wordMatrix, keyMatrix)
	fmt.Println("Hex matrix of the word: ")
	printMatrix(wordMatrix)
	fmt.Printf("<----------------------->")
	fmt.Println("Hex matrix of the key: ")
	printMatrix(keyMatrix)
	fmt.Printf("<----------------------->")
	fmt.Println("Hex matrix of their sum")
	printMatrix(sumMatrix)
	fmt.Printf("<----------------------->")

}
