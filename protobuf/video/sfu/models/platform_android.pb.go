// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: video/sfu/models/platform_android.proto

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

// Different thermal states of an Android device
// Docs: https://source.android.com/docs/core/power/thermal-mitigation#codes
type AndroidThermalStatus int32

const (
	AndroidThermalStatus_ANDROID_THERMAL_STATUS_UNSPECIFIED AndroidThermalStatus = 0
	AndroidThermalStatus_ANDROID_THERMAL_STATUS_NONE        AndroidThermalStatus = 1
	AndroidThermalStatus_ANDROID_THERMAL_STATUS_LIGHT       AndroidThermalStatus = 2
	AndroidThermalStatus_ANDROID_THERMAL_STATUS_MODERATE    AndroidThermalStatus = 3
	AndroidThermalStatus_ANDROID_THERMAL_STATUS_SEVERE      AndroidThermalStatus = 4
	AndroidThermalStatus_ANDROID_THERMAL_STATUS_CRITICAL    AndroidThermalStatus = 5
	AndroidThermalStatus_ANDROID_THERMAL_STATUS_EMERGENCY   AndroidThermalStatus = 6
	AndroidThermalStatus_ANDROID_THERMAL_STATUS_SHUTDOWN    AndroidThermalStatus = 7
)

// Enum value maps for AndroidThermalStatus.
var (
	AndroidThermalStatus_name = map[int32]string{
		0: "ANDROID_THERMAL_STATUS_UNSPECIFIED",
		1: "ANDROID_THERMAL_STATUS_NONE",
		2: "ANDROID_THERMAL_STATUS_LIGHT",
		3: "ANDROID_THERMAL_STATUS_MODERATE",
		4: "ANDROID_THERMAL_STATUS_SEVERE",
		5: "ANDROID_THERMAL_STATUS_CRITICAL",
		6: "ANDROID_THERMAL_STATUS_EMERGENCY",
		7: "ANDROID_THERMAL_STATUS_SHUTDOWN",
	}
	AndroidThermalStatus_value = map[string]int32{
		"ANDROID_THERMAL_STATUS_UNSPECIFIED": 0,
		"ANDROID_THERMAL_STATUS_NONE":        1,
		"ANDROID_THERMAL_STATUS_LIGHT":       2,
		"ANDROID_THERMAL_STATUS_MODERATE":    3,
		"ANDROID_THERMAL_STATUS_SEVERE":      4,
		"ANDROID_THERMAL_STATUS_CRITICAL":    5,
		"ANDROID_THERMAL_STATUS_EMERGENCY":   6,
		"ANDROID_THERMAL_STATUS_SHUTDOWN":    7,
	}
)

func (x AndroidThermalStatus) Enum() *AndroidThermalStatus {
	p := new(AndroidThermalStatus)
	*p = x
	return p
}

func (x AndroidThermalStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AndroidThermalStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_video_sfu_models_platform_android_proto_enumTypes[0].Descriptor()
}

func (AndroidThermalStatus) Type() protoreflect.EnumType {
	return &file_video_sfu_models_platform_android_proto_enumTypes[0]
}

func (x AndroidThermalStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AndroidThermalStatus.Descriptor instead.
func (AndroidThermalStatus) EnumDescriptor() ([]byte, []int) {
	return file_video_sfu_models_platform_android_proto_rawDescGZIP(), []int{0}
}

type AndroidDeviceState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThermalState AndroidThermalStatus `protobuf:"varint,1,opt,name=thermal_state,json=thermalState,proto3,enum=stream.video.sfu.models.AndroidThermalStatus" json:"thermal_state,omitempty"`
}

func (x *AndroidDeviceState) Reset() {
	*x = AndroidDeviceState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_sfu_models_platform_android_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AndroidDeviceState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AndroidDeviceState) ProtoMessage() {}

