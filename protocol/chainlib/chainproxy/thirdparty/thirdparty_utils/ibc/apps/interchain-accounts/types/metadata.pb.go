// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/interchain_accounts/v1/metadata.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Metadata defines a set of protocol specific data encoded into the ICS27 channel version bytestring
// See ICS004: https://github.com/cosmos/ibc/tree/master/spec/core/ics-0channel-and-packet-semantics#Versioning
type Metadata struct {
	// version defines the ICS27 protocol version
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// controller_connection_id is the connection identifier associated with the controller chain
	ControllerConnectionId string `protobuf:"bytes,2,opt,name=controller_connection_id,json=controllerConnectionId,proto3" json:"controller_connection_id,omitempty" yaml:"controller_connection_id"`
	// host_connection_id is the connection identifier associated with the host chain
	HostConnectionId string `protobuf:"bytes,3,opt,name=host_connection_id,json=hostConnectionId,proto3" json:"host_connection_id,omitempty" yaml:"host_connection_id"`
	// address defines the interchain account address to be fulfilled upon the OnChanOpenTry handshake step
	// NOTE: the address field is empty on the OnChanOpenInit handshake step
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// encoding defines the supported codec format
	Encoding string `protobuf:"bytes,5,opt,name=encoding,proto3" json:"encoding,omitempty"`
	// tx_type defines the type of transactions the interchain account can execute
	TxType string `protobuf:"bytes,6,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c29c32e397d1f21e, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Metadata) GetControllerConnectionId() string {
	if m != nil {
		return m.ControllerConnectionId
	}
	return ""
}

func (m *Metadata) GetHostConnectionId() string {
	if m != nil {
		return m.HostConnectionId
	}
	return ""
}

func (m *Metadata) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Metadata) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *Metadata) GetTxType() string {
	if m != nil {
		return m.TxType
	}
	return ""
}

func init() {
	proto.RegisterType((*Metadata)(nil), "ibc.applications.interchain_accounts.v1.Metadata")
}

func init() {
	proto.RegisterFile("ibc/applications/interchain_accounts/v1/metadata.proto", fileDescriptor_c29c32e397d1f21e)
}

var fileDescriptor_c29c32e397d1f21e = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xeb, 0x40,
	0x18, 0xc5, 0x9b, 0xde, 0x6b, 0x5b, 0x67, 0x25, 0x83, 0xe8, 0x58, 0x30, 0x95, 0xb8, 0xd0, 0x4d,
	0x33, 0x54, 0xa1, 0x82, 0xcb, 0x8a, 0x0b, 0x11, 0x37, 0xc5, 0x95, 0x20, 0x61, 0x32, 0x33, 0xa4,
	0x03, 0xc9, 0x7c, 0x21, 0x33, 0x0d, 0xed, 0x5b, 0xf8, 0x06, 0xbe, 0x8e, 0xcb, 0x2e, 0x5d, 0x15,
	0x69, 0xdf, 0xa0, 0x4f, 0x20, 0x49, 0x6b, 0xeb, 0xdf, 0x5d, 0x4e, 0xce, 0x39, 0xbf, 0x8f, 0xe1,
	0xa0, 0xae, 0x0a, 0x39, 0x65, 0x69, 0x1a, 0x2b, 0xce, 0xac, 0x02, 0x6d, 0xa8, 0xd2, 0x56, 0x66,
	0x7c, 0xc0, 0x94, 0x0e, 0x18, 0xe7, 0x30, 0xd4, 0xd6, 0xd0, 0xbc, 0x43, 0x13, 0x69, 0x99, 0x60,
	0x96, 0xf9, 0x69, 0x06, 0x16, 0xf0, 0x89, 0x0a, 0xb9, 0xff, 0xb9, 0xe7, 0xff, 0xd2, 0xf3, 0xf3,
	0x4e, 0x73, 0x37, 0x82, 0x08, 0xca, 0x0e, 0x2d, 0xbe, 0x96, 0x75, 0xef, 0xb9, 0x8a, 0x1a, 0x77,
	0x2b, 0x22, 0x26, 0xa8, 0x9e, 0xcb, 0xcc, 0x28, 0xd0, 0xc4, 0x39, 0x72, 0x4e, 0xb7, 0xfb, 0x1f,
	0x12, 0x3f, 0x22, 0xc2, 0x41, 0xdb, 0x0c, 0xe2, 0x58, 0x66, 0x01, 0x07, 0xad, 0x25, 0x2f, 0xae,
	0x05, 0x4a, 0x90, 0x6a, 0x11, 0xed, 0x1d, 0x2f, 0xa6, 0xad, 0xd6, 0x98, 0x25, 0xf1, 0xa5, 0xf7,
	0x57, 0xd2, 0xeb, 0xef, 0x6d, 0xac, 0xab, 0xb5, 0x73, 0x23, 0xf0, 0x2d, 0xc2, 0x03, 0x30, 0xf6,
	0x1b, 0xf8, 0x5f, 0x09, 0x3e, 0x5c, 0x4c, 0x5b, 0x07, 0x4b, 0xf0, 0xcf, 0x8c, 0xd7, 0xdf, 0x29,
	0x7e, 0x7e, 0x81, 0x11, 0x54, 0x67, 0x42, 0x64, 0xd2, 0x18, 0xf2, 0x7f, 0xf9, 0x8a, 0x95, 0xc4,
	0x4d, 0xd4, 0x90, 0x9a, 0x83, 0x50, 0x3a, 0x22, 0x5b, 0xa5, 0xb5, 0xd6, 0x78, 0x1f, 0xd5, 0xed,
	0x28, 0xb0, 0xe3, 0x54, 0x92, 0x5a, 0x69, 0xd5, 0xec, 0xe8, 0x7e, 0x9c, 0xca, 0x5e, 0xf0, 0x32,
	0x73, 0x9d, 0xc9, 0xcc, 0x75, 0xde, 0x66, 0xae, 0xf3, 0x34, 0x77, 0x2b, 0x93, 0xb9, 0x5b, 0x79,
	0x9d, 0xbb, 0x95, 0x87, 0xeb, 0x48, 0xd9, 0xc1, 0x30, 0xf4, 0x39, 0x24, 0x94, 0x83, 0x49, 0xc0,
	0x50, 0x15, 0xf2, 0x76, 0x04, 0x34, 0xef, 0xd2, 0x04, 0xc4, 0x30, 0x96, 0xa6, 0x98, 0xd4, 0xd0,
	0xb3, 0x8b, 0xf6, 0x66, 0x95, 0xf6, 0x7a, 0xcd, 0xe2, 0x9a, 0x09, 0x6b, 0xe5, 0x12, 0xe7, 0xef,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0x45, 0x0f, 0xdd, 0x02, 0x02, 0x00, 0x00,
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxType) > 0 {
		i -= len(m.TxType)
		copy(dAtA[i:], m.TxType)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.TxType)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Encoding) > 0 {
		i -= len(m.Encoding)
		copy(dAtA[i:], m.Encoding)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Encoding)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.HostConnectionId) > 0 {
		i -= len(m.HostConnectionId)
		copy(dAtA[i:], m.HostConnectionId)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.HostConnectionId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ControllerConnectionId) > 0 {
		i -= len(m.ControllerConnectionId)
		copy(dAtA[i:], m.ControllerConnectionId)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.ControllerConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.ControllerConnectionId)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.HostConnectionId)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.Encoding)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	l = len(m.TxType)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	return n
}

func sovMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ControllerConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ControllerConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encoding", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Encoding = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
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
				return 0, ErrInvalidLengthMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetadata = fmt.Errorf("proto: unexpected end of group")
)
