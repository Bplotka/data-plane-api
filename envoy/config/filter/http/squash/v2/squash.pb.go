// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: squash.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		squash.proto

	It has these top-level messages:
		Squash
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

// [#proto-status: experimental]
type Squash struct {
	// The name of the cluster that hosts the Squash server.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// When the filter requests the Squash server to create a DebugAttachment, it will use this
	// structure as template for the body of the request. It can contain reference to environment
	// variables in the form of '{{ ENV_VAR_NAME }}'. These can be used to provide the Squash server
	// with more information to find the process to attach the debugger to. For example, in a
	// Istio/k8s environment, this will contain information on the pod:
	//
	// .. code-block:: json
	//
	//  {
	//    "spec": {
	//      "attachment": {
	//        "pod": "{{ POD_NAME }}",
	//        "namespace": "{{ POD_NAMESPACE }}"
	//      },
	//      "match_request": true
	//    }
	//  }
	//
	// (where POD_NAME, POD_NAMESPACE are configured in the pod via the Downward API)
	AttachmentTemplate *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=attachment_template,json=attachmentTemplate" json:"attachment_template,omitempty"`
	// The timeout for individual requests sent to the Squash cluster. Defaults to 1 second.
	RequestTimeout *time.Duration `protobuf:"bytes,3,opt,name=request_timeout,json=requestTimeout,stdduration" json:"request_timeout,omitempty"`
	// The total timeout Squash will delay a request and wait for it to be attached. Defaults to 60
	// seconds.
	AttachmentTimeout *time.Duration `protobuf:"bytes,4,opt,name=attachment_timeout,json=attachmentTimeout,stdduration" json:"attachment_timeout,omitempty"`
	// Amount of time to poll for the status of the attachment object in the Squash server
	// (to check if has been attached). Defaults to 1 second.
	AttachmentPollPeriod *time.Duration `protobuf:"bytes,5,opt,name=attachment_poll_period,json=attachmentPollPeriod,stdduration" json:"attachment_poll_period,omitempty"`
}

func (m *Squash) Reset()                    { *m = Squash{} }
func (m *Squash) String() string            { return proto.CompactTextString(m) }
func (*Squash) ProtoMessage()               {}
func (*Squash) Descriptor() ([]byte, []int) { return fileDescriptorSquash, []int{0} }

func (m *Squash) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Squash) GetAttachmentTemplate() *google_protobuf1.Struct {
	if m != nil {
		return m.AttachmentTemplate
	}
	return nil
}

