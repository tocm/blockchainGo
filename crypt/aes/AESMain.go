package aes

import "fmt"

// ========== 对称加密/解密运算 ============== //



func AESEncryptMain()  {

	fmt.Println("======= AES encrypt ======")
	encryptData := AESEncrypt([]byte("hello. welcome to AES encrypt demo"), []byte(AES_KEY_ENCRYPT))
	fmt.Println("AES encrypt data = ",encryptData)

	deEncryptData := AESDecrypt(encryptData, []byte(AES_KEY_ENCRYPT))
	fmt.Println("AES decrypt data =",string(deEncryptData))
}
