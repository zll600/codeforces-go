// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/arc109/submit?taskScreenName=arc109_c
func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`3 2
RPS`,
			`P`,
		},
		{
			`11 1
RPSSPRSPPRS`,
			`P`,
		},
		{
			`1 100
S`,
			`S`,
		},
		// TODO 测试参数的下界和上界
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/arc109/tasks/arc109_c
