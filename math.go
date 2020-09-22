package fraction

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func lcm(a, b int) (res int) {
	res = 1
	if a>b {
		res = a
	}else{
		res = b
	}
	for {
		if res % a == 0 && res % b == 0 {
			break
		}
		res++
	}
	return res
}
