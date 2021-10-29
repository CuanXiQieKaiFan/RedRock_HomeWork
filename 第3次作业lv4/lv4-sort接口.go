package main

import (
	"fmt"
	"sort"
)


type students struct {
	Name string
	Age int
	Score float64
}
type studentsList  []students  // []students别名
func (slist studentsList) CopyList(list studentsList)	*studentsList  {
	for _,stu := range list{
		slist = append(slist,stu)
	}
	return &slist
}

func (slist studentsList)Len()int  {
	return  len(slist)
}
func (slist studentsList)Swap(i,j int){
	slist[i],slist[j] = slist[j],slist[i]
}
func (slist studentsList)Less(i,j int)bool  {  //降序
	return slist[i].Score > slist[j].Score
}

func main() {
	stu1 := students{"Tom",18,59}
	stu2 := students{"Jan",18,66}
	stu3 := students{"Jenifer",18,33}
	stu4 := students{"Tracy",18,90}
	stu5 := students{"Reachle",18,48}
	//list := studensList{stu1,stu2,stu3,stu4,stu5}
	list := make([]students, 4)
	list = []students{stu1, stu2, stu3, stu4,stu5}
	var  Slice studentsList
	Mylist := Slice.CopyList(list)
	fmt.Println("排序前")
	fmt.Println(Mylist)
	sort.Sort(Mylist)
	fmt.Println("排序后")
	fmt.Println(Mylist)

}
