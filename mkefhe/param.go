package mkefhe

import (
	// "fmt"
	"math"
	"mkefhe_lwr/utils"
)

type PublicParams struct {
	q	uint64
	p 	uint64
	n 	int
	m 	int
	k 	int
	l1 	int
	l2 	int
	N 	int
	A	[][]float64
}

func DefaultParams() PublicParams {
	q := uint64(1 << 16) // q = 2^15 = 32768
	p := uint64(1 << 11)  // p = 2^8 = 256
	n := 8             // Lattice dimension 256

	k := 10              

	l1 := int(math.Ceil(math.Log2(float64(q))))
	l2 := int(math.Ceil(math.Log2(float64(p))))
	// m := int(1 * (float64(n)*float64(l1) + float64(l2))) 
	m := n*l1 + l2 // big oh

	N := n*l1 + l2

	A := utils.SampleUniformMatrix(m, n, q)


	// print A in a format which can be copy-pasted into a python file
	// fmt.Print("A = [")
	// for i := 0; i < len(A); i++ {
	// 	fmt.Print("[")
	// 	for j := 0; j < len(A[i]); j++ {
	// 		fmt.Printf("%f", A[i][j])
	// 		if j < len(A[i])-1 {
	// 			fmt.Print(", ")
	// 		}
	// 	}
	// 	fmt.Print("]")
	// 	if i < len(A)-1 {
	// 		fmt.Print(", ")
	// 	}
	// }
	// fmt.Print("]")

	return PublicParams {
		q: q,
		p: p,
		n: n,
		k: k,
		l1: l1,
		l2: l2,
		m: m,
		N: N,
		A: A,
	}
}

func (params PublicParams) GetQ() uint64 {
	return params.q
}
func (params PublicParams) GetP() uint64 {
	return params.p
}
func (params PublicParams) GetSmallN() int {
	return params.n
}
func (params PublicParams) GetM() int {
	return params.m
}
func (params PublicParams) GetK() int {
	return params.k
}
func (params PublicParams) GetL1() int {
	return params.l1
}
func (params PublicParams) GetL2() int {
	return params.l2
}
func (params PublicParams) GetBigN() int {
	return params.N
}
func (params PublicParams) GetA() [][]float64 {
	return params.A
}