// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: calculator-service/proto/calc.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//- - - - - - -
//message defined for sum and product
type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  int64 `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber int64 `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_service_proto_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_service_proto_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_calculator_service_proto_calc_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetFirstNumber() int64 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *Input) GetSecondNumber() int64 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input *Input `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_service_proto_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_service_proto_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_calculator_service_proto_calc_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetInput() *Input {
	if x != nil {
		return x.Input
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_service_proto_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_service_proto_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_calculator_service_proto_calc_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

//- - - - - - -
//message defined for server streaming for getting prime factors in stream
type InputDecomposition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *InputDecomposition) Reset() {
	*x = InputDecomposition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_service_proto_calc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputDecomposition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputDecomposition) ProtoMessage() {}

func (x *InputDecomposition) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_service_proto_calc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputDecomposition.ProtoReflect.Descriptor instead.
func (*InputDecomposition) Descriptor() ([]byte, []int) {
	return file_calculator_service_proto_calc_proto_rawDescGZIP(), []int{3}
}

func (x *InputDecomposition) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type RequestDecomposition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input *InputDecomposition `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *RequestDecomposition) Reset() {
	*x = RequestDecomposition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_service_proto_calc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDecomposition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDecomposition) ProtoMessage() {}

func (x *RequestDecomposition) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_service_proto_calc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDecomposition.ProtoReflect.Descriptor instead.
func (*RequestDecomposition) Descriptor() ([]byte, []int) {
	return file_calculator_service_proto_calc_proto_rawDescGZIP(), []int{4}
}

func (x *RequestDecomposition) GetInput() *InputDecomposition {
	if x != nil {
		return x.Input
	}
	return nil
}

type ResponseDecomposition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimeFactor int64 `protobuf:"varint,1,opt,name=prime_factor,json=primeFactor,proto3" json:"prime_factor,omitempty"`
}

func (x *ResponseDecomposition) Reset() {
	*x = ResponseDecomposition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_service_proto_calc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDecomposition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDecomposition) ProtoMessage() {}

func (x *ResponseDecomposition) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_service_proto_calc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDecomposition.ProtoReflect.Descriptor instead.
func (*ResponseDecomposition) Descriptor() ([]byte, []int) {
	return file_calculator_service_proto_calc_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseDecomposition) GetPrimeFactor() int64 {
	if x != nil {
		return x.PrimeFactor
	}
	return 0
}

//- - - - - - -
//message defined for client streaming for getting average
type Number struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_service_proto_calc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_service_proto_calc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Number.ProtoReflect.Descriptor instead.
func (*Number) Descriptor() ([]byte, []int) {
	return file_calculator_service_proto_calc_proto_rawDescGZIP(), []int{6}
}

