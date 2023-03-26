// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: proto/orbis/crypto/v1alpha1/crypto.proto

package cryptov1alpha1

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

type KeyType int32

const (
	KeyType_RSA       KeyType = 0
	KeyType_Ed25519   KeyType = 1
	KeyType_Secp256k1 KeyType = 2
	KeyType_ECDSA     KeyType = 3
)

// Enum value maps for KeyType.
var (
	KeyType_name = map[int32]string{
		0: "RSA",
		1: "Ed25519",
		2: "Secp256k1",
		3: "ECDSA",
	}
	KeyType_value = map[string]int32{
		"RSA":       0,
		"Ed25519":   1,
		"Secp256k1": 2,
		"ECDSA":     3,
	}
)

func (x KeyType) Enum() *KeyType {
	p := new(KeyType)
	*p = x
	return p
}

func (x KeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_orbis_crypto_v1alpha1_crypto_proto_enumTypes[0].Descriptor()
}

func (KeyType) Type() protoreflect.EnumType {
	return &file_proto_orbis_crypto_v1alpha1_crypto_proto_enumTypes[0]
}

func (x KeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyType.Descriptor instead.
func (KeyType) EnumDescriptor() ([]byte, []int) {
	return file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDescGZIP(), []int{0}
}

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type KeyType `protobuf:"varint,1,opt,name=Type,proto3,enum=orbis.crypto.v1alpha1.KeyType" json:"Type,omitempty"`
	Data []byte  `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orbis_crypto_v1alpha1_crypto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orbis_crypto_v1alpha1_crypto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDescGZIP(), []int{0}
}

func (x *PublicKey) GetType() KeyType {
	if x != nil {
		return x.Type
	}
	return KeyType_RSA
}

func (x *PublicKey) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type KeyType `protobuf:"varint,1,opt,name=Type,proto3,enum=orbis.crypto.v1alpha1.KeyType" json:"Type,omitempty"`
	Data []byte  `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *PrivateKey) Reset() {
	*x = PrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_orbis_crypto_v1alpha1_crypto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKey) ProtoMessage() {}

func (x *PrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_orbis_crypto_v1alpha1_crypto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKey.ProtoReflect.Descriptor instead.
func (*PrivateKey) Descriptor() ([]byte, []int) {
	return file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDescGZIP(), []int{1}
}

func (x *PrivateKey) GetType() KeyType {
	if x != nil {
		return x.Type
	}
	return KeyType_RSA
}

func (x *PrivateKey) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_orbis_crypto_v1alpha1_crypto_proto protoreflect.FileDescriptor

var file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2f, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x72, 0x62, 0x69,
	0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x22, 0x53, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x32,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6f,
	0x72, 0x62, 0x69, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x54, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x2a, 0x39, 0x0a, 0x07,
	0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x53, 0x41, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x45, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x43, 0x44, 0x53, 0x41, 0x10, 0x03, 0x42, 0xf0, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e,
	0x6f, 0x72, 0x62, 0x69, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0b, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6f,
	0x72, 0x62, 0x69, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x43, 0x58, 0xaa, 0x02, 0x15, 0x4f,
	0x72, 0x62, 0x69, 0x73, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x15, 0x4f, 0x72, 0x62, 0x69, 0x73, 0x5c, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x21, 0x4f,
	0x72, 0x62, 0x69, 0x73, 0x5c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x17, 0x4f, 0x72, 0x62, 0x69, 0x73, 0x3a, 0x3a, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDescOnce sync.Once
	file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDescData = file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDesc
)

func file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDescGZIP() []byte {
	file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDescOnce.Do(func() {
		file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDescData)
	})
	return file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDescData
}

var file_proto_orbis_crypto_v1alpha1_crypto_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_orbis_crypto_v1alpha1_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_orbis_crypto_v1alpha1_crypto_proto_goTypes = []interface{}{
	(KeyType)(0),       // 0: orbis.crypto.v1alpha1.KeyType
	(*PublicKey)(nil),  // 1: orbis.crypto.v1alpha1.PublicKey
	(*PrivateKey)(nil), // 2: orbis.crypto.v1alpha1.PrivateKey
}
var file_proto_orbis_crypto_v1alpha1_crypto_proto_depIdxs = []int32{
	0, // 0: orbis.crypto.v1alpha1.PublicKey.Type:type_name -> orbis.crypto.v1alpha1.KeyType
	0, // 1: orbis.crypto.v1alpha1.PrivateKey.Type:type_name -> orbis.crypto.v1alpha1.KeyType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_orbis_crypto_v1alpha1_crypto_proto_init() }
func file_proto_orbis_crypto_v1alpha1_crypto_proto_init() {
	if File_proto_orbis_crypto_v1alpha1_crypto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_orbis_crypto_v1alpha1_crypto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKey); i {
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
		file_proto_orbis_crypto_v1alpha1_crypto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateKey); i {
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
			RawDescriptor: file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_orbis_crypto_v1alpha1_crypto_proto_goTypes,
		DependencyIndexes: file_proto_orbis_crypto_v1alpha1_crypto_proto_depIdxs,
		EnumInfos:         file_proto_orbis_crypto_v1alpha1_crypto_proto_enumTypes,
		MessageInfos:      file_proto_orbis_crypto_v1alpha1_crypto_proto_msgTypes,
	}.Build()
	File_proto_orbis_crypto_v1alpha1_crypto_proto = out.File
	file_proto_orbis_crypto_v1alpha1_crypto_proto_rawDesc = nil
	file_proto_orbis_crypto_v1alpha1_crypto_proto_goTypes = nil
	file_proto_orbis_crypto_v1alpha1_crypto_proto_depIdxs = nil
}