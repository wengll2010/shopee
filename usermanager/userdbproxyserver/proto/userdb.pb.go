// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: userdb.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 添加用户
type AddDBUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Nick     string `protobuf:"bytes,3,opt,name=nick,proto3" json:"nick,omitempty"`
	Picture  string `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	Regtime  int32  `protobuf:"varint,5,opt,name=regtime,proto3" json:"regtime,omitempty"`
}

func (x *AddDBUserRequest) Reset() {
	*x = AddDBUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDBUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDBUserRequest) ProtoMessage() {}

func (x *AddDBUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDBUserRequest.ProtoReflect.Descriptor instead.
func (*AddDBUserRequest) Descriptor() ([]byte, []int) {
	return file_userdb_proto_rawDescGZIP(), []int{0}
}

func (x *AddDBUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddDBUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddDBUserRequest) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

func (x *AddDBUserRequest) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *AddDBUserRequest) GetRegtime() int32 {
	if x != nil {
		return x.Regtime
	}
	return 0
}

type AddDBUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Errmsg  string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
}

func (x *AddDBUserReply) Reset() {
	*x = AddDBUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDBUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDBUserReply) ProtoMessage() {}

func (x *AddDBUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_userdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDBUserReply.ProtoReflect.Descriptor instead.
func (*AddDBUserReply) Descriptor() ([]byte, []int) {
	return file_userdb_proto_rawDescGZIP(), []int{1}
}

func (x *AddDBUserReply) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *AddDBUserReply) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

// 查询用户
type GetDBUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetDBUserRequest) Reset() {
	*x = GetDBUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDBUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDBUserRequest) ProtoMessage() {}

func (x *GetDBUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDBUserRequest.ProtoReflect.Descriptor instead.
func (*GetDBUserRequest) Descriptor() ([]byte, []int) {
	return file_userdb_proto_rawDescGZIP(), []int{2}
}

func (x *GetDBUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetDBUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode  int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Errmsg   string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Nick     string `protobuf:"bytes,5,opt,name=nick,proto3" json:"nick,omitempty"`
	Picture  string `protobuf:"bytes,6,opt,name=picture,proto3" json:"picture,omitempty"`
	Regtime  int32  `protobuf:"varint,7,opt,name=regtime,proto3" json:"regtime,omitempty"`
}

func (x *GetDBUserReply) Reset() {
	*x = GetDBUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDBUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDBUserReply) ProtoMessage() {}

func (x *GetDBUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_userdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDBUserReply.ProtoReflect.Descriptor instead.
func (*GetDBUserReply) Descriptor() ([]byte, []int) {
	return file_userdb_proto_rawDescGZIP(), []int{3}
}

func (x *GetDBUserReply) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetDBUserReply) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *GetDBUserReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetDBUserReply) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GetDBUserReply) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

func (x *GetDBUserReply) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *GetDBUserReply) GetRegtime() int32 {
	if x != nil {
		return x.Regtime
	}
	return 0
}

// 修改用户
type UpdateDBUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Nick    string `protobuf:"bytes,2,opt,name=nick,proto3" json:"nick,omitempty"`
	Picture string `protobuf:"bytes,3,opt,name=picture,proto3" json:"picture,omitempty"`
}

func (x *UpdateDBUserRequest) Reset() {
	*x = UpdateDBUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDBUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDBUserRequest) ProtoMessage() {}

func (x *UpdateDBUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDBUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateDBUserRequest) Descriptor() ([]byte, []int) {
	return file_userdb_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDBUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDBUserRequest) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

func (x *UpdateDBUserRequest) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

type UpdateDBUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Errmsg  string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
}

func (x *UpdateDBUserReply) Reset() {
	*x = UpdateDBUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDBUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDBUserReply) ProtoMessage() {}

func (x *UpdateDBUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_userdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDBUserReply.ProtoReflect.Descriptor instead.
func (*UpdateDBUserReply) Descriptor() ([]byte, []int) {
	return file_userdb_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDBUserReply) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *UpdateDBUserReply) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

var File_userdb_proto protoreflect.FileDescriptor

var file_userdb_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x44, 0x42, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x69,
	0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x67, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x42, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x44, 0x42, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x42, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xba,
	0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x42, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72,
	0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x67, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x42, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x45, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x42,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x32, 0xce, 0x01, 0x0a, 0x06,
	0x44, 0x42, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x44, 0x42, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x44,
	0x42, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x42, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x42, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x42,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x42, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x42,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x42, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x42, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userdb_proto_rawDescOnce sync.Once
	file_userdb_proto_rawDescData = file_userdb_proto_rawDesc
)

func file_userdb_proto_rawDescGZIP() []byte {
	file_userdb_proto_rawDescOnce.Do(func() {
		file_userdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_userdb_proto_rawDescData)
	})
	return file_userdb_proto_rawDescData
}

var file_userdb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_userdb_proto_goTypes = []interface{}{
	(*AddDBUserRequest)(nil),    // 0: proto.AddDBUserRequest
	(*AddDBUserReply)(nil),      // 1: proto.AddDBUserReply
	(*GetDBUserRequest)(nil),    // 2: proto.GetDBUserRequest
	(*GetDBUserReply)(nil),      // 3: proto.GetDBUserReply
	(*UpdateDBUserRequest)(nil), // 4: proto.UpdateDBUserRequest
	(*UpdateDBUserReply)(nil),   // 5: proto.UpdateDBUserReply
}
var file_userdb_proto_depIdxs = []int32{
	0, // 0: proto.DBUser.AddDBUser:input_type -> proto.AddDBUserRequest
	2, // 1: proto.DBUser.GetDBUser:input_type -> proto.GetDBUserRequest
	4, // 2: proto.DBUser.UpdateDBUser:input_type -> proto.UpdateDBUserRequest
	1, // 3: proto.DBUser.AddDBUser:output_type -> proto.AddDBUserReply
	3, // 4: proto.DBUser.GetDBUser:output_type -> proto.GetDBUserReply
	5, // 5: proto.DBUser.UpdateDBUser:output_type -> proto.UpdateDBUserReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userdb_proto_init() }
func file_userdb_proto_init() {
	if File_userdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDBUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_userdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDBUserReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_userdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDBUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_userdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDBUserReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_userdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDBUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_userdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDBUserReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_userdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userdb_proto_goTypes,
		DependencyIndexes: file_userdb_proto_depIdxs,
		MessageInfos:      file_userdb_proto_msgTypes,
	}.Build()
	File_userdb_proto = out.File
	file_userdb_proto_rawDesc = nil
	file_userdb_proto_goTypes = nil
	file_userdb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DBUserClient is the client API for DBUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DBUserClient interface {
	AddDBUser(ctx context.Context, in *AddDBUserRequest, opts ...grpc.CallOption) (*AddDBUserReply, error)
	GetDBUser(ctx context.Context, in *GetDBUserRequest, opts ...grpc.CallOption) (*GetDBUserReply, error)
	UpdateDBUser(ctx context.Context, in *UpdateDBUserRequest, opts ...grpc.CallOption) (*UpdateDBUserReply, error)
}

type dBUserClient struct {
	cc grpc.ClientConnInterface
}

func NewDBUserClient(cc grpc.ClientConnInterface) DBUserClient {
	return &dBUserClient{cc}
}

func (c *dBUserClient) AddDBUser(ctx context.Context, in *AddDBUserRequest, opts ...grpc.CallOption) (*AddDBUserReply, error) {
	out := new(AddDBUserReply)
	err := c.cc.Invoke(ctx, "/proto.DBUser/AddDBUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBUserClient) GetDBUser(ctx context.Context, in *GetDBUserRequest, opts ...grpc.CallOption) (*GetDBUserReply, error) {
	out := new(GetDBUserReply)
	err := c.cc.Invoke(ctx, "/proto.DBUser/GetDBUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBUserClient) UpdateDBUser(ctx context.Context, in *UpdateDBUserRequest, opts ...grpc.CallOption) (*UpdateDBUserReply, error) {
	out := new(UpdateDBUserReply)
	err := c.cc.Invoke(ctx, "/proto.DBUser/UpdateDBUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBUserServer is the server API for DBUser service.
type DBUserServer interface {
	AddDBUser(context.Context, *AddDBUserRequest) (*AddDBUserReply, error)
	GetDBUser(context.Context, *GetDBUserRequest) (*GetDBUserReply, error)
	UpdateDBUser(context.Context, *UpdateDBUserRequest) (*UpdateDBUserReply, error)
}

// UnimplementedDBUserServer can be embedded to have forward compatible implementations.
type UnimplementedDBUserServer struct {
}

func (*UnimplementedDBUserServer) AddDBUser(context.Context, *AddDBUserRequest) (*AddDBUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDBUser not implemented")
}
func (*UnimplementedDBUserServer) GetDBUser(context.Context, *GetDBUserRequest) (*GetDBUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDBUser not implemented")
}
func (*UnimplementedDBUserServer) UpdateDBUser(context.Context, *UpdateDBUserRequest) (*UpdateDBUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDBUser not implemented")
}

func RegisterDBUserServer(s *grpc.Server, srv DBUserServer) {
	s.RegisterService(&_DBUser_serviceDesc, srv)
}

func _DBUser_AddDBUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDBUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBUserServer).AddDBUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBUser/AddDBUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBUserServer).AddDBUser(ctx, req.(*AddDBUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBUser_GetDBUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDBUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBUserServer).GetDBUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBUser/GetDBUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBUserServer).GetDBUser(ctx, req.(*GetDBUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBUser_UpdateDBUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDBUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBUserServer).UpdateDBUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBUser/UpdateDBUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBUserServer).UpdateDBUser(ctx, req.(*UpdateDBUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DBUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DBUser",
	HandlerType: (*DBUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDBUser",
			Handler:    _DBUser_AddDBUser_Handler,
		},
		{
			MethodName: "GetDBUser",
			Handler:    _DBUser_GetDBUser_Handler,
		},
		{
			MethodName: "UpdateDBUser",
			Handler:    _DBUser_UpdateDBUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userdb.proto",
}