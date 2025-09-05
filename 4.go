package main

import "fmt"

func countVotes(votes []string) map[string]int {
 result := map[string]int{}
 for _, v := range votes {
  result[v]++
 }
 return result
}

func main() {
 votes := []string{"Карина", "Матвей", "Карина", "Артём", "Карина", "Матвей"}
 cv := countVotes(votes)
 fmt.Println("Подсчёт голосов:", cv)
}
