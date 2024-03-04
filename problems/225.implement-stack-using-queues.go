package problems

/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
	storage []int
}

func Constructor() MyStack {
	return MyStack{
		storage: make([]int, 0),
	}
}

func (stack *MyStack) Push(x int) {
	stack.storage = append([]int{x}, stack.storage...)
}

func (stack *MyStack) Pop() int {
	res := stack.storage[0]
	stack.storage = stack.storage[1:]
	return res
}

func (stack *MyStack) Top() int {
	return stack.storage[0]
}

func (stack *MyStack) Empty() bool {
	return len(stack.storage) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
