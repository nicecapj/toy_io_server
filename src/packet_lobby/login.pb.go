// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

package packet_lobby

import (
	fmt "fmt"
	math "math"
	. "packet_returncode"

	proto "github.com/golang/protobuf/proto"
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

type LoginReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LoginRes struct {
	RetCode              ReturnCode `protobuf:"varint,1,opt,name=retCode,proto3,enum=packet.returncode.ReturnCode" json:"retCode,omitempty"`
	Uid                  int64      `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LoginRes) Reset()         { *m = LoginRes{} }
func (m *LoginRes) String() string { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()    {}
func (*LoginRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
}

func (m *LoginRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRes.Unmarshal(m, b)
}
func (m *LoginRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRes.Marshal(b, m, deterministic)
}
func (m *LoginRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRes.Merge(m, src)
}
func (m *LoginRes) XXX_Size() int {
	return xxx_messageInfo_LoginRes.Size(m)
}
func (m *LoginRes) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRes.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRes proto.InternalMessageInfo

func (m *LoginRes) GetRetCode() ReturnCode {
	if m != nil {
		return m.RetCode
	}
	return ReturnCode_retOK
}

func (m *LoginRes) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *LoginRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Location struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    int32    `protobuf:"varint,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{2}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Location) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Location) GetZ() int32 {
	if m != nil {
		return m.Z
	}
	return 0
}

type UserInfo struct {
	Uid                  int64     `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location             *Location `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{3}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type RoomEnterReq struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomEnterReq) Reset()         { *m = RoomEnterReq{} }
func (m *RoomEnterReq) String() string { return proto.CompactTextString(m) }
func (*RoomEnterReq) ProtoMessage()    {}
func (*RoomEnterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{4}
}

func (m *RoomEnterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomEnterReq.Unmarshal(m, b)
}
func (m *RoomEnterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomEnterReq.Marshal(b, m, deterministic)
}
func (m *RoomEnterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomEnterReq.Merge(m, src)
}
func (m *RoomEnterReq) XXX_Size() int {
	return xxx_messageInfo_RoomEnterReq.Size(m)
}
func (m *RoomEnterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomEnterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoomEnterReq proto.InternalMessageInfo

func (m *RoomEnterReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type RoomEnterRes struct {
	RetCode              ReturnCode `protobuf:"varint,1,opt,name=retCode,proto3,enum=packet.returncode.ReturnCode" json:"retCode,omitempty"`
	RoomID               int32      `protobuf:"varint,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	RoomName             string     `protobuf:"bytes,3,opt,name=roomName,proto3" json:"roomName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RoomEnterRes) Reset()         { *m = RoomEnterRes{} }
func (m *RoomEnterRes) String() string { return proto.CompactTextString(m) }
func (*RoomEnterRes) ProtoMessage()    {}
func (*RoomEnterRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{5}
}

func (m *RoomEnterRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomEnterRes.Unmarshal(m, b)
}
func (m *RoomEnterRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomEnterRes.Marshal(b, m, deterministic)
}
func (m *RoomEnterRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomEnterRes.Merge(m, src)
}
func (m *RoomEnterRes) XXX_Size() int {
	return xxx_messageInfo_RoomEnterRes.Size(m)
}
func (m *RoomEnterRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomEnterRes.DiscardUnknown(m)
}

var xxx_messageInfo_RoomEnterRes proto.InternalMessageInfo

func (m *RoomEnterRes) GetRetCode() ReturnCode {
	if m != nil {
		return m.RetCode
	}
	return ReturnCode_retOK
}

func (m *RoomEnterRes) GetRoomID() int32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *RoomEnterRes) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

type RoomEnterNfy struct {
	UserInfoList         []*UserInfo `protobuf:"bytes,1,rep,name=UserInfoList,proto3" json:"UserInfoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RoomEnterNfy) Reset()         { *m = RoomEnterNfy{} }
func (m *RoomEnterNfy) String() string { return proto.CompactTextString(m) }
func (*RoomEnterNfy) ProtoMessage()    {}
func (*RoomEnterNfy) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{6}
}

func (m *RoomEnterNfy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomEnterNfy.Unmarshal(m, b)
}
func (m *RoomEnterNfy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomEnterNfy.Marshal(b, m, deterministic)
}
func (m *RoomEnterNfy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomEnterNfy.Merge(m, src)
}
func (m *RoomEnterNfy) XXX_Size() int {
	return xxx_messageInfo_RoomEnterNfy.Size(m)
}
func (m *RoomEnterNfy) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomEnterNfy.DiscardUnknown(m)
}

var xxx_messageInfo_RoomEnterNfy proto.InternalMessageInfo

func (m *RoomEnterNfy) GetUserInfoList() []*UserInfo {
	if m != nil {
		return m.UserInfoList
	}
	return nil
}

type ReadyForGameReq struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadyForGameReq) Reset()         { *m = ReadyForGameReq{} }
func (m *ReadyForGameReq) String() string { return proto.CompactTextString(m) }
func (*ReadyForGameReq) ProtoMessage()    {}
func (*ReadyForGameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{7}
}

func (m *ReadyForGameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadyForGameReq.Unmarshal(m, b)
}
func (m *ReadyForGameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadyForGameReq.Marshal(b, m, deterministic)
}
func (m *ReadyForGameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadyForGameReq.Merge(m, src)
}
func (m *ReadyForGameReq) XXX_Size() int {
	return xxx_messageInfo_ReadyForGameReq.Size(m)
}
func (m *ReadyForGameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadyForGameReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadyForGameReq proto.InternalMessageInfo

func (m *ReadyForGameReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type RoomLeaveReq struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	RoomID               int32    `protobuf:"varint,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomLeaveReq) Reset()         { *m = RoomLeaveReq{} }
func (m *RoomLeaveReq) String() string { return proto.CompactTextString(m) }
func (*RoomLeaveReq) ProtoMessage()    {}
func (*RoomLeaveReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{8}
}

func (m *RoomLeaveReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomLeaveReq.Unmarshal(m, b)
}
func (m *RoomLeaveReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomLeaveReq.Marshal(b, m, deterministic)
}
func (m *RoomLeaveReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomLeaveReq.Merge(m, src)
}
func (m *RoomLeaveReq) XXX_Size() int {
	return xxx_messageInfo_RoomLeaveReq.Size(m)
}
func (m *RoomLeaveReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomLeaveReq.DiscardUnknown(m)
}

var xxx_messageInfo_RoomLeaveReq proto.InternalMessageInfo

func (m *RoomLeaveReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *RoomLeaveReq) GetRoomID() int32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type RoomLeaveRes struct {
	RetCode              ReturnCode `protobuf:"varint,1,opt,name=retCode,proto3,enum=packet.returncode.ReturnCode" json:"retCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RoomLeaveRes) Reset()         { *m = RoomLeaveRes{} }
func (m *RoomLeaveRes) String() string { return proto.CompactTextString(m) }
func (*RoomLeaveRes) ProtoMessage()    {}
func (*RoomLeaveRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{9}
}

func (m *RoomLeaveRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomLeaveRes.Unmarshal(m, b)
}
func (m *RoomLeaveRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomLeaveRes.Marshal(b, m, deterministic)
}
func (m *RoomLeaveRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomLeaveRes.Merge(m, src)
}
func (m *RoomLeaveRes) XXX_Size() int {
	return xxx_messageInfo_RoomLeaveRes.Size(m)
}
func (m *RoomLeaveRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomLeaveRes.DiscardUnknown(m)
}

var xxx_messageInfo_RoomLeaveRes proto.InternalMessageInfo

func (m *RoomLeaveRes) GetRetCode() ReturnCode {
	if m != nil {
		return m.RetCode
	}
	return ReturnCode_retOK
}

type RoomLeaveNfy struct {
	UserInfoList         []*UserInfo `protobuf:"bytes,1,rep,name=UserInfoList,proto3" json:"UserInfoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RoomLeaveNfy) Reset()         { *m = RoomLeaveNfy{} }
func (m *RoomLeaveNfy) String() string { return proto.CompactTextString(m) }
func (*RoomLeaveNfy) ProtoMessage()    {}
func (*RoomLeaveNfy) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{10}
}

func (m *RoomLeaveNfy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomLeaveNfy.Unmarshal(m, b)
}
func (m *RoomLeaveNfy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomLeaveNfy.Marshal(b, m, deterministic)
}
func (m *RoomLeaveNfy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomLeaveNfy.Merge(m, src)
}
func (m *RoomLeaveNfy) XXX_Size() int {
	return xxx_messageInfo_RoomLeaveNfy.Size(m)
}
func (m *RoomLeaveNfy) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomLeaveNfy.DiscardUnknown(m)
}

var xxx_messageInfo_RoomLeaveNfy proto.InternalMessageInfo

func (m *RoomLeaveNfy) GetUserInfoList() []*UserInfo {
	if m != nil {
		return m.UserInfoList
	}
	return nil
}

type MoveStartReq struct {
	Uid                  int64     `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	CurrentPos           *Location `protobuf:"bytes,2,opt,name=currentPos,proto3" json:"currentPos,omitempty"`
	TargetPos            *Location `protobuf:"bytes,3,opt,name=targetPos,proto3" json:"targetPos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MoveStartReq) Reset()         { *m = MoveStartReq{} }
func (m *MoveStartReq) String() string { return proto.CompactTextString(m) }
func (*MoveStartReq) ProtoMessage()    {}
func (*MoveStartReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{11}
}

func (m *MoveStartReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveStartReq.Unmarshal(m, b)
}
func (m *MoveStartReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveStartReq.Marshal(b, m, deterministic)
}
func (m *MoveStartReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveStartReq.Merge(m, src)
}
func (m *MoveStartReq) XXX_Size() int {
	return xxx_messageInfo_MoveStartReq.Size(m)
}
func (m *MoveStartReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveStartReq.DiscardUnknown(m)
}

var xxx_messageInfo_MoveStartReq proto.InternalMessageInfo

func (m *MoveStartReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *MoveStartReq) GetCurrentPos() *Location {
	if m != nil {
		return m.CurrentPos
	}
	return nil
}

func (m *MoveStartReq) GetTargetPos() *Location {
	if m != nil {
		return m.TargetPos
	}
	return nil
}

type MoveStartRes struct {
	RetCode              ReturnCode `protobuf:"varint,1,opt,name=retCode,proto3,enum=packet.returncode.ReturnCode" json:"retCode,omitempty"`
	TargetPos            *Location  `protobuf:"bytes,2,opt,name=targetPos,proto3" json:"targetPos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MoveStartRes) Reset()         { *m = MoveStartRes{} }
func (m *MoveStartRes) String() string { return proto.CompactTextString(m) }
func (*MoveStartRes) ProtoMessage()    {}
func (*MoveStartRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{12}
}

func (m *MoveStartRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveStartRes.Unmarshal(m, b)
}
func (m *MoveStartRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveStartRes.Marshal(b, m, deterministic)
}
func (m *MoveStartRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveStartRes.Merge(m, src)
}
func (m *MoveStartRes) XXX_Size() int {
	return xxx_messageInfo_MoveStartRes.Size(m)
}
func (m *MoveStartRes) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveStartRes.DiscardUnknown(m)
}

var xxx_messageInfo_MoveStartRes proto.InternalMessageInfo

func (m *MoveStartRes) GetRetCode() ReturnCode {
	if m != nil {
		return m.RetCode
	}
	return ReturnCode_retOK
}

func (m *MoveStartRes) GetTargetPos() *Location {
	if m != nil {
		return m.TargetPos
	}
	return nil
}

type MoveStartNfy struct {
	Uid                  int64     `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	CurrentPos           *Location `protobuf:"bytes,2,opt,name=currentPos,proto3" json:"currentPos,omitempty"`
	TargetPos            *Location `protobuf:"bytes,3,opt,name=targetPos,proto3" json:"targetPos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MoveStartNfy) Reset()         { *m = MoveStartNfy{} }
func (m *MoveStartNfy) String() string { return proto.CompactTextString(m) }
func (*MoveStartNfy) ProtoMessage()    {}
func (*MoveStartNfy) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{13}
}

func (m *MoveStartNfy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveStartNfy.Unmarshal(m, b)
}
func (m *MoveStartNfy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveStartNfy.Marshal(b, m, deterministic)
}
func (m *MoveStartNfy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveStartNfy.Merge(m, src)
}
func (m *MoveStartNfy) XXX_Size() int {
	return xxx_messageInfo_MoveStartNfy.Size(m)
}
func (m *MoveStartNfy) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveStartNfy.DiscardUnknown(m)
}

var xxx_messageInfo_MoveStartNfy proto.InternalMessageInfo

func (m *MoveStartNfy) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *MoveStartNfy) GetCurrentPos() *Location {
	if m != nil {
		return m.CurrentPos
	}
	return nil
}

func (m *MoveStartNfy) GetTargetPos() *Location {
	if m != nil {
		return m.TargetPos
	}
	return nil
}

type MoveChangeReq struct {
	Uid                  int64     `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	CurrentPos           *Location `protobuf:"bytes,2,opt,name=currentPos,proto3" json:"currentPos,omitempty"`
	TargetPos            *Location `protobuf:"bytes,3,opt,name=targetPos,proto3" json:"targetPos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MoveChangeReq) Reset()         { *m = MoveChangeReq{} }
func (m *MoveChangeReq) String() string { return proto.CompactTextString(m) }
func (*MoveChangeReq) ProtoMessage()    {}
func (*MoveChangeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{14}
}

func (m *MoveChangeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveChangeReq.Unmarshal(m, b)
}
func (m *MoveChangeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveChangeReq.Marshal(b, m, deterministic)
}
func (m *MoveChangeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveChangeReq.Merge(m, src)
}
func (m *MoveChangeReq) XXX_Size() int {
	return xxx_messageInfo_MoveChangeReq.Size(m)
}
func (m *MoveChangeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveChangeReq.DiscardUnknown(m)
}

var xxx_messageInfo_MoveChangeReq proto.InternalMessageInfo

func (m *MoveChangeReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *MoveChangeReq) GetCurrentPos() *Location {
	if m != nil {
		return m.CurrentPos
	}
	return nil
}

func (m *MoveChangeReq) GetTargetPos() *Location {
	if m != nil {
		return m.TargetPos
	}
	return nil
}

type MoveChangeRes struct {
	RetCode              ReturnCode `protobuf:"varint,1,opt,name=retCode,proto3,enum=packet.returncode.ReturnCode" json:"retCode,omitempty"`
	TargetPos            *Location  `protobuf:"bytes,2,opt,name=targetPos,proto3" json:"targetPos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MoveChangeRes) Reset()         { *m = MoveChangeRes{} }
func (m *MoveChangeRes) String() string { return proto.CompactTextString(m) }
func (*MoveChangeRes) ProtoMessage()    {}
func (*MoveChangeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{15}
}

func (m *MoveChangeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveChangeRes.Unmarshal(m, b)
}
func (m *MoveChangeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveChangeRes.Marshal(b, m, deterministic)
}
func (m *MoveChangeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveChangeRes.Merge(m, src)
}
func (m *MoveChangeRes) XXX_Size() int {
	return xxx_messageInfo_MoveChangeRes.Size(m)
}
func (m *MoveChangeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveChangeRes.DiscardUnknown(m)
}

var xxx_messageInfo_MoveChangeRes proto.InternalMessageInfo

func (m *MoveChangeRes) GetRetCode() ReturnCode {
	if m != nil {
		return m.RetCode
	}
	return ReturnCode_retOK
}

func (m *MoveChangeRes) GetTargetPos() *Location {
	if m != nil {
		return m.TargetPos
	}
	return nil
}

type MoveChangeNfy struct {
	Uid                  int64     `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	TargetPos            *Location `protobuf:"bytes,2,opt,name=targetPos,proto3" json:"targetPos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MoveChangeNfy) Reset()         { *m = MoveChangeNfy{} }
func (m *MoveChangeNfy) String() string { return proto.CompactTextString(m) }
func (*MoveChangeNfy) ProtoMessage()    {}
func (*MoveChangeNfy) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{16}
}

func (m *MoveChangeNfy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveChangeNfy.Unmarshal(m, b)
}
func (m *MoveChangeNfy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveChangeNfy.Marshal(b, m, deterministic)
}
func (m *MoveChangeNfy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveChangeNfy.Merge(m, src)
}
func (m *MoveChangeNfy) XXX_Size() int {
	return xxx_messageInfo_MoveChangeNfy.Size(m)
}
func (m *MoveChangeNfy) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveChangeNfy.DiscardUnknown(m)
}

var xxx_messageInfo_MoveChangeNfy proto.InternalMessageInfo

func (m *MoveChangeNfy) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *MoveChangeNfy) GetTargetPos() *Location {
	if m != nil {
		return m.TargetPos
	}
	return nil
}

type MoveEndReq struct {
	Uid                  int64     `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	TargetPos            *Location `protobuf:"bytes,2,opt,name=targetPos,proto3" json:"targetPos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MoveEndReq) Reset()         { *m = MoveEndReq{} }
func (m *MoveEndReq) String() string { return proto.CompactTextString(m) }
func (*MoveEndReq) ProtoMessage()    {}
func (*MoveEndReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{17}
}

func (m *MoveEndReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveEndReq.Unmarshal(m, b)
}
func (m *MoveEndReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveEndReq.Marshal(b, m, deterministic)
}
func (m *MoveEndReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveEndReq.Merge(m, src)
}
func (m *MoveEndReq) XXX_Size() int {
	return xxx_messageInfo_MoveEndReq.Size(m)
}
func (m *MoveEndReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveEndReq.DiscardUnknown(m)
}

var xxx_messageInfo_MoveEndReq proto.InternalMessageInfo

func (m *MoveEndReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *MoveEndReq) GetTargetPos() *Location {
	if m != nil {
		return m.TargetPos
	}
	return nil
}

type MoveEndRes struct {
	RetCode              ReturnCode `protobuf:"varint,1,opt,name=retCode,proto3,enum=packet.returncode.ReturnCode" json:"retCode,omitempty"`
	EndPos               *Location  `protobuf:"bytes,2,opt,name=endPos,proto3" json:"endPos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MoveEndRes) Reset()         { *m = MoveEndRes{} }
func (m *MoveEndRes) String() string { return proto.CompactTextString(m) }
func (*MoveEndRes) ProtoMessage()    {}
func (*MoveEndRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{18}
}

func (m *MoveEndRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveEndRes.Unmarshal(m, b)
}
func (m *MoveEndRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveEndRes.Marshal(b, m, deterministic)
}
func (m *MoveEndRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveEndRes.Merge(m, src)
}
func (m *MoveEndRes) XXX_Size() int {
	return xxx_messageInfo_MoveEndRes.Size(m)
}
func (m *MoveEndRes) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveEndRes.DiscardUnknown(m)
}

var xxx_messageInfo_MoveEndRes proto.InternalMessageInfo

func (m *MoveEndRes) GetRetCode() ReturnCode {
	if m != nil {
		return m.RetCode
	}
	return ReturnCode_retOK
}

func (m *MoveEndRes) GetEndPos() *Location {
	if m != nil {
		return m.EndPos
	}
	return nil
}

type MoveEndNfy struct {
	StartPos             *Location `protobuf:"bytes,1,opt,name=startPos,proto3" json:"startPos,omitempty"`
	EndPos               *Location `protobuf:"bytes,2,opt,name=endPos,proto3" json:"endPos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MoveEndNfy) Reset()         { *m = MoveEndNfy{} }
func (m *MoveEndNfy) String() string { return proto.CompactTextString(m) }
func (*MoveEndNfy) ProtoMessage()    {}
func (*MoveEndNfy) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{19}
}

func (m *MoveEndNfy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveEndNfy.Unmarshal(m, b)
}
func (m *MoveEndNfy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveEndNfy.Marshal(b, m, deterministic)
}
func (m *MoveEndNfy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveEndNfy.Merge(m, src)
}
func (m *MoveEndNfy) XXX_Size() int {
	return xxx_messageInfo_MoveEndNfy.Size(m)
}
func (m *MoveEndNfy) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveEndNfy.DiscardUnknown(m)
}

var xxx_messageInfo_MoveEndNfy proto.InternalMessageInfo

func (m *MoveEndNfy) GetStartPos() *Location {
	if m != nil {
		return m.StartPos
	}
	return nil
}

func (m *MoveEndNfy) GetEndPos() *Location {
	if m != nil {
		return m.EndPos
	}
	return nil
}

type AttackTo struct {
	TargetUid            int64    `protobuf:"varint,1,opt,name=targetUid,proto3" json:"targetUid,omitempty"`
	SkillID              int32    `protobuf:"varint,2,opt,name=skillID,proto3" json:"skillID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttackTo) Reset()         { *m = AttackTo{} }
func (m *AttackTo) String() string { return proto.CompactTextString(m) }
func (*AttackTo) ProtoMessage()    {}
func (*AttackTo) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{20}
}

func (m *AttackTo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttackTo.Unmarshal(m, b)
}
func (m *AttackTo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttackTo.Marshal(b, m, deterministic)
}
func (m *AttackTo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttackTo.Merge(m, src)
}
func (m *AttackTo) XXX_Size() int {
	return xxx_messageInfo_AttackTo.Size(m)
}
func (m *AttackTo) XXX_DiscardUnknown() {
	xxx_messageInfo_AttackTo.DiscardUnknown(m)
}

var xxx_messageInfo_AttackTo proto.InternalMessageInfo

func (m *AttackTo) GetTargetUid() int64 {
	if m != nil {
		return m.TargetUid
	}
	return 0
}

func (m *AttackTo) GetSkillID() int32 {
	if m != nil {
		return m.SkillID
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "packet.lobby.LoginReq")
	proto.RegisterType((*LoginRes)(nil), "packet.lobby.LoginRes")
	proto.RegisterType((*Location)(nil), "packet.lobby.Location")
	proto.RegisterType((*UserInfo)(nil), "packet.lobby.UserInfo")
	proto.RegisterType((*RoomEnterReq)(nil), "packet.lobby.RoomEnterReq")
	proto.RegisterType((*RoomEnterRes)(nil), "packet.lobby.RoomEnterRes")
	proto.RegisterType((*RoomEnterNfy)(nil), "packet.lobby.RoomEnterNfy")
	proto.RegisterType((*ReadyForGameReq)(nil), "packet.lobby.ReadyForGameReq")
	proto.RegisterType((*RoomLeaveReq)(nil), "packet.lobby.RoomLeaveReq")
	proto.RegisterType((*RoomLeaveRes)(nil), "packet.lobby.RoomLeaveRes")
	proto.RegisterType((*RoomLeaveNfy)(nil), "packet.lobby.RoomLeaveNfy")
	proto.RegisterType((*MoveStartReq)(nil), "packet.lobby.MoveStartReq")
	proto.RegisterType((*MoveStartRes)(nil), "packet.lobby.MoveStartRes")
	proto.RegisterType((*MoveStartNfy)(nil), "packet.lobby.MoveStartNfy")
	proto.RegisterType((*MoveChangeReq)(nil), "packet.lobby.MoveChangeReq")
	proto.RegisterType((*MoveChangeRes)(nil), "packet.lobby.MoveChangeRes")
	proto.RegisterType((*MoveChangeNfy)(nil), "packet.lobby.MoveChangeNfy")
	proto.RegisterType((*MoveEndReq)(nil), "packet.lobby.MoveEndReq")
	proto.RegisterType((*MoveEndRes)(nil), "packet.lobby.MoveEndRes")
	proto.RegisterType((*MoveEndNfy)(nil), "packet.lobby.MoveEndNfy")
	proto.RegisterType((*AttackTo)(nil), "packet.lobby.AttackTo")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0xe5, 0x96, 0x76, 0xd9, 0x5b, 0x80, 0xc9, 0x87, 0xa9, 0x9a, 0x00, 0x55, 0xe6, 0xb2,
	0x53, 0x0e, 0x65, 0x02, 0xc4, 0x0d, 0xc6, 0x98, 0x86, 0xca, 0x84, 0xcc, 0x26, 0xce, 0x6e, 0xe2,
	0x95, 0xa8, 0xa9, 0x5d, 0x1c, 0x77, 0x5a, 0x86, 0xe0, 0x4f, 0x80, 0x7f, 0x19, 0xd9, 0x8d, 0xd3,
	0x44, 0xcd, 0x0a, 0x34, 0x12, 0x70, 0xf3, 0x6b, 0xbe, 0xf7, 0x7e, 0xfe, 0x3e, 0x5b, 0x2e, 0xec,
	0x24, 0x72, 0x1c, 0x8b, 0x60, 0xa6, 0xa4, 0x96, 0xd8, 0x9f, 0xb1, 0x70, 0xc2, 0x75, 0x90, 0xc8,
	0xd1, 0x28, 0xdb, 0xdf, 0x55, 0x5c, 0xcf, 0x95, 0x08, 0x65, 0xc4, 0x17, 0xdf, 0xc9, 0x23, 0xf0,
	0x86, 0x46, 0x4e, 0xf9, 0x67, 0x8c, 0xe1, 0x8e, 0x60, 0x53, 0xde, 0x43, 0x7d, 0x74, 0xb0, 0x4d,
	0xed, 0x9a, 0xc4, 0xc5, 0xf7, 0x14, 0x3f, 0x83, 0x2d, 0xc5, 0xf5, 0x91, 0x8c, 0x16, 0x92, 0x7b,
	0x83, 0x87, 0x41, 0x3e, 0xbd, 0x34, 0x96, 0xda, 0xa5, 0x11, 0x51, 0xa7, 0xc6, 0xbb, 0xd0, 0x9e,
	0xc7, 0x51, 0xaf, 0xd5, 0x47, 0x07, 0x6d, 0x6a, 0x96, 0x05, 0xaa, 0x5d, 0x42, 0x1d, 0x1a, 0x54,
	0xc8, 0x74, 0x2c, 0x05, 0xf6, 0x01, 0x5d, 0x5b, 0x48, 0x87, 0xa2, 0x6b, 0x53, 0x65, 0xb6, 0xbb,
	0x43, 0x51, 0x66, 0xaa, 0x1b, 0xdb, 0xd8, 0xa1, 0xe8, 0x86, 0x44, 0xe0, 0x5d, 0xa4, 0x5c, 0x9d,
	0x8a, 0x4b, 0xe9, 0x38, 0x68, 0x95, 0xd3, 0x5a, 0x72, 0xf0, 0x00, 0xbc, 0x24, 0xe7, 0xd8, 0x31,
	0x3b, 0x83, 0xbd, 0xa0, 0x9c, 0x52, 0xe0, 0x76, 0x41, 0x0b, 0x1d, 0xe9, 0x83, 0x4f, 0xa5, 0x9c,
	0x1e, 0x0b, 0xcd, 0x95, 0x89, 0x6a, 0x85, 0x44, 0xbe, 0x54, 0x14, 0x0d, 0xc2, 0xda, 0x83, 0xae,
	0x92, 0x72, 0x7a, 0xfa, 0x3a, 0x77, 0x9c, 0x57, 0x78, 0x1f, 0x3c, 0xb3, 0x3a, 0x5b, 0xc6, 0x56,
	0xd4, 0xe4, 0x6d, 0x09, 0x7e, 0x76, 0x99, 0xe1, 0x17, 0xe0, 0xbb, 0x50, 0x86, 0x71, 0xaa, 0x7b,
	0xa8, 0xdf, 0x5e, 0xb5, 0xe9, 0x14, 0xb4, 0xa2, 0x25, 0x8f, 0xe1, 0x3e, 0xe5, 0x2c, 0xca, 0xde,
	0x48, 0x75, 0xc2, 0xa6, 0xbc, 0xde, 0xed, 0xf3, 0x05, 0x70, 0xc8, 0xd9, 0x55, 0xbd, 0xe2, 0x36,
	0x1b, 0xe4, 0xa4, 0xd2, 0xb9, 0x79, 0x4e, 0xce, 0xb3, 0x1d, 0xd4, 0xd4, 0xf3, 0x77, 0x04, 0xfe,
	0x3b, 0x79, 0xc5, 0x3f, 0x68, 0xa6, 0x74, 0xbd, 0x9f, 0xa7, 0x00, 0xe1, 0x5c, 0x29, 0x2e, 0xf4,
	0x7b, 0x99, 0x5a, 0x4f, 0xb7, 0xdf, 0x9b, 0x92, 0x12, 0x1f, 0xc2, 0xb6, 0x66, 0x6a, 0xcc, 0x6d,
	0xdb, 0xfa, 0xeb, 0xb6, 0x14, 0x92, 0xaf, 0x95, 0xfd, 0x34, 0xb8, 0x4d, 0x15, 0x7c, 0xeb, 0x77,
	0xf1, 0x95, 0x3c, 0x4c, 0xb8, 0xff, 0x3a, 0x8f, 0x1f, 0x08, 0xee, 0x9a, 0x0d, 0x1d, 0x7d, 0x62,
	0x62, 0xcc, 0xff, 0x87, 0x13, 0xfa, 0x56, 0xdd, 0xd0, 0x5f, 0x3f, 0xa2, 0x8f, 0x65, 0x7e, 0xfd,
	0x11, 0x6d, 0x36, 0xf8, 0x1c, 0xc0, 0x0c, 0x3e, 0x16, 0x51, 0x7d, 0xcc, 0x9b, 0x4d, 0x9d, 0x97,
	0xa6, 0x36, 0xc8, 0x2a, 0x80, 0x2e, 0x17, 0xd1, 0xaf, 0xc9, 0xb9, 0x8a, 0xcc, 0x0a, 0xac, 0x89,
	0x68, 0x00, 0x5e, 0x6a, 0x6e, 0xb4, 0xe9, 0x47, 0xeb, 0x5f, 0x7e, 0xa7, 0xfb, 0x63, 0xe2, 0x2b,
	0xf0, 0x5e, 0x6a, 0xcd, 0xc2, 0xc9, 0xb9, 0xc4, 0x0f, 0x5c, 0x54, 0x17, 0x45, 0x84, 0xcb, 0x1f,
	0x70, 0x0f, 0xb6, 0xd2, 0x49, 0x9c, 0x24, 0xc5, 0x13, 0xe9, 0xca, 0x51, 0xd7, 0xfe, 0x37, 0x3f,
	0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0x02, 0x60, 0x4b, 0xe9, 0xca, 0x07, 0x00, 0x00,
}
