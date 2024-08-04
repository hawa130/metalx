// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: controller.proto

package proto

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

type TaskType int32

const (
	TaskType_FileDownload TaskType = 0 // Download a file from a URL
	TaskType_FileUpload   TaskType = 1 // Upload a file to a URL
	TaskType_Command      TaskType = 2 // Execute a command
)

// Enum value maps for TaskType.
var (
	TaskType_name = map[int32]string{
		0: "FileDownload",
		1: "FileUpload",
		2: "Command",
	}
	TaskType_value = map[string]int32{
		"FileDownload": 0,
		"FileUpload":   1,
		"Command":      2,
	}
)

func (x TaskType) Enum() *TaskType {
	p := new(TaskType)
	*p = x
	return p
}

func (x TaskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_proto_enumTypes[0].Descriptor()
}

func (TaskType) Type() protoreflect.EnumType {
	return &file_controller_proto_enumTypes[0]
}

func (x TaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskType.Descriptor instead.
func (TaskType) EnumDescriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0}
}

type RegisterAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`      // Agent identifier, generated by edge agent from its own information (e.g. CPU ID) or user-defined, maybe duplicated with others nodes
	Ip   string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`      // Agent IP address, can be accessed by controller node
	Port int32  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"` // Agent port, can be accessed by controller node
}

func (x *RegisterAgentRequest) Reset() {
	*x = RegisterAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAgentRequest) ProtoMessage() {}

func (x *RegisterAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAgentRequest.ProtoReflect.Descriptor instead.
func (*RegisterAgentRequest) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterAgentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterAgentRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *RegisterAgentRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type RegisterAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success int32  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 0 for success, others for failure
	Token   string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`      // Token for agent to communicate with controller, generated by controller, should be unique
}

func (x *RegisterAgentResponse) Reset() {
	*x = RegisterAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAgentResponse) ProtoMessage() {}

func (x *RegisterAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAgentResponse.ProtoReflect.Descriptor instead.
func (*RegisterAgentResponse) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterAgentResponse) GetSuccess() int32 {
	if x != nil {
		return x.Success
	}
	return 0
}

func (x *RegisterAgentResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                                   // Token obtained from controller, to identify the agent
	CurrentState string `protobuf:"bytes,2,opt,name=current_state,json=currentState,proto3" json:"current_state,omitempty"` // Current state of the agent, used to determine what task to assign
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_controller_proto_rawDescGZIP(), []int{2}
}

func (x *GetTaskRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetTaskRequest) GetCurrentState() string {
	if x != nil {
		return x.CurrentState
	}
	return ""
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type TaskType `protobuf:"varint,1,opt,name=type,proto3,enum=TaskType" json:"type,omitempty"` // Type of the task
	Url  *string  `protobuf:"bytes,2,opt,name=url,proto3,oneof" json:"url,omitempty"`            // URL of the file to download or upload
	Path string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`                // Path of the file to download or upload, or executable file to execute
	Args []string `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`                // Arguments for the command
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_controller_proto_rawDescGZIP(), []int{3}
}

func (x *Task) GetType() TaskType {
	if x != nil {
		return x.Type
	}
	return TaskType_FileDownload
}

func (x *Task) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *Task) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Task) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type GetTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId           int64  `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`                                // Task ID, used to identify the task
	NextState        string `protobuf:"bytes,2,opt,name=next_state,json=nextState,proto3" json:"next_state,omitempty"`                        // next state of the agent if task succeed. controller DO NOT maintain the state of agent
	NextStateFailure string `protobuf:"bytes,3,opt,name=next_state_failure,json=nextStateFailure,proto3" json:"next_state_failure,omitempty"` // next state of the agent if task failed. controller DO NOT maintain the state of agent
	Task             *Task  `protobuf:"bytes,4,opt,name=task,proto3" json:"task,omitempty"`                                                   // Task to be executed by agent
	StreamingLog     bool   `protobuf:"varint,5,opt,name=streamingLog,proto3" json:"streamingLog,omitempty"`                                  // Whether to stream log to controller
}

func (x *GetTaskResponse) Reset() {
	*x = GetTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskResponse) ProtoMessage() {}

func (x *GetTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskResponse.ProtoReflect.Descriptor instead.
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{4}
}

func (x *GetTaskResponse) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *GetTaskResponse) GetNextState() string {
	if x != nil {
		return x.NextState
	}
	return ""
}

