package rsa

import (
	"encoding/pem"
	"crypto/x509"
	os_rsa "crypto/rsa"
	"crypto/rand"
)

//-------------- 非对称加密运算 ----------------//

var RSA_public_key = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC+O+5SgN/W7AYnms4exT4ZFaxw
/TBb6nYDysi9mWAbTVRd1HlFdVnLam41H/hdxOqKxhuR6fKXKnsG8Ud9g19Vx4kQ
Z7+I1YTsVs66A1yQwtBOi+v+Yu7uvsM86EoS6IKqjrLbRtKlf5uot9Q/xuyEG90p
9YrdqIJJM9b/dU6NDQIDAQAB
-----END PUBLIC KEY-----`)


var RSA_private_key  = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC+O+5SgN/W7AYnms4exT4ZFaxw/TBb6nYDysi9mWAbTVRd1HlF
dVnLam41H/hdxOqKxhuR6fKXKnsG8Ud9g19Vx4kQZ7+I1YTsVs66A1yQwtBOi+v+
Yu7uvsM86EoS6IKqjrLbRtKlf5uot9Q/xuyEG90p9YrdqIJJM9b/dU6NDQIDAQAB
AoGAKXHZHDRjj+lf9eTuHxoNy11DcmJW1M+h4nWdgmEV5RzV0pNl0FIevAh19jD+
5T7vckhy3TlA1to0b3DFiHzEmgUwjVt1jm0dwgCFm4YHUf3cGmkNGx3oX7vaOkWs
CcTqhAMANpTdk/hzsoqDVtM4OGIPn1eNm0rtQK4K0MCRw8ECQQDts2OMN44J0lc9
NsEPYeSawy5qxF1BczdmJLBiuGg1opK3FQsaGwOOzt0S1m0t0KRAi37m3KViQzDw
gTvemqDFAkEAzOERlhdPxTQoehkjGCVqnsepPCpjOCfPzFaTsgdwX9HpfBltisYI
dkZE6salntPgPwIJBg6qlaEbAEBOTcFvqQJAe0JKnKsyPGWWV2fNNOzOXnJX34Vz
1jyovY7I6Gg3oZhr/AQtfZwEfElI2jvW4asPgWjUpWJFYG73Mb/SC9yjAQJBAIOa
5+MT4mf40y9rSkHOdyi7dJhzWfzOhgrqakFnyeWjP2o6I0WLZBAln0t4gxb9XwTa
V9WOQkGPlXqZ1DLaWNECQQDh+a9HpbXfCbb0buuJyfxfjbx9BaDxNFMpUlkHX0oL
zFvJWHLEhSNYY84D3+zXS4l0rbxqyAwXeCWKwH8EX2kS
-----END RSA PRIVATE KEY-----`)


/**
	RSA加密
 */
func RSAEncrypt(orgData []byte) []byte {

	//通过公钥加密
	block,_ := pem.Decode(RSA_public_key)
	//解释公钥
	publicInterface,_ := x509.ParsePKIXPublicKey(block.Bytes)
	//加截公钥
	pub := publicInterface.(*os_rsa.PublicKey)

	//利用公钥pub加密
	bits,_ := os_rsa.EncryptPKCS1v15(rand.Reader,pub,orgData)
	//返回密文
	return bits
}

//RSA解密
func RSADecrypt(cipherTxt []byte) []byte {
	//通过私钥解密
	block,_:=pem.Decode(RSA_private_key)
	//解析私钥
	pri,_:=x509.ParsePKCS1PrivateKey(block.Bytes)
	//解密
	bits,_:=os_rsa.DecryptPKCS1v15(rand.Reader,pri,cipherTxt)
	//返回明文
	return bits

}