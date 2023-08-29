// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: v2raypool.proto

package grpc

import (
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

type ProxyNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index      uint32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Id         string  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	LocalPort  uint32  `protobuf:"varint,3,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	Speed      float32 `protobuf:"fixed32,4,opt,name=speed,proto3" json:"speed,omitempty"`
	Title      string  `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	LocalAddr  string  `protobuf:"bytes,6,opt,name=local_addr,json=localAddr,proto3" json:"local_addr,omitempty"`
	RemoteAddr string  `protobuf:"bytes,7,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	IsRunning  bool    `protobuf:"varint,8,opt,name=is_running,json=isRunning,proto3" json:"is_running,omitempty"`
	IsOk       bool    `protobuf:"varint,9,opt,name=is_ok,json=isOk,proto3" json:"is_ok,omitempty"`
	TestAt     string  `protobuf:"bytes,10,opt,name=test_at,json=testAt,proto3" json:"test_at,omitempty"`
}

func (x *ProxyNode) Reset() {
	*x = ProxyNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2raypool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyNode) ProtoMessage() {}

func (x *ProxyNode) ProtoReflect() protoreflect.Message {
	mi := &file_v2raypool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyNode.ProtoReflect.Descriptor instead.
func (*ProxyNode) Descriptor() ([]byte, []int) {
	return file_v2raypool_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyNode) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ProxyNode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProxyNode) GetLocalPort() uint32 {
	if x != nil {
		return x.LocalPort
	}
	return 0
}

func (x *ProxyNode) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *ProxyNode) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProxyNode) GetLocalAddr() string {
	if x != nil {
		return x.LocalAddr
	}
	return ""
}

func (x *ProxyNode) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *ProxyNode) GetIsRunning() bool {
	if x != nil {
		return x.IsRunning
	}
	return false
}

func (x *ProxyNode) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *ProxyNode) GetTestAt() string {
	if x != nil {
		return x.TestAt
	}
	return ""
}

type OptResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *OptResult) Reset() {
	*x = OptResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2raypool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptResult) ProtoMessage() {}

func (x *OptResult) ProtoReflect() protoreflect.Message {
	mi := &file_v2raypool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptResult.ProtoReflect.Descriptor instead.
func (*OptResult) Descriptor() ([]byte, []int) {
	return file_v2raypool_proto_rawDescGZIP(), []int{1}
}

func (x *OptResult) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OptResult) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type OptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OptRequest) Reset() {
	*x = OptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2raypool_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptRequest) ProtoMessage() {}

func (x *OptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2raypool_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptRequest.ProtoReflect.Descriptor instead.
func (*OptRequest) Descriptor() ([]byte, []int) {
	return file_v2raypool_proto_rawDescGZIP(), []int{2}
}

type OptRequestDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *OptRequestDomain) Reset() {
	*x = OptRequestDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2raypool_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptRequestDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptRequestDomain) ProtoMessage() {}

func (x *OptRequestDomain) ProtoReflect() protoreflect.Message {
	mi := &file_v2raypool_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptRequestDomain.ProtoReflect.Descriptor instead.
func (*OptRequestDomain) Descriptor() ([]byte, []int) {
	return file_v2raypool_proto_rawDescGZIP(), []int{3}
}

func (x *OptRequestDomain) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type OptRequestUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *OptRequestUrl) Reset() {
	*x = OptRequestUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2raypool_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptRequestUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptRequestUrl) ProtoMessage() {}

func (x *OptRequestUrl) ProtoReflect() protoreflect.Message {
	mi := &file_v2raypool_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptRequestUrl.ProtoReflect.Descriptor instead.
func (*OptRequestUrl) Descriptor() ([]byte, []int) {
	return file_v2raypool_proto_rawDescGZIP(), []int{4}
}

func (x *OptRequestUrl) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type UpdateSubscribeResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Total  uint32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Add    uint32 `protobuf:"varint,4,opt,name=add,proto3" json:"add,omitempty"`
}

func (x *UpdateSubscribeResult) Reset() {
	*x = UpdateSubscribeResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2raypool_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSubscribeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSubscribeResult) ProtoMessage() {}

func (x *UpdateSubscribeResult) ProtoReflect() protoreflect.Message {
	mi := &file_v2raypool_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSubscribeResult.ProtoReflect.Descriptor instead.
func (*UpdateSubscribeResult) Descriptor() ([]byte, []int) {
	return file_v2raypool_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSubscribeResult) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateSubscribeResult) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UpdateSubscribeResult) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UpdateSubscribeResult) GetAdd() uint32 {
	if x != nil {
		return x.Add
	}
	return 0
}

type KillNodesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Total   uint32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Runport uint32 `protobuf:"varint,4,opt,name=runport,proto3" json:"runport,omitempty"`
	Kill    uint32 `protobuf:"varint,5,opt,name=kill,proto3" json:"kill,omitempty"`
	Fail    uint32 `protobuf:"varint,6,opt,name=fail,proto3" json:"fail,omitempty"`
}

func (x *KillNodesResult) Reset() {
	*x = KillNodesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2raypool_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillNodesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillNodesResult) ProtoMessage() {}

func (x *KillNodesResult) ProtoReflect() protoreflect.Message {
	mi := &file_v2raypool_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillNodesResult.ProtoReflect.Descriptor instead.
func (*KillNodesResult) Descriptor() ([]byte, []int) {
	return file_v2raypool_proto_rawDescGZIP(), []int{6}
}

func (x *KillNodesResult) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *KillNodesResult) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *KillNodesResult) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *KillNodesResult) GetRunport() uint32 {
	if x != nil {
		return x.Runport
	}
	return 0
}

func (x *KillNodesResult) GetKill() uint32 {
	if x != nil {
		return x.Kill
	}
	return 0
}

func (x *KillNodesResult) GetFail() uint32 {
	if x != nil {
		return x.Fail
	}
	return 0
}

type ProxyNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ProxyNode `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ProxyNodes) Reset() {
	*x = ProxyNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2raypool_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyNodes) ProtoMessage() {}

