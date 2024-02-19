package main

import (
	"fmt"
	"time"
)

func main() {
	var groet = bepaalgroet()
	fmt.Printf("%v! Welkom bij Fonteyn Vakantieparken\n", groet)
}
func bepaalgroet() string {
	hour := time.Now().Hour()
	var groet string
	if hour >= 7 && hour < 12 {
		groet = "goedemorgen"
	} else if hour >= 12 && hour < 18 {
		groet = "goedemiddag"
	} else if hour >= 18 && hour < 23 {
		groet = "goedenavond"
	} else {
		groet = "Sorry, de parkeerplaats is s nachts gesloten"
	}
	return groet
}
