package main

import "fmt"

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

func main() {
	fmt.Println("Hello go")

	word := "Two One Nine Two"
	matrix := convertToHexMatrix(word)
	printMatrix(matrix)

}
