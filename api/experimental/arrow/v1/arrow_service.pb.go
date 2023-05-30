// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: opentelemetry/proto/experimental/arrow/v1/arrow_service.proto

package v1

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

// Enumeration of all the OTel Arrow payload types currently supported by the
// OTel Arrow protocol.
type ArrowPayloadType int32

const (
	ArrowPayloadType_UNKNOWN ArrowPayloadType = 0
	// A payload representing a collection of resource attributes.
	ArrowPayloadType_RESOURCE_ATTRS ArrowPayloadType = 1
	// A payload representing a collection of scope attributes.
	ArrowPayloadType_SCOPE_ATTRS ArrowPayloadType = 2
	// A set of payloads representing a collection of metrics.
	ArrowPayloadType_METRICS                         ArrowPayloadType = 10 // Main metric payload
	ArrowPayloadType_NUMBER_DATA_POINTS              ArrowPayloadType = 11
	ArrowPayloadType_SUMMARY_DATA_POINTS             ArrowPayloadType = 12
	ArrowPayloadType_HISTOGRAM_DATA_POINTS           ArrowPayloadType = 13
	ArrowPayloadType_EXP_HISTOGRAM_DATA_POINTS       ArrowPayloadType = 14
	ArrowPayloadType_NUMBER_DP_ATTRS                 ArrowPayloadType = 15
	ArrowPayloadType_SUMMARY_DP_ATTRS                ArrowPayloadType = 16
	ArrowPayloadType_HISTOGRAM_DP_ATTRS              ArrowPayloadType = 17
	ArrowPayloadType_EXP_HISTOGRAM_DP_ATTRS          ArrowPayloadType = 18
	ArrowPayloadType_NUMBER_DP_EXEMPLARS             ArrowPayloadType = 19
	ArrowPayloadType_HISTOGRAM_DP_EXEMPLARS          ArrowPayloadType = 20
	ArrowPayloadType_EXP_HISTOGRAM_DP_EXEMPLARS      ArrowPayloadType = 21
	ArrowPayloadType_NUMBER_DP_EXEMPLAR_ATTRS        ArrowPayloadType = 22
	ArrowPayloadType_HISTOGRAM_DP_EXEMPLAR_ATTRS     ArrowPayloadType = 23
	ArrowPayloadType_EXP_HISTOGRAM_DP_EXEMPLAR_ATTRS ArrowPayloadType = 24
	// A set of payloads representing a collection of logs.
	ArrowPayloadType_LOGS      ArrowPayloadType = 30
	ArrowPayloadType_LOG_ATTRS ArrowPayloadType = 31
	// A set of payloads representing a collection of traces.
	ArrowPayloadType_SPANS            ArrowPayloadType = 40
	ArrowPayloadType_SPAN_ATTRS       ArrowPayloadType = 41
	ArrowPayloadType_SPAN_EVENTS      ArrowPayloadType = 42
	ArrowPayloadType_SPAN_LINKS       ArrowPayloadType = 43
	ArrowPayloadType_SPAN_EVENT_ATTRS ArrowPayloadType = 44
	ArrowPayloadType_SPAN_LINK_ATTRS  ArrowPayloadType = 45
)

