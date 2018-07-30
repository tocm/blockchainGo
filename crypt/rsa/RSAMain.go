package rsa

import "fmt"

func RSACryptMain()  {


	fmt.Println("-------- RSA crypto main ------")

	//加密
	cipher:=RSAEncrypt([]byte("hello, Welcome to RSA crypt demo " + string(RSA_public_key)))
	fmt.Println(cipher)

	//解密
	plain:=RSADecrypt(cipher)
	fmt.Println(string(plain))


	//fmt.Println("---- RSA 通过go api 生成公私钥 demo ----")
	//privKey,pubKey := RSAGenerateCryptKeyByApi()
	//fmt.Println("---- RSA 加密 demo ----")
	//cipherData := RSAEnCryptByApi(pubKey, []byte("hello,this is RSA crypto demo RSAGenerateCryptKeyByApi"))
	//fmt.Println()
	//RSADeCryptByApi(privKey, cipherData)
	//
	//fmt.Println("---- RSA 签名与验证签名 demo ----")
	//sigedData := RSAEnSigining(privKey)
	//fmt.Println(sigedData)
	//RSAVerifySigining(pubKey, sigedData)
}
