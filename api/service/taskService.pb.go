// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: taskService.proto

package service

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

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//@inject_tag: json:"Id" form:"Id"
	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//@inject_tag: json:"Uid" form:"Uid"
	Uid uint64 `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	//@inject_tag: json:"Title" form:"Title"
	Title string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	//@inject_tag: json:"Content" form:"Content"
	Content string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	//@inject_tag: json:"StartTime" form:"StartTime"
	StartTime int64 `protobuf:"varint,5,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	//@inject_tag: json:"EndTime" form:"EndTime"
	EndTime int64 `protobuf:"varint,6,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	//@inject_tag: json:"Status" form:"Status"
	Status int64 `protobuf:"varint,7,opt,name=Status,proto3" json:"Status,omitempty"`
	// @inject_tag: json:"Start" form:"Start" uri:"Start"
	Start uint32 `protobuf:"varint,8,opt,name=Start,proto3" json:"Start,omitempty"`
	// @inject_tag: json:"Limit" form:"Limit" uri:"Limit"
	Limit uint32 `protobuf:"varint,9,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taskService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_taskService_proto_rawDescGZIP(), []int{0}
}

func (x *TaskRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskRequest) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *TaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TaskRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *TaskRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *TaskRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskRequest) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *TaskRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type TaskListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskList []*TaskModel `protobuf:"bytes,1,rep,name=TaskList,proto3" json:"TaskList,omitempty"`
	// @inject_tag: json:"Count"
	Count uint32 `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *TaskListResponse) Reset() {
	*x = TaskListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse) ProtoMessage() {}

func (x *TaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_taskService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse.ProtoReflect.Descriptor instead.
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return file_taskService_proto_rawDescGZIP(), []int{1}
}

func (x *TaskListResponse) GetTaskList() []*TaskModel {
	if x != nil {
		return x.TaskList
	}
	return nil
}

func (x *TaskListResponse) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type TaskDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskDetail *TaskModel `protobuf:"bytes,1,opt,name=TaskDetail,proto3" json:"TaskDetail,omitempty"`
}

func (x *TaskDetailResponse) Reset() {
	*x = TaskDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDetailResponse) ProtoMessage() {}

func (x *TaskDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_taskService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDetailResponse.ProtoReflect.Descriptor instead.
func (*TaskDetailResponse) Descriptor() ([]byte, []int) {
	return file_taskService_proto_rawDescGZIP(), []int{2}
}

func (x *TaskDetailResponse) GetTaskDetail() *TaskModel {
	if x != nil {
		return x.TaskDetail
	}
	return nil
}

var File_taskService_proto protoreflect.FileDescriptor

var file_taskService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0f, 0x74, 0x61,
	0x73, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01,
	0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x58, 0x0a, 0x10, 0x54,
	0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x48, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x54,
	0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x32,
	0xcf, 0x02, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_taskService_proto_rawDescOnce sync.Once
	file_taskService_proto_rawDescData = file_taskService_proto_rawDesc
)

func file_taskService_proto_rawDescGZIP() []byte {
	file_taskService_proto_rawDescOnce.Do(func() {
		file_taskService_proto_rawDescData = protoimpl.X.CompressGZIP(file_taskService_proto_rawDescData)
	})
	return file_taskService_proto_rawDescData
}

var file_taskService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_taskService_proto_goTypes = []interface{}{
	(*TaskRequest)(nil),        // 0: service.TaskRequest
	(*TaskListResponse)(nil),   // 1: service.TaskListResponse
	(*TaskDetailResponse)(nil), // 2: service.TaskDetailResponse
	(*TaskModel)(nil),          // 3: service.TaskModel
}
var file_taskService_proto_depIdxs = []int32{
	3, // 0: service.TaskListResponse.TaskList:type_name -> service.TaskModel
	3, // 1: service.TaskDetailResponse.TaskDetail:type_name -> service.TaskModel
	0, // 2: service.TaskService.CreateTask:input_type -> service.TaskRequest
	0, // 3: service.TaskService.GetTasksList:input_type -> service.TaskRequest
	0, // 4: service.TaskService.GetTask:input_type -> service.TaskRequest
	0, // 5: service.TaskService.UpdateTask:input_type -> service.TaskRequest
	0, // 6: service.TaskService.DeleteTask:input_type -> service.TaskRequest
	2, // 7: service.TaskService.CreateTask:output_type -> service.TaskDetailResponse
	1, // 8: service.TaskService.GetTasksList:output_type -> service.TaskListResponse
	2, // 9: service.TaskService.GetTask:output_type -> service.TaskDetailResponse
	2, // 10: service.TaskService.UpdateTask:output_type -> service.TaskDetailResponse
	2, // 11: service.TaskService.DeleteTask:output_type -> service.TaskDetailResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_taskService_proto_init() }
func file_taskService_proto_init() {
	if File_taskService_proto != nil {
		return
	}
	file_taskModel_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_taskService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
		file_taskService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListResponse); i {
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
		file_taskService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDetailResponse); i {
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
			RawDescriptor: file_taskService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_taskService_proto_goTypes,
		DependencyIndexes: file_taskService_proto_depIdxs,
		MessageInfos:      file_taskService_proto_msgTypes,
	}.Build()
	File_taskService_proto = out.File
	file_taskService_proto_rawDesc = nil
	file_taskService_proto_goTypes = nil
	file_taskService_proto_depIdxs = nil
}
