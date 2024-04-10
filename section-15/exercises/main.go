package main

import (
	"fmt"
	"strconv"
	"time"
)

type money float64

type book struct {
	title       string
	price       money
	publishedAt interface{}
}

type game struct {
	title string
	price money
}

type puzzle struct {
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

func (b *book) print() {
	fmt.Printf("%-15s: %s - (%v)\n", b.title, b.price.string(), formatDate(b.publishedAt))
}

func formatDate(data interface{}) string {
	// if data == nil {
	// 	return "unknown"
	// }
	// if data, ok := data.(int); ok {
	// 	var t int
	// 	t = data
	// 	u := time.Unix(int64(t), 0)
	// 	return u.String()
	// }

	// if data, ok := data.(string); ok {
	// 	var t int
	// 	t, _ = strconv.Atoi(data)
	// 	u := time.Unix(int64(t), 0)
	// 	return u.String()
	// }
	var t int

	switch e := data.(type) {
	case int:
		t = e
	case string:
		t, _ = strconv.Atoi(e)
	default:
		return "unknown"
	}

	u := time.Unix(int64(t), 0)
	return u.String()
}

func (g *game) print() {
	fmt.Printf("%-15s: %s\n", g.title, g.price.string())
}

func (p *puzzle) print() {
	fmt.Printf("%-15s: %s\n", p.title, p.price.string())
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

func (l list) discount(ratio float64) {
	type discounter interface {
		discount(float64)
	}

	for _, it := range l {
		// g, hasDiscount := it.(*game)
		// g, hasDiscount := it.(interface{ discount(float64) })
		g, hasDiscount := it.(discounter)
		if hasDiscount {
			g.discount(ratio)
		}
	}
}

func main() {
	var (
		mobydick       = book{title: "moby dick", price: 10, publishedAt: 1712755836}
		treasureIsland = book{title: "treasure island", price: 25, publishedAt: "1711891860"}
		oliverTwist    = book{title: "oliver twist", price: 15}
		minecraft      = game{title: "minecraft", price: 20}
		tetris         = game{title: "tetris", price: 5}
		rubik          = puzzle{title: "rubik's cube", price: 5}
	)

	var store list
	store = append(store, &minecraft, &tetris, &mobydick, &rubik, &treasureIsland, &oliverTwist)
	// store.print()
	store.discount(0.5)
	store.print()
}
