// Code generated by protoc-gen-go. DO NOT EDIT.
// source: system_metrics.proto

package internal

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SystemMetrics struct {
	Info                 *SystemMetricsInfo         `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Artifact             *SystemMetricsArtifactInfo `protobuf:"bytes,2,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SystemMetrics) Reset()         { *m = SystemMetrics{} }
func (m *SystemMetrics) String() string { return proto.CompactTextString(m) }
func (*SystemMetrics) ProtoMessage()    {}
func (*SystemMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_6937a64b9bbae5c1, []int{0}
}

func (m *SystemMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemMetrics.Unmarshal(m, b)
}
func (m *SystemMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemMetrics.Marshal(b, m, deterministic)
}
func (m *SystemMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemMetrics.Merge(m, src)
}
func (m *SystemMetrics) XXX_Size() int {
	return xxx_messageInfo_SystemMetrics.Size(m)
}
func (m *SystemMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_SystemMetrics proto.InternalMessageInfo

func (m *SystemMetrics) GetInfo() *SystemMetricsInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *SystemMetrics) GetArtifact() *SystemMetricsArtifactInfo {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type SystemMetricsInfo struct {
	Project              string               `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Version              string               `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Variant              string               `protobuf:"bytes,3,opt,name=variant,proto3" json:"variant,omitempty"`
	TaskName             string               `protobuf:"bytes,4,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	TaskId               string               `protobuf:"bytes,5,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Execution            int32                `protobuf:"varint,6,opt,name=execution,proto3" json:"execution,omitempty"`
	Mainline             bool                 `protobuf:"varint,7,opt,name=mainline,proto3" json:"mainline,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SystemMetricsInfo) Reset()         { *m = SystemMetricsInfo{} }
func (m *SystemMetricsInfo) String() string { return proto.CompactTextString(m) }
func (*SystemMetricsInfo) ProtoMessage()    {}
func (*SystemMetricsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6937a64b9bbae5c1, []int{1}
}

func (m *SystemMetricsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemMetricsInfo.Unmarshal(m, b)
}
func (m *SystemMetricsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemMetricsInfo.Marshal(b, m, deterministic)
}
func (m *SystemMetricsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemMetricsInfo.Merge(m, src)
}
func (m *SystemMetricsInfo) XXX_Size() int {
	return xxx_messageInfo_SystemMetricsInfo.Size(m)
}
func (m *SystemMetricsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemMetricsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SystemMetricsInfo proto.InternalMessageInfo

func (m *SystemMetricsInfo) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *SystemMetricsInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SystemMetricsInfo) GetVariant() string {
	if m != nil {
		return m.Variant
	}
	return ""
}

func (m *SystemMetricsInfo) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *SystemMetricsInfo) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SystemMetricsInfo) GetExecution() int32 {
	if m != nil {
		return m.Execution
	}
	return 0
}

func (m *SystemMetricsInfo) GetMainline() bool {
	if m != nil {
		return m.Mainline
	}
	return false
}

func (m *SystemMetricsInfo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type SystemMetricsArtifactInfo struct {
	Format               DataFormat      `protobuf:"varint,1,opt,name=format,proto3,enum=cedar.DataFormat" json:"format,omitempty"`
	Compression          CompressionType `protobuf:"varint,2,opt,name=compression,proto3,enum=cedar.CompressionType" json:"compression,omitempty"`
	Schema               SchemaType      `protobuf:"varint,3,opt,name=schema,proto3,enum=cedar.SchemaType" json:"schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SystemMetricsArtifactInfo) Reset()         { *m = SystemMetricsArtifactInfo{} }
func (m *SystemMetricsArtifactInfo) String() string { return proto.CompactTextString(m) }
func (*SystemMetricsArtifactInfo) ProtoMessage()    {}
func (*SystemMetricsArtifactInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6937a64b9bbae5c1, []int{2}
}

func (m *SystemMetricsArtifactInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemMetricsArtifactInfo.Unmarshal(m, b)
}
func (m *SystemMetricsArtifactInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemMetricsArtifactInfo.Marshal(b, m, deterministic)
}
func (m *SystemMetricsArtifactInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemMetricsArtifactInfo.Merge(m, src)
}
func (m *SystemMetricsArtifactInfo) XXX_Size() int {
	return xxx_messageInfo_SystemMetricsArtifactInfo.Size(m)
}
func (m *SystemMetricsArtifactInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemMetricsArtifactInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SystemMetricsArtifactInfo proto.InternalMessageInfo

func (m *SystemMetricsArtifactInfo) GetFormat() DataFormat {
	if m != nil {
		return m.Format
	}
	return DataFormat_TEXT
}

func (m *SystemMetricsArtifactInfo) GetCompression() CompressionType {
	if m != nil {
		return m.Compression
	}
	return CompressionType_NONE
}

func (m *SystemMetricsArtifactInfo) GetSchema() SchemaType {
	if m != nil {
		return m.Schema
	}
	return SchemaType_RAW_EVENTS
}

type SystemMetricsData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemMetricsData) Reset()         { *m = SystemMetricsData{} }
func (m *SystemMetricsData) String() string { return proto.CompactTextString(m) }
func (*SystemMetricsData) ProtoMessage()    {}
func (*SystemMetricsData) Descriptor() ([]byte, []int) {
	return fileDescriptor_6937a64b9bbae5c1, []int{3}
}

func (m *SystemMetricsData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemMetricsData.Unmarshal(m, b)
}
func (m *SystemMetricsData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemMetricsData.Marshal(b, m, deterministic)
}
func (m *SystemMetricsData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemMetricsData.Merge(m, src)
}
func (m *SystemMetricsData) XXX_Size() int {
	return xxx_messageInfo_SystemMetricsData.Size(m)
}
func (m *SystemMetricsData) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemMetricsData.DiscardUnknown(m)
}

var xxx_messageInfo_SystemMetricsData proto.InternalMessageInfo

func (m *SystemMetricsData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SystemMetricsData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SystemMetricsSeriesEnd struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompletedAt          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SystemMetricsSeriesEnd) Reset()         { *m = SystemMetricsSeriesEnd{} }
func (m *SystemMetricsSeriesEnd) String() string { return proto.CompactTextString(m) }
func (*SystemMetricsSeriesEnd) ProtoMessage()    {}
func (*SystemMetricsSeriesEnd) Descriptor() ([]byte, []int) {
	return fileDescriptor_6937a64b9bbae5c1, []int{4}
}

func (m *SystemMetricsSeriesEnd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemMetricsSeriesEnd.Unmarshal(m, b)
}
func (m *SystemMetricsSeriesEnd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemMetricsSeriesEnd.Marshal(b, m, deterministic)
}
func (m *SystemMetricsSeriesEnd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemMetricsSeriesEnd.Merge(m, src)
}
func (m *SystemMetricsSeriesEnd) XXX_Size() int {
	return xxx_messageInfo_SystemMetricsSeriesEnd.Size(m)
}
func (m *SystemMetricsSeriesEnd) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemMetricsSeriesEnd.DiscardUnknown(m)
}

var xxx_messageInfo_SystemMetricsSeriesEnd proto.InternalMessageInfo

func (m *SystemMetricsSeriesEnd) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SystemMetricsSeriesEnd) GetCompletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CompletedAt
	}
	return nil
}

type SystemMetricsResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemMetricsResponse) Reset()         { *m = SystemMetricsResponse{} }
func (m *SystemMetricsResponse) String() string { return proto.CompactTextString(m) }
func (*SystemMetricsResponse) ProtoMessage()    {}
func (*SystemMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6937a64b9bbae5c1, []int{5}
}

func (m *SystemMetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemMetricsResponse.Unmarshal(m, b)
}
func (m *SystemMetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemMetricsResponse.Marshal(b, m, deterministic)
}
func (m *SystemMetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemMetricsResponse.Merge(m, src)
}
func (m *SystemMetricsResponse) XXX_Size() int {
	return xxx_messageInfo_SystemMetricsResponse.Size(m)
}
func (m *SystemMetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemMetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SystemMetricsResponse proto.InternalMessageInfo

func (m *SystemMetricsResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*SystemMetrics)(nil), "cedar.SystemMetrics")
	proto.RegisterType((*SystemMetricsInfo)(nil), "cedar.SystemMetricsInfo")
	proto.RegisterType((*SystemMetricsArtifactInfo)(nil), "cedar.SystemMetricsArtifactInfo")
	proto.RegisterType((*SystemMetricsData)(nil), "cedar.SystemMetricsData")
	proto.RegisterType((*SystemMetricsSeriesEnd)(nil), "cedar.SystemMetricsSeriesEnd")
	proto.RegisterType((*SystemMetricsResponse)(nil), "cedar.SystemMetricsResponse")
}

func init() { proto.RegisterFile("system_metrics.proto", fileDescriptor_6937a64b9bbae5c1) }

