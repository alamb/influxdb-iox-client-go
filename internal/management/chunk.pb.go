// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: influxdata/iox/management/v1/chunk.proto

package v1

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

// Which storage system is a chunk located in?
type ChunkStorage int32

const (
	// Not currently returned
	ChunkStorage_CHUNK_STORAGE_UNSPECIFIED ChunkStorage = 0
	// The chunk is still open for new writes, in the Mutable Buffer
	ChunkStorage_CHUNK_STORAGE_OPEN_MUTABLE_BUFFER ChunkStorage = 1
	// The chunk is no longer open for writes, in the Mutable Buffer
	ChunkStorage_CHUNK_STORAGE_CLOSED_MUTABLE_BUFFER ChunkStorage = 2
	// The chunk is in the Read Buffer (where it can not be mutated)
	ChunkStorage_CHUNK_STORAGE_READ_BUFFER ChunkStorage = 3
	// The chunk is in the Read Buffer and Object Store
	ChunkStorage_CHUNK_STORAGE_READ_BUFFER_AND_OBJECT_STORE ChunkStorage = 4
	// The chunk is stored in Object Storage (where it can not be mutated)
	ChunkStorage_CHUNK_STORAGE_OBJECT_STORE_ONLY ChunkStorage = 5
)

// Enum value maps for ChunkStorage.
var (
	ChunkStorage_name = map[int32]string{
		0: "CHUNK_STORAGE_UNSPECIFIED",
		1: "CHUNK_STORAGE_OPEN_MUTABLE_BUFFER",
		2: "CHUNK_STORAGE_CLOSED_MUTABLE_BUFFER",
		3: "CHUNK_STORAGE_READ_BUFFER",
		4: "CHUNK_STORAGE_READ_BUFFER_AND_OBJECT_STORE",
		5: "CHUNK_STORAGE_OBJECT_STORE_ONLY",
	}
	ChunkStorage_value = map[string]int32{
		"CHUNK_STORAGE_UNSPECIFIED":                  0,
		"CHUNK_STORAGE_OPEN_MUTABLE_BUFFER":          1,
		"CHUNK_STORAGE_CLOSED_MUTABLE_BUFFER":        2,
		"CHUNK_STORAGE_READ_BUFFER":                  3,
		"CHUNK_STORAGE_READ_BUFFER_AND_OBJECT_STORE": 4,
		"CHUNK_STORAGE_OBJECT_STORE_ONLY":            5,
	}
)

func (x ChunkStorage) Enum() *ChunkStorage {
	p := new(ChunkStorage)
	*p = x
	return p
}

func (x ChunkStorage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChunkStorage) Descriptor() protoreflect.EnumDescriptor {
	return file_influxdata_iox_management_v1_chunk_proto_enumTypes[0].Descriptor()
}

func (ChunkStorage) Type() protoreflect.EnumType {
	return &file_influxdata_iox_management_v1_chunk_proto_enumTypes[0]
}

