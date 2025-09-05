package main

import "fmt"

func main(){
	
 a := 0
 b := 0
 c := 0

 fmt.Print("Введите вес основного багажа: ")
 fmt.Scan(&a)
 fmt.Print("Введите вес ручной клади: ")
 fmt.Scan(&b)
 fmt.Print("Введите вес доп. ручной клади: ")
 fmt.Scan(&c)

 fmt.Print("Общий вес = ", a + b + c)
}
