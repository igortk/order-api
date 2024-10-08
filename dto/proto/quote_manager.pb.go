// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: message/quote_manager.proto

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

type Type int32

const (
	Type_ORDER_DIRECTION_UNDEFINED Type = 0
	Type_ORDER_ALL                 Type = 1
	Type_ORDER_SINGLE              Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "ORDER_DIRECTION_UNDEFINED",
		1: "ORDER_ALL",
		2: "ORDER_SINGLE",
	}
	Type_value = map[string]int32{
		"ORDER_DIRECTION_UNDEFINED": 0,
		"ORDER_ALL":                 1,
		"ORDER_SINGLE":              2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_message_quote_manager_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_message_quote_manager_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_message_quote_manager_proto_rawDescGZIP(), []int{0}
}

type GetExchangeRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pair string `protobuf:"bytes,2,opt,name=pair,proto3" json:"pair,omitempty"`
	Type Type   `protobuf:"varint,3,opt,name=type,proto3,enum=qm.Type" json:"type,omitempty"`
}

func (x *GetExchangeRateRequest) Reset() {
	*x = GetExchangeRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_quote_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExchangeRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExchangeRateRequest) ProtoMessage() {}

func (x *GetExchangeRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_quote_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExchangeRateRequest.ProtoReflect.Descriptor instead.
func (*GetExchangeRateRequest) Descriptor() ([]byte, []int) {
	return file_message_quote_manager_proto_rawDescGZIP(), []int{0}
}

func (x *GetExchangeRateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetExchangeRateRequest) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *GetExchangeRateRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_ORDER_DIRECTION_UNDEFINED
}

type ExchangeRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair    string  `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	Min     float64 `protobuf:"fixed64,2,opt,name=min,proto3" json:"min,omitempty"`
	Average float64 `protobuf:"fixed64,3,opt,name=average,proto3" json:"average,omitempty"`
	Max     float64 `protobuf:"fixed64,4,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *ExchangeRate) Reset() {
	*x = ExchangeRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_quote_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRate) ProtoMessage() {}

func (x *ExchangeRate) ProtoReflect() protoreflect.Message {
	mi := &file_message_quote_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRate.ProtoReflect.Descriptor instead.
func (*ExchangeRate) Descriptor() ([]byte, []int) {
	return file_message_quote_manager_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeRate) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *ExchangeRate) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *ExchangeRate) GetAverage() float64 {
	if x != nil {
		return x.Average
	}
	return 0
}

func (x *ExchangeRate) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

type ExchangeRateList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate []*ExchangeRate `protobuf:"bytes,1,rep,name=rate,proto3" json:"rate,omitempty"`
}

func (x *ExchangeRateList) Reset() {
	*x = ExchangeRateList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_quote_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRateList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRateList) ProtoMessage() {}

func (x *ExchangeRateList) ProtoReflect() protoreflect.Message {
	mi := &file_message_quote_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRateList.ProtoReflect.Descriptor instead.
func (*ExchangeRateList) Descriptor() ([]byte, []int) {
	return file_message_quote_manager_proto_rawDescGZIP(), []int{2}
}

func (x *ExchangeRateList) GetRate() []*ExchangeRate {
	if x != nil {
		return x.Rate
	}
	return nil
}

type GetExchangeRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Rates []*ExchangeRate `protobuf:"bytes,2,rep,name=rates,proto3" json:"rates,omitempty"`
	Error *Error          `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetExchangeRateResponse) Reset() {
	*x = GetExchangeRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_quote_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExchangeRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExchangeRateResponse) ProtoMessage() {}

func (x *GetExchangeRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_quote_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExchangeRateResponse.ProtoReflect.Descriptor instead.
func (*GetExchangeRateResponse) Descriptor() ([]byte, []int) {
	return file_message_quote_manager_proto_rawDescGZIP(), []int{3}
}

func (x *GetExchangeRateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetExchangeRateResponse) GetRates() []*ExchangeRate {
	if x != nil {
		return x.Rates
	}
	return nil
}

func (x *GetExchangeRateResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_message_quote_manager_proto protoreflect.FileDescriptor

var file_message_quote_manager_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x71,
	0x6d, 0x1a, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x69, 0x72, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x08, 0x2e, 0x71, 0x6d, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x60, 0x0a, 0x0c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6d, 0x61, 0x78, 0x22, 0x38, 0x0a, 0x10, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x6d, 0x2e, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x22, 0x75,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x05, 0x72, 0x61, 0x74,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x6d, 0x2e, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x05, 0x72, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x46, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x19, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4f,
	0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_quote_manager_proto_rawDescOnce sync.Once
	file_message_quote_manager_proto_rawDescData = file_message_quote_manager_proto_rawDesc
)

func file_message_quote_manager_proto_rawDescGZIP() []byte {
	file_message_quote_manager_proto_rawDescOnce.Do(func() {
		file_message_quote_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_quote_manager_proto_rawDescData)
	})
	return file_message_quote_manager_proto_rawDescData
}

var file_message_quote_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_quote_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_message_quote_manager_proto_goTypes = []interface{}{
	(Type)(0),                       // 0: qm.Type
	(*GetExchangeRateRequest)(nil),  // 1: qm.GetExchangeRateRequest
	(*ExchangeRate)(nil),            // 2: qm.ExchangeRate
	(*ExchangeRateList)(nil),        // 3: qm.ExchangeRateList
	(*GetExchangeRateResponse)(nil), // 4: qm.GetExchangeRateResponse
	(*Error)(nil),                   // 5: error.Error
}
var file_message_quote_manager_proto_depIdxs = []int32{
	0, // 0: qm.GetExchangeRateRequest.type:type_name -> qm.Type
	2, // 1: qm.ExchangeRateList.rate:type_name -> qm.ExchangeRate
	2, // 2: qm.GetExchangeRateResponse.rates:type_name -> qm.ExchangeRate
	5, // 3: qm.GetExchangeRateResponse.error:type_name -> error.Error
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_message_quote_manager_proto_init() }
func file_message_quote_manager_proto_init() {
	if File_message_quote_manager_proto != nil {
		return
	}
	file_message_error_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_quote_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExchangeRateRequest); i {
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
		file_message_quote_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRate); i {
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
		file_message_quote_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeRateList); i {
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
		file_message_quote_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExchangeRateResponse); i {
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
			RawDescriptor: file_message_quote_manager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_quote_manager_proto_goTypes,
		DependencyIndexes: file_message_quote_manager_proto_depIdxs,
		EnumInfos:         file_message_quote_manager_proto_enumTypes,
		MessageInfos:      file_message_quote_manager_proto_msgTypes,
	}.Build()
	File_message_quote_manager_proto = out.File
	file_message_quote_manager_proto_rawDesc = nil
	file_message_quote_manager_proto_goTypes = nil
	file_message_quote_manager_proto_depIdxs = nil
}
