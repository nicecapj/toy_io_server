# toy_io_server
 Golang io game server to make fun


1. Ready for work
download protoc.exe
https://repo1.maven.org/maven2/com/google/protobuf/protoc/3.6.1/ 

2. 
go get github.com/golang/protobuf/protoc-gen-go

3.
go get github.com/golang/protobuf/proto


4. make protobuf go file
protoc --go_out=. login.proto






for go clien
protoc --go_out=. login.proto


for c++ client
protoc --cpp_out=. login.proto