var fileDescriptor_6937a64b9bbae5c1 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x95, 0xd3, 0x36, 0x71, 0x26, 0x69, 0xf4, 0x75, 0xbf, 0x52, 0x4c, 0x28, 0x22, 0xf2, 0x85,
	0x20, 0x21, 0x57, 0x0a, 0x07, 0x40, 0x82, 0x43, 0x08, 0x20, 0x15, 0x44, 0x91, 0x9c, 0x9e, 0xb8,
	0x44, 0x5b, 0xef, 0x24, 0x2c, 0x64, 0x77, 0xad, 0xdd, 0x2d, 0xa2, 0xe2, 0x2f, 0xf0, 0x4f, 0xf8,
	0x2f, 0xfc, 0x26, 0xe4, 0x89, 0xe3, 0x34, 0x6d, 0xa0, 0x17, 0x6e, 0x99, 0x99, 0x37, 0xef, 0x4d,
	0xde, 0xf3, 0xc2, 0xbe, 0xbb, 0x70, 0x1e, 0xd5, 0x44, 0xa1, 0xb7, 0x32, 0x73, 0x49, 0x6e, 0x8d,
	0x37, 0x6c, 0x27, 0x43, 0xc1, 0x6d, 0xf7, 0xfe, 0xcc, 0x98, 0xd9, 0x1c, 0x8f, 0xa8, 0x79, 0x76,
	0x3e, 0x3d, 0xf2, 0x52, 0xa1, 0xf3, 0x5c, 0xe5, 0x0b, 0x5c, 0x77, 0x77, 0x6a, 0xac, 0xe2, 0xbe,
	0x5c, 0x8b, 0xbf, 0xc3, 0xee, 0x98, 0xe8, 0xde, 0x2f, 0xd8, 0xd8, 0x23, 0xd8, 0x96, 0x7a, 0x6a,
	0xa2, 0xa0, 0x17, 0xf4, 0x5b, 0x83, 0x28, 0x21, 0xda, 0x64, 0x0d, 0x73, 0xac, 0xa7, 0x26, 0x25,
	0x14, 0x7b, 0x0e, 0x21, 0xb7, 0x5e, 0x4e, 0x79, 0xe6, 0xa3, 0x1a, 0x6d, 0xf4, 0x36, 0x6d, 0x0c,
	0x4b, 0x0c, 0x6d, 0x56, 0x1b, 0xf1, 0x8f, 0x1a, 0xec, 0x5d, 0x63, 0x66, 0x11, 0x34, 0x72, 0x6b,
	0x3e, 0x63, 0xe6, 0xe9, 0x88, 0x66, 0xba, 0x2c, 0x8b, 0xc9, 0x57, 0xb4, 0x4e, 0x1a, 0x4d, 0x62,
	0xcd, 0x74, 0x59, 0xd2, 0x84, 0x5b, 0xc9, 0xb5, 0x8f, 0xb6, 0xca, 0xc9, 0xa2, 0x64, 0x77, 0xa1,
	0xe9, 0xb9, 0xfb, 0x32, 0xd1, 0x5c, 0x61, 0xb4, 0x4d, 0xb3, 0xb0, 0x68, 0x9c, 0x70, 0x85, 0xec,
	0x36, 0x34, 0x68, 0x28, 0x45, 0xb4, 0x43, 0xa3, 0x7a, 0x51, 0x1e, 0x0b, 0x76, 0x08, 0x4d, 0xfc,
	0x86, 0xd9, 0xb9, 0x2f, 0xb4, 0xea, 0xbd, 0xa0, 0xbf, 0x93, 0xae, 0x1a, 0xac, 0x0b, 0xa1, 0xe2,
	0x52, 0xcf, 0xa5, 0xc6, 0xa8, 0xd1, 0x0b, 0xfa, 0x61, 0x5a, 0xd5, 0xec, 0x19, 0x40, 0x66, 0x91,
	0x7b, 0x14, 0x13, 0xee, 0xa3, 0x90, 0x3c, 0xe9, 0x26, 0x8b, 0x54, 0x92, 0x65, 0x2a, 0xc9, 0xe9,
	0x32, 0x95, 0xb4, 0x59, 0xa2, 0x87, 0x3e, 0xfe, 0x19, 0xc0, 0x9d, 0x3f, 0xda, 0xc6, 0x1e, 0x42,
	0x7d, 0x11, 0x1d, 0xb9, 0xd2, 0x19, 0xec, 0x95, 0x46, 0xbf, 0xe2, 0x9e, 0xbf, 0xa1, 0x41, 0x5a,
	0x02, 0xd8, 0x53, 0x68, 0x65, 0x46, 0xe5, 0x16, 0x5d, 0xe5, 0x55, 0x67, 0x70, 0x50, 0xe2, 0x47,
	0xab, 0xc9, 0xe9, 0x45, 0x8e, 0xe9, 0x65, 0x68, 0x21, 0xe2, 0xb2, 0x4f, 0xa8, 0x38, 0xd9, 0xb8,
	0x12, 0x19, 0x53, 0x93, 0xf0, 0x25, 0x20, 0x7e, 0x72, 0x25, 0xbb, 0xe2, 0x0e, 0xd6, 0x81, 0x9a,
	0x14, 0x65, 0x6c, 0x35, 0x29, 0x18, 0x83, 0x6d, 0xc1, 0x3d, 0xa7, 0x13, 0xda, 0x29, 0xfd, 0x8e,
	0x67, 0x70, 0xb0, 0xb6, 0x38, 0x46, 0x2b, 0xd1, 0xbd, 0xd6, 0xe2, 0xda, 0xf6, 0x0b, 0x68, 0x17,
	0xc7, 0xcd, 0xb1, 0x74, 0x73, 0xeb, 0x46, 0x37, 0x5b, 0x15, 0x7e, 0xe8, 0xe3, 0x07, 0x70, 0x6b,
	0x4d, 0x28, 0x45, 0x97, 0x1b, 0xed, 0xf0, 0xaa, 0xce, 0xe0, 0x57, 0x0d, 0xd8, 0xa8, 0xf8, 0x9f,
	0xeb, 0x4f, 0xe1, 0x04, 0xa2, 0x11, 0x85, 0x73, 0xb9, 0x9d, 0x62, 0x66, 0xac, 0x60, 0xfb, 0x9b,
	0x3e, 0xf3, 0xee, 0xe1, 0xa6, 0x6e, 0x25, 0xfb, 0x16, 0xfe, 0x1b, 0x0a, 0xb1, 0xae, 0xb1, 0xf1,
	0x81, 0x15, 0x56, 0xde, 0xc0, 0xf5, 0x01, 0xfe, 0x1f, 0x7b, 0x8b, 0x5c, 0xfd, 0x13, 0xba, 0x7e,
	0xc0, 0xde, 0x41, 0x7b, 0x34, 0x37, 0x0e, 0x97, 0x4c, 0xf7, 0x36, 0xe1, 0xab, 0xa8, 0xfe, 0x4e,
	0xf7, 0x12, 0x3e, 0x86, 0x52, 0x7b, 0xb4, 0x9a, 0xcf, 0xcf, 0xea, 0x14, 0xd3, 0xe3, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x1d, 0x9e, 0x49, 0x01, 0xb7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CedarSystemMetricsClient is the client API for CedarSystemMetrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CedarSystemMetricsClient interface {
	CreateSystemMetricRecord(ctx context.Context, in *SystemMetrics, opts ...grpc.CallOption) (*SystemMetricsResponse, error)
	AddSystemMetrics(ctx context.Context, in *SystemMetricsData, opts ...grpc.CallOption) (*SystemMetricsResponse, error)
	StreamSystemMetrics(ctx context.Context, opts ...grpc.CallOption) (CedarSystemMetrics_StreamSystemMetricsClient, error)
	CloseMetrics(ctx context.Context, in *SystemMetricsSeriesEnd, opts ...grpc.CallOption) (*SystemMetricsResponse, error)
}

