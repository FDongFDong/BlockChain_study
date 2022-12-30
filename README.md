# BlockChain_study
Go 언어로 블록체인 스터디
- [DB에 저장된 블록 GET 메서드로 가져오기](#DB에-저장된-블록-GET-메서드로-가져오기)
- [DB로 부터 저장된 블록 데이터 불러와 콘솔로 출력](#DB로부터-저장된-블록-데이터-불러와-콘솔로-출력)
- [RESTful Blockchain](#RESTful-동작시키기)


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
