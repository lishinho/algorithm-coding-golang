package main

// 单调递减栈
// 此例中栈中元素单调递减
type monoStack struct {
	stack []int
}

// 创建单调栈
func (m *monoStack) NewMonoStack(in []int) *monoStack {
	return &monoStack{stack: in}
}

// Peek：返回栈顶元素
func (m *monoStack) Peek() int {
	if len(m.stack) > 0 {
		return m.stack[len(m.stack)-1]
	}
	return -1
}

// Push: 压栈，如果栈顶元素<=进入元素，则pop栈顶元素，并更新栈顶元素对应value
// 如果栈顶元素>进入元素，则直接push
func (m *monoStack) Push(a int, f func(peek int)) {
	for len(m.stack) > 0 && m.Peek() <= a {
		f(m.Peek())
		m.Pop()
	}
	m.stack = append(m.stack, a)
}

// Pop: 把栈顶元素pop出去
func (m *monoStack) Pop() {
	if len(m.stack) > 0 {
		m.stack = m.stack[:len(m.stack)-1]
		return
	}
}
