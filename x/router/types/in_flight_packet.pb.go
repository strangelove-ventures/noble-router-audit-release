// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: router/in_flight_packet.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

// InFlightPacket contains information about the initially minted funds
// @param source_domain_sender
// @param nonce
type InFlightPacket struct {
	SourceDomainSender string `protobuf:"bytes,1,opt,name=source_domain_sender,json=sourceDomainSender,proto3" json:"source_domain_sender,omitempty"`
	Nonce              uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ChannelId          string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	PortId             string `protobuf:"bytes,4,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	Sequence           uint64 `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *InFlightPacket) Reset()         { *m = InFlightPacket{} }
func (m *InFlightPacket) String() string { return proto.CompactTextString(m) }
func (*InFlightPacket) ProtoMessage()    {}
func (*InFlightPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_93456951126717f2, []int{0}
}
func (m *InFlightPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InFlightPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InFlightPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InFlightPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InFlightPacket.Merge(m, src)
}
func (m *InFlightPacket) XXX_Size() int {
	return m.Size()
}
func (m *InFlightPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_InFlightPacket.DiscardUnknown(m)
}

var xxx_messageInfo_InFlightPacket proto.InternalMessageInfo

func (m *InFlightPacket) GetSourceDomainSender() string {
	if m != nil {
		return m.SourceDomainSender
	}
	return ""
}

func (m *InFlightPacket) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *InFlightPacket) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *InFlightPacket) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *InFlightPacket) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*InFlightPacket)(nil), "noble.router.InFlightPacket")
}

func init() { proto.RegisterFile("router/in_flight_packet.proto", fileDescriptor_93456951126717f2) }

var fileDescriptor_93456951126717f2 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x31, 0x4f, 0x02, 0x41,
	0x10, 0x85, 0x59, 0x05, 0x94, 0x8d, 0xb1, 0xd8, 0x90, 0x48, 0x48, 0xd8, 0x10, 0x2b, 0x1a, 0x58,
	0x89, 0xa5, 0x9d, 0x31, 0x26, 0x74, 0x06, 0x63, 0x63, 0x73, 0xd9, 0xdb, 0x1b, 0x8f, 0x8d, 0xc7,
	0xcc, 0xb9, 0xbb, 0x47, 0xf4, 0x5f, 0xf8, 0x43, 0xfc, 0x21, 0x96, 0x94, 0x96, 0x06, 0xfe, 0x88,
	0xb9, 0x3d, 0x62, 0xf7, 0xde, 0x7c, 0x6f, 0x5e, 0xf1, 0xf8, 0xc8, 0x51, 0x15, 0xc0, 0x29, 0x8b,
	0xc9, 0x4b, 0x61, 0xf3, 0x55, 0x48, 0x4a, 0x6d, 0x5e, 0x21, 0xcc, 0x4a, 0x47, 0x81, 0xc4, 0x19,
	0x52, 0x5a, 0xc0, 0xac, 0x09, 0x0d, 0xa5, 0x21, 0xbf, 0x26, 0xaf, 0x52, 0xed, 0x41, 0x6d, 0xe6,
	0x29, 0x04, 0x3d, 0x57, 0x86, 0x2c, 0x36, 0xe9, 0x61, 0x3f, 0xa7, 0x9c, 0xa2, 0x54, 0xb5, 0x6a,
	0xae, 0x97, 0x5f, 0x8c, 0x9f, 0x2f, 0xf0, 0x3e, 0xb6, 0x3f, 0xc4, 0x72, 0x71, 0xc5, 0xfb, 0x9e,
	0x2a, 0x67, 0x20, 0xc9, 0x68, 0xad, 0x2d, 0x26, 0x1e, 0x30, 0x03, 0x37, 0x60, 0x63, 0x36, 0xe9,
	0x2d, 0x45, 0xc3, 0xee, 0x22, 0x7a, 0x8c, 0x44, 0xf4, 0x79, 0x07, 0x09, 0x0d, 0x0c, 0x8e, 0xc6,
	0x6c, 0xd2, 0x5e, 0x36, 0x46, 0x8c, 0x38, 0x37, 0x2b, 0x8d, 0x08, 0x45, 0x62, 0xb3, 0xc1, 0x71,
	0xfc, 0xee, 0x1d, 0x2e, 0x8b, 0x4c, 0x5c, 0xf0, 0x93, 0x92, 0x5c, 0xa8, 0x59, 0x3b, 0xb2, 0x6e,
	0x6d, 0x17, 0x99, 0x18, 0xf2, 0x53, 0x0f, 0x6f, 0x15, 0xd4, 0x85, 0x9d, 0x58, 0xf8, 0xef, 0x6f,
	0x9f, 0xbe, 0x77, 0x92, 0x6d, 0x77, 0x92, 0xfd, 0xee, 0x24, 0xfb, 0xdc, 0xcb, 0xd6, 0x76, 0x2f,
	0x5b, 0x3f, 0x7b, 0xd9, 0x7a, 0xbe, 0xc9, 0x6d, 0x58, 0x55, 0xe9, 0xcc, 0xd0, 0x5a, 0xf9, 0xe0,
	0x34, 0xe6, 0x50, 0xd0, 0x06, 0xa6, 0x1b, 0xc0, 0x50, 0x39, 0xf0, 0x2a, 0x8e, 0x35, 0x3d, 0x2c,
	0xfa, 0xae, 0x0e, 0x22, 0x7c, 0x94, 0xe0, 0xd3, 0x6e, 0x1c, 0xe3, 0xfa, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0xdc, 0x89, 0x08, 0x81, 0x71, 0x01, 0x00, 0x00,
}

func (m *InFlightPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InFlightPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InFlightPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintInFlightPacket(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintInFlightPacket(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintInFlightPacket(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Nonce != 0 {
		i = encodeVarintInFlightPacket(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x10
	}
	if len(m.SourceDomainSender) > 0 {
		i -= len(m.SourceDomainSender)
		copy(dAtA[i:], m.SourceDomainSender)
		i = encodeVarintInFlightPacket(dAtA, i, uint64(len(m.SourceDomainSender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInFlightPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovInFlightPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InFlightPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourceDomainSender)
	if l > 0 {
		n += 1 + l + sovInFlightPacket(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovInFlightPacket(uint64(m.Nonce))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovInFlightPacket(uint64(l))
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovInFlightPacket(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovInFlightPacket(uint64(m.Sequence))
	}
	return n
}

func sovInFlightPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInFlightPacket(x uint64) (n int) {
	return sovInFlightPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InFlightPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInFlightPacket
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
			return fmt.Errorf("proto: InFlightPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InFlightPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceDomainSender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInFlightPacket
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
				return ErrInvalidLengthInFlightPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInFlightPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceDomainSender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInFlightPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInFlightPacket
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
				return ErrInvalidLengthInFlightPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInFlightPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInFlightPacket
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
				return ErrInvalidLengthInFlightPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInFlightPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInFlightPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInFlightPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInFlightPacket
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
func skipInFlightPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInFlightPacket
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
					return 0, ErrIntOverflowInFlightPacket
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
					return 0, ErrIntOverflowInFlightPacket
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
				return 0, ErrInvalidLengthInFlightPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInFlightPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInFlightPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInFlightPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInFlightPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInFlightPacket = fmt.Errorf("proto: unexpected end of group")
)
