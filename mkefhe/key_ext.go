package mkefhe

import (
	"mkefhe_lwr/utils"
)

type ExtendedPublicKey struct {
	matrix [][]float64 // m x (n+1) matrix
}
func (epk ExtendedPublicKey) GetMatrix() [][]float64 {
	return epk.matrix
}


func KeyExt(pkj []PublicKey, pp PublicParams) ExtendedPublicKey {
    m := pp.GetM()
    n := pp.GetSmallN()
    p := pp.GetP()
    q := pp.q
    Abar := pp.GetA() // m x n


	bBar := make([][]float64, m)
	for i:=0; i<m; i++ {
		bBar[i] = make([]float64, 1)
		sum := float64(0)
		for _, pkj := range pkj {
			sum = sum + pkj.GetBJ()[i][0]
		}
		bBar[i][0] = utils.SignedMod(sum, float64(p))
	}

    epkMatrix := make([][]float64, m)
    for i := 0; i < m; i++ {
        row := make([]float64, n+1)
        row[0] = bBar[i][0]
        for j := 0; j < n; j++ {
            // row[j+1] = Abar[i][j]
            row[j+1] = utils.SignedMod(Abar[i][j], float64(q))
        }
        epkMatrix[i] = row
    }
    epk := ExtendedPublicKey{
		matrix: epkMatrix,
	}

    return epk
}
