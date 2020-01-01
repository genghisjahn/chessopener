package main

//Game struct that holds the game data
type Game struct {
	Opening string
	ECO     string
	Moves   []Move
}

//Move struct that holds info on a specific move and the move order
type Move struct {
	Number int
	White  string
	Black  string
}
