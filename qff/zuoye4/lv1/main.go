package main

import (
	"fmt"
	"sort"
)

func main() {
	i := 0
	k := 0
	var n int
	m := make(map[int]int)
	fmt.Scanf("%d\n", &n)
	s := make([]int, n) //制造一个切片来存放输入的数
	for i = 0; i < n; i++ {
		fmt.Scanf("%d\n", &s[i])
		m[i] = s[i]
	}
	//for range 寻找每个数的个数
	var j = 0
	x := make([]int, n) //新建一个切片来存放不同大小的数
	for i = 0; i < n; i++ {
		//外层所有的数
		for j = 0; j <= i; j++ {
			//遍历数组寻找不重复的数
			if s[i] == s[j] && i != j {
				n--
				break
			}
			if i == j {
				x[k] = s[i]
				k++
			}
		}
	}
	n = n - 2
	c := make([]int, n)
	for k = 0; k < n; k++ {
		c[k] = x[k]
	}

	//给切片中的数排序

	sort.Ints(c)

	//打印
	for k = 0; k < n; k++ {
		i = 0
		for _, v := range m {
			if v == c[k] {
				i++
			}
		}
		fmt.Printf("%d %d\n", c[k], i)
	}

}
