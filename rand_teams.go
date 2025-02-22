package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	players := []string{
		"Veronika",
		"Roberta",
		"Felipe",
		"Jiri",
		"Filip",
		"Prokhor",
		"Alper",
		"Anastasis",
		"Kyrylo",
	}

	teamAA, teamAB := createTeams(players)
	teamBA, teamBB := createTeams(players)

	fmt.Println("Group 1:")
	fmt.Println("Team AA:", teamAA)
	fmt.Println("Team AB:", teamAB)
	fmt.Println("Group 2:")
	fmt.Println("Team BA:", teamBA)
	fmt.Println("Team BB:", teamBB)
}

func createTeams(players []string) ([]string, []string) {
	players1 := make([]string, len(players))
	copy(players1, players)
	players2 := make([]string, len(players))
	copy(players2, players)

	rand.Shuffle(len(players1), func(i, j int) {
		players1[i], players1[j] = players1[j], players1[i]
	})
	rand.Shuffle(len(players2), func(i, j int) {
		players2[i], players2[j] = players2[j], players2[i]
	})

	teamA := players1[:len(players1)/2]
	teamB := players2[len(players2)/2:]
	return teamA, teamB
}
