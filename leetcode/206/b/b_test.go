// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`4`, `[[1, 2, 3], [3, 2, 0], [3, 1, 0], [1, 2, 0]]`, `[[0, 1], [2, 3]]`, 
			`2`,
		},
		{
			`2`, `[[1], [0]]`, `[[1, 0]]`, 
			`0`,
		},
		{
			`4`, `[[1, 3, 2], [2, 3, 0], [1, 3, 0], [0, 2, 1]]`, `[[1, 3], [0, 2]]`, 
			`4`,
		},
		// TODO 测试参数的下界和上界
		{
			`6`,`[[1,4,3,2,5],[0,5,4,3,2],[3,0,1,5,4],[2,1,4,0,5],[2,1,0,3,5],[3,4,2,0,1]]`, `[[3,1],[2,0],[5,4]]`,
			`5`,
		},
	}
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, unhappyFriends, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-206/problems/count-unhappy-friends/
