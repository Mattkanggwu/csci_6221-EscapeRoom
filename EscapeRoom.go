package main //Main package.

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
		fmt.Println("The door is locked. You need a key to open it.")
		return
	} else if nextRoom != "" {
		if (nextRoom == "north" || nextRoom == "west" || nextRoom == "east") && !contains(g.Player.Items, "key") {
			fmt.Println("The door is locked. You need a key to open it.")
			return
		}

		if nextRoom == "north" && !contains(g.Player.Items, "key") {
			g.Search()
			return
		}

		if (nextRoom == "west" || nextRoom == "east") && contains(g.Player.Items, "key") {
			fmt.Println("You used the key to unlock the door.")
			g.Player.Items = removeItem(g.Player.Items, "key")
			if nextRoom == "west" {
				fmt.Println("Now, you can enter the west room.")
			}
		} else if (nextRoom == "west" || nextRoom == "east") && !contains(g.Player.Items, "key") {
			fmt.Println("The door is locked. You need a key to open it.")
			return
		}

		g.Player.Room = nextRoom
		fmt.Println("Now, you are in the", nextRoom, "room.")

		if nextRoom == "north" {
			fmt.Println("Now, there are two doors again: south (starting room) and west.")
		} else if nextRoom == "east" {
			fmt.Println("Congratulations! You have escaped the room. Input 'quit' to end this game.")
			return
		}
	} else {
		fmt.Println("You can't go that way.")
	}
}
func main() { // initial game state and run the game loop
	game := Game{
		Player: Player{
			Room:  "start",
			Items: []string{},
		},
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Escape Room Game!")
	fmt.Println("You can search each room to get the key because the door is locked.")
	fmt.Println("Now, you are in the starting room. Please, input the 'search' to find the key and there are two doors for north and east, select one")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch strings.TrimSpace(input) { // handles different commands based on the user's input
		case "search":
			game.Search()
		case "north", "south", "east", "west": // user enters a direction, move() method of the game is called to move the player to the corresponding room.

			game.Move(input)
		case "quit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid command. Type 'search' to search the room.")
		}
	}
}

func canSearchRoom(room string) bool { // this function determines if the given room is searchable or not.
	searchableRooms := []string{"start", "north", "south"} // string representing the current room and returns a boolean value
	return contains(searchableRooms, room)
}

func getRoomExit(room string, direction string) string { // it is responsible for retrieving the next room based on the current room
	// and the direction the player wants to move in
	exits := map[string]map[string]string{
		"start": {
			"north": "north",
			"east":  "east",
		},
		"north": {
			"south": "start",
			"west":  "game_over",
		},
		"south": {
			"north": "start",
		},
		"east": {
			"west": "start",
		},
	}

	return exits[room][direction]
}

func contains(slice []string, item string) bool { //contains helper function is used within the cansearchroom() and move() to check if
	// an item exists in a slice.
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func removeItem(slice []string, item string) []string {
	index := -1
	for i, s := range slice {
		if s == item {
			index = i

			break
		}
	}
	if index != -1 {
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice

}
