// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: grpc/increment.proto

package incrementValue

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

type BeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BeatRequest) Reset() {
	*x = BeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_increment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeatRequest) ProtoMessage() {}

func (x *BeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_increment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeatRequest.ProtoReflect.Descriptor instead.
func (*BeatRequest) Descriptor() ([]byte, []int) {
	return file_grpc_increment_proto_rawDescGZIP(), []int{0}
}

func (x *BeatRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BeatAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port string `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *BeatAck) Reset() {
	*x = BeatAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_increment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeatAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeatAck) ProtoMessage() {}

func (x *BeatAck) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_increment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeatAck.ProtoReflect.Descriptor instead.
func (*BeatAck) Descriptor() ([]byte, []int) {
	return file_grpc_increment_proto_rawDescGZIP(), []int{1}
}

func (x *BeatAck) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_increment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_increment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_grpc_increment_proto_rawDescGZIP(), []int{2}
}

func (x *AddRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response int32 `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_increment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_increment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grpc_increment_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetResponse() int32 {
	if x != nil {
		return x.Response
	}
	return 0
}

var File_grpc_increment_proto protoreflect.FileDescriptor

var file_grpc_increment_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x1d, 0x0a, 0x07, 0x42, 0x65, 0x61, 0x74, 0x41, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x22,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x26, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe3, 0x01, 0x0a, 0x0e, 0x49,
	0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x43, 0x0a,
	0x09, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x69, 0x6e, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65,
	0x61, 0x74, 0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x2e, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x2e, 0x42, 0x65, 0x61, 0x74, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x54, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x2e, 0x69, 0x6e, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x42, 0x5a, 0x40, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x76, 0x69, 0x68, 0x61, 0x6e, 0x6e, 0x69,
	0x6e, 0x65, 0x6e, 0x2f, 0x44, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x67, 0x69, 0x74, 0x3b, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_increment_proto_rawDescOnce sync.Once
	file_grpc_increment_proto_rawDescData = file_grpc_increment_proto_rawDesc
)

func file_grpc_increment_proto_rawDescGZIP() []byte {
	file_grpc_increment_proto_rawDescOnce.Do(func() {
		file_grpc_increment_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_increment_proto_rawDescData)
	})
	return file_grpc_increment_proto_rawDescData
}

var file_grpc_increment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_increment_proto_goTypes = []interface{}{
	(*BeatRequest)(nil), // 0: incrementValue.BeatRequest
	(*BeatAck)(nil),     // 1: incrementValue.BeatAck
	(*AddRequest)(nil),  // 2: incrementValue.AddRequest
	(*Response)(nil),    // 3: incrementValue.Response
}
var file_grpc_increment_proto_depIdxs = []int32{
	2, // 0: incrementValue.IncrementValue.Increment:input_type -> incrementValue.AddRequest
	0, // 1: incrementValue.IncrementValue.GetHeartBeat:input_type -> incrementValue.BeatRequest
	2, // 2: incrementValue.IncrementValue.AddToValue:input_type -> incrementValue.AddRequest
	3, // 3: incrementValue.IncrementValue.Increment:output_type -> incrementValue.Response
	1, // 4: incrementValue.IncrementValue.GetHeartBeat:output_type -> incrementValue.BeatAck
	3, // 5: incrementValue.IncrementValue.AddToValue:output_type -> incrementValue.Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_increment_proto_init() }
func file_grpc_increment_proto_init() {
	if File_grpc_increment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_increment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeatRequest); i {
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
		file_grpc_increment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeatAck); i {
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
		file_grpc_increment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_grpc_increment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_grpc_increment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_increment_proto_goTypes,
		DependencyIndexes: file_grpc_increment_proto_depIdxs,
		MessageInfos:      file_grpc_increment_proto_msgTypes,
	}.Build()
	File_grpc_increment_proto = out.File
	file_grpc_increment_proto_rawDesc = nil
	file_grpc_increment_proto_goTypes = nil
	file_grpc_increment_proto_depIdxs = nil
}