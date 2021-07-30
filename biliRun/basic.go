package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"time"
	"unsafe"
)

const (
	// Containers is for normal containers
	Containers int = 1 << iota
	// InitContainers is for init containers
	InitContainers
	// EphemeralContainers is for ephemeral containers
	EphemeralContainers

	IntMax  = int(^uint(0) >> 1)
	IntMin  = ^IntMax
	IntMock = ^uint(0)
)

func 前闭后开() {
	a := "string"
	fmt.Println(a[:0])
	fmt.Println(a[:len(a)-1])
}

func whatCanBePrinted() {
	fmt.Println("由於對方隱私設置\n無法查看空間內容")
	fmt.Println("无法查看空间内容\n请将该用户移除黑名单")
	fmt.Println("无法查看空间内容\n请将该用户移除黑名单")
}

func timed() {
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Unix(1627300800, 0).Format("2006-01-02 15:04:05"))
}

func goSlice() {
	for _, i := range []int{1, 2, 3, 4, 5} {
		//i := i
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 5)
}

func a(i int) {
	fmt.Println(i)
}

func outOfArrayGet() {
	const name = "dddd"
	arr := [3]string{"sss", "ddd", "ddd"}
	for i := range arr[:2] {
		fmt.Println(arr[i])
	}
	fmt.Println(name, &arr)
}

func parseInt() {
	id, _ := strconv.ParseInt("16982385563916616011", 10, 64)
	fmt.Println(id)
}

func calculateEvenMedian(a, b int) float64 {
	return float64(a+b) / 2
}

func howManyBitsCanStand() {
	fmt.Println(runtime.GOARCH) //CPU型号
	fmt.Println(unsafe.Sizeof(struct{}{}))
	fmt.Println(unsafe.Sizeof(int64(1)))
	fmt.Println(unsafe.Sizeof(int32(1)))
	fmt.Println(unsafe.Sizeof(int8(1)))
	fmt.Println(unsafe.Sizeof(3.33))
	fmt.Println(unsafe.Sizeof("hello world wwww"))
}

func basicIotaValues() {
	fmt.Println(Containers)
	fmt.Println(InitContainers)
	fmt.Println(EphemeralContainers)
}

func deferAdd1() (i int) {
	i = 10
	defer func() { i++ }()
	return 1
}

func deferAdd2() int {
	i := 10
	defer func() { i++ }()
	return 1
}

func minus(i int) {
	fmt.Println(-i)
}

func initArrayCanAppend() {
	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	fmt.Printf("%+v", arr)
}

func appendNil() {
	var arr []*Hot
	arr = append(arr, nil)
	fmt.Printf("%+v", arr)
}

func arrayToJson() {
	arr := []int64{1, 2, 3, 4, 5}
	m := map[string][]int64{
		"str": arr,
		"abc": arr,
	}
	bt, _ := json.Marshal(m)
	fmt.Printf("%s", bt[1:len(bt)-1])
}

func maxInteger() {
	fmt.Println(IntMax)
	fmt.Println(IntMin)
}

func jsonMarshalT() {
	var res int64
	if err := json.Unmarshal(nil, &res); err != nil {
		fmt.Printf("error %+v", err)
		return
	}
	fmt.Println(res)
}

func testMapNil() {
	var mm map[int64]*ReserveActExtra
	if v, ok := mm[3]; ok {
		fmt.Println(v.IsValid)
	}
	fmt.Println(IntMax)
	fmt.Println(IntMin)
	fmt.Println(IntMock)
}
