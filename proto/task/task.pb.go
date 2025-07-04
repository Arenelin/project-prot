// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: proto/task/task.proto

package task

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	UserId        uint32                 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	IsDone        *bool                  `protobuf:"varint,4,opt,name=isDone,proto3,oneof" json:"isDone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_proto_task_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_task_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Task) GetIsDone() bool {
	if x != nil && x.IsDone != nil {
		return *x.IsDone
	}
	return false
}

type GetTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	mi := &file_proto_task_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_task_proto_rawDescGZIP(), []int{1}
}

func (x *GetTaskRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListTasksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTasksRequest) Reset() {
	*x = ListTasksRequest{}
	mi := &file_proto_task_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksRequest) ProtoMessage() {}

func (x *ListTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksRequest.ProtoReflect.Descriptor instead.
func (*ListTasksRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_task_proto_rawDescGZIP(), []int{2}
}

type ListTasksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTasksResponse) Reset() {
	*x = ListTasksResponse{}
	mi := &file_proto_task_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksResponse) ProtoMessage() {}

func (x *ListTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksResponse.ProtoReflect.Descriptor instead.
func (*ListTasksResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_task_proto_rawDescGZIP(), []int{3}
}

func (x *ListTasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type ListTasksByUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTasksByUserRequest) Reset() {
	*x = ListTasksByUserRequest{}
	mi := &file_proto_task_task_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTasksByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksByUserRequest) ProtoMessage() {}

func (x *ListTasksByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_task_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksByUserRequest.ProtoReflect.Descriptor instead.
func (*ListTasksByUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_task_proto_rawDescGZIP(), []int{4}
}

func (x *ListTasksByUserRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListTasksByUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTasksByUserResponse) Reset() {
	*x = ListTasksByUserResponse{}
	mi := &file_proto_task_task_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTasksByUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksByUserResponse) ProtoMessage() {}

func (x *ListTasksByUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_task_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksByUserResponse.ProtoReflect.Descriptor instead.
func (*ListTasksByUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_task_proto_rawDescGZIP(), []int{5}
}

func (x *ListTasksByUserResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	UserId        uint32                 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	mi := &file_proto_task_task_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_task_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_task_proto_rawDescGZIP(), []int{6}
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	mi := &file_proto_task_task_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_task_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_task_proto_rawDescGZIP(), []int{7}
}

func (x *CreateTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	mi := &file_proto_task_task_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_task_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_task_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type UpdateTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTaskResponse) Reset() {
	*x = UpdateTaskResponse{}
	mi := &file_proto_task_task_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskResponse) ProtoMessage() {}

func (x *UpdateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_task_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskResponse.ProtoReflect.Descriptor instead.
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_task_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_proto_task_task_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_task_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_task_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteTaskRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	mi := &file_proto_task_task_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_task_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_task_proto_rawDescGZIP(), []int{11}
}

var File_proto_task_task_proto protoreflect.FileDescriptor

const file_proto_task_task_proto_rawDesc = "" +
	"\n" +
	"\x15proto/task/task.proto\x12\x04task\"l\n" +
	"\x04Task\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x16\n" +
	"\x06userId\x18\x03 \x01(\rR\x06userId\x12\x1b\n" +
	"\x06isDone\x18\x04 \x01(\bH\x00R\x06isDone\x88\x01\x01B\t\n" +
	"\a_isDone\" \n" +
	"\x0eGetTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\"\x12\n" +
	"\x10ListTasksRequest\"5\n" +
	"\x11ListTasksResponse\x12 \n" +
	"\x05tasks\x18\x01 \x03(\v2\n" +
	".task.TaskR\x05tasks\"(\n" +
	"\x16ListTasksByUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\";\n" +
	"\x17ListTasksByUserResponse\x12 \n" +
	"\x05tasks\x18\x01 \x03(\v2\n" +
	".task.TaskR\x05tasks\"A\n" +
	"\x11CreateTaskRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x16\n" +
	"\x06userId\x18\x02 \x01(\rR\x06userId\"4\n" +
	"\x12CreateTaskResponse\x12\x1e\n" +
	"\x04task\x18\x01 \x01(\v2\n" +
	".task.TaskR\x04task\"3\n" +
	"\x11UpdateTaskRequest\x12\x1e\n" +
	"\x04task\x18\x01 \x01(\v2\n" +
	".task.TaskR\x04task\"4\n" +
	"\x12UpdateTaskResponse\x12\x1e\n" +
	"\x04task\x18\x01 \x01(\v2\n" +
	".task.TaskR\x04task\"#\n" +
	"\x11DeleteTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\"\x14\n" +
	"\x12DeleteTaskResponse2\x8b\x03\n" +
	"\vTaskService\x12+\n" +
	"\aGetTask\x12\x14.task.GetTaskRequest\x1a\n" +
	".task.Task\x12<\n" +
	"\tListTasks\x12\x16.task.ListTasksRequest\x1a\x17.task.ListTasksResponse\x12N\n" +
	"\x0fListTasksByUser\x12\x1c.task.ListTasksByUserRequest\x1a\x1d.task.ListTasksByUserResponse\x12?\n" +
	"\n" +
	"CreateTask\x12\x17.task.CreateTaskRequest\x1a\x18.task.CreateTaskResponse\x12?\n" +
	"\n" +
	"UpdateTask\x12\x17.task.UpdateTaskRequest\x1a\x18.task.UpdateTaskResponse\x12?\n" +
	"\n" +
	"DeleteTask\x12\x17.task.DeleteTaskRequest\x1a\x18.task.DeleteTaskResponseB-Z+github.com/Arenelin/project-prot/proto/taskb\x06proto3"

var (
	file_proto_task_task_proto_rawDescOnce sync.Once
	file_proto_task_task_proto_rawDescData []byte
)

func file_proto_task_task_proto_rawDescGZIP() []byte {
	file_proto_task_task_proto_rawDescOnce.Do(func() {
		file_proto_task_task_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_task_task_proto_rawDesc), len(file_proto_task_task_proto_rawDesc)))
	})
	return file_proto_task_task_proto_rawDescData
}

var file_proto_task_task_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_task_task_proto_goTypes = []any{
	(*Task)(nil),                    // 0: task.Task
	(*GetTaskRequest)(nil),          // 1: task.GetTaskRequest
	(*ListTasksRequest)(nil),        // 2: task.ListTasksRequest
	(*ListTasksResponse)(nil),       // 3: task.ListTasksResponse
	(*ListTasksByUserRequest)(nil),  // 4: task.ListTasksByUserRequest
	(*ListTasksByUserResponse)(nil), // 5: task.ListTasksByUserResponse
	(*CreateTaskRequest)(nil),       // 6: task.CreateTaskRequest
	(*CreateTaskResponse)(nil),      // 7: task.CreateTaskResponse
	(*UpdateTaskRequest)(nil),       // 8: task.UpdateTaskRequest
	(*UpdateTaskResponse)(nil),      // 9: task.UpdateTaskResponse
	(*DeleteTaskRequest)(nil),       // 10: task.DeleteTaskRequest
	(*DeleteTaskResponse)(nil),      // 11: task.DeleteTaskResponse
}
var file_proto_task_task_proto_depIdxs = []int32{
	0,  // 0: task.ListTasksResponse.tasks:type_name -> task.Task
	0,  // 1: task.ListTasksByUserResponse.tasks:type_name -> task.Task
	0,  // 2: task.CreateTaskResponse.task:type_name -> task.Task
	0,  // 3: task.UpdateTaskRequest.task:type_name -> task.Task
	0,  // 4: task.UpdateTaskResponse.task:type_name -> task.Task
	1,  // 5: task.TaskService.GetTask:input_type -> task.GetTaskRequest
	2,  // 6: task.TaskService.ListTasks:input_type -> task.ListTasksRequest
	4,  // 7: task.TaskService.ListTasksByUser:input_type -> task.ListTasksByUserRequest
	6,  // 8: task.TaskService.CreateTask:input_type -> task.CreateTaskRequest
	8,  // 9: task.TaskService.UpdateTask:input_type -> task.UpdateTaskRequest
	10, // 10: task.TaskService.DeleteTask:input_type -> task.DeleteTaskRequest
	0,  // 11: task.TaskService.GetTask:output_type -> task.Task
	3,  // 12: task.TaskService.ListTasks:output_type -> task.ListTasksResponse
	5,  // 13: task.TaskService.ListTasksByUser:output_type -> task.ListTasksByUserResponse
	7,  // 14: task.TaskService.CreateTask:output_type -> task.CreateTaskResponse
	9,  // 15: task.TaskService.UpdateTask:output_type -> task.UpdateTaskResponse
	11, // 16: task.TaskService.DeleteTask:output_type -> task.DeleteTaskResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_task_task_proto_init() }
func file_proto_task_task_proto_init() {
	if File_proto_task_task_proto != nil {
		return
	}
	file_proto_task_task_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_task_task_proto_rawDesc), len(file_proto_task_task_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_task_task_proto_goTypes,
		DependencyIndexes: file_proto_task_task_proto_depIdxs,
		MessageInfos:      file_proto_task_task_proto_msgTypes,
	}.Build()
	File_proto_task_task_proto = out.File
	file_proto_task_task_proto_goTypes = nil
	file_proto_task_task_proto_depIdxs = nil
}
