// Code generated by protoc-gen-go.
// source: common/chain-config.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/chain-config.proto
	common/common.proto
	common/configuration.proto

It has these top-level messages:
	MSPPrincipal
	OrganizationUnit
	MSPRole
	Header
	ChainHeader
	SignatureHeader
	Payload
	Envelope
	Block
	BlockHeader
	BlockData
	BlockMetadata
	ConfigurationEnvelope
	SignedConfigurationItem
	ConfigurationItem
	ConfigurationSignature
	Policy
	SignaturePolicyEnvelope
	SignaturePolicy
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MSPPrincipal_Classification int32

const (
	MSPPrincipal_ByMSPRole MSPPrincipal_Classification = 0
	// one of a member of MSP network, and the one of an
	// administrator of an MSP network
	MSPPrincipal_ByOrganizationUnit MSPPrincipal_Classification = 1
	// groupping of entities, per MSP affiliation
	// E.g., this can well be represented by an MSP's
	// Organization unit
	MSPPrincipal_ByIdentity MSPPrincipal_Classification = 2
)

var MSPPrincipal_Classification_name = map[int32]string{
	0: "ByMSPRole",
	1: "ByOrganizationUnit",
	2: "ByIdentity",
}
var MSPPrincipal_Classification_value = map[string]int32{
	"ByMSPRole":          0,
	"ByOrganizationUnit": 1,
	"ByIdentity":         2,
}

func (x MSPPrincipal_Classification) String() string {
	return proto.EnumName(MSPPrincipal_Classification_name, int32(x))
}
func (MSPPrincipal_Classification) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type MSPRole_MSPRoleType int32

const (
	MSPRole_Member MSPRole_MSPRoleType = 0
	MSPRole_Admin  MSPRole_MSPRoleType = 1
)

var MSPRole_MSPRoleType_name = map[int32]string{
	0: "Member",
	1: "Admin",
}
var MSPRole_MSPRoleType_value = map[string]int32{
	"Member": 0,
	"Admin":  1,
}

func (x MSPRole_MSPRoleType) String() string {
	return proto.EnumName(MSPRole_MSPRoleType_name, int32(x))
}
func (MSPRole_MSPRoleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

// MSPPrincipal aims to represent an MSP-centric set of identities.
// In particular, this structure allows for definition of
//  - a group of identities that are member of the same MSP
//  - a group of identities that are member of the same organization unit
//    in the same MSP
//  - a group of identities that are administering a specific MSP
//  - a specific identity
// Expressing these groups is done given two fields of the fields below
//  - Classification, that defines the type of classification of identities
//    in an MSP this principal would be defined on; Classification can take
//    three values:
//     (i)  ByMSPRole: that represents a classification of identities within
//          MSP based on one of the two pre-defined MSP rules, "member" and "admin"
//     (ii) ByOrganizationUnit: that represents a classification of identities
//          within MSP based on the organization unit an identity belongs to
//     (iii)ByIdentity that denotes that MSPPrincipal is mapped to a single
//          identity/certificate; this would mean that the Principal bytes
//          message
type MSPPrincipal struct {
	// Classification describes the way that one should process
	// Principal. An Classification value of "ByOrganizationUnit" reflects
	// that "Principal" contains the name of an organization this MSP
	// handles. A Classification value "ByIdentity" means that
	// "Principal" contains a specific identity. Default value
	// denotes that Principal contains one of the groups by
	// default supported by all MSPs ("admin" or "member").
	PrincipalClassification MSPPrincipal_Classification `protobuf:"varint,1,opt,name=PrincipalClassification,json=principalClassification,enum=common.MSPPrincipal_Classification" json:"PrincipalClassification,omitempty"`
	// Principal completes the policy principal definition. For the default
	// principal types, Principal can be either "Admin" or "Member".
	// For the ByOrganizationUnit/ByIdentity values of Classification,
	// PolicyPrincipal acquires its value from an organization unit or
	// identity, respectively.
	Principal []byte `protobuf:"bytes,3,opt,name=Principal,json=principal,proto3" json:"Principal,omitempty"`
}

func (m *MSPPrincipal) Reset()                    { *m = MSPPrincipal{} }
func (m *MSPPrincipal) String() string            { return proto.CompactTextString(m) }
func (*MSPPrincipal) ProtoMessage()               {}
func (*MSPPrincipal) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MSPPrincipal) GetPrincipalClassification() MSPPrincipal_Classification {
	if m != nil {
		return m.PrincipalClassification
	}
	return MSPPrincipal_ByMSPRole
}

