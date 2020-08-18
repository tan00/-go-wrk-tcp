package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var oncePrintSend, oncePrintRecv sync.Once

//StartClient  url_: ip:port
func StartClient(url_ string, requestBody []byte, dka bool, responseChan chan *Response, waitGroup *sync.WaitGroup, tc int) {
	defer waitGroup.Done()

	conn, err := net.DialTimeout("tcp", url_, 3*time.Second)
	if err != nil {
		log.Fatalf("can't connect to server %s  :%v\n", url_, err)
	}

	defer conn.Close()

	var (
		respLenField []byte = make([]byte, 2)
		respLen      int
		resp         = make([]byte, 16*1024)
	)
	for {
		if len(responseChan) >= tc {
			break
		}

		timer := NewTimer()
		timer.Reset()

		writeoffset := 0
		for {
			if n, err := conn.Write(requestBody[writeoffset:]); err != nil {
				log.Fatalf("write to serve error : %v  \n", requestBody)
			} else {
				writeoffset += n
			}
			if writeoffset == len(requestBody) {
				break
			}
		}

		//打印发送的消息
		oncePrintSend.Do(func() {
			if *printMsg {
				fmt.Println("==========================Print Msg==========================")
				log.Printf("send: %s \n", hex.EncodeToString(requestBody[2:]))
			}
		})

		//eg 00 0F 3132333435363738 4130 3030
		if n, err := conn.Read(respLenField); err != nil && n != 2 {
			log.Fatalf("read from server error : %v  \n  read %d bytes except %d \n", err, n, 2)
		}
		respLen = bytesTo16Int(respLenField)

		readoffset := 0
		for {
			if n, err := conn.Read(resp[readoffset:]); err != nil {
				log.Fatalf("read from server error : %v  \n read %d bytes except %d \n", err, n, respLen)
			} else {
				readoffset += n
			}

			if readoffset == respLen {
				break
			}
		}

		oncePrintRecv.Do(func() {
			if *printMsg {
				log.Printf("recv: %s \n", hex.EncodeToString(resp[0:respLen]))
			}
		})

		//解析返回
		t1 := timer.Duration()
		respObj := parseRes(resp)
		respObj.Duration = t1
		if *msgSize == -1 {
			respObj.Size = int64(respLen)
		} else {
			respObj.Size = int64(*msgSize)
		}

		responseChan <- respObj
	}

}

func parseRes(tcpRes []byte) *Response {
	respObj := &Response{}

	tcpRes = tcpRes[*headerlen : len(tcpRes)-1]

	respObj.StatusCode = bytesTo16Int(tcpRes[2:3])
	if respObj.StatusCode != 0 {
		respObj.Error = true
	} else {
		respObj.Error = false
		respObj.Body = tcpRes
	}
	return respObj
}
