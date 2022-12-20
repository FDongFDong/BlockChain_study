package main

import (
	"coin/exam09/blockchain"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	// public/private는 template까지 영향을 준다.
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}
func main() {
	// tamplates를 업데이트한다.(templates는 pages/*.gohtml을 가지고 있게되고)
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	// 해당 라인이 실행되면 templates variable은 template Object가 된다.
	// (templates는 partials/*.gohtml도 함께 가지고 있게 된다.)
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	// 에러가 있을때만 실행
	log.Fatal(http.ListenAndServe(port, nil))
}
