// Copyright (c) 2019, Juniper Networks, Inc.
// All rights reserved.

// This IDL defines objects that specify lightweight encapsulation and
// decapsulation parameters for flexible tunnels. This type of tunnel
// is intended to be controlled via JET applications for efficient,
// high-scale dynamic tunneling. The tunnels created by these objects
// do not have control plane functionality, i.e., protocols or other
// control plane features cannot run over these interfaces, but they are
// sent to the forwarding plane where forwarding state is instantiated for
// them.
//
// In the context of this API, an "attribute" is a single parameter of
// the tunnel encapsulation/decapsulation (e.g. source port), while
// a "profile" the complete set of attributes that fully defines all
// required parameters needed to encapsulate or decapsulate traffic.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: jnx_routing_flexible_tunnel_profile.proto

// This is part of the routing package.

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// [brief]: The action performed for the tunnel.
// [detail]: The action performed for the tunnel.
// [default]: ENCAPSULATE.
type TunnelActionType int32

const (
	// [brief]: Use associated TunnelAttributes to encapsulate packets.
	TunnelActionType_ENCAPSULATE TunnelActionType = 0
	// [brief]: Use associated TunnelAttributes to decapsulate packets.
	TunnelActionType_DECAPSULATE TunnelActionType = 1
)

// Enum value maps for TunnelActionType.
var (
	TunnelActionType_name = map[int32]string{
		0: "ENCAPSULATE",
		1: "DECAPSULATE",
	}
	TunnelActionType_value = map[string]int32{
		"ENCAPSULATE": 0,
		"DECAPSULATE": 1,
	}
)

func (x TunnelActionType) Enum() *TunnelActionType {
	p := new(TunnelActionType)
	*p = x
	return p
}

func (x TunnelActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TunnelActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_jnx_routing_flexible_tunnel_profile_proto_enumTypes[0].Descriptor()
}

func (TunnelActionType) Type() protoreflect.EnumType {
	return &file_jnx_routing_flexible_tunnel_profile_proto_enumTypes[0]
}

