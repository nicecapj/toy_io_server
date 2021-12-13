module src/toy_io_server

go 1.17

require (
    github.com/golang/protobuf v1.5.2
    google.golang.org/protobuf v1.26.0 // indirect
	network v0.0.0
	packet_lobby v0.0.0
	packet_protocol v0.0.0
	packet_returncode v0.0.0
	util v0.0.0
)

require google.golang.org/protobuf v1.26.0 // indirect

replace util => ../util

replace network => ../network

replace packet_lobby => ../packet_lobby

replace packet_protocol => ../packet_protocol

replace packet_returncode => ../packet_returncode