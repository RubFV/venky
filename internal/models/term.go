package models

type Term struct {
	ID         int
	Expression string
	Reading    string
	Meaning    string
	Kanjis     []Kanji
	Categories []Category
	Lists      []List
}
