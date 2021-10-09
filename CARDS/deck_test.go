package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		//t.Errorf("1expected deck to have 16 elements, got %d", len(d)) //my version
		t.Errorf("expected deck to have 16 elements, got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected deck to have Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("expected deck to have Four of Clubs for thr last card, but got %s", d[len(d)-1])
	}

	func TesSaveToDeckAndNewDeckFromFile(t *testing.T) {
		os.remove("_decktesting")

		deck := newDeck()
		deck.saceToFile("_decktesting")

		loadedDeck := newDeckFromFile("_decktesting")

		if len(loadedDeck) != 16 {
			t.Errorf("expected 16 cards in deck, got %d", len(loadedDeck))
		}

		os.remove("_decktesting")

	}
}
