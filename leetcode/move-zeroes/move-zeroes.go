package move_zeroes

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]

Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1

Follow up: Could you minimize the total number of operations done?
*/

//func moveZeroes(nums []int) {
//	movedZeroes := 0
//	for j := 0; j < len(nums)-1; j++ {
//		prevMoves := movedZeroes
//		for i := 0; i < len(nums)-1-movedZeroes; i++ {
//			if nums[i] == 0 {
//				moveToTheEnd(nums, i)
//				movedZeroes++
//				break
//			}
//		}
//		if movedZeroes-prevMoves == 0 {
//			break
//		}
//	}
//}

func moveToTheEnd(nums []int, pos int) {
	for i := pos; i < len(nums)-1; i++ {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}

}

func moveZeroes(nums []int) {
	lastNonZerFound := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZerFound], nums[i] = nums[i], nums[lastNonZerFound]
			lastNonZerFound++
		}
	}

}
