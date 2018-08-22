package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expect 16 cards, but got %v", len(d))
	}
}

func TestSavetoFileAndnewDeckFromFile(t *testing.T) {
	os.Remove("_dockTesing")

	d := newDeck()
	d.savetoFile("_dockTesing")

	loadedDeck := newDeckFromFile("_dockTesing")

	if len(loadedDeck) != 16 {
		t.Errorf("Expect 16 cards, but got %v", len(loadedDeck))
	}

	os.Remove("_dockTesing")
}
