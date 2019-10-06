package main

import (
	"fmt"
	"math"
)

func max(a, b int)int{
	if a>=b{
		return a
	}
	return b
}

func zeroOneKnapsackArray(c int, w, v []int)[]float64{
	ans := make([]float64, c+1)
	for i:=1;i<len(ans);i++{
		ans[i] = math.Inf(-1)
	}
	for i:=1;i<len(w);i++{
		for j := c; j >= w[i]; j-- {
			if j>=w[i] {
				ans[j] = math.Max(ans[j], float64(v[i])+ans[j-w[i]])
			}
		}
	}
	fmt.Println(ans)
	return ans
}

func zeroOneKnapsackMatrix(c int, w, v []int)[][]int{
	ans := make([][]int, len(v))
	for i:=0;i<len(ans);i++{
		ans[i] = make([]int, c+1)
	}
	for i:=1;i<len(ans);i++ {
		for j := 1; j < c+1; j++ {
			if j < w[i] {
				ans[i][j] = ans[i-1][j]
			} else {
				ans[i][j] = max(ans[i-1][j], v[i]+ans[i-1][j-w[i]])
			}
		}
	}
	fmt.Println(ans)
	return ans
}

func findComponents(c int, w, v []int, dp [][]int)[]int{
	ans := make([]int, 0)
	for i:=len(dp)-1;i>0;i--{
		if dp[i][c] != dp[i-1][c]{
			ans = append(ans, i)
			c = c - w[i]
		}
		if c == 0{
			break
		}
	}
	fmt.Println(ans)
	return ans
}

func completeKnapsack(c int, w, v []int)[]int{
	ans := make([]int, c+1)
	for i:=1;i<len(w);i++{
		for j:=w[i];j<len(ans);j++{
			ans[j] = max(ans[j], ans[j-w[i]]+v[i])
		}
	}
	fmt.Println(ans)
	return ans
}

func completeKnapsackMatrix(c int, w, v []int)[][]int{
	ans := make([][]int, len(w))
	for i:=0;i<len(ans);i++{
		ans[i] = make([]int, c+1)
	}
	for i:=1;i<len(w);i++{
		for j:=w[i];j<c+1;j++{
			if j < w[i] {
				ans[i][j] = ans[i-1][j]
			}else{
				ans[i][j] = max(ans[i-1][j], ans[i][j-w[i]]+v[i])
			}
		}
	}
	fmt.Println(ans)
	return ans
}

func main()  {
	c := 12
	w := []int{0,2,3,4,5}
	v := []int{0,3,4,5,6}
	dp := zeroOneKnapsackMatrix(c, w, v)
	zeroOneKnapsackArray(c, w, v)
	completeKnapsack(c, w, v)
	completeKnapsackMatrix(c, w, v)
	_ = dp
	//findComponents(c, w, v, dp)
}
