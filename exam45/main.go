package main

import (
	"coin/exam45/cli"
	"coin/exam45/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
