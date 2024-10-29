// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/sequencer/sequencer.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Sequencer defines a sequencer identified by its' address (sequencerAddress).
// The sequencer could be attached to only one rollapp (rollappId).
type Sequencer struct {
	// address is the bech32-encoded address of the sequencer account which is the account that the message was sent from.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// pubkey is the public key of the sequencers' dymint client, as a Protobuf Any.
	DymintPubKey *types.Any `protobuf:"bytes,2,opt,name=dymintPubKey,proto3" json:"dymintPubKey,omitempty"`
	// rollappId defines the rollapp to which the sequencer belongs.
	RollappId string `protobuf:"bytes,3,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// metadata defines the extra information for the sequencer.
	Metadata SequencerMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata"`
	// jailed defined whether the sequencer has been jailed from bonded status or not.
	Jailed   bool `protobuf:"varint,5,opt,name=jailed,proto3" json:"jailed,omitempty"`
	Proposer bool `protobuf:"varint,6,opt,name=proposer,proto3" json:"proposer,omitempty"` // Deprecated: Do not use.
	// status is the sequencer status (bonded/unbonding/unbonded).
	Status OperatingStatus `protobuf:"varint,7,opt,name=status,proto3,enum=dymensionxyz.dymension.sequencer.OperatingStatus" json:"status,omitempty"`
	// tokens define the delegated tokens (incl. self-delegation).
	Tokens github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,8,rep,name=tokens,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"tokens"`
	// unbond_request_height stores the height at which this sequencer has
	// requested to unbond.
	UnbondRequestHeight int64 `protobuf:"varint,9,opt,name=unbond_request_height,json=unbondRequestHeight,proto3" json:"unbond_request_height,omitempty"`
	// unbond_time defines the time when the sequencer will complete unbonding.
	UnbondTime time.Time `protobuf:"bytes,10,opt,name=unbond_time,json=unbondTime,proto3,stdtime" json:"unbond_time"`
	// notice_period_time defines the time when the sequencer will finish it's notice period if started
	NoticePeriodTime time.Time `protobuf:"bytes,11,opt,name=notice_period_time,json=noticePeriodTime,proto3,stdtime" json:"notice_period_time"`
	// RewardAddr is a bech32 encoded sdk acc address
	RewardAddr string `protobuf:"bytes,12,opt,name=reward_addr,json=rewardAddr,proto3" json:"reward_addr,omitempty"`
	// WhitelistedRelayers is an array of the whitelisted relayer addresses. Addresses are bech32-encoded strings.
	WhitelistedRelayers []string `protobuf:"bytes,13,rep,name=whitelisted_relayers,json=whitelistedRelayers,proto3" json:"whitelisted_relayers,omitempty"`
}

func (m *Sequencer) Reset()         { *m = Sequencer{} }
func (m *Sequencer) String() string { return proto.CompactTextString(m) }
func (*Sequencer) ProtoMessage()    {}
func (*Sequencer) Descriptor() ([]byte, []int) {
	return fileDescriptor_997b8663a5fc0f58, []int{0}
}
func (m *Sequencer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sequencer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sequencer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sequencer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sequencer.Merge(m, src)
}
func (m *Sequencer) XXX_Size() int {
	return m.Size()
}
func (m *Sequencer) XXX_DiscardUnknown() {
	xxx_messageInfo_Sequencer.DiscardUnknown(m)
}

var xxx_messageInfo_Sequencer proto.InternalMessageInfo

func (m *Sequencer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Sequencer) GetDymintPubKey() *types.Any {
	if m != nil {
		return m.DymintPubKey
	}
	return nil
}

func (m *Sequencer) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *Sequencer) GetMetadata() SequencerMetadata {
	if m != nil {
		return m.Metadata
	}
	return SequencerMetadata{}
}

func (m *Sequencer) GetJailed() bool {
	if m != nil {
		return m.Jailed
	}
	return false
}

// Deprecated: Do not use.
func (m *Sequencer) GetProposer() bool {
	if m != nil {
		return m.Proposer
	}
	return false
}

func (m *Sequencer) GetStatus() OperatingStatus {
	if m != nil {
		return m.Status
	}
	return Unbonded
}

func (m *Sequencer) GetTokens() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *Sequencer) GetUnbondRequestHeight() int64 {
	if m != nil {
		return m.UnbondRequestHeight
	}
	return 0
}

func (m *Sequencer) GetUnbondTime() time.Time {
	if m != nil {
		return m.UnbondTime
	}
	return time.Time{}
}

func (m *Sequencer) GetNoticePeriodTime() time.Time {
	if m != nil {
		return m.NoticePeriodTime
	}
	return time.Time{}
}

func (m *Sequencer) GetRewardAddr() string {
	if m != nil {
		return m.RewardAddr
	}
	return ""
}

func (m *Sequencer) GetWhitelistedRelayers() []string {
	if m != nil {
		return m.WhitelistedRelayers
	}
	return nil
}

// BondReduction defines an object which holds the information about the sequencer and its queued unbonding amount
type BondReduction struct {
	// sequencer_address is the bech32-encoded address of the sequencer account which is the account that the message was sent from.
	SequencerAddress string `protobuf:"bytes,1,opt,name=sequencer_address,json=sequencerAddress,proto3" json:"sequencer_address,omitempty"`
	// decrease_bond_amount is the amount of tokens to be unbonded.
	DecreaseBondAmount types1.Coin `protobuf:"bytes,2,opt,name=decrease_bond_amount,json=decreaseBondAmount,proto3" json:"decrease_bond_amount"`
	// decrease_bond_time defines, if unbonding, the min time for the sequencer to complete unbonding.
	DecreaseBondTime time.Time `protobuf:"bytes,3,opt,name=decrease_bond_time,json=decreaseBondTime,proto3,stdtime" json:"decrease_bond_time"`
}

func (m *BondReduction) Reset()         { *m = BondReduction{} }
func (m *BondReduction) String() string { return proto.CompactTextString(m) }
func (*BondReduction) ProtoMessage()    {}
func (*BondReduction) Descriptor() ([]byte, []int) {
	return fileDescriptor_997b8663a5fc0f58, []int{1}
}
func (m *BondReduction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BondReduction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BondReduction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BondReduction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BondReduction.Merge(m, src)
}
func (m *BondReduction) XXX_Size() int {
	return m.Size()
}
func (m *BondReduction) XXX_DiscardUnknown() {
	xxx_messageInfo_BondReduction.DiscardUnknown(m)
}

var xxx_messageInfo_BondReduction proto.InternalMessageInfo

func (m *BondReduction) GetSequencerAddress() string {
	if m != nil {
		return m.SequencerAddress
	}
	return ""
}

func (m *BondReduction) GetDecreaseBondAmount() types1.Coin {
	if m != nil {
		return m.DecreaseBondAmount
	}
	return types1.Coin{}
}

func (m *BondReduction) GetDecreaseBondTime() time.Time {
	if m != nil {
		return m.DecreaseBondTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Sequencer)(nil), "dymensionxyz.dymension.sequencer.Sequencer")
	proto.RegisterType((*BondReduction)(nil), "dymensionxyz.dymension.sequencer.BondReduction")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/sequencer/sequencer.proto", fileDescriptor_997b8663a5fc0f58)
}

var fileDescriptor_997b8663a5fc0f58 = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6f, 0xd3, 0x3e,
	0x1c, 0x6d, 0xd6, 0xad, 0x6b, 0xdd, 0xed, 0xaf, 0xfd, 0xbd, 0x02, 0xde, 0x84, 0xd2, 0x68, 0xa7,
	0x48, 0x68, 0xc9, 0xda, 0x49, 0x70, 0x6e, 0x11, 0x12, 0x13, 0x42, 0x8c, 0x0c, 0x2e, 0x5c, 0x22,
	0x37, 0x31, 0x69, 0x58, 0x63, 0x07, 0xdb, 0xd9, 0x16, 0x3e, 0xc5, 0x3e, 0x07, 0x67, 0x3e, 0xc4,
	0xc4, 0x69, 0x47, 0x4e, 0x0c, 0xad, 0x9f, 0x82, 0x1b, 0x8a, 0xe3, 0x66, 0x1d, 0x08, 0xaa, 0x9d,
	0x92, 0x9f, 0x9f, 0xdf, 0xcb, 0xef, 0xf7, 0xfc, 0x1c, 0xb0, 0x17, 0xe6, 0x09, 0xa1, 0x22, 0x66,
	0xf4, 0x2c, 0xff, 0xe4, 0x56, 0x85, 0x2b, 0xc8, 0xc7, 0x8c, 0xd0, 0x80, 0xf0, 0x9b, 0x37, 0x27,
	0xe5, 0x4c, 0x32, 0x68, 0xcd, 0x33, 0x9c, 0xaa, 0x70, 0xaa, 0x7d, 0xdb, 0x5b, 0x01, 0x13, 0x09,
	0x13, 0xbe, 0xda, 0xef, 0x96, 0x45, 0x49, 0xde, 0xde, 0x8a, 0x18, 0x8b, 0x26, 0xc4, 0x55, 0xd5,
	0x28, 0x7b, 0xef, 0x62, 0x9a, 0x6b, 0xa8, 0x13, 0xb1, 0x88, 0x95, 0x94, 0xe2, 0x4d, 0xaf, 0x76,
	0x7f, 0x27, 0xc8, 0x38, 0x21, 0x42, 0xe2, 0x24, 0xd5, 0x1b, 0xcc, 0x52, 0xdf, 0x1d, 0x61, 0x41,
	0xdc, 0x93, 0xde, 0x88, 0x48, 0xdc, 0x73, 0x03, 0x16, 0x53, 0x8d, 0x3f, 0xd0, 0x78, 0x22, 0x22,
	0xf7, 0xa4, 0x57, 0x3c, 0x34, 0xe0, 0x2e, 0x9c, 0x3c, 0x21, 0x12, 0x87, 0x58, 0x62, 0x4d, 0x78,
	0xb2, 0x90, 0xc0, 0x52, 0xc2, 0xb1, 0x8c, 0x69, 0xe4, 0x0b, 0x89, 0x65, 0xa6, 0x87, 0xde, 0xf9,
	0xb9, 0x02, 0x5a, 0x47, 0xb3, 0x4d, 0x10, 0x81, 0x55, 0x1c, 0x86, 0x9c, 0x08, 0x81, 0x0c, 0xcb,
	0xb0, 0x5b, 0xde, 0xac, 0x84, 0x1e, 0x58, 0x0b, 0xf3, 0x24, 0xa6, 0xf2, 0x30, 0x1b, 0xbd, 0x20,
	0x39, 0x5a, 0xb2, 0x0c, 0xbb, 0xdd, 0xef, 0x38, 0xa5, 0x05, 0xce, 0xcc, 0x02, 0x67, 0x40, 0xf3,
	0x21, 0xfa, 0xfa, 0x65, 0xb7, 0xa3, 0xad, 0x0d, 0x78, 0x9e, 0x4a, 0xe6, 0x94, 0x2c, 0xef, 0x96,
	0x06, 0x7c, 0x08, 0x5a, 0x9c, 0x4d, 0x26, 0x38, 0x4d, 0x0f, 0x42, 0x54, 0x57, 0xdf, 0xbb, 0x59,
	0x80, 0x6f, 0x41, 0x73, 0x36, 0x24, 0x5a, 0x56, 0x5f, 0xdb, 0x77, 0x16, 0x1d, 0xaf, 0x53, 0x8d,
	0xf2, 0x52, 0x53, 0x87, 0xcb, 0x17, 0xdf, 0xbb, 0x35, 0xaf, 0x92, 0x82, 0xf7, 0x41, 0xe3, 0x03,
	0x8e, 0x27, 0x24, 0x44, 0x2b, 0x96, 0x61, 0x37, 0x3d, 0x5d, 0x41, 0x13, 0x34, 0x53, 0xce, 0x52,
	0x26, 0x08, 0x47, 0x8d, 0x02, 0x19, 0x2e, 0x21, 0xc3, 0xab, 0xd6, 0xe0, 0x01, 0x68, 0x94, 0xc6,
	0xa1, 0x55, 0xcb, 0xb0, 0xff, 0xeb, 0xf7, 0x16, 0x37, 0xf3, 0x6a, 0x66, 0xf9, 0x91, 0x22, 0x7a,
	0x5a, 0x00, 0x06, 0xa0, 0x21, 0xd9, 0x31, 0xa1, 0x02, 0x35, 0xad, 0xba, 0xdd, 0xee, 0x6f, 0x39,
	0xda, 0xac, 0x22, 0x27, 0x8e, 0xce, 0x89, 0xf3, 0x94, 0xc5, 0x74, 0xb8, 0x57, 0x74, 0xff, 0xf9,
	0xaa, 0x6b, 0x47, 0xb1, 0x1c, 0x67, 0x23, 0x27, 0x60, 0x89, 0x0e, 0xad, 0x7e, 0xec, 0x8a, 0xf0,
	0xd8, 0x95, 0x79, 0x4a, 0x84, 0x22, 0x08, 0x4f, 0x4b, 0xc3, 0x3e, 0xb8, 0x97, 0xd1, 0x11, 0xa3,
	0xa1, 0xcf, 0x8b, 0x86, 0x84, 0xf4, 0xc7, 0x24, 0x8e, 0xc6, 0x12, 0xb5, 0x2c, 0xc3, 0xae, 0x7b,
	0x9b, 0x25, 0xe8, 0x95, 0xd8, 0x73, 0x05, 0xc1, 0x67, 0xa0, 0xad, 0x39, 0x45, 0x92, 0x11, 0x50,
	0xae, 0x6f, 0xff, 0x71, 0xc6, 0x6f, 0x66, 0x31, 0x1f, 0x36, 0x8b, 0xf6, 0xce, 0xaf, 0xba, 0x86,
	0x07, 0x4a, 0x62, 0x01, 0x41, 0x0f, 0x40, 0xca, 0x64, 0x1c, 0x10, 0x3f, 0x25, 0x3c, 0x66, 0x5a,
	0xad, 0x7d, 0x07, 0xb5, 0x8d, 0x92, 0x7f, 0xa8, 0xe8, 0x4a, 0xb3, 0x0b, 0xda, 0x9c, 0x9c, 0x62,
	0x1e, 0xfa, 0x45, 0x22, 0xd1, 0x9a, 0x4a, 0x0b, 0x28, 0x97, 0x06, 0x61, 0xc8, 0x61, 0x0f, 0x74,
	0x4e, 0xc7, 0xb1, 0x24, 0x93, 0x58, 0x48, 0x52, 0x0c, 0x3d, 0xc1, 0x39, 0xe1, 0x02, 0xad, 0x5b,
	0x75, 0xbb, 0xe5, 0x6d, 0xce, 0x61, 0x9e, 0x86, 0x76, 0xa6, 0x06, 0x58, 0x1f, 0x2a, 0x13, 0xc2,
	0x2c, 0x90, 0x31, 0xa3, 0xf0, 0x11, 0xf8, 0xbf, 0x3a, 0x3e, 0xff, 0xf6, 0x4d, 0xd8, 0xa8, 0x80,
	0x81, 0xbe, 0x12, 0xaf, 0x41, 0x27, 0x24, 0x01, 0x27, 0x58, 0x10, 0x5f, 0x99, 0x86, 0x13, 0x96,
	0x51, 0xa9, 0xaf, 0xc6, 0x3f, 0x0e, 0xb5, 0x8c, 0x24, 0x9c, 0x91, 0x8b, 0x16, 0x06, 0x8a, 0x5a,
	0x38, 0x77, 0x5b, 0x52, 0x39, 0x57, 0xbf, 0x8b, 0x73, 0xf3, 0xaa, 0xc5, 0x86, 0xe1, 0xe1, 0xc5,
	0xb5, 0x69, 0x5c, 0x5e, 0x9b, 0xc6, 0x8f, 0x6b, 0xd3, 0x38, 0x9f, 0x9a, 0xb5, 0xcb, 0xa9, 0x59,
	0xfb, 0x36, 0x35, 0x6b, 0xef, 0x1e, 0xcf, 0x85, 0xea, 0x2f, 0xff, 0x8f, 0x93, 0x7d, 0xf7, 0x6c,
	0xee, 0x27, 0xa2, 0x82, 0x36, 0x6a, 0xa8, 0x0e, 0xf6, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x05,
	0x14, 0xe1, 0x41, 0xa0, 0x05, 0x00, 0x00,
}

func (m *Sequencer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sequencer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sequencer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WhitelistedRelayers) > 0 {
		for iNdEx := len(m.WhitelistedRelayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WhitelistedRelayers[iNdEx])
			copy(dAtA[i:], m.WhitelistedRelayers[iNdEx])
			i = encodeVarintSequencer(dAtA, i, uint64(len(m.WhitelistedRelayers[iNdEx])))
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.RewardAddr) > 0 {
		i -= len(m.RewardAddr)
		copy(dAtA[i:], m.RewardAddr)
		i = encodeVarintSequencer(dAtA, i, uint64(len(m.RewardAddr)))
		i--
		dAtA[i] = 0x62
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.NoticePeriodTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.NoticePeriodTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSequencer(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x5a
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.UnbondTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.UnbondTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintSequencer(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x52
	if m.UnbondRequestHeight != 0 {
		i = encodeVarintSequencer(dAtA, i, uint64(m.UnbondRequestHeight))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSequencer(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.Status != 0 {
		i = encodeVarintSequencer(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	if m.Proposer {
		i--
		if m.Proposer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Jailed {
		i--
		if m.Jailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSequencer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintSequencer(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.DymintPubKey != nil {
		{
			size, err := m.DymintPubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSequencer(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSequencer(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BondReduction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BondReduction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BondReduction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n5, err5 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.DecreaseBondTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.DecreaseBondTime):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintSequencer(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.DecreaseBondAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSequencer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.SequencerAddress) > 0 {
		i -= len(m.SequencerAddress)
		copy(dAtA[i:], m.SequencerAddress)
		i = encodeVarintSequencer(dAtA, i, uint64(len(m.SequencerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSequencer(dAtA []byte, offset int, v uint64) int {
	offset -= sovSequencer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Sequencer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSequencer(uint64(l))
	}
	if m.DymintPubKey != nil {
		l = m.DymintPubKey.Size()
		n += 1 + l + sovSequencer(uint64(l))
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovSequencer(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovSequencer(uint64(l))
	if m.Jailed {
		n += 2
	}
	if m.Proposer {
		n += 2
	}
	if m.Status != 0 {
		n += 1 + sovSequencer(uint64(m.Status))
	}
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovSequencer(uint64(l))
		}
	}
	if m.UnbondRequestHeight != 0 {
		n += 1 + sovSequencer(uint64(m.UnbondRequestHeight))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.UnbondTime)
	n += 1 + l + sovSequencer(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.NoticePeriodTime)
	n += 1 + l + sovSequencer(uint64(l))
	l = len(m.RewardAddr)
	if l > 0 {
		n += 1 + l + sovSequencer(uint64(l))
	}
	if len(m.WhitelistedRelayers) > 0 {
		for _, s := range m.WhitelistedRelayers {
			l = len(s)
			n += 1 + l + sovSequencer(uint64(l))
		}
	}
	return n
}

func (m *BondReduction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SequencerAddress)
	if l > 0 {
		n += 1 + l + sovSequencer(uint64(l))
	}
	l = m.DecreaseBondAmount.Size()
	n += 1 + l + sovSequencer(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.DecreaseBondTime)
	n += 1 + l + sovSequencer(uint64(l))
	return n
}

func sovSequencer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSequencer(x uint64) (n int) {
	return sovSequencer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Sequencer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSequencer
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
			return fmt.Errorf("proto: Sequencer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sequencer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DymintPubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DymintPubKey == nil {
				m.DymintPubKey = &types.Any{}
			}
			if err := m.DymintPubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Jailed = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Proposer = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OperatingStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, types1.Coin{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondRequestHeight", wireType)
			}
			m.UnbondRequestHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondRequestHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.UnbondTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoticePeriodTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.NoticePeriodTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistedRelayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhitelistedRelayers = append(m.WhitelistedRelayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSequencer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSequencer
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
func (m *BondReduction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSequencer
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
			return fmt.Errorf("proto: BondReduction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BondReduction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequencerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SequencerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DecreaseBondAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DecreaseBondAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DecreaseBondTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSequencer
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
				return ErrInvalidLengthSequencer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSequencer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.DecreaseBondTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSequencer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSequencer
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
func skipSequencer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSequencer
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
					return 0, ErrIntOverflowSequencer
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
					return 0, ErrIntOverflowSequencer
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
				return 0, ErrInvalidLengthSequencer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSequencer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSequencer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSequencer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSequencer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSequencer = fmt.Errorf("proto: unexpected end of group")
)
