

serve: 
	go-bindata -pkg=main -prefix=data -o=bindata.go data/...
	go run *.go
build:
	go-bindata -pkg=main -prefix=data -o=bindata.go data/...
	go build -o main *.go
