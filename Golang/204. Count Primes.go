package leetcode

// 计数质数
// https://leetcode.cn/problems/count-primes/

// 枚举
func isPrime(x int) bool { // 判断是否是质数
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) (cnt int) {
	// 枚举每个数判断其是不是质数
	for i := 2; i < n; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	return
}

// 埃氏筛
func countPrimes1(n int) (cnt int) {
	// 设isPrime[i]表示数i是不是质数，如果是质数则为1，否则为0
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	// 从小到大遍历每个数
	for i := 2; i < n; i++ {
		// 如果这个数为质数，则将其所有的倍数都标记为合数（除了该质数本身），即0
		if isPrime[i] {
			cnt++

			for j := 2 * i; j < n; j += i { // j为i的倍数
				isPrime[j] = false
			}
		}
	}
	return
}

// 线性筛
func countPrimes2(n int) int {
	primes := []int{}
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= n {
				break
			}
			isPrime[i*p] = false
			if i%p == 0 {
				break
			}
		}
	}
	return len(primes)
}
