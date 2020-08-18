all:linux win loogson install

linux:
	GOOS=linux GOARCH=amd64 go build  -o go-wrk-tcp.linux
win:
	GOOS=windows GOARCH=amd64 go build -o go-wrk-tcp.exe

loogson:
	GOOS=linux GOARCH=mips64le go build  -o go-wrk-tcp.loonson

install:
	cp go-wrk-tcp.* hsm_bench
