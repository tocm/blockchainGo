package rsa

import (
	"encoding/pem"
	"crypto/x509"
	os_rsa "crypto/rsa"
	"crypto/rand"
)

//--------------RSA 非对称加密运算 ----------------//
//---------使用openssl 工具预生成好private key and public key------//

//RSA生成私钥是4086bit长度，所以加密的明文必须少于4086bit
var RSA_public_key = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAx+JqpKZVafq4kzSiQzmw
dYoGvBFo813pSZktMlI4P0VSS6YwIYhP4SxfTDFUjBd344J3Gnbl/Eqln9BhE5ZR
+vz0zXftERQVKgO9KlP8go1B8ZDuR4XPdqJfT/1h5f7JenHqKSMQEbc+7LX6nDj6
whYy+8IFnWFla1EoPu7PymYTum1nFV1ymVKXgokKKQS+p86NN0nwyCFTthzp1+Mp
Ka8FqA0HgvZloAlRebiaafaB9vQqQEr0E9dkFpwEBusxKc3AAz+KDUG3i0TZIbCq
T/PfWlhxGXLmHqRD99R27/fgmPmH0+DljabJI/wQHxR20iW/sugqSudJq6ua7MSN
mwIDAQAB
-----END PUBLIC KEY-----`)

//RSA生成私钥是2048bit长度，所以需要加密的明文必须少于2048bit
var RSA_private_key  = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAx+JqpKZVafq4kzSiQzmwdYoGvBFo813pSZktMlI4P0VSS6Yw
IYhP4SxfTDFUjBd344J3Gnbl/Eqln9BhE5ZR+vz0zXftERQVKgO9KlP8go1B8ZDu
R4XPdqJfT/1h5f7JenHqKSMQEbc+7LX6nDj6whYy+8IFnWFla1EoPu7PymYTum1n
FV1ymVKXgokKKQS+p86NN0nwyCFTthzp1+MpKa8FqA0HgvZloAlRebiaafaB9vQq
QEr0E9dkFpwEBusxKc3AAz+KDUG3i0TZIbCqT/PfWlhxGXLmHqRD99R27/fgmPmH
0+DljabJI/wQHxR20iW/sugqSudJq6ua7MSNmwIDAQABAoIBACHChxqKEXEV7KjI
QgOysbAYgOyw0BYSrX1FxtFx+5gMIEswpX50ZHF2qh3XUx1aNWCSv4UTweJrxS+J
5WuGfxAyOzKc58hzqONwwDlu9/4g15tyTBx+aDz0Tb7X3/oVn3fz14XdKzCiz5jx
X4OVSB0R3MiSy5n6QkPu8Co48ZWQfaNyOvnOQ36TEyBRT9MhY+IcE0cO4/vJDwCx
3uqmMWPCIon13Gg3IWs/unOnORsZh/GiXFiwlVX/cr5pEz3zQJc7ZJl+eRrfni3p
Vkcmp1kmHqM8//zW5H30tvNAcG7ujafk79tkHyHdouaaf6RL0IDLCAzaNr8KDaQs
gyCGieECgYEA54vEk8avOz8d1npqafHji8yfJ8ws9/GgU0mMaXe7KcJfL7+x6TeT
9mQxhkI1PMUKT23MT5uVOniOymOaWZL+fLWlHMOShjfwjzmoLINqOitbkDJMNpJF
zUZRMbgnZNIrp/P02b356L2Ik4E4JLi1qD4J3mICXhY5kZjsyPv4PEsCgYEA3P6g
WdcPC3W2/iYm/qiQ29C1kuhMnTDFKeu5JU9c75vGwEf01t/f7wFZhhRX8Ajpcbd5
K93WOAJGqeRRoEyYn5TWp2/84lAzXiu9dwUr8gJg2qGLs/mdU/COuDyXUJ8ygFR5
XWK1sL8UCetC4no0DG5UPvB+17etokge1tHWgfECgYEAhT0YUTDzhNQ9Bz5KHczj
c3671trWxZZ0+3hHJCmRueXvgDp+1gT1T8M3/Fn86Is585rrBqQwqwTUpgHKv+Cm
MH0WFgc0OafW9ZPoAjVc2zw35DINSRnCsOzVrBacFJgcJvlq3CFFvXfTL44cEIV5
guawJxee5SVE21N0bfMQVSECgYEAlAeva15mAnwgTo9J0Y0iM0vTRIt8PlLbzt/K
DhbiHtIOyN+qPoCJpLQXvntECFI/5N9MODkJdEz5/nUkxCZqmYJ+YnvePc4RCoRs
DhyVui8jp29MgAq0+vQI9ZoIlYV9s35iKU6ke7cSei5vh/rpgGYq4ln/7QlpMmny
1lyKvlECgYA1cvufK/vd+saZHVgGJYTXyLot7P9C/q0yyLjs1Wp/CNnJVFfGByQP
Nbumex7gKa1JuK28vOaZXu/gsjBdT31Z8j6XIOHpa/dP7elbdQNzWBQ0Cx1i0D2Y
EZbLh/GyBnM8Gw761C6R6f5+tqyuVW7kFyErqx6y5OGHQugyEWnBsA==
-----END RSA PRIVATE KEY-----`)


/**
	RSA加密
	使用openssl 工具预生成好private key and public key。
	orgData[]： 需要加密的明文大小必须是少于生成RSA 私钥时的字节长度
 */
func RSAEncrypt(orgData []byte) []byte {

	//通过公钥加密
	block,_ := pem.Decode(RSA_public_key)
	//解释公钥
	publicInterface,_ := x509.ParsePKIXPublicKey(block.Bytes)
	//加截公钥
	pub := publicInterface.(*os_rsa.PublicKey)

	//利用公钥pub加密，明文字节大小必须是少于生成RSA 私钥时的字节长度
	bits,_ := os_rsa.EncryptPKCS1v15(rand.Reader,pub,orgData)
	//返回密文
	return bits
}

//RSA解密
//使用openssl 工具预生成好private key and public key
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