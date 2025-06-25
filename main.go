package main

import (
	"fmt"
	"math/rand"
	"mkefhe_lwr/mkefhe"
	"time"
)

func test(parameters mkefhe.PublicParams, sk1, sk2 mkefhe.SecretKey, epk mkefhe.ExtendedPublicKey, m1, m2 uint8) {
	fmt.Println("Message 1 : ", m1)
	fmt.Println("Message 2 : ", m2)

	C1 := mkefhe.Encrypt(epk, m1, parameters)
	C2 := mkefhe.Encrypt(epk, m2, parameters)

	fmt.Println("Messages Encrypted")

	// Please replace the operation below with the desired operation

	// C := mkefhe.CipherAdd(C1, C2, parameters)
	C := mkefhe.CipherMult(C1, C2, parameters)
	// C := mkefhe.CipherNand(C1, C2, parameters)

	fmt.Println("Ciphertext Evaluated")

	pd1 := mkefhe.PartialDecrypt(C, sk1, parameters)
	pd2 := mkefhe.PartialDecrypt(C, sk2, parameters)

	fmt.Println("Partial Decryptions Done")

	pd_arr := []float64{pd1, pd2}
	result := (mkefhe.Decrypt(C, pd_arr, parameters))

	// Please replace the expected value below with the expected value for the operation selected previously

	// expected := ((m1 + m2) % 2)
	expected := m1 * m2
	// expected := 1-(m1 * m2) 

	fmt.Println("Decryption Done")

	fmt.Println("Observed result: ", result)
	fmt.Println("Expected result: ", expected)
	if (result) != expected {
		panic(fmt.Sprintf("Test failed: expected %d, got %d\n", expected, result))
	}
}

func main() {
	parameters := mkefhe.DefaultParams()

	pk1, sk1 := mkefhe.KeyGen(parameters)
	pk2, sk2 := mkefhe.KeyGen(parameters)

	R0 := make([][]float64, parameters.GetM())
	for i := 0; i < parameters.GetM(); i++ {
		R0[i] = make([]float64, parameters.GetBigN())
	}

	fmt.Println("Key Generated")

	pk_arr := []mkefhe.PublicKey{pk1, pk2}
	epk := mkefhe.KeyExt(pk_arr, parameters)

	fmt.Println("Extended Public Key Generated")

	number_of_tests := 100
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < number_of_tests; i++ {
		m1 := uint8(r.Intn(2)) // 0 or 1
		m2 := uint8(r.Intn(2)) // 0 or 1
		// m1 := uint8(1)
		// m2 := uint8(1)
		test(parameters, sk1, sk2, epk, m1, m2)
	}

}
