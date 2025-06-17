package utils

import (
	"crypto/rand"
	"encoding/binary"
	"math"
	"math/big"
)

func SampleCryptoUint(q uint64) float64 {
	bq := big.NewInt(int64(q))
	r, err := rand.Int(rand.Reader, bq)
	if err != nil {
		panic("crypto/rand failed: " + err.Error())
	}
	return float64(r.Uint64())
}

func SampleUniformVector(n int, q uint64) []float64 {
	vec := make([]float64, n)
	for i := range vec {
		vec[i] = SampleCryptoUint(q)
	}
	return vec
}

func SampleUniformMatrix(m, n int, q uint64) [][]float64 {
	mat := make([][]float64, m)
	for i := range mat {
		mat[i] = SampleUniformVector(n, q)
	}
	return mat
}

// func SampleSmallVector(n int) []int8 {
// 	vec := make([]int8, n)
// 	for i := range vec {
// 		r := SampleCryptoUint(3)  // 0 to 2
// 		vec[i] = int8(r - 1)     // -1 to 1
// 	}
// 	return vec
// }

func SampleBinaryVector(n int) []byte {
	vec := make([]byte, n)
	for i := range vec {
		vec[i] = byte(SampleCryptoUint(2))
	}
	return vec
}


func secureFloat64() float64 {
    var b [8]byte
    _, err := rand.Read(b[:])
    if err != nil {
        panic(err)
    }
    // Use the top 53 bits for a float64 in [0,1)
    u := binary.BigEndian.Uint64(b[:]) >> 11
    return float64(u) / (1 << 53)
}

// SecureDiscreteGaussian samples an integer from a discrete Gaussian with mean 0 and stddev sigma.
func SecureDiscreteGaussian(sigma float64) float64 {
    // Box-Muller transform for normal distribution
    u1 := secureFloat64()
    u2 := secureFloat64()
    z := math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2*math.Pi*u2)
    // Round to nearest integer for discrete output
    return math.Round(z * sigma)
}