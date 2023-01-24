// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/maps/routes/v1/compute_routes_response.proto

package routes

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ComputeRoutes the response message.
type ComputeRoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains an array of computed routes (up to three) when you specify
	// compute_alternatives_routes, and contains just one route when you don't.
	// When this array contains multiple entries, the first one is the most
	// recommended route. If the array is empty, then it means no route could be
	// found.
	Routes []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	// In some cases when the server is not able to compute the route results with
	// all of the input preferences, it may fallback to using a different way of
	// computation. When fallback mode is used, this field contains detailed info
	// about the fallback response. Otherwise this field is unset.
	FallbackInfo *FallbackInfo `protobuf:"bytes,2,opt,name=fallback_info,json=fallbackInfo,proto3" json:"fallback_info,omitempty"`
}

func (x *ComputeRoutesResponse) Reset() {
	*x = ComputeRoutesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routes_v1_compute_routes_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeRoutesResponse) ProtoMessage() {}

func (x *ComputeRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routes_v1_compute_routes_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeRoutesResponse.ProtoReflect.Descriptor instead.
func (*ComputeRoutesResponse) Descriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_compute_routes_response_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeRoutesResponse) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *ComputeRoutesResponse) GetFallbackInfo() *FallbackInfo {
	if x != nil {
		return x.FallbackInfo
	}
	return nil
}

var File_google_maps_routes_v1_compute_routes_response_proto protoreflect.FileDescriptor

var file_google_maps_routes_v1_compute_routes_response_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61,
	0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x29, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x15, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61,
	0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x0d, 0x66, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0xb0, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x1a, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x04, 0x47, 0x4d, 0x52, 0x53, 0xaa, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routes_v1_compute_routes_response_proto_rawDescOnce sync.Once
	file_google_maps_routes_v1_compute_routes_response_proto_rawDescData = file_google_maps_routes_v1_compute_routes_response_proto_rawDesc
)

func file_google_maps_routes_v1_compute_routes_response_proto_rawDescGZIP() []byte {
	file_google_maps_routes_v1_compute_routes_response_proto_rawDescOnce.Do(func() {
		file_google_maps_routes_v1_compute_routes_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routes_v1_compute_routes_response_proto_rawDescData)
	})
	return file_google_maps_routes_v1_compute_routes_response_proto_rawDescData
}

var file_google_maps_routes_v1_compute_routes_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_maps_routes_v1_compute_routes_response_proto_goTypes = []interface{}{
	(*ComputeRoutesResponse)(nil), // 0: google.maps.routes.v1.ComputeRoutesResponse
	(*Route)(nil),                 // 1: google.maps.routes.v1.Route
	(*FallbackInfo)(nil),          // 2: google.maps.routes.v1.FallbackInfo
}
var file_google_maps_routes_v1_compute_routes_response_proto_depIdxs = []int32{
	1, // 0: google.maps.routes.v1.ComputeRoutesResponse.routes:type_name -> google.maps.routes.v1.Route
	2, // 1: google.maps.routes.v1.ComputeRoutesResponse.fallback_info:type_name -> google.maps.routes.v1.FallbackInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_maps_routes_v1_compute_routes_response_proto_init() }
func file_google_maps_routes_v1_compute_routes_response_proto_init() {
	if File_google_maps_routes_v1_compute_routes_response_proto != nil {
		return
	}
	file_google_maps_routes_v1_fallback_info_proto_init()
	file_google_maps_routes_v1_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_maps_routes_v1_compute_routes_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeRoutesResponse); i {
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
			RawDescriptor: file_google_maps_routes_v1_compute_routes_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routes_v1_compute_routes_response_proto_goTypes,
		DependencyIndexes: file_google_maps_routes_v1_compute_routes_response_proto_depIdxs,
		MessageInfos:      file_google_maps_routes_v1_compute_routes_response_proto_msgTypes,
	}.Build()
	File_google_maps_routes_v1_compute_routes_response_proto = out.File
	file_google_maps_routes_v1_compute_routes_response_proto_rawDesc = nil
	file_google_maps_routes_v1_compute_routes_response_proto_goTypes = nil
	file_google_maps_routes_v1_compute_routes_response_proto_depIdxs = nil
}
