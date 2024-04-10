package main

import (
	"fmt"
	"strconv"
	"time"
)

type money float64

type product struct {
	title string
	price money
}

func (p *product) print() {
	fmt.Printf("%-15s: %s\n", p.title, p.price.string())
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}

type book struct {
	product
	publishedAt interface{}
}

type game struct {
	product
}

type puzzle struct {
	product
}

type item interface {
	print()
	discount(float64)
}

type list []item

func (m money) string() string {
	return fmt.Sprintf("$%.2f", float64(m))
}

func (b *book) print() {
	fmt.Printf("%-15s: %s - (%v)\n", b.title, b.price.string(), formatDate(b.publishedAt))
}

func formatDate(data interface{}) string {
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
		mobydick       = book{product: product{title: "moby dick", price: 10}, publishedAt: 1712755836}
		treasureIsland = book{product: product{title: "treasure island", price: 25}, publishedAt: "1711891860"}
		oliverTwist    = book{product: product{title: "oliver twist", price: 15}, publishedAt: nil}
		minecraft      = game{product{title: "minecraft", price: 20}}
		tetris         = game{product{title: "tetris", price: 5}}
		rubik          = puzzle{product{title: "rubik's cube", price: 5}}
	)

	var store list
	store = append(store, &minecraft, &tetris, &mobydick, &rubik, &treasureIsland, &oliverTwist)
	// store.print()
	store.discount(0.5)
	store.print()
}
