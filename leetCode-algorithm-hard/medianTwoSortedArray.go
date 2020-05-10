package main

/*
	Name: Median of Two Sorted Arrays
	Problem: https://leetcode.com/problems/median-of-two-sorted-arrays
 */
func main() {
	nums1 := []int{1,3}
	nums2 := []int{2}

	println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := combine(nums1, nums2)
	return medianOf(nums)
}

func combine(mis, njs []int) []int {
	lenMis, i := len(mis), 0
	lenNjs, j := len(njs), 0
	res := make([]int, lenMis + lenNjs)

	for k := 0; k < lenMis + lenNjs; k++ {
		if i == lenMis ||
			(i < lenMis && j < lenNjs && mis[i] > njs[j]) {
			res[k] = njs[j]
			j++
			continue
		}

		if j == lenNjs ||
			(i < lenMis && j < lenNjs && mis[i] <= njs[j]) {
			res[k] = mis[i]
			i++
		}
	}

	return res
}

func medianOf(nums []int) float64 {
	l := len(nums)

	if l == 0 {
		panic("Error")
	}

	if l % 2 == 0 {
		return float64(nums[l / 2]+nums[l / 2 - 1]) / 2.0
	}

	return float64(nums[l / 2])
}