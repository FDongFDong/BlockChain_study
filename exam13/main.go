package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

// struct field tag 사용
// omitempty : Field가 비어있으면 Field를 숨겨준다.
// "-" : 해당 필드를 무시한다.
type URLDescription struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
	Field       int    `json:"-"`
}

// Client에게 JSON을 보낸다.
func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
			Field:       1,
		},
		{
			URL:         "/blocks",
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
	}
	// Client에게 JSON임을 알려주기 위함
	rw.Header().Add("Content-Type", "application/json")
	// // data를 JSON형태로 인코딩한다.
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)

	// 위와 같은 동작을 한다.
	json.NewEncoder(rw).Encode(data)

}
func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
