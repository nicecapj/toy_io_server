# toy_io_server
- Try to create a real time gameServer using golang, tcp, protobuf, and concurrency programs.
- First goal: to make a moveable .io game. 

![mainFlow_ver1](https://user-images.githubusercontent.com/8508812/59664539-34767a80-91ec-11e9-9183-657a4e991eaa.jpg)


**0. Simple run**
1. Setting GOPATH to toy_io_server root folder (reference : https://github.com/golang/go/wiki/SettingGOPATH  
2. run server : go run toy_io_server\src\toy_io_server\toy_io_server.go  
3. run test client : go run toy_io_server\src\toy_io_testClient\toy_io_testClient.go  
  
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
protoc --go_out=. returncode.proto  
protoc --go_out=. login.proto 
protoc --go_out=. protocol.proto  
    
**7. copy xxx.pb.go to src/packet_xxx folder**  
copy login.pb.go to src\packet_lobby\
copy protocol.pb.go to src\packet_protocol\
copy returncode.pb.pb.go to src\returncode.pb\
  
**8. Modify and Use**    


**for go client**  
protoc --go_out=. login.proto


**for c++ client**  
protoc --cpp_out=. login.proto







reference
serialization/deserialization of header : https://golang.org/src/encoding/binary/example_test.go
singleton : http://marcio.io/2015/07/singleton-pattern-in-go/
