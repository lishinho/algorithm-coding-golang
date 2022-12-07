package leetcode

// https://leetcode.cn/problems/word-ladder-ii
//func findLadders(beginWord string, endWord string, wordList []string) (res [][]string) {
//	n := len(wordList)
//	m := make(map[string][]int, n+1)
//	visited := make([][]bool, n)
//	for i := range visited {
//		visited[i] = make([]bool, n)
//	}
//
//	// O(n)*len(word)
//	for i := range wordList {
//		if canBeConverted(beginWord, wordList[i]) {
//			if len(m[beginWord]) == 0 {
//				m[beginWord] = []int{i}
//			} else {
//				m[beginWord] = append(m[beginWord], i)
//			}
//		}
//	}
//	// O(n^2)*len(word)
//	for i := 0; i < n; i++ {
//		m[wordList[i]] = []int{}
//		for j := n - 1; j >= 0; j-- {
//			if i == j {
//				continue
//			}
//			if visited[i][j] || visited[j][i] || canBeConverted(wordList[i], wordList[j]) {
//				m[wordList[i]] = append(m[wordList[i]], j)
//				visited[i][j] = true
//				visited[j][i] = true
//			}
//		}
//	}
//
//	// O(path)
//	var tmp []string
//	tmp = append(tmp, beginWord)
//	wordVisited := make(map[string]bool, n)
//	minn := len(wordList)
//	var dfs func(s1 string, tmp []string)
//	dfs = func(s1 string, tmp []string) {
//		if s1 == endWord {
//			if len(tmp) < minn {
//				minn = len(tmp)
//			}
//			res = append(res, append([]string{}, tmp...))
//			return
//		}
//		if len(tmp) >= minn || wordVisited[s1] {
//			return
//		}
//		for _, v := range m[s1] {
//			wordVisited[s1] = true
//			tmp = append(tmp, wordList[v])
//			dfs(wordList[v], tmp)
//			wordVisited[s1] = false
//			tmp = tmp[:len(tmp)-1]
//		}
//	}
//	dfs(beginWord, tmp)
//	var realRes [][]string
//	for _, v := range res {
//		if len(v) == minn {
//			realRes = append(realRes, v)
//		}
//	}
//	return realRes
//}

func canBeConverted(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var cnt int
	for i := range s1 {
		if s1[i] != s2[i] {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return cnt == 1
}

// bfs+dfs(如果是双向bfs，效果会更好)
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	//字典表（将wordList中的单词放入hash表中，方便查找）
	dict := make(map[string]bool, 0)
	for _, v := range wordList {
		dict[v] = true
	}
	//如果endWord不在hash表中，表示不存在转换列表，返回空集合
	if !dict[endWord] {
		return [][]string{}
	}
	//将第一个单词放入hash表中，方便实现邻接表（构建图）
	dict[beginWord] = true
	//构建邻接表
	graph := make(map[string][]string, 0)
	//执行bfs搜索，找到每个点到endWord的距离
	distance := bfs(endWord, dict, graph)
	res := make([][]string, 0) //保存结果
	//执行dfs操作
	dfs(beginWord, endWord, &res, []string{}, distance, graph)
	return res
}

// 回溯实现方式一：（个人偏好这个，更符合模板）
func dfs(beginWord string, endWord string, res *[][]string, path []string, distance map[string]int, graph map[string][]string) {
	//出递归条件
	if beginWord == endWord {
		path = append(path, beginWord) //加入endWord节点
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append((*res), tmp)
		path = path[:len(path)-1] //移除endWord节点
		return
	}
	//否则遍历图
	for _, v := range graph[beginWord] {
		//遍历图时，朝着距离与终点越来越近的方向进行（当前节点的距离肯定要比下一个距离大1）
		if distance[beginWord] == distance[v]+1 {
			path = append(path, beginWord)
			dfs(v, endWord, res, path, distance, graph)
			//回溯（执行完上述的所有时，将其回溯回去）
			path = path[:len(path)-1]
		}
	}
}

//回溯实现方式二：
// func dfs(beginWord string,endWord string,res *[][]string,path []string,distance map[string]int,graph map[string][]string){
//     path=append(path,beginWord)
//     //出递归条件
//     if beginWord==endWord{
//         tmp:=make([]string,len(path))
//         copy(tmp,path)
//         (*res)=append((*res),tmp)
//         return
//     }
//     //否则遍历图
//     for _,v:=range graph[beginWord]{
//         //遍历图时，朝着距离与终点越来越近的方向进行（当前节点的距离肯定要比下一个距离大1）
//         if distance[beginWord]==distance[v]+1{
//             dfs(v,endWord,res,path,distance,graph)
//         }
//     }
//     //回溯（执行完上述的所有时，将其回溯回去）
//     path=path[:len(path)-1]
// }

// 从终点出发，进行bfs，计算每一个点到达终点的距离
func bfs(endWord string, dict map[string]bool, graph map[string][]string) map[string]int {
	distance := make(map[string]int, 0) //每个点到终点的距离
	queue := make([]string, 0)
	queue = append(queue, endWord)
	distance[endWord] = 0 //初始值
	for len(queue) != 0 {
		cursize := len(queue)
		for i := 0; i < cursize; i++ {
			word := queue[0]
			queue = queue[1:]
			//找到和word有一位不同的单词列表
			expansion := expand(word, dict)
			for _, v := range expansion {
				//构造邻接表
				//我们是从beginWord到endWord构造邻接表，而bfs是从endWord开始，所以构造时，反过来构造
				//即graph[v]=append(graph[v],word)而不是graph[word]=append(graph[word],v)
				graph[v] = append(graph[v], word)
				//表示没访问过
				if _, ok := distance[v]; !ok {
					distance[v] = distance[word] + 1 //距离加一
					queue = append(queue, v)         //入队列
				}
			}
		}
	}
	return distance
}

// 获得邻接点
func expand(word string, dict map[string]bool) []string {
	expansion := make([]string, 0) //保存word的邻接点
	//从word的每一位开始
	chs := []rune(word)
	for i := 0; i < len(word); i++ {
		tmp := chs[i] //保存当前位，方便后序进行复位
		for c := 'a'; c <= 'z'; c++ {
			//如果一样则直接跳过，之所以用tmp，是因为chs[i]在变
			if tmp == c {
				continue
			}
			chs[i] = c
			newstr := string(chs)
			//新单词在dict中不存在，则跳过
			if dict[newstr] {
				expansion = append(expansion, newstr)
			}
		}
		chs[i] = tmp //单词位复位
	}
	return expansion
}
