package hash
import "strconv"

/**
	实现HashMap 原理
 */

const ARRAY_LEN int = 16

//创建长度为16数组，用于存放散列运算生成的 0〜15 范围的固定散列值
var buckets = make([]*LNode, ARRAY_LEN)

/* 初始化数组，以及每个数组元素默认创建一个链表头，
  链表主要存放散列运算中生成的0~15范围的冲突，
  比如两个都生成是5的散列值，这时候就把这两个值用链表来存放
*/
func InitBuckets()  {
	for i:=0; i< ARRAY_LEN; i++ {
		buckets[i] = CreateHeadNode(KValue{ "hash", strconv.Itoa(i)});
	}
}

/**
	添加用户新值，供外部调用
 */
func AddKeyValue(key string, value string)  {
	//散列算法生成 0~15 值
	hashIndex := HashCode(key);

	//获取数组对应的链表头节点
	var headNote = buckets[hashIndex]

	//先找出尾节点以确保在尾结点上添加节点
	var tailNode = TailNode(headNote)

	//添置新结点
	AddNewNode(KValue{key,value}, tailNode);

	//显示当前链表中所有节点
	ShowAllNode(headNote)
}

/**
	通过key 取出对应值
 */
func GetKeyValue(key string) string {
	hashIndex := HashCode(key)
	var lHeadNode = buckets[hashIndex]
	lnode := FindNode(KValue{key,""}, lHeadNode)
	return  lnode.Data.Value

}