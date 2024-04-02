package main

import "fmt"

type money float64

type book struct {
	title string
	price money
}

type game struct {
	title string
	price money
}

type printer interface {
	print()
}

type list []printer

func (m money) string() string {
	return fmt.Sprintf("$%.2f", float64(m))
}

func printBook(b book) {
	fmt.Printf("%-15s: %s\n", b.title, b.price.string())
}

func (b *book) print() {
	fmt.Printf("%-15s: %s\n", b.title, b.price.string())
}

func printGame(g game) {
	fmt.Printf("%-15s: %s\n", g.title, g.price.string())
}

func (g *game) print() {
	fmt.Printf("%-15s: %s\n", g.title, g.price.string())
}

func (g *game) discount(ratio float64) {
	g.price *= money(1 - ratio)
}

func (b *book) discount(ratio float64) {
	b.price *= money(1 - ratio)
}

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("No games available")
		return
	}
	for _, it := range l {
		it.print()
	}
}

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	// printBook(mobydick)
	// printGame(minecraft)
	// printGame(tetris)

	// minecraft.discount(0.5)
	// mobydick.print()
	// minecraft.print()
	// tetris.print()

	// var items []*game
	// items = append(items, &minecraft, &tetris, &mobydick)

	// myList := list(items)
	// myList.print()

	var store list
	store = append(store, &minecraft, &tetris, &mobydick)
	store.print()
}
