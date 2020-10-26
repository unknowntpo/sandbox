package divide

import (
	"fmt"
	"math"
)

// abs3232 transfer and 32 bit wide integer to it's abs32olute value and return the result as
// unsigned integer
// Because negative of an int32 type math.MinInt32 will overflow, so we use bitwise shift the handle it
func abs32(x int32) uint32 {
	// special case: x == - 2^31
	if x == math.MinInt32 {
		return uint32(1 << 31)
	} else {
		return uint32(-x)
	}
}

// function divide take two number (int32 type) and divide them, returns int32 type result
// when will int32 overflow?
func divide(dividend int32, divisor int32) int32 {
	/* Pseudo code
	IF (out-of bound)
	    RETURN INT_MAX

	DETEMINE sign flag

	dvd <- abs32(dividend)
	dvs <- abs32(divisor)

	WHILE dvd >= dvs
		temp <- dvs
		cnt <- 1
		WHILE temp << 1 <= dvd
		    temp <- temp << 1
		    cnt <- cnt << 1
		ENDWHILE
		ans <- ans + cnt
		dvd <- dvd - temp
	ENDWHILE

	*/

	// special case: -2^31 / 1
	if dividend == math.MinInt32 && divisor == 1 {
		return math.MaxInt32
	}

	// take abs32olute value
	var dvd uint32 = abs32(dividend)
	var dvs uint32 = abs32(divisor)
	var temp uint32 = dvs
	var cnt uint32 = 0
	var ans int32 = 0
	for dvd >= dvs {
		temp = dvs
		cnt = 1
		for dvd >= (temp << 1) {
			temp <<= 1
			cnt <<= 1
		}
		dvd -= temp
		ans += cnt
	}

	if (dividend < 0) ^ (divisor < 0) {
		return -ans
	} else {
		return ans
	}
}
func TestMain(t *testing.T) {
	var tests = []struct {
		dividend, divisor int32
		want              int32
	}{
		{-2147483648, 1, -2147483648},
		{-2147483648, -1, 2147483647},
		{10, 3, 3},
		{7, -3, -2},
		{0, 1, 0},
		{1, 1, 1},
	}
	for _, test := range tests {
		if ans := divide(test.dividend, test.divisor); ans != test.want {
			t.Errof("divide(%d, %d) = %d; want %d", test.dividend, test.divisor, ans, test.want)
		}
	}
}
func main() {
	fmt.Println("vim-go")
}
