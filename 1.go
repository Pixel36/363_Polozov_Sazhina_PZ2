package main

import (
 "fmt"
 "time"
)

func totalHotelCost(dates []string) int {
 const weekdayPrice = 2100
 const weekendPrice = 2850
 total := 0
 layout := "2006-01-02"
 for _, ds := range dates {
  t, err := time.Parse(layout, ds)
  if err != nil {
   continue
  }
  wd := t.Weekday()
  if wd >= time.Monday && wd <= time.Thursday {
   total += weekdayPrice
  } else {
   total += weekendPrice
  }
 }
 return total
}

func main() {
 dates := []string{"2025-09-01", "2025-09-02", "2025-09-05"}
 fmt.Println("Стоимость проживания:", totalHotelCost(dates))
}