func (x *ProxyNodes) ProtoReflect() protoreflect.Message {
	mi := &file_v2raypool_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyNodes.ProtoReflect.Descriptor instead.
func (*ProxyNodes) Descriptor() ([]byte, []int) {
	return file_v2raypool_proto_rawDescGZIP(), []int{7}
}

func (x *ProxyNodes) GetItems() []*ProxyNode {
	if x != nil {
		return x.Items
	}
	return nil
}

type ProxyNodes_OptResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ProxyNodes_OptResult) Reset() {
	*x = ProxyNodes_OptResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2raypool_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyNodes_OptResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyNodes_OptResult) ProtoMessage() {}

func (x *ProxyNodes_OptResult) ProtoReflect() protoreflect.Message {
	mi := &file_v2raypool_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyNodes_OptResult.ProtoReflect.Descriptor instead.
func (*ProxyNodes_OptResult) Descriptor() ([]byte, []int) {
	return file_v2raypool_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ProxyNodes_OptResult) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ProxyNodes_OptResult) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_v2raypool_proto protoreflect.FileDescriptor

var file_v2raypool_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x76, 0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x22, 0x89, 0x02, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x0a, 0x05, 0x69, 0x73,
	0x5f, 0x6f, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x41, 0x74, 0x22, 0x35, 0x0a, 0x09, 0x4f, 0x70, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x0c, 0x0a, 0x0a, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a,
	0x10, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x21, 0x0a, 0x0d, 0x4f, 0x70, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x69, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x61, 0x64, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x4b, 0x69, 0x6c, 0x6c,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x75, 0x6e, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x75,
	0x6e, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x69,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x61, 0x69, 0x6c, 0x22, 0x6f, 0x0a,
	0x0a, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x35, 0x0a, 0x09, 0x4f, 0x70, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xcc,
	0x05, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x15, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x42, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x15, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x55, 0x72, 0x6c,
	0x12, 0x18, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x70, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x72, 0x6c, 0x1a, 0x14, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x50, 0x6f, 0x6f, 0x6c, 0x41, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x41, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x70,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x10, 0x54, 0x65, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x41, 0x6c, 0x6c, 0x12, 0x15, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x15,
	0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x41, 0x6c, 0x6c,
	0x46, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f,
	0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x4b, 0x69, 0x6c, 0x6c, 0x41, 0x6c, 0x6c, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x4f, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0f, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x6f,
	0x64, 0x65, 0x1a, 0x14, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f,
	0x70, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x12, 0x15, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2raypool_proto_rawDescOnce sync.Once
	file_v2raypool_proto_rawDescData = file_v2raypool_proto_rawDesc
)

func file_v2raypool_proto_rawDescGZIP() []byte {
	file_v2raypool_proto_rawDescOnce.Do(func() {
		file_v2raypool_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2raypool_proto_rawDescData)
	})
	return file_v2raypool_proto_rawDescData
}

