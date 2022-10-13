package awesome

func merge_sorted_array(nums1 []int, m int, nums2 []int, n int)  {

	p1 := m - 1
	p2 := n - 1
	tail := m+n - 1
	for ;p1 >=0 || p2 >= 0; tail -- {

		cur := 0
		if p1 == -1 {
			cur = nums2[p2]
			p2 --
		} else if p2 == -1 {
			cur = nums1[p1]
			p1 --
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1 --
		} else  {
			cur = nums2[p2]
			p2 --
		}
		nums1[tail] = cur
	}
}