func (x ChunkStorage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChunkStorage.Descriptor instead.
func (ChunkStorage) EnumDescriptor() ([]byte, []int) {
	return file_influxdata_iox_management_v1_chunk_proto_rawDescGZIP(), []int{0}
}

// Is there any lifecycle action currently outstanding for this chunk?
type ChunkLifecycleAction int32

const (
	// No lifecycle
	ChunkLifecycleAction_CHUNK_LIFECYCLE_ACTION_UNSPECIFIED ChunkLifecycleAction = 0
	// Chunk is in the process of being moved to the read buffer
	ChunkLifecycleAction_CHUNK_LIFECYCLE_ACTION_MOVING ChunkLifecycleAction = 1
	/// Chunk is in the process of being written to object storage
	ChunkLifecycleAction_CHUNK_LIFECYCLE_ACTION_PERSISTING ChunkLifecycleAction = 2
	/// Chunk is in the process of being compacted
	ChunkLifecycleAction_CHUNK_LIFECYCLE_ACTION_COMPACTING ChunkLifecycleAction = 3
	/// Chunk is about to be dropped from memory and (if persisted) from object store.
	ChunkLifecycleAction_CHUNK_LIFECYCLE_ACTION_DROPPING ChunkLifecycleAction = 4
)

// Enum value maps for ChunkLifecycleAction.
var (
	ChunkLifecycleAction_name = map[int32]string{
		0: "CHUNK_LIFECYCLE_ACTION_UNSPECIFIED",
		1: "CHUNK_LIFECYCLE_ACTION_MOVING",
		2: "CHUNK_LIFECYCLE_ACTION_PERSISTING",
		3: "CHUNK_LIFECYCLE_ACTION_COMPACTING",
		4: "CHUNK_LIFECYCLE_ACTION_DROPPING",
	}
	ChunkLifecycleAction_value = map[string]int32{
		"CHUNK_LIFECYCLE_ACTION_UNSPECIFIED": 0,
		"CHUNK_LIFECYCLE_ACTION_MOVING":      1,
		"CHUNK_LIFECYCLE_ACTION_PERSISTING":  2,
		"CHUNK_LIFECYCLE_ACTION_COMPACTING":  3,
		"CHUNK_LIFECYCLE_ACTION_DROPPING":    4,
	}
)

func (x ChunkLifecycleAction) Enum() *ChunkLifecycleAction {
	p := new(ChunkLifecycleAction)
	*p = x
	return p
}

func (x ChunkLifecycleAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChunkLifecycleAction) Descriptor() protoreflect.EnumDescriptor {
	return file_influxdata_iox_management_v1_chunk_proto_enumTypes[1].Descriptor()
}

func (ChunkLifecycleAction) Type() protoreflect.EnumType {
	return &file_influxdata_iox_management_v1_chunk_proto_enumTypes[1]
}

func (x ChunkLifecycleAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChunkLifecycleAction.Descriptor instead.
func (ChunkLifecycleAction) EnumDescriptor() ([]byte, []int) {
	return file_influxdata_iox_management_v1_chunk_proto_rawDescGZIP(), []int{1}
}

// `Chunk` represents part of a partition of data in a database.
// A chunk can contain one or more tables.
type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The partition key of this chunk
	PartitionKey string `protobuf:"bytes,1,opt,name=partition_key,json=partitionKey,proto3" json:"partition_key,omitempty"`
	// The table of this chunk
	TableName string `protobuf:"bytes,8,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// The id of this chunk
	Id uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Which storage system the chunk is located in
	Storage ChunkStorage `protobuf:"varint,3,opt,name=storage,proto3,enum=influxdata.iox.management.v1.ChunkStorage" json:"storage,omitempty"`
	// Is there any outstanding lifecycle action for this chunk?
	LifecycleAction ChunkLifecycleAction `protobuf:"varint,10,opt,name=lifecycle_action,json=lifecycleAction,proto3,enum=influxdata.iox.management.v1.ChunkLifecycleAction" json:"lifecycle_action,omitempty"`
	// The number of bytes used to store this chunk in memory
	MemoryBytes uint64 `protobuf:"varint,4,opt,name=memory_bytes,json=memoryBytes,proto3" json:"memory_bytes,omitempty"`
	// The number of bytes used to store this chunk in object storage
	ObjectStoreBytes uint64 `protobuf:"varint,11,opt,name=object_store_bytes,json=objectStoreBytes,proto3" json:"object_store_bytes,omitempty"`
	// The number of rows in this chunk
	RowCount uint64 `protobuf:"varint,9,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	// The time at which the chunk data was accessed, by a query or a write
	TimeOfLastAccess *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=time_of_last_access,json=timeOfLastAccess,proto3" json:"time_of_last_access,omitempty"`
	// The earliest time at which data contained within this chunk was written
	// into IOx. Note due to the compaction, etc... this may not be the chunk
	// that data was originally written into
	TimeOfFirstWrite *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=time_of_first_write,json=timeOfFirstWrite,proto3" json:"time_of_first_write,omitempty"`
	// The latest time at which data contained within this chunk was written
	// into IOx. Note due to the compaction, etc... this may not be the chunk
	// that data was originally written into
	TimeOfLastWrite *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=time_of_last_write,json=timeOfLastWrite,proto3" json:"time_of_last_write,omitempty"`
	// Time at which this chunk was marked as closed. Note this is not
	// the same as the timestamps on the data itself
	TimeClosed *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=time_closed,json=timeClosed,proto3" json:"time_closed,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_influxdata_iox_management_v1_chunk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_influxdata_iox_management_v1_chunk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_influxdata_iox_management_v1_chunk_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetPartitionKey() string {
	if x != nil {
		return x.PartitionKey
	}
	return ""
}

func (x *Chunk) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *Chunk) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Chunk) GetStorage() ChunkStorage {
	if x != nil {
		return x.Storage
	}
	return ChunkStorage_CHUNK_STORAGE_UNSPECIFIED
}

func (x *Chunk) GetLifecycleAction() ChunkLifecycleAction {
	if x != nil {
		return x.LifecycleAction
	}
	return ChunkLifecycleAction_CHUNK_LIFECYCLE_ACTION_UNSPECIFIED
}

func (x *Chunk) GetMemoryBytes() uint64 {
	if x != nil {
		return x.MemoryBytes
	}
	return 0
}

func (x *Chunk) GetObjectStoreBytes() uint64 {
	if x != nil {
		return x.ObjectStoreBytes
	}
	return 0
}

func (x *Chunk) GetRowCount() uint64 {
	if x != nil {
		return x.RowCount
	}
	return 0
}

func (x *Chunk) GetTimeOfLastAccess() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeOfLastAccess
	}
	return nil
}

func (x *Chunk) GetTimeOfFirstWrite() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeOfFirstWrite
	}
	return nil
}

func (x *Chunk) GetTimeOfLastWrite() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeOfLastWrite
	}
	return nil
}

func (x *Chunk) GetTimeClosed() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeClosed
	}
	return nil
}

var File_influxdata_iox_management_v1_chunk_proto protoreflect.FileDescriptor

