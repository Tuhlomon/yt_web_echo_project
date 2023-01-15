package main

import (
	"log"

	"github.com/Tuhlomon/yt_web_echo_project/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
