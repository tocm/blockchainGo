package crypt

import (
	"fmt"
	"blockchainGo/crypt/hash"
	"blockchainGo/crypt/des"
	"blockchainGo/crypt/aes"
	"blockchainGo/crypt/rsa"
)


func CryptMain()  {

	fmt.Println("=======  Welcome to cryto function test BEGIN ========")
	hash.HashMain()

	fmt.Println()
	des.DesCryptMain();

	fmt.Println()
	aes.AESEncryptMain();

	fmt.Println()
	rsa.RSACryptMain()

	fmt.Println()




	fmt.Println("======= cryto function test End ========")

}

