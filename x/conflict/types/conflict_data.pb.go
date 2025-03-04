// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/conflict/conflict_data.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/lavanet/lava/v2/x/pairing/types"
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

type ResponseConflict struct {
	ConflictRelayData0 *ConflictRelayData `protobuf:"bytes,1,opt,name=conflictRelayData0,proto3" json:"conflictRelayData0,omitempty"`
	ConflictRelayData1 *ConflictRelayData `protobuf:"bytes,2,opt,name=conflictRelayData1,proto3" json:"conflictRelayData1,omitempty"`
}

func (m *ResponseConflict) Reset()         { *m = ResponseConflict{} }
func (m *ResponseConflict) String() string { return proto.CompactTextString(m) }
func (*ResponseConflict) ProtoMessage()    {}
func (*ResponseConflict) Descriptor() ([]byte, []int) {
	return fileDescriptor_db493e54bcd78171, []int{0}
}
func (m *ResponseConflict) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseConflict) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseConflict.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseConflict) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseConflict.Merge(m, src)
}
func (m *ResponseConflict) XXX_Size() int {
	return m.Size()
}
func (m *ResponseConflict) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseConflict.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseConflict proto.InternalMessageInfo

func (m *ResponseConflict) GetConflictRelayData0() *ConflictRelayData {
	if m != nil {
		return m.ConflictRelayData0
	}
	return nil
}

func (m *ResponseConflict) GetConflictRelayData1() *ConflictRelayData {
	if m != nil {
		return m.ConflictRelayData1
	}
	return nil
}

