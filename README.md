# golang_longest_common_subsequence

Given two strings `text1` and `text2`, return *the length of their longest **common subsequence**.* If there is no **common subsequence**, return `0`.

A **subsequence** of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

- For example, `"ace"` is a subsequence of `"abcde"`.

A **common subsequence** of two strings is a subsequence that is common to both strings.

## Examples

**Example 1:**

```
Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.

```

**Example 2:**

```
Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.

```

**Example 3:**

```
Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.

```

**Constraints:**

- `1 <= text1.length, text2.length <= 1000`
- `text1` and `text2` consist of only lowercase English characters.

## 解析

給定兩個字串 text1, text2

要求寫一個演算法找出 text1, text2 最大共同子字串的長度

這邊定義每個字串 s 的子字串 是從 s 的字元中刪除幾個字元所形成的字串

共同子字串 common_substring 代表 同時是 text1, text2 的子字串

len(common_substring) ≤ min(len(text1), len(text2))

要透過動態規劃的作法

先思考如何去找出該問題的子問題

舉例來說： text1: “abcde”, text2: “ace”

當發現 兩個字串的第一個字元相等 

找 “abcde” , “ace” 最長子字串長度 = 1 + “bcde”, “ce” 最長子字串長度

![](https://i.imgur.com/9OwbR4f.png)

如果是順向去找會發現有些字串會重複查找

可以從反向來思考

兩個字串都以從第i 個位置到最後逐步考慮更長字串所能找到的子字串

首先定義 d[i][j] 

第1個字串從第i個位置開始從第 i 字元開始與

第2個字串 從第 j 個開始所形成的最長字串長度

可以發現

d[i][j] = d[i+1][j+1]+1 , if text1[i] == text2[j]

d[i][j] = max(d[i+1][j], d[i][j+1]) if text1[i] ≠ text2[j]  

![](https://i.imgur.com/zsrG7i3.png)

## 程式碼
```go
package sol

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for row := range dp {
		dp[row] = make([]int, n+1)
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// start from last Idx reduce the repetition
	// dp[i][j] = dp[i+1][j+1] if text1[i] == text2[j]
	// dp[i][j] = max(dp[i+1][j], dp[i][j+1]) if text1[i] != text2[j]
	for text1_start := m - 1; text1_start >= 0; text1_start-- {
		for text2_start := n - 1; text2_start >= 0; text2_start-- {
			if text1[text1_start] == text2[text2_start] {
				dp[text1_start][text2_start] = dp[text1_start+1][text2_start+1] + 1
			} else {
				dp[text1_start][text2_start] = max(
					dp[text1_start+1][text2_start],
					dp[text1_start][text2_start+1],
				)
			}
		}
	}
	return dp[0][0]
}
```
## 困難點

1. 要能找出最長子字串的遞迴子關係

## Solve Point

- [x]  建立一個 m+1 by n+1 的整數矩陣 dp 用來儲存動態規劃的中間計算結果
- [x]  透過 dp[i][j] 的遞迴關係式從 i = m+1, j = n+1 逐步往前推算
- [x]  回傳 dp[0][0]