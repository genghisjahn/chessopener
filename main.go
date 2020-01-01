package main

import (
	"flag"
	"log"
	"strings"
)

//WHITE const for the string "white"
const WHITE = "white"

//BLACK const for the string "black"
const BLACK = "black"

var sideMap = map[string]string{"w": WHITE, WHITE: WHITE, "b": BLACK, BLACK: BLACK}
var side = WHITE

func main() {
	var ok bool
	var s = flag.String("side", "w", "type w(hite) or b(lack)")
	flag.Parse()
	ts := *s
	if side, ok = sideMap[strings.ToLower(ts)]; !ok {
		log.Println("Invalid side: ", ts)
		return
	}
	log.Println("The side is: ", side)

}
