package cli

import (
	"coin/exam46/explorer"
	"coin/exam46/rest"
	"flag"
	"fmt"
	"os"
	"runtime"
)

func usage() {
	fmt.Printf("Welcome to FDong Coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port: 		Set the PORT of the server\n")
	fmt.Printf("-mode:			Choose between 'html' and 'rest'\n")
	runtime.Goexit()
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
