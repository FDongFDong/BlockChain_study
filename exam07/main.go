package main

import (
	"coin/exam07/blockchain"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const port string = ":4000"

type homeData struct {
	// public/private는 template까지 영향을 준다.
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	// template.Must() err가 있다면 처리해준다. 에러가 없으면 Template pointer를 반환한다.
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	// 블록체인을 블러와 저장한다.
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}
func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	// 에러가 있을때만 실행
	log.Fatal(http.ListenAndServe(port, nil))
}
