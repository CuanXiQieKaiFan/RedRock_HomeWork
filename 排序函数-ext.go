package main

import(
	"fmt"
	"math/rand"
	"time"
	"sort"
)

var (
	slice []int
	seed int64  //随机数种子
	ran int     //随机数
)

func main() {
	seed = time.Now().Unix()
	rand.Seed(seed)
	for i := 0; i < 100; i++ {
		ran = rand.Intn(101)    //0-100
			slice= append(slice,ran)
	}
	for _,i := range slice{
		fmt.Print(i," ")
	}
	fmt.Println("排序后：")
	sort.Slice(slice, func(i, j int) bool {
		return slice[i]>slice[j]
	})
	for _,i := range slice{
		fmt.Print(i," ")
	}
}
