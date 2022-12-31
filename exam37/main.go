package main

import (
	"coin/exam37/cli"
	"coin/exam37/db"
)

func main() {
	defer db.Close()
	cli.Start()

}
