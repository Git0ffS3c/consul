// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/udp/udp_proxy/v2alpha/udp_proxy.proto

package envoy_config_filter_udp_udp_proxy_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type UdpProxyConfig struct {
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are valid to be assigned to RouteSpecifier:
	//	*UdpProxyConfig_Cluster
	RouteSpecifier       isUdpProxyConfig_RouteSpecifier `protobuf_oneof:"route_specifier"`
	IdleTimeout          *duration.Duration              `protobuf:"bytes,3,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *UdpProxyConfig) Reset()         { *m = UdpProxyConfig{} }
func (m *UdpProxyConfig) String() string { return proto.CompactTextString(m) }
func (*UdpProxyConfig) ProtoMessage()    {}
func (*UdpProxyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_82ee653a514c2977, []int{0}
}

func (m *UdpProxyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UdpProxyConfig.Unmarshal(m, b)
}
func (m *UdpProxyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UdpProxyConfig.Marshal(b, m, deterministic)
}
func (m *UdpProxyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpProxyConfig.Merge(m, src)
}
func (m *UdpProxyConfig) XXX_Size() int {
	return xxx_messageInfo_UdpProxyConfig.Size(m)
}
func (m *UdpProxyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpProxyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UdpProxyConfig proto.InternalMessageInfo

func (m *UdpProxyConfig) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

type isUdpProxyConfig_RouteSpecifier interface {
	isUdpProxyConfig_RouteSpecifier()
}

type UdpProxyConfig_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*UdpProxyConfig_Cluster) isUdpProxyConfig_RouteSpecifier() {}

func (m *UdpProxyConfig) GetRouteSpecifier() isUdpProxyConfig_RouteSpecifier {
	if m != nil {
		return m.RouteSpecifier
	}
	return nil
}

func (m *UdpProxyConfig) GetCluster() string {
	if x, ok := m.GetRouteSpecifier().(*UdpProxyConfig_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *UdpProxyConfig) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UdpProxyConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UdpProxyConfig_Cluster)(nil),
	}
}

func init() {
	proto.RegisterType((*UdpProxyConfig)(nil), "envoy.config.filter.udp.udp_proxy.v2alpha.UdpProxyConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/udp/udp_proxy/v2alpha/udp_proxy.proto", fileDescriptor_82ee653a514c2977)
}

var fileDescriptor_82ee653a514c2977 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3f, 0x4b, 0x73, 0x31,
	0x18, 0xc5, 0xdf, 0xb4, 0x2f, 0x2d, 0xa6, 0xfe, 0xe3, 0x0e, 0x5a, 0x0b, 0x4a, 0xd1, 0xa5, 0x2e,
	0x09, 0xd4, 0x41, 0x04, 0xa7, 0xe8, 0xe0, 0x78, 0xb9, 0xe8, 0x7c, 0x49, 0x9b, 0xdc, 0x1a, 0x88,
	0xf7, 0x09, 0xb9, 0x49, 0x69, 0x37, 0xbf, 0x94, 0x8b, 0x9f, 0xc0, 0xd5, 0x6f, 0x23, 0x4e, 0x92,
	0xa4, 0x97, 0xe2, 0xe6, 0x10, 0x08, 0xcf, 0x39, 0xbf, 0xe4, 0x39, 0x07, 0xdf, 0xc8, 0x7a, 0x09,
	0x6b, 0x3a, 0x87, 0xba, 0x52, 0x0b, 0x5a, 0x29, 0xed, 0xa4, 0xa5, 0x5e, 0x98, 0x70, 0x4a, 0x63,
	0x61, 0xb5, 0xa6, 0xcb, 0x29, 0xd7, 0xe6, 0x99, 0x6f, 0x27, 0xc4, 0x58, 0x70, 0x90, 0x5d, 0x46,
	0x94, 0x24, 0x94, 0x24, 0x94, 0x78, 0x61, 0xc8, 0xd6, 0xb8, 0x41, 0x47, 0x67, 0x0b, 0x80, 0x85,
	0x96, 0x34, 0x82, 0x33, 0x5f, 0x51, 0xe1, 0x2d, 0x77, 0x0a, 0xea, 0xf4, 0xd4, 0xe8, 0xd4, 0x0b,
	0xc3, 0x29, 0xaf, 0x6b, 0x70, 0x71, 0xdc, 0xd0, 0xc6, 0x71, 0xe7, 0x9b, 0x8d, 0x7c, 0xbc, 0xe4,
	0x5a, 0x09, 0xee, 0x24, 0x6d, 0x2f, 0x49, 0x38, 0x7f, 0x43, 0x78, 0xff, 0x49, 0x98, 0x3c, 0x7c,
	0x76, 0x17, 0xf7, 0xc8, 0x26, 0x78, 0x10, 0xd8, 0xd2, 0x58, 0x59, 0xa9, 0xd5, 0x10, 0x8d, 0xd1,
	0x64, 0x87, 0xf5, 0xbf, 0xd9, 0x7f, 0xdb, 0x19, 0xa3, 0x02, 0x07, 0x2d, 0x8f, 0x52, 0x76, 0x81,
	0xfb, 0x73, 0xed, 0x1b, 0x27, 0xed, 0xb0, 0xf3, 0xcb, 0xf5, 0xf0, 0xaf, 0x68, 0x95, 0xec, 0x16,
	0xef, 0x2a, 0xa1, 0x65, 0xe9, 0xd4, 0x8b, 0x04, 0xef, 0x86, 0xdd, 0x31, 0x9a, 0x0c, 0xa6, 0x27,
	0x24, 0x05, 0x22, 0x6d, 0x20, 0x72, 0xbf, 0x09, 0x54, 0x0c, 0x82, 0xfd, 0x31, 0xb9, 0xd9, 0x11,
	0x3e, 0xb0, 0xe0, 0x9d, 0x2c, 0x1b, 0x23, 0xe7, 0xaa, 0x52, 0xd2, 0x66, 0xdd, 0x2f, 0x86, 0x58,
	0xf1, 0xfe, 0xfa, 0xf1, 0xd9, 0xeb, 0x1c, 0x22, 0x7c, 0xad, 0x80, 0xc4, 0x1e, 0x53, 0x61, 0x7f,
	0xae, 0x94, 0xed, 0xb5, 0xb9, 0xf3, 0xb0, 0x42, 0x8e, 0x66, 0xbd, 0xb8, 0xcb, 0xd5, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x2f, 0xa9, 0x74, 0x8b, 0xd2, 0x01, 0x00, 0x00,
}