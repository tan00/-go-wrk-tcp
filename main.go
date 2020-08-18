package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

type Config struct {
	Port  string
	Nodes []string
}

var (
	numThreads        = flag.Int("t", 1, "the numbers of threads used")
	requestBody       = flag.String("b", "", "the request body")
	requestBodyFile   = flag.String("p", "", "the request body data file")
	requestBodyHex    = flag.Bool("hex", false, "the request body using hex coding")
	addLenField       = flag.Bool("addlen", true, "auto add 2 byte field  to index body length")
	headerlen         = flag.Int("headerlen", 0, "n byte msg header")
	numConnections    = flag.Int("c", 1, "the max numbers of connections used")
	totalCalls        = flag.Int("n", 1, "the total number of calls processed")
	disableKeepAlives = flag.Bool("k", true, "if keep-alives are disabled")
	certFile          = flag.String("cert", "someCertFile", "A PEM eoncoded certificate file.")
	keyFile           = flag.String("key", "someKeyFile", "A PEM encoded private key file.")
	caFile            = flag.String("CA", "someCertCAFile", "A PEM eoncoded CA's certificate file.")
	insecure          = flag.Bool("i", false, "TLS checks are disabled")
	printMsg          = flag.Bool("print", false, "print all in and out msg")
	msgSize           = flag.Int("size", -1, "msg size to calc Mbps")
	target            string
)

func init() {
	flag.Parse()
	target = os.Args[len(os.Args)-1]
	runtime.GOMAXPROCS(*numThreads) //一个gorotine 一个连接， 设置线程最大数量是为了限制同时运行的gorotine
}

func setRequestBody() {
	// requestBody has been setup and it has highest priority
	if *requestBody != "" {
		return
	}

	if *requestBodyFile == "" {
		return
	}

	// requestBodyFile has been setup
	data, err := ioutil.ReadFile(*requestBodyFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	body := string(data)
	body = strings.TrimSpace(body)
	requestBody = &body
}

func main() {
	if len(os.Args) == 1 { //If no argument specified
		flag.Usage() //Print usage
		os.Exit(1)
	}
	setRequestBody()
	SingleNode(target)

}
