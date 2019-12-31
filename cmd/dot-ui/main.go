package main

import (
	"log"

	dotui "github.com/kreativka/dot-ui"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	dotui.RunUI()
}
