package main

import (
	"coin/exam32/cli"
	"coin/exam32/db"
)

func main() {
	defer db.Close()
	cli.Start()

}
