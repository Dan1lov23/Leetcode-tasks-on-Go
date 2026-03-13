/*

	Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

	The overall run time complexity should be O(log (m+n)).

	Example 1:

	Input: nums1 = [1,3], nums2 = [2]
	Output: 2.00000
	Explanation: merged array = [1,2,3] and median is 2.
	Example 2:

	Input: nums1 = [1,2], nums2 = [3,4]
	Output: 2.50000
	Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.


	Constraints:

	nums1.length == m
	nums2.length == n
	0 <= m <= 1000
	0 <= n <= 1000
	1 <= m + n <= 2000
	-106 <= nums1[i], nums2[i] <= 106

*/

package main

import "fmt"

func GetMinIndex(arr []int) int {
	minIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func SortArr(arr []int) []int {
	var sortedArr []int

	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	for len(arrCopy) > 0 {
		minIdx := GetMinIndex(arrCopy)
		sortedArr = append(sortedArr, arrCopy[minIdx])
		arrCopy = append(arrCopy[:minIdx], arrCopy[minIdx+1:]...)
	}

	return sortedArr
}

func MedianOfTwoSortedArrays(nums1 []int, nums2 []int) float64 {

	var result float64

	var newArr []int

	for a := 0; a < len(nums1); a++ {
		newArr = append(newArr, nums1[a])
	}

	for a := 0; a < len(nums2); a++ {
		newArr = append(newArr, nums2[a])
	}

	newArr = SortArr(newArr)

	if len(newArr)%2 == 1 {
		result = float64(newArr[(len(newArr)-1)/2])
		result = result / 1.0
		return result

	} else if len(newArr)%2 == 0 {
		var mid1 = newArr[len(newArr)/2]
		var mid2 = newArr[len(newArr)/2-1]
		result = float64(mid1+mid2) / 2.0
		return result
	}

	return 0

}

func main() {

	result := MedianOfTwoSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(result)

}
