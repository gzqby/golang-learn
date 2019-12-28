package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	count := 0
	fibMap := make(map[int]int)
	return func() int {
		count++
		if count==1 || count==2 {
			fibMap[count]=1
		}else{
			fibMap[count]=fibMap[count-1]+fibMap[count-2]
		}
		return fibMap[count]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 24; i++ {
		fmt.Println(f())
	}
}
