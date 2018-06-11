// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: buffer.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		buffer.proto

	It has these top-level messages:
		Buffer
		BufferPerRoute
*/
package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import google_protobuf1 "google/protobuf"
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

type Buffer struct {
	// The maximum request size that the filter will buffer before the connection
	// manager will stop buffering and return a 413 response.
	MaxRequestBytes *google_protobuf1.UInt32Value `protobuf:"bytes,1,opt,name=max_request_bytes,json=maxRequestBytes" json:"max_request_bytes,omitempty"`
	// The maximum number of seconds that the filter will wait for a complete
	// request before returning a 408 response.
	MaxRequestTime *time.Duration `protobuf:"bytes,2,opt,name=max_request_time,json=maxRequestTime,stdduration" json:"max_request_time,omitempty"`
}

func (m *Buffer) Reset()                    { *m = Buffer{} }
func (m *Buffer) String() string            { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()               {}
func (*Buffer) Descriptor() ([]byte, []int) { return fileDescriptorBuffer, []int{0} }

func (m *Buffer) GetMaxRequestBytes() *google_protobuf1.UInt32Value {
	if m != nil {
		return m.MaxRequestBytes
	}
	return nil
}

func (m *Buffer) GetMaxRequestTime() *time.Duration {
	if m != nil {
		return m.MaxRequestTime
	}
	return nil
}

type BufferPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*BufferPerRoute_Disabled
	//	*BufferPerRoute_Buffer
	Override isBufferPerRoute_Override `protobuf_oneof:"override"`
}

func (m *BufferPerRoute) Reset()                    { *m = BufferPerRoute{} }
func (m *BufferPerRoute) String() string            { return proto.CompactTextString(m) }
func (*BufferPerRoute) ProtoMessage()               {}
func (*BufferPerRoute) Descriptor() ([]byte, []int) { return fileDescriptorBuffer, []int{1} }

type isBufferPerRoute_Override interface {
	isBufferPerRoute_Override()
	MarshalTo([]byte) (int, error)
	Size() int
}

type BufferPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}
type BufferPerRoute_Buffer struct {
	Buffer *Buffer `protobuf:"bytes,2,opt,name=buffer,oneof"`
}

func (*BufferPerRoute_Disabled) isBufferPerRoute_Override() {}
func (*BufferPerRoute_Buffer) isBufferPerRoute_Override()   {}

func (m *BufferPerRoute) GetOverride() isBufferPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *BufferPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*BufferPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *BufferPerRoute) GetBuffer() *Buffer {
	if x, ok := m.GetOverride().(*BufferPerRoute_Buffer); ok {
		return x.Buffer
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BufferPerRoute) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BufferPerRoute_OneofMarshaler, _BufferPerRoute_OneofUnmarshaler, _BufferPerRoute_OneofSizer, []interface{}{
		(*BufferPerRoute_Disabled)(nil),
		(*BufferPerRoute_Buffer)(nil),
	}
}

func _BufferPerRoute_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BufferPerRoute)
	// override
	switch x := m.Override.(type) {
	case *BufferPerRoute_Disabled:
		t := uint64(0)
		if x.Disabled {
			t = 1
		}
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *BufferPerRoute_Buffer:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Buffer); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BufferPerRoute.Override has unexpected type %T", x)
	}
	return nil
}

func _BufferPerRoute_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BufferPerRoute)
	switch tag {
	case 1: // override.disabled
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Override = &BufferPerRoute_Disabled{x != 0}
		return true, err
	case 2: // override.buffer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Buffer)
		err := b.DecodeMessage(msg)
		m.Override = &BufferPerRoute_Buffer{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BufferPerRoute_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BufferPerRoute)
	// override
	switch x := m.Override.(type) {
	case *BufferPerRoute_Disabled:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += 1
	case *BufferPerRoute_Buffer:
		s := proto.Size(x.Buffer)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Buffer)(nil), "envoy.config.filter.http.buffer.v2.Buffer")
	proto.RegisterType((*BufferPerRoute)(nil), "envoy.config.filter.http.buffer.v2.BufferPerRoute")
}
func (m *Buffer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Buffer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MaxRequestBytes != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBuffer(dAtA, i, uint64(m.MaxRequestBytes.Size()))
		n1, err := m.MaxRequestBytes.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.MaxRequestTime != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBuffer(dAtA, i, uint64(types.SizeOfStdDuration(*m.MaxRequestTime)))
		n2, err := types.StdDurationMarshalTo(*m.MaxRequestTime, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *BufferPerRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BufferPerRoute) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Override != nil {
		nn3, err := m.Override.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn3
	}
	return i, nil
}

