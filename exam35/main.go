package main

import (
	"coin/exam35/cli"
	"coin/exam35/db"
)

func main() {
	defer db.Close()
	cli.Start()

}
