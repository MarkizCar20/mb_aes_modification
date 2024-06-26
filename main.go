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
			sum := val1 + val2 // Use XOR for addition in GF(2^8)
			//sum := val1 + val2 // Should be XOR instead of ADD, bcs GF(256)
			result[i][j] = fmt.Sprintf("%02x", sum)
		}
	}

	return result
}

func computeInverseMatrix(matrix [4][4]string) [4][4]string {
	var inverseMatrix [4][4]string
	inverses := computeMultiplicationInverses()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			val, err := strconv.ParseInt(matrix[i][j], 16, 64)
			if err != nil {
				fmt.Println("Error converting hex to int:", err)
				continue
			}
			inverseMatrix[i][j] = fmt.Sprintf("%02x", inverses[val])
		}
	}
	return inverseMatrix
}

func computeMultiplicationInverses() [256]byte {
	var inverses [256]byte
	inverses[0] = 0

	for x := 1; x < 256; x++ {
		for y := 1; y < 256; y++ {
			if gfMultiply(byte(x), byte(y)) == 1 {
				inverses[x] = byte(y)
				break
			}
		}
	}

	return inverses
}

func checkInverses(m1, m2 [4][4]string) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			val1, err1 := strconv.ParseInt(m1[i][j], 16, 64)
			val2, err2 := strconv.ParseInt(m2[i][j], 16, 64)

			if err1 != nil || err2 != nil {
				fmt.Println("Error converting hex to int", err1, err2)
				return false
			}
			if gfMultiply(byte(val1), byte(val2)) != 1 {
				fmt.Printf("Mismatch at (%d, %d): %02x * %02x != 01\n", i, j, val1, val2)
				return false
			}
		}
	}
	return true
}

func gfMultiply(a, b byte) byte {
	var result byte = 0
	var polynomial byte = 0x1B // Polynomial for reduction: x^8 + x^4 + x^3 + x + 1 -> 0b11011 in hex is 0x1B

	for i := 0; i < 8; i++ {
		if (b & 1) != 0 {
			result ^= a
		}
		carry := (a & 0x80) != 0
		a <<= 1
		if carry {
			a ^= polynomial
		}
		b >>= 1
	}
	return result
}

func byteToVector(b byte) [8]int {
	var vector [8]int
	binaryStr := fmt.Sprintf("%08b", b)
	for i := 0; i < 8; i++ {
		vector[i] = int(binaryStr[i] - '0')
	}
	return vector
}

func matrixToVectors(m [4][4]string) [4][4][8]int {
	var vectorMatrix [4][4][8]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			val, err := strconv.ParseUint(m[i][j], 16, 8)
			if err != nil {
				fmt.Printf("Error converting hex to int", err)
				continue
			}
			vectorMatrix[i][j] = byteToVector(byte(val))
		}
	}
	return vectorMatrix
}

func printVectorMatrix(vMatrix [4][4][8]int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%v ", vMatrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Hello go")

	word := "Two One Nine Two"
	key := "1234567891234567"
	wordMatrix := convertToHexMatrix(word)
	keyMatrix := convertToHexMatrix(key)
	sumMatrix := addMatrix(wordMatrix, keyMatrix)
	inverseMatrix := computeInverseMatrix(sumMatrix)
	fmt.Println("Hex matrix of the word: ")
	printMatrix(wordMatrix)
	fmt.Println("<----------------------->")
	fmt.Println("Hex matrix of the key: ")
	printMatrix(keyMatrix)
	fmt.Println("<----------------------->")
	fmt.Println("Hex matrix of their sum")
	printMatrix(sumMatrix)
	fmt.Println("<----------------------->")
	fmt.Println("Inverse Matrix: ")
	printMatrix(inverseMatrix)
	fmt.Println("<----------------------->")
	bitVectorMatrix := matrixToVectors(inverseMatrix)
	fmt.Println("<----------------------->")
	fmt.Println("Matrix of 8b vectors of the inverse sum: ")
	printVectorMatrix(bitVectorMatrix)
	fmt.Println("<----------------------->")

}
