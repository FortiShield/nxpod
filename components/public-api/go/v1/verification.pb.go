// Copyright (c) 2024 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: gitpod/v1/verification.proto

package v1

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

// Required fields:
// - phone_number
type SendPhoneNumberVerificationTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// phone_number in E.164 format
	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *SendPhoneNumberVerificationTokenRequest) Reset() {
	*x = SendPhoneNumberVerificationTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_verification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendPhoneNumberVerificationTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendPhoneNumberVerificationTokenRequest) ProtoMessage() {}

func (x *SendPhoneNumberVerificationTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_verification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendPhoneNumberVerificationTokenRequest.ProtoReflect.Descriptor instead.
func (*SendPhoneNumberVerificationTokenRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_verification_proto_rawDescGZIP(), []int{0}
}

func (x *SendPhoneNumberVerificationTokenRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type SendPhoneNumberVerificationTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// verification_id is used to VerifyPhoneNumberVerificationToken
	VerificationId string `protobuf:"bytes,1,opt,name=verification_id,json=verificationId,proto3" json:"verification_id,omitempty"`
}

func (x *SendPhoneNumberVerificationTokenResponse) Reset() {
	*x = SendPhoneNumberVerificationTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_verification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendPhoneNumberVerificationTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendPhoneNumberVerificationTokenResponse) ProtoMessage() {}

func (x *SendPhoneNumberVerificationTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_verification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendPhoneNumberVerificationTokenResponse.ProtoReflect.Descriptor instead.
func (*SendPhoneNumberVerificationTokenResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_verification_proto_rawDescGZIP(), []int{1}
}

func (x *SendPhoneNumberVerificationTokenResponse) GetVerificationId() string {
	if x != nil {
		return x.VerificationId
	}
	return ""
}

// Required fields:
// - phone_number
// - verification_id
// - token
type VerifyPhoneNumberVerificationTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// phone_number in E.164 format
	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// verification_id is returned by SendPhoneNumberVerificationToken
	VerificationId string `protobuf:"bytes,2,opt,name=verification_id,json=verificationId,proto3" json:"verification_id,omitempty"`
	// token is the verification token from providers
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VerifyPhoneNumberVerificationTokenRequest) Reset() {
	*x = VerifyPhoneNumberVerificationTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_verification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPhoneNumberVerificationTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPhoneNumberVerificationTokenRequest) ProtoMessage() {}

func (x *VerifyPhoneNumberVerificationTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_verification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPhoneNumberVerificationTokenRequest.ProtoReflect.Descriptor instead.
func (*VerifyPhoneNumberVerificationTokenRequest) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_verification_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyPhoneNumberVerificationTokenRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *VerifyPhoneNumberVerificationTokenRequest) GetVerificationId() string {
	if x != nil {
		return x.VerificationId
	}
	return ""
}

func (x *VerifyPhoneNumberVerificationTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyPhoneNumberVerificationTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// verified indicates if the verification was successful
	Verified bool `protobuf:"varint,1,opt,name=verified,proto3" json:"verified,omitempty"`
}

func (x *VerifyPhoneNumberVerificationTokenResponse) Reset() {
	*x = VerifyPhoneNumberVerificationTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitpod_v1_verification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPhoneNumberVerificationTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPhoneNumberVerificationTokenResponse) ProtoMessage() {}

func (x *VerifyPhoneNumberVerificationTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gitpod_v1_verification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPhoneNumberVerificationTokenResponse.ProtoReflect.Descriptor instead.
func (*VerifyPhoneNumberVerificationTokenResponse) Descriptor() ([]byte, []int) {
	return file_gitpod_v1_verification_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyPhoneNumberVerificationTokenResponse) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

var File_gitpod_v1_verification_proto protoreflect.FileDescriptor

var file_gitpod_v1_verification_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x4c, 0x0a, 0x27, 0x53, 0x65, 0x6e,
	0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x53, 0x0a, 0x28, 0x53, 0x65, 0x6e, 0x64, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x8d, 0x01, 0x0a,
	0x29, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a,
	0x0f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x48, 0x0a, 0x2a,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x32, 0xbb, 0x02, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8d,
	0x01, 0x0a, 0x20, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x32, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x93,
	0x01, 0x0a, 0x22, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x34, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x67, 0x69,
	0x74, 0x70, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x51, 0x0a, 0x16, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x70, 0x6f,
	0x64, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f,
	0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gitpod_v1_verification_proto_rawDescOnce sync.Once
	file_gitpod_v1_verification_proto_rawDescData = file_gitpod_v1_verification_proto_rawDesc
)

func file_gitpod_v1_verification_proto_rawDescGZIP() []byte {
	file_gitpod_v1_verification_proto_rawDescOnce.Do(func() {
		file_gitpod_v1_verification_proto_rawDescData = protoimpl.X.CompressGZIP(file_gitpod_v1_verification_proto_rawDescData)
	})
	return file_gitpod_v1_verification_proto_rawDescData
}

var file_gitpod_v1_verification_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gitpod_v1_verification_proto_goTypes = []interface{}{
	(*SendPhoneNumberVerificationTokenRequest)(nil),    // 0: gitpod.v1.SendPhoneNumberVerificationTokenRequest
	(*SendPhoneNumberVerificationTokenResponse)(nil),   // 1: gitpod.v1.SendPhoneNumberVerificationTokenResponse
	(*VerifyPhoneNumberVerificationTokenRequest)(nil),  // 2: gitpod.v1.VerifyPhoneNumberVerificationTokenRequest
	(*VerifyPhoneNumberVerificationTokenResponse)(nil), // 3: gitpod.v1.VerifyPhoneNumberVerificationTokenResponse
}
var file_gitpod_v1_verification_proto_depIdxs = []int32{
	0, // 0: gitpod.v1.VerificationService.SendPhoneNumberVerificationToken:input_type -> gitpod.v1.SendPhoneNumberVerificationTokenRequest
	2, // 1: gitpod.v1.VerificationService.VerifyPhoneNumberVerificationToken:input_type -> gitpod.v1.VerifyPhoneNumberVerificationTokenRequest
	1, // 2: gitpod.v1.VerificationService.SendPhoneNumberVerificationToken:output_type -> gitpod.v1.SendPhoneNumberVerificationTokenResponse
	3, // 3: gitpod.v1.VerificationService.VerifyPhoneNumberVerificationToken:output_type -> gitpod.v1.VerifyPhoneNumberVerificationTokenResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gitpod_v1_verification_proto_init() }
func file_gitpod_v1_verification_proto_init() {
	if File_gitpod_v1_verification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gitpod_v1_verification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendPhoneNumberVerificationTokenRequest); i {
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
		file_gitpod_v1_verification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendPhoneNumberVerificationTokenResponse); i {
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
		file_gitpod_v1_verification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPhoneNumberVerificationTokenRequest); i {
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
		file_gitpod_v1_verification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPhoneNumberVerificationTokenResponse); i {
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
			RawDescriptor: file_gitpod_v1_verification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gitpod_v1_verification_proto_goTypes,
		DependencyIndexes: file_gitpod_v1_verification_proto_depIdxs,
		MessageInfos:      file_gitpod_v1_verification_proto_msgTypes,
	}.Build()
	File_gitpod_v1_verification_proto = out.File
	file_gitpod_v1_verification_proto_rawDesc = nil
	file_gitpod_v1_verification_proto_goTypes = nil
	file_gitpod_v1_verification_proto_depIdxs = nil
}
