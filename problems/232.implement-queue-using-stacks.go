package problems

/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

type MyQueue struct {
	normal  Stack
	reverse Stack
}

func Constructor() MyQueue {
	return MyQueue{
		normal:  make([]int, 0),
		reverse: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	for !q.reverse.Empty() {
		q.normal.Push(q.reverse.Pop())
	}
	q.normal.Push(x)
}

func (q *MyQueue) Pop() int {
	for !q.normal.Empty() {
		q.reverse.Push(q.normal.Pop())
	}
	return q.reverse.Pop()
}

func (q *MyQueue) Peek() int {
	for !q.normal.Empty() {
		q.reverse.Push(q.normal.Pop())
	}
	return q.reverse.Top()
}

func (q *MyQueue) Empty() bool {
	return q.normal.Empty() && q.reverse.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
