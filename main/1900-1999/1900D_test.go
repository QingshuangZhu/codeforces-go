// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1900/problem/D
// https://codeforces.com/problemset/status/1900/problem/D
func TestCF1900D(t *testing.T) {
	testCases := [][2]string{
		{
			`2
5
2 3 6 12 17
8
6 12 8 10 15 12 18 16`,
			`24
203`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, CF1900D)
}