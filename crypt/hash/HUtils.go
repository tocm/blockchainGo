package hash


/**
	简单hash，特征： 将不同的输入转成固定长度的值
 */
func SimpleHash(number int) int {
	return (number + 2 + (number << 1)) %8 ^ 5;
}

/**
	典型Hash 散列算法
    特征： 将任何长度输入字符串，通过算法，散列成 0 ~ 15 的整数值返回。
    通过hash出来的 0-15 概率是相等的
 */
func HashCode(key string) int  {
	var index int = 0
	index = int(key[0])
	for k:=0; k<len(key); k++ {
		index *= (1103515245 + int(key[k]))
	}
	index >>= 27
	index &= 16 - 1
	return index
}