func (x TunnelActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TunnelActionType.Descriptor instead.
func (TunnelActionType) EnumDescriptor() ([]byte, []int) {
	return file_jnx_routing_flexible_tunnel_profile_proto_rawDescGZIP(), []int{0}
}

// [brief]: VXLAN tunnel attributes.
// [detail]: A set of VXLAN flexible tunnel attributes contains the encapsulation and/or
// decapsulation parameters specific to a VXLAN flexible tunnel.
//
// A VXLAN decapsulation tunnel is considered unique only if the full tuple
// of FTI interface, VNI, source prefix/length, and destination port is
// unique from other decapsulation tunnels. Adding or updating VXLAN
// decapsulation tunnel with conflicting attributes will result in an error.
type VxlanTunnelAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [mandatory]: required for both encapsulation and decapsulation.
	// [brief]: 24 bit VXLAN Network Identifier (VNI).
	// [range]: [1:16777214].
	Vni uint32 `protobuf:"varint,1,opt,name=vni,proto3" json:"vni,omitempty"`
	// [mandatory]: required for both encapsulation and decapsulation.
	// [brief]: Source address prefix.
	// [detail]: Outer source address prefix for the encapsulated traffic.
	// Must be IPv4 or IPv6.
	SourcePrefix *NetworkAddress `protobuf:"bytes,2,opt,name=source_prefix,json=sourcePrefix,proto3" json:"source_prefix,omitempty"`
	// [mandatory]: required for both encapsulation and decapsulation.
	// [brief]: Prefix length of source address.
	// [detail]: Source address prefix length.
	// For encapsulation, it must be the host
	// address prefix length.
	// prefix length will be 32 for IPv4, 128 for IPv6.
	SourcePrefixLen uint32 `protobuf:"varint,3,opt,name=source_prefix_len,json=sourcePrefixLen,proto3" json:"source_prefix_len,omitempty"`
	// [brief]: Source UDP port range.
	// [detail]: Outer source UDP port range. The source port for each route will be
	// picked from the given range.
	// REQUIRED for encapsulations, OPTIONAL and ignored for decapsulations.
	// [range]: [0:65535].
	SourceUdpPortRange *NumericRange `protobuf:"bytes,4,opt,name=source_udp_port_range,json=sourceUdpPortRange,proto3" json:"source_udp_port_range,omitempty"`
	// [brief]: Source MAC address.
	// [detail]: Source MAC address will be ignored. OPTIONAL. JUNOS will set the
	// source MAC by default for encapsulations.
	// [default]: 00:00:5e:00:52:01.
	SourceMac *MacAddress `protobuf:"bytes,5,opt,name=source_mac,json=sourceMac,proto3" json:"source_mac,omitempty"`
	// [brief]: Destination address
	// [detail]: Outer destination address for the encapsulated traffic.
	// Must be the IPv4 or IPv6 address of a local interface.
	// REQUIRED for encapsulations, OPTIONAL and ignored for decapsulations.
	DestinationAddress *NetworkAddress `protobuf:"bytes,6,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	// [mandatory]: required for both encapsulation and decapsulation.
	// [brief]: Destination UDP port.
	// [Range]: [0:65535]
	// (inclusive).
	DestinationUdpPort uint32 `protobuf:"varint,7,opt,name=destination_udp_port,json=destinationUdpPort,proto3" json:"destination_udp_port,omitempty"`
	// [brief]: Destination MAC address.
	// [detail]: Destination MAC address, a 48-bit MAC value. OPTIONAL.
	// If not specified, JUNOS will set the destination MAC by
	// default for encapsulations.
	// [default_value]: 00:00:5e:00:52:00.
	DestinationMac *MacAddress `protobuf:"bytes,8,opt,name=destination_mac,json=destinationMac,proto3" json:"destination_mac,omitempty"`
	// [brief]: VXLAN flag.
	// [detail]: VXLAN Flags are 8 bits of flags as specified in RFC7348.
	// [default_value]: The default RFC7348 value of 0x08 will
	// be used in encapsulations.
	Flags *wrapperspb.UInt32Value `protobuf:"bytes,9,opt,name=flags,proto3" json:"flags,omitempty"`
	// [brief]: VXLAN header fields reserved as of RFC7348.
	// [detail]: VXLAN header fields reserved as defined in RFC7348.
	// [default]: Zero.
	//
	// Types that are assignable to VxlanHeaderReserved:
	//	*VxlanTunnelAttributes_Rfc7348Reserved
	VxlanHeaderReserved isVxlanTunnelAttributes_VxlanHeaderReserved `protobuf_oneof:"vxlan_header_reserved"`
}

func (x *VxlanTunnelAttributes) Reset() {
	*x = VxlanTunnelAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VxlanTunnelAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VxlanTunnelAttributes) ProtoMessage() {}

func (x *VxlanTunnelAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VxlanTunnelAttributes.ProtoReflect.Descriptor instead.
func (*VxlanTunnelAttributes) Descriptor() ([]byte, []int) {
	return file_jnx_routing_flexible_tunnel_profile_proto_rawDescGZIP(), []int{0}
}

func (x *VxlanTunnelAttributes) GetVni() uint32 {
	if x != nil {
		return x.Vni
	}
	return 0
}

func (x *VxlanTunnelAttributes) GetSourcePrefix() *NetworkAddress {
	if x != nil {
		return x.SourcePrefix
	}
	return nil
}

func (x *VxlanTunnelAttributes) GetSourcePrefixLen() uint32 {
	if x != nil {
		return x.SourcePrefixLen
	}
	return 0
}

func (x *VxlanTunnelAttributes) GetSourceUdpPortRange() *NumericRange {
	if x != nil {
		return x.SourceUdpPortRange
	}
	return nil
}

func (x *VxlanTunnelAttributes) GetSourceMac() *MacAddress {
	if x != nil {
		return x.SourceMac
	}
	return nil
}

func (x *VxlanTunnelAttributes) GetDestinationAddress() *NetworkAddress {
	if x != nil {
		return x.DestinationAddress
	}
	return nil
}

func (x *VxlanTunnelAttributes) GetDestinationUdpPort() uint32 {
	if x != nil {
		return x.DestinationUdpPort
	}
	return 0
}

func (x *VxlanTunnelAttributes) GetDestinationMac() *MacAddress {
	if x != nil {
		return x.DestinationMac
	}
	return nil
}

func (x *VxlanTunnelAttributes) GetFlags() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (m *VxlanTunnelAttributes) GetVxlanHeaderReserved() isVxlanTunnelAttributes_VxlanHeaderReserved {
	if m != nil {
		return m.VxlanHeaderReserved
	}
	return nil
}

func (x *VxlanTunnelAttributes) GetRfc7348Reserved() *Rfc7348VxlanReserved {
	if x, ok := x.GetVxlanHeaderReserved().(*VxlanTunnelAttributes_Rfc7348Reserved); ok {
		return x.Rfc7348Reserved
	}
	return nil
}

type isVxlanTunnelAttributes_VxlanHeaderReserved interface {
	isVxlanTunnelAttributes_VxlanHeaderReserved()
}

type VxlanTunnelAttributes_Rfc7348Reserved struct {
	// [brief]: VXLAN header reserved fields as defined in RFC7348.
	// [detail]: VXLAN header fields reserved as defined in RFC7348.
	// [default]: Zero.
	Rfc7348Reserved *Rfc7348VxlanReserved `protobuf:"bytes,10,opt,name=rfc7348_reserved,json=rfc7348Reserved,proto3,oneof"`
}

func (*VxlanTunnelAttributes_Rfc7348Reserved) isVxlanTunnelAttributes_VxlanHeaderReserved() {}

// [brief]: VXLAN header reserved fields as of RFC7348.
// [detail]: VXLAN header reserved fields as defined in RFC7348.
type Rfc7348VxlanReserved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [brief]: VXLAN 24-bit reserved field as specified in RFC7348.
	// [detail]: This is the RFC7348 VXLAN header 24-bit reserved field that
	// follows the VXLAN flags and preceeds the VNI (bits 8-31).
	// This field is ignored for decapulation profiles.
	// If not set, the bits will be set to all zeros.
	// If set, only the lower 24-bits may be non-zero.
	// [range]: [0:(2^24-1)]
	ReservedBits_8_31 uint32 `protobuf:"varint,1,opt,name=reserved_bits_8_31,json=reservedBits831,proto3" json:"reserved_bits_8_31,omitempty"`
	// [brief]: VXLAN 8-bit reserved field as specified in RFC7348.
	// [detail]: This is the RFC7348 VXLAN header 8-bit reserved field that
	// follows the VNI (bits 55-63).
	// This field is ignored for decapulation profiles.
	// [range]: [0:255]
	ReservedBits_55_63 uint32 `protobuf:"varint,2,opt,name=reserved_bits_55_63,json=reservedBits5563,proto3" json:"reserved_bits_55_63,omitempty"`
}

func (x *Rfc7348VxlanReserved) Reset() {
	*x = Rfc7348VxlanReserved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rfc7348VxlanReserved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rfc7348VxlanReserved) ProtoMessage() {}

func (x *Rfc7348VxlanReserved) ProtoReflect() protoreflect.Message {
	mi := &file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rfc7348VxlanReserved.ProtoReflect.Descriptor instead.
func (*Rfc7348VxlanReserved) Descriptor() ([]byte, []int) {
	return file_jnx_routing_flexible_tunnel_profile_proto_rawDescGZIP(), []int{1}
}

func (x *Rfc7348VxlanReserved) GetReservedBits_8_31() uint32 {
	if x != nil {
		return x.ReservedBits_8_31
	}
	return 0
}

func (x *Rfc7348VxlanReserved) GetReservedBits_55_63() uint32 {
	if x != nil {
		return x.ReservedBits_55_63
	}
	return 0
}

// [brief]: Tunnel attributes.
// [detail]: Encapsulation/Decapsulation attributes for various types of tunnels.
type TunnelAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [brief]: Tunnel types supported by flexible tunnels.
	//
	// Types that are assignable to TunnelAttributesType:
	//	*TunnelAttributes_Vxlan
	TunnelAttributesType isTunnelAttributes_TunnelAttributesType `protobuf_oneof:"tunnel_attributes_type"`
}

func (x *TunnelAttributes) Reset() {
	*x = TunnelAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TunnelAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelAttributes) ProtoMessage() {}

func (x *TunnelAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelAttributes.ProtoReflect.Descriptor instead.
func (*TunnelAttributes) Descriptor() ([]byte, []int) {
	return file_jnx_routing_flexible_tunnel_profile_proto_rawDescGZIP(), []int{2}
}

func (m *TunnelAttributes) GetTunnelAttributesType() isTunnelAttributes_TunnelAttributesType {
	if m != nil {
		return m.TunnelAttributesType
	}
	return nil
}

func (x *TunnelAttributes) GetVxlan() *VxlanTunnelAttributes {
	if x, ok := x.GetTunnelAttributesType().(*TunnelAttributes_Vxlan); ok {
		return x.Vxlan
	}
	return nil
}

type isTunnelAttributes_TunnelAttributesType interface {
	isTunnelAttributes_TunnelAttributesType()
}

type TunnelAttributes_Vxlan struct {
	// [brief]: VXLAN tunnels.
	Vxlan *VxlanTunnelAttributes `protobuf:"bytes,1,opt,name=vxlan,proto3,oneof"`
}

func (*TunnelAttributes_Vxlan) isTunnelAttributes_TunnelAttributesType() {}

// [brief]: Flexible tunnel profile.
// [detail]: A flexible tunnel profile contains a set encapsulation parameters for a
// flexible tunnel. It may specify an encapsulation or decapsulation.
type FlexibleTunnelProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [brief]: Flexible tunnel profile name.
	// [detail]: A unique client-assigned name for this profile.
	// REQUIRED when when used with any RPC from the flexible_tunnel_service.
	// OPTIONAL when FlexibleTunnelProfile is used in RouteGateway.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// [mandatory]:
	// [brief]: The action that this profile will be used for.
	// [detail]: The action that this profile will be used for: encapsulate,
	// decapsulate.
	Action TunnelActionType `protobuf:"varint,2,opt,name=action,proto3,enum=jnx.jet.routing.rib.TunnelActionType" json:"action,omitempty"`
	// statistics, policing, filtering, and default encapsulation parameters.
	// [mandatory]:
	// [brief]: Name of the anchoring tunnel interface.
	// [detail]: Name of the anchoring tunnel interface to use for
	InterfaceName string `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	// [mandatory]:
	// [brief]: The encapsulation or decapsulation parameters defining the profile.
	Attributes *TunnelAttributes `protobuf:"bytes,4,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *FlexibleTunnelProfile) Reset() {
	*x = FlexibleTunnelProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlexibleTunnelProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlexibleTunnelProfile) ProtoMessage() {}

func (x *FlexibleTunnelProfile) ProtoReflect() protoreflect.Message {
	mi := &file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlexibleTunnelProfile.ProtoReflect.Descriptor instead.
func (*FlexibleTunnelProfile) Descriptor() ([]byte, []int) {
	return file_jnx_routing_flexible_tunnel_profile_proto_rawDescGZIP(), []int{3}
}

func (x *FlexibleTunnelProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlexibleTunnelProfile) GetAction() TunnelActionType {
	if x != nil {
		return x.Action
	}
	return TunnelActionType_ENCAPSULATE
}

func (x *FlexibleTunnelProfile) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *FlexibleTunnelProfile) GetAttributes() *TunnelAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

var File_jnx_routing_flexible_tunnel_profile_proto protoreflect.FileDescriptor

var file_jnx_routing_flexible_tunnel_profile_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6a, 0x6e, 0x78, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x6c,
	0x65, 0x78, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6a, 0x6e, 0x78,
	0x2e, 0x6a, 0x65, 0x74, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x69, 0x62,
	0x1a, 0x1b, 0x6a, 0x6e, 0x78, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6a,
	0x6e, 0x78, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6a, 0x6e, 0x78, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x05, 0x0a, 0x15, 0x56, 0x78, 0x6c,
	0x61, 0x6e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x6e, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x76, 0x6e, 0x69, 0x12, 0x49, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6a, 0x6e,
	0x78, 0x2e, 0x6a, 0x65, 0x74, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x2a, 0x0a, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x65, 0x6e, 0x12, 0x4f, 0x0a, 0x15, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6e, 0x78,
	0x2e, 0x6a, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x75, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x55, 0x64, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6a, 0x6e, 0x78, 0x2e, 0x6a, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4d, 0x61, 0x63, 0x12, 0x55, 0x0a, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6a, 0x6e, 0x78, 0x2e, 0x6a, 0x65, 0x74, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x12, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x30,
	0x0a, 0x14, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x64,
	0x70, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x64, 0x70, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x43, 0x0a, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x61, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6a, 0x6e, 0x78, 0x2e,
	0x6a, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x63, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x61, 0x63, 0x12, 0x32, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x56, 0x0a, 0x10, 0x72, 0x66, 0x63,
	0x37, 0x33, 0x34, 0x38, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6a, 0x6e, 0x78, 0x2e, 0x6a, 0x65, 0x74, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x69, 0x62, 0x2e, 0x52, 0x66, 0x63, 0x37, 0x33, 0x34,
	0x38, 0x56, 0x78, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x0f, 0x72, 0x66, 0x63, 0x37, 0x33, 0x34, 0x38, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x42, 0x17, 0x0a, 0x15, 0x76, 0x78, 0x6c, 0x61, 0x6e, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x22, 0x72, 0x0a, 0x14, 0x52, 0x66,
	0x63, 0x37, 0x33, 0x34, 0x38, 0x56, 0x78, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x64, 0x12, 0x2b, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x62,
	0x69, 0x74, 0x73, 0x5f, 0x38, 0x5f, 0x33, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x42, 0x69, 0x74, 0x73, 0x38, 0x33, 0x31, 0x12,
	0x2d, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x69, 0x74, 0x73,
	0x5f, 0x35, 0x35, 0x5f, 0x36, 0x33, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x42, 0x69, 0x74, 0x73, 0x35, 0x35, 0x36, 0x33, 0x22, 0x70,
	0x0a, 0x10, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x42, 0x0a, 0x05, 0x76, 0x78, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x6a, 0x6e, 0x78, 0x2e, 0x6a, 0x65, 0x74, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x72, 0x69, 0x62, 0x2e, 0x56, 0x78, 0x6c, 0x61, 0x6e, 0x54, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x48, 0x00, 0x52,
	0x05, 0x76, 0x78, 0x6c, 0x61, 0x6e, 0x42, 0x18, 0x0a, 0x16, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0xd8, 0x01, 0x0a, 0x15, 0x46, 0x6c, 0x65, 0x78, 0x69, 0x62, 0x6c, 0x65, 0x54, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25,
	0x2e, 0x6a, 0x6e, 0x78, 0x2e, 0x6a, 0x65, 0x74, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x72, 0x69, 0x62, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6a, 0x6e, 0x78, 0x2e, 0x6a,
	0x65, 0x74, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x69, 0x62, 0x2e, 0x54,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2a, 0x34, 0x0a, 0x10, 0x54,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x43, 0x41, 0x50, 0x53, 0x55, 0x4c, 0x41, 0x54, 0x45, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x45, 0x43, 0x41, 0x50, 0x53, 0x55, 0x4c, 0x41, 0x54, 0x45, 0x10,
	0x01, 0x42, 0x24, 0x5a, 0x11, 0x6c, 0x61, 0x62, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x82, 0xb5, 0x18, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x8a,
	0xb5, 0x18, 0x04, 0x31, 0x39, 0x2e, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jnx_routing_flexible_tunnel_profile_proto_rawDescOnce sync.Once
	file_jnx_routing_flexible_tunnel_profile_proto_rawDescData = file_jnx_routing_flexible_tunnel_profile_proto_rawDesc
)

func file_jnx_routing_flexible_tunnel_profile_proto_rawDescGZIP() []byte {
	file_jnx_routing_flexible_tunnel_profile_proto_rawDescOnce.Do(func() {
		file_jnx_routing_flexible_tunnel_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_jnx_routing_flexible_tunnel_profile_proto_rawDescData)
	})
	return file_jnx_routing_flexible_tunnel_profile_proto_rawDescData
}

var file_jnx_routing_flexible_tunnel_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_jnx_routing_flexible_tunnel_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_jnx_routing_flexible_tunnel_profile_proto_goTypes = []interface{}{
	(TunnelActionType)(0),          // 0: jnx.jet.routing.rib.TunnelActionType
	(*VxlanTunnelAttributes)(nil),  // 1: jnx.jet.routing.rib.VxlanTunnelAttributes
	(*Rfc7348VxlanReserved)(nil),   // 2: jnx.jet.routing.rib.Rfc7348VxlanReserved
	(*TunnelAttributes)(nil),       // 3: jnx.jet.routing.rib.TunnelAttributes
	(*FlexibleTunnelProfile)(nil),  // 4: jnx.jet.routing.rib.FlexibleTunnelProfile
	(*NetworkAddress)(nil),         // 5: jnx.jet.routing.base.NetworkAddress
	(*NumericRange)(nil),           // 6: jnx.jet.common.NumericRange
	(*MacAddress)(nil),             // 7: jnx.jet.common.MacAddress
	(*wrapperspb.UInt32Value)(nil), // 8: google.protobuf.UInt32Value
}
var file_jnx_routing_flexible_tunnel_profile_proto_depIdxs = []int32{
	5,  // 0: jnx.jet.routing.rib.VxlanTunnelAttributes.source_prefix:type_name -> jnx.jet.routing.base.NetworkAddress
	6,  // 1: jnx.jet.routing.rib.VxlanTunnelAttributes.source_udp_port_range:type_name -> jnx.jet.common.NumericRange
	7,  // 2: jnx.jet.routing.rib.VxlanTunnelAttributes.source_mac:type_name -> jnx.jet.common.MacAddress
	5,  // 3: jnx.jet.routing.rib.VxlanTunnelAttributes.destination_address:type_name -> jnx.jet.routing.base.NetworkAddress
	7,  // 4: jnx.jet.routing.rib.VxlanTunnelAttributes.destination_mac:type_name -> jnx.jet.common.MacAddress
	8,  // 5: jnx.jet.routing.rib.VxlanTunnelAttributes.flags:type_name -> google.protobuf.UInt32Value
	2,  // 6: jnx.jet.routing.rib.VxlanTunnelAttributes.rfc7348_reserved:type_name -> jnx.jet.routing.rib.Rfc7348VxlanReserved
	1,  // 7: jnx.jet.routing.rib.TunnelAttributes.vxlan:type_name -> jnx.jet.routing.rib.VxlanTunnelAttributes
	0,  // 8: jnx.jet.routing.rib.FlexibleTunnelProfile.action:type_name -> jnx.jet.routing.rib.TunnelActionType
	3,  // 9: jnx.jet.routing.rib.FlexibleTunnelProfile.attributes:type_name -> jnx.jet.routing.rib.TunnelAttributes
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_jnx_routing_flexible_tunnel_profile_proto_init() }
func file_jnx_routing_flexible_tunnel_profile_proto_init() {
	if File_jnx_routing_flexible_tunnel_profile_proto != nil {
		return
	}
	file_jnx_common_addr_types_proto_init()
	file_jnx_common_base_types_proto_init()
	file_jnx_routing_base_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VxlanTunnelAttributes); i {
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
		file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rfc7348VxlanReserved); i {
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
		file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TunnelAttributes); i {
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
		file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlexibleTunnelProfile); i {
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
	file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*VxlanTunnelAttributes_Rfc7348Reserved)(nil),
	}
	file_jnx_routing_flexible_tunnel_profile_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TunnelAttributes_Vxlan)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_jnx_routing_flexible_tunnel_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jnx_routing_flexible_tunnel_profile_proto_goTypes,
		DependencyIndexes: file_jnx_routing_flexible_tunnel_profile_proto_depIdxs,
		EnumInfos:         file_jnx_routing_flexible_tunnel_profile_proto_enumTypes,
		MessageInfos:      file_jnx_routing_flexible_tunnel_profile_proto_msgTypes,
	}.Build()
	File_jnx_routing_flexible_tunnel_profile_proto = out.File
	file_jnx_routing_flexible_tunnel_profile_proto_rawDesc = nil
	file_jnx_routing_flexible_tunnel_profile_proto_goTypes = nil
	file_jnx_routing_flexible_tunnel_profile_proto_depIdxs = nil
}
