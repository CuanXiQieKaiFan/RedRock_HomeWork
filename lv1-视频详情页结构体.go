package main

import (
	"fmt"
	"reflect"
)

type VideoInterface struct {      //视频详情页面
	Comment string         //评论
	Title   string         //标题
	Author                 //作者
}

type Author struct {
	Name string             //名字
	VIP bool                //是否是高贵的带会员
	Icon string             //头像
	Focus int               //关注人数
	Thumbs_up int           //点赞数
	Collector int           //收藏
	Transmiter  int           //转发
	VlogKind []string       //相关视频
}


func main()  {
	Kind :=CreateSlice()  //创建相关视频种类
	BingB :=Author{"吃花椒的喵酱",true,"Q版冰冰",4928000,388000,67000,33000,Kind} //创建作者
	Bing :=VideoInterface{"阿巴阿巴","失踪人口回归",BingB}               //创建视频页面
	printVideo(Bing)  //打印视频页面

}

//创建相关推荐视频的种类
func CreateSlice()(Kind []string) {
	k1 := "生活"
	k2 := "咨讯"
	k3 := "运动"
	Kind = append(Kind,k1)
	Kind = append(Kind,k2)
	Kind = append(Kind,k3)
	return Kind
}
func printAuthor(BingB Author){
	v := reflect.TypeOf(BingB)
	f :=reflect.ValueOf(BingB)
	for i:=0;i<v.NumField();i++{
		t := v.Field(i)
		val := f.Field(i)
		fmt.Printf("%s: %v\n",t.Name,val)
	}
}
func printVideo(Bing VideoInterface){
	fmt.Println("Title:",Bing.Title)
	printAuthor(Bing.Author)
	fmt.Println("Comment:",Bing.Comment)
}
