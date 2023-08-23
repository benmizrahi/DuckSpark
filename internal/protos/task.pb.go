// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.12
// source: protos/task.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataType int32

const (
	DataType_string DataType = 0
	DataType_int    DataType = 1
	DataType_double DataType = 2
	DataType_long   DataType = 3
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0: "string",
		1: "int",
		2: "double",
		3: "long",
	}
	DataType_value = map[string]int32{
		"string": 0,
		"int":    1,
		"double": 2,
		"long":   3,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_task_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_protos_task_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_protos_task_proto_rawDescGZIP(), []int{0}
}

type IPartition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string  `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Data  []*Data `protobuf:"bytes,2,rep,name=Data,proto3" json:"Data,omitempty"`
	Tasks []*Task `protobuf:"bytes,3,rep,name=Tasks,proto3" json:"Tasks,omitempty"`
}

func (x *IPartition) Reset() {
	*x = IPartition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPartition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPartition) ProtoMessage() {}

func (x *IPartition) ProtoReflect() protoreflect.Message {
	mi := &file_protos_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPartition.ProtoReflect.Descriptor instead.
func (*IPartition) Descriptor() ([]byte, []int) {
	return file_protos_task_proto_rawDescGZIP(), []int{0}
}

func (x *IPartition) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *IPartition) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *IPartition) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataTypes    []DataType `protobuf:"varint,1,rep,packed,name=dataTypes,proto3,enum=protos.DataType" json:"dataTypes,omitempty"`
	CompressData []byte     `protobuf:"bytes,2,opt,name=compressData,proto3" json:"compressData,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_protos_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_protos_task_proto_rawDescGZIP(), []int{1}
}

func (x *Data) GetDataTypes() []DataType {
	if x != nil {
		return x.DataTypes
	}
	return nil
}

func (x *Data) GetCompressData() []byte {
	if x != nil {
		return x.CompressData
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Plugin       string                 `protobuf:"bytes,2,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Instruction  []string               `protobuf:"bytes,3,rep,name=instruction,proto3" json:"instruction,omitempty"`
	CreationTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_protos_task_proto_msgTypes[2]
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
	return file_protos_task_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Task) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *Task) GetInstruction() []string {
	if x != nil {
		return x.Instruction
	}
	return nil
}

func (x *Task) GetCreationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

type IPartitionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	TaskResults []*TaskResult          `protobuf:"bytes,2,rep,name=taskResults,proto3" json:"taskResults,omitempty"`
	EndTime     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *IPartitionResult) Reset() {
	*x = IPartitionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPartitionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPartitionResult) ProtoMessage() {}

func (x *IPartitionResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPartitionResult.ProtoReflect.Descriptor instead.
func (*IPartitionResult) Descriptor() ([]byte, []int) {
	return file_protos_task_proto_rawDescGZIP(), []int{3}
}

func (x *IPartitionResult) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *IPartitionResult) GetTaskResults() []*TaskResult {
	if x != nil {
		return x.TaskResults
	}
	return nil
}

func (x *IPartitionResult) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type TaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Status  bool                   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Data    []*Data                `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	EndTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *TaskResult) Reset() {
	*x = TaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResult) ProtoMessage() {}

func (x *TaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResult.ProtoReflect.Descriptor instead.
func (*TaskResult) Descriptor() ([]byte, []int) {
	return file_protos_task_proto_rawDescGZIP(), []int{4}
}

func (x *TaskResult) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *TaskResult) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *TaskResult) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TaskResult) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_protos_task_proto protoreflect.FileDescriptor

var file_protos_task_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x0a,
	0x49, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x20,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x22, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x22, 0x5a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x94, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x10, 0x49, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x34, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x90, 0x01, 0x0a,
	0x0a, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x2a,
	0x35, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x6c, 0x6f, 0x6e, 0x67, 0x10, 0x03, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_task_proto_rawDescOnce sync.Once
	file_protos_task_proto_rawDescData = file_protos_task_proto_rawDesc
)

func file_protos_task_proto_rawDescGZIP() []byte {
	file_protos_task_proto_rawDescOnce.Do(func() {
		file_protos_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_task_proto_rawDescData)
	})
	return file_protos_task_proto_rawDescData
}

var file_protos_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_task_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_task_proto_goTypes = []interface{}{
	(DataType)(0),                 // 0: protos.DataType
	(*IPartition)(nil),            // 1: protos.IPartition
	(*Data)(nil),                  // 2: protos.Data
	(*Task)(nil),                  // 3: protos.Task
	(*IPartitionResult)(nil),      // 4: protos.IPartitionResult
	(*TaskResult)(nil),            // 5: protos.TaskResult
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_protos_task_proto_depIdxs = []int32{
	2, // 0: protos.IPartition.Data:type_name -> protos.Data
	3, // 1: protos.IPartition.Tasks:type_name -> protos.Task
	0, // 2: protos.Data.dataTypes:type_name -> protos.DataType
	6, // 3: protos.Task.creationTime:type_name -> google.protobuf.Timestamp
	5, // 4: protos.IPartitionResult.taskResults:type_name -> protos.TaskResult
	6, // 5: protos.IPartitionResult.endTime:type_name -> google.protobuf.Timestamp
	2, // 6: protos.TaskResult.data:type_name -> protos.Data
	6, // 7: protos.TaskResult.endTime:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_protos_task_proto_init() }
func file_protos_task_proto_init() {
	if File_protos_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPartition); i {
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
		file_protos_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_protos_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPartitionResult); i {
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
		file_protos_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResult); i {
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
			RawDescriptor: file_protos_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_task_proto_goTypes,
		DependencyIndexes: file_protos_task_proto_depIdxs,
		EnumInfos:         file_protos_task_proto_enumTypes,
		MessageInfos:      file_protos_task_proto_msgTypes,
	}.Build()
	File_protos_task_proto = out.File
	file_protos_task_proto_rawDesc = nil
	file_protos_task_proto_goTypes = nil
	file_protos_task_proto_depIdxs = nil
}
