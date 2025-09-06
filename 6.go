package main

import "fmt"

func main(){
	tags :=[][]string {{"go","backend"}, {"git","go","tools"}}
	tagsMap :=map[string]bool{}
	for _,tags := range tags{
		for _,tag := range tags{
			tagsMap[tag]=true
		} 
	}
	fmt.Println(tags)
	fmt.Println(tagsMap)
}
