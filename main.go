package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

const YellowANSICode = "\033[33m"
const TimeToAnswer = time.Second * 3

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Name      string
	Score     int
	Questions []Question
}

func (g *GameState) Init() {
	fmt.Println("Welcome to the game!")
	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading name")
	}

	g.Name = name

	fmt.Printf("Let's start the game, %s", g.Name)
}

func (g *GameState) ProcessCSV() {
	file, err := os.Open("quiz.csv")

	if err != nil {
		panic("Error opening file")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("Error closing file")
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		panic("Error reading csv file")
	}

	for index, record := range records {
		if index > 0 {
			correctAnswer, _ := toInt(record[5])

			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0, errors.New("invalid input. Your answer must be a number")
	}

	return i, nil
}

func (g *GameState) Run() {
	answerCh := make(chan int)

	for i, question := range g.Questions {
		fmt.Println("--------------------------------------------------")
		fmt.Printf("%s %d. %s %s\n", YellowANSICode, i+1, question.Text, YellowANSICode)

		for j, option := range question.Options {
			fmt.Printf("[%d] - %s\n", j+1, option)
		}

		fmt.Println("Your answer (you have 10 seconds to answer):")

		go func() {
			for {
				reader := bufio.NewReader(os.Stdin)

				input, _ := reader.ReadString('\n')
				answer, err := toInt(input[:len(input)-1])

				if err == nil {
					answerCh <- answer
				}
			}
		}()

		select {
		case answer := <-answerCh:
			if answer == question.Answer {
				fmt.Println("Congrats, you got it right!")
				g.Score += 10
			} else {
				fmt.Println("Sorry, you got it wrong!")
			}

			continue

		case <-time.After(5 * time.Second):
			fmt.Println("\nTime's up! Moving to the next question.")
			continue
		}
	}
}

func main() {
	game := &GameState{Score: 0}
	game.ProcessCSV()
	game.Init()
	game.Run()

	fmt.Printf("Game over, Your final score is %d\n", game.Score)
}
