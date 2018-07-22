package crypt

import (
	"fmt"
	"blockchainGo/crypt/hash"
	"blockchainGo/crypt/des"
)


func CryptMain()  {

	fmt.Println("=======  Welcome to cryto function test BEGIN ========")
	hash.HashMain()

	des.DesCryptMain();

	fmt.Println("======= cryto function test End ========")

}

