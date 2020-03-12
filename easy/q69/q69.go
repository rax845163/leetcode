package q69

func mySqrt(x int) int {
	low := 0
	max := x
	for {
		mid := (low + max) / 2
		p := mid * mid
		if p == x {
			return mid
		} else if p < x {

		} else {

		}
	}

	return 0
}
