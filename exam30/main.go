package main

import (
	"coin/exam30/cli"
	"coin/exam30/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
