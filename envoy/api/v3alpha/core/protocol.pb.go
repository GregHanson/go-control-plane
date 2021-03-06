// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/core/protocol.proto

package envoy_api_v3alpha_core

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type TcpProtocolOptions struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpProtocolOptions) Reset()         { *m = TcpProtocolOptions{} }
func (m *TcpProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*TcpProtocolOptions) ProtoMessage()    {}
func (*TcpProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_52a634d5e642c216, []int{0}
}

func (m *TcpProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProtocolOptions.Unmarshal(m, b)
}
func (m *TcpProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProtocolOptions.Marshal(b, m, deterministic)
}
func (m *TcpProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProtocolOptions.Merge(m, src)
}
func (m *TcpProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_TcpProtocolOptions.Size(m)
}
func (m *TcpProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProtocolOptions proto.InternalMessageInfo

type HttpProtocolOptions struct {
	IdleTimeout          *duration.Duration `protobuf:"bytes,1,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HttpProtocolOptions) Reset()         { *m = HttpProtocolOptions{} }
func (m *HttpProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*HttpProtocolOptions) ProtoMessage()    {}
func (*HttpProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_52a634d5e642c216, []int{1}
}

func (m *HttpProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpProtocolOptions.Unmarshal(m, b)
}
func (m *HttpProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpProtocolOptions.Marshal(b, m, deterministic)
}
func (m *HttpProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpProtocolOptions.Merge(m, src)
}
func (m *HttpProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_HttpProtocolOptions.Size(m)
}
func (m *HttpProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_HttpProtocolOptions proto.InternalMessageInfo

func (m *HttpProtocolOptions) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

type Http1ProtocolOptions struct {
	AllowAbsoluteUrl      *wrappers.BoolValue `protobuf:"bytes,1,opt,name=allow_absolute_url,json=allowAbsoluteUrl,proto3" json:"allow_absolute_url,omitempty"`
	AcceptHttp_10         bool                `protobuf:"varint,2,opt,name=accept_http_10,json=acceptHttp10,proto3" json:"accept_http_10,omitempty"`
	DefaultHostForHttp_10 string              `protobuf:"bytes,3,opt,name=default_host_for_http_10,json=defaultHostForHttp10,proto3" json:"default_host_for_http_10,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}            `json:"-"`
	XXX_unrecognized      []byte              `json:"-"`
	XXX_sizecache         int32               `json:"-"`
}

func (m *Http1ProtocolOptions) Reset()         { *m = Http1ProtocolOptions{} }
func (m *Http1ProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*Http1ProtocolOptions) ProtoMessage()    {}
func (*Http1ProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_52a634d5e642c216, []int{2}
}

func (m *Http1ProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http1ProtocolOptions.Unmarshal(m, b)
}
func (m *Http1ProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http1ProtocolOptions.Marshal(b, m, deterministic)
}
func (m *Http1ProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http1ProtocolOptions.Merge(m, src)
}
func (m *Http1ProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_Http1ProtocolOptions.Size(m)
}
func (m *Http1ProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Http1ProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Http1ProtocolOptions proto.InternalMessageInfo

func (m *Http1ProtocolOptions) GetAllowAbsoluteUrl() *wrappers.BoolValue {
	if m != nil {
		return m.AllowAbsoluteUrl
	}
	return nil
}

func (m *Http1ProtocolOptions) GetAcceptHttp_10() bool {
	if m != nil {
		return m.AcceptHttp_10
	}
	return false
}

func (m *Http1ProtocolOptions) GetDefaultHostForHttp_10() string {
	if m != nil {
		return m.DefaultHostForHttp_10
	}
	return ""
}

type Http2ProtocolOptions struct {
	HpackTableSize                               *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=hpack_table_size,json=hpackTableSize,proto3" json:"hpack_table_size,omitempty"`
	MaxConcurrentStreams                         *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_concurrent_streams,json=maxConcurrentStreams,proto3" json:"max_concurrent_streams,omitempty"`
	InitialStreamWindowSize                      *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=initial_stream_window_size,json=initialStreamWindowSize,proto3" json:"initial_stream_window_size,omitempty"`
	InitialConnectionWindowSize                  *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=initial_connection_window_size,json=initialConnectionWindowSize,proto3" json:"initial_connection_window_size,omitempty"`
	AllowConnect                                 bool                  `protobuf:"varint,5,opt,name=allow_connect,json=allowConnect,proto3" json:"allow_connect,omitempty"`
	AllowMetadata                                bool                  `protobuf:"varint,6,opt,name=allow_metadata,json=allowMetadata,proto3" json:"allow_metadata,omitempty"`
	MaxOutboundFrames                            *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=max_outbound_frames,json=maxOutboundFrames,proto3" json:"max_outbound_frames,omitempty"`
	MaxOutboundControlFrames                     *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=max_outbound_control_frames,json=maxOutboundControlFrames,proto3" json:"max_outbound_control_frames,omitempty"`
	MaxConsecutiveInboundFramesWithEmptyPayload  *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=max_consecutive_inbound_frames_with_empty_payload,json=maxConsecutiveInboundFramesWithEmptyPayload,proto3" json:"max_consecutive_inbound_frames_with_empty_payload,omitempty"`
	MaxInboundPriorityFramesPerStream            *wrappers.UInt32Value `protobuf:"bytes,10,opt,name=max_inbound_priority_frames_per_stream,json=maxInboundPriorityFramesPerStream,proto3" json:"max_inbound_priority_frames_per_stream,omitempty"`
	MaxInboundWindowUpdateFramesPerDataFrameSent *wrappers.UInt32Value `protobuf:"bytes,11,opt,name=max_inbound_window_update_frames_per_data_frame_sent,json=maxInboundWindowUpdateFramesPerDataFrameSent,proto3" json:"max_inbound_window_update_frames_per_data_frame_sent,omitempty"`
	StreamErrorOnInvalidHttpMessaging            bool                  `protobuf:"varint,12,opt,name=stream_error_on_invalid_http_messaging,json=streamErrorOnInvalidHttpMessaging,proto3" json:"stream_error_on_invalid_http_messaging,omitempty"`
	XXX_NoUnkeyedLiteral                         struct{}              `json:"-"`
	XXX_unrecognized                             []byte                `json:"-"`
	XXX_sizecache                                int32                 `json:"-"`
}

func (m *Http2ProtocolOptions) Reset()         { *m = Http2ProtocolOptions{} }
func (m *Http2ProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*Http2ProtocolOptions) ProtoMessage()    {}
func (*Http2ProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_52a634d5e642c216, []int{3}
}

func (m *Http2ProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http2ProtocolOptions.Unmarshal(m, b)
}
func (m *Http2ProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http2ProtocolOptions.Marshal(b, m, deterministic)
}
func (m *Http2ProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http2ProtocolOptions.Merge(m, src)
}
func (m *Http2ProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_Http2ProtocolOptions.Size(m)
}
func (m *Http2ProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Http2ProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Http2ProtocolOptions proto.InternalMessageInfo

func (m *Http2ProtocolOptions) GetHpackTableSize() *wrappers.UInt32Value {
	if m != nil {
		return m.HpackTableSize
	}
	return nil
}

func (m *Http2ProtocolOptions) GetMaxConcurrentStreams() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConcurrentStreams
	}
	return nil
}

func (m *Http2ProtocolOptions) GetInitialStreamWindowSize() *wrappers.UInt32Value {
	if m != nil {
		return m.InitialStreamWindowSize
	}
	return nil
}

func (m *Http2ProtocolOptions) GetInitialConnectionWindowSize() *wrappers.UInt32Value {
	if m != nil {
		return m.InitialConnectionWindowSize
	}
	return nil
}

func (m *Http2ProtocolOptions) GetAllowConnect() bool {
	if m != nil {
		return m.AllowConnect
	}
	return false
}

func (m *Http2ProtocolOptions) GetAllowMetadata() bool {
	if m != nil {
		return m.AllowMetadata
	}
	return false
}

func (m *Http2ProtocolOptions) GetMaxOutboundFrames() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxOutboundFrames
	}
	return nil
}

