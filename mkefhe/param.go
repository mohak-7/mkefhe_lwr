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
	q := uint64(1 << 16) // modulus q
	p := uint64(1 << 11)  // modulus p < q
	n := 8             // Lattice dimension

	k := 10              

	l1 := int(math.Ceil(math.Log2(float64(q))))
	l2 := int(math.Ceil(math.Log2(float64(p))))
	m := n*l1 + l2 // big oh

	N := n*l1 + l2

	A := utils.SampleUniformMatrix(m, n, q)


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