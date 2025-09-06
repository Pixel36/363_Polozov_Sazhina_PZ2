package main

import (
 "fmt"
 "strconv"
)

func decToBinHex(n int64) (string, string) {
 return strconv.FormatInt(n, 2), strconv.FormatInt(n, 16)
}

func convertBase(s string, fromBase int, toBase int) (string, error) {
 val, err := strconv.ParseInt(s, fromBase, 64)
 if err != nil {
  return "", err
 }
 return strconv.FormatInt(val, toBase), nil
}

func main() {
 bin, hex := decToBinHex(42)
 fmt.Println("42 -> bin:", bin, "hex:", hex)
 conv, cerr := convertBase("101010", 2, 16)
 fmt.Println("101010(2) ->", conv, "err:", cerr)
}
