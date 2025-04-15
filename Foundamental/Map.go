package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
map 是一种无序的 数据结构 其是引用类型
*/
func main() {
	userInfo := map[string]string{
		"username": "pprof.cn",
		"password": "123456",
	}
	// v 是 userInfo[key]的 value 而 ok 是个布尔变量
	v, ok := userInfo["username"]
	fmt.Println(v, ok)
	// 判断 键 username 是否存在
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key not exist")
	}
	fmt.Println(userInfo) //

	for k, v := range userInfo {
		fmt.Println(k, v)
	}

	for k := range userInfo {
		fmt.Println(k)
	}

	// 删除键值对
	delete(userInfo, "username")
	fmt.Println(userInfo)

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		// 随机生成value
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	// 从 map中取出所有key存入切片sliceKey

	var sliceKey = make([]string, 0, 200)
	fmt.Println(sliceKey)
	// 因为 map是无序的，所以遍历都是乱序
	for key := range scoreMap {
		sliceKey = append(sliceKey, key)
	}

	fmt.Println(sliceKey)
	// 对切片排序
	sort.Strings(sliceKey)
	fmt.Println(sliceKey)
	fmt.Println(scoreMap)
	for _, key := range sliceKey {
		fmt.Println(key, scoreMap[key])
	}
	// 存储map的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "123"
	mapSlice[0]["password"] = "456"
	mapSlice[0]["address"] = "789"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	// map 存储 切片
	sliceMap := make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")

	key := "中国"
	value, ok := sliceMap[key]
	// 如果不存在 siliceMap[key]
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
