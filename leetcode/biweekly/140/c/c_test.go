// Generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_c(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, validSequence, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/biweekly-contest-140/problems/find-the-lexicographically-smallest-valid-sequence/
// https://leetcode.cn/problems/find-the-lexicographically-smallest-valid-sequence/