// Enum value maps for ArrowPayloadType.
var (
	ArrowPayloadType_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "RESOURCE_ATTRS",
		2:  "SCOPE_ATTRS",
		10: "METRICS",
		11: "NUMBER_DATA_POINTS",
		12: "SUMMARY_DATA_POINTS",
		13: "HISTOGRAM_DATA_POINTS",
		14: "EXP_HISTOGRAM_DATA_POINTS",
		15: "NUMBER_DP_ATTRS",
		16: "SUMMARY_DP_ATTRS",
		17: "HISTOGRAM_DP_ATTRS",
		18: "EXP_HISTOGRAM_DP_ATTRS",
		19: "NUMBER_DP_EXEMPLARS",
		20: "HISTOGRAM_DP_EXEMPLARS",
		21: "EXP_HISTOGRAM_DP_EXEMPLARS",
		22: "NUMBER_DP_EXEMPLAR_ATTRS",
		23: "HISTOGRAM_DP_EXEMPLAR_ATTRS",
		24: "EXP_HISTOGRAM_DP_EXEMPLAR_ATTRS",
		30: "LOGS",
		31: "LOG_ATTRS",
		40: "SPANS",
		41: "SPAN_ATTRS",
		42: "SPAN_EVENTS",
		43: "SPAN_LINKS",
		44: "SPAN_EVENT_ATTRS",
		45: "SPAN_LINK_ATTRS",
	}
	ArrowPayloadType_value = map[string]int32{
		"UNKNOWN":                         0,
		"RESOURCE_ATTRS":                  1,
		"SCOPE_ATTRS":                     2,
		"METRICS":                         10,
		"NUMBER_DATA_POINTS":              11,
		"SUMMARY_DATA_POINTS":             12,
		"HISTOGRAM_DATA_POINTS":           13,
		"EXP_HISTOGRAM_DATA_POINTS":       14,
		"NUMBER_DP_ATTRS":                 15,
		"SUMMARY_DP_ATTRS":                16,
		"HISTOGRAM_DP_ATTRS":              17,
		"EXP_HISTOGRAM_DP_ATTRS":          18,
		"NUMBER_DP_EXEMPLARS":             19,
		"HISTOGRAM_DP_EXEMPLARS":          20,
		"EXP_HISTOGRAM_DP_EXEMPLARS":      21,
		"NUMBER_DP_EXEMPLAR_ATTRS":        22,
		"HISTOGRAM_DP_EXEMPLAR_ATTRS":     23,
		"EXP_HISTOGRAM_DP_EXEMPLAR_ATTRS": 24,
		"LOGS":                            30,
		"LOG_ATTRS":                       31,
		"SPANS":                           40,
		"SPAN_ATTRS":                      41,
		"SPAN_EVENTS":                     42,
		"SPAN_LINKS":                      43,
		"SPAN_EVENT_ATTRS":                44,
		"SPAN_LINK_ATTRS":                 45,
	}
)

func (x ArrowPayloadType) Enum() *ArrowPayloadType {
	p := new(ArrowPayloadType)
	*p = x
	return p
}

func (x ArrowPayloadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArrowPayloadType) Descriptor() protoreflect.EnumDescriptor {
	return file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_enumTypes[0].Descriptor()
}

func (ArrowPayloadType) Type() protoreflect.EnumType {
	return &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_enumTypes[0]
}

func (x ArrowPayloadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArrowPayloadType.Descriptor instead.
func (ArrowPayloadType) EnumDescriptor() ([]byte, []int) {
	return file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescGZIP(), []int{0}
}

type StatusCode int32

const (
	StatusCode_OK    StatusCode = 0
	StatusCode_ERROR StatusCode = 1
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	StatusCode_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_enumTypes[1].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_enumTypes[1]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescGZIP(), []int{1}
}

type ErrorCode int32

const (
	ErrorCode_UNAVAILABLE      ErrorCode = 0
	ErrorCode_INVALID_ARGUMENT ErrorCode = 1
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "UNAVAILABLE",
		1: "INVALID_ARGUMENT",
	}
	ErrorCode_value = map[string]int32{
		"UNAVAILABLE":      0,
		"INVALID_ARGUMENT": 1,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_enumTypes[2].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_enumTypes[2]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescGZIP(), []int{2}
}

// A message sent by an exporter to a collector containing a batch of Arrow
// records.
type BatchArrowRecords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [mandatory] Batch ID. Must be unique in the context of the stream.
	BatchId string `protobuf:"bytes,1,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	// [mandatory] A collection of payloads containing the data of the batch.
	ArrowPayloads []*ArrowPayload `protobuf:"bytes,2,rep,name=arrow_payloads,json=arrowPayloads,proto3" json:"arrow_payloads,omitempty"`
	// [optional] Headers associated with this batch, encoded using hpack.
	Headers []byte `protobuf:"bytes,3,opt,name=headers,proto3" json:"headers,omitempty"`
}

func (x *BatchArrowRecords) Reset() {
	*x = BatchArrowRecords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchArrowRecords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchArrowRecords) ProtoMessage() {}

func (x *BatchArrowRecords) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchArrowRecords.ProtoReflect.Descriptor instead.
func (*BatchArrowRecords) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescGZIP(), []int{0}
}

func (x *BatchArrowRecords) GetBatchId() string {
	if x != nil {
		return x.BatchId
	}
	return ""
}

func (x *BatchArrowRecords) GetArrowPayloads() []*ArrowPayload {
	if x != nil {
		return x.ArrowPayloads
	}
	return nil
}

func (x *BatchArrowRecords) GetHeaders() []byte {
	if x != nil {
		return x.Headers
	}
	return nil
}

// Represents a batch of OTel Arrow entities.
type ArrowPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [mandatory] A unique id assigned to a sub-stream of the batch sharing the
	// same schema, and dictionaries.
	SubStreamId string `protobuf:"bytes,1,opt,name=sub_stream_id,json=subStreamId,proto3" json:"sub_stream_id,omitempty"`
	// [mandatory] Type of the OTel Arrow payload.
	Type ArrowPayloadType `protobuf:"varint,2,opt,name=type,proto3,enum=opentelemetry.proto.experimental.arrow.v1.ArrowPayloadType" json:"type,omitempty"`
	// [mandatory] Serialized Arrow Record Batch
	// For a description of the Arrow IPC format see:
	// https://arrow.apache.org/docs/format/Columnar.html#serialization-and-interprocess-communication-ipc
	Record []byte `protobuf:"bytes,3,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *ArrowPayload) Reset() {
	*x = ArrowPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrowPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrowPayload) ProtoMessage() {}

