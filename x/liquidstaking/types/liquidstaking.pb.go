// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/liquidstaking/v1beta1/liquidstaking.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ValidatorStatus enumerates the status of a liquid validator.
type ValidatorStatus int32

const (
	// VALIDATOR_STATUS_UNSPECIFIED defines the default status.
	ValidatorStatusNil ValidatorStatus = 0
	// VALIDATOR_STATUS_ACTIVE defines the ...
	ValidatorStatusActive ValidatorStatus = 1
	// VALIDATOR_STATUS_DELISTING defines the ...
	ValidatorStatusDelisting ValidatorStatus = 2
	// VALIDATOR_STATUS_DELISTED defines the ...
	ValidatorStatusDelisted ValidatorStatus = 3
)

var ValidatorStatus_name = map[int32]string{
	0: "VALIDATOR_STATUS_UNSPECIFIED",
	1: "VALIDATOR_STATUS_ACTIVE",
	2: "VALIDATOR_STATUS_DELISTING",
	3: "VALIDATOR_STATUS_DELISTED",
}

var ValidatorStatus_value = map[string]int32{
	"VALIDATOR_STATUS_UNSPECIFIED": 0,
	"VALIDATOR_STATUS_ACTIVE":      1,
	"VALIDATOR_STATUS_DELISTING":   2,
	"VALIDATOR_STATUS_DELISTED":    3,
}

func (x ValidatorStatus) String() string {
	return proto.EnumName(ValidatorStatus_name, int32(x))
}

func (ValidatorStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f11ef7f6d0889fb0, []int{0}
}

