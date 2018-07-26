package rsa

import "fmt"

func RSACryptMain()  {


	fmt.Println("-------- RSA crypto main ------")
	//加密
	cipher:=RSAEncrypt([]byte("hello, Welcome to RSA crypt demo "))
	fmt.Println(cipher)

	//解密
	plain:=RSADecrypt(cipher)
	fmt.Println(string(plain))

}
