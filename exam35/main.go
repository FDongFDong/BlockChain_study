package main

import (
	"coin/exam34/cli"
	"coin/exam34/db"
)

func main() {
	defer db.Close()
	cli.Start()

}
