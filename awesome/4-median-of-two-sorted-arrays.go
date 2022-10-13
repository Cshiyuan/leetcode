package awesome

//给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的 中位数 。
//
//算法的时间复杂度应该为 O(log (m+n)) 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/median-of-two-sorted-arrays
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)

	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return (float64(getKthElement(nums1, nums2, midIndex1+1)) +
			float64(getKthElement(nums1, nums2, midIndex2+1)))/2.0
	}
}

func getKthElement(nums1 []int, nums2 []int, k int) int {

	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		// 中位数判断
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 < pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
