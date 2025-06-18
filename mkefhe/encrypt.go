package mkefhe

import (
	// "fmt"
	"mkefhe_lwr/utils"
)

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
	G := utils.GadgetMatrix(p, q, n, N, l1, l2)	// [][]float64 of size (n+1) x N

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
