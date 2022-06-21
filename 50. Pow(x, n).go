package main

func main() {
	
}


func myPow(x float64, n int) float64 {
	if n<0{
		return fastMux(-n,1/x)
	}
	return fastMux(n,x)
}


//快速幂+递归


func fastMux(n int,x float64)float64{
	if n == 0{
		return 1
	}

	y := fastMux(n/2,x)

	if n%2==0{
		return y*y
	}
	return y*y*x
}