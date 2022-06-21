package main

func main() {
	
}

//状态压缩dp

func makesquare(matchsticks []int) bool {
	n := len(matchsticks)
	var sum int
	for i:=range matchsticks{
		sum += matchsticks[i]
	}
	if sum % 4 != 0{
		return false
	}
	k := sum/4
	dp := make([]int,1<<n)
	for i:=0;i<len(dp);i++{
		dp[i] = -1
	}
	dp[0] = 0

	for i:=1;i<(1<<n);i++{
		for j:=0;j<n;j++{
			if (i>>j)&1 == 0 {
				continue
			}
			s1 := i&^(1<<j)
			if dp[s1]>=0 && dp[s1]+matchsticks[j] <= k{
				dp[i] = (dp[s1]+matchsticks[j])%k
				break
			}
		}
	}

	return dp[len(dp)-1]==0

}