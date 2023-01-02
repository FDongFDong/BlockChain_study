# BlockChain_study
Go 언어로 블록체인 스터디

- [BlockChain\_study](#blockchain_study)
  - [Genesis Block 만들어보기](#genesis-block-만들어보기)
  - [Genesis Block, Second Block, ... 만들어보기](#genesis-block-second-block--만들어보기)
  - [Refactoring, Singleton](#refactoring-singleton)
    - [block 추가 및 block 정보 가져오기](#block-추가-및-block-정보-가져오기)
    - [Web Server 1](#web-server-1)
  - [Blockchain ↔ 백엔드 서버를 만들기 위한 build up 01](#blockchain-백엔드-서버를-만들기-위한-build-up-01)
  - [Blockchain ↔ 백엔드 서버를 만들기 위한 build up 02](#blockchain-백엔드-서버를-만들기-위한-build-up-02)
    - [Template를 이용한 웹페이지 출력](#template를-이용한-웹페이지-출력)
    - [Blockchain.Data를 웹페이지에 출력](#blockchaindata를-웹페이지에-출력)
    - [앞서 만든 서버를 Refactoring 하기](#앞서-만든-서버를-refactoring-하기)
    - [Block 데이터 출력하기](#block-데이터-출력하기)
    - [앞서 만든 서버 Refactoring 하기](#앞서-만든-서버-refactoring-하기)
    - [Struct 다루기 1](#struct-다루기-1)
    - [Struct 다루기 2](#struct-다루기-2)
    - [Struct 다루기](#struct-다루기)
    - [RESTful하게 수정하기](#restful하게-수정하기)
    - [상기 코드 Refactoring 진행](#상기-코드-refactoring-진행)
    - [Gorilla/mux 패키지 사용하기](#gorillamux-패키지-사용하기)
    - [블록의 Height 값을 받아 해당하는 블록정보 가져오기](#블록의-height-값을-받아-해당하는-블록정보-가져오기)
    - [에러 처리하기](#에러-처리하기)
    - [MiddleWare 적용](#middleware-적용)
- [os.Args 사용](#osargs-사용)
  - [FlagSet 사용](#flagset-사용)
  - [FlagSet 응용](#flagset-응용)
- [DB 처리하기](#db-처리하기)
  - [bolt.db 사용하기 1](#boltdb-사용하기-1)
  - [bolt.db 사용하기 2](#boltdb-사용하기-2)
  - [bolt.db 사용하기 3](#boltdb-사용하기-3)
  - [bold db 확인하는 패키지](#bold-db-확인하는-패키지)
  - [boltbrowser 사용](#boltbrowser-사용)
  - [boltdbweb 사용](#boltdbweb-사용)
  - [DB로부터 저장된 블록 데이터 불러와 콘솔로 출력](#db로부터-저장된-블록-데이터-불러와-콘솔로-출력)
  - [DB에 저장된 블록 GET 메서드로 가져오기](#db에-저장된-블록-get-메서드로-가져오기)
  - [RESTful 동작시키기](#restful-동작시키기)
- [작업 증명 구현하기(PoW)](#작업-증명-구현하기pow)
  - [구현](#구현)
  - [상기 코드를 추가하여 기존 코드에 채굴기능 만들기](#상기-코드를-추가하여-기존-코드에-채굴기능-만들기)
  - [Defficulty 자동으로 수정되게 변경하기](#defficulty-자동으로-수정되게-변경하기)
    - [8분~ 12분 사이 5개 블록이 생성되는 것을 기준으로 코드 작성](#8분-12분-사이-5개-블록이-생성되는-것을-기준으로-코드-작성)
    - [시나리오 진행](#시나리오-진행)
- [트랜잭션 구축](#트랜잭션-구축)
  - [코인베이스에서 채굴자에게 코인을 주도록 만들고 트랜잭션에 기록하기](#코인베이스에서-채굴자에게-코인을-주도록-만들고-트랜잭션에-기록하기)
  - [보유한 자산 조회하기](#보유한-자산-조회하기)
- [Mempoll(메모리풀)](#mempoll메모리풀)
  - [Mempool에 트랜잭션 생성하기](#mempool에-트랜잭션-생성하기)


## Genesis Block 만들어보기
[exam01](https://github.com/FDongFDong/BlockChain_study/tree/main/exam01)
 
## Genesis Block, Second Block, ... 만들어보기
[exam02](https://github.com/FDongFDong/BlockChain_study/tree/main/exam02)

## Refactoring, Singleton 
[exam03]()
[exam04]()

- 기존 함수에 문제가 있다.
  - 하나의 함수에서 블록을 생성하고
  - 블록을 해시하고
  - 새로운 블록을 추가한다.
  
```go
  func (b *blockchain) addBlock(data string) {
    newBlock := block{data, "", b.getLastHash()}
    hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
    newBlock.hash = fmt.Sprintf("%x", hash)
    b.blocks = append(b.blocks, newBlock)
  }
```

### block 추가 및 block 정보 가져오기

[exam05]()

기존에 private하게 만들어둔 block 정보를 public하게 만들어 준 후 AllBlocks() 함수를 통해 데이터를 읽어왔다. 


### Web Server 1
[exam06]()

블록체인 <-> 백엔드 서버를 위한 빌드업 

간단한 웹서버 만들기 


## Blockchain ↔ 백엔드 서버를 만들기 위한 build up 01

- 소스 코드
  - main.go

    ```go
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
    ```

## Blockchain ↔ 백엔드 서버를 만들기 위한 build up 02

### Template를 이용한 웹페이지 출력

- 소스 코드
  - main.go

      ```go
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
    ```

  - blockchain/blockchain.go

    ```go
        package blockchain
        
        import (
         "crypto/sha256"
         "fmt"
         "sync"
        )
        
        // main.go에서 가져다 쓰기위해 임시로 대문자로 변경
        type Block struct {
         Data     string
         Hash     string
         PrevHash string
        }
        type blockchain struct {
         blocks []*Block
        }
        
        var b *blockchain
        
        var once sync.Once
        
        func (b *Block) calculateHash() {
         hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
         b.Hash = fmt.Sprintf("%x", hash)
        }
        
        func getLastHash() string {
         totalBlocks := len(GetBlockchain().blocks)
         if totalBlocks == 0 {
          return ""
         }
         return GetBlockchain().blocks[totalBlocks-1].Hash
        }
        
        func createBlock(data string) *Block {
         newBlock := Block{data, "", getLastHash()}
         newBlock.calculateHash()
         return &newBlock
        }
        
        // export 함수
        func (b *blockchain) AddBlock(data string) {
         b.blocks = append(b.blocks, createBlock(data))
        }
        
        func GetBlockchain() *blockchain {
         if b == nil {
          once.Do(func() {
           b = &blockchain{}
           b.AddBlock("Genesis Block")
          })
         }
         return b
        }
        
        // 사용자에게 field를 드러내주는 function(singleton의 철학)
        func (b *blockchain) AllBlocks() []*Block {
         return b.blocks
         // return GetBlockchain().blocks
        }
    ```


### [Blockchain.Data](http://Blockchain.Data)를 웹페이지에 출력

- 기존 .html 파일을 .gohtml 파일로 수정
- css는 해당 mvp.css 사용

    <link rel="stylesheet" href="[https://unpkg.com/mvp.css@1.12/mvp.css](https://unpkg.com/mvp.css@1.12/mvp.css)">

- 소스 코드
  - main.go

      ```go
        package main
        
        import (
         "coin/exam08/blockchain"
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
         tmpl := template.Must(template.ParseFiles("templates/home.gohtml"))
         data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
         tmpl.Execute(rw, data)
        }
        func main() {
         http.HandleFunc("/", home)
         fmt.Printf("Listening on http://localhost%s\n", port)
         // 에러가 있을때만 실행
         log.Fatal(http.ListenAndServe(port, nil))
        }
      ```

  - blockchain/blockchain.go

      ```go
        package blockchain
        
        import (
         "crypto/sha256"
         "fmt"
         "sync"
        )
        
        // main.go에서 가져다 쓰기위해 임시로 대문자로 변경
        type Block struct {
         Data     string
         Hash     string
         PrevHash string
        }
        type blockchain struct {
         blocks []*Block
        }
        
        var b *blockchain
        
        var once sync.Once
        
        func (b *Block) calculateHash() {
         hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
         b.Hash = fmt.Sprintf("%x", hash)
        }
        
        func getLastHash() string {
         totalBlocks := len(GetBlockchain().blocks)
         if totalBlocks == 0 {
          return ""
         }
         return GetBlockchain().blocks[totalBlocks-1].Hash
        }
        
        func createBlock(data string) *Block {
         newBlock := Block{data, "", getLastHash()}
         newBlock.calculateHash()
         return &newBlock
        }
        
        // export 함수
        func (b *blockchain) AddBlock(data string) {
         b.blocks = append(b.blocks, createBlock(data))
        }
        
        func GetBlockchain() *blockchain {
         if b == nil {
          once.Do(func() {
           b = &blockchain{}
           b.AddBlock("Genesis Block")
          })
         }
         return b
        }
        
        // 사용자에게 field를 드러내주는 function(singleton의 철학)
        func (b *blockchain) AllBlocks() []*Block {
         return b.blocks
         // return GetBlockchain().blocks
        }
      ```

  - pages/add.gohtml
    - 다음에 진행
  - pages/home.gohtml

      ```html
        <!DOCTYPE html>
        <html lang="en">
          <head>
            <meta charset="UTF-8" />
            <meta http-equiv="X-UA-Compatible" content="IE=edge" />
            <meta name="viewport" content="width=device-width, initial-scale=1.0" />
            <link rel="stylesheet" href="https://unpkg.com/mvp.css@1.12/mvp.css"> 
            <title>Coin</title>
          </head>
          <body>
          <header>
            <h1>{{.PageTitle}}</h1>
          </header>
          <main>
            {{range .Blocks}} 
            {{/* .Data는 실제 Blocks 안에 있는 Data를 의미한다. */}}
            {{/* 가져올 값들은 모두 대문자로 시작해야하고 struct에 있는 field명과 같아야한다. */}}
            <section>
              <ul>
                <li>{{.Data}}</li>
                <li>{{.Hash}}</li>
                <li>{{.PrevHash}}</li>
              </ul>
            <section>
            {{end}}
          </main>
          </body>
        </html>
      ```

  - partials/footer.html
    - 다음에 진행
  - partials/head.gohtml
    - 다음에 진행
  - partials/header.gohtml
    - 다음에 진행



### 앞서 만든 서버를 Refactoring 하기

- html/template 패키지를 이용해 .gohtml 파일 경로 설정 및 라우팅
- 소스 코드
  - main.go

      ```go
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
      ```

  - template/pages/add.gohtml

      ```html
        {{define "add"}}
        <!DOCTYPE html>
        <html lang="en">
        {{template "head"}}
          <body>
          {{template "header"}}
            <main>
              <form>
                <input type="text" placeholder="Data for your block" required />
              </form>
            </main> 
            {{template "footer"}}
          </body>
        </html>
        {{end}}
      ```

  - template/pages/home.gohtml

      ```html
        {{define "home"}}
        <!DOCTYPE html>
        <html lang="en">
        {{template "head"}}
          <body>
          {{template "header"}}
            <main>
              {{range .Blocks}} 
        {{template "Block"}}
              {{end}}
            </main>
            {{template "footer"}}
          </body>
        </html>
        {{end}}
      ```

  - template/partials/block.gohtml

      ```html
        {{define "Block"}}
              <div>
                <ul>
                  <li>{{.Data}}</li>
                  <li>{{.Hash}}</li>
                  {{if .PreHash}}
                    <li>{{.PreHash}}</li>
                  {{end}} 
                </ul>
              </div>
            <hr />
        {{end}}
      ```

  - template/partials/footer.gohtml

      ```html
        {{define "footer"}}
        <footer>&copy; 2022</footer>
        {{end}}
      ```

  - template/partials/head.gohtml

      ```html
        {{define "head"}}
          <head>
            <meta charset="UTF-8" />
            <meta http-equiv="X-UA-Compatible" content="IE=edge" />
            <meta name="viewport" content="width=device-width, initial-scale=1.0" />
            <link rel="stylesheet" href="https://unpkg.com/mvp.css@1.12/mvp.css"> 
            <title>Coin</title>
          </head>
          {{end}}
      ```

  - template/partials/header.gohtml

      ```html
        {{define "header"}}
        
        <header>
            <nav>
              <a href="/"><h1>FDongFDong 코인</h1></a>
              <ul>
                <li>
                  <a href="/">Home</a>
                </li>
                <li>
                  <a href="/add">Add</a>
                </li>
              </ul>
            </nav>
            <h1>{{.PageTitle}}</h1>
        </header>
        {{end}}
      ```



### Block 데이터 출력하기

- /add 페이지 get 호출 시 보여줄 화면
- /add 페이지에서 블록 데이터 추가 시 /home으로 리다이렉션 진행
- 추가된 블록 데이터 home template에 렌더링 시켜주기
- 소스 코드
  - main.go

      ```go
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
         PageTitle string
         Blocks    []*blockchain.Block
        }
        
        func home(rw http.ResponseWriter, r *http.Request) {
         data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
         templates.ExecuteTemplate(rw, "home", data)
        }
        func add(rw http.ResponseWriter, r *http.Request) {
         switch r.Method {
         case "GET":
          templates.ExecuteTemplate(rw, "add", nil)
         case "POST":
          r.ParseForm()
          data := r.FormValue("blockData")
          fmt.Println(data)
          blockchain.GetBlockchain().AddBlock(data)
          http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
        
         }
        }
        func main() {
         templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
         templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
         http.HandleFunc("/", home)
         http.HandleFunc("/add", add)
         fmt.Printf("Listening on http://localhost%s\n", port)
         log.Fatal(http.ListenAndServe(port, nil))
        }
      ```

  - add.gohtml

      ```html
        {{define "add"}}
        <!DOCTYPE html>
        <html lang="en">
          {{template "head" "Add"}}
          <body>
            {{template "header" "Add"}}
            <main>
                <form method="POST">
                    <input type="text" placeholder="Data for your block" required name="blockData" />
                    <button>Add Block</button>
                </form>
            </main>
          {{template "footer"}}
          </body>
        </html>
        {{end}}
      ```

  - home.gohtml

      ```html
        {{define "home"}}
        <!DOCTYPE html>
        <html lang="en">
          {{template "head" .PageTitle}}
          <body>
            {{template "header" .PageTitle}}
            <main>
            {{range .Blocks}}
              {{template "block" .}}
            {{end}}
            </main>
          {{template "footer"}}
          </body>
        </html>
        {{end}}
      ```


### 앞서 만든 서버 Refactoring 하기

- 소스 코드
  - main.go

      ```go
        package main
        
        import (
         "coin/exam11/explorer"
        )
        
        func main() {
         explorer.Start()
        }
      ```

  - explorer.go

      ```go
        package explorer
        
        import (
         "coin/exam11/blockchain"
         "fmt"
         "log"
         "net/http"
         "text/template"
        )
        
        const (
         port        string = ":4000"
         templateDir string = "explorer/templates/"
        )
        
        var templates *template.Template
        
        type homeData struct {
         PageTitle string
         Blocks    []*blockchain.Block
        }
        
        func home(rw http.ResponseWriter, r *http.Request) {
         data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
         templates.ExecuteTemplate(rw, "home", data)
        }
        func add(rw http.ResponseWriter, r *http.Request) {
         switch r.Method {
         case "GET":
          templates.ExecuteTemplate(rw, "add", nil)
         case "POST":
          r.ParseForm()
          data := r.FormValue("blockData")
          fmt.Println(data)
          blockchain.GetBlockchain().AddBlock(data)
          http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
        
         }
        }
        func Start() {
         // Standard 라이브러리를 사용하고
         templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
         // templates 변수를 사용했다. (템플릿 위에 템플릿을 얹은 형태)
         templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
         http.HandleFunc("/", home)
         http.HandleFunc("/add", add)
         fmt.Printf("Listening on http://localhost%s\n", port)
         log.Fatal(http.ListenAndServe(port, nil))
        }
      ```

  - block.gohtml

      ```html
        {{define "block"}}
         <div>
            <ul>
                <li><strong>Data: </strong>{{.Data}}</li>
                <li><strong>Hash: </strong>{{.Hash}}</li>
                {{if .PrevHash}}
                <li><strong>Previous Hash: </strong>{{.PrevHash}}</li>
                {{end}}
            </ul>
            </div>
            <hr />
        {{end}}
      ```

---

### Struct 다루기 1

- JSON으로 변경하기(Marshal)
- 소스 코드
  - main.go

      ```go
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
      ```

### Struct 다루기 2

- JSON으로 변경하기
- Struct field tag 사용하기
- 소스 코드
  - main.go

      ```go
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
      ```


---

### Struct 다루기

- Interface 이용하여 출력하기
  - MarshalText() 함수 사용하기
    - JSON의 형태를 사용자가 Custom할 수 있다.
- 소스 코드
  - main.go

      ```go
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
         //  URL:         "/",
         //  Method:      "GET",
         //  Description: "See Documentation",
         // })
         http.HandleFunc("/", documentation)
         fmt.Printf("Listening on http://localhost%s\n", port)
        
         log.Fatal(http.ListenAndServe(port, nil))
        }
      ```


---

### RESTful하게 수정하기

- POST
  - 블록 생성 및 추가하기
- GET
  - 모든 블록 가져오기
- 소스 코드
  - main.go

  ```go
    package main
    
    import (
     "coin/exam16/blockchain"
     "coin/exam16/utils"
     "encoding/json"
     "fmt"
     "log"
     "net/http"
    )
    
    const port string = ":4000"
    
    type URL string
    
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
    
    type AddBlockBody struct {
     Message string `json:"message"`
    }
    
    func (u URLDescription) String() string {
     return "Hello I'm the URL Description"
    }
    
    func documentation(rw http.ResponseWriter, r *http.Request) {
     data := []URLDescription{
      {
       URL:         URL("/"),
       Method:      "GET",
       Description: "See Documentation",
      },
      {
       URL:         URL("/blocks"),
       Method:      "GET",
       Description: "See All Block",
      },
      {
       URL:         URL("/blocks"),
       Method:      "POST",
       Description: "Add A Block",
       Payload:     "data:string",
      },
      {
       URL:         URL("/blocks/{id}"),
       Method:      "GET",
       Description: "See A Block",
      },
     }
    
     rw.Header().Add("Content-Type", "application/json")
     json.NewEncoder(rw).Encode(data)
    
    }
    func blocks(rw http.ResponseWriter, r *http.Request) {
     switch r.Method {
     case "GET":
      rw.Header().Add("Content-Type", "application/json")
      json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
     case "POST":
      var addBlockBody AddBlockBody
    
      utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
      blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
      rw.WriteHeader(http.StatusCreated)
     }
    }
    func main() {
     http.HandleFunc("/", documentation)
     http.HandleFunc("/blocks", blocks)
     fmt.Printf("Listening on http://localhost%s\n", port)
     log.Fatal(http.ListenAndServe(port, nil))
    }
  ```

- 실행 결과
  - 모든 블록 가져오기
  - GET
  - URL : /blocks

  ```json
    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Wed, 28 Dec 2022 04:34:24 GMT
    Content-Length: 115
    Connection: close
    
    [
      {
        "Data": "Genesis Block",
        "Hash": "89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3",
        "PrevHash": ""
      }
    ]
  ```

  - 블록 생성 및 추가하기
  - POST
  - URL : /blocks

      ```json
        HTTP/1.1 201 Created
        Date: Wed, 28 Dec 2022 04:34:51 GMT
        Content-Length: 0
        Connection: close
      ```

  - 추가된 블록 확인하기
  - GET
  - URL : /Blocks

      ```json
        HTTP/1.1 200 OK
        Content-Type: application/json
        Date: Wed, 28 Dec 2022 04:35:43 GMT
        Content-Length: 477
        Connection: close
        
        [
          {
            "Data": "Genesis Block",
            "Hash": "89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3",
            "PrevHash": ""
          },
          {
            "Data": "Data for my block",
            "Hash": "6de940f3a7ead5008e358bdda0ac9b0234a4e8dbc94c31ca1dd91b8798607182",
            "PrevHash": "89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3"
          },
        ]
      ```

---

### 상기 코드 Refactoring 진행

- NewServeMux로 새로운 Mux를 만들어 explorer과 rest를 나눠줬다
- 소스 코드
  - main.go

      ```go
        package main
        
        import (
         "coin/exam17/explorer"
         "coin/exam17/rest"
        )
        
        func main() {
         go explorer.Start(3000)
         rest.Start(4000)
        }
      ```

  - rest/rest.go

      ```go
        package rest
        
        import (
         "coin/exam17/blockchain"
         "coin/exam17/utils"
         "encoding/json"
         "fmt"
         "log"
         "net/http"
        )
        
        var port string
        
        type url string
        
        func (u url) MarshalText() ([]byte, error) {
         url := fmt.Sprintf("http://localhost%s%s", port, u)
         return []byte(url), nil
        }
        
        type urlDescription struct {
         URL         url    `json:"url"`
         Method      string `json:"method"`
         Description string `json:"description"`
         Payload     string `json:"payload,omitempty"`
        }
        
        type addBlockBody struct {
         Message string `json:"message"`
        }
        
        func (u urlDescription) String() string {
         return "Hello I'm the URL Description"
        }
        
        func documentation(rw http.ResponseWriter, r *http.Request) {
         data := []urlDescription{
          {
           URL:         url("/"),
           Method:      "GET",
           Description: "See Documentation",
          },
          {
           URL:         url("/blocks"),
           Method:      "GET",
           Description: "See All Block",
          },
          {
           URL:         url("/blocks"),
           Method:      "POST",
           Description: "Add A Block",
           Payload:     "data:string",
          },
          {
           URL:         url("/blocks/{id}"),
           Method:      "GET",
           Description: "See A Block",
          },
         }
        
         rw.Header().Add("Content-Type", "application/json")
         json.NewEncoder(rw).Encode(data)
        
        }
        func blocks(rw http.ResponseWriter, r *http.Request) {
         switch r.Method {
         case "GET":
          rw.Header().Add("Content-Type", "application/json")
          json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
         case "POST":
          var addBlockBody addBlockBody
        
          utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
          blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
          rw.WriteHeader(http.StatusCreated)
         }
        }
        func Start(aPort int) {
         handler := http.NewServeMux()
         port = fmt.Sprintf(":%d", aPort)
         handler.HandleFunc("/", documentation)
         handler.HandleFunc("/blocks", blocks)
         fmt.Printf("Listening on http://localhost%s\n", port)
         log.Fatal(http.ListenAndServe(port, handler))
        }
      ```

  - erplorer/explorer.go

      ```go
        package explorer
        
        import (
         "coin/exam17/blockchain"
         "fmt"
         "log"
         "net/http"
         "text/template"
        )
        
        const (
         templateDir string = "explorer/templates/"
        )
        
        var templates *template.Template
        
        type homeData struct {
         PageTitle string
         Blocks    []*blockchain.Block
        }
        
        func home(rw http.ResponseWriter, r *http.Request) {
         data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
         templates.ExecuteTemplate(rw, "home", data)
        }
        func add(rw http.ResponseWriter, r *http.Request) {
         switch r.Method {
         case "GET":
          templates.ExecuteTemplate(rw, "add", nil)
         case "POST":
          r.ParseForm()
          data := r.FormValue("blockData")
          fmt.Println(data)
          blockchain.GetBlockchain().AddBlock(data)
          http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
        
         }
        }
        func Start(port int) {
         handle := http.NewServeMux()
         templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
         templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
         handle.HandleFunc("/", home)
         handle.HandleFunc("/add", add)
         fmt.Printf("Listening on http://localhost%d\n", port)
         log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handle))
        }
      ```

---

### Gorilla/mux 패키지 사용하기

- mux 변경
- GET메서드 URL에서 파라미터 가져오기
- 소스 코드
  - main.go

      ```go
        package main
        
        import (
         "coin/exam18/rest"
        )
        
        func main() {
        
         rest.Start(4000)
        }
      ```

  - rest.go

      ```go
        package rest
        
        import (
         "coin/exam18/blockchain"
         "coin/exam18/utils"
         "encoding/json"
         "fmt"
         "log"
         "net/http"
        
         "github.com/gorilla/mux"
        )
        
        //
        // ...
        //
        
        func block(rw http.ResponseWriter, r *http.Request) {
         vars := mux.Vars(r)
         fmt.Println(vars)
         id := vars["id"]
        }
        func Start(aPort int) {
         router := mux.NewRouter()
         port = fmt.Sprintf(":%d", aPort)
         router.HandleFunc("/", documentation).Methods("GET")
         router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
         router.HandleFunc("/blocks/{id:[0-9]+}", block).Methods("GET")
         fmt.Printf("Listening on http://localhost%s\n", port)
         log.Fatal(http.ListenAndServe(port, router))
        }
      ```

- 실행 결과

  ```go
    Listening on http://localhost:4000
    map[id:1]
  ```

---

### 블록의 Height 값을 받아 해당하는 블록정보 가져오기

- 소스 코드
  - main.go

      ```go
        package main
        
        import (
         "coin/exam19/rest"
        )
        
        func main() {
        
         rest.Start(4000)
        }
      ```

  - rest/rest.go

      ```go
        package rest
        
        import (
         "coin/exam19/blockchain"
         "coin/exam19/utils"
         "encoding/json"
         "fmt"
         "log"
         "net/http"
         "strconv"
        
         "github.com/gorilla/mux"
        )
        
        // ...
        
        func documentation(rw http.ResponseWriter, r *http.Request) {
         data := []urlDescription{
          // ... 
          // 블록의 height값을 받는다
          {
           URL:         url("/blocks/{height}"),
           Method:      "GET",
           Description: "See A Block",
          },
         }
        }
        
        // URL의 파라미터를 받아서 해당하는 블록을 찾아 json형태로 출력한다.
        func block(rw http.ResponseWriter, r *http.Request) {
         vars := mux.Vars(r)
         id, err := strconv.Atoi(vars["height"])
         utils.HandleErr(err)
        
         block := blockchain.GetBlockchain().GetBlock(id)
         json.NewEncoder(rw).Encode(block)
        
        }
        
        func Start(aPort int) {
         router := mux.NewRouter()
         // ...
         router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")
         fmt.Printf("Listening on http://localhost%s\n", port)
         log.Fatal(http.ListenAndServe(port, router))
        }
      ```

  - blockchain/blockchain.go

      ```go
        package blockchain
        
        // 블록의 데이터에 height 값을 넣어준다.
        
        type Block struct {
         Data     string `json:"data"`
         Hash     string `json:"hash"`
         PrevHash string `json:"prevhash,omitempty"`
         Height   int    `json:"height"`
        }
        
        // ...
        // 블록을 생성할 때 height 값도 함께 넣어준다.
        
        func createBlock(data string) *Block {
         newBlock := Block{data, "", getLastHash(), len(GetBlockchain().blocks) + 1}
         newBlock.calculateHash()
         return &newBlock
        }
        
        // ...
        
        // height 값을 받아 하나의 블록을 반환한다
        func (b *blockchain) GetBlock(height int) *Block {
         return b.blocks[height-1]
        }
      ```

- 실행 결과
    1. 기존에 저장된 블록체인 정보를 가져온다.

        ```json
            [
              {
                "data": "Genesis Block",
                "hash": "89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3",
                "height": 1
              }
            ]
        ```

    2. 하나의 블록을 생성한다.

        ```json
            HTTP/1.1 201 Created
            Date: Wed, 28 Dec 2022 08:17:07 GMT
            Content-Length: 0
            Connection: close
        ```

    3. 생성된 블록을 확인한다.

        ```json
            [
              {
                "data": "Genesis Block",
                "hash": "89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3",
                "height": 1
              },
              {
                "data": "Data for my block",
                "hash": "6de940f3a7ead5008e358bdda0ac9b0234a4e8dbc94c31ca1dd91b8798607182",
                "prevhash": "89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3",
                "height": 2
              }
            ]
        ```

    4. height가 2번인 블록의 정보를 가져온다

        ```json
            {
              "data": "Data for my block",
              "hash": "6de940f3a7ead5008e358bdda0ac9b0234a4e8dbc94c31ca1dd91b8798607182",
              "prevhash": "89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3",
              "height": 2
            }
        ```

---

### 에러 처리하기

- 소스 코드
  - main.go

        ```go
        package main
        
        import (
         "coin/exam20/rest"
        )
        
        func main() {
        
         rest.Start(4000)
        }
        ```

  - rest/rest.go

        ```go
        package rest
        
        // 에러 메세지를 저장하기 위한 구조체 선언
        type errorResponse struct {
         ErrorMessage string `json:"errormessage"`
        }
        
        // ...
        
        func block(rw http.ResponseWriter, r *http.Request) {
         vars := mux.Vars(r)
         id, err := strconv.Atoi(vars["height"])
         utils.HandleErr(err)
        
         block, err := blockchain.GetBlockchain().GetBlock(id)
         encoder := json.NewEncoder(rw)
         if err == blockchain.ErrNotFound {
          encoder.Encode(errorResponse{fmt.Sprint(err)})
         } else {
          encoder.Encode(block)
         }
        
        }
        
        // ...
        ```

  - blockchain/blockchain.go

        ```go
        // 에러 메세지를 만든다.
        var ErrNotFound = errors.New("block not found")
        
        // Client가 잘못된 인덱스에 접근하면 에러 메세지를 리턴한다.
        func (b *blockchain) GetBlock(height int) (*Block, error) {
         if height > len(b.blocks) {
          return nil, ErrNotFound
         }
         return b.blocks[height-1], nil
        }
        ```

- 실행 결과

    ```go
    HTTP/1.1 200 OK
    Date: Wed, 28 Dec 2022 08:41:00 GMT
    Content-Length: 35
    Content-Type: text/plain; charset=utf-8
    Connection: close
    
    {
      "errormessage": "block not found"
    }
    ```

---

### MiddleWare 적용

- Header에 Content-Type, application/json을 넣기 위해 미들웨어형태로 구현
- 소스 코드
  - rest/rest.go

        ```go
        package rest
        
        // ...
        
        func jsonContentTypeMiddleware(next http.Handler) http.Handler {
         return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
          rw.Header().Add("Content-Type", "application/json")
          next.ServeHTTP(rw, r)
         })
        }
        func Start(aPort int) {
         router := mux.NewRouter()
         router.Use(jsonContentTypeMiddleware)
         port = fmt.Sprintf(":%d", aPort)
         router.HandleFunc("/", documentation).Methods("GET")
         //...
        }
        ```

- 실행 결과

    ```go
    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Wed, 28 Dec 2022 15:11:12 GMT
    Content-Length: 366
    Connection: close
    ```

---
___

# os.Args 사용

- os.Args

    ```go
    var Args []string
    ```

    CLI에서 사용된 문자열 배열을 리턴합니다. 첫번째 인자는 실행프로그램의 이름

- 소스 코드

    ```go
    package main
    
    import (
     "fmt"
     "os"
    )
    
    func usage() {
     fmt.Printf("Welcome to FDong Coin\n\n")
     fmt.Printf("Please use the following commands:\n\n")
     fmt.Printf("explorer:   Start the HTML Explorer\n")
     fmt.Printf("rest:   Start the REST API (recommende)\n")
     os.Exit(0)
    }
    
    func main() {
    
     if len(os.Args) < 2 {
      usage()
     }
    
     switch os.Args[1] {
     case "explorer":
      fmt.Println("Start Explorer")
     case "rest":
      fmt.Println("Start REST API")
     default:
      usage()
     }
    }
    ```

- 실행 결과
  - go run main.go 입력 시

    ```go
    > go run main.go     
    Welcome to FDong Coin
    
    Please use the following commands:
    
    explorer:               Start the HTML Explorer
    rest:                   Start the REST API (recommende)
    ```

  - go run main.go explorer 입력 시

    ```go
    > go run main.go explorer
    Start Explorer
    ```

  - go run main.go rest 입력 시

    ```go
    > go run main.go rest    
    Start REST API
    ```

---

## FlagSet 사용

- Flag를 여러개 사용할 때 쓰기 좋다
- 소스 코드

    ```go
    package main
    
    import (
     "flag"
     "fmt"
     "os"
    )
    
    func usage() {
     fmt.Printf("Welcome to FDong Coin\n\n")
     fmt.Printf("Please use the following commands:\n\n")
     fmt.Printf("explorer:   Start the HTML Explorer\n")
     fmt.Printf("rest:   Start the REST API (recommende)\n")
     os.Exit(0)
    }
    
    func main() {
     // FlagSet은 go에게 어떤 command가 어떤 flag를 가질 것인지 알려주는 역할을 한다.
    
     if len(os.Args) < 2 {
      usage()
     }
    
     rest := flag.NewFlagSet("rest", flag.ExitOnError)
     portFlag := rest.Int("port", 4000, "Sets the port of the server")
    
     switch os.Args[1] {
     case "explorer":
      fmt.Println("Start Explorer")
     case "rest":
      rest.Parse(os.Args[2:])
     default:
      usage()
     }
     if rest.Parsed() {
      fmt.Println(*portFlag)
      fmt.Println("Start Server")
     }
    
    }
    ```

---

## FlagSet 응용

- 소스 코드
  - main.go

        ```go
        package main
        
        import "coin/exam23/cli"
        
        func main() {
         cli.Start()
        }
        ```

  - cli/cli.go

      ```go
        package cli
        
        import (
         "coin/exam23/explorer"
         "coin/exam23/rest"
         "flag"
         "fmt"
         "os"
        )
        
        func usage() {
         fmt.Printf("Welcome to FDong Coin\n\n")
         fmt.Printf("Please use the following flags:\n\n")
         fmt.Printf("-port:   Set the PORT of the server\n")
         fmt.Printf("-mode:   Choose between 'html' and 'rest'\n")
         os.Exit(0)
        }
        func Start() {
         if len(os.Args) == 1 {
          usage()
         }
        
         port := flag.Int("port", 4000, "Set port of the server")
         mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
         flag.Parse()
        
         switch *mode {
         case "rest":
          // start rest api
          rest.Start(*port)
         case "html":
          // start html explorer
          explorer.Start(*port)
         default:
          usage()
         }
        }
      ```

- 실행 결과
  - go run main.go -mode=rest -port=2000

    ```shell
      > go run main.go -mode=rest -port=2000
      Listening on http://localhost:2000
    ```

  - go run main.go -mode=html -port=8000

      ```shell
        > go run main.go -mode=html -port=8000
        Listening on http://localhost8000
      ```

  - 잘못된 flag

      ```shell
        > go run main.go -mode=html -port=asdf
        invalid value "asdf" for flag -port: parse error
        Usage of /var/folders/5s/13x9ywys5wz_w321jgl_f_pw0000gn/T/go-build895749235/b001/exe/main:
          -mode string
                Choose between 'html' and 'rest' (default "rest")
          -port int
                Set port of the server (default 4000)
        exit status 2
      ```
---
# DB 처리하기

--- 
## bolt.db 사용하기 1

[https://github.com/boltdb/bolt](https://github.com/boltdb/bolt)

bolt db는 Key/Value 형태의 저장소이다.

- bolt.db 설치하기
    
  ```go
    go get github.com/boltdb/bolt
  ```
    
- 소스 코드
    - db/db.go
        
        ```go
          package db
          
          import (
            "coin/exam24/utils"
          
            "github.com/boltdb/bolt"
          )
          
          const (
            dbname       = "blockchain.db"
            dataBucket   = "data"
            blocksBucket = "blocks"
          )
          
          var db *bolt.DB
          
          // DB initialize, Singleton pattern형식
          func DB() *bolt.DB {
            if db == nil {
              // init db
              // path는 DB의 이름, 파일이 없으면 자동으로 생성해준다,
              dbPointer, err := bolt.Open(dbname, 0600, nil)
              utils.HandleErr(err)
              db = dbPointer
              // bucket이 존재하지 않으면 생성시켜주는 Transaction, 두개의 bucket을 만들어준다.
              // bucket는 table 같은거다
              err = db.Update(func(tx *bolt.Tx) error {
                _, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
                utils.HandleErr(err)
                _, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
          
                return err
              })
              utils.HandleErr(err)
            }
            return db
          }
        ```
        

---

## bolt.db 사용하기 2

- block 데이터 저장하기
- 소스 코드
    - main.go
        
        ```go
          package main
          
          import (
            "coin/exam26/blockchain"
          )
          
          func main() {
            blockchain.Blockchain()
          }
          ```
          
      - db/db.go
          
          ```go
          package db
          
          import (
            "coin/exam26/utils"
            "fmt"
          
            "github.com/boltdb/bolt"
          )
          
          const (
            dbname       = "blockchain.db"
            dataBucket   = "data"
            blocksBucket = "blocks"
          )
          
          var db *bolt.DB
          
          // DB initialize, Singleton pattern형식
          func DB() *bolt.DB {
            if db == nil {
              // init db
              // path는 DB의 이름, 파일이 없으면 자동으로 생성해준다,
              dbPointer, err := bolt.Open(dbname, 0600, nil)
              utils.HandleErr(err)
              db = dbPointer
              // bucket이 존재하지 않으면 생성시켜주는 Transaction, 두개의 bucket을 만들어준다.
              // bucket는 table 같은거다
              err = db.Update(func(tx *bolt.Tx) error {
                _, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
                utils.HandleErr(err)
                _, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
          
                return err
              })
              utils.HandleErr(err)
            }
            return db
          }
          
          func SaveBlock(hash string, data []byte) {
            fmt.Printf("Saving Block %s\nData: %b", hash, data)
            err := DB().Update(func(tx *bolt.Tx) error {
              bucket := tx.Bucket([]byte(blocksBucket))
              err := bucket.Put([]byte(hash), data)
              return err
            })
            utils.HandleErr(err)
          }
          func SaveBlockchain(data []byte) {
            err := DB().Update(func(tx *bolt.Tx) error {
              bucket := tx.Bucket([]byte(dataBucket))
              err := bucket.Put([]byte("checkpoint"), data)
              return err
            })
            utils.HandleErr(err)
          }
        ```
        
    - blockchain/block.go
        
        ```go
        package blockchain
        
        import (
        	"bytes"
        	"coin/exam26/db"
        	"coin/exam26/utils"
        	"crypto/sha256"
        	"encoding/gob"
        	"fmt"
        )
        
        type Block struct {
        	Data     string `json:"data"`
        	Hash     string `json:"hash"`
        	PrevHash string `json:"prevhash,omitempty"`
        	Height   int    `json:"height"`
        }
        
        func (b *Block) toBytes() []byte {
        	var blockBuffer bytes.Buffer
        	encoder := gob.NewEncoder(&blockBuffer)
        	utils.HandleErr(encoder.Encode(b))
        	return blockBuffer.Bytes()
        }
        
        func (b *Block) persist() {
        	db.SaveBlock(b.Hash, b.toBytes())
        }
        func createBlock(data string, prevHash string, height int) *Block {
        	block := &Block{
        		Data:     data,
        		Hash:     "",
        		PrevHash: prevHash,
        		Height:   height,
        	}
        	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
        	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
        	block.persist()
        	return block
        }
        ```
        
    - blockchain/chain.go
        
        ```go
          package blockchain
          
          import (
            "sync"
          )
          
          type blockchain struct {
            NewestHash string `json:"newestHash"`
            Height     int    `json:"height"`
          }
          
          var b *blockchain
          
          var once sync.Once
          
          // AddBlock receiver
          func (b *blockchain) AddBlock(data string) {
            block := createBlock(data, b.NewestHash, b.Height)
            b.NewestHash = block.Hash
            b.Height = block.Height
          }
          
          func Blockchain() *blockchain {
            if b == nil {
              once.Do(func() {
                b = &blockchain{"", 0}
                b.AddBlock("Genesis Block")
              })
            }
            return b
          }
        ```
        
- 실행 결과
    
    ```shell
      > go run main.go
      Saving Block 8500b59bb5271135cd9bcbf0afd693028d76df3b9c7da58d412b13fc8a8f9394
      Data: [111101 11111111 10000001 11 1 1 101 1000010 1101100 1101111 1100011 1101011 1 11111111 10000010 0 1 100 1 100 1000100 1100001 1110100 1100001 1 1100 0 1 100 1001000 1100001 1110011 1101000 1 1100 0 1 1000 1010000 1110010 1100101 1110110 1001000 1100001 1110011 1101000 1 1100 0 1 110 1001000 1100101 1101001 1100111 1101000 1110100 1 100 0 0 0 1010100 11111111 10000010 1 1101 1000111 1100101 1101110 1100101 1110011 1101001 1110011 100000 1000010 1101100 1101111 1100011 1101011 1 1000000 111000 110101 110000 110000 1100010 110101 111001 1100010 1100010 110101 110010 110111 110001 110001 110011 110101 1100011 1100100 111001 1100010 1100011 1100010 1100110 110000 1100001 1100110 1100100 110110 111001 110011 110000 110010 111000 1100100 110111 110110 1100100 1100110 110011 1100010 111001 1100011 110111 1100100 1100001 110101 111000 1100100 110100 110001 110010 1100010 110001 110011 1100110 1100011 111000 1100001 111000 1100110 111001 110011 111001 110100 0]
    ```
    
---
## bolt.db 사용하기 3

- block 데이터 저장하기
- chain 데이터 저장하기
- []byte로 만들어주는 유틸 함수만들기
- 소스 코드
    - blockchain/chain.go
        
        ```go
          package blockchain
          
          import (
            "coin/exam27/db"
            "coin/exam27/utils"
            "sync"
          )
          
          type blockchain struct {
            NewestHash string `json:"newestHash"`
            Height     int    `json:"height"`
          }
          
          var b *blockchain
          
          var once sync.Once
          
          func (b *blockchain) persist() {
            db.SaveBlockchain(utils.ToBytes(b))
          }
          
          // AddBlock receiver
          func (b *blockchain) AddBlock(data string) {
            block := createBlock(data, b.NewestHash, b.Height+1)
            b.NewestHash = block.Hash
            b.Height = block.Height
            b.persist()
          }
          
          func Blockchain() *blockchain {
            if b == nil {
              once.Do(func() {
                b = &blockchain{"", 0}
                b.AddBlock("Genesis Block")
              })
            }
            return b
          }
        ```
        
    - blockchain/block.go
        
        ```go
          package blockchain
          
          import (
            "coin/exam27/db"
            "coin/exam27/utils"
            "crypto/sha256"
            "fmt"
          )
          
          type Block struct {
            Data     string `json:"data"`
            Hash     string `json:"hash"`
            PrevHash string `json:"prevhash,omitempty"`
            Height   int    `json:"height"`
          }
          
          func (b *Block) persist() {
            db.SaveBlock(b.Hash, utils.ToBytes(b))
          }
          func createBlock(data string, prevHash string, height int) *Block {
            block := &Block{
              Data:     data,
              Hash:     "",
              PrevHash: prevHash,
              Height:   height,
            }
            payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
            block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
            block.persist()
            return block
          }
        ```
        
    - uitls/utils.go
        
        ```go
          package utils
          
          import (
            "bytes"
            "encoding/gob"
            "log"
          )
          
          func HandleErr(err error) {
            if err != nil {
              log.Panic(err)
            }
        }
        
        // 원하는 건 모든지 받을 수 있다.
        func ToBytes(i interface{}) []byte {
        	var aBuffer bytes.Buffer
        	encoder := gob.NewEncoder(&aBuffer)
        	HandleErr(encoder.Encode(i))
        	return aBuffer.Bytes()
        }
        ```
        
---

## bold db 확인하는 패키지

- bolt.db로 생성한 db파일이 있어야한다.
---

## boltbrowser 사용

[boltbrowser](https://pkg.go.dev/github.com/br0xen/boltbrowser@v0.0.0-20210531150353-7f10a81cece0#section-readme)

- 설치
    
    ```go
       go get github.com/br0xen/boltbrowser
    ```
    
- 사용 방법
    
    ```go
        boltbrowser <filename>
    ```
    


## boltdbweb 사용

[](https://pkg.go.dev/github.com/evnix/boltdbweb@v0.0.0-20191029203843-5b16e6623bd9)

[GitHub - evnix/boltdbweb: A web based GUI for BoltDB files](https://github.com/evnix/boltdbweb)

- 설치
    
    ```go
    go get github.com/evnix/boltdbweb
    ```
    
- 사용 방법
    
    ```go
    boltdbweb --db-name=<DBfilename>
    ```
    

___

## DB로부터 저장된 블록 데이터 불러와 콘솔로 출력

- 소스 코드
    - main.go
        
        ```go
        func main() {
        	blockchain.Blockchain()
        }
        ```
        
    - blockchain/chain.go
        
        ```go
        func Blockchain() *blockchain {
        	if b == nil {
        		once.Do(func() {
        			// NewestHash가 없고 Height는 0인 블록체인을 만들고
        			b = &blockchain{"", 0}
        			// checkpoint에 data가 있는지 확인한다.
        			fmt.Printf("NewestHash: %s\nHeight: %d\n", b.NewestHash, b.Height)
        			checkpoint := db.Checkpoint()
        			if checkpoint == nil {
        				//없으면 initialize한다.
        				b.AddBlock("Genesis Block")
        			} else {
        				fmt.Printf("Restoring...\n")
        				b.restore(checkpoint)
        			}
        
        		})
        	}
        	fmt.Printf("NewestHash: %s\nHeight: %d\n", b.NewestHash, b.Height)
        	return b
        }
        ```
        
    - db/db.go
        
        ```go
        // checkpoint가 있는지 없는지 리턴하는 함수
        func Checkpoint() []byte {
        	var data []byte
        	// View : Read Only
        	DB().View(func(tx *bolt.Tx) error {
        		// bucket을 가져온다
        		bucket := tx.Bucket([]byte(dataBucket))
        		data = bucket.Get([]byte(checkpoint))
        		return nil
        	})
        	return data
        }
        ```
        
- 실행 결과
    
    ```go
    > go run main.go
    NewestHash: 
    Height: 0
    Restoring...
    NewestHash: 110e8eb40de97b0929e53a3b0cd2d697a4bd25a38fe1a5f069d6d35a40976e17
    Height: 4
    ```
    
___

## DB에 저장된 블록 GET 메서드로 가져오기

- rest api: url에서 hash를 받는다.
- FindBlock() 함수가 DB에 가서 해당 hash를 key로 가진 블록을 찾는다.
    - 못찾으면 nil과 ErrNotFound 에러를 리턴한다.
    - 찾으면 빈 블록을 만들고 그 블록에 restore함수를 호출하여 찾은 블록 데이터를 채운다.
        - restore함수는 db로 부터 불러온 byte slice를 받아서 encoder를 만들고 data를 읽어서 decode해준다.
- 소스 코드
    - main.go
        
        ```go
        func main() {
        	blockchain.Blockchain()
        	cli.Start()
        }
        ```
        
    - rest/rest.go
        
        ```go
        
        // ...
        
        func documentation(rw http.ResponseWriter, r *http.Request) {
        	data := []urlDescription{
        		{
        			URL:         url("/"),
        			Method:      "GET",
        			Description: "See Documentation",
        		},
        		{
        			URL:         url("/blocks"),
        			Method:      "GET",
        			Description: "See All Block",
        		},
        		{
        			URL:         url("/blocks"),
        			Method:      "POST",
        			Description: "Add A Block",
        			Payload:     "data:string",
        		},
        		{
        			URL:         url("/blocks/{hesh}"),
        			Method:      "GET",
        			Description: "See A Block",
        		},
        	}
        
        	rw.Header().Add("Content-Type", "application/json")
        	json.NewEncoder(rw).Encode(data)
        
        }
        
        func block(rw http.ResponseWriter, r *http.Request) {
        	vars := mux.Vars(r)
        	hash := (vars["hash"])
        
        	block, err := blockchain.FindBlock(hash)
        	encoder := json.NewEncoder(rw)
        
        	if err == blockchain.ErrNotFound {
        		encoder.Encode(errorResponse{fmt.Sprint(err)})
        	} else {
        		encoder.Encode(block)
        	}
        
        }
        
        func Start(aPort int) {
        	router := mux.NewRouter()
        	port = fmt.Sprintf(":%d", aPort)
        	// ...
        	// hex 값을 id로 받아오기 위해 a-f
        	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")
        	// ...
        	log.Fatal(http.ListenAndServe(port, router))
        }
        ```
        
- 실행 결과
    - Method : GET
    - URL : [http://localhost:4000/blocks/](http://localhost:4000/blocks/):hash
    - Test URL : [http://localhost:4000/blocks/110e8eb40de97b0929e53a3b0cd2d697a4bd25a38fe1a5f069d6d35a40976e17](http://localhost:4000/blocks/110e8eb40de97b0929e53a3b0cd2d697a4bd25a38fe1a5f069d6d35a40976e17)
    - 결과
        
        ```go
        HTTP/1.1 200 OK
        Date: Fri, 30 Dec 2022 05:10:50 GMT
        Content-Length: 180
        Content-Type: text/plain; charset=utf-8
        Connection: close
        
        {
          "data": "Third",
          "hash": "110e8eb40de97b0929e53a3b0cd2d697a4bd25a38fe1a5f069d6d35a40976e17",
          "prevhash": "64bf97f63405065b829e45a1f99321bf3f55a02fb352779759c95adc1a83e8f4",
          "height": 4
        }
        ```

___

## RESTful 동작시키기

- DB에 등록된 모든 블록 가져오기
- DB 연결 끊기 추가
- 정상적인 종료 처리
    - runtime.Goexit()
        
        [runtime](https://pkg.go.dev/runtime#Goexit)
        
        ```go
        func Goexit()
        ```
        
        - 해당 함수는 모든 함수를 제거하고 마지막으로 defer를 모두 종료시키고 종료한다.
- RESTful 동작시키기
- 소스 코드
    - rest/rest.go
        
        ```go
        // ...
        func blocks(rw http.ResponseWriter, r *http.Request) {
        	switch r.Method {
        	case "GET":
        
        		json.NewEncoder(rw).Encode(blockchain.Blockchain().Blocks())
        	case "POST":
        
        		var addBlockBody addBlockBody
        
        		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
        		blockchain.Blockchain().AddBlock(addBlockBody.Message)
        		rw.WriteHeader(http.StatusCreated)
        	}
        }
        // ...
        ```
        
    - blockchain/chain.go
        
        ```go
        // 등록된 블록을 가져오는 함수
        
        func (b *blockchain) Blocks() []*Block {
        	// 찾은 블록들을 저장할 Block 포인터를 저장할 슬라이스를 만든다.
        	var blocks []*Block
        	// 최근에 생성된 블록의 해시값을 가져온다.
        	hashCursor := b.NewestHash
        	for {
        		// 최근 블록 부터 가져온다
        		// 무조건 찾을 수 있기 때문에 error 처리를 따로 하지 않는다.
        		block, _ := FindBlock(hashCursor)
        		// 찾은 블록을 []*block에 넣는다.
        		blocks = append(blocks, block)
        		// 가져온 블록의 이전 해시값이 빈값이 아니면 이전 블록이 있는 것이므로
        		// 이전 블록의 해시 값을 가져와 가르킨다.
        		if block.PrevHash != "" {
        			hashCursor = block.PrevHash
        		} else {
        			// PrevHash가 없는 Genesis Block면
        			break
        		}
        	}
        	return blocks
        }
        ```
        
- 실행 결과
    - DB에 등록된 블록 모두 가져오기
        - Method : GET
        - URL : [http://localhost:4000/blocks](http://localhost:4000/blocks)
        - 결과
            
            ```go
            HTTP/1.1 200 OK
            Date: Fri, 30 Dec 2022 07:12:47 GMT
            Content-Length: 845
            Content-Type: text/plain; charset=utf-8
            Connection: close
            
            [
              {
                "data": "Data for my block",
                "hash": "c57fca35b986fc973b8f1b535ae9970c081606d5bffde0c3def73cba9cf50bf9",
                "prevhash": "110e8eb40de97b0929e53a3b0cd2d697a4bd25a38fe1a5f069d6d35a40976e17",
                "height": 5
              },
              {
                "data": "Third",
                "hash": "110e8eb40de97b0929e53a3b0cd2d697a4bd25a38fe1a5f069d6d35a40976e17",
                "prevhash": "64bf97f63405065b829e45a1f99321bf3f55a02fb352779759c95adc1a83e8f4",
                "height": 4
              },
              {
                "data": "Second",
                "hash": "64bf97f63405065b829e45a1f99321bf3f55a02fb352779759c95adc1a83e8f4",
                "prevhash": "3fd9ee8fc98866d3fb2946af3aa9d32048d60b71467f1d6c89f07b4ba77d21cc",
                "height": 3
              },
              {
                "data": "First",
                "hash": "3fd9ee8fc98866d3fb2946af3aa9d32048d60b71467f1d6c89f07b4ba77d21cc",
                "prevhash": "81f2ced897805e5539e750784e8d12bff104712be9bf8845ce52e006b0f3252e",
                "height": 2
              },
              {
                "data": "Genesis Block",
                "hash": "81f2ced897805e5539e750784e8d12bff104712be9bf8845ce52e006b0f3252e",
                "height": 1
              }
            ]
            ```
            
    - 블록 추가하기
        - Method : POST
        - URL : [http://localhost:4000/blocks](http://localhost:4000/blocks)
            
            ```json
            {
              "message" : "Blockchain Test"
            }
            ```
            
        - 결과
            
            ```json
            HTTP/1.1 201 Created
            Date: Fri, 30 Dec 2022 07:14:17 GMT
            Content-Length: 0
            Connection: close
            ```
            
            ```json
            HTTP/1.1 200 OK
            Date: Fri, 30 Dec 2022 07:14:23 GMT
            Content-Length: 1035
            Content-Type: text/plain; charset=utf-8
            Connection: close
            
            [
              {
                "data": "Blockchain Test",
                "hash": "b2b2a8fb45aa0631ee0c7f00d48fe508ac29155b9aac4f21628d320d0e3ce39f",
                "prevhash": "c57fca35b986fc973b8f1b535ae9970c081606d5bffde0c3def73cba9cf50bf9",
                "height": 6
              },
              {
                "data": "Data for my block",
                "hash": "c57fca35b986fc973b8f1b535ae9970c081606d5bffde0c3def73cba9cf50bf9",
                "prevhash": "110e8eb40de97b0929e53a3b0cd2d697a4bd25a38fe1a5f069d6d35a40976e17",
                "height": 5
              },
              {
                "data": "Third",
                "hash": "110e8eb40de97b0929e53a3b0cd2d697a4bd25a38fe1a5f069d6d35a40976e17",
                "prevhash": "64bf97f63405065b829e45a1f99321bf3f55a02fb352779759c95adc1a83e8f4",
                "height": 4
              },
              {
                "data": "Second",
                "hash": "64bf97f63405065b829e45a1f99321bf3f55a02fb352779759c95adc1a83e8f4",
                "prevhash": "3fd9ee8fc98866d3fb2946af3aa9d32048d60b71467f1d6c89f07b4ba77d21cc",
                "height": 3
              },
              {
                "data": "First",
                "hash": "3fd9ee8fc98866d3fb2946af3aa9d32048d60b71467f1d6c89f07b4ba77d21cc",
                "prevhash": "81f2ced897805e5539e750784e8d12bff104712be9bf8845ce52e006b0f3252e",
                "height": 2
              },
              {
                "data": "Genesis Block",
                "hash": "81f2ced897805e5539e750784e8d12bff104712be9bf8845ce52e006b0f3252e",
                "height": 1
              }
            ]
            ```
            
    - Hash 값을 통해 하나의 블록 조회하기
        - Method : GET
        - URL : [http://localhost:4000/blocks](http://localhost:4000/blocks/c57fca35b986fc973b8f1b535ae9970c081606d5bffde0c3def73cba9cf50bf9)/:hash
        - 실행 결과
            
            ```json
            HTTP/1.1 200 OK
            Date: Fri, 30 Dec 2022 07:15:29 GMT
            Content-Length: 190
            Content-Type: text/plain; charset=utf-8
            Connection: close
            
            {
              "data": "Blockchain Test",
              "hash": "b2b2a8fb45aa0631ee0c7f00d48fe508ac29155b9aac4f21628d320d0e3ce39f",
              "prevhash": "c57fca35b986fc973b8f1b535ae9970c081606d5bffde0c3def73cba9cf50bf9",
              "height": 6
            }
            ```

# 작업 증명 구현하기(PoW)

- 컴퓨터가 풀기는 어렵지만 검증하기는 쉬운 방법
- n개의 0으로 시작하는hash를 찾도록 하자
  - n개는 Difficulty에 의해 결정된다.
- 네트워크가 Client의 해시의 시작이 n개의 0으로 되어있는지 검증한다.
- 시간의 흐름에 따라 Difficulty 값을 바뀌도록 한다.
- Ex

    ```go
      package main
      
      import (
      "crypto/sha256"
      "fmt"
      )
      
      func main() {
      hash := sha256.Sum256([]byte("hello"))
      fmt.Printf("%x\n", hash)
      }
    ```

    해당 문자열로 만든 해시값은 앞에 0이 하나도 없다.

    → 해시는 결정론적 함수이기에 입력값으로 수정으로 출력값을 바꿔줘야한다.

- 블록체인에서는 블록의 정보를 수정할 수 없다.
  - 해시값을 수정하면 해당 블록을 사용하지 못한다.
  - Data는 유저가 보내주는 것이라 수정할 수 없다.
- Nonce 값은 채굴자가 변경할 수 있는 유일한 값이다.
  - Nonce값을 변경해가며 해시값을 조건에 맞게 찾아내야한다.
  
      ```go
        func main() {
          hash := sha256.Sum256([]byte("hello1"))
          fmt.Printf("%x\n", hash)
        }
       ```

       ```go
        func main() {
           hash := sha256.Sum256([]byte("hello2"))
           fmt.Printf("%x\n", hash)
        }
       ```

      ```go
        func main() {
          hash := sha256.Sum256([]byte("hello3"))
          fmt.Printf("%x\n", hash)
        }
       ```

      ```go
        func main() {
          hash := sha256.Sum256([]byte("hello4"))
          fmt.Printf("%x\n", hash)
        }
      ```

    - 못찾았다.

## 구현

네트워크 : difficulty의 값만큼 0의 개수를 증가시켜 난이도를 어렵게 만든다

채굴자 : nonce 값을 변경해가며 0의 개수가 맞는 것을 찾아낸다.

- 소스 코드
  - main.go

      ```go
        package main
        
        import (
         "crypto/sha256"
         "fmt"
         "strings"
        )
        
        func main() {
         difficulty := 2
         // target := "0" * 2
         // 첫번째 인자값을 두번째 인자값 만큼 연결해서 출력해준다.
         target := strings.Repeat("0", difficulty)
         nonce := 1
         for {
          // 16진수 string으로 변환
          hash := fmt.Sprintf("%x", sha256.Sum256([]byte("hello"+fmt.Sprint(nonce))))
          fmt.Printf("Hash:%s\nTarget:%s\nNonce:%d\n\n", hash, target, nonce)
          if strings.HasPrefix(hash, target) {
           return
          } else {
           nonce++
          }
         }
        
        }
       ```

- 실행 결과
  - difficulty가 2일 때
      ```json
        Hash:001b92541ed0a22b0cb89018b561d895503206c0082c0ecf2d0b7e5182191eed
        Target:00
        Nonce:227
      ```

    - 227번만에 찾았다.
  - difficulty가 3일 때

      ```json
        Hash:0006bc9ad4253c42e32b546dc17e5ea3fedaecdabef371b09906cea9387e8695
        Target:000
        Nonce:10284
      ```

    - 10284번 걸렸다.
  - difficulty가 4일 때

      ```json
        Hash:0000e49eab06aa7a6b3aef7708991b91a7e01451fd67f520b832b89b18f4e7de
        Target:0000
        Nonce:60067
      ```
    - 60067번
- 난이도가 조금만 올라가도 엄청나게 연산이 많이 필요한 것을 느낄 수 있다.
- 실제 비트코인은 좀 더 복잡하다

---

## 상기 코드를 추가하여 기존 코드에 채굴기능 만들기

- 블록 구조체 자체를 해쉬함수를 통해 해시하기
- 소스 코드
  - blockchain/block.go

      ```go
        func (b *Block) mine() {
         target := strings.Repeat("0", b.Difficulty)
         for {
          blockAsString := fmt.Sprint(b)
          // 블록을 string으로 바꾼 후 해쉬로 변환시킨다음 16진수 string으로 다시 변환한다.
          hash := fmt.Sprintf("%x", sha256.Sum256([]byte(blockAsString)))
          fmt.Printf("Block as String:%s\nHash:%s\nTarget:%s\nNonce:%d\n\n\n", blockAsString, hash, target, b.Nonce)
          if strings.HasPrefix(hash, target) {
           b.Hash = hash
           break
          } else {
           b.Nonce++
          }
        
         }
        }
      ```

- 실행 결과
  - Method : GET
  - URL : [http://localhost:4000/blocks](http://localhost:4000/blocks)

    ```go
    HTTP/1.1 200 OK
    Date: Sat, 31 Dec 2022 02:47:41 GMT
    Content-Length: 138
    Content-Type: text/plain; charset=utf-8
    Connection: close
    
    [
      {
        "data": "Genesis Block",
        "hash": "0056f988ce765e06b2fccc508947a1771b1822f889c3594bc174aa6032fc688c",
        "height": 1,
        "defficulty": 2,
        "nonce": 76
      }
    ]
    ```
___
## Defficulty 자동으로 수정되게 변경하기

- 비트코인에서 착안
  - 1개의 블록 생성에 10분
  - 2016개의 블록을 생성하는데 2주
    - 2주보다 적게 걸렷으면 Difficulty 늘리기
    - 2주보다 많이 걸렸으면 Difficulty 줄이기

### 8분~ 12분 사이 5개 블록이 생성되는 것을 기준으로 코드 작성

- 조건
  - 초기 난이도는 2
    - 해시 값이 맨앞자리 부터 00으로 시작하면 검증완료
  - 5개 블록 생성 시 8분~12분 사이에 생성된 블록은 난이도 유지
  - 8분 미만 시 5개 블록이 생성되면
    - 난이도 높이기
  - 12분 초과 시 5개 블록이 생성되면
    - 난이도 낮추기
- Block에 Timestamp 추가하기
  - 블록의 생성에 얼마나 시간이 걸렸는지 확인하기 위함
- REST API로 확인하기 위해 내용 추가
- 소스 코드
  - chain.go

      ```go
        
        const (
         defaultDifficulty  int = 2
         difficultyInterval int = 5
         blockInterval      int = 2
         allowedRange       int = 2
        )
        
        type blockchain struct {
         // 최근에 등록된 Hash
         NewestHash string `json:"newestHash"`
         // 블록의 수
         Height            int `json:"height"`
         // 현재 난이도
         CurrentDifficulty int `json:"currentdifficulty"`
        }
        
        func (b *blockchain) AddBlock(data string) {
         block := createBlock(data, b.NewestHash, b.Height+1)
         b.NewestHash = block.Hash
         b.Height = block.Height
         b.CurrentDifficulty = block.Difficulty
         b.persist()
        }
        
        func (b *blockchain) recalculateDifficulty() int {
         allBlocks := b.Blocks()
         // 블록 슬라이스에 가장 최근 블록이 앞에 들어간다.
         newestBlock := allBlocks[0]
         // 가장 최근에 난이도가 재설정된 블록은 allBlock[5-1]이다.
         lastRecalculatedBlock := allBlocks[difficultyInterval-1]
         // 두 블록 사이에 걸린시간 : 최근 생성된 블록의 시간 - 블록의 난이도가 재설정된 후 생성된 블록의 시간
         // 타임스탬프가 Unix Time이기에 초단위로 변경해줘야한다.
         actualTime := (newestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60)
         // 예상 시간 : 5 * 2 = 10분 기준으로 난이도를 설정한다.
         expectedTime := difficultyInterval * blockInterval
        
         // 10분을 기준으로 앞뒤로 2분씩 범위안에만 들어오면 난이도를 유지한다.
         // 예상 시간 보다 빠르게 블록이 생성되면 난이도를 1만큼 증가시킨다.
         if actualTime <= (expectedTime - allowedRange) {
          return b.CurrentDifficulty + 1
         } else if actualTime >= (expectedTime + allowedRange) {
          // 예상 시간 보다 느리게 블록이 생성되면 난이도를 1만큼 감소시킨다.
          return b.CurrentDifficulty - 1
         }
         return b.CurrentDifficulty
        }
        
        func (b *blockchain) difficulty() int {
         // 제네시스 블록체인은 Difficulty가 2다.
         if b.Height == 0 {
          return defaultDifficulty
         } else if b.Height%difficultyInterval == 0 {
          // 비트코인은 2016개, 우리의 블록체인은 5개 블록마다 체크하여 난이도를 조정한다
          return b.recalculateDifficulty()
         } else {
          // 난이도가 변경된 후 블록이 5개가 추가되지 않았으면 현재 난이도를 그대로 유지한다.
          return b.CurrentDifficulty
         }
        }
      ```

  - block.go

      ```go
        type Block struct {
         Data       string `json:"data"`
         Hash       string `json:"hash"`
         PrevHash   string `json:"prevhash,omitempty"`
         Height     int    `json:"height"`
         Difficulty int    `json:"defficulty"`
         Nonce      int    `json:"nonce"`
         Timestamp  int    `json:"timestamp"`
        }
        
        func (b *Block) mine() {
         target := strings.Repeat("0", b.Difficulty)
         for {
          // time.Now().Unix() : int64를 반환한다. 1970년 1월 1일 UTC로부터 흐른 시간을 초단위로
          b.Timestamp = int(time.Now().Unix())
          hash := utils.Hash(b)
        
         if strings.HasPrefix(hash, target) {
          b.Hash = hash
          break
         } else {
          b.Nonce++
         }
        }
        
        func createBlock(data string, prevHash string, height int) *Block {
         block := &Block{
          Data:       data,
          Hash:       "",
          PrevHash:   prevHash,
          Height:     height,
          Difficulty: Blockchain().difficulty(),
          Nonce:      0,
         }
         block.mine()
         block.persist()
         return block
        }
        
      ```

- 실행 결과
  - Method : GET
  - URL : [http://localhost:4000/blocks](http://localhost:4000/blocks)
  - 기능 : 현재 생성된 블록 모두 가져오기
  - 결과

    ```go
    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Sat, 31 Dec 2022 04:28:23 GMT
    Content-Length: 162
    Connection: close
    
    [
      {
        "data": "Genesis Block",
        "hash": "007b0d340726556946f5121d17bb7bb2f6b892e037a6e4cd783336aa0ff497fe",
        "height": 1,
        "defficulty": 2,
        "nonce": 234,
        "timestamp": 1672460696
      }
    ]
    ```

  - Method : GET
  - URL : [http://localhost:4000/status](http://localhost:4000/status)
  - 기능 : 현재 체인의 상태 가져오기
  - 결과

      ```go
        HTTP/1.1 200 OK
        Content-Type: application/json
        Date: Sat, 31 Dec 2022 04:27:48 GMT
        Content-Length: 115
        Connection: close
        
        {
          "newestHash": "007b0d340726556946f5121d17bb7bb2f6b892e037a6e4cd783336aa0ff497fe",
          "height": 1,
          "currentdifficulty": 2
        }
      ```

  - Method : POST
  - URL : [http://localhost:4000/blocks](http://localhost:4000/blocks)
  - 기능 : JSON형태의 값을 받아 블록 하나 생성하기
    - 전달 된 값

        ```go
            {
              "message" : "Blockchain Test"
            }
        ```

  - 결과

      ```go
        HTTP/1.1 201 Created
        Content-Type: application/json
        Date: Sat, 31 Dec 2022 04:29:35 GMT
        Content-Length: 0
        Connection: close
      ```

### 시나리오 진행

    총 5개의 블록을 생성했을 때 8분 미만으로 5개의 블록이 생성되면 난이도가 1 올라간다.

  - 초기 상태

      ```json
        {
          "newestHash": "00005e8d76daa1d457e5cde329901a1c1bd352cb00e4d626f3571c6431d2dd36",
          "height": 1,
          "currentdifficulty": 2
        }
      ```

  - 블록이 6번째에 난이도가 변한것을 알 수 있다.

      ```json
        {
          "newestHash": "000e4683d1ca7256e62118d1cffde70b5957f9813d32e6f7f0b4873f5df1d81c",
          "height": 6,
          "currentdifficulty": 3
        }
      ```

  - 난이도가 재설정된 블록 부터 시작해서 8분 미만으로 5개의 블록이 추가로 생성되었기에 난이도가 또 올라간 것을 확인할 수 있다.

      ```json
        {
          "newestHash": "0000be41b16c285e84f5b5deec30317127f1612e4c66e8d5016e3ce484a100ea",
          "height": 11,
          "currentdifficulty": 4
        }
      ```
---

# 트랜잭션 구축

**비트코인은 UTXO를 이용하여 트랜잭션을 만든다.**

Tx

- TxIn[] : 거래를 실행하기 이전에 내 주머니에 있는 돈
- TxOut[] : 거래가 끝났을 때 각각의 사람들이 갖고있는 액수

```go
Tx
 TxIn[5천원(프동프동)]
 TxOut[0원(프동프동), 5천원(A User)]
```

- 만약 5천원을 주고싶은데 1만원 지폐로 가지고 있을 경우

```go
TX
 TxIn[1만원(프동프동)]
 TxOut[5천원(A User)/5천원(프동프동)]
```

- 채굴자(코인베이스의 거래)가 코인을 딱 채굴했을 시 상태
  - 채굴자가 코인을 채굴했을 시 Input값은 블록체인이다.
  - Output 값은 채굴자다.

```go
Tx
 TxIn[$10(blockchain)]
 TXOut[$10(miner)]
```

## 코인베이스에서 채굴자에게 코인을 주도록 만들고 트랜잭션에 기록하기

- 시나리오
  - 블록 생성 시(채굴 시)
    - 프동프동이란 사람이 COINBASE로 부터 채굴한 코인을 받는다

        ```go
        Tx 
         TxIn(소유자 : COINBASE, 가진 코인의 수: 50개)
         TxOut(소유자 : 프동프동, 가진 코인의 수 : 50개
        ```

- 소스 코드
  - blockchain/block.go

      ```go
        func createBlock(prevHash string, height int) *Block {
         block := &Block{
        
          Hash:       "",
          PrevHash:   prevHash,
          Height:     height,
          Difficulty: Blockchain().difficulty(),
          Nonce:      0,
          // 블록 생성 시 채굴자의 이름을 가지고 트랜잭션을 만들고 블록에 포함시킨다.
          Transactions: []*Tx{makeCoinbaseTx("fdongfdong")},
         }
         block.mine()
         block.persist()
         return block
        }
      ```

  - transaction/transaction.go

      ```go
        package blockchain
        
        import (
         "coin/exam36/utils"
         "time"
        )
        
        const (
         minerReward int = 50
        )
        
        type Tx struct {
         Id        string   `json:"id"`
         Timestamp int      `json:"timestamp"`
         TxIns     []*TxIn  `json:"txins"`
         TxOuts    []*TxOut `json:"txouts"`
        }
        
        func (t *Tx) getId() {
         t.Id = utils.Hash(t)
        }
        
        type TxIn struct {
         Owner  string `json:"owner"`
         Amount int    `json:"amount"`
        }
        
        type TxOut struct {
         Owner  string `json:"owner"`
         Amount int    `json:"amount"`
        }
        // 코인베이스에서 채굴자에게 코인을 주기 위해(거래) 트랜잭션을 만든다.
        func makeCoinbaseTx(address string) *Tx {
         txIns := []*TxIn{
          {"COINBASE", minerReward},
         }
         txOuts := []*TxOut{
          {address, minerReward},
         }
         tx := Tx{
          Id:        "",
          Timestamp: int(time.Now().Unix()),
          TxIns:     txIns,
          TxOuts:    txOuts,
         }
         tx.getId()
         return &tx
        }
      ```

- 실행 결과
  - 블록이 생성되었을 때

      ```json
        HTTP/1.1 200 OK
        Content-Type: application/json
        Date: Sat, 31 Dec 2022 05:42:10 GMT
        Content-Length: 342
        Connection: close
        
        [
          {
            "hash": "007870047badd28511c43d3b4d8108fd9e346bc89113636ecfabf396dda01541",
            "height": 1,
            "defficulty": 2,
            "nonce": 226,
            "timestamp": 1672465330,
            "Transactions": [
              {
                "id": "bebd0536f46ded4f0e2c053c7e002dafb0fd1cc68233f73ae41850dd19b39ed8",
                "timestamp": 1672465330,
                "txins": [
                  {
                    "owner": "COINBASE",
                    "amount": 50
                  }
                ],
                "txouts": [
                  {
                    "owner": "fdongfdong",
                    "amount": 50
                  }
                ]
              }
            ]
          }
        ]
      ```
---

## 보유한 자산 조회하기

- 트랜잭션에 포함된 owner로 자산 조회하기
- RESTful하게 가져오기
- 소스 코드
  - main.go

        ```go
        
        // 쿼리에 total 값이 true이면 해당하는 address의 자산을 모두 합쳐서 Client에게 보내준다.
        func balance(rw http.ResponseWriter, r *http.Request) {
         vars := mux.Vars(r)
         address := vars["address"]
         total := r.URL.Query().Get("total")
         switch total {
         case "true":
          amount := blockchain.Blockchain().BalanceByAddress(address)
          json.NewEncoder(rw).Encode(balanceResponse{address, amount})
         default:
          utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blockchain().TxOutsByAddress(address)))
         }
        }
        
        func Start(aPort int) {
         router := mux.NewRouter()
         router.HandleFunc("/balance/{address}", balance).Methods("GET")
         fmt.Printf("Listening on http://localhost%s\n", port)
         log.Fatal(http.ListenAndServe(port, router))
        }
        ```

  - blockchain/chain.go

    ```go
        // 블록의 트랜잭션 출력값을 모두 가져오는 함수
        func (b *blockchain) txOuts() []*TxOut {
         var txOuts []*TxOut
         // 모든 블록을 가져온다.
         blocks := b.Blocks()
         for _, block := range blocks {
          // 모든 블록의 TxOuts을 가져온다.
          for _, tx := range block.Transactions {
           txOuts = append(txOuts, tx.TxOuts...)
          }
         }
         return txOuts
        }
        
        // 블록의 트랜잭션 출력값 리스트에서 address에 해당하는 값들만 찾아오는 함수
        func (b *blockchain) TxOutsByAddress(address string) []*TxOut {
         var ownedTxOuts []*TxOut
         txOuts := b.txOuts()
         for _, txOut := range txOuts {
          if txOut.Owner == address {
           ownedTxOuts = append(ownedTxOuts, txOut)
          }
         }
         return ownedTxOuts
        }
        
        // 해당하는 address의 자산을 모두 합쳐서 반환하는 함수
        func (b *blockchain) BalanceByAddress(address string) int {
         txOuts := b.TxOutsByAddress(address)
         var amount int
         for _, txOut := range txOuts {
          amount += txOut.Amount
         }
         return amount
        }
    ```

- 실행 결과
  - 쿼리로 total을 줬을 때
    - Method : GET
    - URL : [http://localhost:4000/balance/fdongfdong?total=true](http://localhost:4000/balance/fdongfdong?total=true)
    - 기능 : 해당하는 address에 모든 자산을 합쳐서 가져온다.

        ```json
        HTTP/1.1 200 OK
        Content-Type: application/json
        Date: Sat, 31 Dec 2022 09:08:43 GMT
        Content-Length: 39
        Connection: close
        
        {
          "address": "fdongfdong",
          "balance": 100
        }
        ```

  - 쿼리를 주지 않았을 때
    - Method : GET
    - URL : [http://localhost:4000/balance/fdongfdong](http://localhost:4000/balance/fdongfdong)
    - 기능 :  각각의 트랜잭션 출력값을 가져온다.

        ```json
        HTTP/1.1 200 OK
        Content-Type: application/json
        Date: Sat, 31 Dec 2022 09:08:53 GMT
        Content-Length: 72
        Connection: close
        
        [
          {
            "owner": "fdongfdong",
            "amount": 50
          },
          {
            "owner": "fdongfdong",
            "amount": 50
          }
        ]
        ```
---

# Mempoll(메모리풀)

아직 확정되지 않은 거래내역을 보관하는 곳

→ 아직 Confirm 받지 않은 Transaction들이 들어가는 곳

## Mempool에 트랜잭션 생성하기

[소스 코드](https://github.com/FDongFDong/BlockChain_study/tree/main/exam38)
  - blockchain/transaction.go

    ```go
    type mempool struct {
      Txs []*Tx
    }
    
    // 비어있는 mempool 생성
    var Mempool *mempool = &mempool{}
    
    type Tx struct {
      Id        string   `json:"id"`
      Timestamp int      `json:"timestamp"`
      TxIns     []*TxIn  `json:"txins"`
      TxOuts    []*TxOut `json:"txouts"`
    }
    
    type TxIn struct {
      Owner  string `json:"owner"`
      Amount int    `json:"amount"`
    }
    
    type TxOut struct {
      Owner  string `json:"owner"`
      Amount int    `json:"amount"`
    }
    
    // 유효한 트랜잭션을 생성하려면, 유저가 input에 돈이 들어있다는걸 보여주면 된다.
    // Transaction Input을 가져와서 Transaction Output을 만들면 해당 Transaction은 유효해진다.
    func makeTx(from, to string, amount int) (*Tx, error) {
      // from의 잔고가 보내고자 하는 금액 보다 적으면 에러 출력
      if Blockchain().BalanceByAddress(from) < amount {
      return nil, errors.New("not enough money")
      }
      // 새로운 트랜잭션을 만들기 위해 txIns, txOuts을 생성한다.
      var txIns []*TxIn
      var txOuts []*TxOut
    
      total := 0
      // 이전 TxOut으로 TxInput을 만들기 위해 요청한 사용자의 TxOut List를 가져온다.
      oldTxOuts := Blockchain().TxOutsByAddress(from)
      // total값에 TxOut을 모아서 amount 값이 되도록 한다.
      for _, txOut := range oldTxOuts {
      if total > amount {
        break
      }
      // 보내려는 값이 일치할 때 까지 
      txIn := &TxIn{txOut.Owner, txOut.Amount}
      txIns = append(txIns, txIn)
      total += txIn.Amount
      }
      change := total - amount
      // TxOut에게 줄 거스름돈이 있을 경우
      if change != 0 {
      changeTxOut := &TxOut{from, change}
      txOuts = append(txOuts, changeTxOut)
      }
      txOut := &TxOut{to, amount}
      txOuts = append(txOuts, txOut)
      // 새로운 트랜잭션을 만들어준다.
      tx := &Tx{
      Id:        "",
      Timestamp: int(time.Now().Unix()),
      TxIns:     txIns,
      TxOuts:    txOuts,
      }
      tx.getId()
      return tx, nil
    }
    
    // 트랜잭션을 Mempool에 추가한다. 트랜잭션을 생성하지는 않는다.
    // 어떠한 이유로 트랜잭션을 블록에 추가할 수 없으면 error을 리턴해준다.
    func (m *mempool) AddTx(to string, amount int) error {
      tx, err := makeTx("fdongfdong", to, amount)
      if err != nil {
      return err
      }
      // 트랜잭션이 정상적으로 만들어졌다면 Mempool에 추가해준다.
      m.Txs = append(m.Txs, tx)
      return nil
    }
    ```

  - rest/rest.go

    ```go
    func mempool(rw http.ResponseWriter, r *http.Request) {
      utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Mempool.Txs))
    }
    
    func transactions(rw http.ResponseWriter, r *http.Request) {
      var payload addTxPayload
      utils.HandleErr(json.NewDecoder(r.Body).Decode(&payload))
      err := blockchain.Mempool.AddTx(payload.To, payload.Amount)
      if err != nil {
      json.NewEncoder(rw).Encode(errorResponse{"not enough funds"})
      }
      rw.WriteHeader(http.StatusCreated)
    }
    
    func Start(aPort int) {
      router := mux.NewRouter()
      port = fmt.Sprintf(":%d", aPort)
      router.Use(jsonContentTypeMiddleware)
      // ...
      router.HandleFunc("/mempool", mempool).Methods("GET")
      router.HandleFunc("/transactions", transactions).Methods("POST")
      fmt.Printf("Listening on http://localhost%s\n", port)
    }
    ```

- 실행 결과

  - 기능 : mempool에 들어있는 트랜잭션 확인
  - Method : GET
  - URL : [http://localhost:4000/mempool](http://localhost:4000/mempool)

  - 시나리오
    - **기능 : 초기(블록체인) 상태 확인**
    - Method : GET
    - URL : [http://localhost:4000/status](http://localhost:4000/status)

        ```json
        HTTP/1.1 200 OK
        Content-Type: application/json
        Date: Mon, 02 Jan 2023 00:44:25 GMT
        Content-Length: 115
        Connection: close
        
        {
          "newestHash": "00b904fb8d5a30754d81f8362b7bca54ba4d073689e8ab4af1d2be130bd085c1",
          "height": 1,
          "currentdifficulty": 2
        }
        ```

    - **기능 : 블록체인 내에 있는 블록 확인**
    - Method : GET
    - URL : [http://localhost:4000/blocks](http://localhost:4000/blocks)

        ```json
        HTTP/1.1 200 OK
        Content-Type: application/json
        Date: Mon, 02 Jan 2023 00:45:02 GMT
        Content-Length: 342
        Connection: close
        
        [
          {
            "hash": "00b904fb8d5a30754d81f8362b7bca54ba4d073689e8ab4af1d2be130bd085c1",
            "height": 1,
            "defficulty": 2,
            "nonce": 172,
            "timestamp": 1672620199,
            "transactions": [
              {
                "id": "c7186f0a495a53ec522353da9aa9d2ff38e0f5c23e4cfd91b5b209fdd6582932",
                "timestamp": 1672620199,
                "txins": [
                  {
                    "owner": "COINBASE",
                    "amount": 50
                  }
                ],
                "txouts": [
                  {
                    "owner": "fdongfdong",
                    "amount": 50
                  }
                ]
              }
            ]
          }
        ]
        ```

    - **기능 : 잔고가 부족한 상태에서 전송(트랜잭션 생성)**
    - Method : POST
    - URL :  [http://localhost:4000/transactions](http://localhost:4000/transactions)

        ```json
        {
          "to" : "uou",
          "amount" : 80
        }
        ```

    - 결과

        ```json
        HTTP/1.1 200 OK
        Content-Type: application/json
        Date: Mon, 02 Jan 2023 00:46:48 GMT
        Content-Length: 36
        Connection: close
        
        {
          "errorMessage": "not enough funds"
        }
        ```

    - 기능 : 잔고가 부족하므로 **추가적인 채굴 진행**
    - Method : POST
    - URL : [http://localhost:4000/blocks](http://localhost:4000/blocks)

        ```json
        {
          "message" : "Blockchain Test"
        }
        ```

    - 실행 결과

        ```json
        HTTP/1.1 201 Created
        Content-Type: application/json
        Date: Mon, 02 Jan 2023 00:48:32 GMT
        Content-Length: 0
        Connection: close
        ```

    - **기능 : 잔고 확인**
    - Method : GET
    - URL : [http://localhost:4000/balance/fdongfdong](http://localhost:4000/balance/fdongfdong?total=true)

        ```json
        HTTP/1.1 200 OK
        Content-Type: application/json
        Date: Mon, 02 Jan 2023 00:49:18 GMT
        Content-Length: 39
        Connection: close
        
        {
          "address": "fdongfdong",
          "balance": 100
        }
        ```

    - **기능 : 송금 진행(트랜잭션 생성)**
    - Method : POST
    - URL : [http://localhost:4000/transactions](http://localhost:4000/transactions)

        ```json
        {
          "to" : "uou",
          "amount" : 80
        }
        ```

    - 실행 결과

        ```json
        HTTP/1.1 201 Created
        Content-Type: application/json
        Date: Mon, 02 Jan 2023 00:50:34 GMT
        Content-Length: 0
        Connection: close
        ```

    - 기능 : Mempool 확인
    - Method : GET
    - URL : [http://localhost:4000/mempool](http://localhost:4000/mempool)
    - 실행 결과

        ```json
        HTTP/1.1 200 OK
        Content-Type: application/json
        Date: Mon, 02 Jan 2023 00:53:01 GMT
        Content-Length: 253
        Connection: close
        
        [
          {
            "id": "146702f253cf4cc4d5dfa0048302d6d3df07284a70fb940f589e15cb5651d36a",
            "timestamp": 1672620634,
            "txins": [
              {
                "owner": "fdongfdong",
                "amount": 50
              },
              {
                "owner": "fdongfdong",
                "amount": 50
              }
            ],
            "txouts": [
              {
                "owner": "fdongfdong",
                "amount": 20
              },
              {
                "owner": "uou",
                "amount": 80
              }
            ]
          }
        ]
        ```
