package main

import (
	"coin/exam44/cli"
	"coin/exam44/db"
)

func main() {
	defer db.Close()
	cli.Start()

}
