// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: grpc/notifier/notifier.proto

package notifier

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

type SendInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receivers []string          `protobuf:"bytes,1,rep,name=receivers,proto3" json:"receivers,omitempty"`
	Action    string            `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Data      map[string]string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SendInput) Reset() {
	*x = SendInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_notifier_notifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendInput) ProtoMessage() {}

func (x *SendInput) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_notifier_notifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendInput.ProtoReflect.Descriptor instead.
func (*SendInput) Descriptor() ([]byte, []int) {
	return file_grpc_notifier_notifier_proto_rawDescGZIP(), []int{0}
}

func (x *SendInput) GetReceivers() []string {
	if x != nil {
		return x.Receivers
	}
	return nil
}

func (x *SendInput) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *SendInput) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SendResult) Reset() {
	*x = SendResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_notifier_notifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResult) ProtoMessage() {}

func (x *SendResult) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_notifier_notifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResult.ProtoReflect.Descriptor instead.
func (*SendResult) Descriptor() ([]byte, []int) {
	return file_grpc_notifier_notifier_proto_rawDescGZIP(), []int{1}
}

func (x *SendResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_grpc_notifier_notifier_proto protoreflect.FileDescriptor

var file_grpc_notifier_notifier_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4,
	0x01, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x24, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x2d, 0x0a, 0x08, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12,
	0x0a, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0b, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_notifier_notifier_proto_rawDescOnce sync.Once
	file_grpc_notifier_notifier_proto_rawDescData = file_grpc_notifier_notifier_proto_rawDesc
)

func file_grpc_notifier_notifier_proto_rawDescGZIP() []byte {
	file_grpc_notifier_notifier_proto_rawDescOnce.Do(func() {
		file_grpc_notifier_notifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_notifier_notifier_proto_rawDescData)
	})
	return file_grpc_notifier_notifier_proto_rawDescData
}

var file_grpc_notifier_notifier_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_notifier_notifier_proto_goTypes = []interface{}{
	(*SendInput)(nil),  // 0: SendInput
	(*SendResult)(nil), // 1: SendResult
	nil,                // 2: SendInput.DataEntry
}
var file_grpc_notifier_notifier_proto_depIdxs = []int32{
	2, // 0: SendInput.data:type_name -> SendInput.DataEntry
	0, // 1: Notifier.Send:input_type -> SendInput
	1, // 2: Notifier.Send:output_type -> SendResult
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_notifier_notifier_proto_init() }
func file_grpc_notifier_notifier_proto_init() {
	if File_grpc_notifier_notifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_notifier_notifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendInput); i {
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
		file_grpc_notifier_notifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResult); i {
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
			RawDescriptor: file_grpc_notifier_notifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_notifier_notifier_proto_goTypes,
		DependencyIndexes: file_grpc_notifier_notifier_proto_depIdxs,
		MessageInfos:      file_grpc_notifier_notifier_proto_msgTypes,
	}.Build()
	File_grpc_notifier_notifier_proto = out.File
	file_grpc_notifier_notifier_proto_rawDesc = nil
	file_grpc_notifier_notifier_proto_goTypes = nil
	file_grpc_notifier_notifier_proto_depIdxs = nil
}
