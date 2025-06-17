package mkefhe

import (
	// "fmt"
	"math"
	"mkefhe_lwr/utils"
)

type PublicKey struct {
    Matrix [][]float64
    bj [][]float64 // m x 1 vector
}

type SecretKey struct {
    // si64  [][]float64
	Si    [][]float64
    Sk    [][]float64
}

func (pk PublicKey) GetPK() [][]float64 {
    return pk.Matrix
}
func (pk PublicKey) GetBJ() [][]float64 {
    return pk.bj
}
func (sk SecretKey) GetSK() [][]float64 {
    return sk.Sk
}
// func (sk SecretKey) GetSI64() [][]float64 {
//     return sk.si64
// }
func (sk SecretKey) GetSI() [][]float64 {
    return sk.Si
}


func KeyGen(pp PublicParams) (pkj PublicKey, skj SecretKey) {
    n := pp.GetSmallN()
    m := pp.GetM()
    q := pp.GetQ()
    p := pp.GetP()
    Abar := pp.GetA() // m x n matrix

    // Generate a random binary vector sj of length n
    sj := utils.SampleUniformMatrix(n, 1, 2) // [][]byte of length n x 1

    // sj := skk

    // fmt.Println("sj", sj)

    // fmt.Println("A : ", Abar)

    // fmt.Println("sj: ", sj)
	// fmt.Println("----------------------------------")

    // Convert sj to a slice of float64 for calculations
    // sj64 := make([][]float64, n)
    // for i, v := range sj {
    //     sj64[i] = make([]float64, 1)
    //     sj64[i][0] = float64(v[0])
    // }

    // fmt.Println("sj64: ", sj64)
	// fmt.Println("----------------------------------")

    // Asj = Abar * sj64 mod q
    // Abar is m x n, sj64 is n x 1, so Asj will be m x 1
    // Asj := utils.MultiplyMatricesMod(Abar, sj, float64(q)) 
    Asj := utils.MultiplyMatrices(Abar, sj) // m x 1 matrix

    // fmt.Println("Asj: ", Asj)

    // rAsj, cAsj := len(Asj), len(Asj[0])
    // fmt.Println("Asj shape: ", rAsj, cAsj)
    // fmt.Println("Asj: ", Asj)
	// fmt.Println("----------------------------------")

    // bj = round(p * Asj / q) mod p
    bj := make([][]float64, m)
    for i := 0; i < m; i++ {
        bj[i] = make([]float64, 1)
        scaled := math.Round(float64(p)* float64(Asj[i][0]) / float64(q) )
        bj[i][0] = utils.SignedMod(scaled, float64(p)) 
    }

    // fmt.Println("bj: ", bj)



    // rBj, cBj := len(bj), len(bj[0])
    // fmt.Println("bj shape: ", rBj, cBj)
    // fmt.Println("bj: ", bj)
	// fmt.Println("----------------------------------")

    // Public Key pk = [bj | Abar]
    // bj is m x 1, Abar is m x n, so pk will be m x (n+1)
    pkMatrix := make([][]float64, m)
    for i := 0; i < m; i++ {
        row := make([]float64, n+1)
        row[0] = bj[i][0]
        for j := 0; j < n; j++ {
            row[j+1] = Abar[i][j]
        }
        pkMatrix[i] = row
    }
    pkj = PublicKey{
        Matrix: pkMatrix,
        bj: bj,
    }
    // fmt.Println("pk shape: ", len(pkj.Matrix), len(pkj.Matrix[0]))
    // fmt.Println("pk: ", pkj)
	// fmt.Println("----------------------------------")

    // skVector := make([]float64, n+1)
    // skVector[0] = 1.0
    // for i := 0; i < n; i++ {
    //     skVector[i+1] = -float64(p) * float64(sj[i][0])/ float64(q) 
    // }

    // Secret key sk = [1 | -p * sj64 / q]
   sk := make([][]float64, n+1)
   sk[0] = make([]float64, 1)
   sk[0][0] = 1.0
   for i := 0; i < n; i++ {
       sk[i+1] = make([]float64, 1)
       sk[i+1][0] = -float64(p) * float64(sj[i][0]) / float64(q)
   }

    // rSk, cSk := len(sk), len(sk[0])
    // fmt.Println("sk shape: ", rSk, cSk)
    // fmt.Println("sk: ", sk)
	// fmt.Println("----------------------------------")


    skj = SecretKey{
        // si64 : sj64,
		Si : sj,
        Sk : sk,
        // Vector: skVector,
    }

    // pksk := (utils.MultiplyPkSk(pkMatrix, sk, float64(pp.GetP())))
	// fmt.Println("dimension of pksk: ", len(pksk), "x", len(pksk[0]))
	// fmt.Println("pksk: ", pksk)
	// fmt.Println("----------------------------------")

    // // now we need to find (bj - p/q * Asj) mod p
    // tempvar := make([][]float64, m)
    // for i := 0; i < m; i++ {
    //     tempvar[i] = make([]float64, 1)
    //     tempvar[i][0] = utils.SignedMod(bj[i][0] - float64(p)*float64(Asj[i][0])/float64(q), float64(p))
    // }
    // fmt.Println("tempvar shape: ", len(tempvar), "x", len(tempvar[0]))
    // fmt.Println("tempvar: ", tempvar)
    // fmt.Println("----------------------------------")

    // // Check if pksk is equal to tempvar
    // for i := 0; i < m; i++ {
    //     if math.Abs(pksk[i][0] - tempvar[i][0]) > 1e-9 { // Using a small epsilon for float comparison
    //         fmt.Println("Error: pksk does not match tempvar at index", i)
    //         fmt.Println("pksk[i][0]:", pksk[i][0], "tempvar[i][0]:", tempvar[i][0])
    //         return PublicKey{}, SecretKey{} // Return empty keys on error
    //     }
    // }
    // fmt.Println("pksk matches tempvar, key generation successful!")
    // fmt.Println("----------------------------------")


    return pkj, skj
}

