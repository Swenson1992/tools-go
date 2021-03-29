package base

import (
	"errors"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3}
	// 改用地址的传参
	add(&arr)
	fmt.Println(arr)
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误:%w", e)
	fmt.Println(w)
	fmt.Println(errors.Unwrap(w))
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	defer fmt.Println("defer 4")
	fmt.Println(errors.Is(w, e))
}

// 这个函数不能有返回值，也不能定义全局变量， 怎么改变原切片的内容并返回
func add(arr *[]int) {
	*arr = append(*arr, 10)
}

// 实际输出  [1, 2, 3]
// 期望输出  [1, 2, 3, 10]
