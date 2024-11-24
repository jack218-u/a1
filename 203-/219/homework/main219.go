package main

import (
	"fmt"
	"sort"

	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	Score int
}
type ScoreSlice []Student

func (s ScoreSlice) Len() int {
	return len(s)
}
func (s ScoreSlice) Less(i, j int) bool {
	// return s[i].Score < s[j].Score
	return s[i].Name < s[j].Name
}
func (s ScoreSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func main() {
	var a ScoreSlice
	for i := 1; i <= 10; i++ {
		b := Student{
			Name:  fmt.Sprintf("学生~%d", rand.Intn(11)),
			Age:   rand.Intn(13) + 18,
			Score: rand.Intn(100),
		}
		a = append(a, b)
	}
	fmt.Println(a)
	sort.Sort(a)
	fmt.Println("---排序后---")
	fmt.Println(a)
}
