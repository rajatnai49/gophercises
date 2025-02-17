package deck

import (
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d := CreateDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52 cards but got %d\n", len(d))
	}

	if d[0].Rank != Ace || d[0].Suit != Spades {
		t.Errorf("Expected first card to be Ace of Spades but got %s of %s\n", ranks[d[0].Rank], suits[d[0].Suit])
	}

	if d[len(d)-1].Rank != King || d[len(d)-1].Suit != Clubs {
		t.Errorf("Expected last card to be King of Clubs but got %s of %s\n", ranks[d[len(d)-1].Rank], suits[d[len(d)-1].Suit])
	}
}

func TestShuffleDeck(t *testing.T) {
	d := CreateDeck()
	originalDeck := make(Deck, len(d))
	copy(originalDeck, d)

	shuffled := false
	for i := 0; i < len(d); i++ {
		if d[i] != originalDeck[i] {
			shuffled = true
			break
		}
	}

	if !shuffled {
		t.Errorf("Expected deck to be shuffled, atleast two cards should have different positions.\n")
	}
}
