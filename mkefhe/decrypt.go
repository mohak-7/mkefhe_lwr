package mkefhe

import (
	// "fmt"
	"math"
	"mkefhe_lwr/utils"
)

func PartialDecrypt(C CipherText, sk SecretKey, pp PublicParams) float64 {
	wT := make([][]float64, pp.n+1)
	for i:=0;i<pp.n+1;i++ {
		wT[i] = make([]float64, 1)
	}
	wT[0][0] = (math.Ceil(float64(pp.GetP())/2.0))

	C1 := C.GetC1() // n x N matrix

	GinvwT := utils.Ginv(wT, pp.GetL1(), pp.GetL2(), pp.GetSmallN())
	si := sk.GetSI()


	pi := utils.MultiplyMatrices(utils.Transpose(si), utils.MultiplyMatrices(C1, GinvwT))

	return pi[0][0] + 1
}

func Decrypt(C CipherText, partialDecryptions []float64, pp PublicParams) uint8{
	c0 := C.GetC0() 

	wT := make([][]float64, pp.n+1)
	for i:=0;i<pp.n+1;i++ {
		wT[i] = make([]float64, 1)
	}
	wT[0][0] = (math.Ceil(float64(pp.GetP())/2.0))

	Np := len(partialDecryptions)

	p_ := 0
	for i := 0; i < Np; i++ {
		p_ += int(partialDecryptions[i])
	}

	GinvwT := utils.Ginv(wT, pp.GetL1(), pp.GetL2(), pp.GetSmallN())

	v := utils.MultiplyMatrices(c0, GinvwT)[0][0] - float64(pp.p)*float64(p_)/float64(pp.q)

	temp := utils.SignedMod(v, float64(pp.p))
	msg := float64(temp)/math.Ceil(float64(pp.p)/2.0)

	// return msg

	// result := math.Round(msg)
	// result := uint64(math.Round(msg))%2

	result := math.Abs(utils.SignedMod(math.Round(msg), 2.0))
	return uint8(result)
}