package main

import "fmt"

func main() {

	ages := make(map[string]int)
	ages["tom"] = 18
	ages["lily"] = 0

	/*ages := map[string]int{
		"tony": 30,
		"alice": 20,
	}*/
	// len(map) 返回map中元素的个数
	fmt.Println("len(ages)", len(ages))
	// 通过 key 访问 map 中的元素 value，如果 key 不存在返回 value 的默认值
	fmt.Println("tom age is", ages["tom1"])

	// 判断一个key是否存在，ok 为 true 说明存在， 为 false 表示 key 不存在
	age, ok := ages["lily"]
	fmt.Println("lily age", age)
	fmt.Println("lily ok", ok)

	// 移除一个元素
	//delete(ages, "tom")

	_, ok = ages["tom"]
	fmt.Println("tom ok", ok)

	fmt.Println("---------------------------")
	for k, v := range ages {
		fmt.Printf("%s %d\n", k, v)
	}

}