func (m *MSPPrincipal) GetPrincipal() []byte {
	if m != nil {
		return m.Principal
	}
	return nil
}

// OrganizationUnit governs the organization of the Principal
// field of a policy principal when a specific organization unity members
// are to be defined within a policy principal.
type OrganizationUnit struct {
	// MSPIdentifier represents the identifier of the MSP this organization unit
	// refers to
	MSPIdentifier string `protobuf:"bytes,1,opt,name=MSPIdentifier,json=mSPIdentifier" json:"MSPIdentifier,omitempty"`
	// OrganizationUnitIdentifier defines the organization unit under the
	// MSP identified with MSPIdentifier
	OrganizationUnitIdentifier string `protobuf:"bytes,2,opt,name=OrganizationUnitIdentifier,json=organizationUnitIdentifier" json:"OrganizationUnitIdentifier,omitempty"`
}

func (m *OrganizationUnit) Reset()                    { *m = OrganizationUnit{} }
func (m *OrganizationUnit) String() string            { return proto.CompactTextString(m) }
func (*OrganizationUnit) ProtoMessage()               {}
func (*OrganizationUnit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OrganizationUnit) GetMSPIdentifier() string {
	if m != nil {
		return m.MSPIdentifier
	}
	return ""
}

func (m *OrganizationUnit) GetOrganizationUnitIdentifier() string {
	if m != nil {
		return m.OrganizationUnitIdentifier
	}
	return ""
}

// MSPRole governs the organization of the Principal
// field of an MSPPrincipal when it aims to define one of the
// two dedicated roles within an MSP: Admin and Members.
type MSPRole struct {
	// MSPIdentifier represents the identifier of the MSP this principal
	// refers to
	MSPIdentifier string `protobuf:"bytes,1,opt,name=MSPIdentifier,json=mSPIdentifier" json:"MSPIdentifier,omitempty"`
	// MSPRoleType defines which of the available, pre-defined MSP-roles
	// an identiy should posess inside the MSP with identifier MSPidentifier
	Role MSPRole_MSPRoleType `protobuf:"varint,2,opt,name=Role,json=role,enum=common.MSPRole_MSPRoleType" json:"Role,omitempty"`
}

func (m *MSPRole) Reset()                    { *m = MSPRole{} }
func (m *MSPRole) String() string            { return proto.CompactTextString(m) }
func (*MSPRole) ProtoMessage()               {}
func (*MSPRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MSPRole) GetMSPIdentifier() string {
	if m != nil {
		return m.MSPIdentifier
	}
	return ""
}

func (m *MSPRole) GetRole() MSPRole_MSPRoleType {
	if m != nil {
		return m.Role
	}
	return MSPRole_Member
}

func init() {
	proto.RegisterType((*MSPPrincipal)(nil), "common.MSPPrincipal")
	proto.RegisterType((*OrganizationUnit)(nil), "common.OrganizationUnit")
	proto.RegisterType((*MSPRole)(nil), "common.MSPRole")
	proto.RegisterEnum("common.MSPPrincipal_Classification", MSPPrincipal_Classification_name, MSPPrincipal_Classification_value)
	proto.RegisterEnum("common.MSPRole_MSPRoleType", MSPRole_MSPRoleType_name, MSPRole_MSPRoleType_value)
}

