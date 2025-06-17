package mkefhe

import (
	// "fmt"
	"mkefhe_lwr/utils"
)

func GadgetVector(l float64, m uint64) []float64 {
	vec := make([]float64, int(l))
	for i := 0; i < int(l); i++ {
		vec[i] = float64((1 << i) % int(m))
	}
	return vec
}

func GadgetMatrix(p, q uint64, n, N, l1, l2 int) [][]float64 {
	g1 := GadgetVector(float64(l1), q)
	g2 := GadgetVector(float64(l2), p)

	G := make([][]float64, n+1)
	for i := range G {
		G[i] = make([]float64, N)
	}

	for i := 0; i < l2; i++ {
		G[0][i] = g2[i]
	}

	for block := 0; block < n; block++ {
		offset := 1 + block
		for i := 0; i < l1; i++ {
			G[offset][l1*block+l2+i] = g1[i]
		}
	}

	return G
}

type CipherText struct {
	C   [][]float64 // m x N matrix
	// c0 [][]float64 // 1 x N row
	// c1 [][]float64 // n x N matrix
}
func (c CipherText) GetC() [][]float64 {
	return c.C
}
func (c CipherText) GetC0() [][]float64 {
	return c.C[0:1]
}
func (c CipherText) GetC1() [][]float64 {
	return c.C[1:]
}

func Encrypt(epk ExtendedPublicKey, msg uint8, pp PublicParams) CipherText {
	N := pp.GetBigN()
	n := pp.GetSmallN()
	m := pp.GetM()
	p := pp.GetP()
	q := pp.GetQ()
	l1 := pp.GetL1()
	l2 := pp.GetL2()

	R := utils.SampleUniformMatrix(m, N, 2)	// [][]float64 of size m x N
	G := GadgetMatrix(p, q, n, N, l1, l2)	// [][]float64 of size (n+1) x N

	BT := utils.Transpose(epk.GetMatrix())	// [][]float64 of size (n+1) x m

	var C [][]float64
	if msg == 1 {
		C = utils.AddMatrices(G, utils.MultiplyMatrices(BT, R))
	} else {
		C = utils.MultiplyMatrices(BT, R)
	}

	for j:=0;j<N;j++ {
		C[0][j] = utils.UnsignedMod(C[0][j], float64(p))
	}

	for i:= 1; i < n+1; i++ {
		for j := 0; j < N; j++ {
			C[i][j] = utils.UnsignedMod(C[i][j], float64(q))
		}
	}


	return CipherText{
		C : C,
	}
}
