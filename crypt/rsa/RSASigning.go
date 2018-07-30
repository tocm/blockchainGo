package rsa

import (
	"crypto/md5"
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto"
	"crypto/rand"
	"fmt"
)

//--------------RSA 非对称签名运算 ----------------//

const RSA_SIGN_KEY_AUTHOR  = "Andy.chen" //签名作者

/**
	RSA 签名 使用私钥
 */
func RSASign() []byte {
	fmt.Println("todo RSA 数字签名")
	//解释私钥字节, 将字节数组转换成私钥类型
	block,_:=pem.Decode(RSA_private_key) //使用openssl工具生成的私钥
	priv,_:=x509.ParsePKCS1PrivateKey(block.Bytes)

	//签名
	plaintxt:=[]byte(RSA_SIGN_KEY_AUTHOR)//确定签名人
	h:=md5.New()
	h.Write(plaintxt)
	hashed:=h.Sum(nil)

	opts:=&rsa.PSSOptions{rsa.PSSSaltLengthAuto,crypto.MD5}
	sig,_:=rsa.SignPSS(rand.Reader,priv,crypto.MD5,hashed,opts)

	//返回签名结果
	return sig
}

/**
	验证签名： 使用公钥
 */
func RSAVerifySign(sign []byte) bool {
	fmt.Println("todo RSA 验证数字签名")
	//解释公钥字节
	block,_:=pem.Decode(RSA_public_key)//使用openssl工具生成的公钥
	pubInterface,_:=x509.ParsePKIXPublicKey(block.Bytes)
	pub:=pubInterface.(*rsa.PublicKey)

	//验证发送方是否来自签名作者
	h:=md5.New()
	h.Write([]byte(RSA_SIGN_KEY_AUTHOR))
	hashed:=h.Sum(nil)

	e:=rsa.VerifyPSS(pub,crypto.MD5,hashed,sign,nil)
	if e==nil {
		fmt.Printf("接受数据成功，数据来自于作者<%s >发送的，没有受到被攻击",RSA_SIGN_KEY_AUTHOR)
		return true
	}
	return false
}