func (m *BufferPerRoute_Disabled) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x8
	i++
	if m.Disabled {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func (m *BufferPerRoute_Buffer) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Buffer != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBuffer(dAtA, i, uint64(m.Buffer.Size()))
		n4, err := m.Buffer.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func encodeVarintBuffer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Buffer) Size() (n int) {
	var l int
	_ = l
	if m.MaxRequestBytes != nil {
		l = m.MaxRequestBytes.Size()
		n += 1 + l + sovBuffer(uint64(l))
	}
	if m.MaxRequestTime != nil {
		l = types.SizeOfStdDuration(*m.MaxRequestTime)
		n += 1 + l + sovBuffer(uint64(l))
	}
	return n
}

func (m *BufferPerRoute) Size() (n int) {
	var l int
	_ = l
	if m.Override != nil {
		n += m.Override.Size()
	}
	return n
}

func (m *BufferPerRoute_Disabled) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}
func (m *BufferPerRoute_Buffer) Size() (n int) {
	var l int
	_ = l
	if m.Buffer != nil {
		l = m.Buffer.Size()
		n += 1 + l + sovBuffer(uint64(l))
	}
	return n
}

func sovBuffer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBuffer(x uint64) (n int) {
	return sovBuffer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Buffer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuffer
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
			return fmt.Errorf("proto: Buffer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Buffer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRequestBytes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
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
				return ErrInvalidLengthBuffer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxRequestBytes == nil {
				m.MaxRequestBytes = &google_protobuf1.UInt32Value{}
			}
			if err := m.MaxRequestBytes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRequestTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
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
				return ErrInvalidLengthBuffer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxRequestTime == nil {
				m.MaxRequestTime = new(time.Duration)
			}
			if err := types.StdDurationUnmarshal(m.MaxRequestTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBuffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBuffer
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
func (m *BufferPerRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuffer
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
			return fmt.Errorf("proto: BufferPerRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BufferPerRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Override = &BufferPerRoute_Disabled{b}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buffer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuffer
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
				return ErrInvalidLengthBuffer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Buffer{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Override = &BufferPerRoute_Buffer{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBuffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBuffer
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
func skipBuffer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBuffer
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
					return 0, ErrIntOverflowBuffer
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
					return 0, ErrIntOverflowBuffer
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
				return 0, ErrInvalidLengthBuffer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBuffer
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
				next, err := skipBuffer(dAtA[start:])
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
	ErrInvalidLengthBuffer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBuffer   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("buffer.proto", fileDescriptorBuffer) }

var fileDescriptorBuffer = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3f, 0x4e, 0xe3, 0x40,
	0x18, 0xc5, 0x33, 0xce, 0x9f, 0xf5, 0xce, 0xae, 0xb2, 0x8e, 0xb5, 0xd2, 0x66, 0x57, 0x2b, 0x13,
	0xa5, 0x01, 0xa5, 0x18, 0x4b, 0xce, 0x0d, 0x2c, 0x0a, 0xe8, 0x22, 0x43, 0x28, 0x68, 0xa2, 0x31,
	0xfe, 0x6c, 0x06, 0xd9, 0x1e, 0x33, 0x1e, 0x9b, 0xe4, 0x0a, 0x9c, 0x80, 0x8a, 0x03, 0x50, 0x53,
	0x20, 0xaa, 0x94, 0x94, 0xdc, 0x00, 0x94, 0x2e, 0xb7, 0x40, 0xfe, 0x13, 0x40, 0x4a, 0x41, 0xf7,
	0x34, 0x6f, 0xbe, 0xdf, 0x7b, 0xdf, 0x0c, 0xfe, 0xe9, 0x66, 0xbe, 0x0f, 0x82, 0x24, 0x82, 0x4b,
	0xae, 0x0f, 0x21, 0xce, 0xf9, 0x82, 0x9c, 0xf1, 0xd8, 0x67, 0x01, 0xf1, 0x59, 0x28, 0x41, 0x90,
	0x73, 0x29, 0x13, 0x52, 0x5f, 0xcb, 0xad, 0x7f, 0x46, 0xc0, 0x79, 0x10, 0x82, 0x59, 0x4e, 0xb8,
	0x99, 0x6f, 0x7a, 0x99, 0xa0, 0x92, 0xf1, 0xb8, 0x62, 0x6c, 0xfb, 0x57, 0x82, 0x26, 0x09, 0x88,
	0xb4, 0xf6, 0xff, 0xe4, 0x34, 0x64, 0x1e, 0x95, 0x60, 0x6e, 0x44, 0x6d, 0xfc, 0x0e, 0x78, 0xc0,
	0x4b, 0x69, 0x16, 0xaa, 0x3a, 0x1d, 0xde, 0x23, 0xdc, 0xb1, 0xcb, 0x70, 0xfd, 0x08, 0xf7, 0x22,
	0x3a, 0x9f, 0x09, 0xb8, 0xcc, 0x20, 0x95, 0x33, 0x77, 0x21, 0x21, 0xed, 0xa3, 0x01, 0xda, 0xfb,
	0x61, 0xfd, 0x27, 0x55, 0x2a, 0xd9, 0xa4, 0x92, 0xe9, 0x61, 0x2c, 0xc7, 0xd6, 0x09, 0x0d, 0x33,
	0xb0, 0xbf, 0x3f, 0xae, 0x97, 0xcd, 0xd6, 0x48, 0x19, 0x34, 0x9c, 0x5f, 0x11, 0x9d, 0x3b, 0x15,
	0xc0, 0x2e, 0xe6, 0xf5, 0x29, 0xd6, 0x3e, 0x43, 0x25, 0x8b, 0xa0, 0xaf, 0x94, 0xcc, 0xbf, 0x5b,
	0xcc, 0xfd, 0x7a, 0x53, 0x5b, 0xbb, 0x79, 0xd9, 0x41, 0x05, 0xf4, 0xdb, 0x1d, 0x6a, 0xa9, 0x68,
	0xd4, 0x70, 0xba, 0x1f, 0xdc, 0x63, 0x16, 0xc1, 0xf0, 0x16, 0xe1, 0x6e, 0x55, 0x7b, 0x02, 0xc2,
	0xe1, 0x99, 0x04, 0x7d, 0x17, 0xab, 0x1e, 0x4b, 0xa9, 0x1b, 0x82, 0x57, 0xb6, 0x56, 0xeb, 0x5e,
	0x17, 0x8a, 0x8a, 0x0e, 0x1a, 0xce, 0xbb, 0xa9, 0x4f, 0x70, 0xa7, 0x7a, 0xee, 0xba, 0xc8, 0x88,
	0x7c, 0xfd, 0x2d, 0xa4, 0x0a, 0xb3, 0x71, 0x81, 0x6c, 0x5f, 0x23, 0x45, 0x2b, 0x98, 0x35, 0xc7,
	0xee, 0x61, 0x95, 0xe7, 0x20, 0x04, 0xf3, 0x40, 0x6f, 0x3f, 0xac, 0x97, 0x4d, 0x64, 0x6b, 0x4f,
	0x2b, 0x03, 0x3d, 0xaf, 0x0c, 0xf4, 0xba, 0x32, 0xd0, 0xa9, 0x92, 0x5b, 0x6e, 0xa7, 0xdc, 0x73,
	0xfc, 0x16, 0x00, 0x00, 0xff, 0xff, 0x41, 0x92, 0xd6, 0xdf, 0x13, 0x02, 0x00, 0x00,
}