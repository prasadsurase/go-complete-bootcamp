// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "Name": "god of war",
//                  "Genre": "action adventure",
//                  "Price": 50
//          },
//          {
//                  "id": 2,
//                  "Name": "x-com 2",
//                  "Genre": "strategy",
//                  "Price": 40
//          },
//          {
//                  "id": 3,
//                  "Name": "minecraft",
//                  "Genre": "sandbox",
//                  "Price": 20
//          }
//  ]
//
// ---------------------------------------------------------

type Item struct {
	Id    int    `json:"id"`
	Name  string `json:"Name"`
	Price int    `json:"Price"`
}

type Game struct {
	Item  `json:"item"`
	Genre string `json:"Genre"`
}

func main() {
	games := []Game{
		{Item: Item{Id: 1, Name: "god of war", Price: 50}, Genre: "action adventure"},
		{Item: Item{Id: 2, Name: "x-com 2", Price: 30}, Genre: "strategy"},
		{Item: Item{Id: 3, Name: "minecraft", Price: 20}, Genre: "sandbox"},
	}

	fmt.Printf("Games %v\n", games)

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		fmt.Printf(`
			> list   : lists all the games
			> save   : save to json file
			> quit   : quits
		`)
		if !in.Scan() {
			break
		}

		fmt.Println()

		switch in.Text() {
		case "list":
			for _, game := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					game.Id, game.Name, "("+game.Genre+")", game.Price)
			}
		case "save":
			s, err := json.Marshal(games)
			if err != nil {
				log.Print(err)
			}
			fmt.Println(string(s))
			continue
		case "quit":
			return
		}
	}
}
