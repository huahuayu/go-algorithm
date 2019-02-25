package solution

import (
	"fmt"
	"math/rand"
	"time"
)

/*
给定一个 n × n 的二维矩阵表示一个图像。

将图像顺时针旋转 90 度。

说明：

你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
*/

/*
java code:
class Solution {

    public void rotate(int[][] matrix) {
        if (matrix == null || matrix.length == 0 || matrix[0].length == 0) {
            return;
        }

        int col = matrix[0].length-1;
        int row = matrix.length-1;

        int colF = col;
        int rowF = row;

        for (int r = 0; r < rowF ; r++) {
            for (int c = r; c < colF; c++) {
                int x1 = r;
                int y1 = c;

                int x2 = c;
                int y2 = col - r;

                int x3 = row - r;
                int y3 = col - c;

                int x4 = row - c;
                int y4 = r;

                swap(matrix, x1, y1, x2, y2);
                swap(matrix, x1, y1, x3, y3);
                swap(matrix, x1, y1, x4, y4);
            }
            rowF--;
            colF--;
        }
    }

    public void swap(int[][] matrix, int x1, int y1, int x2, int y2) {
        int temp = matrix[x1][y1];
        matrix[x1][y1] = matrix[x2][y2];
        matrix[x2][y2] = temp;
    }
}
*/
// 初始化一个n*n方的矩阵
func initMatrix(n int) [][]int {
	matrix := make([][]int, n)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i] = append(matrix[i], rand.Intn(100))
		}
	}

	return matrix
}

// 打印矩阵
func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}

// 翻转矩阵
func rotate(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	col := len(matrix[0]) - 1
	row := len(matrix) - 1

	colF := col
	rowF := row

	for r := 0; r < rowF; r++ {
		//fmt.Println("level:", r)
		for c := r; c < colF; c++ {
			x1 := r
			y1 := c

			x2 := c
			y2 := col - r

			x3 := row - r
			y3 := col - c

			x4 := row - c
			y4 := r

			//fmt.Println("x1,y1:", x1, y1)
			//fmt.Println("x2,y2:", x2, y2)
			//fmt.Println("x3,y3:", x3, y3)
			//fmt.Println("x4,y4:", x4, y4)

			swap(matrix, x1, y1, x2, y2)
			swap(matrix, x1, y1, x3, y3)
			swap(matrix, x1, y1, x4, y4)

		}
		rowF--
		colF--
	}
}

// 交换矩阵中的点o1(x1,y1), o2(x2,y2)
func swap(matrix [][]int, x1, y1, x2, y2 int) {
	matrix[x1][y1], matrix[x2][y2] = matrix[x2][y2], matrix[x1][y1]
}
