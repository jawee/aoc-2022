package daytwo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func B() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/daytwo/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    totalScore := 0
    for scanner.Scan() {
        game := strings.Split(scanner.Text(), " ")
        opponent := game[0]
        outcome := game[1]
        totalScore = totalScore + getScoreForOutcome(opponent, outcome)
    }
    
    fmt.Println(totalScore)
}

func getScoreForOutcome(opponent, outcome string) int {
    if outcome == "Y" {
        return 3 + getScoreForChoice(opponent)
    }
    
    if outcome == "Z" {
        choice := "A"
        if opponent == "A" {
            choice = "B"
        } else if opponent == "B" {
            choice = "C" 
        }
        return 6 + getScoreForChoice(choice)
    }
    
    choice := "A"
    if opponent == "A" {
        choice = "C"
    } else if opponent == "C" {
        choice = "B"
    }

    return 0 + getScoreForChoice(choice)
}
