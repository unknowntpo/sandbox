package divide

import (
	"math"
)

// abs3232 transfer and 32 bit wide integer to it's abs32olute value and return the result as
// unsigned integer
// Because negative of an int32 type math.MinInt32 will overflow, so we use bitwise shift the handle it
func abs32(x int32) uint32 {
	// special case: x == - 2^31
	if x == math.MinInt32 {
		return uint32(1 << 31)
	} else if x < 0 {
		return uint32(-x)
	}
	return uint32(x)
}

// function divide take two number (int32 type) and divide them, returns int32 type result
// when will int32 overflow?
func divide(dividend int32, divisor int32) int32 {
	// special case: -2^31 / 1
	if dividend == math.MinInt32 && divisor == 1 {
		return math.MinInt32
	}

	// special case: -2^31 / -1
	// cause temp overflow
	// Return math.MaxInt32, follow the instruction of leetCode29
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	// take abs32olute value
	var dvd uint32 = abs32(dividend)
	var dvs uint32 = abs32(divisor)
	var temp uint32 = dvs
	var cnt uint32 = 0
	var ans uint32 = 0

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

	if (dividend < 0) != (divisor < 0) {
		return -int32(ans)
	} else {
		return int32(ans)
	}
}
