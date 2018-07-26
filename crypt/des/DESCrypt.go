package des

import (
	"bytes"
	"crypto/des"
	"crypto/cipher"
	"fmt"
)


// ========== 对称加密/解密运算 ============== //

 
/**
	加码
	orgData: 原始传进来明文: 比如：传进来 abc
	blockSize：设置指定补码后总长度    比如： 5 即：5-减去原来的明文长度需要补多少位
	[]byte: 返回补码后的字节数组
 */
func PKCS5Padding(orgData []byte, blockSize int) []byte  {
	//比如：abc22
	padding := blockSize - len(orgData)%blockSize //原始长度取余得到需要补的长度
	padtxt := bytes.Repeat([]byte{byte(padding)}, padding) //将需要补的长度，以相应的长度值为准重复补上多少位。如：2位即补 22
	return append(orgData, padtxt...) //然后添加到原始明文后面
}

/**
	去码，根据补码的逻辑，反向取出原码
 */
func PKCS5UnPadding(cipherTxt []byte) []byte {
	length :=len(cipherTxt)
	unpadding:=int(cipherTxt[length-1])
	orgEndLen := length-unpadding
	return cipherTxt[:orgEndLen]  //数组切片 [:length-unpadding]相当于取从第0位到第4位的元素返回新数组

}

/**
	DES加密，加密会用到加码
	orig: 明文
	key: 指定密钥 它长度必须为8byte
 */
func DESEnCrypt(orig[] byte, key[] byte) []byte {
	//首先检验秘钥是否合法
	//DES加密算法，秘钥的长度必须为8byte
	block, _ := des.NewCipher(key)
	fmt.Println("EnCrypt key size = ", block.BlockSize())
	//补码
	origData := PKCS5Padding(orig, block.BlockSize())
	//设置加密方式
	blockMode := cipher.NewCBCEncrypter(block, key)
	//加密处理
	crypted := make([]byte, len(origData)) //创建一个存放加密后的密文的数组
	blockMode.CryptBlocks(crypted, origData)
	return crypted
}

/*
	DES解密，解密要用到去码
	cipherTxt: 加密后的码
	key: 密钥 长度必须为8位
 */
func DESDeCrypt(cipherTxt []byte, key []byte) []byte {
	//校验key的有效性
	block, _ := des.NewCipher(key)
	//设置解码方式
	blockMode := cipher.NewCBCDecrypter(block, key)
	//创建缓冲，存放解密后的数据
	orgData := make([]byte, len(cipherTxt))
	//开始解密
	blockMode.CryptBlocks(orgData, cipherTxt)
	//去掉编码
	orgData = PKCS5UnPadding(orgData)
	return orgData
}

