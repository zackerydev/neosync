// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: mgmt/v1alpha1/metrics.proto

package mgmtv1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type RangedMetricName int32

const (
	// If unspecified, an error will be thrown
	RangedMetricName_RANGED_METRIC_NAME_UNSPECIFIED RangedMetricName = 0
	// The input_received metric
	RangedMetricName_RANGED_METRIC_NAME_INPUT_RECEIVED RangedMetricName = 1
)

// Enum value maps for RangedMetricName.
var (
	RangedMetricName_name = map[int32]string{
		0: "RANGED_METRIC_NAME_UNSPECIFIED",
		1: "RANGED_METRIC_NAME_INPUT_RECEIVED",
	}
	RangedMetricName_value = map[string]int32{
		"RANGED_METRIC_NAME_UNSPECIFIED":    0,
		"RANGED_METRIC_NAME_INPUT_RECEIVED": 1,
	}
)

func (x RangedMetricName) Enum() *RangedMetricName {
	p := new(RangedMetricName)
	*p = x
	return p
}

func (x RangedMetricName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RangedMetricName) Descriptor() protoreflect.EnumDescriptor {
	return file_mgmt_v1alpha1_metrics_proto_enumTypes[0].Descriptor()
}

func (RangedMetricName) Type() protoreflect.EnumType {
	return &file_mgmt_v1alpha1_metrics_proto_enumTypes[0]
}

func (x RangedMetricName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RangedMetricName.Descriptor instead.
func (RangedMetricName) EnumDescriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_metrics_proto_rawDescGZIP(), []int{0}
}

type MetricResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// figure out what this is
	Metric map[string]string `protobuf:"bytes,1,rep,name=metric,proto3" json:"metric,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A repeated list of metric values for the given range
	Values []*MetricValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *MetricResult) Reset() {
	*x = MetricResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricResult) ProtoMessage() {}

func (x *MetricResult) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricResult.ProtoReflect.Descriptor instead.
func (*MetricResult) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *MetricResult) GetMetric() map[string]string {
	if x != nil {
		return x.Metric
	}
	return nil
}

func (x *MetricResult) GetValues() []*MetricValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type MetricValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The actual value of the metric
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// The timestamp of when this value was recorded
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *MetricValue) Reset() {
	*x = MetricValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricValue) ProtoMessage() {}

func (x *MetricValue) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricValue.ProtoReflect.Descriptor instead.
func (*MetricValue) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *MetricValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *MetricValue) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type GetRangedMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The start time
	Start *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// The end time
	End *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	// The metric to return
	Metric RangedMetricName `protobuf:"varint,3,opt,name=metric,proto3,enum=mgmt.v1alpha1.RangedMetricName" json:"metric,omitempty"`
	// Types that are assignable to Identifier:
	//
	//	*GetRangedMetricsRequest_AccountId
	//	*GetRangedMetricsRequest_JobId
	//	*GetRangedMetricsRequest_RunId
	Identifier isGetRangedMetricsRequest_Identifier `protobuf_oneof:"identifier"`
}

func (x *GetRangedMetricsRequest) Reset() {
	*x = GetRangedMetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_metrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRangedMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRangedMetricsRequest) ProtoMessage() {}

func (x *GetRangedMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_metrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRangedMetricsRequest.ProtoReflect.Descriptor instead.
func (*GetRangedMetricsRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_metrics_proto_rawDescGZIP(), []int{2}
}

func (x *GetRangedMetricsRequest) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *GetRangedMetricsRequest) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *GetRangedMetricsRequest) GetMetric() RangedMetricName {
	if x != nil {
		return x.Metric
	}
	return RangedMetricName_RANGED_METRIC_NAME_UNSPECIFIED
}

func (m *GetRangedMetricsRequest) GetIdentifier() isGetRangedMetricsRequest_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (x *GetRangedMetricsRequest) GetAccountId() string {
	if x, ok := x.GetIdentifier().(*GetRangedMetricsRequest_AccountId); ok {
		return x.AccountId
	}
	return ""
}

func (x *GetRangedMetricsRequest) GetJobId() string {
	if x, ok := x.GetIdentifier().(*GetRangedMetricsRequest_JobId); ok {
		return x.JobId
	}
	return ""
}

func (x *GetRangedMetricsRequest) GetRunId() string {
	if x, ok := x.GetIdentifier().(*GetRangedMetricsRequest_RunId); ok {
		return x.RunId
	}
	return ""
}

type isGetRangedMetricsRequest_Identifier interface {
	isGetRangedMetricsRequest_Identifier()
}

type GetRangedMetricsRequest_AccountId struct {
	// The account identifier that will be used to filter by
	AccountId string `protobuf:"bytes,4,opt,name=account_id,json=accountId,proto3,oneof"`
}

type GetRangedMetricsRequest_JobId struct {
	// The job identifier that will be used to filter by
	JobId string `protobuf:"bytes,5,opt,name=job_id,json=jobId,proto3,oneof"`
}

type GetRangedMetricsRequest_RunId struct {
	// The run identifier that will be used to filter by
	RunId string `protobuf:"bytes,6,opt,name=run_id,json=runId,proto3,oneof"`
}

func (*GetRangedMetricsRequest_AccountId) isGetRangedMetricsRequest_Identifier() {}

func (*GetRangedMetricsRequest_JobId) isGetRangedMetricsRequest_Identifier() {}

func (*GetRangedMetricsRequest_RunId) isGetRangedMetricsRequest_Identifier() {}

type GetRangedMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*MetricResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GetRangedMetricsResponse) Reset() {
	*x = GetRangedMetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_metrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRangedMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRangedMetricsResponse) ProtoMessage() {}

func (x *GetRangedMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_metrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRangedMetricsResponse.ProtoReflect.Descriptor instead.
func (*GetRangedMetricsResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_metrics_proto_rawDescGZIP(), []int{3}
}

func (x *GetRangedMetricsResponse) GetResults() []*MetricResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type GetMetricCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The start time
	Start *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// The end time
	End *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	// The metric to return
	Metric RangedMetricName `protobuf:"varint,3,opt,name=metric,proto3,enum=mgmt.v1alpha1.RangedMetricName" json:"metric,omitempty"`
	// Types that are assignable to Identifier:
	//
	//	*GetMetricCountRequest_AccountId
	//	*GetMetricCountRequest_JobId
	//	*GetMetricCountRequest_RunId
	Identifier isGetMetricCountRequest_Identifier `protobuf_oneof:"identifier"`
}

func (x *GetMetricCountRequest) Reset() {
	*x = GetMetricCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_metrics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetricCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricCountRequest) ProtoMessage() {}

func (x *GetMetricCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_metrics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricCountRequest.ProtoReflect.Descriptor instead.
func (*GetMetricCountRequest) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_metrics_proto_rawDescGZIP(), []int{4}
}

func (x *GetMetricCountRequest) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *GetMetricCountRequest) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *GetMetricCountRequest) GetMetric() RangedMetricName {
	if x != nil {
		return x.Metric
	}
	return RangedMetricName_RANGED_METRIC_NAME_UNSPECIFIED
}

func (m *GetMetricCountRequest) GetIdentifier() isGetMetricCountRequest_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (x *GetMetricCountRequest) GetAccountId() string {
	if x, ok := x.GetIdentifier().(*GetMetricCountRequest_AccountId); ok {
		return x.AccountId
	}
	return ""
}

func (x *GetMetricCountRequest) GetJobId() string {
	if x, ok := x.GetIdentifier().(*GetMetricCountRequest_JobId); ok {
		return x.JobId
	}
	return ""
}

func (x *GetMetricCountRequest) GetRunId() string {
	if x, ok := x.GetIdentifier().(*GetMetricCountRequest_RunId); ok {
		return x.RunId
	}
	return ""
}

type isGetMetricCountRequest_Identifier interface {
	isGetMetricCountRequest_Identifier()
}

type GetMetricCountRequest_AccountId struct {
	// The account identifier that will be used to filter by
	AccountId string `protobuf:"bytes,4,opt,name=account_id,json=accountId,proto3,oneof"`
}

type GetMetricCountRequest_JobId struct {
	// The job identifier that will be used to filter by
	JobId string `protobuf:"bytes,5,opt,name=job_id,json=jobId,proto3,oneof"`
}

type GetMetricCountRequest_RunId struct {
	// The run identifier that will be used to filter by
	RunId string `protobuf:"bytes,6,opt,name=run_id,json=runId,proto3,oneof"`
}

func (*GetMetricCountRequest_AccountId) isGetMetricCountRequest_Identifier() {}

func (*GetMetricCountRequest_JobId) isGetMetricCountRequest_Identifier() {}

func (*GetMetricCountRequest_RunId) isGetMetricCountRequest_Identifier() {}

type GetMetricCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The summed up count of the metric based on the input query and timerange specified
	Count uint64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetMetricCountResponse) Reset() {
	*x = GetMetricCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mgmt_v1alpha1_metrics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetricCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricCountResponse) ProtoMessage() {}

func (x *GetMetricCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mgmt_v1alpha1_metrics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricCountResponse.ProtoReflect.Descriptor instead.
func (*GetMetricCountResponse) Descriptor() ([]byte, []int) {
	return file_mgmt_v1alpha1_metrics_proto_rawDescGZIP(), []int{5}
}

func (x *GetMetricCountResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_mgmt_v1alpha1_metrics_proto protoreflect.FileDescriptor

var file_mgmt_v1alpha1_metrics_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x0c, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x67,
	0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x32, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5d, 0x0a, 0x0b, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xb1, 0x02, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12,
	0x29, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x48, 0x00, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x06, 0x6a, 0x6f,
	0x62, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba,
	0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64,
	0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x51,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x67,
	0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x22, 0xaf, 0x02, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a,
	0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x67,
	0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x12, 0x29, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0,
	0x01, 0x01, 0x48, 0x00, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x05, 0x6a, 0x6f, 0x62,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x48, 0x00, 0x52, 0x05,
	0x72, 0x75, 0x6e, 0x49, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2a, 0x5d, 0x0a, 0x10, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x41, 0x4e, 0x47, 0x45,
	0x44, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x44, 0x5f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d,
	0x45, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44,
	0x10, 0x01, 0x32, 0xd8, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x26, 0x2e, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24,
	0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xc8, 0x01,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x75, 0x63, 0x6c, 0x65, 0x75, 0x73, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6e, 0x65, 0x6f,
	0x73, 0x79, 0x6e, 0x63, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6d, 0x67, 0x6d, 0x74, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4d, 0x67,
	0x6d, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x0d, 0x4d, 0x67,
	0x6d, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x19, 0x4d, 0x67,
	0x6d, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4d, 0x67, 0x6d, 0x74, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mgmt_v1alpha1_metrics_proto_rawDescOnce sync.Once
	file_mgmt_v1alpha1_metrics_proto_rawDescData = file_mgmt_v1alpha1_metrics_proto_rawDesc
)

func file_mgmt_v1alpha1_metrics_proto_rawDescGZIP() []byte {
	file_mgmt_v1alpha1_metrics_proto_rawDescOnce.Do(func() {
		file_mgmt_v1alpha1_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_mgmt_v1alpha1_metrics_proto_rawDescData)
	})
	return file_mgmt_v1alpha1_metrics_proto_rawDescData
}

var file_mgmt_v1alpha1_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mgmt_v1alpha1_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mgmt_v1alpha1_metrics_proto_goTypes = []interface{}{
	(RangedMetricName)(0),            // 0: mgmt.v1alpha1.RangedMetricName
	(*MetricResult)(nil),             // 1: mgmt.v1alpha1.MetricResult
	(*MetricValue)(nil),              // 2: mgmt.v1alpha1.MetricValue
	(*GetRangedMetricsRequest)(nil),  // 3: mgmt.v1alpha1.GetRangedMetricsRequest
	(*GetRangedMetricsResponse)(nil), // 4: mgmt.v1alpha1.GetRangedMetricsResponse
	(*GetMetricCountRequest)(nil),    // 5: mgmt.v1alpha1.GetMetricCountRequest
	(*GetMetricCountResponse)(nil),   // 6: mgmt.v1alpha1.GetMetricCountResponse
	nil,                              // 7: mgmt.v1alpha1.MetricResult.MetricEntry
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
}
var file_mgmt_v1alpha1_metrics_proto_depIdxs = []int32{
	7,  // 0: mgmt.v1alpha1.MetricResult.metric:type_name -> mgmt.v1alpha1.MetricResult.MetricEntry
	2,  // 1: mgmt.v1alpha1.MetricResult.values:type_name -> mgmt.v1alpha1.MetricValue
	8,  // 2: mgmt.v1alpha1.MetricValue.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 3: mgmt.v1alpha1.GetRangedMetricsRequest.start:type_name -> google.protobuf.Timestamp
	8,  // 4: mgmt.v1alpha1.GetRangedMetricsRequest.end:type_name -> google.protobuf.Timestamp
	0,  // 5: mgmt.v1alpha1.GetRangedMetricsRequest.metric:type_name -> mgmt.v1alpha1.RangedMetricName
	1,  // 6: mgmt.v1alpha1.GetRangedMetricsResponse.results:type_name -> mgmt.v1alpha1.MetricResult
	8,  // 7: mgmt.v1alpha1.GetMetricCountRequest.start:type_name -> google.protobuf.Timestamp
	8,  // 8: mgmt.v1alpha1.GetMetricCountRequest.end:type_name -> google.protobuf.Timestamp
	0,  // 9: mgmt.v1alpha1.GetMetricCountRequest.metric:type_name -> mgmt.v1alpha1.RangedMetricName
	3,  // 10: mgmt.v1alpha1.MetricsService.GetRangedMetrics:input_type -> mgmt.v1alpha1.GetRangedMetricsRequest
	5,  // 11: mgmt.v1alpha1.MetricsService.GetMetricCount:input_type -> mgmt.v1alpha1.GetMetricCountRequest
	4,  // 12: mgmt.v1alpha1.MetricsService.GetRangedMetrics:output_type -> mgmt.v1alpha1.GetRangedMetricsResponse
	6,  // 13: mgmt.v1alpha1.MetricsService.GetMetricCount:output_type -> mgmt.v1alpha1.GetMetricCountResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_mgmt_v1alpha1_metrics_proto_init() }
func file_mgmt_v1alpha1_metrics_proto_init() {
	if File_mgmt_v1alpha1_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mgmt_v1alpha1_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricResult); i {
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
		file_mgmt_v1alpha1_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricValue); i {
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
		file_mgmt_v1alpha1_metrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRangedMetricsRequest); i {
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
		file_mgmt_v1alpha1_metrics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRangedMetricsResponse); i {
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
		file_mgmt_v1alpha1_metrics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetricCountRequest); i {
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
		file_mgmt_v1alpha1_metrics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetricCountResponse); i {
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
	file_mgmt_v1alpha1_metrics_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*GetRangedMetricsRequest_AccountId)(nil),
		(*GetRangedMetricsRequest_JobId)(nil),
		(*GetRangedMetricsRequest_RunId)(nil),
	}
	file_mgmt_v1alpha1_metrics_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*GetMetricCountRequest_AccountId)(nil),
		(*GetMetricCountRequest_JobId)(nil),
		(*GetMetricCountRequest_RunId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mgmt_v1alpha1_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mgmt_v1alpha1_metrics_proto_goTypes,
		DependencyIndexes: file_mgmt_v1alpha1_metrics_proto_depIdxs,
		EnumInfos:         file_mgmt_v1alpha1_metrics_proto_enumTypes,
		MessageInfos:      file_mgmt_v1alpha1_metrics_proto_msgTypes,
	}.Build()
	File_mgmt_v1alpha1_metrics_proto = out.File
	file_mgmt_v1alpha1_metrics_proto_rawDesc = nil
	file_mgmt_v1alpha1_metrics_proto_goTypes = nil
	file_mgmt_v1alpha1_metrics_proto_depIdxs = nil
}
