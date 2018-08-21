package base

import (
	"fmt"
	"bytes"
	"encoding/binary"
)

//注意公用方法要前面大写
func CalMax(num1 int, num2 int) int{
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

/*
	32位整型转4个字节byte数组
 */
func IntToBytesBigendian(numb int32) []byte {

	//for int(4byte) is big-endian size
	numbBytes := make([]byte, 4)
	numbBytes[0] = byte((numb >> 24) & 0xFF);   //高位为0
	numbBytes[1] = byte((numb >> 16) & 0xFF);
	numbBytes[2] = byte((numb >> 8) & 0xFF);
	numbBytes[3] = byte(numb & 0xFF); //低位为值

	return numbBytes;

}

func BytesToInt32Bigendian(numbBytes[]byte) int32 {
	//byte 转成 int 型
	numbInt := ((int32)(numbBytes[3] & 0xff)) | ((int32)(numbBytes[2] & 0xff) << 8) | ((int32)(numbBytes[1] & 0xff) << 16) | ((int32)(numbBytes[0] & 0xff) << 24)
	return numbInt

}


//字节转换成整形
func BytesToInt(b []byte) int {
	var tmp int32
	fmt.Println("----BytesToInt ---",b)
	bytesBuffer := bytes.NewBuffer(b)
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}