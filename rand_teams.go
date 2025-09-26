package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	players := []string{
		"Veronika",
		"Roberta",
		"Felipe",
		"Jiri",
		"Filip",
		"Prokhor",
		"Anastasis",
		"Kyrylo",
	}

	rand.NewSource(time.Now().UnixNano())
	shuffledPlayers := shufflePlayers(players)
	teamAA, teamAB := createTeams(shuffledPlayers)

	fmt.Println("Let's make sure we do it fair and square and good luck on hunting ghosts!!!")
	fmt.Println("The Haunted House", teamAA)
	fmt.Println("The Haunted House 2", teamAB)
}

func shufflePlayers(players []string) []string {
	shuffled := make([]string, len(players))
	copy(shuffled, players)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}

func createTeams(players []string) ([]string, []string) {
	teamA := players[:len(players)/2]
	teamB := players[len(players)/2:]
	return teamA, teamB
}
