package awesome

//给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指在第 i 天之后，才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/daily-temperatures
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type minStack struct {
	val []int
}


func (s *minStack) top() int {
	return s.val[len(s.val) - 1]
}

func (s *minStack) push(i int) {
	s.val = append(s.val, i)
}

func (s *minStack) isEmpty() bool {
	return len(s.val) == 0
}

func (s *minStack) pop() {
	if len(s.val) == 1 {
		s.val = []int{}
		return
	}
	s.val = s.val[:len(s.val) - 1]
}

func (s *minStack) len() int {
	return len(s.val)
}

func dailyTemperatures(temperatures []int) []int {

	// 最小栈
	s := minStack{}
	ans := make([]int, len(temperatures), len(temperatures))
	for i, t := range temperatures {

		if  s.isEmpty() {
			s.push(i)
			continue
		}
		for {
			if s.isEmpty() {
				s.push(i)
				break
			}
			// 栈顶
			top := s.top()
			// 直接入栈，然后退出
			if temperatures[top] >= t {
				s.push(i)
				break
			}
			// [73,74,75,71,69,72,76,73]
			//fmt.Printf("t: %v top: %v s: %v\n", i, top, s)
			//fmt.Printf("index: %v \n",  s.len() - 1)
			ans[top] =  i - top
			s.pop()
		}
	}

	return ans

}