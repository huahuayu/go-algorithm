package solution

import (
	"fmt"
	"testing"
)

func TestInitMatrixAndPrint(t *testing.T) {
	matrix := initMatrix(5)
	printMatrix(matrix)
}

func TestRotate(t *testing.T) {
	matrix := initMatrix(5)
	fmt.Println("original matrix:")
	printMatrix(matrix)
	rotate(matrix)
	fmt.Println("rotated matrix:")
	printMatrix(matrix)
}
