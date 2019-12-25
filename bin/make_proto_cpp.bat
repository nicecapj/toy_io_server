protoc --cpp_out=../src/packet_returncode/. -I ../src/proto returncode.proto
protoc --cpp_out=../src/packet_lobby/. -I ../src/proto login.proto 
protoc --cpp_out=../src/packet_protocol/. -I ../src/proto protocol.proto