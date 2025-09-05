package main

import (
 "errors"
 "fmt"
 "strings"
)

func validateUser(name string, age int, email string) error {
 if strings.TrimSpace(name) == "" {
  return errors.New("поле имя не должно быить пустым")
 }
 if len([]rune(name)) >= 50 {
  return errors.New("длина имени должна быть не более 50")
 }
 if age < 18 || age > 120 {
  return errors.New("возраст должен быть от 18 до 120 лет")
 }
 if !strings.Contains(email, "@") {
  return errors.New("email должен содержать '@'")
 }
 return nil
}

func main() {
 fmt.Println("validateUser Иван:", validateUser("Иван", 25, "ivan@example.com"))
 fmt.Println("validateUser invalid:", validateUser("", 130, "у"))
 fmt.Println("validateUser invalid:", validateUser("Аааааааааааааааааааааааааааааааааааааааааааааандрей", 18, "у"))
 fmt.Println("validateUser invalid:", validateUser("Андрей", 130, "у"))
 fmt.Println("validateUser invalid:", validateUser("Андрей", 18, "у"))
 fmt.Println("validateUser invalid:", validateUser("Андрей", 18, "andrey@example.comу"))
}