type cedarSystemMetricsClient struct {
	cc *grpc.ClientConn
}

func NewCedarSystemMetricsClient(cc *grpc.ClientConn) CedarSystemMetricsClient {
	return &cedarSystemMetricsClient{cc}
}

func (c *cedarSystemMetricsClient) CreateSystemMetricRecord(ctx context.Context, in *SystemMetrics, opts ...grpc.CallOption) (*SystemMetricsResponse, error) {
	out := new(SystemMetricsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarSystemMetrics/CreateSystemMetricRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cedarSystemMetricsClient) AddSystemMetrics(ctx context.Context, in *SystemMetricsData, opts ...grpc.CallOption) (*SystemMetricsResponse, error) {
	out := new(SystemMetricsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarSystemMetrics/AddSystemMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cedarSystemMetricsClient) StreamSystemMetrics(ctx context.Context, opts ...grpc.CallOption) (CedarSystemMetrics_StreamSystemMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CedarSystemMetrics_serviceDesc.Streams[0], "/cedar.CedarSystemMetrics/StreamSystemMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &cedarSystemMetricsStreamSystemMetricsClient{stream}
	return x, nil
}

type CedarSystemMetrics_StreamSystemMetricsClient interface {
	Send(*SystemMetricsData) error
	CloseAndRecv() (*SystemMetricsResponse, error)
	grpc.ClientStream
}

type cedarSystemMetricsStreamSystemMetricsClient struct {
	grpc.ClientStream
}

func (x *cedarSystemMetricsStreamSystemMetricsClient) Send(m *SystemMetricsData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cedarSystemMetricsStreamSystemMetricsClient) CloseAndRecv() (*SystemMetricsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SystemMetricsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cedarSystemMetricsClient) CloseMetrics(ctx context.Context, in *SystemMetricsSeriesEnd, opts ...grpc.CallOption) (*SystemMetricsResponse, error) {
	out := new(SystemMetricsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarSystemMetrics/CloseMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CedarSystemMetricsServer is the server API for CedarSystemMetrics service.
type CedarSystemMetricsServer interface {
	CreateSystemMetricRecord(context.Context, *SystemMetrics) (*SystemMetricsResponse, error)
	AddSystemMetrics(context.Context, *SystemMetricsData) (*SystemMetricsResponse, error)
	StreamSystemMetrics(CedarSystemMetrics_StreamSystemMetricsServer) error
	CloseMetrics(context.Context, *SystemMetricsSeriesEnd) (*SystemMetricsResponse, error)
}

// UnimplementedCedarSystemMetricsServer can be embedded to have forward compatible implementations.
type UnimplementedCedarSystemMetricsServer struct {
}

func (*UnimplementedCedarSystemMetricsServer) CreateSystemMetricRecord(ctx context.Context, req *SystemMetrics) (*SystemMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSystemMetricRecord not implemented")
}
func (*UnimplementedCedarSystemMetricsServer) AddSystemMetrics(ctx context.Context, req *SystemMetricsData) (*SystemMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSystemMetrics not implemented")
}
func (*UnimplementedCedarSystemMetricsServer) StreamSystemMetrics(srv CedarSystemMetrics_StreamSystemMetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSystemMetrics not implemented")
}
func (*UnimplementedCedarSystemMetricsServer) CloseMetrics(ctx context.Context, req *SystemMetricsSeriesEnd) (*SystemMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseMetrics not implemented")
}

func RegisterCedarSystemMetricsServer(s *grpc.Server, srv CedarSystemMetricsServer) {
	s.RegisterService(&_CedarSystemMetrics_serviceDesc, srv)
}

func _CedarSystemMetrics_CreateSystemMetricRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMetrics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarSystemMetricsServer).CreateSystemMetricRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarSystemMetrics/CreateSystemMetricRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarSystemMetricsServer).CreateSystemMetricRecord(ctx, req.(*SystemMetrics))
	}
	return interceptor(ctx, in, info, handler)
}

