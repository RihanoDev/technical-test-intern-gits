package main

import "fmt"

func denseRanking(scores []int, playerScores []int) []int {

	uniqueScores := []int{}
	for _, score := range scores {
		if len(uniqueScores) == 0 || uniqueScores[len(uniqueScores)-1] != score {
			uniqueScores = append(uniqueScores, score)
		}
	}

	ranks := make([]int, len(playerScores))

	for i, playerScore := range playerScores {
		rank := 1
		for _, uniqueScore := range uniqueScores {
			if playerScore >= uniqueScore {
				break
			}
			rank++
		}
		ranks[i] = rank
	}

	return ranks
}

func main() {

	score := []int{100, 100, 50, 40, 40, 20, 10}
	gitScore := []int{5, 25, 50, 120}

	result := denseRanking(score, gitScore)

	for _, rank := range result {
		fmt.Print(rank, " ")
	}
}
