package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Peers map[string]*peer = make(map[string]*peer)

// 연결을 저장하기 위함
type peer struct {
	key     string
	address string
	port    string
	conn    *websocket.Conn
	inbox   chan []byte
}

// clean up 함수
func (p *peer) close() {
	p.conn.Close()
	delete(Peers, p.key)
}
func (p *peer) read() {
	// Error이 나오면 해당 peer를 삭제한다.
	defer p.close()
	for {
		_, m, err := p.conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("%s", m)
	}
}
func (p *peer) write() {
	defer p.close()
	for {
		// 메세지가 오기를 기다린다.
		m, ok := <-p.inbox
		if !ok {
			break
		}
		p.conn.WriteMessage(websocket.TextMessage, m)
	}
}
func initPeer(conn *websocket.Conn, address, port string) *peer {
	key := fmt.Sprintf("%s:%s", address, port)
	p := &peer{
		conn:    conn,
		inbox:   make(chan []byte),
		address: address,
		key:     key,
		port:    port,
	}
	go p.read()
	go p.write()
	Peers[key] = p
	return p
}

// :3000 포트의 peers
// {
// 	"127.0.0.1:4000": conn
// }
// :4000 포트의 peers
// {
// 	"127.0.0.1:4000": conn
// }
