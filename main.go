package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

//WHITE const for the string "white"
const WHITE = "white"

//BLACK const for the string "black"
const BLACK = "black"

var sideMap = map[string]string{"w": WHITE, WHITE: WHITE, "b": BLACK, BLACK: BLACK}
var side = WHITE

var games []Game

func main() {
	var ok bool
	var s = flag.String("side", "w", "type w(hite) or b(lack)")
	flag.Parse()
	ts := *s
	if side, ok = sideMap[strings.ToLower(ts)]; !ok {
		log.Println("Invalid side: ", ts)
		return
	}
	d := "pgns/" + side
	fileInfo, err := ioutil.ReadDir(d)
	if err != nil {
		log.Println(err)
		return
	}

	for _, file := range fileInfo {
		lines := getFileLines(d, file)
		games = append(games, getGameFromData(lines))
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	gi := r1.Intn(len(games))
	quiz(games[gi])
}

func quiz(g Game) {
	//TODO: Use this link to get FEN's from PGN
	fmt.Println("Opening: ", g.Opening)
	for _, v := range g.Moves {
		fmt.Println("Move:", v.Number)
		if side == WHITE {
			var text string
			fmt.Scanln(&text)
			if text == v.White {
				fmt.Printf("%s\n", v.Black)
			} else {
				fmt.Println("Incorrect.  The move is " + v.White)
				return
			}
		} else {
			var text string
			fmt.Printf("%s\n", v.White)
			fmt.Scanln(&text)
			if text == v.Black {
				fmt.Print("\n")
			} else {
				fmt.Println("Incorrect.  The move is " + v.Black)
				return
			}
		}
	}
	fmt.Println("Congrats!  You completed the quiz for: ", g.ECO, g.Opening)
}

func getGameFromData(lines []string) Game {
	var g Game
	for _, v := range lines {
		if strings.HasPrefix(v, "[ECO") {
			g.ECO = getVal(v)
		}
		if strings.HasPrefix(v, "[Opening") {
			g.Opening = getVal(v)
		}
		if strings.HasPrefix(v, "1.") {
			g.Moves = getMoves(v)
		}
	}
	return g
}

func getMoves(s string) []Move {
	moves := []Move{}
	parts := strings.Split(s, ".")
	for k, v := range parts {
		if len(v) > 1 {
			t := strings.Split(v, " ")
			m := Move{}
			m.Number = k
			m.White = t[1]
			if len(t) > 3 {
				m.Black = t[2]
			}
			moves = append(moves, m)
		}
	}
	return moves
}

func getVal(s string) string {
	parts := strings.Split(s, "\"")
	r := parts[1]
	return r
}

func getFileLines(dir string, f os.FileInfo) []string {
	file, err := os.Open(dir + "/" + f.Name())
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	return txtlines
}
