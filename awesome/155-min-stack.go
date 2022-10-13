package awesome

import "math"

//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//实现 MinStack 类:
//
//MinStack() 初始化堆栈对象。
//void push(int val) 将元素val推入堆栈。
//void pop() 删除堆栈顶部的元素。
//int top() 获取堆栈顶部的元素。
//int getMin() 获取堆栈中的最小元素。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/min-stack
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type MinStack struct {
	slice  []int //具体的数据结构
	minPos int   //记录最小的位置
}

func Constructor1() MinStack {
	return MinStack{
		slice:  []int{},
		minPos: -1,
	}
}

func (s *MinStack) Push(val int) {

	// 压栈
	s.slice = append(s.slice, val)
	if s.minPos == -1 {
		s.minPos = 0
		return
	}
	// 说明最小的位置
	if s.slice[s.minPos] > val {
		s.minPos = len(s.slice) - 1
	}
}

// 出栈
func (s *MinStack) Pop() {
	oPos := len(s.slice)-1
	s.slice = s.slice[:len(s.slice)-1]
	if s.minPos ==  oPos{
		minVal := math.MaxInt64
		for pos, val := range s.slice {
			if val < minVal {
				minVal = val
				s.minPos = pos
			}
		}
	}
}

func (s *MinStack) Top() int {
	return s.slice[len(s.slice)-1]
}

func (s *MinStack) GetMin() int {
	return s.slice[s.minPos]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
