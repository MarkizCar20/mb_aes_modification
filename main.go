package main

import (
	"fmt"

	"github.com/chronos-tachyon/gf256"
)

func ConvertStringToMatrix(s string) [4][4]byte {
	var matrix [4][4]byte

	// Convert string to byte slice
	bytes := []byte(s)

	// Fill the matrix
	for i := 0; i < 16 && i < len(bytes); i++ {
		matrix[i/4][i%4] = bytes[i]
	}

	return matrix
}

// PrintMatrix prints the 4x4 matrix in hex format
func PrintMatrix(matrix [4][4]byte) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%02x ", matrix[i][j])
		}
		fmt.Println()
	}
}

func inverseMatrix(m [4][4]byte) [4][4]byte {
	var result [4][4]byte
	field := gf256.New(gf256.Poly11B)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			val, err := field.Inv(m[i][j])
			if err != true {
				fmt.Println("error: ", err)
				continue
			}
			result[i][j] = val
		}
	}

	return result
}

func byteToBits(b byte) [8]byte {
	var bits [8]byte
	for i := 0; i < 8; i++ {
		bits[i] = (b >> (7 - i)) & 1
	}
	return bits
}

func vectorMatrix(m [4][4]byte) [4][4][8]byte {
	var bitMatrix [4][4][8]byte
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			bitMatrix[i][j] = byteToBits(byte(m[i][j]))
		}
	}
	return bitMatrix
}

func MatrixVectorMultiply(matrix [8][8]byte, vector [8]byte) [8]byte {
	var result [8]byte
	for i := 0; i < 8; i++ {
		sum := byte(0)
		for j := 0; j < 8; j++ {
			sum ^= matrix[i][j] & vector[j]
		}
		result[i] = sum
	}
	return result
}

func AddVectors(vec1, vec2 [8]byte) [8]byte {
	var result [8]byte
	for i := 0; i < 8; i++ {
		result[i] = vec1[i] ^ vec2[i]
	}
	return result
}

func processBitMatrix(m [4][4][8]byte) [4][4][8]byte {
	var transformMatrix = [8][8]byte{
		{1, 0, 0, 0, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 1, 1, 1, 1, 1},
	}

	var additionVector = [8]byte{1, 1, 0, 0, 0, 1, 1, 0}

	var resultMatrix [4][4][8]byte
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			intermediate := MatrixVectorMultiply(transformMatrix, m[i][j])
			resultMatrix[i][j] = AddVectors(intermediate, additionVector)
		}
	}

	return resultMatrix
}

// PrintBitMatrix prints the 4x4x8 bit matrix
func PrintBitMatrix(matrix [4][4][8]byte) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(matrix[i][j])
		}
		fmt.Println()
	}
}

func BitsToByte(bits [8]byte) byte {
	var b byte
	for i := 0; i < 8; i++ {
		b |= bits[i] << (7 - i)
	}
	return b
}

// BitVectorsToMatrix converts a 4x4x8 bit matrix back to a 4x4 byte matrix
func BitVectorsToMatrix(bitMatrix [4][4][8]byte) [4][4]byte {
	var matrix [4][4]byte
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			matrix[i][j] = BitsToByte(bitMatrix[i][j])
		}
	}
	return matrix
}

// MatrixToString converts a 4x4 byte matrix back to a string
func MatrixToString(matrix [4][4]byte) string {
	var bytes []byte
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if matrix[i][j] != 0 {
				bytes = append(bytes, matrix[i][j])
			}
		}
	}
	return string(bytes)
}

// SubtractVectors subtracts two 8x1 vectors
func SubtractVectors(a, b [8]byte) [8]byte {
	var result [8]byte
	for i := 0; i < 8; i++ {
		result[i] = a[i] ^ b[i]
	}
	return result
}

func InverseProcessBitMatrix(bitMatrix [4][4][8]byte) [4][4][8]byte {
	var inverseTransformMatrix = [8][8]byte{
		{0, 0, 1, 0, 0, 1, 0, 1},
		{1, 0, 0, 1, 0, 0, 1, 0},
		{0, 1, 0, 0, 1, 0, 0, 1},
		{1, 0, 1, 0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0, 0, 1, 0},
		{0, 0, 1, 0, 1, 0, 0, 1},
		{1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 1, 0},
	}

	var subtractionVector = [8]byte{1, 1, 0, 0, 0, 1, 1, 0}

	var resultMatrix [4][4][8]byte
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			intermediate := SubtractVectors(bitMatrix[i][j], subtractionVector)
			resultMatrix[i][j] = MatrixVectorMultiply(inverseTransformMatrix, intermediate)
		}
	}
	return resultMatrix
}

func main() {
	startingWord := "Swo One Nine Two"
	fmt.Println(startingWord)

	startMatrix := ConvertStringToMatrix(startingWord)
	PrintMatrix(startMatrix)

	// field := gf256.New(gf256.Poly11B)
	invMatrix := inverseMatrix(startMatrix)
	PrintMatrix(invMatrix)
	bitMatrix := vectorMatrix(invMatrix)
	fmt.Println("Bit matrix before processing: ")
	PrintBitMatrix(bitMatrix)
	processedMatrix := processBitMatrix(bitMatrix)
	fmt.Println("Bit matrix after processing: ")
	PrintBitMatrix(processedMatrix)
	finalMatrix := BitVectorsToMatrix(processedMatrix)
	fmt.Println("Final Byte Matrix: ")
	PrintMatrix(finalMatrix)
	finalString := MatrixToString(finalMatrix)
	fmt.Println(finalString)

	matrixAfterChannel := ConvertStringToMatrix(finalString)
	PrintMatrix(matrixAfterChannel)
	bitMatrixAfterChannel := vectorMatrix(matrixAfterChannel)
	PrintBitMatrix(bitMatrixAfterChannel)
	inverseProcessMatrix := InverseProcessBitMatrix(bitMatrixAfterChannel)
	fmt.Println("Inverse Processed Matrix: ")
	PrintBitMatrix(inverseProcessMatrix)
	byteMatrix := BitVectorsToMatrix(inverseProcessMatrix)
	PrintMatrix(byteMatrix)
	finalfinalMatrix := inverseMatrix(byteMatrix)
	finalInverseString := MatrixToString(finalfinalMatrix)
	fmt.Println(finalInverseString)
}
