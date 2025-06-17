package mkefhe

import (
	// "fmt"
	"math"
	"mkefhe_lwr/utils"
)

func PartialDecrypt(C CipherText, sk SecretKey, pp PublicParams) float64 {
	w := make([][]float64, pp.n+1)
	for i:=0;i<pp.n+1;i++ {
		w[i] = make([]float64, 1)
	}
	w[0][0] = (math.Ceil(float64(pp.GetP())/2.0))
	// wT := utils.Transpose(w)

	C1 := C.GetC1() // n x N matrix

	Ginvw := utils.Ginv(w, pp.GetL1(), pp.GetL2(), pp.GetSmallN())
	si := sk.GetSI()

	// print the dimension of si, c1, and w
	// fmt.Println("si dimension: ", len(si), "x", len(si[0]))
	// fmt.Println("C1 dimension: ", len(C1), "x", len(C1[0]))
	// fmt.Println("Ginvw dimension: ", len(Ginvw), "x", len(Ginvw[0]))

	pi := utils.MultiplyMatrices(utils.Transpose(si), utils.MultiplyMatrices(C1, Ginvw))

	// print the dimension of pi
	// fmt.Println("pi dimension: ", len(pi), "x", len(pi[0]))

    // return (pi[0][0]) + utils.SecureDiscreteGaussian(1)
	return pi[0][0] + 1
}

func Decrypt(C CipherText, partialDecryptions []float64, pp PublicParams) uint8{
	c0 := C.GetC0() 

	w := make([][]float64, pp.n+1)
	for i:=0;i<pp.n+1;i++ {
		w[i] = make([]float64, 1)
	}
	w[0][0] = (math.Ceil(float64(pp.GetP())/2.0))
	// wT := utils.Transpose(w)

	Np := len(partialDecryptions)

	p_ := 0
	for i := 0; i < Np; i++ {
		p_ += int(partialDecryptions[i])
	}

	Ginvw := utils.Ginv(w, pp.GetL1(), pp.GetL2(), pp.GetSmallN())

	v := utils.MultiplyMatrices(c0, Ginvw)[0][0] - float64(pp.p)*float64(p_)/float64(pp.q)

	// v := float64(utils.MultiplyVectors(c0, utils.GinvVector(wT, pp.p, pp.q,pp.l1,pp.l2,pp.n))) - float64(pp.p)*float64(p_)/float64(pp.q)

	// temp := uint64(v) % pp.p
	temp := utils.SignedMod(v, float64(pp.p))
	msg := float64(temp)/math.Ceil(float64(pp.p)/2.0)

	// return msg

	// result := math.Round(msg)
	// result := uint64(math.Round(msg))%2


	result := math.Abs(utils.SignedMod(math.Round(msg), 2.0))
	return uint8(result)


	// return (msg)
}