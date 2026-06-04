// prime_test.go

/*
 To run this, in the same directory type this command:

   $ go test PASS ok      examples/primes 0.180s

The "go test" automatically looks for files then with _test.go, and runs their
test functions.

A test function name must start with Test, and *testing.T as a parameter.

*/

package main

import "testing"

func TestIsPrime(t *testing.T) {
    // tests := []struct {
    //     n        int
    //     expected bool
    // }{
    //     {2, true},
    //     {3, true},
    //     {4, false},
    //     {17, true},
    //     {100, false},
    //     {1, false},
    //     {0, false},
    //     {-5, false},
    // }

    // for _, tt := range tests {
    //     result := IsPrime(tt.n)
    //     if result != tt.expected {
    //         t.Errorf("IsPrime(%d) = %v, want %v", tt.n, result, tt.expected)
    //     }
    // }
	some_primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 37, 101}
	some_non_primes := []int{-2, -1, 0, 1, 4, 6, 8, 9, 10, 25, 1001}

	// check primes
	for _, p := range some_primes {
		result := IsPrime(p)
		if result == false {
			t.Errorf("IsPrime(%v) = false, should be true", p)
		}
	}

	// check non-primes
	for _, n := range some_non_primes {
		result := IsPrime(n)
		if result == true {
			t.Errorf("IsPrime(%v) = true, should be false", result)
		}
	}
} // TestIsPrime


func TestPrimesLessThan(t *testing.T) {
	tests := [] struct {
		n int
		count int
	}{
		{-1, 0},
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 3},
		{7, 3},
		{8, 4},
		{99, 25},
		{100, 25},
		{101, 25},
		{1000, 168},
		{10000, 1229},
	}

	for _, tc := range tests {
		result := PrimesLessThan(tc.n)
		if result != tc.count {
			t.Errorf("PrimesLessThan(%v) = %v, should be %v", tc.n, result, tc.count)
		}
	}
} // TestPrimesLessThan
