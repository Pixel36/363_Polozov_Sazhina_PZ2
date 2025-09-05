package main

import "fmt"

func totalBaggageWeight(weights []float64) float64 {
 sum := 0.0
 for _, w := range weights {
  sum += w
 }
 return sum
}

func main() {
 weights := []float64{7.5, 2.0, 1.2}
 fmt.Printf("Общий вес багажа: %.3f кг\n", totalBaggageWeight(weights))
}
