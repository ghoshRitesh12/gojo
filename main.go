package main

import (
	"log"

	"github.com/ghoshRitesh12/gojo/src"
)

func main() {
	const PORT string = ":4011"

	app := src.Init()
	log.Fatalln(app.Run(PORT))
}
