package main

import "net/http"

const port string = ":4000"

func main() {
	http.ListenAndServe(port, nil)
}
