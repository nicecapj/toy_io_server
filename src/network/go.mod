module network

go 1.17

require github.com/golang/protobuf v1.5.2

require google.golang.org/protobuf v1.26.0 // indirect

require (
	github.com/google/uuid v1.3.0
	packet_protocol v0.0.0
)

replace packet_protocol => ../packet_protocol
