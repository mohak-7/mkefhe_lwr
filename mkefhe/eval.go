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
}

func CipherMult(C1, C2 CipherText, parameters PublicParams) CipherText{
	c1 := C1.GetC()
	c2 := C2.GetC()

	GinvC2 := utils.Ginv(c2, parameters.GetL1(), parameters.GetL2(), parameters.GetSmallN())

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

func CipherNand(C1, C2 CipherText, parameters PublicParams) CipherText {
	c1 := C1.GetC()
	c2 := C2.GetC()

	GinvC2 := utils.Ginv(c2, parameters.GetL1(), parameters.GetL2(), parameters.GetSmallN())

	Cnand := utils.SubtractMatrices(utils.GadgetMatrix(parameters.GetP(), parameters.GetQ(), parameters.GetSmallN(), parameters.GetBigN(), parameters.GetL1(), parameters.GetL2()), utils.MultiplyMatrices(c1, GinvC2))

	for i := 0; i < len(Cnand[0]); i++ {
		Cnand[0][i] = utils.UnsignedMod(Cnand[0][i], float64(parameters.GetP()))
	}
	for i := 1; i < len(Cnand); i++ {
		for j := 0; j < len(Cnand[i]); j++ {
			Cnand[i][j] = utils.UnsignedMod(Cnand[i][j], float64(parameters.GetQ()))
		}
	}

	return CipherText{
		C : Cnand,
	}
}