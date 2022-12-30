# BlockChain_study
Go 언어로 블록체인 스터디

- [BlockChain\_study](#blockchain_study)
  - [Genesis Block 만들어보기](#genesis-block-만들어보기)
  - [Genesis Block, Second Block, ... 만들어보기](#genesis-block-second-block--만들어보기)
  - [Refactoring, Singleton](#refactoring-singleton)
    - [block 추가 및 block 정보 가져오기](#block-추가-및-block-정보-가져오기)
    - [Web Server 1](#web-server-1)
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

        ```go
        > go run main.go -mode=rest -port=2000
        Listening on http://localhost:2000
        ```

  - go run main.go -mode=html -port=8000

        ```go
        > go run main.go -mode=html -port=8000
        Listening on http://localhost8000
        ```

  - 잘못된 flag

        ```go
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
    
    ```go
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
    
- 실행 결과
    - 터미널에 boltbrowser “dbname”
        
        ![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/df741855-b08f-4f91-9750-7bbd74ada7ec/Untitled.png)
        
        ![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b6aa90f6-372b-4b79-973f-ddefa4b4b237/Untitled.png)
---

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
    
- 실행 결과
    
    ![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d01cf690-783c-4ca6-b677-9ac78f47b739/Untitled.png)
    
    ![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/06edadd3-c649-453a-8fef-beb52aca3149/Untitled.png)
    
    ![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/810d5214-5de1-446a-98b3-0e1bcd25c587/Untitled.png)
    
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
