package main

import (
	"fmt"
	"math"
	"mosn.io/holmes"
	"time"
)

type Agreement struct {
	Title string `json:"title,omitempty"` // 协议标题
	Url   string `json:"url,omitempty"`   // 协议跳链
}

func main() {
	//fmt.Println(AvToBv(602109295))
	//fmt.Println(BvToAv("BV16G411K7Hq"))
	//前闭后开()
	m := []int{1, 2, 3, 4, 5}
	m = rotate(m)
	fmt.Printf("%+v\n", m)
	//fmt.Printf("%p\n", m)
	//a := m
	//a[0] = 100
	//fmt.Printf("%+v", a)
}

func calculateDiffDays(a, b time.Time) int64 {
	return int64(math.Round(a.Sub(b).Abs().Hours()))
}

func rotate(m []int) []int {
	fmt.Printf("%p\n", m)
	return []int{1, 2, 3}
}

func initHolmes() *holmes.Holmes {
	h, _ := holmes.New(
		holmes.WithCollectInterval("5s"),
		holmes.WithDumpPath("./tmp"),
		holmes.WithCPUDump(20, 25, 80, time.Minute),
		holmes.WithCPUMax(90),
	)
	h.EnableCPUDump()
	return h
}

func maxProduct(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}
	m := nums[0]
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if i < n-1 {
				dp[i][j] = dp[i+1][j] * nums[i]
			} else {
				dp[i][j] = dp[i][j-1] * nums[j]
			}
			m = max(m, dp[i][j])
		}
	}
	return m
}
