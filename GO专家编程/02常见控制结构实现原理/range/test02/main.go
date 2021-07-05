package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := []int{1, 2, 3}
	for i:= range v {
		fmt.Println(v[i])
		v = append(v, i)
	}
	fmt.Println(v) // [1 2 3 0 1 2]

	m := map[int]string{1:"hello", 2:"world"}
	var i int = 3
	for key, value := range m{
		fmt.Printf("key:%d, value:%s\n", key, value)
		m[i] = value + "_" + strconv.Itoa(i)
		i++
	}
	fmt.Println(m)

}

/*
1
2
3
[1 2 3 0 1 2]
key:2, value:world
key:3, value:world_3
key:4, value:world_3_4
key:5, value:world_3_4_5
key:6, value:world_3_4_5_6
key:7, value:world_3_4_5_6_7
key:8, value:world_3_4_5_6_7_8
key:1, value:hello
map[1:hello 2:world 3:world_3 4:world_3_4 5:world_3_4_5 6:world_3_4_5_6 7:world_3_4_5_6_7 8:world_3_4_5_6_7_8 9:world_3_4_5_6_7_8_9 10:hello_10]
 */

// 能够正常结束
// 遍历 slice 时循环内改变切片的长度，不影响循环次数，循环次数在循环开始前就已经确定了
// 遍历 map 时循环内改变追加元素时新元素会被循环出来
// map底层使用hash表实现，插入数据位置是随机的，所以遍历过程中新插入的数据不能保证遍历到


// 这点与 python 不一样 python 遍历列表时追加元素会一直无休止遍历
// python 遍历字典时追加元素会报 RuntimeError: dictionary changed size during iteration
/*
def rangeList():
	v = [1, 2, 3]
	for i in v:
		print(i)
		v.append(i)

def rangeDict():
	i = 3
	m = {1:"hello",2:"world"}
	for k,v in m.items():
		print(v)
		m[i] = v + str(i)
		i = i + 1
 */