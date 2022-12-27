package main

import (
	"coin/exam12/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	URL         string
	Method      string
	Description string
}

// Client에게 JSON을 보낸다.
func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	// data를 JSON형태로 인코딩한다.
	b, err := json.Marshal(data)
	utils.HandleErr(err)
	fmt.Printf("%s", b)

}
func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
