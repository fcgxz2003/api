//
//     Copyright 2023 The Dragonfly Authors
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
// source: pkg/apis/tfserving/v1/predict.proto

package tfserving

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

// PredictRequest specifies which TensorFlow model to run, as well as
// how inputs are mapped to tensors and how outputs are filtered before
// returning to user.
type PredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Model Specification. If version is not specified, will use the latest
	// (numerical) version.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Input tensors.
	// Names of input tensor are alias names. The mapping from aliases to real
	// input tensor names is stored in the SavedModel export as a prediction
	// SignatureDef under the 'inputs' field.
	Inputs map[string]*TensorProto `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output filter.
	// Names specified are alias names. The mapping from aliases to real output
	// tensor names is stored in the SavedModel export as a prediction
	// SignatureDef under the 'outputs' field.
	// Only tensors specified here will be run/fetched and returned, with the
	// exception that when none is specified, all tensors specified in the
	// named signature will be run/fetched and returned.
	OutputFilter []string `protobuf:"bytes,3,rep,name=output_filter,json=outputFilter,proto3" json:"output_filter,omitempty"`
}

func (x *PredictRequest) Reset() {
	*x = PredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_tfserving_v1_predict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictRequest) ProtoMessage() {}

func (x *PredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_tfserving_v1_predict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictRequest.ProtoReflect.Descriptor instead.
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_tfserving_v1_predict_proto_rawDescGZIP(), []int{0}
}

func (x *PredictRequest) GetModelSpec() *ModelSpec {
	if x != nil {
		return x.ModelSpec
	}
	return nil
}

func (x *PredictRequest) GetInputs() map[string]*TensorProto {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *PredictRequest) GetOutputFilter() []string {
	if x != nil {
		return x.OutputFilter
	}
	return nil
}

// Response for PredictRequest on successful run.
type PredictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Effective Model Specification used to process PredictRequest.
	ModelSpec *ModelSpec `protobuf:"bytes,2,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Output tensors.
	Outputs map[string]*TensorProto `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PredictResponse) Reset() {
	*x = PredictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_tfserving_v1_predict_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictResponse) ProtoMessage() {}

func (x *PredictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_tfserving_v1_predict_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictResponse.ProtoReflect.Descriptor instead.
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_tfserving_v1_predict_proto_rawDescGZIP(), []int{1}
}

func (x *PredictResponse) GetModelSpec() *ModelSpec {
	if x != nil {
		return x.ModelSpec
	}
	return nil
}

func (x *PredictResponse) GetOutputs() map[string]*TensorProto {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_pkg_apis_tfserving_v1_predict_proto protoreflect.FileDescriptor

var file_pkg_apis_tfserving_v1_predict_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x66, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x1a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x66,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x02, 0x0a, 0x0e, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x40, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x54, 0x0a, 0x0b,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xe6, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x66, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x44,
	0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x1a, 0x55, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x64,
	0x37, 0x79, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b,
	0x74, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_apis_tfserving_v1_predict_proto_rawDescOnce sync.Once
	file_pkg_apis_tfserving_v1_predict_proto_rawDescData = file_pkg_apis_tfserving_v1_predict_proto_rawDesc
)

func file_pkg_apis_tfserving_v1_predict_proto_rawDescGZIP() []byte {
	file_pkg_apis_tfserving_v1_predict_proto_rawDescOnce.Do(func() {
		file_pkg_apis_tfserving_v1_predict_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apis_tfserving_v1_predict_proto_rawDescData)
	})
	return file_pkg_apis_tfserving_v1_predict_proto_rawDescData
}

var file_pkg_apis_tfserving_v1_predict_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_apis_tfserving_v1_predict_proto_goTypes = []interface{}{
	(*PredictRequest)(nil),  // 0: tfserving.v1.PredictRequest
	(*PredictResponse)(nil), // 1: tfserving.v1.PredictResponse
	nil,                     // 2: tfserving.v1.PredictRequest.InputsEntry
	nil,                     // 3: tfserving.v1.PredictResponse.OutputsEntry
	(*ModelSpec)(nil),       // 4: tfserving.v1.ModelSpec
	(*TensorProto)(nil),     // 5: tfserving.v1.TensorProto
}
var file_pkg_apis_tfserving_v1_predict_proto_depIdxs = []int32{
	4, // 0: tfserving.v1.PredictRequest.model_spec:type_name -> tfserving.v1.ModelSpec
	2, // 1: tfserving.v1.PredictRequest.inputs:type_name -> tfserving.v1.PredictRequest.InputsEntry
	4, // 2: tfserving.v1.PredictResponse.model_spec:type_name -> tfserving.v1.ModelSpec
	3, // 3: tfserving.v1.PredictResponse.outputs:type_name -> tfserving.v1.PredictResponse.OutputsEntry
	5, // 4: tfserving.v1.PredictRequest.InputsEntry.value:type_name -> tfserving.v1.TensorProto
	5, // 5: tfserving.v1.PredictResponse.OutputsEntry.value:type_name -> tfserving.v1.TensorProto
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_apis_tfserving_v1_predict_proto_init() }
func file_pkg_apis_tfserving_v1_predict_proto_init() {
	if File_pkg_apis_tfserving_v1_predict_proto != nil {
		return
	}
	file_pkg_apis_tfserving_v1_tensor_proto_init()
	file_pkg_apis_tfserving_v1_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_tfserving_v1_predict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictRequest); i {
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
		file_pkg_apis_tfserving_v1_predict_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictResponse); i {
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
			RawDescriptor: file_pkg_apis_tfserving_v1_predict_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_apis_tfserving_v1_predict_proto_goTypes,
		DependencyIndexes: file_pkg_apis_tfserving_v1_predict_proto_depIdxs,
		MessageInfos:      file_pkg_apis_tfserving_v1_predict_proto_msgTypes,
	}.Build()
	File_pkg_apis_tfserving_v1_predict_proto = out.File
	file_pkg_apis_tfserving_v1_predict_proto_rawDesc = nil
	file_pkg_apis_tfserving_v1_predict_proto_goTypes = nil
	file_pkg_apis_tfserving_v1_predict_proto_depIdxs = nil
}