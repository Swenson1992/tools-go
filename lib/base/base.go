package base

import (
	"fmt"
	"strings"
)

func Base() {
	// 由于 Go 语言的字符串都以 UTF-8 格式保存，每个中文占用 3 个字节
	s := strings.Index("飞雪无情", "无")
	fmt.Println(s, len("飞"))

	// if 条件语句
	if a := 10; a > 10 {
		fmt.Println("i>10")
	} else if a == 10 {
		fmt.Println("i=10")
	} else {
		fmt.Println("i<10")
	}

	// switch 语句
	// switch 在默认情况下，case 最后自带 break
	switch i := 1; i {
	case 1:
		fallthrough
	case 2:
		fmt.Println("1")
	default:
		fmt.Println("没有匹配")
	}

	// for 循环语句
	sum := 0
	for j := 1; j <= 100; j++ {
		sum += j
	}
	fmt.Println("the sum is ", sum)

	//使用for模拟while
	total := 0
	index := 1
	for index <= 100 {
		total += index
		index++
	}
	fmt.Println("the sum is ", total)

	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array, len(array), array[2])
	arr := [...]int{1, 2, 3}
	fmt.Println(arr, len(arr))

	for k := 0; k < 5; k++ {
		fmt.Printf("数组索引:%d,对应值:%s\n", k, array[k])
	}
	// 数组循环
	for i, v := range arr {
		fmt.Printf("数组索引:%d,对应值:%d\n", i, v)
	}

	slice := array[2:5]
	fmt.Println(slice, len(slice))

	slice1 := make([]string, 4, 8)
	fmt.Println(slice1)

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice2), cap(slice2))

	nameAgeMap := map[string]int{"宋健": 28}
	//添加键值对或者更新对应 Key 的 Value
	nameAgeMap["飞雪无情"] = 20

	//获取指定 Key 对应的 Value
	age, ok := nameAgeMap["飞雪无情"]
	if ok {
		fmt.Println(age)
	}

	delete(nameAgeMap, "飞雪无情")
	fmt.Println(nameAgeMap, len(nameAgeMap))

	nameSexMap := make(map[string]string)
	nameSexMap["宋健"] = "男"
	fmt.Println(nameSexMap, len(nameSexMap))

	ss := "Hello飞雪无情"
	bs := []byte(ss)
	fmt.Println(bs)
	fmt.Println(ss[0], ss[1], ss[15])

	for i, r := range ss {
		fmt.Println(i, r)
	}

	aa := [3][3]int{}
	aa[0][0] = 1

	fmt.Println(aa)
}