func init() { proto.RegisterFile("common/chain-config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xff, 0xdd, 0x7f, 0x4e, 0xfa, 0xba, 0x95, 0x92, 0x83, 0x9b, 0xd3, 0x83, 0xd4, 0x1d,
	0x04, 0x59, 0x0b, 0xf3, 0x2e, 0x58, 0x0f, 0xe2, 0x61, 0x38, 0xaa, 0x5e, 0x04, 0x0f, 0x6d, 0x96,
	0x76, 0x2f, 0xb4, 0x49, 0xc9, 0x22, 0x58, 0x3f, 0x80, 0x9f, 0xd0, 0x0f, 0x64, 0x96, 0xea, 0xa8,
	0x45, 0xc1, 0x53, 0xf2, 0xbe, 0xf9, 0x3d, 0xcf, 0x93, 0xbc, 0x04, 0x0e, 0xa8, 0x28, 0x0a, 0xc1,
	0x03, 0xba, 0x8a, 0x91, 0x4f, 0xa9, 0xe0, 0x29, 0x66, 0x7e, 0x29, 0x85, 0x12, 0xa4, 0x57, 0x1f,
	0x79, 0xef, 0x16, 0xf4, 0xe7, 0x77, 0x8b, 0x85, 0x44, 0x4e, 0xb1, 0x8c, 0x73, 0xf2, 0x04, 0xc3,
	0x6d, 0x71, 0x95, 0xc7, 0xeb, 0x35, 0xa6, 0x48, 0x63, 0x85, 0x82, 0x8f, 0xac, 0x63, 0xeb, 0xd4,
	0x99, 0x9d, 0xf8, 0xb5, 0xd4, 0x6f, 0xca, 0xfc, 0xef, 0x68, 0x34, 0x2c, 0x7f, 0xf6, 0x20, 0x47,
	0x60, 0x6f, 0x45, 0xa3, 0xff, 0xda, 0xb0, 0x1f, 0xd9, 0x5b, 0xd6, 0xbb, 0x06, 0xa7, 0xc5, 0x0f,
	0xc0, 0x0e, 0x2b, 0x9d, 0x14, 0x89, 0x9c, 0xb9, 0xff, 0xc8, 0x3e, 0x90, 0xb0, 0xba, 0x95, 0x59,
	0xcc, 0xf1, 0xd5, 0x00, 0x0f, 0x1c, 0x95, 0x6b, 0x11, 0x07, 0x20, 0xac, 0x6e, 0x96, 0x8c, 0x2b,
	0x54, 0x95, 0xdb, 0xf1, 0x5e, 0xc0, 0x6d, 0x53, 0x64, 0x02, 0x03, 0x6d, 0x54, 0x43, 0x29, 0x32,
	0x69, 0xde, 0x63, 0x47, 0x83, 0xa2, 0xd9, 0x24, 0x17, 0x30, 0x6e, 0x2b, 0x1b, 0x92, 0x8e, 0x91,
	0x8c, 0xc5, 0xaf, 0x84, 0xf7, 0x66, 0xc1, 0xee, 0xe7, 0x7d, 0xff, 0x98, 0x18, 0x40, 0x77, 0x43,
	0x1b, 0x6f, 0x67, 0x76, 0xd8, 0x18, 0xef, 0xa6, 0xfd, 0xb5, 0xde, 0x57, 0x25, 0x8b, 0xba, 0x52,
	0xef, 0xbc, 0x09, 0xec, 0x35, 0x9a, 0x04, 0xa0, 0x37, 0x67, 0x45, 0xc2, 0xa4, 0x9e, 0x8f, 0x0d,
	0x3b, 0x97, 0xcb, 0x02, 0xb9, 0x6b, 0x85, 0xd3, 0xc7, 0xb3, 0x0c, 0xd5, 0xea, 0x39, 0xd9, 0x18,
	0x06, 0x2b, 0x4d, 0xca, 0x9c, 0x2d, 0x33, 0x26, 0x83, 0x34, 0x4e, 0x24, 0xd2, 0xc0, 0x7c, 0x84,
	0x75, 0x50, 0xc7, 0x25, 0x3d, 0x53, 0x9e, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x84, 0xd1,
	0x7b, 0x34, 0x02, 0x00, 0x00,
}
