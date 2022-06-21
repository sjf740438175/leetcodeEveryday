package main

import "math/rand"

func main() {

}



type Solution struct {
	rects [][]int
	each_r []int
}


func Constructor(rects [][]int) Solution {
	n := len(rects)
	each_r := make([]int,n+1)
	each_r[0]=0
	for i:= range rects{
		each_r[i+1] = each_r[i]+(rects[i][2]-rects[i][0]+1)*(1+rects[i][3]-rects[i][1])
	}


	return Solution{rects: rects,each_r: each_r}

}


func (this *Solution) Pick() []int {
	all := this.each_r[len(this.each_r)-1]
	x := rand.Intn(all)
	i := 0
	for ;i<len(this.each_r) && x < this.each_r[i];i++{
	}
	i--
	least := all - this.each_r[i]
	l_x := this.rects[i][2]-this.rects[i][0]

	tx := this.rects[i][0]+least%l_x
	ty := this.rects[i][1]+least/l_x
	return []int{tx,ty}
}