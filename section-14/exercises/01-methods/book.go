// package main

// import "fmt"

// type book struct {
// 	title string
// 	price float64
// }

// func printBook(b book) {
// 	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
// }

// func (b *book) print() {
// 	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
// }

// func (b *book) discount(ratio float64) {
// 	b.price = b.price * (1 - ratio)
// }
