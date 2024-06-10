// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: protos/task.proto

package protos

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TasksPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Tasks   []*Task              `protobuf:"bytes,2,rep,name=Tasks,proto3" json:"Tasks,omitempty"`
	EndTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *TasksPlan) Reset() {
	*x = TasksPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TasksPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksPlan) ProtoMessage() {}

func (x *TasksPlan) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TasksPlan.ProtoReflect.Descriptor instead.
func (*TasksPlan) Descriptor() ([]byte, []int) {
	return file_protos_task_proto_rawDescGZIP(), []int{0}
}

func (x *TasksPlan) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *TasksPlan) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *TasksPlan) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Plugin       string               `protobuf:"bytes,2,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Instactions  []string             `protobuf:"bytes,3,rep,name=instactions,proto3" json:"instactions,omitempty"`
	CreationTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_protos_task_proto_rawDescGZIP(), []int{1}
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

func (x *Task) GetInstactions() []string {
	if x != nil {
		return x.Instactions
	}
	return nil
}

func (x *Task) GetCreationTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

type TaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Status  bool                 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Data    []*DataRow           `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	EndTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *TaskResult) Reset() {
	*x = TaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResult) ProtoMessage() {}

func (x *TaskResult) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TaskResult.ProtoReflect.Descriptor instead.
func (*TaskResult) Descriptor() ([]byte, []int) {
	return file_protos_task_proto_rawDescGZIP(), []int{2}
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

func (x *TaskResult) GetData() []*DataRow {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TaskResult) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type DataRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *DataRow) Reset() {
	*x = DataRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRow) ProtoMessage() {}

func (x *DataRow) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DataRow.ProtoReflect.Descriptor instead.
func (*DataRow) Descriptor() ([]byte, []int) {
	return file_protos_task_proto_rawDescGZIP(), []int{3}
}

func (x *DataRow) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_protos_task_proto protoreflect.FileDescriptor

var file_protos_task_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x09,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3e,
	0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x93,
	0x01, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x34,
	0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x52, 0x6f, 0x77, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_protos_task_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_task_proto_goTypes = []interface{}{
	(*TasksPlan)(nil),           // 0: protos.TasksPlan
	(*Task)(nil),                // 1: protos.Task
	(*TaskResult)(nil),          // 2: protos.TaskResult
	(*DataRow)(nil),             // 3: protos.DataRow
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_protos_task_proto_depIdxs = []int32{
	1, // 0: protos.TasksPlan.Tasks:type_name -> protos.Task
	4, // 1: protos.TasksPlan.endTime:type_name -> google.protobuf.Timestamp
	4, // 2: protos.Task.creationTime:type_name -> google.protobuf.Timestamp
	3, // 3: protos.TaskResult.data:type_name -> protos.DataRow
	4, // 4: protos.TaskResult.endTime:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_task_proto_init() }
func file_protos_task_proto_init() {
	if File_protos_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TasksPlan); i {
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
		file_protos_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRow); i {
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
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_task_proto_goTypes,
		DependencyIndexes: file_protos_task_proto_depIdxs,
		MessageInfos:      file_protos_task_proto_msgTypes,
	}.Build()
	File_protos_task_proto = out.File
	file_protos_task_proto_rawDesc = nil
	file_protos_task_proto_goTypes = nil
	file_protos_task_proto_depIdxs = nil
}
