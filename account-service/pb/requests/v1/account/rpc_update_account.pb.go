// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: requests/v1/account/rpc_update_account.proto

package account

import (
	models "github.com/nicodanke/inventapp_v3/account-service/pb/models"
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

type UpdateAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyName   *string                `protobuf:"bytes,2,opt,name=companyName,proto3,oneof" json:"companyName,omitempty"`
	Email         *string                `protobuf:"bytes,3,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Active        *bool                  `protobuf:"varint,4,opt,name=active,proto3,oneof" json:"active,omitempty"`
	Phone         *string                `protobuf:"bytes,5,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	WebUrl        *string                `protobuf:"bytes,6,opt,name=webUrl,proto3,oneof" json:"webUrl,omitempty"`
	Country       *string                `protobuf:"bytes,7,opt,name=country,proto3,oneof" json:"country,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAccountRequest) Reset() {
	*x = UpdateAccountRequest{}
	mi := &file_requests_v1_account_rpc_update_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountRequest) ProtoMessage() {}

func (x *UpdateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requests_v1_account_rpc_update_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return file_requests_v1_account_rpc_update_account_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateAccountRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAccountRequest) GetCompanyName() string {
	if x != nil && x.CompanyName != nil {
		return *x.CompanyName
	}
	return ""
}

func (x *UpdateAccountRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UpdateAccountRequest) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

func (x *UpdateAccountRequest) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *UpdateAccountRequest) GetWebUrl() string {
	if x != nil && x.WebUrl != nil {
		return *x.WebUrl
	}
	return ""
}

func (x *UpdateAccountRequest) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

type UpdateAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Account       *models.Account        `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAccountResponse) Reset() {
	*x = UpdateAccountResponse{}
	mi := &file_requests_v1_account_rpc_update_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountResponse) ProtoMessage() {}

func (x *UpdateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_requests_v1_account_rpc_update_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountResponse.ProtoReflect.Descriptor instead.
func (*UpdateAccountResponse) Descriptor() ([]byte, []int) {
	return file_requests_v1_account_rpc_update_account_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateAccountResponse) GetAccount() *models.Account {
	if x != nil {
		return x.Account
	}
	return nil
}

var File_requests_v1_account_rpc_update_account_proto protoreflect.FileDescriptor

var file_requests_v1_account_rpc_update_account_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f, 0x76, 0x33, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x1a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x02, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x25, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x02, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x62,
	0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x06, 0x77, 0x65, 0x62,
	0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x77, 0x65, 0x62, 0x55, 0x72, 0x6c, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x5f, 0x0a, 0x15, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70,
	0x70, 0x5f, 0x76, 0x33, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x4a, 0x5a, 0x48,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x63, 0x6f, 0x64,
	0x61, 0x6e, 0x6b, 0x65, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x5f, 0x76,
	0x33, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_requests_v1_account_rpc_update_account_proto_rawDescOnce sync.Once
	file_requests_v1_account_rpc_update_account_proto_rawDescData = file_requests_v1_account_rpc_update_account_proto_rawDesc
)

func file_requests_v1_account_rpc_update_account_proto_rawDescGZIP() []byte {
	file_requests_v1_account_rpc_update_account_proto_rawDescOnce.Do(func() {
		file_requests_v1_account_rpc_update_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_requests_v1_account_rpc_update_account_proto_rawDescData)
	})
	return file_requests_v1_account_rpc_update_account_proto_rawDescData
}

var file_requests_v1_account_rpc_update_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_requests_v1_account_rpc_update_account_proto_goTypes = []any{
	(*UpdateAccountRequest)(nil),  // 0: inventapp_v3.account_service.requests.v1.account.UpdateAccountRequest
	(*UpdateAccountResponse)(nil), // 1: inventapp_v3.account_service.requests.v1.account.UpdateAccountResponse
	(*models.Account)(nil),        // 2: inventapp_v3.account_service.models.Account
}
var file_requests_v1_account_rpc_update_account_proto_depIdxs = []int32{
	2, // 0: inventapp_v3.account_service.requests.v1.account.UpdateAccountResponse.account:type_name -> inventapp_v3.account_service.models.Account
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_requests_v1_account_rpc_update_account_proto_init() }
func file_requests_v1_account_rpc_update_account_proto_init() {
	if File_requests_v1_account_rpc_update_account_proto != nil {
		return
	}
	file_requests_v1_account_rpc_update_account_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_requests_v1_account_rpc_update_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_requests_v1_account_rpc_update_account_proto_goTypes,
		DependencyIndexes: file_requests_v1_account_rpc_update_account_proto_depIdxs,
		MessageInfos:      file_requests_v1_account_rpc_update_account_proto_msgTypes,
	}.Build()
	File_requests_v1_account_rpc_update_account_proto = out.File
	file_requests_v1_account_rpc_update_account_proto_rawDesc = nil
	file_requests_v1_account_rpc_update_account_proto_goTypes = nil
	file_requests_v1_account_rpc_update_account_proto_depIdxs = nil
}
