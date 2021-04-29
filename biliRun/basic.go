package bili

import (
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
)

func whatCanBePrinted() {
	fmt.Println("由於對方隱私設置\n無法查看空間內容")
	fmt.Println("无法查看空间内容\n请将该用户移除黑名单")
}

func timed() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
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
	fmt.Println(unsafe.Sizeof(1))
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