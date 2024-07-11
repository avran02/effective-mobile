package main

import (
	"log"

	"github.com/avran02/effective-mobile/internal/app"
)

func main() {
	a := app.New()
	if err := a.Run(); err != nil {
		log.Fatal("can't run app", err)
	}
}
