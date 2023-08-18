// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fiattokenfactory/blacklister.proto

package types

import (
	fmt "fmt"
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

type Blacklister struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Blacklister) Reset()         { *m = Blacklister{} }
func (m *Blacklister) String() string { return proto.CompactTextString(m) }
func (*Blacklister) ProtoMessage()    {}
func (*Blacklister) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb6b8f9253711167, []int{0}
}
func (m *Blacklister) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Blacklister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Blacklister.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Blacklister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blacklister.Merge(m, src)
}
func (m *Blacklister) XXX_Size() int {
	return m.Size()
}
func (m *Blacklister) XXX_DiscardUnknown() {
	xxx_messageInfo_Blacklister.DiscardUnknown(m)
}

var xxx_messageInfo_Blacklister proto.InternalMessageInfo

func (m *Blacklister) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Blacklister)(nil), "noble.fiattokenfactory.Blacklister")
}

func init() {
	proto.RegisterFile("fiattokenfactory/blacklister.proto", fileDescriptor_eb6b8f9253711167)
}

var fileDescriptor_eb6b8f9253711167 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xce, 0x31, 0x4a, 0xc7, 0x30,
	0x14, 0xc7, 0xf1, 0x66, 0x51, 0xac, 0xdb, 0x7f, 0x90, 0x4e, 0x41, 0xba, 0xe8, 0x92, 0x66, 0xf0,
	0x06, 0x3d, 0x42, 0x47, 0xb7, 0x24, 0x7d, 0xad, 0xa1, 0x31, 0x2f, 0xbc, 0xbc, 0x16, 0x7b, 0x0b,
	0x8f, 0xe5, 0xd8, 0xd1, 0x51, 0xda, 0x8b, 0x08, 0x05, 0x11, 0x74, 0xfc, 0xc1, 0x17, 0x7e, 0x9f,
	0xb2, 0x1e, 0xbc, 0x61, 0xc6, 0x09, 0xe2, 0x60, 0x1c, 0x23, 0xad, 0xda, 0x06, 0xe3, 0xa6, 0xe0,
	0x33, 0x03, 0x35, 0x89, 0x90, 0xf1, 0x72, 0x17, 0xd1, 0x06, 0x68, 0xfe, 0x96, 0xf5, 0x43, 0x79,
	0xdb, 0xfe, 0xc6, 0x97, 0xaa, 0xbc, 0x36, 0x7d, 0x4f, 0x90, 0x73, 0x25, 0xee, 0xc5, 0xe3, 0x4d,
	0xf7, 0x33, 0xdb, 0xf0, 0xb1, 0x4b, 0xb1, 0xed, 0x52, 0x7c, 0xed, 0x52, 0xbc, 0x1f, 0xb2, 0xd8,
	0x0e, 0x59, 0x7c, 0x1e, 0xb2, 0x78, 0xee, 0x46, 0xcf, 0x2f, 0xb3, 0x6d, 0x1c, 0xbe, 0xea, 0xcc,
	0x64, 0xe2, 0x08, 0x01, 0x17, 0x50, 0x0b, 0x44, 0x9e, 0x09, 0xb2, 0x3e, 0xaf, 0x95, 0x73, 0x9c,
	0x14, 0xe1, 0xcc, 0x40, 0x2a, 0x91, 0x5f, 0x0c, 0x83, 0x7e, 0xd3, 0xff, 0xe8, 0xbc, 0x26, 0xc8,
	0xf6, 0xea, 0x54, 0x3f, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x38, 0xb8, 0x45, 0xdb, 0x00,
	0x00, 0x00,
}

func (m *Blacklister) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Blacklister) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Blacklister) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintBlacklister(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlacklister(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlacklister(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Blacklister) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovBlacklister(uint64(l))
	}
	return n
}

func sovBlacklister(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlacklister(x uint64) (n int) {
	return sovBlacklister(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Blacklister) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlacklister
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
			return fmt.Errorf("proto: Blacklister: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Blacklister: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlacklister
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
				return ErrInvalidLengthBlacklister
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlacklister
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlacklister(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlacklister
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
func skipBlacklister(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlacklister
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
					return 0, ErrIntOverflowBlacklister
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
					return 0, ErrIntOverflowBlacklister
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
				return 0, ErrInvalidLengthBlacklister
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlacklister
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlacklister
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlacklister        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlacklister          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlacklister = fmt.Errorf("proto: unexpected end of group")
)
