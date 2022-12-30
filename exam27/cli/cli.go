package cli

import (
	"coin/exam27/explorer"
	"coin/exam27/rest"
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welcome to FDong Coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port: 		Set the PORT of the server\n")
	fmt.Printf("-mode:			Choose between 'html' and 'rest'\n")
	os.Exit(0)
}
func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	flag.Parse()

	switch *mode {
	case "rest":
		// start rest api
		rest.Start(*port)
	case "html":
		// start html explorer
		explorer.Start(*port)
	default:
		usage()
	}
}
