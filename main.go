package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	game := &GameState{}
	game.Init()
}
