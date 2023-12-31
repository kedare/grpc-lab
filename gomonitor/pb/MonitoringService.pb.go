// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: MonitoringService.proto

package pb

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

// The CPU usage information request
type CpuUsageInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The interval in seconds between each CPU usage information update
	// The average during this time interval will be returned
	Interval int32 `protobuf:"varint,1,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *CpuUsageInfoRequest) Reset() {
	*x = CpuUsageInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MonitoringService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuUsageInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuUsageInfoRequest) ProtoMessage() {}

func (x *CpuUsageInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_MonitoringService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuUsageInfoRequest.ProtoReflect.Descriptor instead.
func (*CpuUsageInfoRequest) Descriptor() ([]byte, []int) {
	return file_MonitoringService_proto_rawDescGZIP(), []int{0}
}

func (x *CpuUsageInfoRequest) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

// The CPU usage information response
type CpuUsageInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The CPU system time (used by kernel)
	SystemTime int32 `protobuf:"varint,1,opt,name=system_time,json=systemTime,proto3" json:"system_time,omitempty"`
	// The CPU user time (used by user processes)
	UserTime int32 `protobuf:"varint,2,opt,name=user_time,json=userTime,proto3" json:"user_time,omitempty"`
	// The CPU idle time (not used)
	IdleTime int32 `protobuf:"varint,3,opt,name=idle_time,json=idleTime,proto3" json:"idle_time,omitempty"`
	// The CPU I/O wait time (waiting for I/O operations to complete)
	IowaitTime int32 `protobuf:"varint,4,opt,name=iowait_time,json=iowaitTime,proto3" json:"iowait_time,omitempty"`
	// The CPU IRQ time (servicing interrupts)
	IrqTime int32 `protobuf:"varint,5,opt,name=irq_time,json=irqTime,proto3" json:"irq_time,omitempty"`
	// The CPU soft IRQ time (servicing soft interrupts)
	SoftirqTime int32 `protobuf:"varint,6,opt,name=softirq_time,json=softirqTime,proto3" json:"softirq_time,omitempty"`
	// The CPU steal time (time spent in other operating systems when running in a virtualized environment)
	StealTime int32 `protobuf:"varint,7,opt,name=steal_time,json=stealTime,proto3" json:"steal_time,omitempty"`
	// The CPU guest time (time spent running a virtual CPU for guest operating systems under the control of the Linux kernel)
	GuestTime int32 `protobuf:"varint,8,opt,name=guest_time,json=guestTime,proto3" json:"guest_time,omitempty"`
	// The CPU guest nice time (time spent running a niced guest (virtual CPU for guest operating systems under the control of the Linux kernel))
	GuestNiceTime int32 `protobuf:"varint,9,opt,name=guest_nice_time,json=guestNiceTime,proto3" json:"guest_nice_time,omitempty"`
}

func (x *CpuUsageInfoResponse) Reset() {
	*x = CpuUsageInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MonitoringService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuUsageInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuUsageInfoResponse) ProtoMessage() {}

func (x *CpuUsageInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_MonitoringService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuUsageInfoResponse.ProtoReflect.Descriptor instead.
func (*CpuUsageInfoResponse) Descriptor() ([]byte, []int) {
	return file_MonitoringService_proto_rawDescGZIP(), []int{1}
}

func (x *CpuUsageInfoResponse) GetSystemTime() int32 {
	if x != nil {
		return x.SystemTime
	}
	return 0
}

func (x *CpuUsageInfoResponse) GetUserTime() int32 {
	if x != nil {
		return x.UserTime
	}
	return 0
}

func (x *CpuUsageInfoResponse) GetIdleTime() int32 {
	if x != nil {
		return x.IdleTime
	}
	return 0
}

func (x *CpuUsageInfoResponse) GetIowaitTime() int32 {
	if x != nil {
		return x.IowaitTime
	}
	return 0
}

func (x *CpuUsageInfoResponse) GetIrqTime() int32 {
	if x != nil {
		return x.IrqTime
	}
	return 0
}

func (x *CpuUsageInfoResponse) GetSoftirqTime() int32 {
	if x != nil {
		return x.SoftirqTime
	}
	return 0
}

func (x *CpuUsageInfoResponse) GetStealTime() int32 {
	if x != nil {
		return x.StealTime
	}
	return 0
}

func (x *CpuUsageInfoResponse) GetGuestTime() int32 {
	if x != nil {
		return x.GuestTime
	}
	return 0
}

func (x *CpuUsageInfoResponse) GetGuestNiceTime() int32 {
	if x != nil {
		return x.GuestNiceTime
	}
	return 0
}

var File_MonitoringService_proto protoreflect.FileDescriptor

var file_MonitoringService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x31, 0x0a, 0x13, 0x43, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x22, 0xb6, 0x02, 0x0a, 0x14, 0x43, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x64,
	0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69,
	0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6f, 0x77, 0x61, 0x69,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6f,
	0x77, 0x61, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x72, 0x71, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x72, 0x71, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x66, 0x74, 0x69, 0x72, 0x71, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x6f, 0x66, 0x74, 0x69,
	0x72, 0x71, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x65, 0x61, 0x6c, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x65, 0x61,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x67, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x69,
	0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x4e, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xb4, 0x01, 0x0a,
	0x11, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x70,
	0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x70, 0x75, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x51, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x70, 0x75, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x70, 0x75, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x1a, 0x5a, 0x03, 0x70, 0x62, 0x2f, 0xaa, 0x02, 0x12, 0x42, 0x6c, 0x61,
	0x7a, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MonitoringService_proto_rawDescOnce sync.Once
	file_MonitoringService_proto_rawDescData = file_MonitoringService_proto_rawDesc
)

func file_MonitoringService_proto_rawDescGZIP() []byte {
	file_MonitoringService_proto_rawDescOnce.Do(func() {
		file_MonitoringService_proto_rawDescData = protoimpl.X.CompressGZIP(file_MonitoringService_proto_rawDescData)
	})
	return file_MonitoringService_proto_rawDescData
}

var file_MonitoringService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_MonitoringService_proto_goTypes = []interface{}{
	(*CpuUsageInfoRequest)(nil),  // 0: proto.CpuUsageInfoRequest
	(*CpuUsageInfoResponse)(nil), // 1: proto.CpuUsageInfoResponse
}
var file_MonitoringService_proto_depIdxs = []int32{
	0, // 0: proto.MonitoringService.GetCpuUsageInfo:input_type -> proto.CpuUsageInfoRequest
	0, // 1: proto.MonitoringService.StreamCpuUsageInfo:input_type -> proto.CpuUsageInfoRequest
	1, // 2: proto.MonitoringService.GetCpuUsageInfo:output_type -> proto.CpuUsageInfoResponse
	1, // 3: proto.MonitoringService.StreamCpuUsageInfo:output_type -> proto.CpuUsageInfoResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MonitoringService_proto_init() }
func file_MonitoringService_proto_init() {
	if File_MonitoringService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MonitoringService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuUsageInfoRequest); i {
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
		file_MonitoringService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuUsageInfoResponse); i {
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
			RawDescriptor: file_MonitoringService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_MonitoringService_proto_goTypes,
		DependencyIndexes: file_MonitoringService_proto_depIdxs,
		MessageInfos:      file_MonitoringService_proto_msgTypes,
	}.Build()
	File_MonitoringService_proto = out.File
	file_MonitoringService_proto_rawDesc = nil
	file_MonitoringService_proto_goTypes = nil
	file_MonitoringService_proto_depIdxs = nil
}
