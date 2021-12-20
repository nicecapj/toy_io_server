module main

go 1.17

replace util => ../util

replace network => ../network

replace packet_lobby => ../packet_lobby

replace packet_protocol => ../packet_protocol

replace packet_returncode => ../packet_returncode

require network v0.0.0-00010101000000-000000000000

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	packet_protocol v0.0.0 // indirect
)
