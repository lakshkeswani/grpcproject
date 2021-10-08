// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: average/averagepb/average.proto

package averagepb

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

type Average struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num float64 `protobuf:"fixed64,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Average) Reset() {
	*x = Average{}
	if protoimpl.UnsafeEnabled {
		mi := &file_average_averagepb_average_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Average) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Average) ProtoMessage() {}

func (x *Average) ProtoReflect() protoreflect.Message {
	mi := &file_average_averagepb_average_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Average.ProtoReflect.Descriptor instead.
func (*Average) Descriptor() ([]byte, []int) {
	return file_average_averagepb_average_proto_rawDescGZIP(), []int{0}
}

func (x *Average) GetNum() float64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type AverageRequestt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average *Average `protobuf:"bytes,1,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *AverageRequestt) Reset() {
	*x = AverageRequestt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_average_averagepb_average_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageRequestt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRequestt) ProtoMessage() {}

func (x *AverageRequestt) ProtoReflect() protoreflect.Message {
	mi := &file_average_averagepb_average_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRequestt.ProtoReflect.Descriptor instead.
func (*AverageRequestt) Descriptor() ([]byte, []int) {
	return file_average_averagepb_average_proto_rawDescGZIP(), []int{1}
}

func (x *AverageRequestt) GetAverage() *Average {
	if x != nil {
		return x.Average
	}
	return nil
}

type AverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AverageResponse) Reset() {
	*x = AverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_average_averagepb_average_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageResponse) ProtoMessage() {}

func (x *AverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_average_averagepb_average_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageResponse.ProtoReflect.Descriptor instead.
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return file_average_averagepb_average_proto_rawDescGZIP(), []int{2}
}

func (x *AverageResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_average_averagepb_average_proto protoreflect.FileDescriptor

var file_average_averagepb_average_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x70, 0x62, 0x2f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x1b, 0x0a, 0x07, 0x41, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x3d, 0x0a, 0x0f, 0x41, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07, 0x61,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0x56, 0x0a, 0x0e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x72, 0x12, 0x18, 0x2e, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x61,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_average_averagepb_average_proto_rawDescOnce sync.Once
	file_average_averagepb_average_proto_rawDescData = file_average_averagepb_average_proto_rawDesc
)

func file_average_averagepb_average_proto_rawDescGZIP() []byte {
	file_average_averagepb_average_proto_rawDescOnce.Do(func() {
		file_average_averagepb_average_proto_rawDescData = protoimpl.X.CompressGZIP(file_average_averagepb_average_proto_rawDescData)
	})
	return file_average_averagepb_average_proto_rawDescData
}

var file_average_averagepb_average_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_average_averagepb_average_proto_goTypes = []interface{}{
	(*Average)(nil),         // 0: average.Average
	(*AverageRequestt)(nil), // 1: average.AverageRequestt
	(*AverageResponse)(nil), // 2: average.AverageResponse
}
var file_average_averagepb_average_proto_depIdxs = []int32{
	0, // 0: average.AverageRequestt.average:type_name -> average.Average
	1, // 1: average.AverageService.AverageSer:input_type -> average.AverageRequestt
	2, // 2: average.AverageService.AverageSer:output_type -> average.AverageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_average_averagepb_average_proto_init() }
func file_average_averagepb_average_proto_init() {
	if File_average_averagepb_average_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_average_averagepb_average_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Average); i {
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
		file_average_averagepb_average_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageRequestt); i {
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
		file_average_averagepb_average_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageResponse); i {
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
			RawDescriptor: file_average_averagepb_average_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_average_averagepb_average_proto_goTypes,
		DependencyIndexes: file_average_averagepb_average_proto_depIdxs,
		MessageInfos:      file_average_averagepb_average_proto_msgTypes,
	}.Build()
	File_average_averagepb_average_proto = out.File
	file_average_averagepb_average_proto_rawDesc = nil
	file_average_averagepb_average_proto_goTypes = nil
	file_average_averagepb_average_proto_depIdxs = nil
}