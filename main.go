package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/fatih/color"
	. "github.com/rsteube/tinyelk/config"
	"github.com/rsteube/tinyelk/db"
	"github.com/rsteube/tinyelk/server"
	"github.com/vjeantet/grok"
	"log"
	"os"
	"time"
)

func LogTransform(config *Config, cache *db.Cache) {
	//g, _ := grok.NewWithConfig(&grok.Config{NamedCapturesOnly: true})
	g, _ := grok.New()
	for name, pattern := range config.Grok.Patterns.Base {
		g.AddPattern(name, pattern)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		// TODO precompile patterns
		// TODO iterate over patterns
		// TODO add multiline
		if values, err := g.Parse(config.Grok.Patterns.Line["javalog"], scanner.Text()); err == nil {
			if len(values) == 0 {
				continue
			}
			values["_message"] = scanner.Text()

			if t, err := dateparse.ParseAny(values["_timestamp"]); err != nil {
				// TODO
				log.Println("TODO failed to parse timestamp")
			} else {
				values["_timestamp"] = t.Format(time.RFC3339)
			}

			for _, drop := range config.Grok.Drop {
				delete(values, drop)
			}

			if j, err := json.Marshal(values); err == nil {

				// TODO verify timestamp is set
				cache.Put(values["_timestamp"], j)

				fmt.Println(string(j))
			}
		} else {
			log.Println(err)
		}
	}

}

func main() {
	log.SetPrefix(color.CyanString("tinyelk: "))
	log.SetOutput(os.Stderr)

	// TODO error handling
	config, err := Load()
	if err != nil {
		fmt.Println("error parsing")
		return
	}

	cache, _ := db.Open()
	defer cache.Close()

	//go LogTransform(config)
	go LogTransform(&config, cache)
	server.Serve(&config, cache)
}