// Params defines the set of params for the liquidstaking module.
type Params struct {
	LiquidBondDenom       string                 `protobuf:"bytes,1,opt,name=liquid_bond_denom,json=liquidBondDenom,proto3" json:"liquid_bond_denom,omitempty" yaml:"liquid_bond_denom"`
	WhitelistedValidators []WhitelistedValidator `protobuf:"bytes,2,rep,name=whitelisted_validators,json=whitelistedValidators,proto3" json:"whitelisted_validators" yaml:"whitelisted_validators"`
	// unstake_fee_rate specifies the ...
	UnstakeFeeRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=unstake_fee_rate,json=unstakeFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"unstake_fee_rate" yaml:"unstake_fee_rate"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f11ef7f6d0889fb0, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// WhitelistedValidator defines a ... TBD
type WhitelistedValidator struct {
	// validator_address defines the bech32-encoded address that whitelisted validator
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`
	// weight specifies the ... TBD
	Weight github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight" yaml:"weight"`
}

func (m *WhitelistedValidator) Reset()         { *m = WhitelistedValidator{} }
func (m *WhitelistedValidator) String() string { return proto.CompactTextString(m) }
func (*WhitelistedValidator) ProtoMessage()    {}
func (*WhitelistedValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_f11ef7f6d0889fb0, []int{1}
}
func (m *WhitelistedValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhitelistedValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhitelistedValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhitelistedValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhitelistedValidator.Merge(m, src)
}
func (m *WhitelistedValidator) XXX_Size() int {
	return m.Size()
}
func (m *WhitelistedValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_WhitelistedValidator.DiscardUnknown(m)
}

var xxx_messageInfo_WhitelistedValidator proto.InternalMessageInfo

// LiquidValidator defines a ... TBD
type LiquidValidator struct {
	// operator_address defines the address of the validator's operator; bech encoded in JSON.
	OperatorAddress string `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty" yaml:"operator_address"`
	// status is the liquid validator status
	Status ValidatorStatus `protobuf:"varint,2,opt,name=status,proto3,enum=crescent.liquidstaking.v1beta1.ValidatorStatus" json:"status,omitempty"`
	// liquid tokens define the liquid staked tokens
	LiquidTokens github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=liquid_tokens,json=liquidTokens,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"liquid_tokens"`
	Weight       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight" yaml:"weight"`
}

func (m *LiquidValidator) Reset()         { *m = LiquidValidator{} }
func (m *LiquidValidator) String() string { return proto.CompactTextString(m) }
func (*LiquidValidator) ProtoMessage()    {}
func (*LiquidValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_f11ef7f6d0889fb0, []int{2}
}
func (m *LiquidValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidValidator.Merge(m, src)
}
func (m *LiquidValidator) XXX_Size() int {
	return m.Size()
}
func (m *LiquidValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidValidator.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidValidator proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("crescent.liquidstaking.v1beta1.ValidatorStatus", ValidatorStatus_name, ValidatorStatus_value)
	proto.RegisterType((*Params)(nil), "crescent.liquidstaking.v1beta1.Params")
	proto.RegisterType((*WhitelistedValidator)(nil), "crescent.liquidstaking.v1beta1.WhitelistedValidator")
	proto.RegisterType((*LiquidValidator)(nil), "crescent.liquidstaking.v1beta1.LiquidValidator")
}

func init() {
	proto.RegisterFile("crescent/liquidstaking/v1beta1/liquidstaking.proto", fileDescriptor_f11ef7f6d0889fb0)
}

var fileDescriptor_f11ef7f6d0889fb0 = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdf, 0x4b, 0xdb, 0x50,
	0x14, 0xc7, 0x93, 0x2a, 0x65, 0xbb, 0x9b, 0xb6, 0x06, 0x7f, 0xd4, 0xe8, 0x92, 0x12, 0xd8, 0x90,
	0x81, 0x09, 0xba, 0x31, 0x46, 0x19, 0x8c, 0xd6, 0x54, 0x17, 0x10, 0x27, 0x69, 0x54, 0xd8, 0x4b,
	0x48, 0x93, 0x6b, 0x0c, 0x6d, 0x72, 0xbb, 0xdc, 0x5b, 0x3b, 0x9f, 0xf7, 0x22, 0x7d, 0x1a, 0x7b,
	0xda, 0x4b, 0x41, 0xd8, 0x3f, 0xb2, 0x47, 0x1f, 0x7d, 0x1c, 0x3e, 0x94, 0xa1, 0x30, 0xf6, 0xec,
	0x5f, 0x30, 0x7a, 0x93, 0x56, 0x4d, 0xbb, 0x0d, 0x61, 0x4f, 0xed, 0x3d, 0xe7, 0x7c, 0xce, 0xfd,
	0x9e, 0xef, 0x3d, 0x04, 0xac, 0xda, 0x21, 0xc4, 0x36, 0x0c, 0x88, 0x52, 0xf7, 0xde, 0x37, 0x3d,
	0x07, 0x13, 0xab, 0xe6, 0x05, 0xae, 0x72, 0xb8, 0x52, 0x85, 0xc4, 0x5a, 0xb9, 0x1d, 0x95, 0x1b,
	0x21, 0x22, 0x88, 0x13, 0xfa, 0x8c, 0x7c, 0x3b, 0x1b, 0x33, 0xfc, 0xb4, 0x8b, 0x5c, 0x44, 0x4b,
	0x95, 0xde, 0xbf, 0x88, 0xe2, 0xe7, 0x6d, 0x84, 0x7d, 0x84, 0xcd, 0x28, 0x11, 0x1d, 0xe2, 0x94,
	0x10, 0x9d, 0x94, 0xaa, 0x85, 0xe1, 0xe0, 0x66, 0x1b, 0x79, 0x41, 0x9c, 0x17, 0x5d, 0x84, 0xdc,
	0x3a, 0x54, 0xe8, 0xa9, 0xda, 0xdc, 0x57, 0x88, 0xe7, 0x43, 0x4c, 0x2c, 0xbf, 0x11, 0x15, 0x48,
	0x3f, 0x53, 0x20, 0xbd, 0x6d, 0x85, 0x96, 0x8f, 0xb9, 0x37, 0x60, 0x2a, 0x52, 0x65, 0x56, 0x51,
	0xe0, 0x98, 0x0e, 0x0c, 0x90, 0x9f, 0x63, 0xf3, 0xec, 0xd2, 0xfd, 0xd2, 0xe2, 0x55, 0x57, 0xcc,
	0x1d, 0x59, 0x7e, 0xbd, 0x20, 0x0d, 0x95, 0x48, 0x7a, 0x26, 0x8a, 0x95, 0x50, 0xe0, 0xa8, 0xbd,
	0x08, 0xf7, 0x99, 0x05, 0xb3, 0xad, 0x03, 0x8f, 0xc0, 0xba, 0x87, 0x09, 0x74, 0xcc, 0x43, 0xab,
	0xee, 0x39, 0x16, 0x41, 0x21, 0xce, 0xa5, 0xf2, 0x63, 0x4b, 0x0f, 0x56, 0x9f, 0xcb, 0x7f, 0x37,
	0x42, 0xde, 0xbb, 0xa6, 0x77, 0xfb, 0x70, 0xe9, 0xf1, 0x69, 0x57, 0x64, 0xae, 0xba, 0xe2, 0xa3,
	0x48, 0xc9, 0xe8, 0x1b, 0x24, 0x7d, 0xa6, 0x35, 0x02, 0xc6, 0x1c, 0x06, 0xd9, 0x66, 0xd0, 0xbb,
	0x07, 0x9a, 0xfb, 0x10, 0x9a, 0xa1, 0x45, 0x60, 0x6e, 0x8c, 0x4e, 0xa7, 0xf5, 0xfa, 0x9e, 0x77,
	0xc5, 0x27, 0xae, 0x47, 0x0e, 0x9a, 0x55, 0xd9, 0x46, 0x7e, 0xec, 0x72, 0xfc, 0xb3, 0x8c, 0x9d,
	0x9a, 0x42, 0x8e, 0x1a, 0x10, 0xcb, 0x2a, 0xb4, 0xaf, 0xba, 0xe2, 0x5c, 0xa4, 0x20, 0xd9, 0x4f,
	0xd2, 0x27, 0xe3, 0xd0, 0x3a, 0x84, 0xba, 0x45, 0x60, 0xe1, 0xde, 0xf1, 0x89, 0xc8, 0x7c, 0x39,
	0x11, 0x19, 0xe9, 0x1b, 0x0b, 0xa6, 0x47, 0x4d, 0xc5, 0x69, 0x60, 0x6a, 0xa0, 0xde, 0xb4, 0x1c,
	0x27, 0x84, 0x18, 0x0f, 0xdb, 0x3e, 0x54, 0x22, 0xe9, 0xd9, 0x41, 0xac, 0x18, 0x85, 0xb8, 0x3d,
	0x90, 0x6e, 0x41, 0xcf, 0x3d, 0x20, 0xb9, 0x14, 0xe5, 0x5f, 0xdf, 0x79, 0xb0, 0x89, 0xd8, 0x5a,
	0xda, 0x45, 0xd2, 0xe3, 0x76, 0x85, 0xf1, 0xde, 0x18, 0xd2, 0x79, 0x0a, 0x64, 0x36, 0xe9, 0x73,
	0x5d, 0xab, 0x5f, 0x07, 0x59, 0xd4, 0x80, 0xe1, 0x08, 0xf1, 0x0b, 0xd7, 0x3e, 0x25, 0x2b, 0x24,
	0x3d, 0xd3, 0x0f, 0xf5, 0xa5, 0x6f, 0x80, 0x34, 0x26, 0x16, 0x69, 0x62, 0x2a, 0x7d, 0x72, 0x55,
	0xf9, 0xd7, 0x86, 0x0c, 0x24, 0x54, 0x28, 0xa6, 0xc7, 0x38, 0x57, 0x01, 0x13, 0xf1, 0x8a, 0x12,
	0x54, 0x83, 0x01, 0x8e, 0xdf, 0x58, 0xbe, 0x83, 0x15, 0x5a, 0x40, 0xf4, 0x87, 0x51, 0x13, 0x83,
	0xf6, 0xb8, 0x61, 0xec, 0xf8, 0xff, 0x35, 0x96, 0xee, 0xc7, 0xaf, 0x13, 0x91, 0x79, 0xfa, 0x31,
	0x05, 0x32, 0x89, 0x99, 0xb8, 0x97, 0x60, 0x71, 0xb7, 0xb8, 0xa9, 0xa9, 0x45, 0xe3, 0xad, 0x6e,
	0x56, 0x8c, 0xa2, 0xb1, 0x53, 0x31, 0x77, 0xb6, 0x2a, 0xdb, 0xe5, 0x35, 0x6d, 0x5d, 0x2b, 0xab,
	0x59, 0x86, 0x9f, 0x6d, 0x77, 0xf2, 0x5c, 0x02, 0xdb, 0xf2, 0xea, 0xdc, 0x0b, 0x30, 0x37, 0x44,
	0x16, 0xd7, 0x0c, 0x6d, 0xb7, 0x9c, 0x65, 0xf9, 0xf9, 0x76, 0x27, 0x3f, 0x93, 0x80, 0x8a, 0x36,
	0xf1, 0x0e, 0x21, 0xf7, 0x0a, 0xf0, 0x43, 0x9c, 0x5a, 0xde, 0xd4, 0x2a, 0x86, 0xb6, 0xb5, 0x91,
	0x4d, 0xf1, 0x8b, 0xed, 0x4e, 0x3e, 0x97, 0x40, 0x55, 0xba, 0xd2, 0x5e, 0xe0, 0x72, 0x05, 0x30,
	0xff, 0x07, 0xba, 0xac, 0x66, 0xc7, 0xf8, 0x85, 0x76, 0x27, 0x3f, 0x37, 0x12, 0x86, 0x0e, 0x3f,
	0x7e, 0xfc, 0x55, 0x60, 0x4a, 0xc6, 0xe9, 0x85, 0xc0, 0x9e, 0x5d, 0x08, 0xec, 0x8f, 0x0b, 0x81,
	0xfd, 0x74, 0x29, 0x30, 0x67, 0x97, 0x02, 0xf3, 0xfd, 0x52, 0x60, 0xde, 0x15, 0x6e, 0x5a, 0x1d,
	0xaf, 0xc6, 0x72, 0x00, 0x49, 0x0b, 0x85, 0xb5, 0x41, 0x40, 0xf9, 0x90, 0xf8, 0x18, 0xd3, 0x27,
	0xa8, 0xa6, 0xe9, 0xb7, 0xee, 0xd9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xdd, 0xa3, 0x67,
	0xb3, 0x05, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.UnstakeFeeRate.Size()
		i -= size
		if _, err := m.UnstakeFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidstaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.WhitelistedValidators) > 0 {
		for iNdEx := len(m.WhitelistedValidators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WhitelistedValidators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLiquidstaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.LiquidBondDenom) > 0 {
		i -= len(m.LiquidBondDenom)
		copy(dAtA[i:], m.LiquidBondDenom)
		i = encodeVarintLiquidstaking(dAtA, i, uint64(len(m.LiquidBondDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WhitelistedValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhitelistedValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhitelistedValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidstaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintLiquidstaking(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LiquidValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidstaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.LiquidTokens.Size()
		i -= size
		if _, err := m.LiquidTokens.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidstaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Status != 0 {
		i = encodeVarintLiquidstaking(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintLiquidstaking(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLiquidstaking(dAtA []byte, offset int, v uint64) int {
	offset -= sovLiquidstaking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LiquidBondDenom)
	if l > 0 {
		n += 1 + l + sovLiquidstaking(uint64(l))
	}
	if len(m.WhitelistedValidators) > 0 {
		for _, e := range m.WhitelistedValidators {
			l = e.Size()
			n += 1 + l + sovLiquidstaking(uint64(l))
		}
	}
	l = m.UnstakeFeeRate.Size()
	n += 1 + l + sovLiquidstaking(uint64(l))
	return n
}

func (m *WhitelistedValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovLiquidstaking(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovLiquidstaking(uint64(l))
	return n
}

func (m *LiquidValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovLiquidstaking(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovLiquidstaking(uint64(m.Status))
	}
	l = m.LiquidTokens.Size()
	n += 1 + l + sovLiquidstaking(uint64(l))
	l = m.Weight.Size()
	n += 1 + l + sovLiquidstaking(uint64(l))
	return n
}

func sovLiquidstaking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLiquidstaking(x uint64) (n int) {
	return sovLiquidstaking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidstaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidBondDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidstaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidBondDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistedValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidstaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhitelistedValidators = append(m.WhitelistedValidators, WhitelistedValidator{})
			if err := m.WhitelistedValidators[len(m.WhitelistedValidators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakeFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidstaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UnstakeFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidstaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WhitelistedValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidstaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WhitelistedValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhitelistedValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidstaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidstaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidstaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LiquidValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidstaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LiquidValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidstaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidstaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ValidatorStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidstaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidTokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidstaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidstaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidstaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLiquidstaking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiquidstaking
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLiquidstaking
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLiquidstaking
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLiquidstaking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLiquidstaking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLiquidstaking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLiquidstaking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiquidstaking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLiquidstaking = fmt.Errorf("proto: unexpected end of group")
)
