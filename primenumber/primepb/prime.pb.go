// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: primenumber/primepb/prime.proto

package primepb

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

type Number struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primenumber_primepb_prime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_primenumber_primepb_prime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Number.ProtoReflect.Descriptor instead.
func (*Number) Descriptor() ([]byte, []int) {
	return file_primenumber_primepb_prime_proto_rawDescGZIP(), []int{0}
}

func (x *Number) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type PrimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number *Number `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PrimeRequest) Reset() {
	*x = PrimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primenumber_primepb_prime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeRequest) ProtoMessage() {}

func (x *PrimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_primenumber_primepb_prime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeRequest.ProtoReflect.Descriptor instead.
func (*PrimeRequest) Descriptor() ([]byte, []int) {
	return file_primenumber_primepb_prime_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeRequest) GetNumber() *Number {
	if x != nil {
		return x.Number
	}
	return nil
}

type PrimemultipleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimemultipleResponse) Reset() {
	*x = PrimemultipleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primenumber_primepb_prime_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimemultipleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimemultipleResponse) ProtoMessage() {}

func (x *PrimemultipleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_primenumber_primepb_prime_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimemultipleResponse.ProtoReflect.Descriptor instead.
func (*PrimemultipleResponse) Descriptor() ([]byte, []int) {
	return file_primenumber_primepb_prime_proto_rawDescGZIP(), []int{2}
}

func (x *PrimemultipleResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_primenumber_primepb_prime_proto protoreflect.FileDescriptor

var file_primenumber_primepb_prime_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x1a,
	0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x3b, 0x0a, 0x0c, 0x70, 0x72,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x15, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x5a, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x6d,
	0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x69, 0x6d, 0x65,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_primenumber_primepb_prime_proto_rawDescOnce sync.Once
	file_primenumber_primepb_prime_proto_rawDescData = file_primenumber_primepb_prime_proto_rawDesc
)

func file_primenumber_primepb_prime_proto_rawDescGZIP() []byte {
	file_primenumber_primepb_prime_proto_rawDescOnce.Do(func() {
		file_primenumber_primepb_prime_proto_rawDescData = protoimpl.X.CompressGZIP(file_primenumber_primepb_prime_proto_rawDescData)
	})
	return file_primenumber_primepb_prime_proto_rawDescData
}

var file_primenumber_primepb_prime_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_primenumber_primepb_prime_proto_goTypes = []interface{}{
	(*Number)(nil),                // 0: primenumber.Number
	(*PrimeRequest)(nil),          // 1: primenumber.primeRequest
	(*PrimemultipleResponse)(nil), // 2: primenumber.primemultipleResponse
}
var file_primenumber_primepb_prime_proto_depIdxs = []int32{
	0, // 0: primenumber.primeRequest.number:type_name -> primenumber.Number
	1, // 1: primenumber.primeService.Prime:input_type -> primenumber.primeRequest
	2, // 2: primenumber.primeService.Prime:output_type -> primenumber.primemultipleResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_primenumber_primepb_prime_proto_init() }
func file_primenumber_primepb_prime_proto_init() {
	if File_primenumber_primepb_prime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_primenumber_primepb_prime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Number); i {
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
		file_primenumber_primepb_prime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeRequest); i {
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
		file_primenumber_primepb_prime_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimemultipleResponse); i {
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
			RawDescriptor: file_primenumber_primepb_prime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_primenumber_primepb_prime_proto_goTypes,
		DependencyIndexes: file_primenumber_primepb_prime_proto_depIdxs,
		MessageInfos:      file_primenumber_primepb_prime_proto_msgTypes,
	}.Build()
	File_primenumber_primepb_prime_proto = out.File
	file_primenumber_primepb_prime_proto_rawDesc = nil
	file_primenumber_primepb_prime_proto_goTypes = nil
	file_primenumber_primepb_prime_proto_depIdxs = nil
}
