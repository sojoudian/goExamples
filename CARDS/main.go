package main

// func main() {
// 	//var card string = "Ace of Spades"
// 	//card := "Ace of Spades"
// 	// card := newCard()
// 	// fmt.Println(card)

// 	// cards := deck{"Ace of Diamonds", newCard()}
// 	// cards = append(cards, "Six of Spades")
// 	//fmt.Println(cards)

// 	// for i, card := range cards {
// 	// 	fmt.Println(i, card)
// 	// }
// 	cards := newDeck()
// 	//cards.print()

// 	hand, remainingCards := deal(cards, 5)
// 	hand.print()
// 	remainingCards.print()
// }

// // func newCard() string {
// // 	return "Five of Diamonds"
// // }

// func main() {
// 	greeting := "Hi there!"
// 	fmt.Println([]byte(greeting))
// }

//////////////////////////  27
// func main() {
// 	cards := newDeck()
// 	// fmt.Println(cards.toString())
// 	cards.saveToFile("my_cards")

// }
////////////////////////// 28
// func main() {
// 	cards := newDeckFromFile("my_cards")
// 	cards.print()
// }

////////////////////////// 29
func main() {
	cards := newDeck()
	cards.shuffle() // shuffle
	cards.print()   // print
}
