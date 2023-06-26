package main //Main package

import (
	"bufio"   // bufio provides buffered I/O operations
	"fmt"     // it is used for formatted input and output operations
	"os"      // provides a platform-independent interface to the operating system
	"strings" // provides functions for manipulating strings
)

type Player struct { //defines a struct Player that represents the player in the game
	Room  string   // Room -> current room of the player
	Items []string // items that player collect
}

type Game struct { //defines a struct Game that represents the game itself.
	Player Player
}

func (g *Game) Search() { /*  defines a method Search() on the Game struct.
	It allows the player to search the current room for a key.*/
	if canSearchRoom(g.Player.Room) {
		if !contains(g.Player.Items, "key") {
			g.Player.Items = append(g.Player.Items, "key")
			fmt.Println("You found a key!")
		} else {
			fmt.Println("You have already found a key in this room.")
		}
	} else {
		fmt.Println("There is nothing to search in this room.")
	}
}

func (g *Game) Move(direction string) {
	nextRoom := getRoomExit(g.Player.Room, direction)
	if nextRoom == "game_over" {
		fmt.Println("Game Over! You entered the wrong room.")
		return
	} else if nextRoom != "" {
		g.Player.Room = nextRoom
		if nextRoom == "north" && !contains(g.Player.Items, "key") {
			g.Search()
		}
		if nextRoom == "east" && contains(g.Player.Items, "key") {
			fmt.Println("Congratulations! You have escaped the room.")
			return
		}
	} else {
		fmt.Println("You can't go that way.")
	}
}

func main() {
	game := Game{
		Player: Player{
			Room:  "start",
			Items: []string{},
		},
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Escape Room Game! You can search each room to get the key because the door is locked. Now, you are in the starting room.")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "search":
			game.Search()
		case "north", "south", "east", "west":
			game.Move(input)
		case "quit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid command. Type 'search' to search the room.")
		}
	}
}

func canSearchRoom(room string) bool {
	searchableRooms := []string{"start", "north", "south"}
	return contains(searchableRooms, room)
}

func getRoomExit(room string, direction string) string {
	exits := map[string]map[string]string{
		"start": {
			"north": "north",
			"east":  "east",
		},
		"north": {
			"south": "start",
		},
		"south": {
			"north": "start",
		},
		"east": {
			"west": "start",
		},
	}

	if roomExits, ok := exits[room]; ok {
		if nextRoom, ok := roomExits[direction]; ok {
			return nextRoom
		}
	}
	return ""
}

func contains(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}
