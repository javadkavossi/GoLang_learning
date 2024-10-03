package sqrt

import (
	"errors"
)

// Abs محاسبه مقدار مطلق یک عدد اعشاری
func Abs(val float64) float64 {
	if val < 0 {
		return -val
	}
	return val
}

// Sqrt محاسبه ریشه دوم یک عدد
func Sqrt(val float64) (float64, error) {
	// اگر ورودی منفی است، خطا برگردانده می‌شود
	if val < 0.0 {
		return 0, errors.New("cannot compute the square root of a negative number")
	}
	if val == 0.0 {
		return 0.0, nil
	}

	// الگوریتم نیوتن-رافسون برای تقریب ریشه دوم
	guess, epsilon := 1.0, 0.00001
	for i := 0; i < 10000; i++ {
		if Abs(guess*guess-val) < epsilon {
			return guess, nil
		}
		guess = (val/guess + guess) / 2.0
	}
	return guess, nil
}
