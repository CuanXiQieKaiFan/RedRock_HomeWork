package main

import "fmt"

var (
	a,b float64
	str string

	slice []float64
)
func main()  {
	fmt.Println("请输入表达式：")
	for{
		fmt.Scanln(&a,&str,&b)
		slice = append(slice,calculate(a,str,b))
		fmt.Println(slice)
	}


}

func calculate(a float64,str string,b float64)(sum float64){
	if str == "+" {
		sum = a+b
	} else if str == "-" {
		sum = a-b
	} else if str == "*" {
		sum = a*b
	} else if str == "/"{
		if b == 0{
			fmt.Println("除数不能为0！")

		}else {
			sum = a/b
		}

	}else {
		fmt.Println("输入有误，宝！")
	}
	return sum
}