// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`[[0,1],[1,2]]`, `[0,2,1]`, 
			`8`,
		},
		{
			`[[0,1],[0,2],[1,2]]`, `[0,10,10]`, 
			`3`,
		},
		
	}
	targetCaseNum := 1
	if err := testutil.RunLeetCodeFuncWithExamples(t, networkBecomesIdle, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-63/problems/the-time-when-the-network-becomes-idle/
