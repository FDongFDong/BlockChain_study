package main

import (
	"coin/exam46/cli"
	"coin/exam46/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
