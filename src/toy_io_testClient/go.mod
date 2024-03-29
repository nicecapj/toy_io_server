module main

go 1.17

require (
	github.com/golang/protobuf v1.5.2
	network v0.0.0-00010101000000-000000000000
	packet_lobby v0.0.0
	packet_protocol v0.0.0
	packet_returncode v0.0.0
	util v0.0.0
)

require (
	github.com/google/uuid v1.3.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)

replace util => ../util

replace network => ../network

replace packet_lobby => ../packet_lobby

replace packet_protocol => ../packet_protocol

replace packet_returncode => ../packet_returncode
