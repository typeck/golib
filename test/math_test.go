package test

import (
	"fmt"
	"testing"
)


func TestGcd(t *testing.T) {
	x, y := 18, 12
	fmt.Println(gcd(x, y))
	fmt.Println(lcm(x, y))
}

// 辗转相除法求最大公因数
// greatest common divisor，gcd
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 最小公倍数: 两个数相乘再除以最大公因数即可得到最小公倍数
// least common multiple, lcm
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

// 求x^n
//  n < 0, y = 1 / x^(-n)
// n为奇数, y = x * x^(n-1)
// n为偶数, y =  x^(2*(n/2))
func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}else if n < 0 {
		return 1/pow(x, -n)
	}else if n%2 == 0 {
		return pow(x*x, n/2)
	}else {
		return x * pow(x, n-1)
	}
}

func TestPrintNums(t *testing.T) {
	fmt.Println(printNums(3))
}


func printNums(n int) []string {
	var s []byte
	var res []string
	dfs(&s, &res, n)
	return res
}
//剑指 Offer 17. 打印从1到最大的n位数, dfs全排列解法
func dfs(s *[]byte, res *[]string, k int) {
	if len(*s) == k {
		*res = append(*res, string(*s))
		return
	}
	for i := 0; i < 10; i++ {
		*s = append(*s, '0' + byte(i))
		dfs(s, res, k)
		*s = (*s)[:len(*s)-1]
	}
}