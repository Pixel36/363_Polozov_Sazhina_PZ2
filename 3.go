package main

import "fmt"

type Order struct {
 ID          int
 Items       []int
 Total       float64
 Address     string
 IsCompleted bool
}

func addOrder(m map[int]Order, o Order) {
 m[o.ID] = o
}

func main() {
 orders := make(map[int]Order)
 o1 := Order{ID: 1, Items: []int{101, 102}, Total: 3500.50, Address: "Ленина, 1", IsCompleted: false}
 addOrder(orders, o1)
 fmt.Println("Список заказов:", orders)
}
