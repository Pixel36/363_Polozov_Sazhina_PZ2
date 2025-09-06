package main

import (
 "fmt"
 "strings"
 "unicode"
)

type TextStats struct {
 Chars     int
 Words     int
 Sentences int
}

func textStats(text string) TextStats {
 chars := 0
 for range text {
  chars++
 }
 words := len(strings.Fields(text))
 sentences := 0
 for _, r := range text {
  if r == '.' || r == '!' || r == '?' {
   sentences++
  }
 }
 if strings.TrimSpace(text) == "" {
  words = 0
  sentences = 0
  chars = 0
 }
 _ = unicode.IsLetter
 return TextStats{Chars: chars, Words: words, Sentences: sentences}
}

func main() {
 txt := "Привет! Это тест. Сколько предложений?"
 ts := textStats(txt)
 fmt.Println("TextStats:", ts)
}
