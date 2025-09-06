package main

import "fmt"

func demoBudget() map[string]float64 {
 m := map[string]float64{
  "Еда":         15000,
  "Транспорт":   5000,
  "Развлечения": 3000,
 }

 m["Еда"] += 2000
 return m
}

func totalBudget(m map[string]float64) float64 {
 total := 0.0
 for _, v := range m {
  total += v
 }
 return total
}

func main() {
 b := demoBudget()
 fmt.Println("Траты:", b)
 fmt.Println("Бюджет:", totalBudget(b))
}
