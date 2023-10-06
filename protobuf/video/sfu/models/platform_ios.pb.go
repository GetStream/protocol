// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: video/sfu/models/platform_ios.proto

package models

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

// Represents the thermal state of the device running iOS SDKs. The full list
// of values is documented at https://developer.apple.com/documentation/foundation/processinfo/thermalstate
type IOsThermalState int32

const (
	IOsThermalState_IOS_THERMAL_STATE_UNKNOWN  IOsThermalState = 0
	IOsThermalState_IOS_THERMAL_STATE_NOMIMAL  IOsThermalState = 1
	IOsThermalState_IOS_THERMAL_STATE_FAIR     IOsThermalState = 2
	IOsThermalState_IOS_THERMAL_STATE_SERIOUS  IOsThermalState = 3
	IOsThermalState_IOS_THERMAL_STATE_CRITICAL IOsThermalState = 4
)

// Enum value maps for IOsThermalState.
var (
	IOsThermalState_name = map[int32]string{
		0: "IOS_THERMAL_STATE_UNKNOWN",
		1: "IOS_THERMAL_STATE_NOMIMAL",
		2: "IOS_THERMAL_STATE_FAIR",
		3: "IOS_THERMAL_STATE_SERIOUS",
		4: "IOS_THERMAL_STATE_CRITICAL",
	}
	IOsThermalState_value = map[string]int32{
		"IOS_THERMAL_STATE_UNKNOWN":  0,
		"IOS_THERMAL_STATE_NOMIMAL":  1,
		"IOS_THERMAL_STATE_FAIR":     2,
		"IOS_THERMAL_STATE_SERIOUS":  3,
		"IOS_THERMAL_STATE_CRITICAL": 4,
	}
)

func (x IOsThermalState) Enum() *IOsThermalState {
	p := new(IOsThermalState)
	*p = x
	return p
}

func (x IOsThermalState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IOsThermalState) Descriptor() protoreflect.EnumDescriptor {
	return file_video_sfu_models_platform_ios_proto_enumTypes[0].Descriptor()
}

func (IOsThermalState) Type() protoreflect.EnumType {
	return &file_video_sfu_models_platform_ios_proto_enumTypes[0]
}

func (x IOsThermalState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IOsThermalState.Descriptor instead.
func (IOsThermalState) EnumDescriptor() ([]byte, []int) {
	return file_video_sfu_models_platform_ios_proto_rawDescGZIP(), []int{0}
}

type IOsDeviceState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThermalState IOsThermalState `protobuf:"varint,1,opt,name=thermal_state,json=thermalState,proto3,enum=stream.video.sfu.models.IOsThermalState" json:"thermal_state,omitempty"`
	// https://developer.apple.com/documentation/foundation/processinfo/1617047-islowpowermodeenabled
	IsLowPowerModeEnabled bool `protobuf:"varint,2,opt,name=is_low_power_mode_enabled,json=isLowPowerModeEnabled,proto3" json:"is_low_power_mode_enabled,omitempty"`
}

func (x *IOsDeviceState) Reset() {
	*x = IOsDeviceState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_sfu_models_platform_ios_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IOsDeviceState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOsDeviceState) ProtoMessage() {}

func (x *IOsDeviceState) ProtoReflect() protoreflect.Message {
	mi := &file_video_sfu_models_platform_ios_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOsDeviceState.ProtoReflect.Descriptor instead.
func (*IOsDeviceState) Descriptor() ([]byte, []int) {
	return file_video_sfu_models_platform_ios_proto_rawDescGZIP(), []int{0}
}

func (x *IOsDeviceState) GetThermalState() IOsThermalState {
	if x != nil {
		return x.ThermalState
	}
	return IOsThermalState_IOS_THERMAL_STATE_UNKNOWN
}

func (x *IOsDeviceState) GetIsLowPowerModeEnabled() bool {
	if x != nil {
		return x.IsLowPowerModeEnabled
	}
	return false
}

var File_video_sfu_models_platform_ios_proto protoreflect.FileDescriptor

var file_video_sfu_models_platform_ios_proto_rawDesc = []byte{
	0x0a, 0x23, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x66, 0x75, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x6f, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x73, 0x66, 0x75, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x99,
	0x01, 0x0a, 0x0e, 0x49, 0x4f, 0x73, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x4d, 0x0a, 0x0d, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x66, 0x75, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x49, 0x4f, 0x73, 0x54, 0x68, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x0c, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x38, 0x0a, 0x19, 0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x73, 0x4c, 0x6f, 0x77, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4d,
	0x6f, 0x64, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x2a, 0xaa, 0x01, 0x0a, 0x0f, 0x49,
	0x4f, 0x73, 0x54, 0x68, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x0a, 0x19, 0x49, 0x4f, 0x53, 0x5f, 0x54, 0x48, 0x45, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1d, 0x0a,
	0x19, 0x49, 0x4f, 0x53, 0x5f, 0x54, 0x48, 0x45, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4d, 0x49, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16,
	0x49, 0x4f, 0x53, 0x5f, 0x54, 0x48, 0x45, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x46, 0x41, 0x49, 0x52, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4f, 0x53, 0x5f,
	0x54, 0x48, 0x45, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x45,
	0x52, 0x49, 0x4f, 0x55, 0x53, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4f, 0x53, 0x5f, 0x54,
	0x48, 0x45, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x49,
	0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x04, 0x42, 0x65, 0x42, 0x0b, 0x53, 0x66, 0x75, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x66, 0x75, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0xaa, 0x02, 0x1a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x66, 0x75, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_sfu_models_platform_ios_proto_rawDescOnce sync.Once
	file_video_sfu_models_platform_ios_proto_rawDescData = file_video_sfu_models_platform_ios_proto_rawDesc
)

func file_video_sfu_models_platform_ios_proto_rawDescGZIP() []byte {
	file_video_sfu_models_platform_ios_proto_rawDescOnce.Do(func() {
		file_video_sfu_models_platform_ios_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_sfu_models_platform_ios_proto_rawDescData)
	})
	return file_video_sfu_models_platform_ios_proto_rawDescData
}

var file_video_sfu_models_platform_ios_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_video_sfu_models_platform_ios_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_video_sfu_models_platform_ios_proto_goTypes = []interface{}{
	(IOsThermalState)(0),   // 0: stream.video.sfu.models.IOsThermalState
	(*IOsDeviceState)(nil), // 1: stream.video.sfu.models.IOsDeviceState
}
var file_video_sfu_models_platform_ios_proto_depIdxs = []int32{
	0, // 0: stream.video.sfu.models.IOsDeviceState.thermal_state:type_name -> stream.video.sfu.models.IOsThermalState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_video_sfu_models_platform_ios_proto_init() }
func file_video_sfu_models_platform_ios_proto_init() {
	if File_video_sfu_models_platform_ios_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_sfu_models_platform_ios_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IOsDeviceState); i {
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
			RawDescriptor: file_video_sfu_models_platform_ios_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_video_sfu_models_platform_ios_proto_goTypes,
		DependencyIndexes: file_video_sfu_models_platform_ios_proto_depIdxs,
		EnumInfos:         file_video_sfu_models_platform_ios_proto_enumTypes,
		MessageInfos:      file_video_sfu_models_platform_ios_proto_msgTypes,
	}.Build()
	File_video_sfu_models_platform_ios_proto = out.File
	file_video_sfu_models_platform_ios_proto_rawDesc = nil
	file_video_sfu_models_platform_ios_proto_goTypes = nil
	file_video_sfu_models_platform_ios_proto_depIdxs = nil
}
