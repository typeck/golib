package test

import (
	"fmt"
	"strconv"
	"testing"
	"unsafe"
)

//字符转换
func TestStrconv(t *testing.T) {
	// ParseInt 将字符串转换为 int 类型
	// s：要转换的字符串
	// base：进位制（2 进制到 36 进制）
	// bitSize：指定整数类型（0:int、8:int8、16:int16、32:int32、64:int64）
	// 返回转换后的结果和转换时遇到的错误
	// 如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
	res, err := strconv.ParseInt("1111", 2, 32)
	fmt.Println(res, err)
	res, err = strconv.ParseInt("0b1111", 0, 32)
	fmt.Println(res, err)
}

//数组操作（go中尽量不要用数组）
func TestArry(t *testing.T) {
	var a = [2]int{1,2}
	fmt.Printf("%x\n", unsafe.Pointer(&a))
	//数组是深拷贝，（和c数组指针传递不一样）
	arryOp(a)
	fmt.Println("a[1]=", a[1])
}
func arryOp(a [2]int) {
	fmt.Printf("%x\n", unsafe.Pointer(&a))
	a[1] = -1
}