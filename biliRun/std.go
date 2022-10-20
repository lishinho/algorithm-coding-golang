package main

import (
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type errorMsg struct {
	err string
	col int
}

func std() {
	var in []*errorMsg
	for {
		var (
			a string
			b int
		)
		_, err := fmt.Scan(&a, &b)
		if err == io.EOF {
			break
		}
		in = append(in, &errorMsg{
			err: a,
			col: b,
		})
	}
	m, appear := makeResMap(in)
	for _, key := range appear {
		if _, ok := m[key]; ok {
			fmt.Printf("%s %d\n", key, m[key])
		}
	}
}

func makeResMap(in []*errorMsg) (map[string]int, []string) {
	res := make(map[string]int, len(in))
	appear := make([]string, 8)
	for _, v := range in {
		t := makeSubString(v.err)
		if _, ok := res[fmt.Sprintf("%s %d", t, v.col)]; ok {
			res[fmt.Sprintf("%s %d", t, v.col)]++
			continue
		}
		res[fmt.Sprintf("%s %d", t, v.col)] = 1
		if len(appear) == 8 {
			appear = appear[1:]
		}
		appear = append(appear, fmt.Sprintf("%s %d", t, v.col))
	}
	return res, appear
}

func isPrimeCouple(a, b int) bool {
	return gcd(a, b) == b
}

func resolve(nums []int) int {
	n := len(nums)
	var res int
	used := make([]bool, n)
	bucket := make([][]int, n/2)
	for i := range bucket {
		bucket[i] = make([]int, 2)
	}
	sort.Ints(nums)
	var backtrack func(start, k, now int)
	backtrack = func(start, k, now int) {
		if k == n/2 {
			return
		}
		if len(bucket[k]) == 2 {
			backtrack(0, k+1, now)
			return
		}
		for i := start; i < n; i++ {
			if used[i] {
				continue
			}
			if len(bucket[k]) == 0 {
				bucket[k][0] = nums[i]
				used[i] = true
				backtrack(i+1, k, now)
				used[i] = false
				continue
			}
			if len(bucket[k]) == 1 {
				if !isPrimeCouple(nums[i], bucket[k][0]) {
					continue
				} else {
					bucket[k][1] = nums[i]
					used[i] = true
					now++
					res = max(res, now)
					backtrack(0, k+1, now)
					used[i] = false
					now--
				}
			}
		}
	}
	backtrack(0, 0, 0)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func resolveCacheTimeout() time.Duration {
	rand.Seed(time.Now().UnixNano())
	return 24*time.Hour + time.Duration(rand.Int63n(100))*time.Second
}

func makeSubString(str string) string {
	sub := str[strings.LastIndexByte(str, '\\')+1:]
	if len(sub) > 16 {
		return sub[len(sub)-16:]
	}
	return sub
}

func reSort(raw []int) []int {
	cnt := [1000]bool{}
	for _, v := range raw {
		cnt[v] = true
	}
	var res []int
	for i, v := range cnt {
		if v {
			res = append(res, i)
		}
	}
	return res
}
