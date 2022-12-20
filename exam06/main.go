package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

// request가 포인터인 이유(복사할 이유가 없기 떄문이다)
// file이 될수도 있고 빅데이터가 될 수도 있기 떄문이다.
func home(rw http.ResponseWriter, r *http.Request) {
	// Fprint()는 io.Writer을 첫번째 인자로 받아 Writer에게 출력한다.
	fmt.Fprint(rw, "Hello from home!")
}
func main() {
	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)
	// 에러가 있을때만 실행
	log.Fatal(http.ListenAndServe(port, nil))
}
