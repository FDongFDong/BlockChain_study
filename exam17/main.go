package main

import (
	"coin/exam17/explorer"
	"coin/exam17/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
