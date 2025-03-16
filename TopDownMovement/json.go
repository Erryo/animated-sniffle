package main

import (
	"encoding/json"
	"log"
	"os"
)

func (l *level) loadMap() {
	var data []byte
	var err error

	data, err = os.ReadFile(l.pathToMap)
	if err != nil {
		log.Panicln("could not open file:", l.pathToMap, "||", err)
	}

	if err = json.Unmarshal(data, &l.tiles); err != nil {
		log.Panicln("couldn't unmarshal json to l.tile from:", l.pathToMap, "||", err)
	}
}
