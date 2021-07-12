// Copyright 2020 Google LLC
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
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: google/ads/googleads/v2/enums/page_one_promoted_strategy_goal.proto

package enums

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Enum describing possible strategy goals.
type PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal int32

const (
	// Not specified.
	PageOnePromotedStrategyGoalEnum_UNSPECIFIED PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal = 0
	// Used for return value only. Represents value unknown in this version.
	PageOnePromotedStrategyGoalEnum_UNKNOWN PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal = 1
	// First page on google.com.
	PageOnePromotedStrategyGoalEnum_FIRST_PAGE PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal = 2
	// Top slots of the first page on google.com.
	PageOnePromotedStrategyGoalEnum_FIRST_PAGE_PROMOTED PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal = 3
)

// Enum value maps for PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal.
var (
	PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "FIRST_PAGE",
		3: "FIRST_PAGE_PROMOTED",
	}
	PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal_value = map[string]int32{
		"UNSPECIFIED":         0,
		"UNKNOWN":             1,
		"FIRST_PAGE":          2,
		"FIRST_PAGE_PROMOTED": 3,
	}
)

func (x PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal) Enum() *PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal {
	p := new(PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal)
	*p = x
	return p
}

func (x PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_enumTypes[0].Descriptor()
}

func (PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_enumTypes[0]
}

func (x PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal.Descriptor instead.
func (PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible strategy goals: where impressions are
// desired to be shown on search result pages.
type PageOnePromotedStrategyGoalEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PageOnePromotedStrategyGoalEnum) Reset() {
	*x = PageOnePromotedStrategyGoalEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageOnePromotedStrategyGoalEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageOnePromotedStrategyGoalEnum) ProtoMessage() {}

func (x *PageOnePromotedStrategyGoalEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageOnePromotedStrategyGoalEnum.ProtoReflect.Descriptor instead.
func (*PageOnePromotedStrategyGoalEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x1f, 0x50, 0x61, 0x67, 0x65, 0x4f, 0x6e, 0x65, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x47, 0x6f,
	0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x64, 0x0a, 0x1b, 0x50, 0x61, 0x67, 0x65, 0x4f, 0x6e,
	0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x47, 0x6f, 0x61, 0x6c, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x50, 0x41, 0x47,
	0x45, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x50, 0x41, 0x47,
	0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x4f, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0xf5, 0x01, 0x0a,
	0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x42, 0x20, 0x50, 0x61, 0x67, 0x65, 0x4f, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x47, 0x6f, 0x61, 0x6c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x32, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a, 0x3a, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDescData = file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDesc
)

func file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDescData
}

var file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_goTypes = []interface{}{
	(PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal)(0), // 0: google.ads.googleads.v2.enums.PageOnePromotedStrategyGoalEnum.PageOnePromotedStrategyGoal
	(*PageOnePromotedStrategyGoalEnum)(nil),                          // 1: google.ads.googleads.v2.enums.PageOnePromotedStrategyGoalEnum
}
var file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_init() }
func file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_init() {
	if File_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageOnePromotedStrategyGoalEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto = out.File
	file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_rawDesc = nil
	file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_goTypes = nil
	file_google_ads_googleads_v2_enums_page_one_promoted_strategy_goal_proto_depIdxs = nil
}