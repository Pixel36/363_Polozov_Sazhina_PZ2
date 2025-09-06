package main

import (
    "fmt"
)

type Product struct {
    Name     string
    Category string
    Price    float64
}

var products = []Product{
    {"Ноутбук", "Электроника", 15000.0},
    {"Смартфон", "Электроника", 10000.0},
    {"Футбоьный мяч", "Спорт", 1000.0},
    {"Монитор", "Электроника", 12000.0},
    {"Повербанк", "Электроника", 3000.0},
}

func filter(products []Product, maxPrice float64, category string) []Product {
    var filtered []Product
    for _, p := range products {
        if p.Price < maxPrice && p.Category == category {
            filtered = append(filtered, p)
        }
    }
    return filtered
}

func main() {
    result := filter(products, 10000.0, "Электроника")
    for _, p := range result {
        fmt.Println(p.Name, p.Category, p.Price)
    }
}
