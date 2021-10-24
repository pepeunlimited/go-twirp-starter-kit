// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: resources/withdrawal.proto

package resources

import (
	enums "github.com/pepeunlimited/go-twirp-starter-kit/pkg/api/v1/enums"
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

type Withdrawal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status     enums.WithdrawalStatus `protobuf:"varint,2,opt,name=status,proto3,enum=pepeunlimited.twirpkit.enums.WithdrawalStatus" json:"status,omitempty"`
	Amount     int64                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	WorkflowId string                 `protobuf:"bytes,4,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	RunId      string                 `protobuf:"bytes,5,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *Withdrawal) Reset() {
	*x = Withdrawal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_withdrawal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Withdrawal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Withdrawal) ProtoMessage() {}

func (x *Withdrawal) ProtoReflect() protoreflect.Message {
	mi := &file_resources_withdrawal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Withdrawal.ProtoReflect.Descriptor instead.
func (*Withdrawal) Descriptor() ([]byte, []int) {
	return file_resources_withdrawal_proto_rawDescGZIP(), []int{0}
}

func (x *Withdrawal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Withdrawal) GetStatus() enums.WithdrawalStatus {
	if x != nil {
		return x.Status
	}
	return enums.WithdrawalStatus(0)
}

func (x *Withdrawal) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Withdrawal) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *Withdrawal) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

var File_resources_withdrawal_proto protoreflect.FileDescriptor

var file_resources_withdrawal_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x70, 0x65,
	0x70, 0x65, 0x75, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x2e, 0x74, 0x77, 0x69, 0x72,
	0x70, 0x6b, 0x69, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1d,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01,
	0x0a, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x46, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x70,
	0x65, 0x70, 0x65, 0x75, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x2e, 0x74, 0x77, 0x69,
	0x72, 0x70, 0x6b, 0x69, 0x74, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72,
	0x75, 0x6e, 0x49, 0x64, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x70, 0x65, 0x75, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64,
	0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x72, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_resources_withdrawal_proto_rawDescOnce sync.Once
	file_resources_withdrawal_proto_rawDescData = file_resources_withdrawal_proto_rawDesc
)

func file_resources_withdrawal_proto_rawDescGZIP() []byte {
	file_resources_withdrawal_proto_rawDescOnce.Do(func() {
		file_resources_withdrawal_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_withdrawal_proto_rawDescData)
	})
	return file_resources_withdrawal_proto_rawDescData
}

var file_resources_withdrawal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_withdrawal_proto_goTypes = []interface{}{
	(*Withdrawal)(nil),          // 0: pepeunlimited.twirpkit.resources.Withdrawal
	(enums.WithdrawalStatus)(0), // 1: pepeunlimited.twirpkit.enums.WithdrawalStatus
}
var file_resources_withdrawal_proto_depIdxs = []int32{
	1, // 0: pepeunlimited.twirpkit.resources.Withdrawal.status:type_name -> pepeunlimited.twirpkit.enums.WithdrawalStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_resources_withdrawal_proto_init() }
func file_resources_withdrawal_proto_init() {
	if File_resources_withdrawal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_withdrawal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Withdrawal); i {
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
			RawDescriptor: file_resources_withdrawal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_withdrawal_proto_goTypes,
		DependencyIndexes: file_resources_withdrawal_proto_depIdxs,
		MessageInfos:      file_resources_withdrawal_proto_msgTypes,
	}.Build()
	File_resources_withdrawal_proto = out.File
	file_resources_withdrawal_proto_rawDesc = nil
	file_resources_withdrawal_proto_goTypes = nil
	file_resources_withdrawal_proto_depIdxs = nil
}
