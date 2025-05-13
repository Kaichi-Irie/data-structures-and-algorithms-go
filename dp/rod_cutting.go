package dp

import "fmt"

type RodCuttingProblem struct {
	N      int
	Prices []int
}

func (p RodCuttingProblem) SolveBottomUp() int {
	dp := make([]int, p.N+1)
	for j := 1; j <= p.N; j++ {
		for i, price := range p.Prices {
			length := i + 1
			if j-length >= 0 {
				dp[j] = max(dp[j], dp[j-length]+price)
			}
		}
	}
	fmt.Println(dp)
	return dp[p.N]
}

func (p RodCuttingProblem) SolveTopDown() int {
	return solve(p.N, p.Prices, make(map[int]int))
}

func solve(n int, prices []int, memo map[int]int) int {
	if n == 0 {
		return 0
	}
	if ans, ok := memo[n]; ok {
		return ans
	}
	var ans int
	for i, price := range prices {
		length := i + 1
		if n-length >= 0 {
			ans = max(ans, solve(n-length, prices, memo)+price)
		}
	}
	return ans
}
