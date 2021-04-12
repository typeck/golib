package test

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	nums := getRandNums()
	quickSort(nums, 0, len(nums))
	fmt.Println(nums)
	//nums = getRandNums()
}

func getRandNums()[]int {
	return []int {1,3,5,7,2,6,4,8,9,2,8,7,6,0,3,5,9,4,1,0}
}

// 通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小， 
// 分割点称为轴心，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，
// 以此达到整个数据变成有序序列。
// 由于l, r的移动方向相反， 故原处于左端较大的元素将按颠倒的次序转移至右端， 因此快排并不稳定。
// 复杂度 O(nlog(n))
// [l, r)
func quickSort(args []int, l, r int) {
	if r - l < 2 {
		return
	}
	pivot := partition(args, l, r - 1)
	quickSort(args, l, pivot)
	quickSort(args, pivot + 1, r)
}

// [l, r]
func partition(args []int, l, r int) int {
	key := args[l]
	for l < r {
		for l < r && args[r] >= key {
			r--
		}
		args[l] = args[r]
		for l < r && args[l] <= key {
			l++
		}
		args[r] = args[l]
	}
	// 最终first == last； （轴心）
    args[l] = key;
	return l
}