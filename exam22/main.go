package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welcome to FDong Coin\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer: 		Start the HTML Explorer\n")
	fmt.Printf("rest:			Start the REST API (recommende)\n")
	os.Exit(0)
}

func main() {
	// FlagSet은 go에게 어떤 command가 어떤 flag를 가질 것인지 알려주는 역할을 한다.

	if len(os.Args) < 2 {
		usage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)
	portFlag := rest.Int("port", 4000, "Sets the port of the server")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		rest.Parse(os.Args[2:])
	default:
		usage()
	}
	if rest.Parsed() {
		fmt.Println(*portFlag)
		fmt.Println("Start Server")
	}

}
