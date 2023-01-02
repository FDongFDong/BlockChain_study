package main

import (
	"coin/exam39/cli"
	"coin/exam39/db"
)

func main() {
	defer db.Close()
	cli.Start()

}
