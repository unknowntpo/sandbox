package main


import (
    "fmt"
    "testing"
)

// Hard-coded testing function
func TestIntMinBasic(t *testing.T) { // what does testing.T means?
    ans := IntMin(2, -2)

    if ans != -2 {
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

// Prefered testing function
func TestIntMinTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {2, -2, -2},
        {2, 3, 2},
        {6, 10, 6},
        {2, -7, -7},
        {200, -100, -100},
    }
    for _, test := range tests {
        // check one by one
        if ans := IntMin(test.a, test.b); ans != test.want {
            t.Errof("IntMin(%d, %d) = %d; want %d", test.a, test.b, ans, test.want)
        }
    }
}


