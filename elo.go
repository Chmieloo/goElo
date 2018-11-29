package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func elo(player1Elo int, player1Matches int, player1Score int, player2Elo int, player2Matches int, player2Score int) (int, int) {
	if player1Matches == 0 {
		player1Matches = 1
	}
	if player2Matches == 0 {
		player2Matches = 1
	}

	if player1Score > player2Score {
		multiplier := math.Log(math.Abs(float64(player1Score-player2Score))+1) * (2.2 / ((float64(player1Elo-player2Elo))*0.001 + 2.2))
		player1Elo, player2Elo = calculateElo(player1Elo, player1Matches, player2Elo, player2Matches, multiplier, false)
	} else if player1Score < player2Score {
		multiplier := math.Log(math.Abs(float64(player1Score-player2Score))+1) * (2.2 / ((float64(player2Elo-player1Elo))*0.001 + 2.2))
		player2Elo, player1Elo = calculateElo(player2Elo, player2Matches, player1Elo, player1Matches, multiplier, false)
	} else {
		player1Elo, player2Elo = calculateElo(player1Elo, player1Matches, player2Elo, player2Matches, 1.0, true)
	}
	return player1Elo, player2Elo
}

func calculateElo(winnerElo int, winnerMatches int, looserElo int, looserMatches int, multiplier float64, isDraw bool) (int, int) {
	constant := 800.0

	Wwinner := 1.0
	Wlooser := 0.0
	if isDraw {
		Wwinner = 0.5
		Wlooser = 0.5
	}
	changeWinner := int((constant / float64(winnerMatches) * (Wwinner - (1 / (1 + math.Pow(10, float64(looserElo-winnerElo)/400))))) * multiplier)
	calculatedWinner := winnerElo + changeWinner

	changeLooser := int((constant / float64(looserMatches) * (Wlooser - (1 / (1 + math.Pow(10, float64(winnerElo-looserElo)/400))))) * multiplier)
	calculatedLooser := looserElo + changeLooser

	return calculatedWinner, calculatedLooser
}

func main() {
	numArguments := len(os.Args) - 1

	if numArguments != 6 {
		panic("Not enough arguments")
	}

	var err error
	nums := make([]int, numArguments)
	for i := 0; i < numArguments; i++ {
		if nums[i], err = strconv.Atoi(os.Args[i+1]); err != nil {
			panic(err)
		}
		fmt.Println(nums[i])
	}

	player1Elo, player2Elo := elo(nums[0], nums[1], nums[2], nums[3], nums[4], nums[5])
	fmt.Println("player 1 ELO: " + strconv.Itoa(player1Elo))
	fmt.Println("player 2 ELO: " + strconv.Itoa(player2Elo))
}
