# toy_io_server
 Golang io game server to make fun


**1. Ready for work**  
install golang 1.12.5

**2. install ide(in my case, used vs code and plugin)

**3.download protoc.exe**  
https://repo1.maven.org/maven2/com/google/protobuf/protoc/3.6.1/ 

**4.download protobuff package**  
go get github.com/golang/protobuf/proto

**5. download protoc go version**  
go get github.com/golang/protobuf/protoc-gen-go

**6. make protobuf go file**  
protoc --go_out=. login.proto  
protoc --go_out=. returncode.proto






**for go client
protoc --go_out=. login.proto


**for c++ client
protoc --cpp_out=. login.proto







reference
serialization/deserialization of header : https://golang.org/src/encoding/binary/example_test.go
singleton : http://marcio.io/2015/07/singleton-pattern-in-go/
