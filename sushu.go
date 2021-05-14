package main

import "fmt"

//获取100以内的素数
func main(){
	a := []int{1}
	for i:=0; i<=100; i++ {
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
		}
	}
	fmt.Println(a)
}
