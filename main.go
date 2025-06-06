package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

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

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		panic("Error reading csv file")
	}

	for index, record := range records {
		if index > 0 {
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  toInt(record[5]),
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic("Error converting string to int")
	}

	return i
}

func main() {
	game := &GameState{Score: 0}
	go game.ProcessCSV()
	game.Init()
}
