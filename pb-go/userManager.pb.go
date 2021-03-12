// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: userManager.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEmailMobile string `protobuf:"bytes,1,opt,name=UserEmailMobile,proto3" json:"UserEmailMobile,omitempty"`
	Password        string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userManager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_userManager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_userManager_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUserEmailMobile() string {
	if x != nil {
		return x.UserEmailMobile
	}
	return ""
}

func (x *UserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID      string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	NewString string `protobuf:"bytes,2,opt,name=NewString,proto3" json:"NewString,omitempty"`
}

func (x *UpdateInfo) Reset() {
	*x = UpdateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userManager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfo) ProtoMessage() {}

func (x *UpdateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_userManager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfo.ProtoReflect.Descriptor instead.
func (*UpdateInfo) Descriptor() ([]byte, []int) {
	return file_userManager_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateInfo) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *UpdateInfo) GetNewString() string {
	if x != nil {
		return x.NewString
	}
	return ""
}

type UpdateAvatar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID     string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	NewBytes []byte `protobuf:"bytes,2,opt,name=NewBytes,proto3" json:"NewBytes,omitempty"`
}

func (x *UpdateAvatar) Reset() {
	*x = UpdateAvatar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userManager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAvatar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAvatar) ProtoMessage() {}

func (x *UpdateAvatar) ProtoReflect() protoreflect.Message {
	mi := &file_userManager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAvatar.ProtoReflect.Descriptor instead.
func (*UpdateAvatar) Descriptor() ([]byte, []int) {
	return file_userManager_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAvatar) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *UpdateAvatar) GetNewBytes() []byte {
	if x != nil {
		return x.NewBytes
	}
	return nil
}

type UserLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Token string `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *UserLoginResponse) Reset() {
	*x = UserLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userManager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginResponse) ProtoMessage() {}

func (x *UserLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userManager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginResponse.ProtoReflect.Descriptor instead.
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return file_userManager_proto_rawDescGZIP(), []int{3}
}

func (x *UserLoginResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UserLoginResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UserLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_userManager_proto protoreflect.FileDescriptor

var file_userManager_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x28, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x65, 0x77,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x65,
	0x77, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x3e, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4e,
	0x65, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x4e,
	0x65, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x4f, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xc4, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x16, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6e, 0x6d, 0x65, 0x12,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x70, 0x62,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userManager_proto_rawDescOnce sync.Once
	file_userManager_proto_rawDescData = file_userManager_proto_rawDesc
)

func file_userManager_proto_rawDescGZIP() []byte {
	file_userManager_proto_rawDescOnce.Do(func() {
		file_userManager_proto_rawDescData = protoimpl.X.CompressGZIP(file_userManager_proto_rawDescData)
	})
	return file_userManager_proto_rawDescData
}

var file_userManager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_userManager_proto_goTypes = []interface{}{
	(*UserInfo)(nil),          // 0: pb.UserInfo
	(*UpdateInfo)(nil),        // 1: pb.UpdateInfo
	(*UpdateAvatar)(nil),      // 2: pb.UpdateAvatar
	(*UserLoginResponse)(nil), // 3: pb.UserLoginResponse
	(*OperationResponse)(nil), // 4: pb.OperationResponse
}
var file_userManager_proto_depIdxs = []int32{
	0, // 0: pb.UserManager.RegisterUserWithUserInfo:input_type -> pb.UserInfo
	0, // 1: pb.UserManager.LoginWithUserLoginInfo:input_type -> pb.UserInfo
	1, // 2: pb.UserManager.UpdateUserNanme:input_type -> pb.UpdateInfo
	1, // 3: pb.UserManager.UpdateUserEmail:input_type -> pb.UpdateInfo
	1, // 4: pb.UserManager.UpdateUserMobile:input_type -> pb.UpdateInfo
	1, // 5: pb.UserManager.UpdateUserPassword:input_type -> pb.UpdateInfo
	2, // 6: pb.UserManager.UpdateUserAvatar:input_type -> pb.UpdateAvatar
	4, // 7: pb.UserManager.RegisterUserWithUserInfo:output_type -> pb.OperationResponse
	3, // 8: pb.UserManager.LoginWithUserLoginInfo:output_type -> pb.UserLoginResponse
	4, // 9: pb.UserManager.UpdateUserNanme:output_type -> pb.OperationResponse
	4, // 10: pb.UserManager.UpdateUserEmail:output_type -> pb.OperationResponse
	4, // 11: pb.UserManager.UpdateUserMobile:output_type -> pb.OperationResponse
	4, // 12: pb.UserManager.UpdateUserPassword:output_type -> pb.OperationResponse
	4, // 13: pb.UserManager.UpdateUserAvatar:output_type -> pb.OperationResponse
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userManager_proto_init() }
func file_userManager_proto_init() {
	if File_userManager_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_userManager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_userManager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfo); i {
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
		file_userManager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAvatar); i {
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
		file_userManager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginResponse); i {
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
			RawDescriptor: file_userManager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userManager_proto_goTypes,
		DependencyIndexes: file_userManager_proto_depIdxs,
		MessageInfos:      file_userManager_proto_msgTypes,
	}.Build()
	File_userManager_proto = out.File
	file_userManager_proto_rawDesc = nil
	file_userManager_proto_goTypes = nil
	file_userManager_proto_depIdxs = nil
}
