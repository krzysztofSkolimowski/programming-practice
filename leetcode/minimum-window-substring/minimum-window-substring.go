package minimum_window_substring

/*
Given two strings s and t of lengths m and n respectively,
\return the minimum window substring of s such that every character in t (including duplicates)
is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

A substring is a contiguous sequence of characters within the string.



Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.


Constraints:

m == s.length
n == t.length
1 <= m, n <= 105
s and t consist of uppercase and lowercase English letters.


Follow up: Could you find an algorithm that runs in O(m + n) time?


hint:
Use two pointers to create a window of letters in S, which would have all the characters from T.

hint:
Since you have to find the minimum window in S which has all the characters from T, you need to expand and contract the window using the two pointers and keep checking the window for all the characters. This approach is also called Sliding Window Approach.

L ------------------------ R , Suppose this is the window that contains all characters of T

        L----------------- R , this is the contracted window. We found a smaller window that still contains all the characters in T

When the window is no longer valid, start expanding again using the right pointer.

*/

//func minWindow(s string, t string) string {
//	if len(s) < 1 {
//		return ""
//	}
//
//	if len(t) > 105 {
//		return ""
//	}
//
//	for windowSize := len(t); windowSize <= len(s); windowSize++ {
//		for beginning := 0; beginning < len(s)-windowSize+1; beginning++ {
//			end := beginning + windowSize - 1
//
//			window := s[beginning : end+1]
//
//			m := map[rune]int{}
//			for _, v := range t {
//				m[v]++
//			}
//
//			for _, ch := range window {
//				v, ok := m[ch]
//				if ok && v > 0 {
//					m[ch]--
//				}
//			}
//
//			mapSum := 0
//			for _, v := range m {
//				mapSum += v
//			}
//
//			if mapSum == 0 {
//				return s[beginning : end+1]
//			}
//
//		}
//
//	}
//
//	return ""
//}

func minWindow(s string, t string) string {

	start := 0
	end := len(s)

	for end-start < len(s) {
		window := s[start:end]
		if containsCharacters(window, t) {
			start++
		}
	}

	return ""
}

func containsCharacters(window string, t string) bool {
	//todo
	return false
}
