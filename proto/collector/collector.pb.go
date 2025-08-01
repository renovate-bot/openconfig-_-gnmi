// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: github.com/openconfig/gnmi/proto/collector/collector.proto

package gnmi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReconnectRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Target        []string               `protobuf:"bytes,1,rep,name=target,proto3" json:"target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReconnectRequest) Reset() {
	*x = ReconnectRequest{}
	mi := &file_github_com_openconfig_gnmi_proto_collector_collector_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconnectRequest) ProtoMessage() {}

func (x *ReconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnmi_proto_collector_collector_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconnectRequest.ProtoReflect.Descriptor instead.
func (*ReconnectRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDescGZIP(), []int{0}
}

func (x *ReconnectRequest) GetTarget() []string {
	if x != nil {
		return x.Target
	}
	return nil
}

type Nil struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Nil) Reset() {
	*x = Nil{}
	mi := &file_github_com_openconfig_gnmi_proto_collector_collector_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Nil) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nil) ProtoMessage() {}

func (x *Nil) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnmi_proto_collector_collector_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nil.ProtoReflect.Descriptor instead.
func (*Nil) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDescGZIP(), []int{1}
}

var File_github_com_openconfig_gnmi_proto_collector_collector_proto protoreflect.FileDescriptor

const file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDesc = "" +
	"\n" +
	":github.com/openconfig/gnmi/proto/collector/collector.proto\x12\x04gnmi\"*\n" +
	"\x10ReconnectRequest\x12\x16\n" +
	"\x06target\x18\x01 \x03(\tR\x06target\"\x05\n" +
	"\x03Nil2=\n" +
	"\tCollector\x120\n" +
	"\tReconnect\x12\x16.gnmi.ReconnectRequest\x1a\t.gnmi.Nil\"\x00B1Z/github.com/openconfig/gnmi/proto/collector;gnmib\x06proto3"

var (
	file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDescOnce sync.Once
	file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDescData []byte
)

func file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDescGZIP() []byte {
	file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDescOnce.Do(func() {
		file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDesc), len(file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDesc)))
	})
	return file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDescData
}

var file_github_com_openconfig_gnmi_proto_collector_collector_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_openconfig_gnmi_proto_collector_collector_proto_goTypes = []any{
	(*ReconnectRequest)(nil), // 0: gnmi.ReconnectRequest
	(*Nil)(nil),              // 1: gnmi.Nil
}
var file_github_com_openconfig_gnmi_proto_collector_collector_proto_depIdxs = []int32{
	0, // 0: gnmi.Collector.Reconnect:input_type -> gnmi.ReconnectRequest
	1, // 1: gnmi.Collector.Reconnect:output_type -> gnmi.Nil
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_openconfig_gnmi_proto_collector_collector_proto_init() }
func file_github_com_openconfig_gnmi_proto_collector_collector_proto_init() {
	if File_github_com_openconfig_gnmi_proto_collector_collector_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDesc), len(file_github_com_openconfig_gnmi_proto_collector_collector_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_openconfig_gnmi_proto_collector_collector_proto_goTypes,
		DependencyIndexes: file_github_com_openconfig_gnmi_proto_collector_collector_proto_depIdxs,
		MessageInfos:      file_github_com_openconfig_gnmi_proto_collector_collector_proto_msgTypes,
	}.Build()
	File_github_com_openconfig_gnmi_proto_collector_collector_proto = out.File
	file_github_com_openconfig_gnmi_proto_collector_collector_proto_goTypes = nil
	file_github_com_openconfig_gnmi_proto_collector_collector_proto_depIdxs = nil
}
