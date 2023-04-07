// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/oracle/price.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Price struct {
	Asset     string                                 `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	Price     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
	Source    string                                 `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Provider  string                                 `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	Timestamp uint64                                 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c916a7954166b60, []int{0}
}
func (m *Price) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Price.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return m.Size()
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func (m *Price) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *Price) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Price) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Price) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Price)(nil), "elysnetwork.elys.oracle.Price")
}

func init() { proto.RegisterFile("elys/oracle/price.proto", fileDescriptor_0c916a7954166b60) }

var fileDescriptor_0c916a7954166b60 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x37, 0xda, 0x2d, 0x36, 0xc7, 0x50, 0x34, 0x14, 0x49, 0x8b, 0x07, 0x29, 0x48, 0x93,
	0x83, 0x6f, 0x50, 0xea, 0x5d, 0xf6, 0xe8, 0x6d, 0x9b, 0x0e, 0xeb, 0xd2, 0xae, 0x13, 0x92, 0x54,
	0xed, 0x5b, 0xf8, 0x32, 0xbe, 0x43, 0x8f, 0x3d, 0x8a, 0x87, 0x22, 0xbb, 0x2f, 0x22, 0xc9, 0xae,
	0x7f, 0x4e, 0x99, 0xdf, 0xcc, 0x37, 0xc3, 0x97, 0x8f, 0x5e, 0xc0, 0x66, 0xe7, 0x14, 0xda, 0x5c,
	0x6f, 0x40, 0x19, 0x5b, 0x6a, 0x90, 0xc6, 0xa2, 0x47, 0x16, 0x07, 0x4f, 0xe0, 0x5f, 0xd0, 0xae,
	0x65, 0xa8, 0x65, 0x2b, 0x1a, 0x0d, 0x0b, 0x2c, 0x30, 0x6a, 0x54, 0xa8, 0x5a, 0xf9, 0xd5, 0x3b,
	0xa1, 0xe9, 0x7d, 0x58, 0x67, 0x43, 0x9a, 0xe6, 0xce, 0x81, 0xe7, 0x64, 0x42, 0xa6, 0x83, 0xac,
	0x05, 0xb6, 0xa0, 0x69, 0xbc, 0xce, 0x4f, 0x42, 0x77, 0x2e, 0xf7, 0xc7, 0x71, 0xf2, 0x79, 0x1c,
	0x5f, 0x17, 0xa5, 0x7f, 0xdc, 0x2e, 0xa5, 0xc6, 0x4a, 0x69, 0x74, 0x15, 0xba, 0xee, 0x99, 0xb9,
	0xd5, 0x5a, 0xf9, 0x9d, 0x01, 0x27, 0x17, 0xa0, 0xb3, 0x76, 0x99, 0x9d, 0xd3, 0xbe, 0xc3, 0xad,
	0xd5, 0xc0, 0x4f, 0xe3, 0xf1, 0x8e, 0xd8, 0x88, 0x9e, 0x19, 0x8b, 0xcf, 0xe5, 0x0a, 0x2c, 0xef,
	0xc5, 0xc9, 0x2f, 0xb3, 0x4b, 0x3a, 0xf0, 0x65, 0x05, 0xce, 0xe7, 0x95, 0xe1, 0xe9, 0x84, 0x4c,
	0x7b, 0xd9, 0x5f, 0x63, 0x7e, 0xb7, 0xaf, 0x05, 0x39, 0xd4, 0x82, 0x7c, 0xd5, 0x82, 0xbc, 0x35,
	0x22, 0x39, 0x34, 0x22, 0xf9, 0x68, 0x44, 0xf2, 0x70, 0xf3, 0xcf, 0x5a, 0xf8, 0xff, 0xac, 0x0b,
	0x23, 0x82, 0x7a, 0xfd, 0xc9, 0x2c, 0x7a, 0x5c, 0xf6, 0x63, 0x0a, 0xb7, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x0a, 0x48, 0xc1, 0xf8, 0x4f, 0x01, 0x00, 0x00,
}

func (m *Price) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Price) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Price) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintPrice(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintPrice(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Source) > 0 {
		i -= len(m.Source)
		copy(dAtA[i:], m.Source)
		i = encodeVarintPrice(dAtA, i, uint64(len(m.Source)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPrice(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintPrice(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPrice(dAtA []byte, offset int, v uint64) int {
	offset -= sovPrice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Price) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovPrice(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovPrice(uint64(l))
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovPrice(uint64(l))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovPrice(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovPrice(uint64(m.Timestamp))
	}
	return n
}

func sovPrice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPrice(x uint64) (n int) {
	return sovPrice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Price) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrice
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
			return fmt.Errorf("proto: Price: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Price: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
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
				return ErrInvalidLengthPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrice
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
func skipPrice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrice
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
					return 0, ErrIntOverflowPrice
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
					return 0, ErrIntOverflowPrice
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
				return 0, ErrInvalidLengthPrice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPrice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPrice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPrice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPrice = fmt.Errorf("proto: unexpected end of group")
)