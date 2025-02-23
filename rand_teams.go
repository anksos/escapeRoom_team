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
	fmt.Println("The Star Element:", teamAA)
	fmt.Println("Apocalypse Zombie 2213:", teamAB)
	fmt.Println("Group 2:")
	fmt.Println("Moriarty's Phantom Trap 1:", teamBA)
	fmt.Println("Moriarty's Phantom Trap 2:", teamBB)
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
