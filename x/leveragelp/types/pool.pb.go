// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/leveragelp/pool.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Pool struct {
	AmmPoolId         uint64                                 `protobuf:"varint,1,opt,name=amm_pool_id,json=ammPoolId,proto3" json:"amm_pool_id,omitempty"`
	Health            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=health,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"health"`
	Enabled           bool                                   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Closed            bool                                   `protobuf:"varint,4,opt,name=closed,proto3" json:"closed,omitempty"`
	LeveragedLpAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=leveraged_lp_amount,json=leveragedLpAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"leveraged_lp_amount"`
	LeverageMax       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=leverage_max,json=leverageMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"leverage_max"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3d30fce7b985054, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetAmmPoolId() uint64 {
	if m != nil {
		return m.AmmPoolId
	}
	return 0
}

func (m *Pool) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Pool) GetClosed() bool {
	if m != nil {
		return m.Closed
	}
	return false
}

func init() {
	proto.RegisterType((*Pool)(nil), "elys.leveragelp.Pool")
}

func init() { proto.RegisterFile("elys/leveragelp/pool.proto", fileDescriptor_c3d30fce7b985054) }

var fileDescriptor_c3d30fce7b985054 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x4a, 0xfb, 0x30,
	0x1c, 0xc7, 0x9b, 0xfd, 0xf7, 0xaf, 0x2e, 0x13, 0xc4, 0x28, 0x12, 0x76, 0xc8, 0x86, 0x07, 0xd9,
	0x65, 0xcd, 0xc1, 0x27, 0x70, 0x88, 0x30, 0x50, 0xd0, 0x1e, 0x3d, 0x58, 0xb2, 0x36, 0x74, 0x63,
	0x49, 0x7f, 0x65, 0xc9, 0x74, 0x7b, 0x0b, 0x1f, 0x6b, 0xc7, 0x81, 0x17, 0xf1, 0x30, 0x64, 0x7b,
	0x11, 0x69, 0x6c, 0x75, 0x57, 0x3d, 0x25, 0xdf, 0x7c, 0x92, 0x0f, 0xf9, 0xf1, 0xc5, 0x2d, 0xa9,
	0x16, 0x86, 0x2b, 0xf9, 0x24, 0xa7, 0x22, 0x95, 0x2a, 0xe7, 0x39, 0x80, 0x0a, 0xf2, 0x29, 0x58,
	0x20, 0x87, 0x05, 0x0b, 0x7e, 0x58, 0xeb, 0x24, 0x85, 0x14, 0x1c, 0xe3, 0xc5, 0xee, 0xeb, 0xda,
	0xd9, 0x6b, 0x0d, 0xd7, 0xef, 0x00, 0x14, 0x61, 0xb8, 0x29, 0xb4, 0x8e, 0x0a, 0x43, 0x34, 0x4e,
	0x28, 0xea, 0xa0, 0x6e, 0x3d, 0x6c, 0x08, 0xad, 0x0b, 0x3a, 0x48, 0xc8, 0x35, 0xf6, 0x47, 0x52,
	0x28, 0x3b, 0xa2, 0xb5, 0x0e, 0xea, 0x36, 0xfa, 0xc1, 0x72, 0xdd, 0xf6, 0xde, 0xd7, 0xed, 0xf3,
	0x74, 0x6c, 0x47, 0xb3, 0x61, 0x10, 0x83, 0xe6, 0x31, 0x18, 0x0d, 0xa6, 0x5c, 0x7a, 0x26, 0x99,
	0x70, 0xbb, 0xc8, 0xa5, 0x09, 0xae, 0x64, 0x1c, 0x96, 0xaf, 0x09, 0xc5, 0x7b, 0x32, 0x13, 0x43,
	0x25, 0x13, 0xfa, 0xaf, 0x83, 0xba, 0xfb, 0x61, 0x15, 0xc9, 0x29, 0xf6, 0x63, 0x05, 0x46, 0x26,
	0xb4, 0xee, 0x40, 0x99, 0xc8, 0x23, 0x3e, 0xae, 0xc6, 0x48, 0x22, 0x95, 0x47, 0x42, 0xc3, 0x2c,
	0xb3, 0xf4, 0xff, 0xaf, 0xbf, 0x31, 0xc8, 0x6c, 0x78, 0xf4, 0xad, 0xba, 0xc9, 0x2f, 0x9d, 0x88,
	0xdc, 0xe3, 0x83, 0xea, 0x30, 0xd2, 0x62, 0x4e, 0xfd, 0x3f, 0xcd, 0xd7, 0xac, 0x1c, 0xb7, 0x62,
	0xde, 0x1f, 0x2c, 0x37, 0x0c, 0xad, 0x36, 0x0c, 0x7d, 0x6c, 0x18, 0x7a, 0xd9, 0x32, 0x6f, 0xb5,
	0x65, 0xde, 0xdb, 0x96, 0x79, 0x0f, 0x7c, 0x47, 0x57, 0x34, 0xd4, 0xcb, 0xa4, 0x7d, 0x86, 0xe9,
	0xc4, 0x05, 0x3e, 0xdf, 0x2d, 0xd3, 0xb9, 0x87, 0xbe, 0xeb, 0xe9, 0xe2, 0x33, 0x00, 0x00, 0xff,
	0xff, 0xaa, 0xe7, 0xeb, 0xc7, 0xec, 0x01, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.LeverageMax.Size()
		i -= size
		if _, err := m.LeverageMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.LeveragedLpAmount.Size()
		i -= size
		if _, err := m.LeveragedLpAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Closed {
		i--
		if m.Closed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Health.Size()
		i -= size
		if _, err := m.Health.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AmmPoolId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.AmmPoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AmmPoolId != 0 {
		n += 1 + sovPool(uint64(m.AmmPoolId))
	}
	l = m.Health.Size()
	n += 1 + l + sovPool(uint64(l))
	if m.Enabled {
		n += 2
	}
	if m.Closed {
		n += 2
	}
	l = m.LeveragedLpAmount.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.LeverageMax.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmmPoolId", wireType)
			}
			m.AmmPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AmmPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Health", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Health.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
			m.Enabled = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Closed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
			m.Closed = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeveragedLpAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LeveragedLpAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeverageMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LeverageMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
