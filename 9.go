package main

import (
 "errors"
 "fmt"
)

type RoomType string
type RoomStatus string

const (
 Single RoomType = "одноместный"
 Double RoomType = "двухместный"
 Suite  RoomType = "люкс"

 Free        RoomStatus = "свободно"
 Booked      RoomStatus = "забронировано"
 Maintenance RoomStatus = "на обслуживании"
)

type HotelRoom struct {
 Type   RoomType
 Status RoomStatus
 Price  float64
}

func bookRoom(rooms map[string]*HotelRoom, number string) error {
 r, ok := rooms[number]
 if !ok {
  return errors.New("комната забронирована")
 }
 if r.Status != Free {
  return errors.New("комната занята")
 }
 r.Status = Booked
 return nil
}

func main() {
 rooms := map[string]*HotelRoom{
  "101": {Type: Single, Status: Free, Price: 2100},
  "102": {Type: Double, Status: Free, Price: 2850},
 }
 if err := bookRoom(rooms, "101"); err != nil {
  fmt.Println("Ошибка бронирования:", err)
 } else {
  fmt.Println("101, статус:", rooms["101"].Status)
 }
}
