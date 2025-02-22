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
		"Alper",
		"Anastasis",
		"Kyrylo",
	}

	rand.NewSource(time.Now().UnixNano())
	shuffledPlayers := shufflePlayers(players)
	teamAA, teamAB := createTeams(shuffledPlayers)
	teamBA, teamBB := createTeams(shufflePlayers(players))

	fmt.Println("Group 1:")
	fmt.Println("Team AA:", teamAA)
	fmt.Println("Team AB:", teamAB)
	fmt.Println("Group 2:")
	fmt.Println("Team BA:", teamBA)
	fmt.Println("Team BB:", teamBB)
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
