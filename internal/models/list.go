package models

type List struct {
	ID          int
	Name        string
	Description string
	Kanjis      []Kanji
	Terms       []Term
}
