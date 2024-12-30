package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {

	l := League{
		Teams: []Team{
			{name: "RBS"},
			{name: "Rapid"},
			{name: "Austria"},
			{name: "SV A"},
		},
		Wins: map[string]int{},
	}

	l.MatchResult("RBS", 2, "Rapid", 0)
	l.MatchResult("RBS", 5, "Austria", 0)
	l.MatchResult("RBS", 4, "Rapid", 0)
	l.MatchResult("Austria", 2, "Rapid", 0)
	l.MatchResult("Austria", 2, "Rapid", 3)
	l.MatchResult("Austria", 2, "SV A", 3)

	fmt.Println(l.Ranking())

	w := os.Stdout
	RankPrinter(l, w)
}

func RankPrinter(r Ranker, w io.Writer) {
	for _, v := range r.Ranking() {
		if _, err := io.WriteString(w, v+"\n"); err != nil {
			log.Fatal(err)
		}
	}
}

type Ranker interface {
	Ranking() []string
}

type Team struct {
	name        string
	playerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l League) MatchResult(nameT1 string, scoreT1 int, nameT2 string, scoreT2 int) {
	if scoreT1 > scoreT2 {
		l.Wins[nameT1] = l.Wins[nameT1] + 1
		return
	}

	if scoreT2 > scoreT1 {
		l.Wins[nameT2] = l.Wins[nameT2] + 1
		return
	}

	fmt.Println("draw...")
}

func (l League) Ranking() []string {
	ranking := make([]string, len(l.Teams))

	for i, t := range l.Teams {
		ranking[i] = t.name
	}

	sort.Slice(ranking, func(i, j int) bool {
		return l.Wins[ranking[i]] > l.Wins[ranking[j]]
	})

	return ranking
}