func (x *AndroidDeviceState) ProtoReflect() protoreflect.Message {
	mi := &file_video_sfu_models_platform_android_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AndroidDeviceState.ProtoReflect.Descriptor instead.
func (*AndroidDeviceState) Descriptor() ([]byte, []int) {
	return file_video_sfu_models_platform_android_proto_rawDescGZIP(), []int{0}
}

func (x *AndroidDeviceState) GetThermalState() AndroidThermalStatus {
	if x != nil {
		return x.ThermalState
	}
	return AndroidThermalStatus_ANDROID_THERMAL_STATUS_UNSPECIFIED
}

var File_video_sfu_models_platform_android_proto protoreflect.FileDescriptor

var file_video_sfu_models_platform_android_proto_rawDesc = []byte{
	0x0a, 0x27, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x66, 0x75, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x61, 0x6e, 0x64, 0x72,
	0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x66, 0x75, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x22, 0x68, 0x0a, 0x12, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x74, 0x68, 0x65, 0x72,
	0x6d, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73,
	0x66, 0x75, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69,
	0x64, 0x54, 0x68, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c,
	0x74, 0x68, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2a, 0xb9, 0x02, 0x0a,
	0x14, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x54, 0x68, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x22, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44,
	0x5f, 0x54, 0x48, 0x45, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a,
	0x1b, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x54, 0x48, 0x45, 0x52, 0x4d, 0x41, 0x4c,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x20,
	0x0a, 0x1c, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x54, 0x48, 0x45, 0x52, 0x4d, 0x41,
	0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x10, 0x02,
	0x12, 0x23, 0x0a, 0x1f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x54, 0x48, 0x45, 0x52,
	0x4d, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x52,
	0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44,
	0x5f, 0x54, 0x48, 0x45, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x53, 0x45, 0x56, 0x45, 0x52, 0x45, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x4e, 0x44, 0x52,
	0x4f, 0x49, 0x44, 0x5f, 0x54, 0x48, 0x45, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x24, 0x0a,
	0x20, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x54, 0x48, 0x45, 0x52, 0x4d, 0x41, 0x4c,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x4e, 0x43,
	0x59, 0x10, 0x06, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x54,
	0x48, 0x45, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x48,
	0x55, 0x54, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x07, 0x42, 0x65, 0x42, 0x0b, 0x53, 0x66, 0x75, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x66, 0x75, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0xaa, 0x02, 0x1a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x66, 0x75, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_sfu_models_platform_android_proto_rawDescOnce sync.Once
	file_video_sfu_models_platform_android_proto_rawDescData = file_video_sfu_models_platform_android_proto_rawDesc
)

func file_video_sfu_models_platform_android_proto_rawDescGZIP() []byte {
	file_video_sfu_models_platform_android_proto_rawDescOnce.Do(func() {
		file_video_sfu_models_platform_android_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_sfu_models_platform_android_proto_rawDescData)
	})
	return file_video_sfu_models_platform_android_proto_rawDescData
}

var file_video_sfu_models_platform_android_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_video_sfu_models_platform_android_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_video_sfu_models_platform_android_proto_goTypes = []interface{}{
	(AndroidThermalStatus)(0),  // 0: stream.video.sfu.models.AndroidThermalStatus
	(*AndroidDeviceState)(nil), // 1: stream.video.sfu.models.AndroidDeviceState
}
var file_video_sfu_models_platform_android_proto_depIdxs = []int32{
	0, // 0: stream.video.sfu.models.AndroidDeviceState.thermal_state:type_name -> stream.video.sfu.models.AndroidThermalStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_video_sfu_models_platform_android_proto_init() }
func file_video_sfu_models_platform_android_proto_init() {
	if File_video_sfu_models_platform_android_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_sfu_models_platform_android_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AndroidDeviceState); i {
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
			RawDescriptor: file_video_sfu_models_platform_android_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_video_sfu_models_platform_android_proto_goTypes,
		DependencyIndexes: file_video_sfu_models_platform_android_proto_depIdxs,
		EnumInfos:         file_video_sfu_models_platform_android_proto_enumTypes,
		MessageInfos:      file_video_sfu_models_platform_android_proto_msgTypes,
	}.Build()
	File_video_sfu_models_platform_android_proto = out.File
	file_video_sfu_models_platform_android_proto_rawDesc = nil
	file_video_sfu_models_platform_android_proto_goTypes = nil
	file_video_sfu_models_platform_android_proto_depIdxs = nil
}
