// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package packet_protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProtocolID int32

const (
	ProtocolID_ProtocolBegin   ProtocolID = 0
	ProtocolID_LoginReq        ProtocolID = 1
	ProtocolID_LoginRes        ProtocolID = 2
	ProtocolID_RoomEnterReq    ProtocolID = 3
	ProtocolID_RoomEnterRes    ProtocolID = 4
	ProtocolID_RoomEnterNfy    ProtocolID = 5
	ProtocolID_ReadyForGameReq ProtocolID = 6
	ProtocolID_RoomLeaveReq    ProtocolID = 7
	ProtocolID_RoomLeaveRes    ProtocolID = 8
	ProtocolID_RoomLeaveNfy    ProtocolID = 9
	ProtocolID_MoveStartReq    ProtocolID = 10
	ProtocolID_MoveStartRes    ProtocolID = 11
	ProtocolID_MoveStartNfy    ProtocolID = 12
	ProtocolID_MoveChangeReq   ProtocolID = 13
	ProtocolID_MoveChangeRes   ProtocolID = 14
	ProtocolID_MoveChangeNfy   ProtocolID = 15
	ProtocolID_MoveEndReq      ProtocolID = 16
	ProtocolID_MoveEndRes      ProtocolID = 17
	ProtocolID_MoveEndNfy      ProtocolID = 18
	ProtocolID_AttackReq       ProtocolID = 19
	ProtocolID_AttackRes       ProtocolID = 20
	ProtocolID_AttackNfy       ProtocolID = 21
	ProtocolID_FisingStartReq  ProtocolID = 22
	ProtocolID_FisingStartRes  ProtocolID = 23
	ProtocolID_FisingStartNfy  ProtocolID = 24
	ProtocolID_FisingHitNfy    ProtocolID = 25
)

var ProtocolID_name = map[int32]string{
	0:  "ProtocolBegin",
	1:  "LoginReq",
	2:  "LoginRes",
	3:  "RoomEnterReq",
	4:  "RoomEnterRes",
	5:  "RoomEnterNfy",
	6:  "ReadyForGameReq",
	7:  "RoomLeaveReq",
	8:  "RoomLeaveRes",
	9:  "RoomLeaveNfy",
	10: "MoveStartReq",
	11: "MoveStartRes",
	12: "MoveStartNfy",
	13: "MoveChangeReq",
	14: "MoveChangeRes",
	15: "MoveChangeNfy",
	16: "MoveEndReq",
	17: "MoveEndRes",
	18: "MoveEndNfy",
	19: "AttackReq",
	20: "AttackRes",
	21: "AttackNfy",
	22: "FisingStartReq",
	23: "FisingStartRes",
	24: "FisingStartNfy",
	25: "FisingHitNfy",
}

var ProtocolID_value = map[string]int32{
	"ProtocolBegin":   0,
	"LoginReq":        1,
	"LoginRes":        2,
	"RoomEnterReq":    3,
	"RoomEnterRes":    4,
	"RoomEnterNfy":    5,
	"ReadyForGameReq": 6,
	"RoomLeaveReq":    7,
	"RoomLeaveRes":    8,
	"RoomLeaveNfy":    9,
	"MoveStartReq":    10,
	"MoveStartRes":    11,
	"MoveStartNfy":    12,
	"MoveChangeReq":   13,
	"MoveChangeRes":   14,
	"MoveChangeNfy":   15,
	"MoveEndReq":      16,
	"MoveEndRes":      17,
	"MoveEndNfy":      18,
	"AttackReq":       19,
	"AttackRes":       20,
	"AttackNfy":       21,
	"FisingStartReq":  22,
	"FisingStartRes":  23,
	"FisingStartNfy":  24,
	"FisingHitNfy":    25,
}

func (x ProtocolID) String() string {
	return proto.EnumName(ProtocolID_name, int32(x))
}

func (ProtocolID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

func init() {
	proto.RegisterEnum("packet.protocol.ProtocolID", ProtocolID_name, ProtocolID_value)
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd1, 0x4b, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0xa0, 0x50, 0xda, 0x21, 0x8f, 0xe9, 0x94, 0xe7, 0x15, 0x58, 0x74, 0xc3, 0x09,
	0x78, 0xb4, 0x80, 0x54, 0x10, 0x0a, 0x27, 0x30, 0xa9, 0x09, 0x56, 0xa9, 0x5d, 0x62, 0xab, 0x52,
	0x2e, 0xca, 0x79, 0xd0, 0x18, 0x02, 0xb1, 0xd9, 0xe5, 0xff, 0x32, 0xbf, 0xad, 0x91, 0x21, 0x5b,
	0xd7, 0xc6, 0x99, 0xd2, 0xbc, 0x4f, 0xfc, 0x07, 0xe5, 0x6b, 0x51, 0x2e, 0xa5, 0x9b, 0xb4, 0x7c,
	0xfe, 0xd9, 0x03, 0x78, 0xfa, 0x09, 0xf7, 0x37, 0x34, 0x82, 0xb4, 0x4d, 0x57, 0xb2, 0x52, 0x1a,
	0xb7, 0x28, 0x81, 0xc1, 0xdc, 0x54, 0x4a, 0x17, 0xf2, 0x03, 0xb7, 0x3b, 0xc9, 0xe2, 0x0e, 0x21,
	0x24, 0x85, 0x31, 0xab, 0xa9, 0x76, 0xb2, 0xe6, 0xff, 0xbd, 0x48, 0x2c, 0xee, 0x06, 0xf2, 0xf8,
	0xda, 0xe0, 0x1e, 0x8d, 0x21, 0x2f, 0xa4, 0x58, 0x34, 0x33, 0x53, 0xdf, 0x8a, 0x95, 0xe4, 0x62,
	0xbf, 0x1d, 0x9b, 0x4b, 0xb1, 0xf1, 0xb2, 0x1f, 0x89, 0xc5, 0x41, 0x20, 0x7c, 0xd4, 0x90, 0xe5,
	0xc1, 0x6c, 0xe4, 0xb3, 0x13, 0xb5, 0xe3, 0x16, 0x44, 0x62, 0xf1, 0x20, 0x10, 0x6e, 0x25, 0xbc,
	0x25, 0xcb, 0xf5, 0x9b, 0xd0, 0x95, 0xbf, 0x2c, 0x8d, 0xc9, 0x62, 0x16, 0x12, 0x17, 0x73, 0xca,
	0x00, 0x98, 0xa6, 0x7a, 0xc1, 0x2d, 0x0c, 0xb2, 0xc5, 0x51, 0x27, 0xf3, 0x3c, 0x51, 0x0a, 0xc3,
	0x4b, 0xe7, 0x44, 0xb9, 0xe4, 0xf1, 0x71, 0x37, 0x5a, 0x3c, 0xfc, 0x8b, 0x3c, 0x7c, 0x44, 0x04,
	0xd9, 0x4c, 0x59, 0xa5, 0xab, 0xdf, 0x6d, 0x8e, 0xff, 0x99, 0xc5, 0x93, 0xc8, 0xb8, 0x7b, 0xca,
	0x3b, 0x7e, 0xdb, 0x9d, 0xf2, 0x72, 0xf6, 0xd2, 0xf7, 0x4f, 0x7c, 0xf1, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x04, 0x9d, 0x9f, 0x12, 0x02, 0x02, 0x00, 0x00,
}
