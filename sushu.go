package main

import (
	"fmt"
	"time"
)

//获取100以内的素数
func suShu(sum int) []int {
	var a []int
	for i:=1; i<=sum; i++ {
		if i > 1{
			num := i / 2
			for {
				if num > 1{
					if i % num == 0{
						break
					}else{
						num --
					}
				}else{
					a = append(a, i)
					break
				}
			}
		}else {
			a = append(a, i)
		}
	}
	return a
}

func main(){
	st := time.Now()
	result := suShu(1000)
	elapsed := time.Since(st)
	fmt.Println(result)
	fmt.Printf("耗时：%v\n", elapsed)
}
