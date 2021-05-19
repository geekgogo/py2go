package main

import "fmt"

func array() []int {
	var a []int
	for i:=0; i<100; i++{
		a = append(a, i)
	}
	return a
}


func main(){
	sourceData := array()
	for i:=0; i<len(sourceData) - 1; i++ {
		for j:= 0; j<len(sourceData) - 1; j++{
			if sourceData[j] < sourceData[j+1] {
				sourceData[j], sourceData[j+1] = sourceData[j+1], sourceData[j]
			}
		}
	}
	fmt.Println(sourceData)
}
