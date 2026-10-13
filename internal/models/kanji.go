package models

type Kanji struct {
	ID        int
	Character string
	Meaning   string
	OnYomi    string
	KunYomi   string
	Level     string
	Terms     []Term
	Lists     []List
}
