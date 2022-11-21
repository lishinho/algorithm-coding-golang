package main

import (
	"encoding/json"
	"fmt"
	"go-common/library/ecode"
	"math"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const (
	// Containers are for normal containers
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
	b := 1
	fmt.Println(-b)
}

func whatCanBePrinted() {
	fmt.Println("由於對方隱私設置\n無法查看空間內容")
	fmt.Println("无法查看空间内容\n请将该用户移除黑名单")
	fmt.Println("无法查看空间内容\n请将该用户移除黑名单")
}

func timed() {
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Unix())
	loc, _ := time.LoadLocation("Brazil/DeNoronha")
	fmt.Println(time.Now().In(loc).Format(time.RFC3339))
	fmt.Println(time.Unix(-866131200, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(strconv.FormatInt(time.Now().UnixNano()/1000000, 10))
}

func goSlice() {
	m := map[int]int{}
	for _, i := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("sdss=%d\n", i)
		m[i] = i + 1
		//i := i
		//go func() {
		//	fmt.Println(i)
		//}()
	}
	fmt.Printf("map %+v\n", m)
	time.Sleep(time.Second * 1)
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

func printLenCap(nums []int) {
	fmt.Printf("len: %d, cap: %d %v\n", len(nums), cap(nums), nums)
}

func basicMapNilTest() {
	m := map[string]int{}
	m = nil
	fmt.Printf("sss %+v", m["sss"])
	fmt.Printf("111")
}

func RoundHalfUp(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

func testSlice() {
	nums := make([]int, 0, 8)
	nums = append(nums, 1, 2, 3, 4, 5)
	nums2 := nums[2:4]
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 5]
	printLenCap(nums2) // len: 2, cap: 6 [3 4]

	nums2 = append(nums2, 50, 60)
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 50]
	printLenCap(nums2) // len: 4, cap: 6 [3 4 50 60]

	nums2 = append(nums2, 1, 2, 3, 4, 5)
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 50]
	printLenCap(nums2) // len: 9, cap: 12 [3 4 50 60, 1, 2, 3, 4, 5]
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
	arr = append(arr, nil, nil)
	fmt.Printf("%+v", arr)
}

func maxInteger() {
	fmt.Println(IntMax)
	fmt.Println(IntMin)
}

func checkEcodeNil() {
	Myecode := (*ecode.Code)(nil)
	var err error
	err = Myecode
	if err == nil {
		fmt.Println("sss")
	}
	if Myecode == nil {
		fmt.Println("s")
	}
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

func stringCount() {
	s := "[星][星][星][星]\n好看[星]"
	ss := "[星]"

	fmt.Println(strings.Count(s[:strings.Index(s, "\n")], ss))
	fmt.Println(ss[strings.Index(ss, "\n")+1:])
}

func searchTarget() {
	arr := []int{1, 2, 3, 44, 5, 66, 7, 88, 9}
	target := 100
	fmt.Println(sort.SearchInts(arr, target))
}

func mapPaixu() {
	dynamicIdMap := map[int64]int64{1: 1, 2: 2, 3: 3}
	for k, v := range dynamicIdMap {
		fmt.Println(k, v)
	}
}
