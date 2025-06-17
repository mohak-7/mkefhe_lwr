package utils

import (
	// "fmt"
	"math"
	"golang.org/x/exp/constraints"
)

func SplitCiphertext(Cstar [][]uint64, n int) ([]uint64, [][]uint64) {
    c0 := Cstar[0] // 1 x N row
    C1 := Cstar[1 : n+1] // n x N matrix
    return c0, C1
}


// func Ginv(v [][]float64, p, q uint64, l1, l2, n int) [][]float64 {
// 	mat := make([][]float64, l2+(n)*l1)

// 	// Decompose the first element
// 	for i := 0;i<len(v[0]);i++ {
// 	    x := math.Mod(v[0][i], float64(p))
// 	    for j := 0; j < l2; j++ {
// 		    mat[j] = make([]float64, 1)
// 		    mat[j][i] = math.Mod(x, 2)
// 		 	x = math.Floor(x / 2)
// 		}
// 	}

// 	// Decompose the remaining elements
// 	for k:= 0; k < len(v[0]); k++ {
// 	    for j := 1; j < n+1; j++ {
// 	    	x := math.Mod(v[j][k], float64(q))
// 	    	offset := l2 + (j-1)*l1
// 	    	for i := 0; i < l1; i++ {
// 	    		mat[offset+i] = make([]float64, 1)
// 	    		mat[offset+i][k] = math.Mod(x, 2)
// 	    		x = math.Floor(x / 2)
// 	    	}
// 	    }
// 	}
// 	return mat
// }

func Ginv(mat [][]float64, l1, l2, n int) [][]float64 {
    cols := len(mat[0])
    rows := l2 + n*l1

    result := make([][]float64, rows)
    for i := range result {
        result[i] = make([]float64, cols)
    }

    // decompose the first row
    for i := 0; i < cols; i++ {
        // x := math.Mod(mat[0][i], float64(p))
        x := mat[0][i]
        for j := 0; j < l2; j++ {
            result[j][i] = math.Mod(x, 2)
            x = math.Floor(x / 2)
        }
    }

    // decompose the remaining rows
    for k := 0; k < cols; k++ {
        for j := 1; j < n+1; j++ {
            // x := math.Mod(mat[j][k], float64(q))
            x := mat[j][k]
            offset := l2 + (j-1)*l1
            for i := 0; i < l1; i++ {
                result[offset+i][k] = math.Mod(x, 2)
                x = math.Floor(x / 2)
            }
        }
    }

    return result
}



func AddMatrices[T constraints.Integer | constraints.Float](A, B [][]T) [][]T {
	m := len(A)
	n := len(A[0])
	result := make([][]T, m)
	for i := 0; i < m; i++ {
		result[i] = make([]T, n)
		for j := 0; j < n; j++ {
			result[i][j] = A[i][j] + B[i][j]
		}
	}
	return result
}

func MultiplyMatrices[T constraints.Integer | constraints.Float](A, B [][]T) [][]T {
    rowsA, colsA := len(A), len(A[0])
    rowsB, colsB := len(B), len(B[0])

    if colsA != rowsB {
        panic("Matrix multiplication not possible: incompatible dimensions")
    }

    C := make([][]T, rowsA)
    for i := range C {
        C[i] = make([]T, colsB)
    }

    for i := 0; i < rowsA; i++ {
        for j := 0; j < colsB; j++ {
            var sum T
            for k := 0; k < colsA; k++ {
                sum += A[i][k] * B[k][j]
            }
            C[i][j] = sum
        }
    }
    return C
}

func MultiplyMatricesMod[T constraints.Integer | constraints.Float](A, B [][]T, mod T) [][]T {
    rowsA, colsA := len(A), len(A[0])
    rowsB, colsB := len(B), len(B[0])

    if colsA != rowsB {
        panic("Matrix multiplication not possible: incompatible dimensions")
    }

    C := make([][]T, rowsA)
    for i := range C {
        C[i] = make([]T, colsB)
    }

    for i := 0; i < rowsA; i++ {
        for j := 0; j < colsB; j++ {
            var sum T
            for k := 0; k < colsA; k++ {
                sum += A[i][k] * B[k][j]
            }
            C[i][j] = SignedMod(sum, mod)
			// C[i][j] = sum
        }
    }
    return C
}

func MultiplyPkSk(A, B [][]float64, mod float64) [][]float64 {
    rowsA, colsA := len(A), len(A[0])
    rowsB, colsB := len(B), len(B[0])

    if colsA != rowsB {
        panic("Matrix multiplication not possible: incompatible dimensions")
    }

    C := make([][]float64, rowsA)
    for i := range C {
        C[i] = make([]float64, colsB)
    }

    for i := 0; i < rowsA; i++ {
        for j := 0; j < colsB; j++ {
            var sum float64
            for k := 0; k < colsA; k++ {
                sum += A[i][k] * B[k][j]
            }
            C[i][j] = SignedMod(sum, mod)
			// C[i][j] = sum
			// C[i][j] = math.Mod(sum , mod)
        }
    }
    return C
}

func MatVecMulMod[T constraints.Integer | constraints.Float](A [][]T, x []T, q T) []T {
	m := len(A)
	if m == 0 {
		return nil
	}
	n := len(x)
	result := make([]T, m)
	for i := 0; i < m; i++ {
		var sum T = 0
		for j := 0; j < n; j++ {
			sum = T(SignedMod(float64(sum + A[i][j]*x[j]), float64(q)))
		}
		result[i] = sum
	}
	return result
}

func Transpose[T any](matrix [][]T) [][]T {
    if len(matrix) == 0 {
        return [][]T{}
    }
    rows, cols := len(matrix), len(matrix[0])
    transposed := make([][]T, cols)
    for i := 0; i < cols; i++ {
        transposed[i] = make([]T, rows)
        for j := 0; j < rows; j++ {
            transposed[i][j] = matrix[j][i]
        }
    }
    return transposed
}

func MultiplyVectors[T constraints.Integer | constraints.Float](a, b []T) T {
    if len(a) != len(b) {
        panic("vectors must be of same length")
    }
    var result T
    for i := range a {
        result += a[i] * b[i]
    }
    return result
}


func SignedMod[T constraints.Integer | constraints.Float](x, q T) T {
	r := math.Mod(float64(x), float64(q))
    if r > float64(q)/2 {
        return T(r - float64(q))
    }
	if r <= -float64(q)/2 {
		return T(r + float64(q))
	}
    return T(r)
}

func UnsignedMod[T constraints.Integer | constraints.Float](x, q T) T {
    r := math.Mod(float64(x), float64(q))
    if r < 0 {
        return T(r + float64(q))
    }
    return T(r)
}