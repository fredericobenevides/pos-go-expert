package main

import (
	"fmt"

	"github.com/fredericobenevides/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
