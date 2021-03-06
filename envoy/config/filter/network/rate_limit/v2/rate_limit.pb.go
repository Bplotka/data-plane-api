// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rate_limit.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		rate_limit.proto

	It has these top-level messages:
		RateLimit
*/
package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_ratelimit "envoy/api/v2/ratelimit"

import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RateLimit struct {
	// The prefix to use when emitting :ref:`statistics <config_network_filters_rate_limit_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The rate limit domain to use in the rate limit service request.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// The rate limit descriptor list to use in the rate limit service request.
	Descriptors []*envoy_api_v2_ratelimit.RateLimitDescriptor `protobuf:"bytes,3,rep,name=descriptors" json:"descriptors,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *time.Duration `protobuf:"bytes,4,opt,name=timeout,stdduration" json:"timeout,omitempty"`
}

func (m *RateLimit) Reset()                    { *m = RateLimit{} }
func (m *RateLimit) String() string            { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()               {}
func (*RateLimit) Descriptor() ([]byte, []int) { return fileDescriptorRateLimit, []int{0} }

func (m *RateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetDescriptors() []*envoy_api_v2_ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimit) GetTimeout() *time.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.network.rate_limit.v2.RateLimit")
}
func (m *RateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimit) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StatPrefix) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.StatPrefix)))
		i += copy(dAtA[i:], m.StatPrefix)
	}
	if len(m.Domain) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.Domain)))
		i += copy(dAtA[i:], m.Domain)
	}
	if len(m.Descriptors) > 0 {
		for _, msg := range m.Descriptors {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRateLimit(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Timeout != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(types.SizeOfStdDuration(*m.Timeout)))
		n1, err := types.StdDurationMarshalTo(*m.Timeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintRateLimit(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RateLimit) Size() (n int) {
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if len(m.Descriptors) > 0 {
		for _, e := range m.Descriptors {
			l = e.Size()
			n += 1 + l + sovRateLimit(uint64(l))
		}
	}
	if m.Timeout != nil {
		l = types.SizeOfStdDuration(*m.Timeout)
		n += 1 + l + sovRateLimit(uint64(l))
	}
	return n
}

func sovRateLimit(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRateLimit(x uint64) (n int) {
	return sovRateLimit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRateLimit
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Descriptors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Descriptors = append(m.Descriptors, &envoy_api_v2_ratelimit.RateLimitDescriptor{})
			if err := m.Descriptors[len(m.Descriptors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timeout == nil {
				m.Timeout = new(time.Duration)
			}
			if err := types.StdDurationUnmarshal(m.Timeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRateLimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRateLimit
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
func skipRateLimit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRateLimit
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
					return 0, ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRateLimit
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRateLimit
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRateLimit
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRateLimit(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRateLimit = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRateLimit   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("rate_limit.proto", fileDescriptorRateLimit) }

var fileDescriptorRateLimit = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbd, 0x4e, 0xfb, 0x30,
	0x14, 0xc5, 0xe5, 0xb4, 0xff, 0xfe, 0xa9, 0xb3, 0x54, 0x11, 0x12, 0xa1, 0x43, 0x08, 0x0c, 0xa8,
	0x80, 0x64, 0x4b, 0x61, 0x62, 0x8d, 0x3a, 0x32, 0xa0, 0x6c, 0xb0, 0x54, 0x6e, 0xe3, 0x44, 0x57,
	0xa4, 0xb9, 0x91, 0x7b, 0x1b, 0xe0, 0x35, 0x98, 0x78, 0x16, 0x26, 0x46, 0x46, 0xde, 0x00, 0xd4,
	0x8d, 0x8d, 0x47, 0x40, 0xf9, 0xe8, 0x87, 0xd8, 0x8e, 0xfd, 0x3b, 0xf7, 0xd8, 0xf7, 0xf0, 0x81,
	0x51, 0xa4, 0x27, 0x19, 0xcc, 0x81, 0x44, 0x61, 0x90, 0xd0, 0x39, 0xd3, 0x79, 0x89, 0x4f, 0x62,
	0x86, 0x79, 0x02, 0xa9, 0x48, 0x20, 0x23, 0x6d, 0x44, 0xae, 0xe9, 0x01, 0xcd, 0xbd, 0xd8, 0x71,
	0x97, 0xc1, 0xf0, 0xb4, 0xb6, 0x4a, 0x55, 0x80, 0x2c, 0x03, 0x59, 0xb1, 0x1a, 0x6d, 0x55, 0x13,
	0x39, 0xf4, 0x52, 0xc4, 0x34, 0xd3, 0xb2, 0x3e, 0x4d, 0x97, 0x89, 0x8c, 0x97, 0x46, 0x11, 0x60,
	0xde, 0xf2, 0x83, 0x52, 0x65, 0x10, 0x2b, 0xd2, 0x72, 0x2d, 0x5a, 0xb0, 0x9f, 0x62, 0x8a, 0xb5,
	0x94, 0x95, 0x6a, 0x6e, 0x4f, 0x7e, 0x18, 0xef, 0x47, 0x8a, 0xf4, 0x75, 0xf5, 0x84, 0x73, 0xce,
	0xed, 0x05, 0x29, 0x9a, 0x14, 0x46, 0x27, 0xf0, 0xe8, 0x32, 0x9f, 0x8d, 0xfa, 0x61, 0xff, 0xf5,
	0xfb, 0xad, 0xd3, 0x35, 0x96, 0xcf, 0x22, 0x5e, 0xd1, 0x9b, 0x1a, 0x3a, 0xc7, 0xbc, 0x17, 0xe3,
	0x5c, 0x41, 0xee, 0x5a, 0x7f, 0x6d, 0x2d, 0x70, 0x6e, 0xb9, 0x1d, 0xeb, 0xc5, 0xcc, 0x40, 0x41,
	0x68, 0x16, 0x6e, 0xc7, 0xef, 0x8c, 0xec, 0xe0, 0x42, 0x34, 0xa5, 0xa8, 0x02, 0x44, 0x19, 0x88,
	0xed, 0x7e, 0x9b, 0x6f, 0x8c, 0x37, 0x33, 0x21, 0xaf, 0x42, 0xff, 0x3d, 0x33, 0x6b, 0x8f, 0x45,
	0xbb, 0x59, 0xce, 0x15, 0xff, 0x4f, 0x30, 0xd7, 0xb8, 0x24, 0xb7, 0xeb, 0xb3, 0x91, 0x1d, 0x1c,
	0x8a, 0xa6, 0x18, 0xb1, 0x2e, 0x46, 0x8c, 0xdb, 0x62, 0xc2, 0xee, 0xcb, 0xe7, 0x11, 0x8b, 0xd6,
	0xfe, 0x70, 0xf0, 0xbe, 0xf2, 0xd8, 0xc7, 0xca, 0x63, 0x5f, 0x2b, 0x8f, 0xdd, 0x59, 0x65, 0x30,
	0xed, 0xd5, 0x33, 0x97, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x8b, 0x20, 0xb5, 0xc1, 0x01,
	0x00, 0x00,
}
