// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: orbis/rabin/v1alpha1/dkg.proto

package rabinv1alpha1

import (
	pb "github.com/libp2p/go-libp2p/core/crypto/pb"
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

type SuiteType int32

const (
	SuiteType_NONE      SuiteType = 0
	SuiteType_Ed25519   SuiteType = 1
	SuiteType_Secp256k1 SuiteType = 2
)

// Enum value maps for SuiteType.
var (
	SuiteType_name = map[int32]string{
		0: "NONE",
		1: "Ed25519",
		2: "Secp256k1",
	}
	SuiteType_value = map[string]int32{
		"NONE":      0,
		"Ed25519":   1,
		"Secp256k1": 2,
	}
)

func (x SuiteType) Enum() *SuiteType {
	p := new(SuiteType)
	*p = x
	return p
}

func (x SuiteType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SuiteType) Descriptor() protoreflect.EnumDescriptor {
	return file_orbis_rabin_v1alpha1_dkg_proto_enumTypes[0].Descriptor()
}

func (SuiteType) Type() protoreflect.EnumType {
	return &file_orbis_rabin_v1alpha1_dkg_proto_enumTypes[0]
}

func (x SuiteType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SuiteType.Descriptor instead.
func (SuiteType) EnumDescriptor() ([]byte, []int) {
	return file_orbis_rabin_v1alpha1_dkg_proto_rawDescGZIP(), []int{0}
}

type State int32

const (
	State_STATE_INITIALIZED         State = 0
	State_STATE_STARTED             State = 1
	State_STATE_CERTIFIED           State = 2
	State_STATE_PROCESSED_DEALS     State = 128
	State_STATE_PROCESSED_RESPONSES State = 129
	State_STATE_PROCESSED_COMMITS   State = 130
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0:   "STATE_INITIALIZED",
		1:   "STATE_STARTED",
		2:   "STATE_CERTIFIED",
		128: "STATE_PROCESSED_DEALS",
		129: "STATE_PROCESSED_RESPONSES",
		130: "STATE_PROCESSED_COMMITS",
	}
	State_value = map[string]int32{
		"STATE_INITIALIZED":         0,
		"STATE_STARTED":             1,
		"STATE_CERTIFIED":           2,
		"STATE_PROCESSED_DEALS":     128,
		"STATE_PROCESSED_RESPONSES": 129,
		"STATE_PROCESSED_COMMITS":   130,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_orbis_rabin_v1alpha1_dkg_proto_enumTypes[1].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_orbis_rabin_v1alpha1_dkg_proto_enumTypes[1]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_orbis_rabin_v1alpha1_dkg_proto_rawDescGZIP(), []int{1}
}

type DKG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RingId    string    `protobuf:"bytes,1,opt,name=ring_id,json=ringId,proto3" json:"ring_id,omitempty"`
	Index     int32     `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Num       int32     `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
	Threshold int32     `protobuf:"varint,4,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Suite     SuiteType `protobuf:"varint,5,opt,name=suite,proto3,enum=orbis.rabin.v1alpha1.SuiteType" json:"suite,omitempty"`
	State     State     `protobuf:"varint,6,opt,name=state,proto3,enum=orbis.rabin.v1alpha1.State" json:"state,omitempty"`
	Pubkey    []byte    `protobuf:"bytes,7,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	PriShare  *PriShare `protobuf:"bytes,8,opt,name=pri_share,json=priShare,proto3" json:"pri_share,omitempty"`
	Nodes     []*Node   `protobuf:"bytes,9,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *DKG) Reset() {
	*x = DKG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_rabin_v1alpha1_dkg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DKG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DKG) ProtoMessage() {}

func (x *DKG) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_rabin_v1alpha1_dkg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DKG.ProtoReflect.Descriptor instead.
func (*DKG) Descriptor() ([]byte, []int) {
	return file_orbis_rabin_v1alpha1_dkg_proto_rawDescGZIP(), []int{0}
}

func (x *DKG) GetRingId() string {
	if x != nil {
		return x.RingId
	}
	return ""
}

func (x *DKG) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *DKG) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *DKG) GetThreshold() int32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *DKG) GetSuite() SuiteType {
	if x != nil {
		return x.Suite
	}
	return SuiteType_NONE
}

func (x *DKG) GetState() State {
	if x != nil {
		return x.State
	}
	return State_STATE_INITIALIZED
}

func (x *DKG) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

func (x *DKG) GetPriShare() *PriShare {
	if x != nil {
		return x.PriShare
	}
	return nil
}

func (x *DKG) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address   string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"` // multiaddress
	PublicKey *pb.PublicKey `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_rabin_v1alpha1_dkg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_rabin_v1alpha1_dkg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_orbis_rabin_v1alpha1_dkg_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Node) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Node) GetPublicKey() *pb.PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type PriShare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	V     []byte `protobuf:"bytes,2,opt,name=v,proto3" json:"v,omitempty"`
}

func (x *PriShare) Reset() {
	*x = PriShare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orbis_rabin_v1alpha1_dkg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriShare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriShare) ProtoMessage() {}

func (x *PriShare) ProtoReflect() protoreflect.Message {
	mi := &file_orbis_rabin_v1alpha1_dkg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriShare.ProtoReflect.Descriptor instead.
func (*PriShare) Descriptor() ([]byte, []int) {
	return file_orbis_rabin_v1alpha1_dkg_proto_rawDescGZIP(), []int{2}
}

func (x *PriShare) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *PriShare) GetV() []byte {
	if x != nil {
		return x.V
	}
	return nil
}

var File_orbis_rabin_v1alpha1_dkg_proto protoreflect.FileDescriptor

var file_orbis_rabin_v1alpha1_dkg_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2f, 0x72, 0x61, 0x62, 0x69, 0x6e, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x72, 0x61, 0x62, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1d, 0x6c, 0x69, 0x62, 0x70, 0x32, 0x70, 0x2f, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x02, 0x0a, 0x03, 0x44, 0x4b, 0x47, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x35, 0x0a, 0x05,
	0x73, 0x75, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6f, 0x72,
	0x62, 0x69, 0x73, 0x2e, 0x72, 0x61, 0x62, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x53, 0x75, 0x69, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x73, 0x75,
	0x69, 0x74, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x72, 0x61, 0x62, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x3b,
	0x0a, 0x09, 0x70, 0x72, 0x69, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x72, 0x61, 0x62, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x52, 0x08, 0x70, 0x72, 0x69, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x72, 0x62,
	0x69, 0x73, 0x2e, 0x72, 0x61, 0x62, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x6c, 0x0a,
	0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x3a, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x62, 0x70, 0x32, 0x70, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x2e, 0x0a, 0x08, 0x50,
	0x72, 0x69, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x76, 0x2a, 0x31, 0x0a, 0x09, 0x53,
	0x75, 0x69, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x10, 0x02, 0x2a, 0xa0,
	0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x45, 0x52, 0x54,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x15, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x41, 0x4c, 0x53,
	0x10, 0x80, 0x01, 0x12, 0x1e, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f,
	0x43, 0x45, 0x53, 0x53, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x53,
	0x10, 0x81, 0x01, 0x12, 0x1c, 0x0a, 0x17, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f,
	0x43, 0x45, 0x53, 0x53, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x53, 0x10, 0x82,
	0x01, 0x42, 0xe6, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e,
	0x72, 0x61, 0x62, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x08,
	0x44, 0x6b, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2f, 0x72, 0x61,
	0x62, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x72, 0x61, 0x62,
	0x69, 0x6e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x52, 0x58,
	0xaa, 0x02, 0x14, 0x4f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x52, 0x61, 0x62, 0x69, 0x6e, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x14, 0x4f, 0x72, 0x62, 0x69, 0x73, 0x5c,
	0x52, 0x61, 0x62, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02,
	0x20, 0x4f, 0x72, 0x62, 0x69, 0x73, 0x5c, 0x52, 0x61, 0x62, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x16, 0x4f, 0x72, 0x62, 0x69, 0x73, 0x3a, 0x3a, 0x52, 0x61, 0x62, 0x69, 0x6e,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_orbis_rabin_v1alpha1_dkg_proto_rawDescOnce sync.Once
	file_orbis_rabin_v1alpha1_dkg_proto_rawDescData = file_orbis_rabin_v1alpha1_dkg_proto_rawDesc
)

func file_orbis_rabin_v1alpha1_dkg_proto_rawDescGZIP() []byte {
	file_orbis_rabin_v1alpha1_dkg_proto_rawDescOnce.Do(func() {
		file_orbis_rabin_v1alpha1_dkg_proto_rawDescData = protoimpl.X.CompressGZIP(file_orbis_rabin_v1alpha1_dkg_proto_rawDescData)
	})
	return file_orbis_rabin_v1alpha1_dkg_proto_rawDescData
}

var file_orbis_rabin_v1alpha1_dkg_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_orbis_rabin_v1alpha1_dkg_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orbis_rabin_v1alpha1_dkg_proto_goTypes = []interface{}{
	(SuiteType)(0),       // 0: orbis.rabin.v1alpha1.SuiteType
	(State)(0),           // 1: orbis.rabin.v1alpha1.State
	(*DKG)(nil),          // 2: orbis.rabin.v1alpha1.DKG
	(*Node)(nil),         // 3: orbis.rabin.v1alpha1.Node
	(*PriShare)(nil),     // 4: orbis.rabin.v1alpha1.PriShare
	(*pb.PublicKey)(nil), // 5: libp2p.crypto.v1.PublicKey
}
var file_orbis_rabin_v1alpha1_dkg_proto_depIdxs = []int32{
	0, // 0: orbis.rabin.v1alpha1.DKG.suite:type_name -> orbis.rabin.v1alpha1.SuiteType
	1, // 1: orbis.rabin.v1alpha1.DKG.state:type_name -> orbis.rabin.v1alpha1.State
	4, // 2: orbis.rabin.v1alpha1.DKG.pri_share:type_name -> orbis.rabin.v1alpha1.PriShare
	3, // 3: orbis.rabin.v1alpha1.DKG.nodes:type_name -> orbis.rabin.v1alpha1.Node
	5, // 4: orbis.rabin.v1alpha1.Node.public_key:type_name -> libp2p.crypto.v1.PublicKey
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_orbis_rabin_v1alpha1_dkg_proto_init() }
func file_orbis_rabin_v1alpha1_dkg_proto_init() {
	if File_orbis_rabin_v1alpha1_dkg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orbis_rabin_v1alpha1_dkg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DKG); i {
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
		file_orbis_rabin_v1alpha1_dkg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_orbis_rabin_v1alpha1_dkg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriShare); i {
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
			RawDescriptor: file_orbis_rabin_v1alpha1_dkg_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orbis_rabin_v1alpha1_dkg_proto_goTypes,
		DependencyIndexes: file_orbis_rabin_v1alpha1_dkg_proto_depIdxs,
		EnumInfos:         file_orbis_rabin_v1alpha1_dkg_proto_enumTypes,
		MessageInfos:      file_orbis_rabin_v1alpha1_dkg_proto_msgTypes,
	}.Build()
	File_orbis_rabin_v1alpha1_dkg_proto = out.File
	file_orbis_rabin_v1alpha1_dkg_proto_rawDesc = nil
	file_orbis_rabin_v1alpha1_dkg_proto_goTypes = nil
	file_orbis_rabin_v1alpha1_dkg_proto_depIdxs = nil
}