var file_influxdata_iox_management_v1_chunk_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6f, 0x78,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x69, 0x6e, 0x66, 0x6c,
	0x75, 0x78, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6f, 0x78, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x05, 0x0a, 0x05, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x44, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75,
	0x78, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6f, 0x78, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x5d, 0x0a,
	0x10, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x69, 0x6f, 0x78, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x6c, 0x69, 0x66,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x2c, 0x0a, 0x12, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x72, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x49, 0x0a, 0x13, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x4c, 0x61, 0x73, 0x74, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x49, 0x0a, 0x13, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66,
	0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10,
	0x74, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x46, 0x69, 0x72, 0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x47, 0x0a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x66,
	0x4c, 0x61, 0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x2a, 0xf1, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x48, 0x55, 0x4e, 0x4b,
	0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x5f,
	0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x4d, 0x55, 0x54,
	0x41, 0x42, 0x4c, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x45, 0x52, 0x10, 0x01, 0x12, 0x27, 0x0a,
	0x23, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x43,
	0x4c, 0x4f, 0x53, 0x45, 0x44, 0x5f, 0x4d, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x42, 0x55,
	0x46, 0x46, 0x45, 0x52, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x5f,
	0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x42, 0x55, 0x46,
	0x46, 0x45, 0x52, 0x10, 0x03, 0x12, 0x2e, 0x0a, 0x2a, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x5f, 0x53,
	0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x42, 0x55, 0x46, 0x46,
	0x45, 0x52, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54,
	0x4f, 0x52, 0x45, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x5f, 0x53,
	0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54,
	0x4f, 0x52, 0x45, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x05, 0x2a, 0xd4, 0x01, 0x0a, 0x14, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x5f, 0x4c, 0x49, 0x46,
	0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x43,
	0x48, 0x55, 0x4e, 0x4b, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x25,
	0x0a, 0x21, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x43, 0x59, 0x43, 0x4c,
	0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x45, 0x52, 0x53, 0x49, 0x53, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x48, 0x55, 0x4e, 0x4b, 0x5f, 0x4c,
	0x49, 0x46, 0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f,
	0x43, 0x48, 0x55, 0x4e, 0x4b, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x52, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x10,
	0x04, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x6c, 0x75, 0x78, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6f, 0x78, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_influxdata_iox_management_v1_chunk_proto_rawDescOnce sync.Once
	file_influxdata_iox_management_v1_chunk_proto_rawDescData = file_influxdata_iox_management_v1_chunk_proto_rawDesc
)

func file_influxdata_iox_management_v1_chunk_proto_rawDescGZIP() []byte {
	file_influxdata_iox_management_v1_chunk_proto_rawDescOnce.Do(func() {
		file_influxdata_iox_management_v1_chunk_proto_rawDescData = protoimpl.X.CompressGZIP(file_influxdata_iox_management_v1_chunk_proto_rawDescData)
	})
	return file_influxdata_iox_management_v1_chunk_proto_rawDescData
}

var file_influxdata_iox_management_v1_chunk_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_influxdata_iox_management_v1_chunk_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_influxdata_iox_management_v1_chunk_proto_goTypes = []interface{}{
	(ChunkStorage)(0),             // 0: influxdata.iox.management.v1.ChunkStorage
	(ChunkLifecycleAction)(0),     // 1: influxdata.iox.management.v1.ChunkLifecycleAction
	(*Chunk)(nil),                 // 2: influxdata.iox.management.v1.Chunk
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_influxdata_iox_management_v1_chunk_proto_depIdxs = []int32{
	0, // 0: influxdata.iox.management.v1.Chunk.storage:type_name -> influxdata.iox.management.v1.ChunkStorage
	1, // 1: influxdata.iox.management.v1.Chunk.lifecycle_action:type_name -> influxdata.iox.management.v1.ChunkLifecycleAction
	3, // 2: influxdata.iox.management.v1.Chunk.time_of_last_access:type_name -> google.protobuf.Timestamp
	3, // 3: influxdata.iox.management.v1.Chunk.time_of_first_write:type_name -> google.protobuf.Timestamp
	3, // 4: influxdata.iox.management.v1.Chunk.time_of_last_write:type_name -> google.protobuf.Timestamp
	3, // 5: influxdata.iox.management.v1.Chunk.time_closed:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_influxdata_iox_management_v1_chunk_proto_init() }
func file_influxdata_iox_management_v1_chunk_proto_init() {
	if File_influxdata_iox_management_v1_chunk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_influxdata_iox_management_v1_chunk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
			RawDescriptor: file_influxdata_iox_management_v1_chunk_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_influxdata_iox_management_v1_chunk_proto_goTypes,
		DependencyIndexes: file_influxdata_iox_management_v1_chunk_proto_depIdxs,
		EnumInfos:         file_influxdata_iox_management_v1_chunk_proto_enumTypes,
		MessageInfos:      file_influxdata_iox_management_v1_chunk_proto_msgTypes,
	}.Build()
	File_influxdata_iox_management_v1_chunk_proto = out.File
	file_influxdata_iox_management_v1_chunk_proto_rawDesc = nil
	file_influxdata_iox_management_v1_chunk_proto_goTypes = nil
	file_influxdata_iox_management_v1_chunk_proto_depIdxs = nil
}
