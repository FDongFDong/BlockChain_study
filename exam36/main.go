package main

import (
	"coin/exam36/cli"
	"coin/exam36/db"
)

func main() {
	defer db.Close()
	cli.Start()

}
