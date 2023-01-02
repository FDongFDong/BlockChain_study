package main

import (
	"coin/exam38/cli"
	"coin/exam38/db"
)

func main() {
	defer db.Close()
	cli.Start()

}
