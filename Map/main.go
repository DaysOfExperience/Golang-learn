package main

import (
	"fmt"
	"sort"
	"strings"
)

// map 引用类型, 声明后如果不初始化就是nil
func main() {
	// map1()
	// map2()
	// map3()
	// map4()
	// map5()
	map6()
}

func map1() {
	var m0 map[string]int // nil
	fmt.Println(m0 == nil)
	m1 := map[string]int{
		"111": 1,
		"222": 2,
	}
	var m2 = make(map[string]int, 10) // 10个元素, 不是必须的, 即使不填容量也不是nil
	m3 := make(map[string]int)
	fmt.Println(m3 == nil) // 非nil
	m2["hehe"] = 1
	m2["haha"] = 2
	fmt.Println(len(m1), m1)
	fmt.Println(len(m2), m2)

	// 判断某个key是否存在
	value, ok := m2["hehe"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("hehe is not exist")
	}
	value, ok = m2["zzz"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("zzz is not exist")
	}

	// map 遍历
	// 不能使用下标, 只能for range
	for key, value := range m2 {
		println(key, value)
	}

	// delete删除map中的元素
	_, ok = m2["hehe"]
	fmt.Println(ok)
	delete(m2, "hehe")
	_, ok = m2["hehe"]
	fmt.Println(ok)
}

func map2() {
	// map本身遍历是无序的, 如何有序遍历?
	m := make(map[string]int, 100) // 100个容量, 但是目前没有元素
	// 插入50个, 过程中不会扩容
	for i := 1; i <= 50; i++ {
		m[fmt.Sprintf("stu:%02d", i)] = i
	}
	fmt.Println(m)        // 有序
	for k, v := range m { // 无序
		fmt.Println(k, v)
	}
	sli := make([]string, 0, 100)
	for k := range m {
		sli = append(sli, k)
	}
	sort.Strings(sli) // 对这个分片进行排序
	for _, s := range sli {
		fmt.Println(s, "-", m[s])
	}
}

func map3() {
	// 切片保存的元素类型为map
	sli := []map[string]int{} // 不是nil但是切片数组中没有元素
	fmt.Println(sli == nil)
	sli = append(sli, make(map[string]int, 10)) // 添加了一个map, 此map初始容量为10, 但是没有元素
	fmt.Println(sli)
	// 1个map元素, 且map可以容纳10个
	for i := 0; i < 10; i++ {
		sli[0][fmt.Sprintf("stu-%-2d", i)] = i
	}
	sli = append(sli, make(map[string]int))
	sli = append(sli, make(map[string]int))
	sli = append(sli, make(map[string]int))
	fmt.Println(sli)

	mapSlice := make([]map[string]int, 3) // 切片元素为三个nil
	fmt.Println(mapSlice)
	for _, m := range mapSlice {
		if m == nil {
			fmt.Println("nil")
		} else {
			fmt.Println(m)
		}
	}
}

func map4() {
	// 值为切片类型的map
	m := make(map[string][]int, 10) // 0个元素
	m["hhh"] = make([]int, 3)       // "hhh" : 切片元素为3个0
	value := m["hhh"]
	fmt.Printf("type of value:%T\n", value)
	for i := 3; i <= 10; i++ {
		value = append(value, i)
	}
	fmt.Println(value)
	m["hhh"] = value // 必须赋值
	fmt.Print(m)
}

func map5() {
	// 统计一个字符串中的单词出现次数
	s := "son of a bitch bullshit what is wrong with you"
	sli := strings.Split(s, " ") // 返回一个切片
	m := make(map[string]int, 20)
	for _, str := range sli {
		_, ok := m[str]
		// value, ok := m[str]
		if ok {
			m[str] = m[str] + 1 // 第一个元素本来就是value
		} else {
			m[str] = 1
		}
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("_________________________")
	m2 := make(map[int]int)
	fmt.Println(m2[1]) // 0?????
	fmt.Println(m2)
}

// 简单来说, 两个切片就有两个len, 即使是同一个底层数组!!!懂了吗?
func map6() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)       // 123
	fmt.Printf("%+v\n", s) // 123
	m["zzz"] = s           // "q1mi" : [1,2,3]
	fmt.Printf("%+v\n", m["zzz"])
	s = append(s[:1], s[2:]...) // [1,3]
	fmt.Printf("%+v\n", s)      // [1,3]
	fmt.Printf("%+v\n", m["zzz"])
	fmt.Printf("%p, %p", s, m["zzz"])
}
