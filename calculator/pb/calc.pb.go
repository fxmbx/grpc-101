// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: calc.proto

package calculator

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

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumer   int32 `protobuf:"varint,1,opt,name=first_numer,json=firstNumer,proto3" json:"first_numer,omitempty"`
	SecondNumber int32 `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calc_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetFirstNumer() int32 {
	if x != nil {
		return x.FirstNumer
	}
	return 0
}

func (x *SumRequest) GetSecondNumber() int32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calc_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type PrimeNumberDecomposerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	K      int32 `protobuf:"varint,2,opt,name=k,proto3" json:"k,omitempty"`
}

func (x *PrimeNumberDecomposerRequest) Reset() {
	*x = PrimeNumberDecomposerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberDecomposerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberDecomposerRequest) ProtoMessage() {}

func (x *PrimeNumberDecomposerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberDecomposerRequest.ProtoReflect.Descriptor instead.
func (*PrimeNumberDecomposerRequest) Descriptor() ([]byte, []int) {
	return file_calc_proto_rawDescGZIP(), []int{2}
}

func (x *PrimeNumberDecomposerRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *PrimeNumberDecomposerRequest) GetK() int32 {
	if x != nil {
		return x.K
	}
	return 0
}

type PrimeNumberDecomposeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimeNumberDecomposeResponse) Reset() {
	*x = PrimeNumberDecomposeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberDecomposeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberDecomposeResponse) ProtoMessage() {}

func (x *PrimeNumberDecomposeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberDecomposeResponse.ProtoReflect.Descriptor instead.
func (*PrimeNumberDecomposeResponse) Descriptor() ([]byte, []int) {
	return file_calc_proto_rawDescGZIP(), []int{3}
}

func (x *PrimeNumberDecomposeResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calc_proto protoreflect.FileDescriptor

var file_calc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x52, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x44, 0x0a, 0x1c, 0x50,
	0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x6b, 0x22, 0x36, 0x0a, 0x1c, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x9e, 0x01, 0x0a, 0x11, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x28, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x15, 0x50, 0x72, 0x69,
	0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x65, 0x72, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x78, 0x6d, 0x62, 0x78, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x31, 0x30, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calc_proto_rawDescOnce sync.Once
	file_calc_proto_rawDescData = file_calc_proto_rawDesc
)

func file_calc_proto_rawDescGZIP() []byte {
	file_calc_proto_rawDescOnce.Do(func() {
		file_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calc_proto_rawDescData)
	})
	return file_calc_proto_rawDescData
}

var file_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_calc_proto_goTypes = []interface{}{
	(*SumRequest)(nil),                   // 0: pb.SumRequest
	(*SumResponse)(nil),                  // 1: pb.SumResponse
	(*PrimeNumberDecomposerRequest)(nil), // 2: pb.PrimeNumberDecomposerRequest
	(*PrimeNumberDecomposeResponse)(nil), // 3: pb.PrimeNumberDecomposeResponse
}
var file_calc_proto_depIdxs = []int32{
	0, // 0: pb.CalculatorService.Sum:input_type -> pb.SumRequest
	2, // 1: pb.CalculatorService.PrimeNumberDecomposer:input_type -> pb.PrimeNumberDecomposerRequest
	1, // 2: pb.CalculatorService.Sum:output_type -> pb.SumResponse
	3, // 3: pb.CalculatorService.PrimeNumberDecomposer:output_type -> pb.PrimeNumberDecomposeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calc_proto_init() }
func file_calc_proto_init() {
	if File_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberDecomposerRequest); i {
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
		file_calc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberDecomposeResponse); i {
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
			RawDescriptor: file_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calc_proto_goTypes,
		DependencyIndexes: file_calc_proto_depIdxs,
		MessageInfos:      file_calc_proto_msgTypes,
	}.Build()
	File_calc_proto = out.File
	file_calc_proto_rawDesc = nil
	file_calc_proto_goTypes = nil
	file_calc_proto_depIdxs = nil
}
