// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: tech/storezhang/dnsabre/message/rule.proto

package message

import (
	rule "github.com/dnsabre/core/rule"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_tech_storezhang_dnsabre_message_rule_proto protoreflect.FileDescriptor

var file_tech_storezhang_dnsabre_message_rule_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e,
	0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x74, 0x65,
	0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e,
	0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x74, 0x65, 0x63,
	0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x6e, 0x73,
	0x61, 0x62, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72,
	0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72,
	0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xea, 0x04, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x74,
	0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x29, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x29, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61,
	0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x72, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x73, 0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x62, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x06, 0x2f, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x76, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x74, 0x65,
	0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e,
	0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x73,
	0x70, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x62, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12,
	0x0b, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x74, 0x0a, 0x06,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x1a, 0x2c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52,
	0x73, 0x70, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e,
	0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e,
	0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x75, 0x6c,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2c, 0x2e, 0x74, 0x65,
	0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e,
	0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x3a, 0x01, 0x2a, 0x62, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x0b, 0x2f, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x79, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x2c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68,
	0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x2c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e,
	0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72,
	0x75, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x42, 0x47, 0x0a, 0x1f, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x62, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_tech_storezhang_dnsabre_message_rule_proto_goTypes = []interface{}{
	(*rule.AddReq)(nil),    // 0: tech.storezhang.dnsabre.core.rule.AddReq
	(*rule.GetReq)(nil),    // 1: tech.storezhang.dnsabre.core.rule.GetReq
	(*rule.PagingReq)(nil), // 2: tech.storezhang.dnsabre.core.rule.PagingReq
	(*rule.UpdateReq)(nil), // 3: tech.storezhang.dnsabre.core.rule.UpdateReq
	(*rule.DeleteReq)(nil), // 4: tech.storezhang.dnsabre.core.rule.DeleteReq
	(*rule.AddRsp)(nil),    // 5: tech.storezhang.dnsabre.core.rule.AddRsp
	(*rule.GetRsp)(nil),    // 6: tech.storezhang.dnsabre.core.rule.GetRsp
	(*rule.PagingRsp)(nil), // 7: tech.storezhang.dnsabre.core.rule.PagingRsp
	(*rule.UpdateRsp)(nil), // 8: tech.storezhang.dnsabre.core.rule.UpdateRsp
	(*rule.DeleteRsp)(nil), // 9: tech.storezhang.dnsabre.core.rule.DeleteRsp
}
var file_tech_storezhang_dnsabre_message_rule_proto_depIdxs = []int32{
	0, // 0: tech.storezhang.dnsabre.message.Rule.Add:input_type -> tech.storezhang.dnsabre.core.rule.AddReq
	1, // 1: tech.storezhang.dnsabre.message.Rule.Get:input_type -> tech.storezhang.dnsabre.core.rule.GetReq
	2, // 2: tech.storezhang.dnsabre.message.Rule.Paging:input_type -> tech.storezhang.dnsabre.core.rule.PagingReq
	3, // 3: tech.storezhang.dnsabre.message.Rule.Update:input_type -> tech.storezhang.dnsabre.core.rule.UpdateReq
	4, // 4: tech.storezhang.dnsabre.message.Rule.Delete:input_type -> tech.storezhang.dnsabre.core.rule.DeleteReq
	5, // 5: tech.storezhang.dnsabre.message.Rule.Add:output_type -> tech.storezhang.dnsabre.core.rule.AddRsp
	6, // 6: tech.storezhang.dnsabre.message.Rule.Get:output_type -> tech.storezhang.dnsabre.core.rule.GetRsp
	7, // 7: tech.storezhang.dnsabre.message.Rule.Paging:output_type -> tech.storezhang.dnsabre.core.rule.PagingRsp
	8, // 8: tech.storezhang.dnsabre.message.Rule.Update:output_type -> tech.storezhang.dnsabre.core.rule.UpdateRsp
	9, // 9: tech.storezhang.dnsabre.message.Rule.Delete:output_type -> tech.storezhang.dnsabre.core.rule.DeleteRsp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tech_storezhang_dnsabre_message_rule_proto_init() }
func file_tech_storezhang_dnsabre_message_rule_proto_init() {
	if File_tech_storezhang_dnsabre_message_rule_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tech_storezhang_dnsabre_message_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tech_storezhang_dnsabre_message_rule_proto_goTypes,
		DependencyIndexes: file_tech_storezhang_dnsabre_message_rule_proto_depIdxs,
	}.Build()
	File_tech_storezhang_dnsabre_message_rule_proto = out.File
	file_tech_storezhang_dnsabre_message_rule_proto_rawDesc = nil
	file_tech_storezhang_dnsabre_message_rule_proto_goTypes = nil
	file_tech_storezhang_dnsabre_message_rule_proto_depIdxs = nil
}
