package main

import "fmt"

func Receiver(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("这个是string")
	case int:
		fmt.Println("这个是int")
	case float64:
		fmt.Println("这个是float64")
	case bool:
		fmt.Println("这个是bool")
	default:
		fmt.Println("这个是自建数据类型或者啥也不是")
	}
}

func main() {

	Receiver("str")
	Receiver(32)
	Receiver(true)
}
