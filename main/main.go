package main


import (
	"fmt"
	"math"
	"blockchainGo/oop"
	baseBc "blockchainGo/utils"

	"blockchainGo/concurrency"
	"blockchainGo/crypt"
)
//常量
const ARRAY_MAX int = 5

//数组初始化
var arraysData = [ARRAY_MAX] int{1,2,3,4,5}
//切片初始化
var slcieData = [] int{6,7,8,9,0}

//此处传递数组参数定义为固定数组，所以不能传非固定长度数组
func printArrays(pArrays [ARRAY_MAX]int, arrayLen int){
	for i:=0; i< arrayLen; i++ {
		fmt.Println("测试数组： %d,", pArrays[i]);
	}

}

//	array是值传递,所以修改后是不生效
func modifyArrays(pArray[ARRAY_MAX] int){
	fmt.Println("preper modify array [1] = 88")
	newArrays := pArray;
	newArrays[0] = 88;//值传递，将不会改变原来数组的元素
}

//切片
func printSlice(pSclie []int, len int){
	for i:=0; i< len; i++ {
		fmt.Println("测试切片： %d,", pSclie[i]);
	}
}
//切片修改，任何时候都是引用传递
func modifySlice(pSlice[] int){
	fmt.Println("preper modify slice [1] = 88")
	newSlice := pSlice;
	newSlice[0] = 88;
}

func printByRange(sliceData[] int){
	fmt.Println("Text Range")
	for index, value := range sliceData{
		fmt.Println(index,value);
	}
}
//函数 测试值传递
func calMaxValue(n1 int, n2 int) int{
	if(n1 > n2) {
		return n1;
	} else {
		return n2;
	}
}

//函数 测试引用传递指针
func calMinValue(p1 *int, p2 *int){
	var temp int;
	if(*p1 < *p2){
		temp = *p1;
		*p1 = *p2;
		*p2 = temp;
	}
}

/*函数 闭包 注意func无参数，
	调用：
	funcSeq := getSequence(1,2) //代表getSequence()
	fmt.Println(funcSeq())//代表func()
*/
func getSequence(x1, x2 int) func() (int, int)  {
	i:= 0
	return func() (int,int) {
		i++
		return i, x1+x2
	}
}



func testMap() {
	//方法一：初始化
	var phone = map[string] int64{"iphone8": 5555, "iphone6": 3698}
	for key,value := range phone{
		fmt.Println(key,value);
	}

	//方法二：make
	mapData := make(map[string] string)
	mapData["china"] = "cn"
	mapData["America"] = "usa"
	for key,value := range mapData{
		fmt.Println(key, value);
	}
}
//测试跨包，结构类
func testOOP() {
	p := new(oop.Person)
	p.SetInfo("andy", 35, "gz")
	fmt.Println("this name is ",p.GetInfo())

}

func testInterface()  {
	var phone oop.Phone;
	phone = new(oop.Iphone)
	phone.Call();
	uuid := phone.GetUUID();
	fmt.Println("get ipone uuid is %d ",uuid)

	phone = new(oop.AndrodPhone)
	phone.Call();
	uuid = phone.GetUUID();
	fmt.Println("get androidphone uuid is %d ",uuid)
}


func main() {
	fmt.Println("-----HelloWorld-------")

	printArrays(arraysData, ARRAY_MAX);
	modifyArrays(arraysData)
	printArrays(arraysData, ARRAY_MAX);


	max := calMaxValue(arraysData[2], arraysData[3])
	fmt.Println("max value = %d", max);

	calMinValue(&arraysData[2], &arraysData[4]);
	fmt.Println("max value = %d", arraysData[2], arraysData[4]);




	//函数作为值
	fucnAsValue := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println("Test the value as a function %d",fucnAsValue(9));

	//测试闭包
	funcSeq := getSequence(1,2)
	fmt.Println(funcSeq())

	//切片
	var slcie1 = make([]int,0)
	slcie1 = append(slcie1, 11,12,13)//追加值
	printSlice(slcie1, len(slcie1))
	modifySlice(slcie1)
	//printSlice(slcie1, len(slcie1))
	printByRange(slcie1)

	testMap();

	testOOP();

	testInterface()

	baseBc.CalMax(2, 3);

	concurrency.Entry()


	//test crypt
	crypt.CryptMain()
}