func (x *ArrowPayload) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrowPayload.ProtoReflect.Descriptor instead.
func (*ArrowPayload) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescGZIP(), []int{1}
}

func (x *ArrowPayload) GetSubStreamId() string {
	if x != nil {
		return x.SubStreamId
	}
	return ""
}

func (x *ArrowPayload) GetType() ArrowPayloadType {
	if x != nil {
		return x.Type
	}
	return ArrowPayloadType_UNKNOWN
}

func (x *ArrowPayload) GetRecord() []byte {
	if x != nil {
		return x.Record
	}
	return nil
}

// A message sent by a Collector to the exporter that opened the data stream.
type BatchStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statuses []*StatusMessage `protobuf:"bytes,1,rep,name=statuses,proto3" json:"statuses,omitempty"`
}

func (x *BatchStatus) Reset() {
	*x = BatchStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchStatus) ProtoMessage() {}

func (x *BatchStatus) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchStatus.ProtoReflect.Descriptor instead.
func (*BatchStatus) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescGZIP(), []int{2}
}

func (x *BatchStatus) GetStatuses() []*StatusMessage {
	if x != nil {
		return x.Statuses
	}
	return nil
}

type StatusMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchId      string     `protobuf:"bytes,1,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	StatusCode   StatusCode `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3,enum=opentelemetry.proto.experimental.arrow.v1.StatusCode" json:"status_code,omitempty"`
	ErrorCode    ErrorCode  `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3,enum=opentelemetry.proto.experimental.arrow.v1.ErrorCode" json:"error_code,omitempty"`
	ErrorMessage string     `protobuf:"bytes,4,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	RetryInfo    *RetryInfo `protobuf:"bytes,5,opt,name=retry_info,json=retryInfo,proto3" json:"retry_info,omitempty"`
}

func (x *StatusMessage) Reset() {
	*x = StatusMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusMessage) ProtoMessage() {}

func (x *StatusMessage) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusMessage.ProtoReflect.Descriptor instead.
func (*StatusMessage) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescGZIP(), []int{3}
}

func (x *StatusMessage) GetBatchId() string {
	if x != nil {
		return x.BatchId
	}
	return ""
}

func (x *StatusMessage) GetStatusCode() StatusCode {
	if x != nil {
		return x.StatusCode
	}
	return StatusCode_OK
}

func (x *StatusMessage) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_UNAVAILABLE
}

func (x *StatusMessage) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *StatusMessage) GetRetryInfo() *RetryInfo {
	if x != nil {
		return x.RetryInfo
	}
	return nil
}

type RetryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetryDelay int64 `protobuf:"varint,1,opt,name=retry_delay,json=retryDelay,proto3" json:"retry_delay,omitempty"`
}

func (x *RetryInfo) Reset() {
	*x = RetryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryInfo) ProtoMessage() {}

func (x *RetryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryInfo.ProtoReflect.Descriptor instead.
func (*RetryInfo) Descriptor() ([]byte, []int) {
	return file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescGZIP(), []int{4}
}

func (x *RetryInfo) GetRetryDelay() int64 {
	if x != nil {
		return x.RetryDelay
	}
	return 0
}

var File_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto protoreflect.FileDescriptor

var file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2f, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x72, 0x6f,
	0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x29, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x22, 0xa8, 0x01, 0x0a, 0x11, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x5e, 0x0a, 0x0e, 0x61,
	0x72, 0x72, 0x6f, 0x77, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x72, 0x72, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0d, 0x61, 0x72,
	0x72, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x0c, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x75, 0x62, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x72, 0x72, 0x6f,
	0x77, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x22, 0x63, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x54, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0xd1, 0x02, 0x0a, 0x0d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x56, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x72,
	0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x53, 0x0a,
	0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x34, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x61,
	0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x09, 0x72, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2c, 0x0a, 0x09,
	0x52, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74,
	0x72, 0x79, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x72, 0x65, 0x74, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x2a, 0xd4, 0x04, 0x0a, 0x10, 0x41,
	0x72, 0x72, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x53, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x53, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x10, 0x0a, 0x12, 0x16,
	0x0a, 0x12, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x50, 0x4f,
	0x49, 0x4e, 0x54, 0x53, 0x10, 0x0b, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52,
	0x59, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x10, 0x0c, 0x12,
	0x19, 0x0a, 0x15, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x5f, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x10, 0x0d, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x58,
	0x50, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x5f, 0x44, 0x41, 0x54, 0x41,
	0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x53, 0x10, 0x0e, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x55, 0x4d,
	0x42, 0x45, 0x52, 0x5f, 0x44, 0x50, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x53, 0x10, 0x0f, 0x12, 0x14,
	0x0a, 0x10, 0x53, 0x55, 0x4d, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x44, 0x50, 0x5f, 0x41, 0x54, 0x54,
	0x52, 0x53, 0x10, 0x10, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x47, 0x52, 0x41,
	0x4d, 0x5f, 0x44, 0x50, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x53, 0x10, 0x11, 0x12, 0x1a, 0x0a, 0x16,
	0x45, 0x58, 0x50, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x5f, 0x44, 0x50,
	0x5f, 0x41, 0x54, 0x54, 0x52, 0x53, 0x10, 0x12, 0x12, 0x17, 0x0a, 0x13, 0x4e, 0x55, 0x4d, 0x42,
	0x45, 0x52, 0x5f, 0x44, 0x50, 0x5f, 0x45, 0x58, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x52, 0x53, 0x10,
	0x13, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x5f, 0x44,
	0x50, 0x5f, 0x45, 0x58, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x52, 0x53, 0x10, 0x14, 0x12, 0x1e, 0x0a,
	0x1a, 0x45, 0x58, 0x50, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x5f, 0x44,
	0x50, 0x5f, 0x45, 0x58, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x52, 0x53, 0x10, 0x15, 0x12, 0x1c, 0x0a,
	0x18, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x44, 0x50, 0x5f, 0x45, 0x58, 0x45, 0x4d, 0x50,
	0x4c, 0x41, 0x52, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x53, 0x10, 0x16, 0x12, 0x1f, 0x0a, 0x1b, 0x48,
	0x49, 0x53, 0x54, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x5f, 0x44, 0x50, 0x5f, 0x45, 0x58, 0x45, 0x4d,
	0x50, 0x4c, 0x41, 0x52, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x53, 0x10, 0x17, 0x12, 0x23, 0x0a, 0x1f,
	0x45, 0x58, 0x50, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x47, 0x52, 0x41, 0x4d, 0x5f, 0x44, 0x50,
	0x5f, 0x45, 0x58, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x52, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x53, 0x10,
	0x18, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x47, 0x53, 0x10, 0x1e, 0x12, 0x0d, 0x0a, 0x09, 0x4c,
	0x4f, 0x47, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x53, 0x10, 0x1f, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x50,
	0x41, 0x4e, 0x53, 0x10, 0x28, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x41, 0x54,
	0x54, 0x52, 0x53, 0x10, 0x29, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x53, 0x10, 0x2a, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x4c,
	0x49, 0x4e, 0x4b, 0x53, 0x10, 0x2b, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x53, 0x10, 0x2c, 0x12, 0x13, 0x0a, 0x0f,
	0x53, 0x50, 0x41, 0x4e, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x53, 0x10,
	0x2d, 0x2a, 0x1f, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x01, 0x2a, 0x32, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x32, 0xa0, 0x01, 0x0a, 0x12, 0x41, 0x72, 0x72, 0x6f, 0x77,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01,
	0x0a, 0x0b, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x3c, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x41,
	0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x36, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x61,
	0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0xa0, 0x01, 0x0a, 0x12, 0x41, 0x72,
	0x72, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x89, 0x01, 0x0a, 0x0b, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73,
	0x12, 0x3c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x36,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0x9c, 0x01, 0x0a,
	0x10, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x87, 0x01, 0x0a, 0x09, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x12,
	0x3c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x36, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0xa2, 0x01, 0x0a, 0x13,
	0x41, 0x72, 0x72, 0x6f, 0x77, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x0c, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x3c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x1a, 0x36, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x7f, 0x0a, 0x2c, 0x69, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31,
	0x42, 0x11, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x35, 0x2f, 0x6f, 0x74, 0x65, 0x6c, 0x2d, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2d,
	0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescOnce sync.Once
	file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescData = file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDesc
)

func file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescGZIP() []byte {
	file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescOnce.Do(func() {
		file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescData)
	})
	return file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDescData
}

var file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_goTypes = []interface{}{
	(ArrowPayloadType)(0),     // 0: opentelemetry.proto.experimental.arrow.v1.ArrowPayloadType
	(StatusCode)(0),           // 1: opentelemetry.proto.experimental.arrow.v1.StatusCode
	(ErrorCode)(0),            // 2: opentelemetry.proto.experimental.arrow.v1.ErrorCode
	(*BatchArrowRecords)(nil), // 3: opentelemetry.proto.experimental.arrow.v1.BatchArrowRecords
	(*ArrowPayload)(nil),      // 4: opentelemetry.proto.experimental.arrow.v1.ArrowPayload
	(*BatchStatus)(nil),       // 5: opentelemetry.proto.experimental.arrow.v1.BatchStatus
	(*StatusMessage)(nil),     // 6: opentelemetry.proto.experimental.arrow.v1.StatusMessage
	(*RetryInfo)(nil),         // 7: opentelemetry.proto.experimental.arrow.v1.RetryInfo
}
var file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_depIdxs = []int32{
	4,  // 0: opentelemetry.proto.experimental.arrow.v1.BatchArrowRecords.arrow_payloads:type_name -> opentelemetry.proto.experimental.arrow.v1.ArrowPayload
	0,  // 1: opentelemetry.proto.experimental.arrow.v1.ArrowPayload.type:type_name -> opentelemetry.proto.experimental.arrow.v1.ArrowPayloadType
	6,  // 2: opentelemetry.proto.experimental.arrow.v1.BatchStatus.statuses:type_name -> opentelemetry.proto.experimental.arrow.v1.StatusMessage
	1,  // 3: opentelemetry.proto.experimental.arrow.v1.StatusMessage.status_code:type_name -> opentelemetry.proto.experimental.arrow.v1.StatusCode
	2,  // 4: opentelemetry.proto.experimental.arrow.v1.StatusMessage.error_code:type_name -> opentelemetry.proto.experimental.arrow.v1.ErrorCode
	7,  // 5: opentelemetry.proto.experimental.arrow.v1.StatusMessage.retry_info:type_name -> opentelemetry.proto.experimental.arrow.v1.RetryInfo
	3,  // 6: opentelemetry.proto.experimental.arrow.v1.ArrowStreamService.ArrowStream:input_type -> opentelemetry.proto.experimental.arrow.v1.BatchArrowRecords
	3,  // 7: opentelemetry.proto.experimental.arrow.v1.ArrowTracesService.ArrowTraces:input_type -> opentelemetry.proto.experimental.arrow.v1.BatchArrowRecords
	3,  // 8: opentelemetry.proto.experimental.arrow.v1.ArrowLogsService.ArrowLogs:input_type -> opentelemetry.proto.experimental.arrow.v1.BatchArrowRecords
	3,  // 9: opentelemetry.proto.experimental.arrow.v1.ArrowMetricsService.ArrowMetrics:input_type -> opentelemetry.proto.experimental.arrow.v1.BatchArrowRecords
	5,  // 10: opentelemetry.proto.experimental.arrow.v1.ArrowStreamService.ArrowStream:output_type -> opentelemetry.proto.experimental.arrow.v1.BatchStatus
	5,  // 11: opentelemetry.proto.experimental.arrow.v1.ArrowTracesService.ArrowTraces:output_type -> opentelemetry.proto.experimental.arrow.v1.BatchStatus
	5,  // 12: opentelemetry.proto.experimental.arrow.v1.ArrowLogsService.ArrowLogs:output_type -> opentelemetry.proto.experimental.arrow.v1.BatchStatus
	5,  // 13: opentelemetry.proto.experimental.arrow.v1.ArrowMetricsService.ArrowMetrics:output_type -> opentelemetry.proto.experimental.arrow.v1.BatchStatus
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_init() }
func file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_init() {
	if File_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchArrowRecords); i {
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
		file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrowPayload); i {
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
		file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchStatus); i {
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
		file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusMessage); i {
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
		file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryInfo); i {
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
			RawDescriptor: file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_goTypes,
		DependencyIndexes: file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_depIdxs,
		EnumInfos:         file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_enumTypes,
		MessageInfos:      file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_msgTypes,
	}.Build()
	File_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto = out.File
	file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_rawDesc = nil
	file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_goTypes = nil
	file_opentelemetry_proto_experimental_arrow_v1_arrow_service_proto_depIdxs = nil
}