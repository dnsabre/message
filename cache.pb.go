// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: tech/storezhang/dnsabre/message/cache.proto

package message

import (
	cache "github.com/dnsabre/core/cache"
	vo "github.com/dnsabre/core/vo"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_tech_storezhang_dnsabre_message_cache_proto protoreflect.FileDescriptor

var file_tech_storezhang_dnsabre_message_cache_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e,
	0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x74,
	0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64,
	0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x74, 0x65, 0x63, 0x68, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x74, 0x65, 0x63, 0x68,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61,
	0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x6f, 0x2f, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xae, 0x02, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x12, 0x56, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x26, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61,
	0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x6f, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09,
	0x12, 0x07, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x12, 0x81, 0x01, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x2d, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x73, 0x70, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x62, 0x05, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x1a, 0x07, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x12, 0x49, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x2a,
	0x07, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x73, 0x42, 0x47, 0x0a, 0x1f, 0x74, 0x65, 0x63, 0x68,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61,
	0x62, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x01, 0x5a, 0x22, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72,
	0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tech_storezhang_dnsabre_message_cache_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),   // 0: google.protobuf.Empty
	(*cache.UpdateReq)(nil), // 1: tech.storezhang.dnsabre.core.cache.UpdateReq
	(*vo.Cache)(nil),        // 2: tech.storezhang.dnsabre.core.vo.Cache
	(*cache.UpdateRsp)(nil), // 3: tech.storezhang.dnsabre.core.cache.UpdateRsp
}
var file_tech_storezhang_dnsabre_message_cache_proto_depIdxs = []int32{
	0, // 0: tech.storezhang.dnsabre.message.Cache.Get:input_type -> google.protobuf.Empty
	1, // 1: tech.storezhang.dnsabre.message.Cache.Update:input_type -> tech.storezhang.dnsabre.core.cache.UpdateReq
	0, // 2: tech.storezhang.dnsabre.message.Cache.Delete:input_type -> google.protobuf.Empty
	2, // 3: tech.storezhang.dnsabre.message.Cache.Get:output_type -> tech.storezhang.dnsabre.core.vo.Cache
	3, // 4: tech.storezhang.dnsabre.message.Cache.Update:output_type -> tech.storezhang.dnsabre.core.cache.UpdateRsp
	0, // 5: tech.storezhang.dnsabre.message.Cache.Delete:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tech_storezhang_dnsabre_message_cache_proto_init() }
func file_tech_storezhang_dnsabre_message_cache_proto_init() {
	if File_tech_storezhang_dnsabre_message_cache_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tech_storezhang_dnsabre_message_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tech_storezhang_dnsabre_message_cache_proto_goTypes,
		DependencyIndexes: file_tech_storezhang_dnsabre_message_cache_proto_depIdxs,
	}.Build()
	File_tech_storezhang_dnsabre_message_cache_proto = out.File
	file_tech_storezhang_dnsabre_message_cache_proto_rawDesc = nil
	file_tech_storezhang_dnsabre_message_cache_proto_goTypes = nil
	file_tech_storezhang_dnsabre_message_cache_proto_depIdxs = nil
}
