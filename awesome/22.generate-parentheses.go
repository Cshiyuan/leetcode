package awesome

func generateParenthesis(n int) []string {

	if n == 0 {
		return []string{}
	}
	totalL := make([][]string, n+1, n+1)
	//for i := range totalL {
		//totalL[i] = make([]string, n)
	//}
	totalL[0] = []string{""}
	totalL[1] = []string{"()"}
	// 开始遍历计算i组括号的组合情况
	for i := 2; i <= n; i++ {

		// 当前遍历结果
		l := []string{}
		for j := 0; j < i; j++ {
			now_list1 := totalL[j]     // p = j 时的括号组合情况
			now_list2 := totalL[i-1-j] // q = (i-1) - j 时的括号组合情况

			for _, v1 := range now_list1 {
				for _, v2 := range now_list2 {

					l = append(l, "("+v1+")"+v2)
				}
			}
		}
		totalL[i] = l
	}
	return totalL[n]
}

func generateParenthesis2(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	list := generateParenthesis(n - 1)
	sets := make(map[string]bool)
	var resultList []string
	for _, s := range list {

		for _, s := range []string{s + "()", "()" + s, "(" + s + ")"} {
			if _, ok := sets[s]; !ok {
				resultList = append(resultList, s)
				sets[s] = true
			}
		}
	}
	return resultList
}
