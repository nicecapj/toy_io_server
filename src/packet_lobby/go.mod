module packet_lobby

go 1.17

require github.com/golang/protobuf v1.5.2
require packet_returncode v0.0.0

require google.golang.org/protobuf v1.26.0 // indirect

replace packet_returncode => ../packet_returncode