func (m *Squash) GetRequestTimeout() *time.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentTimeout() *time.Duration {
	if m != nil {
		return m.AttachmentTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentPollPeriod() *time.Duration {
	if m != nil {
		return m.AttachmentPollPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*Squash)(nil), "envoy.config.filter.http.squash.v2.Squash")
}
func (m *Squash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Squash) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Cluster) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSquash(dAtA, i, uint64(len(m.Cluster)))
		i += copy(dAtA[i:], m.Cluster)
	}
	if m.AttachmentTemplate != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSquash(dAtA, i, uint64(m.AttachmentTemplate.Size()))
		n1, err := m.AttachmentTemplate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.RequestTimeout != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSquash(dAtA, i, uint64(types.SizeOfStdDuration(*m.RequestTimeout)))
		n2, err := types.StdDurationMarshalTo(*m.RequestTimeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.AttachmentTimeout != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSquash(dAtA, i, uint64(types.SizeOfStdDuration(*m.AttachmentTimeout)))
		n3, err := types.StdDurationMarshalTo(*m.AttachmentTimeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.AttachmentPollPeriod != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSquash(dAtA, i, uint64(types.SizeOfStdDuration(*m.AttachmentPollPeriod)))
		n4, err := types.StdDurationMarshalTo(*m.AttachmentPollPeriod, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeVarintSquash(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Squash) Size() (n int) {
	var l int
	_ = l
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentTemplate != nil {
		l = m.AttachmentTemplate.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.RequestTimeout != nil {
		l = types.SizeOfStdDuration(*m.RequestTimeout)
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentTimeout != nil {
		l = types.SizeOfStdDuration(*m.AttachmentTimeout)
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentPollPeriod != nil {
		l = types.SizeOfStdDuration(*m.AttachmentPollPeriod)
		n += 1 + l + sovSquash(uint64(l))
	}
	return n
}

func sovSquash(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSquash(x uint64) (n int) {
	return sovSquash(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Squash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSquash
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
			return fmt.Errorf("proto: Squash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Squash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentTemplate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentTemplate == nil {
				m.AttachmentTemplate = &google_protobuf1.Struct{}
			}
			if err := m.AttachmentTemplate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RequestTimeout == nil {
				m.RequestTimeout = new(time.Duration)
			}
			if err := types.StdDurationUnmarshal(m.RequestTimeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentTimeout == nil {
				m.AttachmentTimeout = new(time.Duration)
			}
			if err := types.StdDurationUnmarshal(m.AttachmentTimeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentPollPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentPollPeriod == nil {
				m.AttachmentPollPeriod = new(time.Duration)
			}
			if err := types.StdDurationUnmarshal(m.AttachmentPollPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSquash(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSquash
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
func skipSquash(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSquash
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
					return 0, ErrIntOverflowSquash
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
					return 0, ErrIntOverflowSquash
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
				return 0, ErrInvalidLengthSquash
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSquash
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
				next, err := skipSquash(dAtA[start:])
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
	ErrInvalidLengthSquash = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSquash   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("squash.proto", fileDescriptorSquash) }

var fileDescriptorSquash = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4e, 0x2a, 0x31,
	0x18, 0x85, 0xd3, 0x81, 0xcb, 0x0d, 0xbd, 0x37, 0xf7, 0x6a, 0x25, 0x32, 0x12, 0x33, 0x12, 0xdc,
	0xb0, 0xea, 0x24, 0xf8, 0x06, 0xc4, 0x05, 0x2b, 0x43, 0x00, 0x37, 0x6e, 0x48, 0x19, 0xca, 0x30,
	0x49, 0x99, 0x7f, 0xe8, 0xfc, 0x9d, 0xc4, 0x37, 0xf1, 0x59, 0x5c, 0xb9, 0x74, 0xe9, 0x1b, 0x68,
	0x58, 0xe9, 0x5b, 0x18, 0x3b, 0x25, 0x10, 0xdd, 0xb0, 0x3b, 0xe9, 0xe9, 0xf7, 0xf5, 0xa4, 0xf4,
	0x6f, 0xbe, 0x36, 0x22, 0x5f, 0xf2, 0x4c, 0x03, 0x02, 0xeb, 0xc8, 0xb4, 0x80, 0x7b, 0x1e, 0x41,
	0xba, 0x48, 0x62, 0xbe, 0x48, 0x14, 0x4a, 0xcd, 0x97, 0x88, 0x19, 0x77, 0xd7, 0x8a, 0x5e, 0x2b,
	0x88, 0x01, 0x62, 0x25, 0x43, 0x4b, 0xcc, 0xcc, 0x22, 0x9c, 0x1b, 0x2d, 0x30, 0x81, 0xb4, 0x74,
	0xb4, 0xce, 0xbf, 0xf7, 0x39, 0x6a, 0x13, 0xa1, 0x6b, 0x9b, 0x85, 0x50, 0xc9, 0x5c, 0xa0, 0x0c,
	0xb7, 0xc1, 0x15, 0x8d, 0x18, 0x62, 0xb0, 0x31, 0xfc, 0x4a, 0xe5, 0x69, 0xe7, 0xdd, 0xa3, 0xb5,
	0xb1, 0x7d, 0x9a, 0x5d, 0xd2, 0xdf, 0x91, 0x32, 0x39, 0x4a, 0xed, 0x93, 0x36, 0xe9, 0xd6, 0xfb,
	0xf5, 0xc7, 0x8f, 0xa7, 0x4a, 0x55, 0x7b, 0x6d, 0x32, 0xda, 0x36, 0x6c, 0x40, 0x4f, 0x04, 0xa2,
	0x88, 0x96, 0x2b, 0x99, 0xe2, 0x14, 0xe5, 0x2a, 0x53, 0x02, 0xa5, 0xef, 0xb5, 0x49, 0xf7, 0x4f,
	0xaf, 0xc9, 0xcb, 0x69, 0x7c, 0x3b, 0x8d, 0x8f, 0xed, 0xb4, 0x11, 0xdb, 0x31, 0x13, 0x87, 0xb0,
	0x01, 0xfd, 0xaf, 0xe5, 0xda, 0xc8, 0x1c, 0xa7, 0x98, 0xac, 0x24, 0x18, 0xf4, 0x2b, 0xd6, 0x72,
	0xf6, 0xc3, 0x72, 0xed, 0x3e, 0xa0, 0x5f, 0x7d, 0x78, 0xbd, 0x20, 0xa3, 0x7f, 0x8e, 0x9b, 0x94,
	0x18, 0xbb, 0xa1, 0x6c, 0x7f, 0x93, 0x93, 0x55, 0x0f, 0x93, 0x1d, 0xef, 0x4d, 0x73, 0xbe, 0x5b,
	0x7a, 0xba, 0xe7, 0xcb, 0x40, 0xa9, 0x69, 0x26, 0x75, 0x02, 0x73, 0xff, 0xd7, 0x61, 0xce, 0xc6,
	0x0e, 0x1f, 0x82, 0x52, 0x43, 0x0b, 0xf7, 0x8f, 0x9e, 0x37, 0x01, 0x79, 0xd9, 0x04, 0xe4, 0x6d,
	0x13, 0x90, 0x3b, 0xaf, 0xe8, 0xcd, 0x6a, 0x56, 0x70, 0xf5, 0x19, 0x00, 0x00, 0xff, 0xff, 0x85,
	0x24, 0x3d, 0x10, 0x24, 0x02, 0x00, 0x00,
}