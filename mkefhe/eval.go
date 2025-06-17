package mkefhe

import (
	"fmt"
	"mkefhe_lwr/utils"
)

func CipherAdd(c1, c2 CipherText, parameters PublicParams) CipherText {
	c1C := c1.GetC()
	c2C := c2.GetC()
	Cadd := utils.AddMatrices(c1C, c2C)

	for i := 0; i < len(Cadd[0]); i++ {
		Cadd[0][i] = utils.UnsignedMod(Cadd[0][i], float64(parameters.GetP()))
	}
	// the remaining rows of Cadd should be reduced modulo q
	for i := 1; i < len(Cadd); i++ {
		for j := 0; j < len(Cadd[i]); j++ {
			Cadd[i][j] = utils.UnsignedMod(Cadd[i][j], float64(parameters.GetQ()))
		}
	}
	return CipherText{
		C : Cadd,
	}

	// if len(c1.GetC0()) != len(c2.GetC0()) || len(c1.GetC1()) != len(c2.GetC1()) {
	// 	panic("Ciphertexts must have the same dimensions")
	// }

	// N := len(c1.GetC0()[0])
	// n := len(c1.GetC1())

	// C0 := make([][]float64, 1)
	// C0[0] = make([]float64, N)

	// for i := 0; i < N; i++ {
	// 	C0[0][i] = c1.GetC0()[0][i] + c2.GetC0()[0][i]
	// }

	// C1 := make([][]float64, n)
	// for i := 0; i < n; i++ {
	// 	C1[i] = make([]float64, N)
	// 	for j := 0; j < N; j++ {
	// 		C1[i][j] = c1.GetC1()[i][j] + c2.GetC1()[i][j]
	// 	}
	// }

	// return CipherText{
	// 	c0: C0,
	// 	c1: C1,
	// }
}

func CipherMult(C1, C2 CipherText, parameters PublicParams) CipherText{
	c1 := C1.GetC()
	c2 := C2.GetC()

	// fmt.Println("C2 : ", c2)

	GinvC2 := utils.Ginv(c2, parameters.GetL1(), parameters.GetL2(), parameters.GetSmallN())

	// fmt.Println("GinvC2 : ", GinvC2)

	// Multiply c1 with Ginv(c2)
	Cmult := utils.MultiplyMatrices(c1, GinvC2)
	fmt.Println("Cmult dimension: ", len(Cmult), "x", len(Cmult[0]))

	// the first row of Cmult should be reduced modulo p
	for i := 0; i < len(Cmult[0]); i++ {
		Cmult[0][i] = utils.UnsignedMod(Cmult[0][i], float64(parameters.GetP()))
	}

	// the remaining rows of Cmult should be reduced modulo q
	for i := 1; i < len(Cmult); i++ {
		for j := 0; j < len(Cmult[i]); j++ {
			Cmult[i][j] = utils.UnsignedMod(Cmult[i][j], float64(parameters.GetQ()))
		}
	}

	// fmt.Println(Cmult)
	// fmt.Println("Cmult dimension: ", len(Cmult), "x", len(Cmult[0]))

	return CipherText{
		C : Cmult,
	}
}