type ConflictRelayData struct {
	Request *types.RelayRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Reply   *ReplyMetadata      `protobuf:"bytes,3,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (m *ConflictRelayData) Reset()         { *m = ConflictRelayData{} }
func (m *ConflictRelayData) String() string { return proto.CompactTextString(m) }
func (*ConflictRelayData) ProtoMessage()    {}
func (*ConflictRelayData) Descriptor() ([]byte, []int) {
	return fileDescriptor_db493e54bcd78171, []int{1}
}
func (m *ConflictRelayData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConflictRelayData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConflictRelayData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConflictRelayData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConflictRelayData.Merge(m, src)
}
func (m *ConflictRelayData) XXX_Size() int {
	return m.Size()
}
func (m *ConflictRelayData) XXX_DiscardUnknown() {
	xxx_messageInfo_ConflictRelayData.DiscardUnknown(m)
}

var xxx_messageInfo_ConflictRelayData proto.InternalMessageInfo

func (m *ConflictRelayData) GetRequest() *types.RelayRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ConflictRelayData) GetReply() *ReplyMetadata {
	if m != nil {
		return m.Reply
	}
	return nil
}

type ReplyMetadata struct {
	HashAllDataHash       []byte `protobuf:"bytes,1,opt,name=hash_all_data_hash,json=hashAllDataHash,proto3" json:"hash_all_data_hash,omitempty"`
	Sig                   []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	LatestBlock           int64  `protobuf:"varint,3,opt,name=latest_block,json=latestBlock,proto3" json:"latest_block,omitempty"`
	FinalizedBlocksHashes []byte `protobuf:"bytes,4,opt,name=finalized_blocks_hashes,json=finalizedBlocksHashes,proto3" json:"finalized_blocks_hashes,omitempty"`
	SigBlocks             []byte `protobuf:"bytes,5,opt,name=sig_blocks,json=sigBlocks,proto3" json:"sig_blocks,omitempty"`
}

func (m *ReplyMetadata) Reset()         { *m = ReplyMetadata{} }
func (m *ReplyMetadata) String() string { return proto.CompactTextString(m) }
func (*ReplyMetadata) ProtoMessage()    {}
func (*ReplyMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_db493e54bcd78171, []int{2}
}
func (m *ReplyMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReplyMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReplyMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplyMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMetadata.Merge(m, src)
}
func (m *ReplyMetadata) XXX_Size() int {
	return m.Size()
}
func (m *ReplyMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMetadata proto.InternalMessageInfo

func (m *ReplyMetadata) GetHashAllDataHash() []byte {
	if m != nil {
		return m.HashAllDataHash
	}
	return nil
}

func (m *ReplyMetadata) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *ReplyMetadata) GetLatestBlock() int64 {
	if m != nil {
		return m.LatestBlock
	}
	return 0
}

func (m *ReplyMetadata) GetFinalizedBlocksHashes() []byte {
	if m != nil {
		return m.FinalizedBlocksHashes
	}
	return nil
}

func (m *ReplyMetadata) GetSigBlocks() []byte {
	if m != nil {
		return m.SigBlocks
	}
	return nil
}

type FinalizationConflict struct {
	RelayReply0 *types.RelayReply `protobuf:"bytes,1,opt,name=relayReply0,proto3" json:"relayReply0,omitempty"`
	RelayReply1 *types.RelayReply `protobuf:"bytes,2,opt,name=relayReply1,proto3" json:"relayReply1,omitempty"`
}

func (m *FinalizationConflict) Reset()         { *m = FinalizationConflict{} }
func (m *FinalizationConflict) String() string { return proto.CompactTextString(m) }
func (*FinalizationConflict) ProtoMessage()    {}
func (*FinalizationConflict) Descriptor() ([]byte, []int) {
	return fileDescriptor_db493e54bcd78171, []int{3}
}
func (m *FinalizationConflict) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FinalizationConflict) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FinalizationConflict.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FinalizationConflict) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinalizationConflict.Merge(m, src)
}
func (m *FinalizationConflict) XXX_Size() int {
	return m.Size()
}
func (m *FinalizationConflict) XXX_DiscardUnknown() {
	xxx_messageInfo_FinalizationConflict.DiscardUnknown(m)
}

var xxx_messageInfo_FinalizationConflict proto.InternalMessageInfo

func (m *FinalizationConflict) GetRelayReply0() *types.RelayReply {
	if m != nil {
		return m.RelayReply0
	}
	return nil
}

func (m *FinalizationConflict) GetRelayReply1() *types.RelayReply {
	if m != nil {
		return m.RelayReply1
	}
	return nil
}

func init() {
	proto.RegisterType((*ResponseConflict)(nil), "lavanet.lava.conflict.ResponseConflict")
	proto.RegisterType((*ConflictRelayData)(nil), "lavanet.lava.conflict.ConflictRelayData")
	proto.RegisterType((*ReplyMetadata)(nil), "lavanet.lava.conflict.ReplyMetadata")
	proto.RegisterType((*FinalizationConflict)(nil), "lavanet.lava.conflict.FinalizationConflict")
}

func init() {
	proto.RegisterFile("lavanet/lava/conflict/conflict_data.proto", fileDescriptor_db493e54bcd78171)
}

var fileDescriptor_db493e54bcd78171 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x5f, 0xab, 0xd3, 0x30,
	0x18, 0xc6, 0x97, 0xd3, 0x73, 0xfc, 0x93, 0x4d, 0x9c, 0xe1, 0x1c, 0x2c, 0x03, 0xcb, 0x2c, 0x5e,
	0x4c, 0x06, 0xad, 0x9b, 0xe0, 0x85, 0x78, 0xe3, 0xa6, 0x32, 0x04, 0x6f, 0x72, 0x25, 0xde, 0x94,
	0xac, 0xcb, 0xda, 0x60, 0x6c, 0x6a, 0x93, 0x0d, 0xeb, 0xa7, 0x10, 0xbc, 0xf7, 0xe3, 0xc8, 0x2e,
	0x77, 0xe9, 0xa5, 0x6c, 0x5f, 0x44, 0x92, 0xb4, 0xd3, 0xea, 0x14, 0xe4, 0x5c, 0xf5, 0x6d, 0xf2,
	0x7b, 0x9e, 0x3c, 0xbc, 0xc9, 0x0b, 0xef, 0x73, 0xb2, 0x26, 0x19, 0x55, 0xa1, 0xfe, 0x86, 0xb1,
	0xc8, 0x96, 0x9c, 0xc5, 0xea, 0x50, 0x44, 0x0b, 0xa2, 0x48, 0x90, 0x17, 0x42, 0x09, 0x74, 0x51,
	0xa1, 0x81, 0xfe, 0x06, 0x35, 0xd1, 0x3b, 0x4f, 0x44, 0x22, 0x0c, 0x11, 0xea, 0xca, 0xc2, 0xbd,
	0x7e, 0xc3, 0x37, 0x27, 0xac, 0x60, 0x59, 0x12, 0x16, 0x94, 0x93, 0xd2, 0x12, 0xfe, 0x57, 0x00,
	0xbb, 0x98, 0xca, 0x5c, 0x64, 0x92, 0x4e, 0x2b, 0x33, 0xf4, 0x1a, 0xa2, 0xda, 0x18, 0x6b, 0xf6,
	0x19, 0x51, 0xe4, 0x81, 0x0b, 0xfa, 0x60, 0xd0, 0x1e, 0x0f, 0x82, 0xa3, 0x01, 0x82, 0xe9, 0xef,
	0x02, 0x7c, 0xc4, 0xe3, 0xa8, 0xf3, 0xc8, 0x3d, 0xb9, 0xb4, 0xf3, 0xc8, 0xff, 0x0c, 0xe0, 0xad,
	0x3f, 0x48, 0xf4, 0x04, 0x5e, 0x2d, 0xe8, 0xfb, 0x15, 0x95, 0xaa, 0x8a, 0xef, 0x37, 0x0f, 0xa9,
	0x5a, 0x12, 0x18, 0x05, 0xb6, 0x24, 0xae, 0x25, 0xe8, 0x31, 0x3c, 0x2b, 0x68, 0xce, 0x4b, 0xd7,
	0x31, 0xda, 0x7b, 0x7f, 0x09, 0x88, 0x35, 0xf3, 0x8a, 0x2a, 0xa2, 0xaf, 0x09, 0x5b, 0xc9, 0xcb,
	0xd3, 0x6b, 0x27, 0x5d, 0xc7, 0xdf, 0x00, 0x78, 0xa3, 0xb1, 0x8d, 0x86, 0x10, 0xa5, 0x44, 0xa6,
	0x11, 0xe1, 0xdc, 0x5c, 0x6b, 0xa4, 0xff, 0x4c, 0xb8, 0x0e, 0xbe, 0xa9, 0xeb, 0xa7, 0x9c, 0xeb,
	0xe8, 0x33, 0x22, 0x53, 0xd4, 0x85, 0x8e, 0x64, 0x89, 0xe9, 0x4f, 0x07, 0xeb, 0x12, 0xdd, 0x85,
	0x1d, 0x4e, 0x14, 0x95, 0x2a, 0x9a, 0x73, 0x11, 0xbf, 0x35, 0xc9, 0x1c, 0xdc, 0xb6, 0x6b, 0x13,
	0xbd, 0x84, 0x1e, 0xc1, 0xdb, 0x4b, 0x96, 0x11, 0xce, 0x3e, 0xd2, 0x85, 0xa5, 0xa4, 0x39, 0x84,
	0x4a, 0xf7, 0xd4, 0x18, 0x5d, 0x1c, 0xb6, 0x8d, 0x40, 0xce, 0xcc, 0x26, 0xba, 0x03, 0xa1, 0x64,
	0x49, 0xa5, 0x70, 0xcf, 0x0c, 0x7a, 0x5d, 0xb2, 0xc4, 0x42, 0xfe, 0x17, 0x00, 0xcf, 0x5f, 0x58,
	0x21, 0x51, 0x4c, 0x64, 0x87, 0xd7, 0x32, 0x81, 0xed, 0xc2, 0xb6, 0x2f, 0xe7, 0x65, 0xfd, 0x4c,
	0xfa, 0xff, 0xec, 0x73, 0xce, 0x4b, 0xfc, 0xab, 0xa8, 0xe9, 0x51, 0x3f, 0x88, 0xff, 0xf2, 0x18,
	0x4d, 0x9e, 0x6f, 0x76, 0x1e, 0xd8, 0xee, 0x3c, 0xf0, 0x7d, 0xe7, 0x81, 0x4f, 0x7b, 0xaf, 0xb5,
	0xdd, 0x7b, 0xad, 0x6f, 0x7b, 0xaf, 0xf5, 0x66, 0x98, 0x30, 0x95, 0xae, 0xe6, 0x41, 0x2c, 0xde,
	0x85, 0x8d, 0x89, 0x58, 0x8f, 0xc3, 0x0f, 0x3f, 0xc7, 0x4d, 0x95, 0x39, 0x95, 0xf3, 0x2b, 0x66,
	0x30, 0x1e, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0x07, 0x5f, 0xd5, 0x60, 0x94, 0x03, 0x00, 0x00,
}

func (m *ResponseConflict) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseConflict) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseConflict) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ConflictRelayData1 != nil {
		{
			size, err := m.ConflictRelayData1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConflictData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ConflictRelayData0 != nil {
		{
			size, err := m.ConflictRelayData0.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConflictData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConflictRelayData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConflictRelayData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConflictRelayData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Reply != nil {
		{
			size, err := m.Reply.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConflictData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Request != nil {
		{
			size, err := m.Request.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConflictData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReplyMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplyMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReplyMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SigBlocks) > 0 {
		i -= len(m.SigBlocks)
		copy(dAtA[i:], m.SigBlocks)
		i = encodeVarintConflictData(dAtA, i, uint64(len(m.SigBlocks)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.FinalizedBlocksHashes) > 0 {
		i -= len(m.FinalizedBlocksHashes)
		copy(dAtA[i:], m.FinalizedBlocksHashes)
		i = encodeVarintConflictData(dAtA, i, uint64(len(m.FinalizedBlocksHashes)))
		i--
		dAtA[i] = 0x22
	}
	if m.LatestBlock != 0 {
		i = encodeVarintConflictData(dAtA, i, uint64(m.LatestBlock))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(dAtA[i:], m.Sig)
		i = encodeVarintConflictData(dAtA, i, uint64(len(m.Sig)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.HashAllDataHash) > 0 {
		i -= len(m.HashAllDataHash)
		copy(dAtA[i:], m.HashAllDataHash)
		i = encodeVarintConflictData(dAtA, i, uint64(len(m.HashAllDataHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FinalizationConflict) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FinalizationConflict) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FinalizationConflict) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RelayReply1 != nil {
		{
			size, err := m.RelayReply1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConflictData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.RelayReply0 != nil {
		{
			size, err := m.RelayReply0.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConflictData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConflictData(dAtA []byte, offset int, v uint64) int {
	offset -= sovConflictData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResponseConflict) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConflictRelayData0 != nil {
		l = m.ConflictRelayData0.Size()
		n += 1 + l + sovConflictData(uint64(l))
	}
	if m.ConflictRelayData1 != nil {
		l = m.ConflictRelayData1.Size()
		n += 1 + l + sovConflictData(uint64(l))
	}
	return n
}

func (m *ConflictRelayData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovConflictData(uint64(l))
	}
	if m.Reply != nil {
		l = m.Reply.Size()
		n += 1 + l + sovConflictData(uint64(l))
	}
	return n
}

func (m *ReplyMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HashAllDataHash)
	if l > 0 {
		n += 1 + l + sovConflictData(uint64(l))
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovConflictData(uint64(l))
	}
	if m.LatestBlock != 0 {
		n += 1 + sovConflictData(uint64(m.LatestBlock))
	}
	l = len(m.FinalizedBlocksHashes)
	if l > 0 {
		n += 1 + l + sovConflictData(uint64(l))
	}
	l = len(m.SigBlocks)
	if l > 0 {
		n += 1 + l + sovConflictData(uint64(l))
	}
	return n
}

func (m *FinalizationConflict) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RelayReply0 != nil {
		l = m.RelayReply0.Size()
		n += 1 + l + sovConflictData(uint64(l))
	}
	if m.RelayReply1 != nil {
		l = m.RelayReply1.Size()
		n += 1 + l + sovConflictData(uint64(l))
	}
	return n
}

func sovConflictData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConflictData(x uint64) (n int) {
	return sovConflictData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResponseConflict) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConflictData
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
			return fmt.Errorf("proto: ResponseConflict: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseConflict: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConflictRelayData0", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
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
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConflictRelayData0 == nil {
				m.ConflictRelayData0 = &ConflictRelayData{}
			}
			if err := m.ConflictRelayData0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConflictRelayData1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
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
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConflictRelayData1 == nil {
				m.ConflictRelayData1 = &ConflictRelayData{}
			}
			if err := m.ConflictRelayData1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConflictData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConflictData
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
func (m *ConflictRelayData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConflictData
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
			return fmt.Errorf("proto: ConflictRelayData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConflictRelayData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
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
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &types.RelayRequest{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
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
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Reply == nil {
				m.Reply = &ReplyMetadata{}
			}
			if err := m.Reply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConflictData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConflictData
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
func (m *ReplyMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConflictData
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
			return fmt.Errorf("proto: ReplyMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplyMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashAllDataHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashAllDataHash = append(m.HashAllDataHash[:0], dAtA[iNdEx:postIndex]...)
			if m.HashAllDataHash == nil {
				m.HashAllDataHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlock", wireType)
			}
			m.LatestBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizedBlocksHashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalizedBlocksHashes = append(m.FinalizedBlocksHashes[:0], dAtA[iNdEx:postIndex]...)
			if m.FinalizedBlocksHashes == nil {
				m.FinalizedBlocksHashes = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigBlocks", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SigBlocks = append(m.SigBlocks[:0], dAtA[iNdEx:postIndex]...)
			if m.SigBlocks == nil {
				m.SigBlocks = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConflictData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConflictData
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
func (m *FinalizationConflict) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConflictData
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
			return fmt.Errorf("proto: FinalizationConflict: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FinalizationConflict: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayReply0", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
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
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RelayReply0 == nil {
				m.RelayReply0 = &types.RelayReply{}
			}
			if err := m.RelayReply0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayReply1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
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
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RelayReply1 == nil {
				m.RelayReply1 = &types.RelayReply{}
			}
			if err := m.RelayReply1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConflictData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConflictData
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
func skipConflictData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConflictData
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
					return 0, ErrIntOverflowConflictData
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
					return 0, ErrIntOverflowConflictData
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
				return 0, ErrInvalidLengthConflictData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConflictData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConflictData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConflictData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConflictData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConflictData = fmt.Errorf("proto: unexpected end of group")
)
