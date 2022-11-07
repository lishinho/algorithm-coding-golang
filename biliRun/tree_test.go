package main

import "testing"

func BenchmarkLevelOrder(b *testing.B) {
	testTree := makeTree(3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8)
	for i := 0; i < b.N; i++ {
		levelOrder(testTree)
	}
}

func BenchmarkLevelOrder1(b *testing.B) {
	testTree := makeTree(3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8)
	for i := 0; i < b.N; i++ {
		levelOrder1(testTree)
	}
}
