package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type player struct {
	name          string
	points_scored int
}
type game struct {
	players         []player
	points_to_play  int
	throw           int
	darts_per_throw int
}

func _read_int_from_console(texts []string) (int, error) {
	if len(texts) < 1 {
		log.Fatal("Param-Array 'texts' is empty")
	}
	for i := 0; i < len(texts)-1; i++ {
		fmt.Println(texts[i])
	}
	fmt.Print(texts[len(texts)-1])
	var int_from_console int
	_, err := fmt.Scanf("%d\n", &int_from_console)
	if err != nil {
		log.Fatal(err)
	}
	return int_from_console, nil
}

func _read_players_from_console(player_count int) ([]player, error) {
	var players []player
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < player_count; i++ {
		fmt.Print("Please enter the name of player ", (i + 1), ": ")
		player_name, _ := reader.ReadString('\n')
		player := player{
			name:          player_name,
			points_scored: 0,
		}
		players = append(players, player)
	}

	return players, nil
}

func _init_game(debug bool) (game, error) {
	if debug {
		return game{
			players: []player{
				{name: "Alice",
					points_scored: 0},
				{name: "Bob",
					points_scored: 0},
				{name: "Eve",
					points_scored: 0},
			},
			points_to_play:  501,
			throw:           0,
			darts_per_throw: 3,
		}, nil
	}

	fmt.Println("Welcome to the Darts GO App!")

	// Get Points to be played
	var points, _ = _read_int_from_console([]string{"Please Enter the points to be played.", "Common values are 501 or 301: "})

	// Get Player-Count
	var player_count, _ = _read_int_from_console([]string{"Please Enter the count of players: "})

	// Initialize Players
	var players, _ = _read_players_from_console(player_count)

	var game = game{
		players:         players,
		points_to_play:  points,
		throw:           0,
		darts_per_throw: 3,
	}

	return game, nil

}

func run_round(game game) (game, error) {
	game.throw = game.throw + 1
	fmt.Println("This is game Number", game.throw)
	for playercount := 0; playercount < len(game.players); playercount++ {
		points_of_throw := 0
		for dart := 0; dart < game.darts_per_throw; dart++ {
			prompt := game.players[playercount].name + " please enter the points scored with dart " + strconv.Itoa(dart+1) + " "
			points_of_dart, _ := _read_int_from_console([]string{prompt}) // Add Number of dart to text
			points_of_throw += points_of_dart
		}
		fmt.Println("you have thrown:", points_of_throw, "Points total")
		game.players[playercount].points_scored += points_of_throw
	}
	return game, nil
}

func run_game(game game) {
	again := true
	for again {
		game, _ = run_round(game)
		for playercount := 0; playercount < len(game.players); playercount++ {
			if game.players[playercount].points_scored == game.points_to_play {
				fmt.Println("WINNER WINNER CHICKEN DINNER for", game.players[playercount].name)
				again = false
			}
		}
	}
}

func main() {
	//var DEBUG, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	const DEBUG = true
	var gameconfig, err = _init_game(DEBUG)
	fmt.Println("Mode:", gameconfig.points_to_play, "Points to play")
	if err == nil {
		fmt.Println("Gameconfig sucessfully loaded...")
	}
	run_game(gameconfig)
}
