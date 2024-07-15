//
//     Copyright 2022 The Dragonfly Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: pkg/apis/security/v1/security.proto

package security

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Certificate request type.
// Dragonfly supports peers authentication with Mutual TLS(mTLS)
// For mTLS, all peers need to request TLS certificates for communicating
// The server side may overwrite ant requested certificate filed based on its policies.
type CertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ASN.1 DER form certificate request.
	// The public key in the CSR is used to generate the certificate,
	// and other fields in the generated certificate may be overwritten by the CA.
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	// Optional: requested certificate validity period.
	ValidityPeriod *durationpb.Duration `protobuf:"bytes,2,opt,name=validity_period,json=validityPeriod,proto3" json:"validity_period,omitempty"`
}

func (x *CertificateRequest) Reset() {
	*x = CertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_security_v1_security_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateRequest) ProtoMessage() {}

func (x *CertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_security_v1_security_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateRequest.ProtoReflect.Descriptor instead.
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_security_v1_security_proto_rawDescGZIP(), []int{0}
}

func (x *CertificateRequest) GetCsr() []byte {
	if x != nil {
		return x.Csr
	}
	return nil
}

func (x *CertificateRequest) GetValidityPeriod() *durationpb.Duration {
	if x != nil {
		return x.ValidityPeriod
	}
	return nil
}

// Certificate response type.
type CertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ASN.1 DER form certificate chain.
	CertificateChain [][]byte `protobuf:"bytes,1,rep,name=certificate_chain,json=certificateChain,proto3" json:"certificate_chain,omitempty"`
}

func (x *CertificateResponse) Reset() {
	*x = CertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_security_v1_security_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateResponse) ProtoMessage() {}

func (x *CertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_security_v1_security_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateResponse.ProtoReflect.Descriptor instead.
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_security_v1_security_proto_rawDescGZIP(), []int{1}
}

func (x *CertificateResponse) GetCertificateChain() [][]byte {
	if x != nil {
		return x.CertificateChain
	}
	return nil
}

var File_pkg_apis_security_v1_security_proto protoreflect.FileDescriptor

var file_pkg_apis_security_v1_security_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x12, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x03, 0x63, 0x73, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x7a, 0x02, 0x10, 0x01, 0x52, 0x03, 0x63, 0x73, 0x72, 0x12, 0x4c, 0x0a, 0x0f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x4c, 0x0a, 0x13, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01,
	0x02, 0x08, 0x01, 0x52, 0x10, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x32, 0x60, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x63, 0x67, 0x78, 0x7a, 0x32, 0x30, 0x30, 0x33, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_apis_security_v1_security_proto_rawDescOnce sync.Once
	file_pkg_apis_security_v1_security_proto_rawDescData = file_pkg_apis_security_v1_security_proto_rawDesc
)

func file_pkg_apis_security_v1_security_proto_rawDescGZIP() []byte {
	file_pkg_apis_security_v1_security_proto_rawDescOnce.Do(func() {
		file_pkg_apis_security_v1_security_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apis_security_v1_security_proto_rawDescData)
	})
	return file_pkg_apis_security_v1_security_proto_rawDescData
}

var file_pkg_apis_security_v1_security_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_apis_security_v1_security_proto_goTypes = []interface{}{
	(*CertificateRequest)(nil),  // 0: security.CertificateRequest
	(*CertificateResponse)(nil), // 1: security.CertificateResponse
	(*durationpb.Duration)(nil), // 2: google.protobuf.Duration
}
var file_pkg_apis_security_v1_security_proto_depIdxs = []int32{
	2, // 0: security.CertificateRequest.validity_period:type_name -> google.protobuf.Duration
	0, // 1: security.Certificate.IssueCertificate:input_type -> security.CertificateRequest
	1, // 2: security.Certificate.IssueCertificate:output_type -> security.CertificateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_apis_security_v1_security_proto_init() }
func file_pkg_apis_security_v1_security_proto_init() {
	if File_pkg_apis_security_v1_security_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_security_v1_security_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateRequest); i {
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
		file_pkg_apis_security_v1_security_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateResponse); i {
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
			RawDescriptor: file_pkg_apis_security_v1_security_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_apis_security_v1_security_proto_goTypes,
		DependencyIndexes: file_pkg_apis_security_v1_security_proto_depIdxs,
		MessageInfos:      file_pkg_apis_security_v1_security_proto_msgTypes,
	}.Build()
	File_pkg_apis_security_v1_security_proto = out.File
	file_pkg_apis_security_v1_security_proto_rawDesc = nil
	file_pkg_apis_security_v1_security_proto_goTypes = nil
	file_pkg_apis_security_v1_security_proto_depIdxs = nil
}
