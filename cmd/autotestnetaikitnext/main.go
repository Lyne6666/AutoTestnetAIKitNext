// cmd/autotestnetaikitnext/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"autotestnetaikitnext/internal/autotestnetaikitnext"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := autotestnetaikitnext.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
