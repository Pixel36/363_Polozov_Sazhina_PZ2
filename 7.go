package main

import "fmt"

type Employee struct {
 ID       int
 Name     string
 Position string
 Salary   float64
}

func payrollStats(employees []Employee) (total float64, avg float64) {
 for _, e := range employees {
  total += e.Salary
 }
 if len(employees) > 0 {
  avg = total / float64(len(employees))
 }
 return
}
func main() {
 emps := []Employee{
  {ID: 1, Name: "Карина", Position: "Dev", Salary: 50000},
  {ID: 2, Name: "Матвей", Position: "QA", Salary: 40000},
 }
 total, avg := payrollStats(emps)
 fmt.Printf("Общий фонд: %.2f, средняя зарплата: %.2f\n", total, avg)
}
