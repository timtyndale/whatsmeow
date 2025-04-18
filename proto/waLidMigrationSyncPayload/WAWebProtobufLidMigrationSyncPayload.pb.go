// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: waLidMigrationSyncPayload/WAWebProtobufLidMigrationSyncPayload.proto

package waLidMigrationSyncPayload

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

type LIDMigrationMapping struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pn            *uint64                `protobuf:"varint,1,req,name=pn" json:"pn,omitempty"`
	AssignedLid   *uint64                `protobuf:"varint,2,req,name=assignedLid" json:"assignedLid,omitempty"`
	LatestLid     *uint64                `protobuf:"varint,3,opt,name=latestLid" json:"latestLid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LIDMigrationMapping) Reset() {
	*x = LIDMigrationMapping{}
	mi := &file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LIDMigrationMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LIDMigrationMapping) ProtoMessage() {}

func (x *LIDMigrationMapping) ProtoReflect() protoreflect.Message {
	mi := &file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LIDMigrationMapping.ProtoReflect.Descriptor instead.
func (*LIDMigrationMapping) Descriptor() ([]byte, []int) {
	return file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDescGZIP(), []int{0}
}

func (x *LIDMigrationMapping) GetPn() uint64 {
	if x != nil && x.Pn != nil {
		return *x.Pn
	}
	return 0
}

func (x *LIDMigrationMapping) GetAssignedLid() uint64 {
	if x != nil && x.AssignedLid != nil {
		return *x.AssignedLid
	}
	return 0
}

func (x *LIDMigrationMapping) GetLatestLid() uint64 {
	if x != nil && x.LatestLid != nil {
		return *x.LatestLid
	}
	return 0
}

type LIDMigrationMappingSyncPayload struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	PnToLidMappings          []*LIDMigrationMapping `protobuf:"bytes,1,rep,name=pnToLidMappings" json:"pnToLidMappings,omitempty"`
	ChatDbMigrationTimestamp *uint64                `protobuf:"varint,2,opt,name=chatDbMigrationTimestamp" json:"chatDbMigrationTimestamp,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *LIDMigrationMappingSyncPayload) Reset() {
	*x = LIDMigrationMappingSyncPayload{}
	mi := &file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LIDMigrationMappingSyncPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LIDMigrationMappingSyncPayload) ProtoMessage() {}

func (x *LIDMigrationMappingSyncPayload) ProtoReflect() protoreflect.Message {
	mi := &file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LIDMigrationMappingSyncPayload.ProtoReflect.Descriptor instead.
func (*LIDMigrationMappingSyncPayload) Descriptor() ([]byte, []int) {
	return file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDescGZIP(), []int{1}
}

func (x *LIDMigrationMappingSyncPayload) GetPnToLidMappings() []*LIDMigrationMapping {
	if x != nil {
		return x.PnToLidMappings
	}
	return nil
}

func (x *LIDMigrationMappingSyncPayload) GetChatDbMigrationTimestamp() uint64 {
	if x != nil && x.ChatDbMigrationTimestamp != nil {
		return *x.ChatDbMigrationTimestamp
	}
	return 0
}

var File_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto protoreflect.FileDescriptor

const file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDesc = "" +
	"\n" +
	"DwaLidMigrationSyncPayload/WAWebProtobufLidMigrationSyncPayload.proto\x12$WAWebProtobufLidMigrationSyncPayload\"e\n" +
	"\x13LIDMigrationMapping\x12\x0e\n" +
	"\x02pn\x18\x01 \x02(\x04R\x02pn\x12 \n" +
	"\vassignedLid\x18\x02 \x02(\x04R\vassignedLid\x12\x1c\n" +
	"\tlatestLid\x18\x03 \x01(\x04R\tlatestLid\"\xc1\x01\n" +
	"\x1eLIDMigrationMappingSyncPayload\x12c\n" +
	"\x0fpnToLidMappings\x18\x01 \x03(\v29.WAWebProtobufLidMigrationSyncPayload.LIDMigrationMappingR\x0fpnToLidMappings\x12:\n" +
	"\x18chatDbMigrationTimestamp\x18\x02 \x01(\x04R\x18chatDbMigrationTimestampBAZ?github.com/timtyndale/whatsmeow/proto/waLidMigrationSyncPayload"

var (
	file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDescOnce sync.Once
	file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDescData []byte
)

func file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDescGZIP() []byte {
	file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDescOnce.Do(func() {
		file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDesc), len(file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDesc)))
	})
	return file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDescData
}

var file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_goTypes = []any{
	(*LIDMigrationMapping)(nil),            // 0: WAWebProtobufLidMigrationSyncPayload.LIDMigrationMapping
	(*LIDMigrationMappingSyncPayload)(nil), // 1: WAWebProtobufLidMigrationSyncPayload.LIDMigrationMappingSyncPayload
}
var file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_depIdxs = []int32{
	0, // 0: WAWebProtobufLidMigrationSyncPayload.LIDMigrationMappingSyncPayload.pnToLidMappings:type_name -> WAWebProtobufLidMigrationSyncPayload.LIDMigrationMapping
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_init() }
func file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_init() {
	if File_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDesc), len(file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_goTypes,
		DependencyIndexes: file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_depIdxs,
		MessageInfos:      file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_msgTypes,
	}.Build()
	File_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto = out.File
	file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_goTypes = nil
	file_waLidMigrationSyncPayload_WAWebProtobufLidMigrationSyncPayload_proto_depIdxs = nil
}
