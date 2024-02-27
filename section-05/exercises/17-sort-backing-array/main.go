package main

import (
	"fmt"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort the backing array
//
//  1. Sort only the middle 3 items.
//
//  2. All the slices should see your changes.
//
//
// RESTRICTION
//
//  Do not sort manually. Sort by using the sort package.
//
//
// EXPECTED OUTPUT
//
//  Original: [pacman mario tetris doom galaga frogger asteroids simcity metroid defender rayman tempest ultima]
//
//  Sorted  : [pacman mario tetris doom galaga asteroids frogger simcity metroid defender rayman tempest ultima]
//
//
// HINT:
//
//   Middle items are         : [frogger asteroids simcity]
//
//   After sorting they become: [asteroids frogger simcity]
//
// ---------------------------------------------------------

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	mid := len(items) / 2

	// [1,2,3,4,5] => mid is 2. [2-1 : 2+4] => [2,3,4]
	// [1,2,3,4] => mid is 2. [2-1 : 2+4] =>
	from := mid - 1
	var to int
	if len(items)%2 == 0 {
		to = mid + 1
	} else {
		to = mid + 2
	}

	sortableData := items[from:to]
	sort.Strings(sortableData)
	fmt.Println()
	fmt.Println("Sorted  :", items)
}