var file_v2raypool_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v2raypool_proto_goTypes = []interface{}{
	(*ProxyNode)(nil),             // 0: v2raypool.ProxyNode
	(*OptResult)(nil),             // 1: v2raypool.OptResult
	(*OptRequest)(nil),            // 2: v2raypool.OptRequest
	(*OptRequestDomain)(nil),      // 3: v2raypool.OptRequestDomain
	(*OptRequestUrl)(nil),         // 4: v2raypool.OptRequestUrl
	(*UpdateSubscribeResult)(nil), // 5: v2raypool.UpdateSubscribeResult
	(*KillNodesResult)(nil),       // 6: v2raypool.KillNodesResult
	(*ProxyNodes)(nil),            // 7: v2raypool.ProxyNodes
	(*ProxyNodes_OptResult)(nil),  // 8: v2raypool.ProxyNodes.OptResult
}
var file_v2raypool_proto_depIdxs = []int32{
	0,  // 0: v2raypool.ProxyNodes.items:type_name -> v2raypool.ProxyNode
	0,  // 1: v2raypool.ProxyPoolService.GetProxyNodes:input_type -> v2raypool.ProxyNode
	3,  // 2: v2raypool.ProxyPoolService.GetProxyNodesByDomain:input_type -> v2raypool.OptRequestDomain
	4,  // 3: v2raypool.ProxyPoolService.SetTestUrl:input_type -> v2raypool.OptRequestUrl
	2,  // 4: v2raypool.ProxyPoolService.StartProxyPoolAll:input_type -> v2raypool.OptRequest
	2,  // 5: v2raypool.ProxyPoolService.StopProxyPoolAll:input_type -> v2raypool.OptRequest
	2,  // 6: v2raypool.ProxyPoolService.TestProxyPoolAll:input_type -> v2raypool.OptRequest
	2,  // 7: v2raypool.ProxyPoolService.TestProxyPoolAllForce:input_type -> v2raypool.OptRequest
	2,  // 8: v2raypool.ProxyPoolService.KillAllNodes:input_type -> v2raypool.OptRequest
	0,  // 9: v2raypool.ProxyPoolService.ActiveProxyNode:input_type -> v2raypool.ProxyNode
	2,  // 10: v2raypool.ProxyPoolService.UpdateProxySubscribe:input_type -> v2raypool.OptRequest
	7,  // 11: v2raypool.ProxyPoolService.GetProxyNodes:output_type -> v2raypool.ProxyNodes
	7,  // 12: v2raypool.ProxyPoolService.GetProxyNodesByDomain:output_type -> v2raypool.ProxyNodes
	1,  // 13: v2raypool.ProxyPoolService.SetTestUrl:output_type -> v2raypool.OptResult
	1,  // 14: v2raypool.ProxyPoolService.StartProxyPoolAll:output_type -> v2raypool.OptResult
	1,  // 15: v2raypool.ProxyPoolService.StopProxyPoolAll:output_type -> v2raypool.OptResult
	1,  // 16: v2raypool.ProxyPoolService.TestProxyPoolAll:output_type -> v2raypool.OptResult
	1,  // 17: v2raypool.ProxyPoolService.TestProxyPoolAllForce:output_type -> v2raypool.OptResult
	6,  // 18: v2raypool.ProxyPoolService.KillAllNodes:output_type -> v2raypool.KillNodesResult
	1,  // 19: v2raypool.ProxyPoolService.ActiveProxyNode:output_type -> v2raypool.OptResult
	5,  // 20: v2raypool.ProxyPoolService.UpdateProxySubscribe:output_type -> v2raypool.UpdateSubscribeResult
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_v2raypool_proto_init() }
func file_v2raypool_proto_init() {
	if File_v2raypool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2raypool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyNode); i {
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
		file_v2raypool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptResult); i {
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
		file_v2raypool_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptRequest); i {
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
		file_v2raypool_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptRequestDomain); i {
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
		file_v2raypool_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptRequestUrl); i {
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
		file_v2raypool_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSubscribeResult); i {
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
		file_v2raypool_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillNodesResult); i {
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
		file_v2raypool_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyNodes); i {
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
		file_v2raypool_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyNodes_OptResult); i {
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
			RawDescriptor: file_v2raypool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v2raypool_proto_goTypes,
		DependencyIndexes: file_v2raypool_proto_depIdxs,
		MessageInfos:      file_v2raypool_proto_msgTypes,
	}.Build()
	File_v2raypool_proto = out.File
	file_v2raypool_proto_rawDesc = nil
	file_v2raypool_proto_goTypes = nil
	file_v2raypool_proto_depIdxs = nil
}
