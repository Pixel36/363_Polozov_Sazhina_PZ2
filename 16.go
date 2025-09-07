package main

import (
 "fmt"
 "time"
)

type SensorData struct {
 SensorID    string
 Temperature float64
 Humidity    float64
 Timestamp   time.Time
}

func AverageTemperature(data []SensorData) float64 {
 if len(data) == 0 {
  return 0
 }
 var sum float64
 for _, d := range data {
  sum += d.Temperature
 }
 return sum / float64(len(data))
}

func main() {
 var readings []SensorData

 readings = append(readings, SensorData{"Sensor1", 22.5, 40, time.Now()})
 readings = append(readings, SensorData{"Sensor2", 23.0, 42, time.Now()})
 readings = append(readings, SensorData{"Sensor3", 21.8, 38, time.Now()})
 readings = append(readings, SensorData{"Sensor4", 24.1, 45, time.Now()})
 readings = append(readings, SensorData{"Sensor5", 22.9, 41, time.Now()})
 readings = append(readings, SensorData{"Sensor6", 23.3, 43, time.Now()})

 avgTemp := AverageTemperature(readings)

 fmt.Printf("Средняя температура за сутки: %.2f°C\n", avgTemp)
}