func (x *GetTaskResponse) GetNextStateFailure() string {
	if x != nil {
		return x.NextStateFailure
	}
	return ""
}

func (x *GetTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *GetTaskResponse) GetStreamingLog() bool {
	if x != nil {
		return x.StreamingLog
	}
	return false
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                  // Token obtained from controller, to identify the agent
	TaskId  int64  `protobuf:"varint,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"` // Task ID
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`              // Log message
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{5}
}

func (x *Log) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Log) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *Log) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // Token obtained from controller, to identify the agent
	Event string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"` // Event type
	Args  [][]byte `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`   // Event arguments
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{6}
}

func (x *Event) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Event) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *Event) GetArgs() [][]byte {
	if x != nil {
		return x.Args
	}
	return nil
}

var File_controller_proto protoreflect.FileDescriptor

var file_controller_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4a, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x47, 0x0a, 0x15,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x6c, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c,
	0x22, 0xb6, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x61,
	0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x22, 0x4e, 0x0a, 0x03, 0x4c, 0x6f, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x72,
	0x67, 0x73, 0x2a, 0x39, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x10, 0x02, 0x32, 0xfa, 0x01,
	0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x04,
	0x70, 0x69, 0x6e, 0x67, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x06, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x0f, 0x75, 0x6e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2e, 0x0a,
	0x07, 0x67, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x20, 0x0a,
	0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x12, 0x04, 0x2e,
	0x4c, 0x6f, 0x67, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x77, 0x61, 0x31, 0x33, 0x30,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_proto_rawDescOnce sync.Once
	file_controller_proto_rawDescData = file_controller_proto_rawDesc
)

func file_controller_proto_rawDescGZIP() []byte {
	file_controller_proto_rawDescOnce.Do(func() {
		file_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_proto_rawDescData)
	})
	return file_controller_proto_rawDescData
}

var file_controller_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_controller_proto_goTypes = []any{
	(TaskType)(0),                 // 0: TaskType
	(*RegisterAgentRequest)(nil),  // 1: RegisterAgentRequest
	(*RegisterAgentResponse)(nil), // 2: RegisterAgentResponse
	(*GetTaskRequest)(nil),        // 3: GetTaskRequest
	(*Task)(nil),                  // 4: Task
	(*GetTaskResponse)(nil),       // 5: GetTaskResponse
	(*Log)(nil),                   // 6: Log
	(*Event)(nil),                 // 7: Event
	(*Empty)(nil),                 // 8: Empty
}
var file_controller_proto_depIdxs = []int32{
	0, // 0: Task.type:type_name -> TaskType
	4, // 1: GetTaskResponse.task:type_name -> Task
	8, // 2: Controller.ping:input_type -> Empty
	1, // 3: Controller.registerAgent:input_type -> RegisterAgentRequest
	8, // 4: Controller.unregisterAgent:input_type -> Empty
	3, // 5: Controller.getTask:input_type -> GetTaskRequest
	6, // 6: Controller.streamingLog:input_type -> Log
	7, // 7: Controller.event:input_type -> Event
	8, // 8: Controller.ping:output_type -> Empty
	2, // 9: Controller.registerAgent:output_type -> RegisterAgentResponse
	8, // 10: Controller.unregisterAgent:output_type -> Empty
	5, // 11: Controller.getTask:output_type -> GetTaskResponse
	8, // 12: Controller.streamingLog:output_type -> Empty
	8, // 13: Controller.event:output_type -> Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_controller_proto_init() }
func file_controller_proto_init() {
	if File_controller_proto != nil {
		return
	}
	file_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controller_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterAgentRequest); i {
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
		file_controller_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterAgentResponse); i {
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
		file_controller_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetTaskRequest); i {
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
		file_controller_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Task); i {
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
		file_controller_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetTaskResponse); i {
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
		file_controller_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Log); i {
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
		file_controller_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Event); i {
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
	file_controller_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controller_proto_goTypes,
		DependencyIndexes: file_controller_proto_depIdxs,
		EnumInfos:         file_controller_proto_enumTypes,
		MessageInfos:      file_controller_proto_msgTypes,
	}.Build()
	File_controller_proto = out.File
	file_controller_proto_rawDesc = nil
	file_controller_proto_goTypes = nil
	file_controller_proto_depIdxs = nil
}
