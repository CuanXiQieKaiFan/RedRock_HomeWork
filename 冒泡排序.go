package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	slice []int
	seed int64  //随机数种子
	ran int     //随机数
)


func main() {
	seed = time.Now().Unix()
	rand.Seed(seed)
	//排序前
	for i := 0; i < 100; i++ {
		ran = rand.Intn(101)    //0-100
		slice= append(slice,ran)
	}
	for _,i := range slice{
		fmt.Print(i," ")
	}
	//排序后
	fmt.Println("排序后")
	mySort(&slice)
	for _,i := range slice{
		fmt.Print(i," ")
	}
}

func mySort(slice *[]int) {
	for i := 0; i < len(*slice);i++ {
		for j:=0;j<len(*slice)-1-i;j++ {
			if(*slice)[j]<(*slice)[j+1]{     //降序
				temp :=(*slice)[j]
				(*slice)[j] = (*slice)[j+1]
				(*slice)[j+1]=temp
			}
		}
	}

}
