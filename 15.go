package main

import "fmt"
import "strings"

type Movie struct {
 Title  string
 Year   int
 Rating float64
 Genres []string
}

func highestRated(movies []Movie) *Movie {
 if len(movies) == 0 {
  return nil
 }
 best := &movies[0]
 for i := 1; i < len(movies); i++ {
  if movies[i].Rating > best.Rating {
   best = &movies[i]
  }
 }
 return best
}

func moviesByGenre(movies []Movie, genre string) []Movie {
 out := []Movie{}
 for _, m := range movies {
  for _, g := range m.Genres {
   if strings.EqualFold(g, genre) {
    out = append(out, m)
    break
   }
  }
 }
 return out
}

func main() {
 movies := []Movie{
  {Title: "ЛЕГО Фильм", Year: 2014, Rating: 7.2, Genres: []string{"мультфильм", "фантастика", "фэнтези", "боевик", "комедия", "приключения", "семейный"}},
  {Title: "ЛЕГО Фильм 2", Year: 2019, Rating: 6.6, Genres: []string{"мультфильм", "комедия", "приключения", "семейный"}},
  {Title: "Выход 8", Year: 2025, Rating: 6.7, Genres: []string{"триллер", "ужасы"}},
  {Title: "Зверополис", Year: 2016, Rating: 8.3, Genres: []string{"мультфильм", "боевик", "комедия", "криминал", "детектив", "приключения", "семейный"}},
  {Title: "Суперсемейка", Year: 2004, Rating: 7.6, Genres: []string{"мультфильм", "фантастика", "семейный", "боевик"}},
 }
 best := highestRated(movies)
 if best != nil {
  fmt.Println("Самый высокий рейтинг:", best.Title, best.Rating)
 }
 action := moviesByGenre(movies, "мультфильм")
 fmt.Println("Мультфильмы:", action)
}
