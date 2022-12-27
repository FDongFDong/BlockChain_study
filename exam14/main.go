package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URL string

// MarshalText()는 인터페이스이다.
func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func (u URLDescription) String() string {
	return "Hello I'm the URL Description"
}

// Client에게 JSON을 보낸다.
func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
	}
	// [사용법 2]
	// fmt.Println(data)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)

}
func main() {
	// [사용법 1]
	// fmt 패키지는 String()가 구현되어 있으면 호출해준다.
	// fmt.Println(URLDescription{
	// 	URL:         "/",
	// 	Method:      "GET",
	// 	Description: "See Documentation",
	// })
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