func (x *Number) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type Average struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num float64 `protobuf:"fixed64,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Average) Reset() {
	*x = Average{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_service_proto_calc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Average) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Average) ProtoMessage() {}

func (x *Average) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_service_proto_calc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Average.ProtoReflect.Descriptor instead.
func (*Average) Descriptor() ([]byte, []int) {
	return file_calculator_service_proto_calc_proto_rawDescGZIP(), []int{7}
}

func (x *Average) GetNum() float64 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_calculator_service_proto_calc_proto protoreflect.FileDescriptor

var file_calculator_service_proto_calc_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x22, 0x4f, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x32, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x22, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2c, 0x0a, 0x12, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x4c, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x34, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x3a, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x22, 0x1a, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x1b,
	0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x32, 0x9a, 0x02, 0x0a, 0x11,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x32, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x13, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x79, 0x12, 0x13, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d,
	0x0a, 0x12, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x21, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12, 0x39, 0x0a,
	0x0a, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x76, 0x63, 0x12, 0x12, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a,
	0x13, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x1a, 0x5a, 0x18, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_service_proto_calc_proto_rawDescOnce sync.Once
	file_calculator_service_proto_calc_proto_rawDescData = file_calculator_service_proto_calc_proto_rawDesc
)

func file_calculator_service_proto_calc_proto_rawDescGZIP() []byte {
	file_calculator_service_proto_calc_proto_rawDescOnce.Do(func() {
		file_calculator_service_proto_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_service_proto_calc_proto_rawDescData)
	})
	return file_calculator_service_proto_calc_proto_rawDescData
}

var file_calculator_service_proto_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_calculator_service_proto_calc_proto_goTypes = []interface{}{
	(*Input)(nil),                 // 0: calculator.Input
	(*Request)(nil),               // 1: calculator.Request
	(*Response)(nil),              // 2: calculator.Response
	(*InputDecomposition)(nil),    // 3: calculator.InputDecomposition
	(*RequestDecomposition)(nil),  // 4: calculator.RequestDecomposition
	(*ResponseDecomposition)(nil), // 5: calculator.ResponseDecomposition
	(*Number)(nil),                // 6: calculator.Number
	(*Average)(nil),               // 7: calculator.Average
}
var file_calculator_service_proto_calc_proto_depIdxs = []int32{
	0, // 0: calculator.Request.input:type_name -> calculator.Input
	3, // 1: calculator.RequestDecomposition.input:type_name -> calculator.InputDecomposition
	1, // 2: calculator.CalculatorService.Sum:input_type -> calculator.Request
	1, // 3: calculator.CalculatorService.Multiply:input_type -> calculator.Request
	4, // 4: calculator.CalculatorService.PrimeDecomposition:input_type -> calculator.RequestDecomposition
	6, // 5: calculator.CalculatorService.AverageSvc:input_type -> calculator.Number
	2, // 6: calculator.CalculatorService.Sum:output_type -> calculator.Response
	2, // 7: calculator.CalculatorService.Multiply:output_type -> calculator.Response
	5, // 8: calculator.CalculatorService.PrimeDecomposition:output_type -> calculator.ResponseDecomposition
	7, // 9: calculator.CalculatorService.AverageSvc:output_type -> calculator.Average
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_calculator_service_proto_calc_proto_init() }
func file_calculator_service_proto_calc_proto_init() {
	if File_calculator_service_proto_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_service_proto_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_calculator_service_proto_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_calculator_service_proto_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_calculator_service_proto_calc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputDecomposition); i {
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
		file_calculator_service_proto_calc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDecomposition); i {
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
		file_calculator_service_proto_calc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseDecomposition); i {
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
		file_calculator_service_proto_calc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Number); i {
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
		file_calculator_service_proto_calc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Average); i {
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
			RawDescriptor: file_calculator_service_proto_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_service_proto_calc_proto_goTypes,
		DependencyIndexes: file_calculator_service_proto_calc_proto_depIdxs,
		MessageInfos:      file_calculator_service_proto_calc_proto_msgTypes,
	}.Build()
	File_calculator_service_proto_calc_proto = out.File
	file_calculator_service_proto_calc_proto_rawDesc = nil
	file_calculator_service_proto_calc_proto_goTypes = nil
	file_calculator_service_proto_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	PrimeDecomposition(ctx context.Context, in *RequestDecomposition, opts ...grpc.CallOption) (CalculatorService_PrimeDecompositionClient, error)
	AverageSvc(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageSvcClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeDecomposition(ctx context.Context, in *RequestDecomposition, opts ...grpc.CallOption) (CalculatorService_PrimeDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeDecompositionClient interface {
	Recv() (*ResponseDecomposition, error)
	grpc.ClientStream
}

type calculatorServicePrimeDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeDecompositionClient) Recv() (*ResponseDecomposition, error) {
	m := new(ResponseDecomposition)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) AverageSvc(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageSvcClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/AverageSvc", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAverageSvcClient{stream}
	return x, nil
}

type CalculatorService_AverageSvcClient interface {
	Send(*Number) error
	CloseAndRecv() (*Average, error)
	grpc.ClientStream
}

type calculatorServiceAverageSvcClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAverageSvcClient) Send(m *Number) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAverageSvcClient) CloseAndRecv() (*Average, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Average)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Sum(context.Context, *Request) (*Response, error)
	Multiply(context.Context, *Request) (*Response, error)
	PrimeDecomposition(*RequestDecomposition, CalculatorService_PrimeDecompositionServer) error
	AverageSvc(CalculatorService_AverageSvcServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) Multiply(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (*UnimplementedCalculatorServiceServer) PrimeDecomposition(*RequestDecomposition, CalculatorService_PrimeDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeDecomposition not implemented")
}
func (*UnimplementedCalculatorServiceServer) AverageSvc(CalculatorService_AverageSvcServer) error {
	return status.Errorf(codes.Unimplemented, "method AverageSvc not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Multiply(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestDecomposition)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeDecomposition(m, &calculatorServicePrimeDecompositionServer{stream})
}

type CalculatorService_PrimeDecompositionServer interface {
	Send(*ResponseDecomposition) error
	grpc.ServerStream
}

type calculatorServicePrimeDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeDecompositionServer) Send(m *ResponseDecomposition) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_AverageSvc_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).AverageSvc(&calculatorServiceAverageSvcServer{stream})
}

type CalculatorService_AverageSvcServer interface {
	SendAndClose(*Average) error
	Recv() (*Number, error)
	grpc.ServerStream
}

type calculatorServiceAverageSvcServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAverageSvcServer) SendAndClose(m *Average) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAverageSvcServer) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _CalculatorService_Multiply_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeDecomposition",
			Handler:       _CalculatorService_PrimeDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AverageSvc",
			Handler:       _CalculatorService_AverageSvc_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator-service/proto/calc.proto",
}
