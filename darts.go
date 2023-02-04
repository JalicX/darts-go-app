package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strings"
)

type player struct {
	name          string
	points_scored int
}
type game struct {
	players        []player
	points_to_play int
	throw          int
}

func init_game() game {
	// Get Points to be played
	fmt.Println("Welcome to the Darts GO App!")
	fmt.Println("Please Enter the points to be played.")
	fmt.Println("Common values are 501 or 301:")
	var points int
	_, err := fmt.Scanf("%d\n", &points)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Points are:", points)
	fmt.Println("Error is:", err)
	//var points = 2

	// Get Player-Count
	fmt.Println("Please Enter the count of players:")
	var player_count int
	_, err = fmt.Scanf("%d\n", &player_count)
	if err != nil {
		log.Fatal(err)
	}
	//var player_count = 2

	// Initialize Players
	//var players := make(map[int]player)
	var players []player
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < player_count; i++ {
		fmt.Print("Please enter the name of player", (i + 1), ": ")
		player_name, _ := reader.ReadString('\n')
		player := player{
			name:          player_name,
			points_scored: 0,
		}
		players = append(players, player)
	}

	var game = game{
		players:        players,
		points_to_play: points,
		throw:          0,
	}

	return game

}

func main() {
	fmt.Println("Hello, World!")

	//text = strings.Replace(text, "\n", "", -1)
	//fmt.Println("hello", text)
	var gameconfig = init_game()
	fmt.Println("Gameconfig is:")
	fmt.Println(gameconfig)
	/*if err != nil {
		log.Fatal(err)
	}*/
	fmt.Println("Initialization finished")
}
