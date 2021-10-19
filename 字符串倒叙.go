package main

import "fmt"

//1、利用顺序表倒叙

type sqList struct {
	flashBS []byte
	length int       //计数器
}

func main() {
	fmt.Println("请输入字符串")
	var s string
	fmt.Scan(&s)
	b := []byte(s)                         //利用byte型切片接受字符串s
	fmt.Println("倒置后")

	var L sqList                           // 创建顺序表L
	listInit(&L)
	flashBackStr(b,&L)
	fmt.Printf("%s",L.flashBS[:])
}

func listInit(L *sqList)*sqList{           //对顺序表初始化
	L.flashBS = make([]byte,100)
	L.length = 0
	return L
}

func flashBackStr(b []byte,L *sqList)*sqList{   //倒叙字符串存到顺序表
	for i :=0;i<len(b);i++{
		L.flashBS[L.length] = b[len(b)-1-i]     //将切片b从最后一位元素开始存入L.flashBs
		L.length++
	}
	return L
}


//2、利用for循环   第一个元素与最后一个交换，第二个与倒数第二个交换...
/*
func flashBackStr(b []byte){
	var temp byte            //创建临时变量
	for i:= 0;i<len(b)/2;i++{
		temp = b[i]
		b[i] = b[len(b)-1-i]
		b[len(b)-1-i] = temp
	}

}

func main() {
	fmt.Println("请输入字符串")
	var s string
	fmt.Scan(&s)

	b := []byte(s)          //利用byte型切片接受字符串s
	fmt.Println("倒置前")
	fmt.Printf("%s\n",b)

	fmt.Println("倒置后")
	flashBackStr(b[:])
	fmt.Printf("%s",b)

}
*/