func (m *Http2ProtocolOptions) GetMaxOutboundControlFrames() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxOutboundControlFrames
	}
	return nil
}

func (m *Http2ProtocolOptions) GetMaxConsecutiveInboundFramesWithEmptyPayload() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConsecutiveInboundFramesWithEmptyPayload
	}
	return nil
}

func (m *Http2ProtocolOptions) GetMaxInboundPriorityFramesPerStream() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxInboundPriorityFramesPerStream
	}
	return nil
}

func (m *Http2ProtocolOptions) GetMaxInboundWindowUpdateFramesPerDataFrameSent() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxInboundWindowUpdateFramesPerDataFrameSent
	}
	return nil
}

func (m *Http2ProtocolOptions) GetStreamErrorOnInvalidHttpMessaging() bool {
	if m != nil {
		return m.StreamErrorOnInvalidHttpMessaging
	}
	return false
}

type GrpcProtocolOptions struct {
	Http2ProtocolOptions *Http2ProtocolOptions `protobuf:"bytes,1,opt,name=http2_protocol_options,json=http2ProtocolOptions,proto3" json:"http2_protocol_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GrpcProtocolOptions) Reset()         { *m = GrpcProtocolOptions{} }
func (m *GrpcProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcProtocolOptions) ProtoMessage()    {}
func (*GrpcProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_52a634d5e642c216, []int{4}
}

func (m *GrpcProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcProtocolOptions.Unmarshal(m, b)
}
func (m *GrpcProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcProtocolOptions.Marshal(b, m, deterministic)
}
func (m *GrpcProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcProtocolOptions.Merge(m, src)
}
func (m *GrpcProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_GrpcProtocolOptions.Size(m)
}
func (m *GrpcProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcProtocolOptions proto.InternalMessageInfo

func (m *GrpcProtocolOptions) GetHttp2ProtocolOptions() *Http2ProtocolOptions {
	if m != nil {
		return m.Http2ProtocolOptions
	}
	return nil
}

func init() {
	proto.RegisterType((*TcpProtocolOptions)(nil), "envoy.api.v3alpha.core.TcpProtocolOptions")
	proto.RegisterType((*HttpProtocolOptions)(nil), "envoy.api.v3alpha.core.HttpProtocolOptions")
	proto.RegisterType((*Http1ProtocolOptions)(nil), "envoy.api.v3alpha.core.Http1ProtocolOptions")
	proto.RegisterType((*Http2ProtocolOptions)(nil), "envoy.api.v3alpha.core.Http2ProtocolOptions")
	proto.RegisterType((*GrpcProtocolOptions)(nil), "envoy.api.v3alpha.core.GrpcProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/core/protocol.proto", fileDescriptor_52a634d5e642c216)
}

var fileDescriptor_52a634d5e642c216 = []byte{
	// 797 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcf, 0x53, 0xdb, 0x46,
	0x14, 0xc7, 0x47, 0x21, 0x0d, 0xb0, 0x18, 0x9a, 0x0a, 0x0f, 0x51, 0x49, 0x27, 0x43, 0xdc, 0x24,
	0xe3, 0xa1, 0x19, 0x39, 0x98, 0x4e, 0x7b, 0xe9, 0xa5, 0x26, 0xa1, 0x70, 0x60, 0x70, 0x65, 0x5c,
	0x8e, 0x3b, 0x6b, 0x69, 0x6d, 0x6f, 0xbb, 0xda, 0xb7, 0xb3, 0x5a, 0xf9, 0x07, 0xe7, 0x5e, 0x7a,
	0xec, 0x3f, 0xd4, 0x3f, 0x8c, 0x8b, 0x3b, 0xda, 0x5d, 0xb9, 0xc6, 0xd0, 0xa9, 0xe9, 0xc9, 0xd2,
	0xbe, 0xef, 0xfb, 0x7c, 0xdf, 0x93, 0x9e, 0x9f, 0xd0, 0x5b, 0x2a, 0x46, 0x30, 0x6d, 0x10, 0xc9,
	0x1a, 0xa3, 0x63, 0xc2, 0xe5, 0x90, 0x34, 0x62, 0x50, 0xb4, 0x21, 0x15, 0x68, 0x88, 0x81, 0x87,
	0xe6, 0xc2, 0xdf, 0x33, 0xb2, 0x90, 0x48, 0x16, 0x3a, 0x59, 0x58, 0xc8, 0xf6, 0x5f, 0x0d, 0x00,
	0x06, 0xdc, 0xc9, 0x7b, 0x79, 0xbf, 0x91, 0xe4, 0x8a, 0x68, 0x06, 0xc2, 0xe6, 0xdd, 0x8f, 0x8f,
	0x15, 0x91, 0x92, 0xaa, 0xcc, 0xc5, 0x5f, 0x8c, 0x08, 0x67, 0x09, 0xd1, 0xb4, 0x51, 0x5e, 0xd8,
	0x40, 0xad, 0x8a, 0xfc, 0xab, 0x58, 0xb6, 0x5d, 0x15, 0x97, 0xb2, 0x60, 0x66, 0xb5, 0x0e, 0xda,
	0x3d, 0xd3, 0x7a, 0xf9, 0xd8, 0xff, 0x01, 0x55, 0x58, 0xc2, 0x29, 0xd6, 0x2c, 0xa5, 0x90, 0xeb,
	0xc0, 0x3b, 0xf0, 0xea, 0x5b, 0xcd, 0x2f, 0x43, 0x6b, 0x1e, 0x96, 0xe6, 0xe1, 0x47, 0x57, 0x5c,
	0xb4, 0x55, 0xc8, 0xaf, 0xac, 0xba, 0xf6, 0x97, 0x87, 0xaa, 0x05, 0xf5, 0x68, 0x19, 0x7b, 0x86,
	0x7c, 0xc2, 0x39, 0x8c, 0x31, 0xe9, 0x65, 0xc0, 0x73, 0x4d, 0x71, 0xae, 0xb8, 0x83, 0xef, 0xdf,
	0x83, 0xb7, 0x00, 0xf8, 0x2f, 0x84, 0xe7, 0x34, 0x7a, 0x6e, 0xb2, 0x7e, 0x74, 0x49, 0x5d, 0xc5,
	0xfd, 0x37, 0x68, 0x87, 0xc4, 0x31, 0x95, 0x1a, 0x0f, 0xb5, 0x96, 0xf8, 0xe8, 0x43, 0xf0, 0xe4,
	0xc0, 0xab, 0x6f, 0x44, 0x15, 0x7b, 0x6a, 0xdc, 0x3f, 0xf8, 0xdf, 0xa1, 0x20, 0xa1, 0x7d, 0x92,
	0x73, 0x8d, 0x87, 0x90, 0x69, 0xdc, 0x07, 0x35, 0xd7, 0xaf, 0x1d, 0x78, 0xf5, 0xcd, 0xa8, 0xea,
	0xe2, 0x67, 0x90, 0xe9, 0x53, 0x50, 0x36, 0xaf, 0xf6, 0xc7, 0xa6, 0x6d, 0xa0, 0xb9, 0xdc, 0xc0,
	0x29, 0x7a, 0x3e, 0x94, 0x24, 0xfe, 0x0d, 0x6b, 0xd2, 0xe3, 0x14, 0x67, 0xec, 0x86, 0xba, 0xf2,
	0xbf, 0xba, 0x57, 0x7e, 0xf7, 0x5c, 0xe8, 0xe3, 0xa6, 0x6d, 0x60, 0xc7, 0x64, 0x5d, 0x15, 0x49,
	0x1d, 0x76, 0x43, 0x7d, 0x82, 0xf6, 0x52, 0x32, 0xc1, 0x31, 0x88, 0x38, 0x57, 0x8a, 0x0a, 0x8d,
	0x33, 0xad, 0x28, 0x49, 0x33, 0xd3, 0xc6, 0x7f, 0xd0, 0x5a, 0xdb, 0xb7, 0x2d, 0x74, 0xb8, 0x11,
	0xcc, 0x66, 0xb3, 0xd9, 0x7a, 0xdd, 0x8b, 0xaa, 0x29, 0x99, 0x9c, 0xcc, 0x49, 0x1d, 0x0b, 0xf2,
	0x7f, 0x45, 0xfb, 0x4c, 0x30, 0xcd, 0x08, 0x77, 0x6c, 0x3c, 0x66, 0x22, 0x81, 0xb1, 0x2d, 0x7a,
	0x6d, 0x05, 0x9b, 0xcf, 0x6f, 0x5b, 0x95, 0x43, 0xe4, 0x6c, 0x66, 0xb3, 0xb5, 0xe8, 0x85, 0x03,
	0x5a, 0x8b, 0x6b, 0x83, 0x33, 0xed, 0x28, 0xf4, 0xaa, 0xf4, 0x8a, 0x41, 0x08, 0x1a, 0x17, 0x4f,
	0xeb, 0x8e, 0xdf, 0xd3, 0xff, 0xe3, 0xf7, 0xd2, 0x41, 0x4f, 0xe6, 0xcc, 0x05, 0xcf, 0xaf, 0xd1,
	0xb6, 0x9d, 0x25, 0xe7, 0x18, 0x7c, 0xe6, 0x06, 0xa0, 0x38, 0x74, 0x19, 0xfe, 0x5b, 0xb4, 0x63,
	0x45, 0x29, 0xd5, 0x24, 0x21, 0x9a, 0x04, 0xcf, 0x8c, 0xca, 0xa6, 0x5e, 0xb8, 0x43, 0xbf, 0x8b,
	0x76, 0x8b, 0xd7, 0x01, 0xb9, 0xee, 0x41, 0x2e, 0x12, 0xdc, 0x57, 0x24, 0xa5, 0x59, 0xb0, 0xbe,
	0x42, 0xd1, 0xeb, 0xb7, 0xad, 0xa7, 0x87, 0x4f, 0xea, 0x5e, 0xf4, 0x45, 0x4a, 0x26, 0x97, 0x0e,
	0x70, 0x6a, 0xf2, 0xfd, 0x04, 0xbd, 0xbc, 0x83, 0x8d, 0x41, 0x68, 0x05, 0xbc, 0xc4, 0x6f, 0x3c,
	0x06, 0x1f, 0x2c, 0xe0, 0x4f, 0x2c, 0xc7, 0xb9, 0xfc, 0xee, 0xa1, 0x23, 0x37, 0x4c, 0x19, 0x8d,
	0x73, 0xcd, 0x46, 0x14, 0x33, 0xb1, 0xd8, 0x08, 0x1e, 0x33, 0x3d, 0xc4, 0x34, 0x95, 0x7a, 0x8a,
	0x25, 0x99, 0x72, 0x20, 0x49, 0xb0, 0xb9, 0xc2, 0xd4, 0x7e, 0x63, 0x07, 0xab, 0xa4, 0x9e, 0x8b,
	0x85, 0xe6, 0xae, 0x99, 0x1e, 0x7e, 0x2a, 0x88, 0x6d, 0x0b, 0xf4, 0x01, 0xbd, 0x2b, 0xaa, 0x28,
	0x9d, 0xa5, 0x62, 0xa0, 0x98, 0x9e, 0x96, 0x25, 0x48, 0xaa, 0xdc, 0x1c, 0x06, 0x68, 0x05, 0xeb,
	0xd7, 0x29, 0x99, 0x38, 0xbf, 0xb6, 0x23, 0x59, 0xdf, 0x36, 0x55, 0x76, 0xfc, 0xfc, 0x3f, 0x3d,
	0xf4, 0xed, 0xa2, 0xa3, 0x1b, 0xb7, 0x5c, 0x16, 0x4b, 0x6f, 0xd1, 0xb6, 0x78, 0xc5, 0xf6, 0x1e,
	0x67, 0x54, 0xe8, 0x60, 0xeb, 0x31, 0xcf, 0xfd, 0xfd, 0x3f, 0x85, 0xd8, 0xe1, 0xeb, 0x1a, 0xfe,
	0xbc, 0x98, 0x8f, 0x44, 0x13, 0x73, 0xd3, 0xa1, 0x42, 0xfb, 0x3f, 0xa3, 0x77, 0xee, 0xcf, 0x46,
	0x95, 0x02, 0x85, 0x41, 0x60, 0x26, 0xcc, 0x22, 0xb6, 0x7b, 0x27, 0xa5, 0x59, 0x46, 0x06, 0x4c,
	0x0c, 0x82, 0x8a, 0x99, 0xc3, 0xd7, 0x56, 0xfd, 0xa9, 0x10, 0x5f, 0x8a, 0x73, 0x2b, 0x2d, 0x36,
	0xcf, 0x45, 0x29, 0xac, 0x4d, 0xd1, 0xee, 0x4f, 0x4a, 0xc6, 0xcb, 0x9b, 0xa8, 0x87, 0xf6, 0x0a,
	0x62, 0x13, 0x97, 0xdf, 0x15, 0x0c, 0x36, 0xe2, 0xf6, 0xd1, 0xfb, 0xf0, 0xe1, 0x0f, 0x4c, 0xf8,
	0xd0, 0x5e, 0x8b, 0xaa, 0xc3, 0x07, 0x4e, 0x5b, 0xdf, 0xa3, 0x37, 0x0c, 0x2c, 0x47, 0x2a, 0x98,
	0x4c, 0xff, 0x05, 0xd9, 0xda, 0x2e, 0x13, 0xcd, 0x6f, 0xdb, 0xeb, 0x3d, 0x33, 0x45, 0x1d, 0xff,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xb9, 0xf1, 0xa3, 0x0c, 0x07, 0x00, 0x00,
}
