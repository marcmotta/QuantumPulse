// cmd/quantumpulse/main.go
package main

import (
	"flag"
	"log"
	"os"

	"quantumpulse/internal/quantumpulse"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := quantumpulse.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
