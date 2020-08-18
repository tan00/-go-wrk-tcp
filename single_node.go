package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"log"
	"sync"
)

func SingleNode(toCall string) []byte {
	var (
		req             []byte
		err             error
		responseChannel = make(chan *Response, *totalCalls*2)
	)

	if *requestBodyHex { //hex
		req, err = hex.DecodeString(*requestBody)
		if err != nil {
			log.Fatalf("requestBody is invalid  err= %v  \n hex=%s\n", err, *requestBody)
		}
	} else {
		req = []byte(*requestBody)
	}

	if *addLenField {
		req = add2ByteLen(req)
	}

	benchTime := NewTimer()
	benchTime.Reset()
	//TODO check ulimit
	wg := &sync.WaitGroup{}

	for i := 0; i < *numConnections; i++ {
		go StartClient(
			toCall,
			req,
			*disableKeepAlives,
			responseChannel,
			wg,
			*totalCalls,
		)
		wg.Add(1)
	}

	wg.Wait()

	result := CalcStats(
		responseChannel,
		benchTime.Duration(),
	)
	return result
}

//为byte添加2字节长度头
func add2ByteLen(in []byte) []byte {
	b1 := intTo2Bytes(len(in))
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write(b1)
	buffer.Write(in)

	return buffer.Bytes()
}

//intTo2Bytes 转换int为2字节byte
func intTo2Bytes(i int) []byte {
	buf := bytes.NewBuffer([]byte{})
	tmp := uint16(i)
	binary.Write(buf, binary.BigEndian, tmp)
	return buf.Bytes()
}

// bytes to int 16
func bytesTo16Int(b []byte) int {
	buf := bytes.NewBuffer(b)
	var tmp uint16
	binary.Read(buf, binary.BigEndian, &tmp)
	return int(tmp)
}
