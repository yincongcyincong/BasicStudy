// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.5.1
// source: github.com/yincongcyincong/BasicStudy/proto/api/basic_study.proto

package api

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

type StudyQuestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`                                   // id
	CreateTime *uint64 `protobuf:"varint,2,opt,name=create_time,json=createTime" json:"create_time,omitempty"` // create time
	Question   *string `protobuf:"bytes,3,opt,name=question" json:"question,omitempty"`                        // question
	Answer     *string `protobuf:"bytes,4,opt,name=answer" json:"answer,omitempty"`                            // answer
	TestId     *uint64 `protobuf:"varint,5,opt,name=test_id,json=testId" json:"test_id,omitempty"`             // test
	Type       *uint32 `protobuf:"varint,6,opt,name=type" json:"type,omitempty"`                               // 1-math 2-english
}

func (x *StudyQuestion) Reset() {
	*x = StudyQuestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudyQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudyQuestion) ProtoMessage() {}

func (x *StudyQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudyQuestion.ProtoReflect.Descriptor instead.
func (*StudyQuestion) Descriptor() ([]byte, []int) {
	return file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDescGZIP(), []int{0}
}

func (x *StudyQuestion) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *StudyQuestion) GetCreateTime() uint64 {
	if x != nil && x.CreateTime != nil {
		return *x.CreateTime
	}
	return 0
}

func (x *StudyQuestion) GetQuestion() string {
	if x != nil && x.Question != nil {
		return *x.Question
	}
	return ""
}

func (x *StudyQuestion) GetAnswer() string {
	if x != nil && x.Answer != nil {
		return *x.Answer
	}
	return ""
}

func (x *StudyQuestion) GetTestId() uint64 {
	if x != nil && x.TestId != nil {
		return *x.TestId
	}
	return 0
}

func (x *StudyQuestion) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

type StudyUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`                                   // id
	CreateTime *uint64 `protobuf:"varint,2,opt,name=create_time,json=createTime" json:"create_time,omitempty"` // create time
	Username   *string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`                        // username
	Password   *string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`                        // password
	Status     *uint32 `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`                           // 1-normal 2-deleted
}

func (x *StudyUser) Reset() {
	*x = StudyUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudyUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudyUser) ProtoMessage() {}

func (x *StudyUser) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudyUser.ProtoReflect.Descriptor instead.
func (*StudyUser) Descriptor() ([]byte, []int) {
	return file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDescGZIP(), []int{1}
}

func (x *StudyUser) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *StudyUser) GetCreateTime() uint64 {
	if x != nil && x.CreateTime != nil {
		return *x.CreateTime
	}
	return 0
}

func (x *StudyUser) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *StudyUser) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *StudyUser) GetStatus() uint32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

type StudyTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`    // id
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"` // username
}

func (x *StudyTest) Reset() {
	*x = StudyTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudyTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudyTest) ProtoMessage() {}

func (x *StudyTest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudyTest.ProtoReflect.Descriptor instead.
func (*StudyTest) Descriptor() ([]byte, []int) {
	return file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDescGZIP(), []int{2}
}

func (x *StudyTest) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *StudyTest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var File_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto protoreflect.FileDescriptor

var file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x69, 0x6e,
	0x63, 0x6f, 0x6e, 0x67, 0x63, 0x79, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x67, 0x2f, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x53, 0x74, 0x75, 0x64, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x22, 0xa1, 0x01, 0x0a, 0x0d, 0x73, 0x74, 0x75, 0x64, 0x79, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x2f, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64, 0x79, 0x54, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x67, 0x63, 0x79, 0x69, 0x6e, 0x63, 0x6f,
	0x6e, 0x67, 0x2f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x53, 0x74, 0x75, 0x64, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69,
}

var (
	file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDescOnce sync.Once
	file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDescData = file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDesc
)

func file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDescGZIP() []byte {
	file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDescOnce.Do(func() {
		file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDescData)
	})
	return file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDescData
}

var file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_goTypes = []interface{}{
	(*StudyQuestion)(nil), // 0: basic_study.studyQuestion
	(*StudyUser)(nil),     // 1: basic_study.studyUser
	(*StudyTest)(nil),     // 2: basic_study.studyTest
}
var file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_init() }
func file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_init() {
	if File_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudyQuestion); i {
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
		file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudyUser); i {
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
		file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudyTest); i {
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
			RawDescriptor: file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_goTypes,
		DependencyIndexes: file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_depIdxs,
		MessageInfos:      file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_msgTypes,
	}.Build()
	File_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto = out.File
	file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_rawDesc = nil
	file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_goTypes = nil
	file_github_com_yincongcyincong_BasicStudy_proto_api_basic_study_proto_depIdxs = nil
}
