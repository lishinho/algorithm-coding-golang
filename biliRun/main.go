package main

import (
	"crypto/md5"
	"encoding/hex"
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
	fmt.Println(BvToAv("BV1z24y137CU"))
	//前闭后开()
	//fmt.Println(calculateDiffDays2(time.Now(), time.Now().Add(33*time.Hour)))
	//fmt.Println(calculateDiffDays(time.Now(), time.Now().Add(33*time.Hour)))
	//fmt.Printf("%p\n", m)
	//a := m
	//a[0] = 100
	//fmt.Printf("%+v", a)
	m := make(map[int64]bool, 4)
	m[3] = true
	fmt.Printf("%+v", m)
	fmt.Printf("%+v", Md5("[UPOWER_2054502104988661_新增4]", []byte{}))
}

//get upower:emo:text:356136e203e15ce522fcdf0ac3200b34

func Md5(plaintext string, salt []byte) string {
	if plaintext == "" {
		return ""
	}
	hash := md5.New()
	hash.Write([]byte(plaintext))
	hash.Write([]byte(salt))
	mh := hash.Sum(nil)
	return hex.EncodeToString(mh[:])
}

func calculateDiffDays2(a, b time.Time) int64 {
	return int64(math.Abs(truncateToDay(a).Sub(truncateToDay(b)).Hours()) / 24)
}

func truncateToDay(a time.Time) time.Time {
	return time.Date(a.Year(), a.Month(), a.Day(), 0, 0, 0, 0, time.Local)
}

func calculateDiffDays(a, b time.Time) int64 {

	return int64(math.Abs(float64(a.Truncate(24 * time.Hour).Sub(b.Truncate(24 * time.Hour)))))
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