func _CedarSystemMetrics_AddSystemMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMetricsData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarSystemMetricsServer).AddSystemMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarSystemMetrics/AddSystemMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarSystemMetricsServer).AddSystemMetrics(ctx, req.(*SystemMetricsData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CedarSystemMetrics_StreamSystemMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CedarSystemMetricsServer).StreamSystemMetrics(&cedarSystemMetricsStreamSystemMetricsServer{stream})
}

type CedarSystemMetrics_StreamSystemMetricsServer interface {
	SendAndClose(*SystemMetricsResponse) error
	Recv() (*SystemMetricsData, error)
	grpc.ServerStream
}

type cedarSystemMetricsStreamSystemMetricsServer struct {
	grpc.ServerStream
}

func (x *cedarSystemMetricsStreamSystemMetricsServer) SendAndClose(m *SystemMetricsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cedarSystemMetricsStreamSystemMetricsServer) Recv() (*SystemMetricsData, error) {
	m := new(SystemMetricsData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CedarSystemMetrics_CloseMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMetricsSeriesEnd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarSystemMetricsServer).CloseMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarSystemMetrics/CloseMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarSystemMetricsServer).CloseMetrics(ctx, req.(*SystemMetricsSeriesEnd))
	}
	return interceptor(ctx, in, info, handler)
}

var _CedarSystemMetrics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cedar.CedarSystemMetrics",
	HandlerType: (*CedarSystemMetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSystemMetricRecord",
			Handler:    _CedarSystemMetrics_CreateSystemMetricRecord_Handler,
		},
		{
			MethodName: "AddSystemMetrics",
			Handler:    _CedarSystemMetrics_AddSystemMetrics_Handler,
		},
		{
			MethodName: "CloseMetrics",
			Handler:    _CedarSystemMetrics_CloseMetrics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSystemMetrics",
			Handler:       _CedarSystemMetrics_StreamSystemMetrics_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "system_metrics.proto",
}