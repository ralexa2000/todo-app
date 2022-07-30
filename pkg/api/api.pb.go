// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TaskCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Task    string `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	DueDate string `protobuf:"bytes,3,opt,name=dueDate,proto3" json:"dueDate,omitempty"`
}

func (x *TaskCreateRequest) Reset() {
	*x = TaskCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreateRequest) ProtoMessage() {}

func (x *TaskCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreateRequest.ProtoReflect.Descriptor instead.
func (*TaskCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *TaskCreateRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *TaskCreateRequest) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *TaskCreateRequest) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

type TaskCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskCreateResponse) Reset() {
	*x = TaskCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreateResponse) ProtoMessage() {}

func (x *TaskCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreateResponse.ProtoReflect.Descriptor instead.
func (*TaskCreateResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

type TaskListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *TaskListRequest) Reset() {
	*x = TaskListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListRequest) ProtoMessage() {}

func (x *TaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListRequest.ProtoReflect.Descriptor instead.
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *TaskListRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type TaskListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*TaskListResponse_Task `protobuf:"bytes,1,rep,name=Tasks,proto3" json:"Tasks,omitempty"`
}

func (x *TaskListResponse) Reset() {
	*x = TaskListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse) ProtoMessage() {}

func (x *TaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
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
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *TaskListResponse) GetTasks() []*TaskListResponse_Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type TaskUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId  uint64 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	User    string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Task    string `protobuf:"bytes,3,opt,name=task,proto3" json:"task,omitempty"`
	DueDate string `protobuf:"bytes,4,opt,name=dueDate,proto3" json:"dueDate,omitempty"`
}

func (x *TaskUpdateRequest) Reset() {
	*x = TaskUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskUpdateRequest) ProtoMessage() {}

func (x *TaskUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskUpdateRequest.ProtoReflect.Descriptor instead.
func (*TaskUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *TaskUpdateRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskUpdateRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *TaskUpdateRequest) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *TaskUpdateRequest) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

type TaskUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskUpdateResponse) Reset() {
	*x = TaskUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskUpdateResponse) ProtoMessage() {}

func (x *TaskUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskUpdateResponse.ProtoReflect.Descriptor instead.
func (*TaskUpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

type TaskDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	TaskId uint64 `protobuf:"varint,2,opt,name=taskId,proto3" json:"taskId,omitempty"`
}

func (x *TaskDeleteRequest) Reset() {
	*x = TaskDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDeleteRequest) ProtoMessage() {}

func (x *TaskDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDeleteRequest.ProtoReflect.Descriptor instead.
func (*TaskDeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *TaskDeleteRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *TaskDeleteRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type TaskDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskDeleteResponse) Reset() {
	*x = TaskDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDeleteResponse) ProtoMessage() {}

func (x *TaskDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDeleteResponse.ProtoReflect.Descriptor instead.
func (*TaskDeleteResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

type TaskListResponse_Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId  uint64 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Task    string `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	DueDate string `protobuf:"bytes,3,opt,name=dueDate,proto3" json:"dueDate,omitempty"`
}

func (x *TaskListResponse_Task) Reset() {
	*x = TaskListResponse_Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListResponse_Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResponse_Task) ProtoMessage() {}

func (x *TaskListResponse_Task) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResponse_Task.ProtoReflect.Descriptor instead.
func (*TaskListResponse_Task) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3, 0}
}

func (x *TaskListResponse_Task) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskListResponse_Task) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

func (x *TaskListResponse_Task) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x11, 0x54,
	0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x75, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x9f, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d,
	0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x1a, 0x4c, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x22, 0x6d, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x22, 0x14, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xed, 0x02,
	0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x59, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x53, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x59, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x23, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76,
	0x2f, 0x72, 0x61, 0x6c, 0x65, 0x78, 0x61, 0x32, 0x30, 0x30, 0x30, 0x2f, 0x74, 0x6f, 0x64, 0x6f,
	0x2d, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_proto_goTypes = []interface{}{
	(*TaskCreateRequest)(nil),     // 0: ozon.dev.mc2.api.TaskCreateRequest
	(*TaskCreateResponse)(nil),    // 1: ozon.dev.mc2.api.TaskCreateResponse
	(*TaskListRequest)(nil),       // 2: ozon.dev.mc2.api.TaskListRequest
	(*TaskListResponse)(nil),      // 3: ozon.dev.mc2.api.TaskListResponse
	(*TaskUpdateRequest)(nil),     // 4: ozon.dev.mc2.api.TaskUpdateRequest
	(*TaskUpdateResponse)(nil),    // 5: ozon.dev.mc2.api.TaskUpdateResponse
	(*TaskDeleteRequest)(nil),     // 6: ozon.dev.mc2.api.TaskDeleteRequest
	(*TaskDeleteResponse)(nil),    // 7: ozon.dev.mc2.api.TaskDeleteResponse
	(*TaskListResponse_Task)(nil), // 8: ozon.dev.mc2.api.TaskListResponse.Task
}
var file_api_proto_depIdxs = []int32{
	8, // 0: ozon.dev.mc2.api.TaskListResponse.Tasks:type_name -> ozon.dev.mc2.api.TaskListResponse.Task
	0, // 1: ozon.dev.mc2.api.Admin.TaskCreate:input_type -> ozon.dev.mc2.api.TaskCreateRequest
	2, // 2: ozon.dev.mc2.api.Admin.TaskList:input_type -> ozon.dev.mc2.api.TaskListRequest
	4, // 3: ozon.dev.mc2.api.Admin.TaskUpdate:input_type -> ozon.dev.mc2.api.TaskUpdateRequest
	6, // 4: ozon.dev.mc2.api.Admin.TaskDelete:input_type -> ozon.dev.mc2.api.TaskDeleteRequest
	1, // 5: ozon.dev.mc2.api.Admin.TaskCreate:output_type -> ozon.dev.mc2.api.TaskCreateResponse
	3, // 6: ozon.dev.mc2.api.Admin.TaskList:output_type -> ozon.dev.mc2.api.TaskListResponse
	5, // 7: ozon.dev.mc2.api.Admin.TaskUpdate:output_type -> ozon.dev.mc2.api.TaskUpdateResponse
	7, // 8: ozon.dev.mc2.api.Admin.TaskDelete:output_type -> ozon.dev.mc2.api.TaskDeleteResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskCreateRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskCreateResponse); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskUpdateRequest); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskUpdateResponse); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDeleteRequest); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDeleteResponse); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListResponse_Task); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
