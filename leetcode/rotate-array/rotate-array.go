package rotate_array

/*
Given an array, rotate the array to the right by k steps, where k is non-negative.



Example 1:

Input: expectedNums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: expectedNums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]


Constraints:

1 <= expectedNums.length <= 105
-231 <= expectedNums[i] <= 231 - 1
0 <= k <= 105


Follow up:

Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
Could you do it in-place with O(1) extra space?

*/
func rotate(nums []int, k int) {
	if len(nums) < 1 {
		return
	}

	if k < 0 {
		return
	}

	k = k % len(nums)

	result := append(nums[len(nums)-k:], nums[:len(nums)-k]...)

	for i, n := range result {
		nums[i] = n
	}
}
