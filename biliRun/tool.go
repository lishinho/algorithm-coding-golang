package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"time"
)

var str = "b3Blbl9hcHBfZnJvbV90eXBlPWRlZXBsaW5rX3R0bGgtYW5kLXJkLTMxeWF6LTE2Nzc5NTY3Mzc5MjgxOTkmb3Blbl9hcHBfdXJsPXNzMzkzMTE="

type TmpSlice struct {
	Num  int64
	Name string
}

func sortArray() {
	//tmp := []int{0, 0, 0, 4, 3, 2, 5, 43}
	tmpSlice := []*TmpSlice{
		{
			Num:  0,
			Name: "aaa",
		},
		{
			Num:  0,
			Name: "bbb",
		},
		{
			Num:  0,
			Name: "ccc",
		},
		{
			Num:  26,
			Name: "ddd",
		},
		{
			Num:  13,
			Name: "eee",
		},
		{
			Num:  21,
			Name: "fff",
		},
		{
			Num:  7,
			Name: "ggg",
		},
		{
			Num:  21,
			Name: "hhh",
		},
	}
	//从大到小
	//sort.SliceStable(tmp, func(i, j int) bool {
	//	if tmp[i] == 0 || tmp[j] == 0 {
	//		return false
	//	}
	//	return tmp[i] > tmp[j]
	//})

	sort.SliceStable(tmpSlice, func(i, j int) bool {
		if tmpSlice[i].Num == 0 || tmpSlice[j].Num == 0 {
			return false
		}
		return tmpSlice[i].Num > tmpSlice[j].Num
	})

	// 从小到大
	//sort.Ints(tmp)
	for _, v := range tmpSlice {
		fmt.Printf("num=%+v, name=%+v\n", v.Num, v.Name)
	}

}

func baseDecode() {
	bs, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("Error %+v", err)
		return
	}
	params, err := url.ParseQuery(string(bs))
	if err != nil {
		fmt.Printf("Error2 %+v", err)
		return
	}
	fmt.Printf("params %+v\n", params.Get("open_app_from_type"))
	fmt.Printf("params %+v\n", params.Get("open_app_url"))

}

func timeStampToTime() {
	timestamp := time.Now().Unix()
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	fmt.Printf("%d年%02d月%02d日 %02d:%02d开始", year, month, day, hour, minute)
}

func switchTestcase() {
	id := 1
	switch {
	case true:
		fmt.Println("顺序来")
	case id < 2:
		fmt.Println("不一定顺序来")
	default:
		fmt.Println("222")
	}
}
